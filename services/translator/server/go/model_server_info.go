/*
 * Translator service between a srb2kart server and json
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ServerInfo struct {

	PacketVersion float32 `json:"PacketVersion,omitempty"`

	Application []float32 `json:"Application,omitempty"`

	Version float32 `json:"Version,omitempty"`

	SubVersion float32 `json:"SubVersion,omitempty"`

	NumberOfPlayer float32 `json:"NumberOfPlayer,omitempty"`

	MaxPlayers float32 `json:"MaxPlayers,omitempty"`

	GameType float32 `json:"GameType,omitempty"`

	ModifiedGame bool `json:"ModifiedGame,omitempty"`

	CheatsEnabled bool `json:"CheatsEnabled,omitempty"`

	KartVars float32 `json:"KartVars,omitempty"`

	FileNeededNum float32 `json:"FileNeededNum,omitempty"`

	Time float32 `json:"Time,omitempty"`

	LevelTime float32 `json:"LevelTime,omitempty"`

	ServerNameRaw string `json:"ServerNameRaw,omitempty"`

	ServerName string `json:"ServerName,omitempty"`

	MapName string `json:"MapName,omitempty"`

	MapTitle string `json:"MapTitle,omitempty"`

	MapMD5 []float32 `json:"MapMD5,omitempty"`

	ActNum float32 `json:"ActNum,omitempty"`

	IsZone float32 `json:"IsZone,omitempty"`

	HttpSource string `json:"HttpSource,omitempty"`

	FileNeeded []FileNeededInner `json:"FileNeeded,omitempty"`
}

// AssertServerInfoRequired checks if the required fields are not zero-ed
func AssertServerInfoRequired(obj ServerInfo) error {
	for _, el := range obj.FileNeeded {
		if err := AssertFileNeededInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseServerInfoRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ServerInfo (e.g. [][]ServerInfo), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseServerInfoRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aServerInfo, ok := obj.(ServerInfo)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertServerInfoRequired(aServerInfo)
	})
}