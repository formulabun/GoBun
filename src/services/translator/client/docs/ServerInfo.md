# ServerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PacketVersion** | Pointer to **float32** |  | [optional] 
**Application** | Pointer to **[]float32** |  | [optional] 
**Version** | Pointer to **float32** |  | [optional] 
**SubVersion** | Pointer to **float32** |  | [optional] 
**NumberOfPlayer** | Pointer to **float32** |  | [optional] 
**MaxPlayers** | Pointer to **float32** |  | [optional] 
**GameType** | Pointer to **float32** |  | [optional] 
**ModifiedGame** | Pointer to **bool** |  | [optional] 
**CheatsEnabled** | Pointer to **bool** |  | [optional] 
**KartVars** | Pointer to **float32** |  | [optional] 
**FileNeededNum** | Pointer to **float32** |  | [optional] 
**Time** | Pointer to **float32** |  | [optional] 
**LevelTime** | Pointer to **float32** |  | [optional] 
**ServerNameRaw** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**MapName** | Pointer to **string** |  | [optional] 
**MapTitle** | Pointer to **string** |  | [optional] 
**MapMD5** | Pointer to **[]float32** |  | [optional] 
**ActNum** | Pointer to **float32** |  | [optional] 
**IsZone** | Pointer to **float32** |  | [optional] 
**HttpSource** | Pointer to **string** |  | [optional] 
**FileNeeded** | Pointer to [**[]FileNeededInner**](FileNeededInner.md) |  | [optional] 

## Methods

### NewServerInfo

`func NewServerInfo() *ServerInfo`

NewServerInfo instantiates a new ServerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerInfoWithDefaults

`func NewServerInfoWithDefaults() *ServerInfo`

NewServerInfoWithDefaults instantiates a new ServerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPacketVersion

`func (o *ServerInfo) GetPacketVersion() float32`

GetPacketVersion returns the PacketVersion field if non-nil, zero value otherwise.

### GetPacketVersionOk

`func (o *ServerInfo) GetPacketVersionOk() (*float32, bool)`

GetPacketVersionOk returns a tuple with the PacketVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPacketVersion

`func (o *ServerInfo) SetPacketVersion(v float32)`

SetPacketVersion sets PacketVersion field to given value.

### HasPacketVersion

`func (o *ServerInfo) HasPacketVersion() bool`

HasPacketVersion returns a boolean if a field has been set.

### GetApplication

`func (o *ServerInfo) GetApplication() []float32`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ServerInfo) GetApplicationOk() (*[]float32, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ServerInfo) SetApplication(v []float32)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ServerInfo) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetVersion

`func (o *ServerInfo) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ServerInfo) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ServerInfo) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ServerInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSubVersion

`func (o *ServerInfo) GetSubVersion() float32`

GetSubVersion returns the SubVersion field if non-nil, zero value otherwise.

### GetSubVersionOk

`func (o *ServerInfo) GetSubVersionOk() (*float32, bool)`

GetSubVersionOk returns a tuple with the SubVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubVersion

`func (o *ServerInfo) SetSubVersion(v float32)`

SetSubVersion sets SubVersion field to given value.

### HasSubVersion

`func (o *ServerInfo) HasSubVersion() bool`

HasSubVersion returns a boolean if a field has been set.

### GetNumberOfPlayer

`func (o *ServerInfo) GetNumberOfPlayer() float32`

GetNumberOfPlayer returns the NumberOfPlayer field if non-nil, zero value otherwise.

### GetNumberOfPlayerOk

`func (o *ServerInfo) GetNumberOfPlayerOk() (*float32, bool)`

GetNumberOfPlayerOk returns a tuple with the NumberOfPlayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfPlayer

`func (o *ServerInfo) SetNumberOfPlayer(v float32)`

SetNumberOfPlayer sets NumberOfPlayer field to given value.

### HasNumberOfPlayer

`func (o *ServerInfo) HasNumberOfPlayer() bool`

HasNumberOfPlayer returns a boolean if a field has been set.

### GetMaxPlayers

`func (o *ServerInfo) GetMaxPlayers() float32`

GetMaxPlayers returns the MaxPlayers field if non-nil, zero value otherwise.

### GetMaxPlayersOk

`func (o *ServerInfo) GetMaxPlayersOk() (*float32, bool)`

GetMaxPlayersOk returns a tuple with the MaxPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPlayers

`func (o *ServerInfo) SetMaxPlayers(v float32)`

SetMaxPlayers sets MaxPlayers field to given value.

### HasMaxPlayers

`func (o *ServerInfo) HasMaxPlayers() bool`

HasMaxPlayers returns a boolean if a field has been set.

### GetGameType

`func (o *ServerInfo) GetGameType() float32`

GetGameType returns the GameType field if non-nil, zero value otherwise.

### GetGameTypeOk

`func (o *ServerInfo) GetGameTypeOk() (*float32, bool)`

GetGameTypeOk returns a tuple with the GameType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameType

`func (o *ServerInfo) SetGameType(v float32)`

SetGameType sets GameType field to given value.

### HasGameType

`func (o *ServerInfo) HasGameType() bool`

HasGameType returns a boolean if a field has been set.

### GetModifiedGame

`func (o *ServerInfo) GetModifiedGame() bool`

GetModifiedGame returns the ModifiedGame field if non-nil, zero value otherwise.

### GetModifiedGameOk

`func (o *ServerInfo) GetModifiedGameOk() (*bool, bool)`

GetModifiedGameOk returns a tuple with the ModifiedGame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedGame

`func (o *ServerInfo) SetModifiedGame(v bool)`

SetModifiedGame sets ModifiedGame field to given value.

### HasModifiedGame

`func (o *ServerInfo) HasModifiedGame() bool`

HasModifiedGame returns a boolean if a field has been set.

### GetCheatsEnabled

`func (o *ServerInfo) GetCheatsEnabled() bool`

GetCheatsEnabled returns the CheatsEnabled field if non-nil, zero value otherwise.

### GetCheatsEnabledOk

`func (o *ServerInfo) GetCheatsEnabledOk() (*bool, bool)`

GetCheatsEnabledOk returns a tuple with the CheatsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheatsEnabled

`func (o *ServerInfo) SetCheatsEnabled(v bool)`

SetCheatsEnabled sets CheatsEnabled field to given value.

### HasCheatsEnabled

`func (o *ServerInfo) HasCheatsEnabled() bool`

HasCheatsEnabled returns a boolean if a field has been set.

### GetKartVars

`func (o *ServerInfo) GetKartVars() float32`

GetKartVars returns the KartVars field if non-nil, zero value otherwise.

### GetKartVarsOk

`func (o *ServerInfo) GetKartVarsOk() (*float32, bool)`

GetKartVarsOk returns a tuple with the KartVars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKartVars

`func (o *ServerInfo) SetKartVars(v float32)`

SetKartVars sets KartVars field to given value.

### HasKartVars

`func (o *ServerInfo) HasKartVars() bool`

HasKartVars returns a boolean if a field has been set.

### GetFileNeededNum

`func (o *ServerInfo) GetFileNeededNum() float32`

GetFileNeededNum returns the FileNeededNum field if non-nil, zero value otherwise.

### GetFileNeededNumOk

`func (o *ServerInfo) GetFileNeededNumOk() (*float32, bool)`

GetFileNeededNumOk returns a tuple with the FileNeededNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNeededNum

`func (o *ServerInfo) SetFileNeededNum(v float32)`

SetFileNeededNum sets FileNeededNum field to given value.

### HasFileNeededNum

`func (o *ServerInfo) HasFileNeededNum() bool`

HasFileNeededNum returns a boolean if a field has been set.

### GetTime

`func (o *ServerInfo) GetTime() float32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ServerInfo) GetTimeOk() (*float32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ServerInfo) SetTime(v float32)`

SetTime sets Time field to given value.

### HasTime

`func (o *ServerInfo) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetLevelTime

`func (o *ServerInfo) GetLevelTime() float32`

GetLevelTime returns the LevelTime field if non-nil, zero value otherwise.

### GetLevelTimeOk

`func (o *ServerInfo) GetLevelTimeOk() (*float32, bool)`

GetLevelTimeOk returns a tuple with the LevelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelTime

`func (o *ServerInfo) SetLevelTime(v float32)`

SetLevelTime sets LevelTime field to given value.

### HasLevelTime

`func (o *ServerInfo) HasLevelTime() bool`

HasLevelTime returns a boolean if a field has been set.

### GetServerNameRaw

`func (o *ServerInfo) GetServerNameRaw() string`

GetServerNameRaw returns the ServerNameRaw field if non-nil, zero value otherwise.

### GetServerNameRawOk

`func (o *ServerInfo) GetServerNameRawOk() (*string, bool)`

GetServerNameRawOk returns a tuple with the ServerNameRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNameRaw

`func (o *ServerInfo) SetServerNameRaw(v string)`

SetServerNameRaw sets ServerNameRaw field to given value.

### HasServerNameRaw

`func (o *ServerInfo) HasServerNameRaw() bool`

HasServerNameRaw returns a boolean if a field has been set.

### GetServerName

`func (o *ServerInfo) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ServerInfo) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ServerInfo) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ServerInfo) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetMapName

`func (o *ServerInfo) GetMapName() string`

GetMapName returns the MapName field if non-nil, zero value otherwise.

### GetMapNameOk

`func (o *ServerInfo) GetMapNameOk() (*string, bool)`

GetMapNameOk returns a tuple with the MapName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapName

`func (o *ServerInfo) SetMapName(v string)`

SetMapName sets MapName field to given value.

### HasMapName

`func (o *ServerInfo) HasMapName() bool`

HasMapName returns a boolean if a field has been set.

### GetMapTitle

`func (o *ServerInfo) GetMapTitle() string`

GetMapTitle returns the MapTitle field if non-nil, zero value otherwise.

### GetMapTitleOk

`func (o *ServerInfo) GetMapTitleOk() (*string, bool)`

GetMapTitleOk returns a tuple with the MapTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapTitle

`func (o *ServerInfo) SetMapTitle(v string)`

SetMapTitle sets MapTitle field to given value.

### HasMapTitle

`func (o *ServerInfo) HasMapTitle() bool`

HasMapTitle returns a boolean if a field has been set.

### GetMapMD5

`func (o *ServerInfo) GetMapMD5() []float32`

GetMapMD5 returns the MapMD5 field if non-nil, zero value otherwise.

### GetMapMD5Ok

`func (o *ServerInfo) GetMapMD5Ok() (*[]float32, bool)`

GetMapMD5Ok returns a tuple with the MapMD5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapMD5

`func (o *ServerInfo) SetMapMD5(v []float32)`

SetMapMD5 sets MapMD5 field to given value.

### HasMapMD5

`func (o *ServerInfo) HasMapMD5() bool`

HasMapMD5 returns a boolean if a field has been set.

### GetActNum

`func (o *ServerInfo) GetActNum() float32`

GetActNum returns the ActNum field if non-nil, zero value otherwise.

### GetActNumOk

`func (o *ServerInfo) GetActNumOk() (*float32, bool)`

GetActNumOk returns a tuple with the ActNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActNum

`func (o *ServerInfo) SetActNum(v float32)`

SetActNum sets ActNum field to given value.

### HasActNum

`func (o *ServerInfo) HasActNum() bool`

HasActNum returns a boolean if a field has been set.

### GetIsZone

`func (o *ServerInfo) GetIsZone() float32`

GetIsZone returns the IsZone field if non-nil, zero value otherwise.

### GetIsZoneOk

`func (o *ServerInfo) GetIsZoneOk() (*float32, bool)`

GetIsZoneOk returns a tuple with the IsZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsZone

`func (o *ServerInfo) SetIsZone(v float32)`

SetIsZone sets IsZone field to given value.

### HasIsZone

`func (o *ServerInfo) HasIsZone() bool`

HasIsZone returns a boolean if a field has been set.

### GetHttpSource

`func (o *ServerInfo) GetHttpSource() string`

GetHttpSource returns the HttpSource field if non-nil, zero value otherwise.

### GetHttpSourceOk

`func (o *ServerInfo) GetHttpSourceOk() (*string, bool)`

GetHttpSourceOk returns a tuple with the HttpSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpSource

`func (o *ServerInfo) SetHttpSource(v string)`

SetHttpSource sets HttpSource field to given value.

### HasHttpSource

`func (o *ServerInfo) HasHttpSource() bool`

HasHttpSource returns a boolean if a field has been set.

### GetFileNeeded

`func (o *ServerInfo) GetFileNeeded() []FileNeededInner`

GetFileNeeded returns the FileNeeded field if non-nil, zero value otherwise.

### GetFileNeededOk

`func (o *ServerInfo) GetFileNeededOk() (*[]FileNeededInner, bool)`

GetFileNeededOk returns a tuple with the FileNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNeeded

`func (o *ServerInfo) SetFileNeeded(v []FileNeededInner)`

SetFileNeeded sets FileNeeded field to given value.

### HasFileNeeded

`func (o *ServerInfo) HasFileNeeded() bool`

HasFileNeeded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


