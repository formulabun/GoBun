package database

import (
	addonsDB "GoBun/database/addons"
	serversDB "GoBun/database/servers"
	"GoBun/srb2kart"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DBClient) CreateAddonGroup(groupName string, content []string) error {
	d := (*mongo.Database)(db)
	return addonsDB.SetAddonGroupContent(d.Collection(addonCollection), groupName, content)
}

func (db *DBClient) InsertAddonGroup(toGroup string, newGroups []string) error {
	d := (*mongo.Database)(db)
	return addonsDB.AddAddonGroup(d.Collection(addonCollection), toGroup, newGroups)
}

func (db *DBClient) InsertAddonFile(toGroup string, newFiles []string) error {
	d := (*mongo.Database)(db)
	return addonsDB.AddAddonFile(d.Collection(addonCollection), toGroup, newFiles)
}

func (db *DBClient) RemoveItemFromGroup(groupName string, item string) error {
	d := (*mongo.Database)(db)
	return addonsDB.RemoveItemFromGroup(d.Collection(addonCollection), groupName, item)
}

func (db *DBClient) RemoveGroup(groupName string) error {
	d := (*mongo.Database)(db)
	return addonsDB.RemoveGroup(d.Collection(addonCollection), groupName)
}

func (db *DBClient) SetServer(server srb2kart.Srb2kart) error {
	d := (*mongo.Database)(db)
	return serversDB.SetServer(d.Collection(serverCollection), server)
}

func (db *DBClient) RemoveServer(serverName string) error {
	d := (*mongo.Database)(db)
	return serversDB.RemoveServer(d.Collection(serverCollection), serverName)
}
