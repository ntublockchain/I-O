rm -rf ./examples/ecomm/tmp/auction_info.json
cd kafka
docker-compose down
cd ../fabric-samples/test-network
./network.sh down
#docker volume prune