package kodo

import (
	"fmt"
	"testing"
)

var bucket1 = "mxdblcf"
var filePath1 = "iptables.jpg"
var key = filePath1

func TestUpLoad(t *testing.T) {

	load, i := UpLoad(filePath, key)
	fmt.Println(load)
	fmt.Println(i)
}

func TestDelete(t *testing.T) {
	Delete(bucket1, key)
}

func TestFileInfo(t *testing.T) {

	FileInfo(bucket, key)
}

func TestFetchNet(t *testing.T) {
	resURL := "http://devtools.qiniu.com/qiniu.png"
	FetchNet(bucket1, resURL)
}
