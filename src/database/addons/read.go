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
		result.Items = append(result.Items, addons.Addon{item.name})
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
	err = found.Decode(&item)
	if err != nil {
		return result, fmt.Errorf("Error decoding addon group %s: %s", addonGroup, err)
	}
	result.GroupName = item.name

	result.Items = make([]addons.AddonCollection, len(item.content))

	for i, v := range item.content {
		if v.kind == groupType {
			var group, err = _addonGroup(ctx, collection, v.value)
			if err != nil {
				return result, fmt.Errorf("Error while decoding item %s for group %s: %s", v.value, addonGroup, err)
			}
			result.Items[i] = group
		} else if v.kind == fileType {
			result.Items[i] = contentToCollection(v)
		} else {
			return result, fmt.Errorf("Incorrect type of item %s for addon group %s", v.value, addonGroup)
		}
	}

	return result, nil
}
