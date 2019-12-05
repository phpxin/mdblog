package conf

import (
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

var (
	ConfigInst = &Config{}
)

type Config struct {
	RpcHost string `json:"rpchost"`
	Storagepath string `json:"storagepath"`
}

func ParseConfig(cpath string) error {
	fp,err := os.Open(cpath)
	if err!=nil {
		return err
	}

	content,err := ioutil.ReadAll(fp)
	if err!=nil {
		return err
	}
	fmt.Println(string(content))

	err = json.Unmarshal(content, ConfigInst)
	if err!=nil {
		return err
	}

	fmt.Println(ConfigInst)
	return nil
}