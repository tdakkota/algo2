# algo2 ![go2go](https://github.com/tdakkota/algo2/workflows/go2go/badge.svg) [![codecov](https://codecov.io/gh/tdakkota/algo2/branch/master/graph/badge.svg?token=9QWNCNZHO2)](https://codecov.io/gh/tdakkota/algo2)
Algorithms written in Go2. 

Some packages are ported from [go2go testdata](https://github.com/golang/go/tree/dev.go2go/src/cmd/go2go/testdata/go2path/src).

# Installation 
### Requirements
- Existing >=go1.4 installation
- `go2go` translation tool
- [Taskfile](https://taskfile.dev/)

### Clone and build Go from dev.go2go ([Installing from source guide](https://golang.org/doc/install/source))
```sh
$ git clone https://go.googlesource.com/go go2goroot
$ cd go2goroot
$ git checkout dev.go2go
$ cd ./go2goroot/src && ./make.bash
```

### Clone the repo into your `GOPATH/GO2PATH`
```sh
$ git clone https://github.com/tdakkota/algo2 $GO2PATH/src/github.com/tdakkota/algo2
```

### Install [Taskfile](https://taskfile.dev/#/installation)
```sh
$ wget https://taskfile.dev/install.sh 
$ chmod +x install.sh
$ ./install.sh -b /usr/local/bin
$ rm install.sh
```

### Test
```sh
$ cd $GO2PATH/src/github.com/tdakkota/algo2
$ task test
```
