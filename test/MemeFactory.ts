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
      hre.ethers.parseEther("1000000")
    );

    return { factory, owner, otherAccount };
  }

  describe("Deployment", function () {
    it("Should set funding goal", async function () {
      const { factory } = await loadFixture(deployMemeFactory);

      expect(await factory.fundingGoal()).to.equal(
        hre.ethers.parseEther("2000000")
      );
    });
  });

  describe("Token creation", function () {
    it("Should create a token", async function () {
      const { factory } = await loadFixture(deployMemeFactory);
      const factoryAddress = await factory.getAddress();
      console.log("factoryAddress:", factoryAddress);
      const tx = await factory.createMeme("Test", "TEST");
      const receipt = await tx.wait();
      //   console.log(receipt?.logs);
      const event = receipt?.logs.find(
        (log) => factory.interface.parseLog(log)?.name === "TokenCreated"
      );
      const parsedEvent = factory.interface.parseLog(event as EventLog);
      const tokenAddress = parsedEvent?.args.token;
      console.log("tokenAddress:", tokenAddress);
      expect(tokenAddress).to.not.be.null;
      const token = await hre.ethers.getContractAt("MemeToken", tokenAddress);
      expect(await token.name()).to.equal("Test");
      expect(await token.symbol()).to.equal("TEST");
      expect(await token.decimals()).to.equal(18);
      expect(await token.totalSupply()).to.equal(0);
      const owner = await token.owner();
      expect(owner).to.equal(factoryAddress);
    });
  });

  describe("Token buy", function () {
    it("Should purchase a token", async function () {
      const { factory } = await loadFixture(deployMemeFactory);
      const tokenAddress = await factory.createMeme("Test", "TEST");
    });
  });
});
