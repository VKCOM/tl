#!/bin/bash

path=speed_proto
fast_path=speed_proto_fast

rm -rf $path/pb
rm -rf $fast_path/pb_fast


mkdir -p $path/pb
protoc \
  -I ./schemas \
  --go_out=./$path/pb \
  --go_opt=paths=source_relative \
  ./schemas/speed.proto

mkdir -p $fast_path/pb_fast
protoc \
  -I ./schemas \
  --gofast_out=paths=source_relative:./$fast_path/pb_fast \
  ./schemas/speed_fast.proto