package addons

import (
	"GoBun/srb2kart/addons"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func ListAddonGroups(collection *mongo.Collection) (result addons.AddonGroup, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return addons.AddonGroup{}, fmt.Errorf("Could not search for groups: %s", err)
	}
	defer cursor.Close(ctx)

	result = addons.AddonGroup{
		GroupName: "Database Result",
	}
	for cursor.Next(ctx) {
		var item group
		err := cursor.Decode(&item)
		if err != nil {
			fmt.Printf("Failed at decoding while listing all groups: %s", err)
		}
		result.Items = append(result.Items, addons.Addon{item.Name})
	}

	return result, nil
}

func AddonGroup(collection *mongo.Collection, addonGroup string) (result addons.AddonGroup, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	return _addonGroup(ctx, collection, addonGroup)
}

func _addonGroup(ctx context.Context, collection *mongo.Collection, addonGroup string) (result addons.AddonGroup, err error) {
	result = addons.AddonGroup{}
	found := collection.FindOne(ctx, groupFilter{addonGroup})

	var item group
	found.Decode(&item)
	result.GroupName = item.Name

	result.Items = make([]addons.AddonCollection, len(item.Content))

	for i, v := range item.Content {
		if v.Kind == "group" {
			var group, _ = _addonGroup(ctx, collection, v.Value)
			result.Items[i] = group
		} else {
			result.Items[i] = contentToCollection(v)
		}
	}

	return result, nil
}
