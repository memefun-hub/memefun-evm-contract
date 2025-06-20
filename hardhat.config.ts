import { HardhatUserConfig, vars } from "hardhat/config";
import "@nomicfoundation/hardhat-toolbox";
import "@openzeppelin/hardhat-upgrades";
import "@nomiclabs/hardhat-solhint";
import "hardhat-abi-exporter";

const config: HardhatUserConfig = {
  solidity: {
    version: "0.8.28",
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  networks: {
    hardhat: {
      accounts: {
        accountsBalance: "10000000000000000000000000000" // 1000000 ETH
      }
    },
    abCoreTestnet: {
      hardfork: "berlin",
      url: `https://rpc.core.testnet.ab.org`,
      chainId: 26888,
      accounts: [vars.get("MEME_FUN_PRIVATE_KEY")],
    },
    abCoreMainnet: {
      hardfork: "berlin",
      url: `https://rpc.core.ab.org`,
      chainId: 36888,
      accounts: [vars.get("MEME_FUN_PRIVATE_KEY")],
    },
  },
  etherscan: {
    apiKey: {
      abCoreTestnet: "empty",
      abCoreMainnet: "empty",
    },
    customChains: [
      {
        network: "abCoreTestnet",
        chainId: 26888,
        urls: {
          apiURL: "https://explorer.core.testnet.ab.org/api",
          browserURL: "https://explorer.core.testnet.ab.org",
        },
      },
      {
        network: "abCoreMainnet",
        chainId: 36888,
        urls: {
          apiURL: "https://explorer.core.ab.org/api",
          browserURL: "https://explorer.core.ab.org",
        },
      },
    ],
  },
  abiExporter: {
    runOnCompile: true,
    clear: true,
    flat: true,
    // pretty: true,
    only: [":MemeFactory$", ":MemeToken$"],
    format: "json",
  },
};

export default config;
