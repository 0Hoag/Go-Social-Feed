package mongo

import (
	"context"
	"errors"

	"github.com/hoag/go-social-feed/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connect connects to MongoDB.
func Connect(ctx context.Context, opts ClientOptions) (Client, error) {
	cl, err := mongo.Connect(ctx, opts.clo)
	if err != nil {
		return nil, err
	}

	return &mongoClient{cl: cl}, nil
}

//go:generate mockery --name=Database --output=mocks --case=underscore
type Database interface {
	Collection(string) Collection
	Client() Client
	NewObjectID() primitive.ObjectID
}

//go:generate mockery --name=Collection --output=mocks --case=underscore
type Collection interface {
	FindOne(ctx context.Context, filter interface{}) SingleResult
	InsertOne(ctx context.Context, document interface{}) (interface{}, error)
	InsertMany(ctx context.Context, document []interface{}) ([]interface{}, error)
	DeleteOne(ctx context.Context, filter interface{}) (int64, error)
	DeleteMany(ctx context.Context, filter interface{}) (int64, error)
	DeleteSoftOne(ctx context.Context, filter interface{}) (int64, error)
	DeleteSoftMany(ctx context.Context, filter interface{}) (int64, error)
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error)
	CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error)
	Aggregate(ctx context.Context, pipeline interface{}) (Cursor, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}

//go:generate mockery --name=SingleResult --output=mocks --case=underscore
type SingleResult interface {
	Decode(interface{}) error
}

//go:generate mockery --name=Cursor --output=mocks --case=underscore
type Cursor interface {
	Close(context.Context) error
	Next(context.Context) bool
	Decode(interface{}) error
	All(context.Context, interface{}) error
	One(context.Context, interface{}) error
}

//go:generate mockery --name=Client --output=mocks --case=underscore
type Client interface {
	Database(string) Database
	Disconnect(context.Context) error
	StartSession() (mongo.Session, error)
	UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error
	Ping(context.Context) error
}

type mongoClient struct {
	cl *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}
type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoCursor struct {
	mc *mongo.Cursor
}

type mongoSession struct {
	mongo.Session
}

func (mc *mongoClient) Ping(ctx context.Context) error {
	return mc.cl.Ping(ctx, readpref.Primary())
}

func (mc *mongoClient) Database(dbName string) Database {
	db := mc.cl.Database(dbName)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return mc.cl.UseSession(ctx, fn)
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.cl.StartSession()
	return &mongoSession{session}, err
}

func (mc *mongoClient) Disconnect(ctx context.Context) error {
	return mc.cl.Disconnect(ctx)
}

func (md *mongoDatabase) Collection(colName string) Collection {
	collection := md.db.Collection(colName)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) NewObjectID() primitive.ObjectID {
	return primitive.NewObjectID()
}

func (md *mongoDatabase) Client() Client {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.coll.UpdateOne(ctx, filter, update, opts[:]...)
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	id, err := mc.coll.InsertOne(ctx, document)
	return id.InsertedID, err
}

func (mc *mongoCollection) InsertMany(ctx context.Context, document []interface{}) ([]interface{}, error) {
	res, err := mc.coll.InsertMany(ctx, document)
	return res.InsertedIDs, err
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := mc.coll.DeleteOne(ctx, filter)
	return count.DeletedCount, err
}

func (mc *mongoCollection) DeleteMany(ctx context.Context, filter interface{}) (int64, error) {
	count, err := mc.coll.DeleteMany(ctx, filter)
	return count.DeletedCount, err
}

func (mc *mongoCollection) DeleteSoftOne(ctx context.Context, filter interface{}) (int64, error) {
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deleted_at", Value: util.Now()},
		}},
	}

	count, err := mc.coll.UpdateOne(ctx, filter, update)
	return count.ModifiedCount, err
}

func (mc *mongoCollection) DeleteSoftMany(ctx context.Context, filter interface{}) (int64, error) {
	if mc == nil || mc.coll == nil {
		return 0, errors.New("mongodb collection is nil")
	}

	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "deleted_at", Value: util.Now()},
		}},
	}

	count, err := mc.coll.UpdateMany(ctx, filter, update)
	if err != nil {
		return 0, err
	}

	if count == nil {
		return 0, errors.New("update result is nil")
	}

	return count.ModifiedCount, err
}

func (mc *mongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	findResult, err := mc.coll.Find(ctx, filter, opts...)
	return &mongoCursor{mc: findResult}, err
}

func (mc *mongoCollection) Aggregate(ctx context.Context, pipeline interface{}) (Cursor, error) {
	aggregateResult, err := mc.coll.Aggregate(ctx, pipeline)
	return &mongoCursor{mc: aggregateResult}, err
}

func (mc *mongoCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.coll.UpdateMany(ctx, filter, update, opts[:]...)
}

func (mc *mongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return mc.coll.CountDocuments(ctx, filter, opts...)
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}

func (mr *mongoCursor) Close(ctx context.Context) error {
	return mr.mc.Close(ctx)
}

func (mr *mongoCursor) Next(ctx context.Context) bool {
	return mr.mc.Next(ctx)
}

func (mr *mongoCursor) Decode(v interface{}) error {
	return mr.mc.Decode(v)
}

func (mr *mongoCursor) One(ctx context.Context, result interface{}) error {
	if mr.mc.Next(ctx) {
		return mr.mc.Decode(result)
	}
	return mongo.ErrNoDocuments
}

func (mr *mongoCursor) All(ctx context.Context, result interface{}) error {
	return mr.mc.All(ctx, result)
}
