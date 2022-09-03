package database

import (
	addonsDB "GoBun/database/addons"
	serversDB "GoBun/database/servers"
	"GoBun/srb2kart"
	"GoBun/srb2kart/addons"
	"go.mongodb.org/mongo-driver/mongo"
)

func (db *DBClient) ListAddonGroups() (addons.AddonGroup, error) {
	d := (*mongo.Database)(db)
	return addonsDB.ListAddonGroups(d.Collection(addonCollection))
}

func (db *DBClient) AddonGroup(groupName string) (addons.AddonGroup, error) {
	d := (*mongo.Database)(db)
	return addonsDB.AddonGroup(d.Collection(addonCollection), groupName)
}

func (db *DBClient) ListServers() ([]srb2kart.Srb2kart, error) {
	d := (*mongo.Database)(db)
	return serversDB.ListServers(d.Collection(serverCollection))
}

func (db *DBClient) Server(serverName string) (srb2kart.Srb2kart, error) {
	d := (*mongo.Database)(db)
	return serversDB.Server(d.Collection(serverCollection), serverName)
}
