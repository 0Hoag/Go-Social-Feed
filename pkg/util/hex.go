package util

import "go.mongodb.org/mongo-driver/bson/primitive"

func ObjectIDsToHex(ids []primitive.ObjectID) []string {
	if len(ids) == 0 {
		return []string{}
	}
	hexes := make([]string, 0, len(ids))
	for _, id := range ids {
		hexes = append(hexes, id.Hex())
	}
	return hexes
}
