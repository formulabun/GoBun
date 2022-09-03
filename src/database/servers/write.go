package servers

import (
	"GoBun/database/common"
	"GoBun/srb2kart"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetServer(collection *mongo.Collection, server srb2kart.Srb2kart) error {
	ctx, cancel := common.MakeContext()
	defer cancel()
	filter := makeKeyFilter(server)

	_, err := collection.UpdateOne(ctx, filter, bson.D{{"$set", server}}, options.Update().SetUpsert(true))
	return err
}

func RemoveServer(collection *mongo.Collection, name string) error {
	ctx, cancel := common.MakeContext()
	defer cancel()
	filter := keyFilter{name}

	_, err := collection.DeleteOne(ctx, filter)
	return err
}
