package main

import (
	"os"

	"github.com/hirukiyo/gin-sample/apiserver"
)

func main() {
	os.Exit(apiserver.StartAPIServer())
}
