package set

import (
	"GoBun/database"
	"GoBun/docker/volume"
	"GoBun/functional/strings"
	"GoBun/scripts/common/subcommand"
	"GoBun/srb2kart"
	"fmt"
	"strconv"
)

var Set = subcommand.Subcommand{"set", []string{"set", "s"}, set}

const (
	portFlag    = "-p"
	volumesFlag = "-v"
	addonFlag   = "-g"
)

func usage() string {
	return fmt.Sprintf("<server name> [%s <port>] [%s <local files volume> <local files volume>] [%s <addon group>]", portFlag, volumesFlag, addonFlag)
}

func parsePort(index *int, args []string) (int, error) {
	p, err := strconv.ParseUint(args[*index], 10, 64)
	*index += 1
	return int(p), err
}

func parseVolumes(index *int, args []string) (volume.VolumeSet, error) {
	res := volume.VolumeSet{[]string{args[*index], args[*index+1]}}
	*index += 2
	return res, nil
}

func parseAddonGroup(index *int, args []string) (string, error) {
	res := args[*index]
	*index += 1
	return res, nil
}

func parseArguments(args []string) (srb2kart.Srb2kart, error) {
	res := srb2kart.DefaultSrb2kart()
	res.Name = args[0]
	var i = 1
	for i < len(args) {
		var err error
		flag := args[i]
		i += 1
		switch flag {
		case portFlag:
			var p int
			p, err = parsePort(&i, args)
			res.Port = p
		case volumesFlag:
			var v volume.VolumeSet
			v, err = parseVolumes(&i, args)
			res.Volumes = v
		case addonFlag:
			var a string
			a, err = parseAddonGroup(&i, args)
			res.AddonGroup = a
		default:
			return res, fmt.Errorf("Unknown flag \"%s\"", flag)
		}
		if err != nil {
			return res, err
		}
	}
	return res, nil
}

func set(args []string) (fmt.Stringer, error) {
	if len(args) == 0 {
		return strings.Stringer{usage()}, nil
	}

	server, err := parseArguments(args)
	if err != nil {
		return nil, err
	}

	client, disconnect, err := database.CreateClient()
	defer disconnect()
	if err != nil {
		return nil, fmt.Errorf("Could create db client: %s", err)
	}

	err = client.SetServer(server)
	if err != nil {
		return nil, fmt.Errorf("Could not create server: %s", err)
	}
	return nil, nil
}
