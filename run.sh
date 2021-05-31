#!/bin/bash
filePath=../data/fng.1000.csv.rot128

go build -o build/app
cd build
./app -f $filePath
cd ..