package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	dbinfra "github.com/GigioSanches/Desafio3Golang/internal/infra/db"
	graphqlinfra "github.com/GigioSanches/Desafio3Golang/internal/infra/graphql"
	grpcinfra "github.com/GigioSanches/Desafio3Golang/internal/infra/grpc"
	pb "github.com/GigioSanches/Desafio3Golang/internal/infra/grpc/proto"
	httpinfra "github.com/GigioSanches/Desafio3Golang/internal/infra/http"
	"github.com/GigioSanches/Desafio3Golang/internal/usecase"
	_ "github.com/lib/pq"
	grpcpkg "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func waitForDB(dsn string) *sql.DB {
	var db *sql.DB
	var err error
	for i := 0; i < 20; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			if err = db.Ping(); err == nil {
				return db
			}
		}
		log.Println("Waiting for database...", err)
		time.Sleep(2 * time.Second)
	}
	log.Fatalf("Could not connect to database: %v", err)
	return nil
}

func main() {
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN env var required")
	}
	db := waitForDB(dsn)
	defer db.Close()

	repo := dbinfra.NewOrderRepositoryDB(db)
	listOrdersUC := &usecase.ListOrdersUseCase{OrderRepo: repo}

	// REST
	handler := httpinfra.NewOrderHandler(listOrdersUC)
	router := httpinfra.NewRouter(handler)
	go func() {
		log.Println("REST API running on :8080")
		log.Fatal(http.ListenAndServe(":8080", router))
	}()

	// gRPC
	grpcServer := grpcpkg.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, grpcinfra.NewOrderServiceServer(listOrdersUC))
	reflection.Register(grpcServer)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("gRPC server running on :50051")
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve gRPC: %v", err)
		}
	}()

	// GraphQL
	gqlHandler := graphqlinfra.NewGraphQLHandler(listOrdersUC)
	go func() {
		log.Println("GraphQL server running on :8081")
		log.Fatal(http.ListenAndServe(":8081", gqlHandler))
	}()

	select {}
}
