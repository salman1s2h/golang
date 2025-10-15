## Shared Library

### MacOS

```
clang -dynamiclib -o libmylib.dylib mylib.c
export DYLD_LIBRARY_PATH=.
go run main.go
```

### Linux

```
clang -shared -o libmylib.so mylib.c
export LD_LIBRARY_PATH=.
go run main.go
```

### Windows

```
clang -shared -o mylib.dll mylib.c
set PATH=%PATH%;.
go run main.go
```


## Static Library

### Mac/Linux

```
# Compile the C code into an object file
clang -c mylib.c -o mylib.o

# Create a static library from the object file
ar rcs libmylib.a mylib.o
```

### Windows

```
# Compile the C code into an object file
clang -c mylib.c -o mylib.o

# Create a static library from the object file
ar rcs libmylib.lib mylib.o
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