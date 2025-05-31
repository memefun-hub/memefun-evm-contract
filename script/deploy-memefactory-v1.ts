import hre, { ethers, upgrades } from "hardhat";

async function main() {
  const MemeFactory = await ethers.getContractFactory("MemeFactory");
  const memeFactory = await upgrades.deployProxy(MemeFactory, [
    ethers.parseEther("30000"),
    ethers.parseEther("1073000191"),
    ethers.parseEther("1000000"),
  ], {
    initializer: 'initialize',
    kind: 'uups'
  });

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
