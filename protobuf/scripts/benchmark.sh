#!/usr/bin/env bash

# setup data
echo ""
echo "Setup data"
redis-load -u localhost:6379 < $(pwd)/scripts/integration/data/redis.json
sleep 1

# run test
echo ""
echo "Run Tests"
echo ""
wrk -t12 -c400 -d30s --script=$(pwd)/scripts/benchmark/post.lua http://localhost:8080/rta/exposure


# cleanup data
echo ""
echo "Clean data"
echo ""
cat $(pwd)/scripts/integration/flush_db.txt | redis-cli --pipe
sleep 1