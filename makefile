# Author: d1y<chenhonzhou@gmail.com>

# creat `ntsd.exe` bin file
# before use: go get github.com/twitter/go-bindata/go-bindata
bindata:
	go-bindata ntsd.exe

build_windows:
	go build -o build/kill.exe .