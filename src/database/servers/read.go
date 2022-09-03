package servers

import (
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "GoBun/database/common"
  "GoBun/srb2kart"
  "fmt"
)

func ListServers(collection *mongo.Collection) ([]srb2kart.Srb2kart, error) {
  ctx, cancel := common.MakeContext()
  defer cancel()

  result := []srb2kart.Srb2kart{}

  filter := bson.D{{}}
  cursor, err := collection.Find(ctx, filter)
  if err != nil {
    return result, fmt.Errorf("could not list servers: %s", err)
  }
  defer cursor.Close(ctx)

  var item srb2kart.Srb2kart
  for cursor.Next(ctx) {
    err := cursor.Decode(&item)
    if err != nil {
      fmt.Printf("Failed to decode an server document: %s", err)
    }
    result = append(result, item)
  }

  return result, nil
}

func Server(collection *mongo.Collection, name string) (srb2kart.Srb2kart, error) {
  ctx, cancel := common.MakeContext()
  defer cancel()

  result := srb2kart.Srb2kart{}
  filter := keyFilter{name}
  found := collection.FindOne(ctx, filter)

  err := found.Decode(&result)
  if err != nil {
    return result, fmt.Errorf("could not decode item: %s", err)
  }
  return result, nil
}
