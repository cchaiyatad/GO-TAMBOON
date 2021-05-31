#!/bin/bash

go build -o build/app
cd build
./app -f ../data/fng.1000.csv.rot128
cd ..