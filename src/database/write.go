package database

import (
  "go.mongodb.org/mongo-driver/mongo"
  addonsDB "GoBun/database/addons" 
)

func (db DBClient) CreateAddonGroup(groupName string, content []string) (error) {
  d := (mongo.Database)(db)
  return addonsDB.SetAddonGroupContent(d.Collection(addonCollection), groupName, content)
}

func (db DBClient) InsertAddonGroup(toGroup string, newGroups []string) (error) {
  d := (mongo.Database)(db)
  return addonsDB.AddAddonGroup(d.Collection(addonCollection), toGroup, newGroups)
}

func (db DBClient) InsertAddonFile(toGroup string, newFiles []string) (error) {
  d := (mongo.Database)(db)
  return addonsDB.AddAddonFile(d.Collection(addonCollection), toGroup, newFiles)
}

func (db DBClient) RemoveItemFromGroup(groupName string, item string) (error) {
  d := (mongo.Database)(db)
  return addonsDB.RemoveItemFromGroup(d.Collection(addonCollection), groupName, item)
}

func (db DBClient) RemoveGroup(groupName string) (error) {
  d := (mongo.Database)(db)
  return addonsDB.RemoveGroup(d.Collection(addonCollection), groupName)
}
