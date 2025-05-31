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

    uint256 public constant MAX_SUPPLY = 10 ** 9 * 1 ether;
    uint256 public constant FEE_DENOMINATOR = 10000;

    // bonding curve parameters
    uint256 public K;
    uint256 public x;
    uint256 public y;

    uint256 public fundingGoal;
    mapping(address => TokenState) public tokens;
    mapping(address => uint256) public collateral; // total collateral for each token
    uint256 public feePercent; // bp
    uint256 public fee;
    uint256 public createFee = 1000 ether;

    /// events
    event TokenCreated(address indexed token, uint256 timestamp);
    event TokenBuy(
        address indexed token,
        uint256 timestamp,
        uint256 fundAmount,
        uint256 returnTokenAmount,
        uint256 feeAmount
    );
    event TokenSell(
        address indexed token,
        uint256 timestamp,
        uint256 tokenAmount,
        uint256 returnFundAmount
    );

    // uups initializer
    function initialize(
        uint256 _x,
        uint256 _y,
        uint256 _fundingGoal
    ) public initializer {
        __UUPSUpgradeable_init();
        __Ownable_init(msg.sender);
        __ReentrancyGuard_init();

        fundingGoal = _fundingGoal;
        x = _x;
        y = _y;
        K = (x * y) / 1 ether;
    }

    function _authorizeUpgrade(
        address newImplementation
    ) internal virtual override {}

    /// user functions

    function createMeme(
        string memory name,
        string memory symbol
    ) external payable nonReentrant returns (address) {
        require(msg.value >= createFee, "Insufficient fee");
        fee += createFee;

        MemeToken token = new MemeToken(name, symbol);
        address tokenAddress = address(token);
        tokens[tokenAddress] = TokenState.FUNDING;
        emit TokenCreated(tokenAddress, block.timestamp);
        return tokenAddress;
    }

    // Handle token purchase
    function buy(address tokenAddress) external payable nonReentrant {
        require(
            tokens[tokenAddress] == TokenState.FUNDING ||
                tokens[tokenAddress] == TokenState.WAIT_TRADING,
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

        uint256 remainingEthNeeded = fundingGoal - tokenCollateral;
        uint256 contributionWithoutFee = (valueToBuy * FEE_DENOMINATOR) /
            (FEE_DENOMINATOR + feePercent);
        if (contributionWithoutFee > remainingEthNeeded) {
            contributionWithoutFee = remainingEthNeeded;
        }
        uint256 _fee = calculateFee(contributionWithoutFee, feePercent);
        uint256 totalCharged = contributionWithoutFee + _fee;
        valueToReturn = valueToBuy > totalCharged
            ? valueToBuy - totalCharged
            : 0;
        fee += _fee;
        MemeToken token = MemeToken(tokenAddress);
        uint256 amount = calculateTokensReceived(
            K,
            x,
            y,
            contributionWithoutFee
        );

        // uint256 availableSupply = FUNDING_SUPPLY - token.totalSupply();
        // require(amount <= availableSupply, "Token supply not enough");
        tokenCollateral += contributionWithoutFee;
        token.mint(msg.sender, amount);
        // update x and y
        y = y - amount;
        x = x + contributionWithoutFee;

        // when reached FUNDING_GOAL
        if (collateral[tokenAddress] >= fundingGoal) {
            // token.mint(address(this), INITIAL_SUPPLY);
            tokens[tokenAddress] = TokenState.WAIT_TRADING;
        }

        // return left
        if (valueToReturn > 0) {
            (bool success, ) = msg.sender.call{value: msg.value - valueToBuy}(
                new bytes(0)
            );
            require(success, "ETH send failed");
        }
        emit TokenBuy(
            tokenAddress,
            block.timestamp,
            msg.value,
            valueToReturn,
            _fee
        );
    }

    function sell(address tokenAddress, uint256 amount) external nonReentrant {
        require(
            tokens[tokenAddress] == TokenState.FUNDING ||
                tokens[tokenAddress] == TokenState.WAIT_TRADING,
            "Token is not funding"
        );
        require(amount > 0, "Amount should be greater than zero");
        MemeToken token = MemeToken(tokenAddress);
        uint256 receivedETH = calculateFundsReceived(K, x, y, amount);
        // calculate fee
        uint256 _fee = calculateFee(receivedETH, feePercent);
        receivedETH -= _fee;
        fee += _fee;
        token.burn(msg.sender, amount);
        collateral[tokenAddress] -= receivedETH;
        // send ether
        (bool success, ) = msg.sender.call{value: receivedETH}(new bytes(0));
        require(success, "ETH send failed");
        emit TokenSell(tokenAddress, block.timestamp, amount, receivedETH);
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
