package main

import "fmt"

type Storage interface {
	Store(data string)
}

type Database interface {
	Save(data string)
}

type CloudFactory interface {
	CreateStorage() Storage
	CreateDatabase() Database
}

// AWS

type S3 struct{}

func (s S3) Store(data string) {
	fmt.Println("Stored in AWS S3:", data)
}

type RDS struct{}

func (r RDS) Save(data string) {
	fmt.Println("Saved in AWS RDS:", data)
}

type AWSFactory struct{}

func (AWSFactory) CreateStorage() Storage {
	return S3{}
}

func (AWSFactory) CreateDatabase() Database {
	return RDS{}
}

// Azure

type BlobStorage struct{}

func (b BlobStorage) Store(data string) {
	fmt.Println("Stored in Azure Blob:", data)
}

type CosmosDB struct{}

func (c CosmosDB) Save(data string) {
	fmt.Println("Saved in Cosmos DB:", data)
}

type AzureFactory struct{}

func (AzureFactory) CreateStorage() Storage {
	return BlobStorage{}
}

func (AzureFactory) CreateDatabase() Database {
	return CosmosDB{}
}

func main() {
	var factory CloudFactory

	factory = AWSFactory{}

	storage := factory.CreateStorage()
	db := factory.CreateDatabase()

	storage.Store("hello")
	db.Save("hello")

	factory = AzureFactory{}

	storage = factory.CreateStorage()
	db = factory.CreateDatabase()

	storage.Store("hello")
	db.Save("hello")
}

/*
	Abstract Factory creates families of related objects.

Example:

AWS Factory
 ├── AWS Storage
 └── AWS Database

Azure Factory
 ├── Azure Storage
 └── Azure Database

*/
