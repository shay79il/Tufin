package main

import (
	"fmt"
	"main/cluster"
	"os"
)

func main() {

	if err := cluster.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
