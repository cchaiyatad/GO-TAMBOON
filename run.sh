#!/bin/bash
filePath=../data/fng.1000.csv.rot128
omisePublicKey="pkey_test_5o0xzraosm7v7cvl2vh"
omiseSecretKey="skey_test_5o0w33kibp2hq86knfy"

go build -o build/app || exit 1
cd build
./app -f $filePath \
    -pk $omisePublicKey \
    -sk $omiseSecretKey \
    -n 4
cd ..