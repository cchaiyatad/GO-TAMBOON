#!/bin/bash
filePath=../data/fng.1000.csv.rot128
omisePublicKey="pkey_test_5o0w33kibttrj8mbs85"
omiseSecretKey="skey_test_5o0w33kibp2hq86knfy"

go build -o build/app
cd build
./app -f $filePath \
    -pk $omisePublicKey \
    -sk $omiseSecretKey
cd ..