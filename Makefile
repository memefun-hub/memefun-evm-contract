compile:
	npx hardhat compile

deploy-abCoreTestnet:
	npx hardhat run script/deploy-memefactory-v1.ts --network abCoreTestnet
