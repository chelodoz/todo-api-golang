package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type ID = primitive.ObjectID

func IDToString(id ID) string {
	return id.Hex()
}
func StringToID(id string) ID {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	return objID
}

func NewID() ID {
	return primitive.NewObjectID()
}
