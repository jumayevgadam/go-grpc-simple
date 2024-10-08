package initialization

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/jumayevgadam/go-grpc-simple/gapi"
	"github.com/jumayevgadam/go-grpc-simple/proto/product"
	"github.com/jumayevgadam/go-grpc-simple/repository"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func DBInit() (*gorm.DB, error) {
	var dialect gorm.Dialector

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	}

	dsn := "host=localhost port=5432 user=postgres dbname=grpc_simple password=12345 sslmode=disable"
	dialect = postgres.Open(dsn)

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		return nil, fmt.Errorf("error to open: %v", err.Error())
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("db.DB(): error: %v", err.Error())
	}

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)

	tm := time.Minute * time.Duration(20)
	sqlDB.SetConnMaxLifetime(tm)

	return db, nil
}

func ServeInit(db *gorm.DB) error {
	network := "tcp"
	address := ":1200"
	listen, err := net.Listen(network, address)
	if err != nil {
		return fmt.Errorf("error to listen: %v", err.Error())
	}
	defer listen.Close()

	grpcServer := grpc.NewServer()

	// register
	productSvc := repository.NewDataProduct(db)
	productServer := gapi.NewGrpcProduct(productSvc)

	// put register here
	product.RegisterProductGapiServer(grpcServer, productServer)
	err = grpcServer.Serve(listen)
	if err != nil {
		return fmt.Errorf("error to serve: %v", err.Error())
	}
	log.Println("Server connected!!!")

	return nil
}
