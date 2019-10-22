package main

import (
	"./core"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	core.HTTP()
}
