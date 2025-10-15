#!/bin/bash

# Create output directories
mkdir -p lib/darwin lib/linux lib/windows

# Compile for macOS (native)
echo "Compiling for macOS..."
clang -c csrc/mylib.c -o mylib.o -Iinclude
ar rcs lib/darwin/libmylib.a mylib.o
rm mylib.o

# Compile for Linux (x86_64)
echo "Compiling for Linux..."
x86_64-linux-musl-gcc -c csrc/mylib.c -o mylib.o -Iinclude
ar rcs lib/linux/libmylib.a mylib.o
rm mylib.o

# Compile for Windows (x86_64)
echo "Compiling for Windows..."
x86_64-w64-mingw32-gcc -c csrc/mylib.c -o mylib.o -Iinclude
x86_64-w64-mingw32-ar rcs lib/windows/mylib.lib mylib.o
rm mylib.o

# Build Go program
echo "Building Go program..."
go build -o myprogram main.go

echo "Done!"