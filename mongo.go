package goqrcode

import (
	"context"
	// "encoding/json"
	"log"
	"time"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/bson"
)

// mongodb

func GetConnectionMongo(MONGOSTRING, dbname string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MONGOSTRING)))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	return client.Database(dbname), nil
}

func SetConnection(MONGOSTRINGENV, dbname string) (*mongo.Database, error) {
    client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MONGOSTRINGENV)))
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
    }
    return client.Database(dbname), nil
}


func MongoConnect(MONGOSTRINGENV, dbname string) *mongo.Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv(MONGOSTRINGENV)))
	if err != nil {
		fmt.Printf("MongoConnect: %v\n", err)
		return nil
	}
	return client.Database(dbname)
}

var collection *mongo.Collection

func insertPayload2() {
	// Establish a connection to your MongoDB instance
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://faisalTampan:9byL9bOl3rhqbSrO@soren.uwshwr6.mongodb.net/test"))
	if err != nil {
		panic(err)
	}

	// Assign the collection to your global variable
	collection = client.Database("Pakarbi").Collection("codeqr")
}

func saveToMongo(formData FormData) error {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://faisalTampan:9byL9bOl3rhqbSrO@soren.uwshwr6.mongodb.net/test"))
	if err != nil {
		return fmt.Errorf("failed to create MongoDB client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %v", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("PakArbi").Collection("codeqr")

	_, err = collection.InsertOne(ctx, formData)
	if err != nil {
		return fmt.Errorf("failed to insert data into MongoDB: %v", err)
	}

	return nil
}

func saveToMongoDB(data FormData) error {
	// Inisialisasi koneksi ke MongoDB
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	// Pilih database dan koleksi MongoDB
	collection := client.Database("nama_database").Collection("nama_koleksi")

	// Simpan data ke dalam MongoDB
	_, err = collection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}

	return nil
}

func InsertDataToMongoDB(formData FormData) error {
    // Establish a connection to MongoDB
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://faisalTampan:9byL9bOl3rhqbSrO@soren.uwshwr6.mongodb.net/test"))
    if err != nil {
        return fmt.Errorf("failed to connect to MongoDB: %v", err)
    }
    defer func() {
        if err = client.Disconnect(context.Background()); err != nil {
            log.Fatal(err)
        }
    }()

    // Access your database and collection
    collection := client.Database("PakArbi").Collection("codeqr")

    // Insert data into MongoDB
    _, err = collection.InsertOne(context.Background(), formData)
    if err != nil {
        return fmt.Errorf("failed to insert data to MongoDB: %v", err)
    }

    return nil
}

func establishMongoDBConnection() (*mongo.Client, error) {
    // Establish a connection to MongoDB
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://faisalTampan:9byL9bOl3rhqbSrO@soren.uwshwr6.mongodb.net/test"))
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
    }

    return client, nil
}

func insertDataToMongoDB(client *mongo.Client, formData FormData) error {
    // Access your database and collection
    collection := client.Database("PakArbi").Collection("codeqr")

    // Insert data into MongoDB
    _, err := collection.InsertOne(context.Background(), formData)
    if err != nil {
        return fmt.Errorf("failed to insert data to MongoDB: %v", err)
    }

    return nil
}

//insert
// func InsertDataToMongoDB(formData FormData) error {
// 	// Convert struct to JSON
// 	dataJSON, err := json.Marshal(formData)
//     if err != nil {
//         return fmt.Errorf("failed to marshal JSON: %v", err)
//     }

// 	// Establish a connection to MongoDB
// 	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://faisalTampan:9byL9bOl3rhqbSrO@soren.uwshwr6.mongodb.net/test"))
// 	if err != nil {
// 		return fmt.Errorf("failed to connect to MongoDB: %v", err)
// 	}
// 	defer func() {
// 		if err = client.Disconnect(context.Background()); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	// Access your database and collection
// 	collection := client.Database("PakArbi").Collection("codeqr")

// 	// Insert data into MongoDB
// 	_, err = collection.InsertOne(context.Background(), formData)
// 	if err != nil {
// 		return fmt.Errorf("failed to insert data to MongoDB: %v", err)
// 	}

// 	return nil
// }