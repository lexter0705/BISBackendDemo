package main

import (
	"BISBackend/Database"
	"BISBackend/Web"
)

func main() {
	Database.NewDatabase()
	Web.RunServer()
}
