#!/bin/sh

set -eux

cd ./c-lib/

# Linux transpile.

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

rm cdb-linux.json || true

CC=$(which gcc) \
AR=$(which ar) \
ccgo -compiledb cdb-linux.json make

CC=$(which gcc) \
AR=$(which ar) \
ccgo -pkgname webp -trace-translation-units -o ../lib/transpiled/webp/libwebp_linux_amd64.go cdb-linux.json src/.libs/libwebp.a

gofmt -s -w ../lib/transpiled/webp/libwebp_linux_amd64.go

# Linux arm64.

LINUX_ARM_HOST=aarch64-linux-gnu

make distclean || true
./autogen.sh
./configure \
    --host=$LINUX_ARM_HOST \
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

rm cdb-linux-arm64.json || true

CC=$(which $LINUX_ARM_HOST-gcc) \
AR=$(which $LINUX_ARM_HOST-ar) \
ccgo -compiledb cdb-linux-arm64.json make

CC=$(which $LINUX_ARM_HOST-gcc) \
AR=$(which $LINUX_ARM_HOST-ar) \
ccgo -pkgname webp -trace-translation-units -o ../lib/transpiled/webp/libwebp_linux_arm64.go cdb-linux-arm64.json src/.libs/libwebp.a

gofmt -s -w ../lib/transpiled/webp/libwebp_linux_arm64.go

# macOS transpile.

MACOS_HOST=x86_64-apple-darwin15

make distclean || true
./autogen.sh
./configure \
    CC=$MACOS_HOST-cc \
    CXX=$MACOS_HOST-c++ \
    --host=$MACOS_HOST \
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

rm cdb-darwin.json || true

CC=$(which $MACOS_HOST-cc) \
AR=$(which $MACOS_HOST-ar) \
ccgo -compiledb cdb-darwin.json make

CC=$(which $MACOS_HOST-cc) \
AR=$(which $MACOS_HOST-ar) \
ccgo -pkgname webp -trace-translation-units -o ../lib/transpiled/webp/libwebp_darwin_amd64.go cdb-darwin.json src/.libs/libwebp.a

gofmt -s -w ../lib/transpiled/webp/libwebp_darwin_amd64.go

# Windows transpile.

WINDOWS_HOST=x86_64-w64-mingw32

make distclean || true
./autogen.sh
./configure \
    --host=$WINDOWS_HOST \
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

rm cdb-windows.json || true

CC=$(which $WINDOWS_HOST-gcc) \
AR=$(which $WINDOWS_HOST-ar) \
ccgo -compiledb cdb-windows.json make

CC=$(which $WINDOWS_HOST-gcc) \
AR=$(which $WINDOWS_HOST-ar) \
ccgo -windows -pkgname webp -trace-translation-units -o ../lib/transpiled/webp/libwebp_windows_amd64.go cdb-windows.json src/.libs/libwebp.a

gofmt -s -w ../lib/transpiled/webp/libwebp_windows_amd64.go


