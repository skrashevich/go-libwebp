#!/bin/sh

set -eux

cd ./c-lib/
make distclean || true
./autogen.sh
./configure \
    --disable-libwebpmux \
    --disable-libwebpdemux \
    --disable-sse4.1 \
    --disable-sse2 \
    --disable-neon \
    --disable-neon-rtcd \
    --disable-threading \
    --disable-gl \
    --disable-sdl \
    --disable-png \
    --disable-jpeg \
    --disable-tiff  \
    --disable-gif  \
    --disable-wic \
    --disable-shared
ccgo -compiledb cdb.json make
CC=$(which gcc) ccgo -pkgname lib -trace-translation-units -o ../lib/libwebp.go cdb.json src/.libs/libwebp.a
gofmt -s -w ../lib/libwebp.go
