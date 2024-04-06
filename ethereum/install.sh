set -x

git clone https://github.com/ethereum/go-ethereum.git
cd go-ethereum
git checkout v1.12.0
go build -o geth ./cmd/geth
sudo mv geth /usr/local/bin

cd ../

git clone https://github.com/Consensys/quorum.git
git checkout v23.4.0
cd quorum
go build -o quorum ./cmd/geth
sudo mv quorum /usr/local/bin

geth version
quorum version