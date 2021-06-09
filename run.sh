#!/bin/bash

filePath=../data/fng.1000.csv.rot128
omisePublicKey="pkey_test_5o4fc4wux9ec01nmmot"
omiseSecretKey="skey_test_5o4fc5etrop6k804sjv"
debug="false"

go build -o build/app || exit 1
cd build

./app -f $filePath \
    -pk $omisePublicKey \
    -sk $omiseSecretKey \
    -n 3 \
    -t 3 \
    -d=$debug

cd ..