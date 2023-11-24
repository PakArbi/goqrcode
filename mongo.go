package goqrcode

import (
	"context"
	"log"
	"time"

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

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://faisalTampan:9byL9bOl3rhqbSrO@soren.uwshwr6.mongodb.net/test"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("ulbi").Collection("users")
}

// InsertPayload inserts a user into MongoDB
func InsertPayload(payload Payload) error {
	_, err := collection.InsertOne(context.Background(), payload)
	return err
}

func insertDataIntoMongo(email string) error {
	_, err := collection.InsertOne(context.Background(), bson.M{"email": email})
	return err
}

// func (m *MockDatabase) InsertPayload(payload Payload) error {
// 	if m.InsertPayloadFunc != nil {
// 		return m.InsertPayloadFunc(payload)
// 	}
// 	return nil
// }

// MockDatabase is a mock for the Database interface
type MockDatabase struct {
	InsertPayloadFuncCalled bool
}

// InsertPayload simulates inserting a payload into the database
func (m *MockDatabase) InsertPayload(payload Payload) error {
	m.InsertPayloadFuncCalled = true
	// Implement logic to simulate payload insertion
	return nil
}

// MockEmailSender is a mock for the EmailSender interface
type MockEmailSender struct {
	SendVerificationEmailFuncCalled bool
}

// SendVerificationEmail mocks sending a verification email
func (m *MockEmailSender) SendVerificationEmail(email string) error {
	m.SendVerificationEmailFuncCalled = true
	// Implement logic to simulate email sending
	return nil
}


// InsertPayload simulates inserting a payload into the database
func (m *MockDB) InsertPayload(payload Payload) error {
	// Implement the logic to simulate the payload insertion
	m.EmailSent = true // Simulate successful payload insertion
	return nil
}

// // SendVerificationEmail mocks sending a verification email
// func (m *MockEmailSender) SendVerificationEmail(email string) error {
// 	// Here you can check if the email sending is triggered correctly
// 	m.VerificationEmailSent = true
// 	return nil
// }