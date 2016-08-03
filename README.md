# go-sandbox
Sample Go applications will be placed here.

Running
-------
```
mkdir -p $GOPATH/src/github.com/az3/
cd $GOPATH/src/github.com/az3/
git clone --depth=1 https://github.com/az3/go-sandbox.git
cd $GOPATH/src/github.com/az3/go-sandbox
# go get github.com/tools/godep
# go get gopkg.in/alexcesaro/quotedprintable.v3
# cd $GOPATH/src/github.com/tools/godep
# git checkout v32
cd $GOPATH/src/github.com/az3/go-sandbox
go build ./cmd/hello/
./hello
```

