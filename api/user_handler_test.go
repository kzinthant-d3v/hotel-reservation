package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/kzinthant-d3v/hotel-reservation/db"
	"github.com/kzinthant-d3v/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type testdb struct {
	db.UserStore
}

const (
	testmongouri = "mongodb://localhost:27017"
	dbname       = "hotel-reservation-testdb"
)

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup() *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testmongouri))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client, dbname),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup()
	defer tdb.teardown(t)
	app := fiber.New()
	userHandler := NewUserHandler(tdb.UserStore)

	app.Post("/", userHandler.HandlePostUser)

	params := types.CreateUserParams{
		FirstName: "kaskar",
		LastName:  "kaskar",
		Email:     "some@foo.com",
		Password:  "2323232323",
	}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Error(err)
	}
	var user types.User = types.User{}
	json.NewDecoder(resp.Body).Decode(&user)
	if len(user.EncryptedPassword) > 0 {
		t.Error("password should not be returned")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected firstname %s but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected lastname %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected email %s but got %s", params.Email, user.Email)
	}

	//bb, err := io.ReadAll(resp.Body)

}
