package db

const (
	DBNAME     = "hotel-reservation-db"
	DBURI      = "mongodb://localhost:27017"
	TestDBNAME = "hotel-reservation-test-db"
)

// func ToObjectId(id string) (primitive.ObjectID) {
// 	oid, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return oid
// }
