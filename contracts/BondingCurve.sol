// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import {FixedPointMathLib} from "solady/src/utils/FixedPointMathLib.sol";

contract BondingCurve {
    using FixedPointMathLib for uint256;

    /// @dev Calculate tokens received for a given amount of ETH
    /// @param K The K value of the bonding curve K = x * y
    /// @param x fund amount
    /// @param y token amount
    /// @param dx input fund amount
    /// @return The amount of tokens received (dy = y - K / (x + dx))
    function calculateTokensReceived(
        uint256 K,
        uint256 x,
        uint256 y,
        uint256 dx
    ) public pure returns (uint256) {
        // K = x * y
        // K = (x + dx) * (y - dy)
        // dy = y - K / (x + dx)
        // (K * ethAmount) / (currentSupply * (currentSupply + ethAmount))

        return y - K.divWad((x + dx));
    }

    /// @dev Calculate ETH received for selling tokens
    /// @param K The K value of the bonding curve K = x * y
    /// @param x fund amount
    /// @param y token amount
    /// @param dy input fund amount
    /// @return The amount of ETH received (dx = x - K / (y + dy))
    function calculateFundsReceived(
        uint256 K,
        uint256 x,
        uint256 y,
        uint256 dy
    ) public pure returns (uint256) {
        // K = x * y
        // K = (x - dx) * (y + dy)
        // dx = dx = x - K / (y + dy)
        return x - K.divWad((y + dy));
    }
}
