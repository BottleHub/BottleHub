#!/bash
solc --abi contracts/Token.sol -o binaries --overwrite
solc --bin contracts/Token.sol -o binaries --overwrite
abigen --bin=binaries/BottleDrops.bin --abi=binaries/BottleDrops.abi --pkg=basic --out='contract builds/bottledrops.go'
