import {
  time,
  loadFixture,
} from "@nomicfoundation/hardhat-toolbox/network-helpers";
import { anyValue } from "@nomicfoundation/hardhat-chai-matchers/withArgs";
import { expect } from "chai";
import hre from "hardhat";
import { EventLog } from "ethers";

describe("MemeFactory", function () {
  // We define a fixture to reuse the same setup in every test.
  // We use loadFixture to run this setup once, snapshot that state,
  // and reset Hardhat Network to that snapshot in every test.
  async function deployMemeFactory() {
    // Contracts are deployed using the first signer/account by default
    const [owner, otherAccount] = await hre.ethers.getSigners();

    const MemeFactory = await hre.ethers.getContractFactory("MemeFactory");
    const factory = await MemeFactory.deploy();
    await factory.initialize(
      hre.ethers.parseEther("30000"),
      hre.ethers.parseEther("1073000191"),
      hre.ethers.parseEther("1000000"),
      hre.ethers.parseEther("1000"),
      100
    );

    return { factory, owner, otherAccount };
  }

  describe("Deployment", function () {
    it("Should set funding goal", async function () {
      const { factory } = await loadFixture(deployMemeFactory);

      expect(await factory.fundingGoal()).to.equal(
        hre.ethers.parseEther("1000000")
      );
    });
  });

  describe("Token creation and buy", function () {
    it("Should create a token", async function () {
      const [signer] = await hre.ethers.getSigners();
      const { factory } = await loadFixture(deployMemeFactory);
      const factoryAddress = await factory.getAddress();
      console.log("factoryAddress:", factoryAddress);
      const tx = await factory.createMeme("Test", "TEST", "", "icon", "", "", "","", "",{
        value: hre.ethers.parseEther("1000")
      });
      const receipt = await tx.wait();
      //   console.log(receipt?.logs);
      const event = receipt?.logs.find(
        (log) => factory.interface.parseLog(log)?.name === "TokenCreated"
      );
      const parsedEvent = factory.interface.parseLog(event as EventLog);
      const tokenAddress = parsedEvent?.args.token;
      console.log("tokenAddress:", tokenAddress);
      console.log(await factory.fee())
      expect(tokenAddress).to.not.be.null;
      const token = await hre.ethers.getContractAt("MemeToken", tokenAddress);
      expect(await token.name()).to.equal("Test");
      expect(await token.symbol()).to.equal("TEST");
      expect(await token.decimals()).to.equal(18);
      expect(await token.totalSupply()).to.equal(0);
      const owner = await token.owner();
      expect(owner).to.equal(factoryAddress);
      const ethBalanceBefore = await hre.ethers.provider.getBalance(signer.address);
      console.log("buy before ethBalance:", ethBalanceBefore);
      console.log("collateral:", await factory.collateral(tokenAddress))
      // buy
      for (let i = 0; i < 9; i++) {
        const buy = await factory.buy(tokenAddress, {
          value: hre.ethers.parseEther("100000")
        });
        await buy.wait()
        console.log(await factory.fee())
      }

      // console.log("================", await hre.ethers.provider.getBalance(signer.address))
      // const buy = await factory.buy(tokenAddress, {
      //   value: hre.ethers.parseEther("300000")
      // });
      // await buy.wait()
      // console.log("================", await hre.ethers.provider.getBalance(signer.address))

      const ethBalanceAfter = await hre.ethers.provider.getBalance(signer.address);
      console.log("buy after ethBalance:", ethBalanceAfter);
      console.log("collateral:", await factory.collateral(tokenAddress));

      let tokenBalance = await token.balanceOf(signer.address);
      console.log("tokenBalance:", tokenBalance);
      console.log("tokenX:", await factory.tokenX(tokenAddress))
      console.log("tokenY:", await factory.tokenY(tokenAddress))

      await factory.sell(tokenAddress, tokenBalance/2n);

      tokenBalance = await token.balanceOf(signer.address);
      await factory.sell(tokenAddress, tokenBalance);

      const ethBalanceAfterSell = await hre.ethers.provider.getBalance(signer.address);
      console.log("sell after ethBalance:", ethBalanceAfterSell);
      console.log("tokenX:", await factory.tokenX(tokenAddress))
      console.log("tokenY:", await factory.tokenY(tokenAddress))
      console.log("collateral:", await factory.collateral(tokenAddress));
      
    });
  });


});
