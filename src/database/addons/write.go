package addons

import (
	"GoBun/functional/array"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func makeContext() (context.Context, func()) {
	return context.WithTimeout(context.Background(), 10*time.Second)
}

func SetAddonGroupContent(collection *mongo.Collection, groupName string, items []string) error {
	ctx, cancel := makeContext()
	defer cancel()

	filter := groupFilter{groupName}
	groupItems := array.Map(items, func(item string) content {
		return content{
			Kind:  "group",
			Value: item,
		}
	})

	update := bson.D{{"$set", bson.D{{"content", groupItems}}}}

	_, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}

func AddAddonGroup(collection *mongo.Collection, toGroup string, newGroups []string) error {
	return addAddonGroup(collection, toGroup, "group", newGroups)
}

func AddAddonFile(collection *mongo.Collection, toGroup string, newFiles []string) error {
	return addAddonGroup(collection, toGroup, "file", newFiles)
}

func RemoveItemFromGroup(collection *mongo.Collection, groupName string, item string) error {
	ctx, cancel := makeContext()
	defer cancel()

	var filter = groupFilter{groupName}
	var itemFilter = bson.D{{"value", item}}
	var update = bson.D{{"$pull", bson.D{{"content", itemFilter}}}}
	_, err := collection.UpdateMany(ctx, filter, update)
	return err
}

func RemoveGroup(collection *mongo.Collection, groupName string) error {
	ctx, cancel := makeContext()
	defer cancel()
	var filter = groupFilter{groupName}
	_, err := collection.DeleteOne(ctx, filter)
	return err
}

func addAddonGroup(collection *mongo.Collection, toGroup, kind string, items []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := groupFilter{toGroup}
	toAdd := array.Map(items, func(item string) content {
		return content{
			Kind:  kind,
			Value: item,
		}
	})
	update := bson.D{{"$push", bson.D{{"content", bson.D{{"$each", toAdd}}}}}}

	_, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return err
}
