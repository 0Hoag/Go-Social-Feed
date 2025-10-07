package mongo

import (
	"context"
	"math"
	"time"

	"github.com/hoag/go-social-feed/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ObjectIDFromHexOrNil returns an ObjectID from the provided hex representation.
func ObjectIDFromHexOrNil(id string) primitive.ObjectID {
	objID, _ := primitive.ObjectIDFromHex(id)
	return objID
}

func ObjectIDsFromHexOrNil(ids []string) []primitive.ObjectID {
	objIDs := make([]primitive.ObjectID, len(ids))
	for i, id := range ids {
		objIDs[i] = ObjectIDFromHexOrNil(id)
	}
	return objIDs
}

func HexFromObjectIDOrNil(id primitive.ObjectID) string {
	return id.Hex()
}

func HexFromObjectIDsOrNil(ids []primitive.ObjectID) []string {
	hexIds := make([]string, len(ids))
	for i, id := range ids {
		hexIds[i] = HexFromObjectIDOrNil(id)
	}
	return hexIds
}

func BuildQueryWithSoftDelete(query bson.M) bson.M {
	query["deleted_at"] = nil
	return query
}

func BuildScopeQuery(ctx context.Context) (bson.M, error) {
	filter := bson.M{}

	return filter, nil
}

func MergeAFilter(filterA bson.A, filterB bson.A) bson.A {
	for _, v := range filterB {
		filterA = append(filterA, v)
	}
	return filterA
}

func MergeMFilter(filterA bson.M, filterB bson.M) bson.M {
	for k, v := range filterB {
		filterA[k] = v
	}
	return filterA
}

func MergeAToMFilter(filterA bson.M, filterB bson.A) bson.M {
	for _, v := range filterB {
		filterA = MergeMFilter(filterA, v.(bson.M))
	}
	return filterA
}

func MergeMToAFilter(filterA bson.A, filterB bson.M) bson.A {
	for k, v := range filterB {
		filterA = MergeAFilter(filterA, bson.A{bson.M{k: v}})
	}
	return filterA
}

func GetPeriodAndYearFromObjectID(id primitive.ObjectID) (int32, int32) {
	t := id.Timestamp()
	m := float64(t.Month())
	y := int32(t.Year())
	return int32(math.Ceil(m / 3)), y
}

func GetPreciousAndYearFromObjectID(id primitive.ObjectID) (int32, int32) {
	t := id.Timestamp()
	m := t.Month()
	y := int32(t.Year())
	p := int32((m-1)/3 + 1)
	return p, y
}

func GetPeriodAndYearFromTime(t time.Time) (int32, int32) {
	y := int32(t.Year())
	p := int32(math.Ceil(float64(t.Month()) / 3))
	return p, y
}

func SetDeletedAt() bson.M {
	return bson.M{
		"$set": bson.M{
			"deleted_at": primitive.NewDateTimeFromTime(util.Now()),
		},
	}
}

func IsObjectID(id string) bool {
	_, err := primitive.ObjectIDFromHex(id)
	return err == nil
}

func ObjectIDsFromHexs(ids []string) ([]primitive.ObjectID, error) {
	objIDs := make([]primitive.ObjectID, 0)
	for _, id := range ids {
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			continue
		}
		objIDs = append(objIDs, objID)
	}

	if len(objIDs) == 0 {
		return nil, ErrInvalidObjectID
	}

	return objIDs, nil
}
