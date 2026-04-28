#!/bin/bash

path=speed_tl
tl2gen=../../target/bin/tl2gen

mkdir -p $path

$tl2gen --language=go \
  --tl2WhiteList=* \
  --outdir=$path \
  --pkgPath=github.com/VKCOM/tl/internal/speed/$path/tl \
  schemas/speed.tl