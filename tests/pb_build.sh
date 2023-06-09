#!/bin/bash

protofile="./testa.proto"

GEN_DIR=./gen

mkdir $GEN_DIR

protoc --reatang-demo_out=$GEN_DIR \
      --reatang-demo_opt=logtostderr=true,loglevel=debug \
      $protofile