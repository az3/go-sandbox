# go-sandbox
Sample Go applications will be placed here.

Running
-------
```
mkdir -p $GOPATH/src/github.com/az3/
# go get -u github.com/FiloSottile/gvt
cd $GOPATH/src/github.com/az3/
git clone --depth=1 https://github.com/az3/go-sandbox.git
cd $GOPATH/src/github.com/az3/go-sandbox
go build ./cmd/hello/
./hello
```

