package goqrcode

import (
	"context"
	// "encoding/json"
	"log"
	"time"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

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

func SaveQRScanResult(qrData QRScan) error {
    // Koneksi ke MongoDB
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://faisalTampan:9byL9bOl3rhqbSrO@soren.uwshwr6.mongodb.net/test"))
    if err != nil {
        return err
    }
    defer func() {
        if err = client.Disconnect(context.Background()); err != nil {
            panic(err)
        }
    }()

    // Pilih database dan koleksi
    database := client.Database("PakArbiQR")
    collection := database.Collection("QRData")

    // Menyimpan qrData ke dalam MongoDB
    _, err = collection.InsertOne(context.Background(), qrData)
    if err != nil {
        return err
    }

    return nil
}


func saveToMongo(dataParkir Parkiran) error {
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

	_, err = collection.InsertOne(ctx, dataParkir)
	if err != nil {
		return fmt.Errorf("failed to insert data into MongoDB: %v", err)
	}

	return nil
}

func InsertDataToMongoDB(dataParkir Parkiran) error {
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
    _, err = collection.InsertOne(context.Background(), dataParkir)
    if err != nil {
        return fmt.Errorf("failed to insert data to MongoDB: %v", err)
    }

    return nil
}

func SequenceAutoIncrement(mongoconn *mongo.Database, sequenceName string) int {
	filter := bson.M{"_id": sequenceName}
	update := bson.M{"$inc": bson.M{"seq": 1}}

	var result struct {
		Seq int `bson:"seq"`
	}

	after := options.After
	opt := &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	}
	collection := mongoconn.Collection("counters")
	err := collection.FindOneAndUpdate(context.TODO(), filter, update, opt).Decode(&result)
	if err != nil {
		// handle error
	}
	return result.Seq
}

func establishMongoDBConnection() (*mongo.Client, error) {
    // Establish a connection to MongoDB
    client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://faisalTampan:9byL9bOl3rhqbSrO@soren.uwshwr6.mongodb.net/test"))
    if err != nil {
        return nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
    }

    return client, nil
}

func insertDataToMongoDB(client *mongo.Client, dataParkir Parkiran) error {
    // Access your database and collection
    collection := client.Database("PakArbi").Collection("codeqr")

    // Insert data into MongoDB
    _, err := collection.InsertOne(context.Background(), dataParkir)
    if err != nil {
        return fmt.Errorf("failed to insert data to MongoDB: %v", err)
    }

    return nil
}

//insert
// func InsertDataToMongoDB(dataParkir Parkiran) error {
// 	// Convert struct to JSON
// 	dataJSON, err := json.Marshal(dataParkir)
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
// 	_, err = collection.InsertOne(context.Background(), dataParkir)
// 	if err != nil {
// 		return fmt.Errorf("failed to insert data to MongoDB: %v", err)
// 	}

// 	return nil
// }