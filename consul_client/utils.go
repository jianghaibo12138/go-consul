package consul_client

import (
	"io/ioutil"
	"os"
)

func ReadJsonConf(filePath string) []byte {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return bytes
}