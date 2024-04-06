#!/bin/sh

set -x

nohup geth \
--datadir data --networkid 15 --nodiscover \
--unlock 0x2b50a4fe2a6300116edbcf2632d079a12abf5f2d --password password.txt \
--http --http.api eth,web3,net,debug --http.addr "localhost" --http.corsdomain "*" \
--ws --ws.addr "localhost" --ws.port 8545 --ws.origins "*" \
--allow-insecure-unlock --gcmode archive \
--miner.etherbase '0x2b50a4fe2a6300116edbcf2632d079a12abf5f2d' --mine >& log.txt 2>&1 &