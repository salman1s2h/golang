## Shared Library

### MacOS

```
clang -dynamiclib -o build/libmylib.dylib csrc/mylib.c
export DYLD_LIBRARY_PATH=.
go run main.go
```

### Linux

```
clang -shared -o build/libmylib.so csrc/mylib.c
export LD_LIBRARY_PATH=.
go run main.go
```

### Windows

```
clang -shared -o build/mylib.dll csrc/mylib.c
set PATH=%PATH%;.
go run main.go
```


## Static Library

### Mac/Linux

```
# Compile the C code into an object file
clang -c csrc/mylib.c -o build/mylib.o

# Create a static library from the object file
ar rcs build/libmylib.a build/mylib.o
```

### Windows

```
# Compile the C code into an object file
clang -c csrc/mylib.c -o build/mylib.o

# Create a static library from the object file
ar rcs build/libmylib.lib build/mylib.o
```


## Install cross compiled stuff

### Linux
brew install FiloSottile/musl-cross/musl-cross
- verify
x86_64-linux-musl-gcc --version

### Windows 
brew install mingw-w64

-- verify
x86_64-w64-mingw32-gcc --version