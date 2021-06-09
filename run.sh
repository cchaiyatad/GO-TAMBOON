#!/bin/bash
filePath=../data/fng.1000.csv.rot128
omisePublicKey="pkey_test_5o0y2isu7yq96qbfrz3"
omiseSecretKey="skey_test_5o0y2isu7yd5ruh6vri"

go build -o build/app || exit 1
cd build
./app -f $filePath \
    -pk $omisePublicKey \
    -sk $omiseSecretKey \
    -n 4 \
    -t 4
cd ..