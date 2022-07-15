#!/bin/sh

git checkout .
git pull
pkill -f nft
go build
nohup /root/nft/nft >> nft.log &
tail -f /root/nft/nft.log