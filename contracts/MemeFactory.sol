// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/ReentrancyGuardUpgradeable.sol";
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import {MemeToken} from "./MemeToken.sol";
import {BondingCurve} from "./BondingCurve.sol";

contract MemeFactory is
    UUPSUpgradeable,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable,
    BondingCurve
{
    enum TokenState {
        NOT_CREATED, // 0
        FUNDING, // 1
        WAIT_TRADING, // 2
        TRADING // 3
    }

    uint256 public constant FEE_DENOMINATOR = 10000;

    // bonding curve parameters
    uint256 public K;
    uint256 public x;
    uint256 public y;
    mapping(address => uint256) public tokenX; // x value for each token
    mapping(address => uint256) public tokenY; // y value for each token

    uint256 public fundingGoal;
    mapping(address => TokenState) public tokens;
    mapping(address => uint256) public collateral; // total collateral for each token
    uint256 public feePercent; // bp
    uint256 public fee;
    uint256 public createFee;

    /// events
    event TokenCreated(
        address indexed token,
        address indexed creator,
        uint256 timestamp,
        string name,
        string symbol,
        string description,
        string icon,
        string website,
        string twitter,
        string telegram,
        string discord,
        string github
    );

    /// @notice when user buy token
    /// @param token token address
    /// @param timestamp timestamp
    /// @param baseAmount meme token amount
    /// @param quoteAmount eth quote amount
    /// @param feeAmount fee amount
    event TokenBuy(
        address indexed token,
        uint256 timestamp,
        uint256 baseAmount,
        uint256 quoteAmount,
        uint256 feeAmount
    );

    /// @notice when user sell token
    /// @param token token address
    /// @param timestamp timestamp
    /// @param baseAmount meme token amount
    /// @param quoteAmount eth quote amount
    /// @param feeAmount fee amount
    event TokenSell(
        address indexed token,
        uint256 timestamp,
        uint256 baseAmount,
        uint256 quoteAmount,
        uint256 feeAmount
    );

    // uups initializer
    function initialize(
        uint256 _x,
        uint256 _y,
        uint256 _fundingGoal,
        uint256 _createFee,
        uint256 _feePercent
    ) public initializer {
        __UUPSUpgradeable_init();
        __Ownable_init(msg.sender);
        __ReentrancyGuard_init();

        fundingGoal = _fundingGoal;
        x = _x;
        y = _y;
        K = (_x * _y) / 1 ether;
        createFee = _createFee;
        feePercent = _feePercent;
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal virtual override {}

    /// user functions

    function createMeme(
        string memory name,
        string memory symbol,
        string memory description,
        string memory icon,
        string memory website,
        string memory twitter,
        string memory telegram,
        string memory discord,
        string memory github
    ) external payable nonReentrant returns (address) {
        require(msg.value >= createFee, "Insufficient fee");
        require(
            bytes(name).length > 0 &&
                bytes(symbol).length > 0 &&
                bytes(icon).length > 0,
            "name, symbol and icon are required"
        );
        fee += createFee;

        MemeToken token = new MemeToken(name, symbol);
        address tokenAddress = address(token);
        tokens[tokenAddress] = TokenState.FUNDING;
        tokenX[tokenAddress] = x;
        tokenY[tokenAddress] = y;
        emit TokenCreated(
            tokenAddress,
            msg.sender,
            block.timestamp,
            name,
            symbol,
            description,
            icon,
            website,
            twitter,
            telegram,
            discord,
            github
        );
        return tokenAddress;
    }

    // Handle token purchase
    function buy(address tokenAddress) external payable nonReentrant {
        require(
            tokens[tokenAddress] == TokenState.FUNDING,
            "Token is not funding"
        );
        require(msg.value > 0, "Zero value not allowed");
        require(tokenAddress != address(0), "Zero address");
        require(
            tokens[tokenAddress] != TokenState.NOT_CREATED,
            "Token not created"
        );

        // calculate fee
        uint256 valueToBuy = msg.value;
        uint256 valueToReturn;
        uint256 tokenCollateral = collateral[tokenAddress];
        uint256 _fee;

        uint256 remainingEthNeeded = fundingGoal - tokenCollateral;
        uint256 contributionWithoutFee = valueToBuy -
            calculateFee(valueToBuy, feePercent);
        if (contributionWithoutFee > remainingEthNeeded) {
            contributionWithoutFee = remainingEthNeeded;
            _fee = calculateFee(contributionWithoutFee, feePercent);
            uint256 totalCharged = contributionWithoutFee + _fee;
            if (valueToBuy > totalCharged) {
                valueToReturn = valueToBuy - totalCharged;
            } else {
                valueToReturn = 0;
                _fee = valueToBuy - remainingEthNeeded;
            }
        } else {
            _fee = calculateFee(valueToBuy, feePercent);
            valueToReturn = 0;
        }
        fee += _fee;

        MemeToken token = MemeToken(tokenAddress);
        uint256 amount = calculateTokensReceived(
            K,
            tokenX[tokenAddress],
            tokenY[tokenAddress],
            contributionWithoutFee
        );

        tokenCollateral += contributionWithoutFee;
        token.mint(msg.sender, amount);
        // update token-specific x and y
        tokenY[tokenAddress] = tokenY[tokenAddress] - amount;
        tokenX[tokenAddress] = tokenX[tokenAddress] + contributionWithoutFee;

        // when reached FUNDING_GOAL
        if (tokenCollateral >= fundingGoal) {
            // token.mint(address(this), INITIAL_SUPPLY);
            tokens[tokenAddress] = TokenState.WAIT_TRADING;
        }
        collateral[tokenAddress] = tokenCollateral;
        // return left
        if (valueToReturn > 0) {
            (bool success, ) = msg.sender.call{value: valueToReturn}(
                new bytes(0)
            );
            require(success, "ETH send failed");
        }
        emit TokenBuy(
            tokenAddress,
            block.timestamp,
            amount,
            contributionWithoutFee,
            _fee
        );
    }

    function sell(address tokenAddress, uint256 amount) external nonReentrant {
        require(
            tokens[tokenAddress] == TokenState.FUNDING,
            "Token is not funding"
        );
        require(amount > 0, "Amount should be greater than zero");
        MemeToken token = MemeToken(tokenAddress);
        require(
            token.balanceOf(msg.sender) >= amount,
            "Insufficient token balance"
        );

        uint256 receivedETH = calculateFundsReceived(
            K,
            tokenX[tokenAddress],
            tokenY[tokenAddress],
            amount
        );
        // calculate fee
        uint256 _fee = calculateFee(receivedETH, feePercent);
        uint256 receivedETHWithoutFee = receivedETH - _fee;
        fee += _fee;
        token.burn(msg.sender, amount);
        collateral[tokenAddress] -= receivedETH;
        // Update token-specific x and y
        tokenY[tokenAddress] = tokenY[tokenAddress] + amount;
        tokenX[tokenAddress] = tokenX[tokenAddress] - receivedETH;
        // send ether
        (bool success, ) = msg.sender.call{value: receivedETHWithoutFee}(
            new bytes(0)
        );
        require(success, "ETH send failed");
        emit TokenSell(
            tokenAddress,
            block.timestamp,
            amount,
            receivedETHWithoutFee,
            _fee
        );
    }

    function calculateFee(
        uint256 _amount,
        uint256 _feePercent
    ) internal pure returns (uint256) {
        return (_amount * _feePercent) / FEE_DENOMINATOR;
    }

    /// admin functions

    function setFundingGoal(uint256 newFundingGoal) external onlyOwner {
        fundingGoal = newFundingGoal;
    }

    function setFeePercent(uint256 _feePercent) external onlyOwner {
        feePercent = _feePercent;
    }

    function setCreateFee(uint256 _createFee) external onlyOwner {
        createFee = _createFee;
    }

    function claimFee() external onlyOwner {
        (bool success, ) = msg.sender.call{value: fee}(new bytes(0));
        require(success, "ETH send failed");
        fee = 0;
    }
}
