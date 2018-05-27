package main

import (
	"os"

	"github.com/Kamaropoulos/goctapus"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	goctapus.Init(os.Args, "debug")

	// Your application code goes here

	goctapus.Start()
}
