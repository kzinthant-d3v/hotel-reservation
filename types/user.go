package types

type User struct {
	//omit field
	//ID string `bson:"_id" json:"_"`
	ID string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
  LastName string `bson:"lastName" json:"lastName"`
}