package main

import (
	"github.com/jumayevgadam/go-grpc-simple/initialization"
	"github.com/jumayevgadam/go-grpc-simple/model"
)

func main() {
	db, err := initialization.DBInit()
	if err != nil {
		panic(err.Error())
	}

	_ = db.AutoMigrate(&model.Product{})

	err = initialization.ServeInit(db)
	if err != nil {
		panic(err)
	}
}
