#!/bin/sh

set -x

export PRIVATE_CONFIG=ignore; \
nohup quorum \
--datadir data --networkid 15 --nodiscover \
--raft --raftport 50000 --verbosity 5 \
--http --http.addr "localhost" --http.port 8546 --http.corsdomain "*" \
--ws --ws.addr "localhost" --ws.port 8546 --ws.origins "*" \
--http.api admin,db,eth,debug,miner,net,shh,txpool,personal,web3,quorum,raft \
--emitcheckpoints --port 21000 >& log.txt 2>&1 &
