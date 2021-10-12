#!/usr/bin/env bash


# run test
echo ""
echo "Run Tests"
echo ""
go test -count=1 -tags=integration ./... -v
sleep 1

# cleanup data
echo ""
echo "Clean data"
echo ""
< $(pwd)/scripts/integration/flush_db.txt redis-cli --pipe
sleep 1