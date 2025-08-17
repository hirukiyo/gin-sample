package main

import (
	"ginapp/apiserver"
	"os"
)

func main() {
	os.Exit(apiserver.StartAPIServer())
}
