package main

import (
	"_/C_/golang/todo"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run(port: "8000"); err != nil{
		log.Fatalf(format:"error occured while running http server: %s", err.Error())
	}

}
