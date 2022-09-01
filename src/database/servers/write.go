package servers

import (
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
  "GoBun/srb2kart"
  "GoBun/database/common"
)

func SetServer(collection *mongo.Collection, server srb2kart.Srb2kart) (error) {
  ctx, cancel := common.MakeContext()
  defer cancel()
  filter := makeKeyFilter(server)

  _, err := collection.UpdateOne(ctx, filter, bson.D{{"$set", server}}, options.Update().SetUpsert(true))
  return err
}
