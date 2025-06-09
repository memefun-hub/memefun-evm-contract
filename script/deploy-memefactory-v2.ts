import hre, { ethers, upgrades } from "hardhat";

async function main() {
  const MemeFactoryV2 = await ethers.getContractFactory("MemeFactory");
  const memeFactory = await upgrades.upgradeProxy(
    "0x3665460E6d26FB30949c2038f94262aE03885843",
    MemeFactoryV2,
    {
      kind: "uups",
    }
  );

  await memeFactory.waitForDeployment();
  console.log("MemeFactory deployed to:", memeFactory.target);

  await hre.run("verify:verify", {
    address: memeFactory.target,
    constructorArguments: [],
  });
}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
