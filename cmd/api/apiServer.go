package main

import (
	"fmt"

	"github.com/n0byk/marbu/config"
)

func init() {

	config.InitConfig()
}

func main() {
	fmt.Println(config.Config)
	// now do something with s3 or whatever
}
