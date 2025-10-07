package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func testMySQL() {
	dsn := "root:root@tcp(localhost:3307)/app"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("MySQL connection error: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("MySQL ping failed: %v", err)
	}
	fmt.Println("✅ Connected to MySQL")
}

func testPostgres() {
	connStr := "host=localhost port=5433 user=user password=password dbname=app sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Postgres connection error: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Postgres ping failed: %v", err)
	}
	fmt.Println("✅ Connected to Postgres")
}

func testMongo() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://root:example@localhost:27018/app?authSource=admin")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}
	fmt.Println("✅ Connected to MongoDB")
}

func main() {
	fmt.Println("Testing database connections...")

	testMySQL()
	testPostgres()
	testMongo()

	fmt.Println("🎉 All database connections successful!")
}
