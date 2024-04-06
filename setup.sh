cd kafka
docker-compose up -d
cd ../fabric-samples/test-network

./network.sh up createChannel -ca
sudo ./network.sh deployCC -ccn MDai -ccp ../../contracts/fabric_erc20/chaincode -ccl go
sudo ./network.sh deployCC -ccn asset -ccp ../../contracts/fabric_asset/chaincode -ccl go

