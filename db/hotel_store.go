package db

import (
	"context"

	"github.com/kzinthant-d3v/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HotelStore interface {
	GetHotels(context.Context, bson.M) ([]*types.Hotel, error)
	InsertHotel(context.Context, *types.Hotel) (*types.Hotel, error)
	Update(ctx context.Context, filter bson.M, update bson.M) error
}

type MongoHotelStore struct {
	client     *mongo.Client
	collection *mongo.Collection
}

const hotelCollection = "hotels"

func NewMongoHotelStore(client *mongo.Client) *MongoHotelStore {
	return &MongoHotelStore{
		client:     client,
		collection: client.Database(DBNAME).Collection(hotelCollection),
	}
}

func (s *MongoHotelStore) GetHotels(ctx context.Context, filter bson.M) ([]*types.Hotel, error) {
	 res, err := s.collection.Find(ctx, filter)
	 if err != nil {
		return nil, err
	 }
	 var hotels []*types.Hotel
	 if err := res.All(ctx, &hotels); err != nil {
		return nil, err
	 }
	  
	 return hotels, nil
}

func (s *MongoHotelStore) InsertHotel(ctx context.Context, hotel *types.Hotel) (*types.Hotel, error) {
	res, err := s.collection.InsertOne(ctx, hotel)
	if err != nil {
		return nil, err
	}
	hotel.ID = res.InsertedID.(primitive.ObjectID)
	return hotel, nil
}

func (s *MongoHotelStore) Update(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := s.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}
