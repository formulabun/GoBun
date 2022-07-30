package database

import (
  "go.mongodb.org/mongo-driver/mongo"
  addonsDB "GoBun/database/addons" 
  "GoBun/srb2kart/addons"
)

func (db *DBClient)ListAddonGroups() (addons.AddonGroup, error) {
  d := (*mongo.Database)(db)
  return addonsDB.ListAddonGroups(d.Collection("addons"))
}

func (db *DBClient)AddonGroup(groupName string) (addons.AddonGroup, error) {
  d := (*mongo.Database)(db)
  return addonsDB.AddonGroup(d.Collection("addons"), groupName)
}
