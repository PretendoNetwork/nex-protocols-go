package datastore_super_smash_bros_4

import (
	"bytes"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
)

type DataStoreReqGetAdditionalMeta struct {
	nex.Structure
	OwnerID    uint32
	DataType   uint16
	Version    uint16
	MetaBinary []byte
}

// ExtractFromStream extracts a DataStoreReqGetAdditionalMeta structure from a stream
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreReqGetAdditionalMeta.OwnerID = stream.ReadUInt32LE()
	dataStoreReqGetAdditionalMeta.DataType = stream.ReadUInt16LE()
	dataStoreReqGetAdditionalMeta.Version = stream.ReadUInt16LE()

	metaBinary, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	dataStoreReqGetAdditionalMeta.MetaBinary = metaBinary

	return nil
}

// Bytes encodes the DataStoreReqGetAdditionalMeta and returns a byte array
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreReqGetAdditionalMeta.OwnerID)
	stream.WriteUInt16LE(dataStoreReqGetAdditionalMeta.DataType)
	stream.WriteUInt16LE(dataStoreReqGetAdditionalMeta.Version)
	stream.WriteQBuffer(dataStoreReqGetAdditionalMeta.MetaBinary)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqGetAdditionalMeta
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) Copy() nex.StructureInterface {
	copied := NewDataStoreReqGetAdditionalMeta()

	copied.OwnerID = dataStoreReqGetAdditionalMeta.OwnerID
	copied.DataType = dataStoreReqGetAdditionalMeta.DataType
	copied.Version = dataStoreReqGetAdditionalMeta.Version
	copied.MetaBinary = make([]byte, len(dataStoreReqGetAdditionalMeta.MetaBinary))

	copy(copied.MetaBinary, dataStoreReqGetAdditionalMeta.MetaBinary)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetAdditionalMeta *DataStoreReqGetAdditionalMeta) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqGetAdditionalMeta)

	if dataStoreReqGetAdditionalMeta.OwnerID != other.OwnerID {
		return false
	}

	if dataStoreReqGetAdditionalMeta.DataType != other.DataType {
		return false
	}

	if dataStoreReqGetAdditionalMeta.Version != other.Version {
		return false
	}

	if !bytes.Equal(dataStoreReqGetAdditionalMeta.MetaBinary, other.MetaBinary) {
		return false
	}

	return true
}

// NewDataStoreReqGetAdditionalMeta returns a new DataStoreReqGetAdditionalMeta
func NewDataStoreReqGetAdditionalMeta() *DataStoreReqGetAdditionalMeta {
	return &DataStoreReqGetAdditionalMeta{}
}

type DataStorePostProfileParam struct {
	nex.Structure
	Profile []byte
}

// ExtractFromStream extracts a DataStorePostProfileParam structure from a stream
func (dataStorePostProfileParam *DataStorePostProfileParam) ExtractFromStream(stream *nex.StreamIn) error {

	profile, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	dataStorePostProfileParam.Profile = profile

	return nil
}

// Bytes encodes the DataStorePostProfileParam and returns a byte array
func (dataStorePostProfileParam *DataStorePostProfileParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteQBuffer(dataStorePostProfileParam.Profile)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePostProfileParam
func (dataStorePostProfileParam *DataStorePostProfileParam) Copy() nex.StructureInterface {
	copied := NewDataStorePostProfileParam()

	copied.Profile = make([]byte, len(dataStorePostProfileParam.Profile))

	copy(copied.Profile, dataStorePostProfileParam.Profile)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePostProfileParam *DataStorePostProfileParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePostProfileParam)

	if !bytes.Equal(dataStorePostProfileParam.Profile, other.Profile) {
		return false
	}

	return true
}

// NewDataStorePostProfileParam returns a new DataStorePostProfileParam
func NewDataStorePostProfileParam() *DataStorePostProfileParam {
	return &DataStorePostProfileParam{}
}

type DataStoreProfileInfo struct {
	nex.Structure
	Pid     uint32
	Profile []byte
}

// ExtractFromStream extracts a DataStoreProfileInfo structure from a stream
func (dataStoreProfileInfo *DataStoreProfileInfo) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreProfileInfo.Pid = stream.ReadUInt32LE()

	profile, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	dataStoreProfileInfo.Profile = profile

	return nil
}

// Bytes encodes the DataStoreProfileInfo and returns a byte array
func (dataStoreProfileInfo *DataStoreProfileInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreProfileInfo.Pid)
	stream.WriteQBuffer(dataStoreProfileInfo.Profile)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreProfileInfo
func (dataStoreProfileInfo *DataStoreProfileInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreProfileInfo()

	copied.Pid = dataStoreProfileInfo.Pid
	copied.Profile = make([]byte, len(dataStoreProfileInfo.Profile))

	copy(copied.Profile, dataStoreProfileInfo.Profile)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreProfileInfo *DataStoreProfileInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreProfileInfo)

	if dataStoreProfileInfo.Pid != other.Pid {
		return false
	}

	if !bytes.Equal(dataStoreProfileInfo.Profile, other.Profile) {
		return false
	}

	return true
}

// NewDataStoreProfileInfo returns a new DataStoreProfileInfo
func NewDataStoreProfileInfo() *DataStoreProfileInfo {
	return &DataStoreProfileInfo{}
}

type DataStoreReplayPlayer struct {
	nex.Structure
	Fighter     uint8
	Health      uint8
	WinningRate uint16
	Color       uint8
	Color2      uint8
	PrincipalID uint32
	Country     uint32
	Region      uint8
	Number      uint8
}

// ExtractFromStream extracts a DataStoreReplayPlayer structure from a stream
func (dataStoreReplayPlayer *DataStoreReplayPlayer) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreReplayPlayer.Fighter = stream.ReadUInt8()
	dataStoreReplayPlayer.Health = stream.ReadUInt8()
	dataStoreReplayPlayer.WinningRate = stream.ReadUInt16LE()
	dataStoreReplayPlayer.Color = stream.ReadUInt8()
	dataStoreReplayPlayer.Color2 = stream.ReadUInt8()
	dataStoreReplayPlayer.PrincipalID = stream.ReadUInt32LE()
	dataStoreReplayPlayer.Country = stream.ReadUInt32LE()
	dataStoreReplayPlayer.Region = stream.ReadUInt8()
	dataStoreReplayPlayer.Number = stream.ReadUInt8()

	return nil
}

// Bytes encodes the DataStoreReplayPlayer and returns a byte array
func (dataStoreReplayPlayer *DataStoreReplayPlayer) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStoreReplayPlayer.Fighter)
	stream.WriteUInt8(dataStoreReplayPlayer.Health)
	stream.WriteUInt16LE(dataStoreReplayPlayer.WinningRate)
	stream.WriteUInt8(dataStoreReplayPlayer.Color)
	stream.WriteUInt8(dataStoreReplayPlayer.Color2)
	stream.WriteUInt32LE(dataStoreReplayPlayer.PrincipalID)
	stream.WriteUInt32LE(dataStoreReplayPlayer.Country)
	stream.WriteUInt8(dataStoreReplayPlayer.Region)
	stream.WriteUInt8(dataStoreReplayPlayer.Number)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReplayPlayer
func (dataStoreReplayPlayer *DataStoreReplayPlayer) Copy() nex.StructureInterface {
	copied := NewDataStoreReplayPlayer()

	copied.Fighter = dataStoreReplayPlayer.Fighter
	copied.Health = dataStoreReplayPlayer.Health
	copied.WinningRate = dataStoreReplayPlayer.WinningRate
	copied.Color = dataStoreReplayPlayer.Color
	copied.Color2 = dataStoreReplayPlayer.Color2
	copied.PrincipalID = dataStoreReplayPlayer.PrincipalID
	copied.Country = dataStoreReplayPlayer.Country
	copied.Region = dataStoreReplayPlayer.Region
	copied.Number = dataStoreReplayPlayer.Number

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReplayPlayer *DataStoreReplayPlayer) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReplayPlayer)

	if dataStoreReplayPlayer.Fighter != other.Fighter {
		return false
	}

	if dataStoreReplayPlayer.Health != other.Health {
		return false
	}

	if dataStoreReplayPlayer.WinningRate != other.WinningRate {
		return false
	}

	if dataStoreReplayPlayer.Color != other.Color {
		return false
	}

	if dataStoreReplayPlayer.Color2 != other.Color2 {
		return false
	}

	if dataStoreReplayPlayer.PrincipalID != other.PrincipalID {
		return false
	}

	if dataStoreReplayPlayer.Country != other.Country {
		return false
	}

	if dataStoreReplayPlayer.Region != other.Region {
		return false
	}

	if dataStoreReplayPlayer.Number != other.Number {
		return false
	}

	return true
}

// NewDataStoreReplayPlayer returns a new DataStoreReplayPlayer
func NewDataStoreReplayPlayer() *DataStoreReplayPlayer {
	return &DataStoreReplayPlayer{}
}

type DataStoreReplayMetaInfo struct {
	nex.Structure
	ReplayID   uint64
	Size       uint32
	Mode       uint8
	Style      uint8
	Rule       uint8
	Stage      uint8
	ReplayType uint8
	Players    []*DataStoreReplayPlayer
	Winners    []uint32
}

// ExtractFromStream extracts a DataStoreReplayMetaInfo structure from a stream
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreReplayMetaInfo.ReplayID = stream.ReadUInt64LE()
	dataStoreReplayMetaInfo.Size = stream.ReadUInt32LE()
	dataStoreReplayMetaInfo.Mode = stream.ReadUInt8()
	dataStoreReplayMetaInfo.Style = stream.ReadUInt8()
	dataStoreReplayMetaInfo.Rule = stream.ReadUInt8()
	dataStoreReplayMetaInfo.Stage = stream.ReadUInt8()
	dataStoreReplayMetaInfo.ReplayType = stream.ReadUInt8()

	players, err := stream.ReadListStructure(NewDataStoreReplayPlayer())
	if err != nil {
		return err
	}

	dataStoreReplayMetaInfo.Players = players.([]*DataStoreReplayPlayer)
	dataStoreReplayMetaInfo.Winners = stream.ReadListUInt32LE()

	return nil
}

// Bytes encodes the DataStoreReplayMetaInfo and returns a byte array
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreReplayMetaInfo.ReplayID)
	stream.WriteUInt32LE(dataStoreReplayMetaInfo.Size)
	stream.WriteUInt8(dataStoreReplayMetaInfo.Mode)
	stream.WriteUInt8(dataStoreReplayMetaInfo.Style)
	stream.WriteUInt8(dataStoreReplayMetaInfo.Rule)
	stream.WriteUInt8(dataStoreReplayMetaInfo.Stage)
	stream.WriteUInt8(dataStoreReplayMetaInfo.ReplayType)
	stream.WriteListStructure(dataStoreReplayMetaInfo.Players)
	stream.WriteListUInt32LE(dataStoreReplayMetaInfo.Winners)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReplayMetaInfo
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReplayMetaInfo()

	copied.ReplayID = dataStoreReplayMetaInfo.ReplayID
	copied.Size = dataStoreReplayMetaInfo.Size
	copied.Mode = dataStoreReplayMetaInfo.Mode
	copied.Style = dataStoreReplayMetaInfo.Style
	copied.Rule = dataStoreReplayMetaInfo.Rule
	copied.Stage = dataStoreReplayMetaInfo.Stage
	copied.ReplayType = dataStoreReplayMetaInfo.ReplayType
	copied.Players = make([]*DataStoreReplayPlayer, len(dataStoreReplayMetaInfo.Players))

	for i := 0; i < len(dataStoreReplayMetaInfo.Players); i++ {
		copied.Players[i] = dataStoreReplayMetaInfo.Players[i].Copy().(*DataStoreReplayPlayer)
	}

	copied.Winners = make([]uint32, len(dataStoreReplayMetaInfo.Winners))

	copy(copied.Winners, dataStoreReplayMetaInfo.Winners)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReplayMetaInfo *DataStoreReplayMetaInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReplayMetaInfo)

	if dataStoreReplayMetaInfo.ReplayID != other.ReplayID {
		return false
	}

	if dataStoreReplayMetaInfo.Size != other.Size {
		return false
	}

	if dataStoreReplayMetaInfo.Mode != other.Mode {
		return false
	}

	if dataStoreReplayMetaInfo.Style != other.Style {
		return false
	}

	if dataStoreReplayMetaInfo.Rule != other.Rule {
		return false
	}

	if dataStoreReplayMetaInfo.Stage != other.Stage {
		return false
	}

	if dataStoreReplayMetaInfo.ReplayType != other.ReplayType {
		return false
	}

	if len(dataStoreReplayMetaInfo.Players) != len(other.Players) {
		return false
	}

	for i := 0; i < len(dataStoreReplayMetaInfo.Players); i++ {
		if !dataStoreReplayMetaInfo.Players[i].Equals(other.Players[i]) {
			return false
		}
	}

	if len(dataStoreReplayMetaInfo.Winners) != len(other.Winners) {
		return false
	}

	for i := 0; i < len(dataStoreReplayMetaInfo.Winners); i++ {
		if dataStoreReplayMetaInfo.Players[i] != other.Players[i] {
			return false
		}
	}

	return true
}

// NewDataStoreReplayMetaInfo returns a new DataStoreReplayMetaInfo
func NewDataStoreReplayMetaInfo() *DataStoreReplayMetaInfo {
	return &DataStoreReplayMetaInfo{}
}

type DataStoreGetReplayMetaParam struct {
	nex.Structure
	ReplayID uint64
	MetaType uint8
}

// ExtractFromStream extracts a DataStoreGetReplayMetaParam structure from a stream
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreGetReplayMetaParam.ReplayID = stream.ReadUInt64LE()
	dataStoreGetReplayMetaParam.MetaType = stream.ReadUInt8()

	return nil
}

// Bytes encodes the DataStoreGetReplayMetaParam and returns a byte array
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreGetReplayMetaParam.ReplayID)
	stream.WriteUInt8(dataStoreGetReplayMetaParam.MetaType)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetReplayMetaParam
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetReplayMetaParam()

	copied.ReplayID = dataStoreGetReplayMetaParam.ReplayID
	copied.MetaType = dataStoreGetReplayMetaParam.MetaType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetReplayMetaParam *DataStoreGetReplayMetaParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetReplayMetaParam)

	if dataStoreGetReplayMetaParam.ReplayID != other.ReplayID {
		return false
	}

	if dataStoreGetReplayMetaParam.MetaType != other.MetaType {
		return false
	}

	return true
}

// NewDataStoreGetReplayMetaParam returns a new DataStoreGetReplayMetaParam
func NewDataStoreGetReplayMetaParam() *DataStoreGetReplayMetaParam {
	return &DataStoreGetReplayMetaParam{}
}

type DataStorePrepareGetReplayParam struct {
	nex.Structure
	ReplayID  uint64
	ExtraData []string
}

// ExtractFromStream extracts a DataStorePrepareGetReplayParam structure from a stream
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePrepareGetReplayParam.ReplayID = stream.ReadUInt64LE()
	dataStorePrepareGetReplayParam.ExtraData = stream.ReadListString()

	return nil
}

// Bytes encodes the DataStorePrepareGetReplayParam and returns a byte array
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStorePrepareGetReplayParam.ReplayID)
	stream.WriteListString(dataStorePrepareGetReplayParam.ExtraData)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePrepareGetReplayParam
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) Copy() nex.StructureInterface {
	copied := NewDataStorePrepareGetReplayParam()

	copied.ReplayID = dataStorePrepareGetReplayParam.ReplayID
	copied.ExtraData = make([]string, len(dataStorePrepareGetReplayParam.ExtraData))

	copy(copied.ExtraData, dataStorePrepareGetReplayParam.ExtraData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePrepareGetReplayParam *DataStorePrepareGetReplayParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePrepareGetReplayParam)

	if dataStorePrepareGetReplayParam.ReplayID != other.ReplayID {
		return false
	}

	if len(dataStorePrepareGetReplayParam.ExtraData) != len(other.ExtraData) {
		return false
	}

	for i := 0; i < len(dataStorePrepareGetReplayParam.ExtraData); i++ {
		if dataStorePrepareGetReplayParam.ExtraData[i] != other.ExtraData[i] {
			return false
		}
	}

	return true
}

// NewDataStorePrepareGetReplayParam returns a new DataStorePrepareGetReplayParam
func NewDataStorePrepareGetReplayParam() *DataStorePrepareGetReplayParam {
	return &DataStorePrepareGetReplayParam{}
}

type DataStorePreparePostReplayParam struct {
	nex.Structure
	Size          uint32
	Mode          uint8
	Style         uint8
	Rule          uint8
	Stage         uint8
	ReplayType    uint8
	CompetitionID uint64
	Score         int32
	Players       []*DataStoreReplayPlayer
	Winners       []uint32
	KeyVersion    uint16
	ExtraData     []string
}

// ExtractFromStream extracts a DataStorePreparePostReplayParam structure from a stream
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePreparePostReplayParam.Size = stream.ReadUInt32LE()
	dataStorePreparePostReplayParam.Mode = stream.ReadUInt8()
	dataStorePreparePostReplayParam.Style = stream.ReadUInt8()
	dataStorePreparePostReplayParam.Rule = stream.ReadUInt8()
	dataStorePreparePostReplayParam.Stage = stream.ReadUInt8()
	dataStorePreparePostReplayParam.ReplayType = stream.ReadUInt8()
	dataStorePreparePostReplayParam.CompetitionID = stream.ReadUInt64LE()
	dataStorePreparePostReplayParam.Score = int32(stream.ReadUInt32LE())

	players, err := stream.ReadListStructure(NewDataStoreReplayPlayer())
	if err != nil {
		return err
	}

	dataStorePreparePostReplayParam.Players = players.([]*DataStoreReplayPlayer)

	dataStorePreparePostReplayParam.Winners = stream.ReadListUInt32LE()
	dataStorePreparePostReplayParam.KeyVersion = stream.ReadUInt16LE()
	dataStorePreparePostReplayParam.ExtraData = stream.ReadListString()

	return nil
}

// Bytes encodes the DataStorePreparePostReplayParam and returns a byte array
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStorePreparePostReplayParam.Size)
	stream.WriteUInt8(dataStorePreparePostReplayParam.Mode)
	stream.WriteUInt8(dataStorePreparePostReplayParam.Style)
	stream.WriteUInt8(dataStorePreparePostReplayParam.Rule)
	stream.WriteUInt8(dataStorePreparePostReplayParam.Stage)
	stream.WriteUInt8(dataStorePreparePostReplayParam.ReplayType)
	stream.WriteUInt64LE(dataStorePreparePostReplayParam.CompetitionID)
	stream.WriteInt32LE(dataStorePreparePostReplayParam.Score)
	stream.WriteListStructure(dataStorePreparePostReplayParam.Players)
	stream.WriteListUInt32LE(dataStorePreparePostReplayParam.Winners)
	stream.WriteUInt16LE(dataStorePreparePostReplayParam.KeyVersion)
	stream.WriteListString(dataStorePreparePostReplayParam.ExtraData)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePreparePostReplayParam
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) Copy() nex.StructureInterface {
	copied := NewDataStorePreparePostReplayParam()

	copied.Size = dataStorePreparePostReplayParam.Size
	copied.Mode = dataStorePreparePostReplayParam.Mode
	copied.Style = dataStorePreparePostReplayParam.Style
	copied.Rule = dataStorePreparePostReplayParam.Rule
	copied.Stage = dataStorePreparePostReplayParam.Stage
	copied.ReplayType = dataStorePreparePostReplayParam.ReplayType
	copied.CompetitionID = dataStorePreparePostReplayParam.CompetitionID
	copied.Score = dataStorePreparePostReplayParam.Score
	copied.Players = make([]*DataStoreReplayPlayer, len(dataStorePreparePostReplayParam.Players))

	for i := 0; i < len(dataStorePreparePostReplayParam.Players); i++ {
		copied.Players[i] = dataStorePreparePostReplayParam.Players[i].Copy().(*DataStoreReplayPlayer)
	}

	copied.Winners = make([]uint32, len(dataStorePreparePostReplayParam.Winners))

	copy(copied.Winners, dataStorePreparePostReplayParam.Winners)

	copied.KeyVersion = dataStorePreparePostReplayParam.KeyVersion
	copied.ExtraData = make([]string, len(dataStorePreparePostReplayParam.ExtraData))

	copy(copied.ExtraData, dataStorePreparePostReplayParam.ExtraData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePreparePostReplayParam *DataStorePreparePostReplayParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePreparePostReplayParam)

	if dataStorePreparePostReplayParam.Size != other.Size {
		return false
	}

	if dataStorePreparePostReplayParam.Mode != other.Mode {
		return false
	}

	if dataStorePreparePostReplayParam.Style != other.Style {
		return false
	}

	if dataStorePreparePostReplayParam.Rule != other.Rule {
		return false
	}

	if dataStorePreparePostReplayParam.Stage != other.Stage {
		return false
	}

	if dataStorePreparePostReplayParam.ReplayType != other.ReplayType {
		return false
	}

	if dataStorePreparePostReplayParam.CompetitionID != other.CompetitionID {
		return false
	}

	if dataStorePreparePostReplayParam.Score != other.Score {
		return false
	}

	if len(dataStorePreparePostReplayParam.Players) != len(other.Players) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostReplayParam.Players); i++ {
		if !dataStorePreparePostReplayParam.Players[i].Equals(other.Players[i]) {
			return false
		}
	}

	if len(dataStorePreparePostReplayParam.Winners) != len(other.Winners) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostReplayParam.Winners); i++ {
		if dataStorePreparePostReplayParam.Winners[i] != other.Winners[i] {
			return false
		}
	}

	if dataStorePreparePostReplayParam.KeyVersion != other.KeyVersion {
		return false
	}

	if len(dataStorePreparePostReplayParam.ExtraData) != len(other.ExtraData) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostReplayParam.ExtraData); i++ {
		if dataStorePreparePostReplayParam.ExtraData[i] != other.ExtraData[i] {
			return false
		}
	}

	return true
}

// NewDataStorePreparePostReplayParam returns a new DataStorePreparePostReplayParam
func NewDataStorePreparePostReplayParam() *DataStorePreparePostReplayParam {
	return &DataStorePreparePostReplayParam{}
}

type DataStoreCompletePostReplayParam struct {
	nex.Structure
	ReplayID      uint64
	CompleteParam *datastore.DataStoreCompletePostParam
	PrepareParam  *DataStorePreparePostReplayParam
}

// ExtractFromStream extracts a DataStoreCompletePostReplayParam structure from a stream
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreCompletePostReplayParam.ReplayID = stream.ReadUInt64LE()

	completeParam, err := stream.ReadStructure(datastore.NewDataStoreCompletePostParam())
	if err != nil {
		return err
	}

	dataStoreCompletePostReplayParam.CompleteParam = completeParam.(*datastore.DataStoreCompletePostParam)

	prepareParam, err := stream.ReadStructure(NewDataStorePreparePostReplayParam())
	if err != nil {
		return err
	}

	dataStoreCompletePostReplayParam.PrepareParam = prepareParam.(*DataStorePreparePostReplayParam)

	return nil
}

// Bytes encodes the DataStoreCompletePostReplayParam and returns a byte array
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreCompletePostReplayParam.ReplayID)
	stream.WriteStructure(dataStoreCompletePostReplayParam.CompleteParam)
	stream.WriteStructure(dataStoreCompletePostReplayParam.PrepareParam)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreCompletePostReplayParam
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) Copy() nex.StructureInterface {
	copied := NewDataStoreCompletePostReplayParam()

	copied.ReplayID = dataStoreCompletePostReplayParam.ReplayID
	copied.CompleteParam = dataStoreCompletePostReplayParam.CompleteParam.Copy().(*datastore.DataStoreCompletePostParam)
	copied.PrepareParam = dataStoreCompletePostReplayParam.PrepareParam.Copy().(*DataStorePreparePostReplayParam)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompletePostReplayParam)

	if dataStoreCompletePostReplayParam.ReplayID != other.ReplayID {
		return false
	}

	if !dataStoreCompletePostReplayParam.CompleteParam.Equals(other.CompleteParam) {
		return false
	}

	if !dataStoreCompletePostReplayParam.PrepareParam.Equals(other.PrepareParam) {
		return false
	}

	return true
}

// NewDataStoreCompletePostReplayParam returns a new DataStoreCompletePostReplayParam
func NewDataStoreCompletePostReplayParam() *DataStoreCompletePostReplayParam {
	return &DataStoreCompletePostReplayParam{}
}

type DataStorePreparePostSharedDataParam struct {
	nex.Structure
	DataType   uint8
	Region     uint8
	Attribute1 uint8
	Attribute2 uint8
	Fighter    []byte
	Size       uint32
	Comment    string
	MetaBinary []byte
	ExtraData  []string
}

// ExtractFromStream extracts a DataStorePreparePostSharedDataParam structure from a stream
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePreparePostSharedDataParam.DataType = stream.ReadUInt8()
	dataStorePreparePostSharedDataParam.Region = stream.ReadUInt8()
	dataStorePreparePostSharedDataParam.Attribute1 = stream.ReadUInt8()
	dataStorePreparePostSharedDataParam.Attribute2 = stream.ReadUInt8()

	fighter, err := stream.ReadBuffer()
	if err != nil {
		return err
	}

	dataStorePreparePostSharedDataParam.Fighter = fighter
	dataStorePreparePostSharedDataParam.Size = stream.ReadUInt32LE()

	comment, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStorePreparePostSharedDataParam.Comment = comment

	metaBinary, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	dataStorePreparePostSharedDataParam.MetaBinary = metaBinary
	dataStorePreparePostSharedDataParam.ExtraData = stream.ReadListString()

	return nil
}

// Bytes encodes the DataStorePreparePostSharedDataParam and returns a byte array
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStorePreparePostSharedDataParam.DataType)
	stream.WriteUInt8(dataStorePreparePostSharedDataParam.Region)
	stream.WriteUInt8(dataStorePreparePostSharedDataParam.Attribute1)
	stream.WriteUInt8(dataStorePreparePostSharedDataParam.Attribute2)
	stream.WriteBuffer(dataStorePreparePostSharedDataParam.Fighter)
	stream.WriteUInt32LE(dataStorePreparePostSharedDataParam.Size)
	stream.WriteString(dataStorePreparePostSharedDataParam.Comment)
	stream.WriteQBuffer(dataStorePreparePostSharedDataParam.MetaBinary)
	stream.WriteListString(dataStorePreparePostSharedDataParam.ExtraData)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePreparePostSharedDataParam
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) Copy() nex.StructureInterface {
	copied := NewDataStorePreparePostSharedDataParam()

	copied.DataType = dataStorePreparePostSharedDataParam.DataType
	copied.Region = dataStorePreparePostSharedDataParam.Region
	copied.Attribute1 = dataStorePreparePostSharedDataParam.Attribute1
	copied.Attribute2 = dataStorePreparePostSharedDataParam.Attribute2
	copied.Fighter = make([]byte, len(dataStorePreparePostSharedDataParam.Fighter))

	copy(copied.Fighter, dataStorePreparePostSharedDataParam.Fighter)

	copied.Size = dataStorePreparePostSharedDataParam.Size
	copied.Comment = dataStorePreparePostSharedDataParam.Comment
	copied.MetaBinary = make([]byte, len(dataStorePreparePostSharedDataParam.MetaBinary))

	copy(copied.MetaBinary, dataStorePreparePostSharedDataParam.MetaBinary)

	copied.ExtraData = make([]string, len(dataStorePreparePostSharedDataParam.ExtraData))

	copy(copied.ExtraData, dataStorePreparePostSharedDataParam.ExtraData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePreparePostSharedDataParam *DataStorePreparePostSharedDataParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePreparePostSharedDataParam)

	if dataStorePreparePostSharedDataParam.DataType != other.DataType {
		return false
	}

	if dataStorePreparePostSharedDataParam.Region != other.Region {
		return false
	}

	if dataStorePreparePostSharedDataParam.Attribute1 != other.Attribute1 {
		return false
	}

	if dataStorePreparePostSharedDataParam.Attribute2 != other.Attribute2 {
		return false
	}

	if !bytes.Equal(dataStorePreparePostSharedDataParam.Fighter, other.Fighter) {
		return false
	}

	if dataStorePreparePostSharedDataParam.Size != other.Size {
		return false
	}

	if dataStorePreparePostSharedDataParam.Comment != other.Comment {
		return false
	}

	if !bytes.Equal(dataStorePreparePostSharedDataParam.MetaBinary, other.MetaBinary) {
		return false
	}

	if len(dataStorePreparePostSharedDataParam.ExtraData) != len(other.ExtraData) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostSharedDataParam.ExtraData); i++ {
		if dataStorePreparePostSharedDataParam.ExtraData[i] != other.ExtraData[i] {
			return false
		}
	}

	return true
}

// NewDataStorePreparePostSharedDataParam returns a new DataStorePreparePostSharedDataParam
func NewDataStorePreparePostSharedDataParam() *DataStorePreparePostSharedDataParam {
	return &DataStorePreparePostSharedDataParam{}
}

type DataStoreCompletePostSharedDataParam struct {
	nex.Structure
	DataID        uint64
	CompleteParam *datastore.DataStoreCompletePostParam
	PrepareParam  *DataStorePreparePostSharedDataParam
}

// ExtractFromStream extracts a DataStoreCompletePostSharedDataParam structure from a stream
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreCompletePostSharedDataParam.DataID = stream.ReadUInt64LE()

	completeParam, err := stream.ReadStructure(datastore.NewDataStoreCompletePostParam())
	if err != nil {
		return err
	}

	dataStoreCompletePostSharedDataParam.CompleteParam = completeParam.(*datastore.DataStoreCompletePostParam)

	prepareParam, err := stream.ReadStructure(NewDataStorePreparePostSharedDataParam())
	if err != nil {
		return err
	}

	dataStoreCompletePostSharedDataParam.PrepareParam = prepareParam.(*DataStorePreparePostSharedDataParam)

	return nil
}

// Bytes encodes the DataStoreCompletePostSharedDataParam and returns a byte array
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreCompletePostSharedDataParam.DataID)
	stream.WriteStructure(dataStoreCompletePostSharedDataParam.CompleteParam)
	stream.WriteStructure(dataStoreCompletePostSharedDataParam.PrepareParam)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreCompletePostSharedDataParam
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) Copy() nex.StructureInterface {
	copied := NewDataStoreCompletePostSharedDataParam()

	copied.DataID = dataStoreCompletePostSharedDataParam.DataID
	copied.CompleteParam = dataStoreCompletePostSharedDataParam.CompleteParam.Copy().(*datastore.DataStoreCompletePostParam)
	copied.PrepareParam = dataStoreCompletePostSharedDataParam.PrepareParam.Copy().(*DataStorePreparePostSharedDataParam)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompletePostSharedDataParam)

	if dataStoreCompletePostSharedDataParam.DataID != other.DataID {
		return false
	}

	if !dataStoreCompletePostSharedDataParam.CompleteParam.Equals(other.CompleteParam) {
		return false
	}

	if !dataStoreCompletePostSharedDataParam.PrepareParam.Equals(other.PrepareParam) {
		return false
	}

	return true
}

// NewDataStoreCompletePostSharedDataParam returns a new DataStoreCompletePostSharedDataParam
func NewDataStoreCompletePostSharedDataParam() *DataStoreCompletePostSharedDataParam {
	return &DataStoreCompletePostSharedDataParam{}
}

type DataStoreSearchSharedDataParam struct {
	nex.Structure
	DataType    uint8
	Owner       uint32
	Region      uint8
	Attribute1  uint8
	Attribute2  uint8
	Fighter     uint8
	ResultRange *nex.ResultRange
}

// ExtractFromStream extracts a DataStoreSearchSharedDataParam structure from a stream
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreSearchSharedDataParam.DataType = stream.ReadUInt8()
	dataStoreSearchSharedDataParam.Owner = stream.ReadUInt32LE()
	dataStoreSearchSharedDataParam.Region = stream.ReadUInt8()
	dataStoreSearchSharedDataParam.Attribute1 = stream.ReadUInt8()
	dataStoreSearchSharedDataParam.Attribute2 = stream.ReadUInt8()
	dataStoreSearchSharedDataParam.Fighter = stream.ReadUInt8()

	resultRange, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return err
	}

	dataStoreSearchSharedDataParam.ResultRange = resultRange.(*nex.ResultRange)

	return nil
}

// Bytes encodes the DataStoreSearchSharedDataParam and returns a byte array
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStoreSearchSharedDataParam.DataType)
	stream.WriteUInt32LE(dataStoreSearchSharedDataParam.Owner)
	stream.WriteUInt8(dataStoreSearchSharedDataParam.Region)
	stream.WriteUInt8(dataStoreSearchSharedDataParam.Attribute1)
	stream.WriteUInt8(dataStoreSearchSharedDataParam.Attribute2)
	stream.WriteUInt8(dataStoreSearchSharedDataParam.Fighter)
	stream.WriteStructure(dataStoreSearchSharedDataParam.ResultRange)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreSearchSharedDataParam
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) Copy() nex.StructureInterface {
	copied := NewDataStoreSearchSharedDataParam()

	copied.DataType = dataStoreSearchSharedDataParam.DataType
	copied.Owner = dataStoreSearchSharedDataParam.Owner
	copied.Region = dataStoreSearchSharedDataParam.Region
	copied.Attribute1 = dataStoreSearchSharedDataParam.Attribute1
	copied.Attribute2 = dataStoreSearchSharedDataParam.Attribute2
	copied.Fighter = dataStoreSearchSharedDataParam.Fighter
	copied.ResultRange = dataStoreSearchSharedDataParam.ResultRange.Copy().(*nex.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchSharedDataParam *DataStoreSearchSharedDataParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSearchSharedDataParam)

	if dataStoreSearchSharedDataParam.DataType != other.DataType {
		return false
	}

	if dataStoreSearchSharedDataParam.Owner != other.Owner {
		return false
	}

	if dataStoreSearchSharedDataParam.Region != other.Region {
		return false
	}

	if dataStoreSearchSharedDataParam.Attribute1 != other.Attribute1 {
		return false
	}

	if dataStoreSearchSharedDataParam.Attribute2 != other.Attribute2 {
		return false
	}

	if dataStoreSearchSharedDataParam.Fighter != other.Fighter {
		return false
	}

	if !dataStoreSearchSharedDataParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// NewDataStoreSearchSharedDataParam returns a new DataStoreSearchSharedDataParam
func NewDataStoreSearchSharedDataParam() *DataStoreSearchSharedDataParam {
	return &DataStoreSearchSharedDataParam{}
}

type DataStoreSharedDataInfo struct {
	nex.Structure
	DataID      uint64
	OwnerID     uint32
	DataType    uint8
	Comment     string
	MetaBinary  []byte
	Profile     []byte
	Rating      int64
	CreatedTime *nex.DateTime
	Info        *DataStoreFileServerObjectInfo
}

// ExtractFromStream extracts a DataStoreSharedDataInfo structure from a stream
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreSharedDataInfo.DataID = stream.ReadUInt64LE()
	dataStoreSharedDataInfo.OwnerID = stream.ReadUInt32LE()
	dataStoreSharedDataInfo.DataType = stream.ReadUInt8()

	comment, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreSharedDataInfo.Comment = comment

	metaBinary, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	dataStoreSharedDataInfo.MetaBinary = metaBinary

	profile, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	dataStoreSharedDataInfo.Profile = profile
	dataStoreSharedDataInfo.Rating = int64(stream.ReadUInt64LE())
	dataStoreSharedDataInfo.CreatedTime = nex.NewDateTime(stream.ReadUInt64LE())

	info, err := stream.ReadStructure(NewDataStoreFileServerObjectInfo())
	if err != nil {
		return err
	}

	dataStoreSharedDataInfo.Info = info.(*DataStoreFileServerObjectInfo)

	return nil
}

// Bytes encodes the DataStoreSharedDataInfo and returns a byte array
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreSharedDataInfo.DataID)
	stream.WriteUInt32LE(dataStoreSharedDataInfo.OwnerID)
	stream.WriteUInt8(dataStoreSharedDataInfo.DataType)
	stream.WriteString(dataStoreSharedDataInfo.Comment)
	stream.WriteQBuffer(dataStoreSharedDataInfo.MetaBinary)
	stream.WriteQBuffer(dataStoreSharedDataInfo.Profile)
	stream.WriteInt64LE(dataStoreSharedDataInfo.Rating)
	stream.WriteDateTime(dataStoreSharedDataInfo.CreatedTime)
	stream.WriteStructure(dataStoreSharedDataInfo.Info)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreSharedDataInfo
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreSharedDataInfo()

	copied.DataID = dataStoreSharedDataInfo.DataID
	copied.OwnerID = dataStoreSharedDataInfo.OwnerID
	copied.DataType = dataStoreSharedDataInfo.DataType
	copied.Comment = dataStoreSharedDataInfo.Comment
	copied.MetaBinary = make([]byte, len(dataStoreSharedDataInfo.MetaBinary))

	copy(copied.MetaBinary, dataStoreSharedDataInfo.MetaBinary)

	copied.Profile = make([]byte, len(dataStoreSharedDataInfo.Profile))

	copy(copied.Profile, dataStoreSharedDataInfo.Profile)

	copied.Rating = dataStoreSharedDataInfo.Rating
	copied.CreatedTime = dataStoreSharedDataInfo.CreatedTime.Copy()
	copied.Info = dataStoreSharedDataInfo.Info.Copy().(*DataStoreFileServerObjectInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSharedDataInfo)

	if dataStoreSharedDataInfo.DataType != other.DataType {
		return false
	}

	if dataStoreSharedDataInfo.DataID != other.DataID {
		return false
	}

	if dataStoreSharedDataInfo.OwnerID != other.OwnerID {
		return false
	}

	if dataStoreSharedDataInfo.DataType != other.DataType {
		return false
	}

	if dataStoreSharedDataInfo.Comment != other.Comment {
		return false
	}

	if !bytes.Equal(dataStoreSharedDataInfo.MetaBinary, other.MetaBinary) {
		return false
	}

	if !bytes.Equal(dataStoreSharedDataInfo.Profile, other.Profile) {
		return false
	}

	if dataStoreSharedDataInfo.Rating != other.Rating {
		return false
	}

	if !dataStoreSharedDataInfo.CreatedTime.Equals(other.CreatedTime) {
		return false
	}

	if !dataStoreSharedDataInfo.Info.Equals(other.Info) {
		return false
	}

	return true
}

// NewDataStoreSharedDataInfo returns a new DataStoreSharedDataInfo
func NewDataStoreSharedDataInfo() *DataStoreSharedDataInfo {
	return &DataStoreSharedDataInfo{}
}

type DataStoreSearchReplayParam struct {
	nex.Structure
	Mode        uint8
	Style       uint8
	Fighter     uint8
	ResultRange *nex.ResultRange
}

// ExtractFromStream extracts a DataStoreSearchReplayParam structure from a stream
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreSearchReplayParam.Mode = stream.ReadUInt8()
	dataStoreSearchReplayParam.Style = stream.ReadUInt8()
	dataStoreSearchReplayParam.Fighter = stream.ReadUInt8()

	resultRange, err := stream.ReadStructure(nex.NewResultRange())
	if err != nil {
		return err
	}

	dataStoreSearchReplayParam.ResultRange = resultRange.(*nex.ResultRange)

	return nil
}

// Bytes encodes the DataStoreSearchReplayParam and returns a byte array
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStoreSearchReplayParam.Mode)
	stream.WriteUInt8(dataStoreSearchReplayParam.Style)
	stream.WriteUInt8(dataStoreSearchReplayParam.Fighter)
	stream.WriteStructure(dataStoreSearchReplayParam.ResultRange)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreSearchReplayParam
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) Copy() nex.StructureInterface {
	copied := NewDataStoreSearchReplayParam()

	copied.Mode = dataStoreSearchReplayParam.Mode
	copied.Style = dataStoreSearchReplayParam.Style
	copied.Fighter = dataStoreSearchReplayParam.Fighter
	copied.ResultRange = dataStoreSearchReplayParam.ResultRange.Copy().(*nex.ResultRange)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchReplayParam *DataStoreSearchReplayParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSearchReplayParam)

	if dataStoreSearchReplayParam.Mode != other.Mode {
		return false
	}

	if dataStoreSearchReplayParam.Style != other.Style {
		return false
	}

	if dataStoreSearchReplayParam.Fighter != other.Fighter {
		return false
	}

	if !dataStoreSearchReplayParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	return true
}

// NewDataStoreSearchReplayParam returns a new DataStoreSearchReplayParam
func NewDataStoreSearchReplayParam() *DataStoreSearchReplayParam {
	return &DataStoreSearchReplayParam{}
}

type DataStorePostFightingPowerScoreParam struct {
	nex.Structure
	Mode             uint8
	Score            uint32
	IsWorldHighScore bool
}

// ExtractFromStream extracts a DataStorePostFightingPowerScoreParam structure from a stream
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePostFightingPowerScoreParam.Mode = stream.ReadUInt8()
	dataStorePostFightingPowerScoreParam.Score = stream.ReadUInt32LE()
	dataStorePostFightingPowerScoreParam.IsWorldHighScore = stream.ReadUInt8() == 1

	return nil
}

// Bytes encodes the DataStorePostFightingPowerScoreParam and returns a byte array
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStorePostFightingPowerScoreParam.Mode)
	stream.WriteUInt32LE(dataStorePostFightingPowerScoreParam.Score)
	stream.WriteBool(dataStorePostFightingPowerScoreParam.IsWorldHighScore)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePostFightingPowerScoreParam
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) Copy() nex.StructureInterface {
	copied := NewDataStorePostFightingPowerScoreParam()

	copied.Mode = dataStorePostFightingPowerScoreParam.Mode
	copied.Score = dataStorePostFightingPowerScoreParam.Score
	copied.IsWorldHighScore = dataStorePostFightingPowerScoreParam.IsWorldHighScore

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePostFightingPowerScoreParam *DataStorePostFightingPowerScoreParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePostFightingPowerScoreParam)

	if dataStorePostFightingPowerScoreParam.Mode != other.Mode {
		return false
	}

	if dataStorePostFightingPowerScoreParam.Score != other.Score {
		return false
	}

	if dataStorePostFightingPowerScoreParam.IsWorldHighScore != other.IsWorldHighScore {
		return false
	}

	return true
}

// NewDataStorePostFightingPowerScoreParam returns a new DataStorePostFightingPowerScoreParam
func NewDataStorePostFightingPowerScoreParam() *DataStorePostFightingPowerScoreParam {
	return &DataStorePostFightingPowerScoreParam{}
}

type DataStoreFightingPowerScore struct {
	nex.Structure
	Score uint32
	Rank  uint32
}

// ExtractFromStream extracts a DataStoreFightingPowerScore structure from a stream
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreFightingPowerScore.Score = stream.ReadUInt32LE()
	dataStoreFightingPowerScore.Rank = stream.ReadUInt32LE()

	return nil
}

// Bytes encodes the DataStoreFightingPowerScore and returns a byte array
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreFightingPowerScore.Score)
	stream.WriteUInt32LE(dataStoreFightingPowerScore.Rank)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreFightingPowerScore
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) Copy() nex.StructureInterface {
	copied := NewDataStoreFightingPowerScore()

	copied.Score = dataStoreFightingPowerScore.Score
	copied.Rank = dataStoreFightingPowerScore.Rank

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreFightingPowerScore *DataStoreFightingPowerScore) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreFightingPowerScore)

	if dataStoreFightingPowerScore.Score != other.Score {
		return false
	}

	if dataStoreFightingPowerScore.Rank != other.Rank {
		return false
	}

	return true
}

// NewDataStoreFightingPowerScore returns a new DataStoreFightingPowerScore
func NewDataStoreFightingPowerScore() *DataStoreFightingPowerScore {
	return &DataStoreFightingPowerScore{}
}

type DataStoreFightingPowerChart struct {
	nex.Structure
	UserNum uint32
	Chart   []*DataStoreFightingPowerScore
}

// ExtractFromStream extracts a DataStoreFightingPowerChart structure from a stream
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreFightingPowerChart.UserNum = stream.ReadUInt32LE()

	chart, err := stream.ReadListStructure(NewDataStoreFightingPowerScore())
	if err != nil {
		return err
	}

	dataStoreFightingPowerChart.Chart = chart.([]*DataStoreFightingPowerScore)

	return nil
}

// Bytes encodes the DataStoreFightingPowerChart and returns a byte array
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreFightingPowerChart.UserNum)
	stream.WriteListStructure(dataStoreFightingPowerChart.Chart)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreFightingPowerChart
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) Copy() nex.StructureInterface {
	copied := NewDataStoreFightingPowerChart()

	copied.UserNum = dataStoreFightingPowerChart.UserNum
	copied.Chart = make([]*DataStoreFightingPowerScore, len(dataStoreFightingPowerChart.Chart))

	for i := 0; i < len(dataStoreFightingPowerChart.Chart); i++ {
		copied.Chart[i] = dataStoreFightingPowerChart.Chart[i].Copy().(*DataStoreFightingPowerScore)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreFightingPowerChart *DataStoreFightingPowerChart) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreFightingPowerChart)

	if dataStoreFightingPowerChart.UserNum != other.UserNum {
		return false
	}

	if len(dataStoreFightingPowerChart.Chart) != len(other.Chart) {
		return false
	}

	for i := 0; i < len(dataStoreFightingPowerChart.Chart); i++ {
		if !dataStoreFightingPowerChart.Chart[i].Equals(other.Chart[i]) {
			return false
		}
	}

	return true
}

// NewDataStoreFightingPowerChart returns a new DataStoreFightingPowerChart
func NewDataStoreFightingPowerChart() *DataStoreFightingPowerChart {
	return &DataStoreFightingPowerChart{}
}

// DataStoreFileServerObjectInfo is sent in the GetObjectInfos method
type DataStoreFileServerObjectInfo struct {
	nex.Structure
	DataID  uint64
	GetInfo *datastore.DataStoreReqGetInfo
}

// ExtractFromStream extracts a DataStoreFileServerObjectInfo structure from a stream
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size
	dataStoreFileServerObjectInfo.DataID = stream.ReadUInt64LE()

	getInfo, err := stream.ReadStructure(datastore.NewDataStoreReqGetInfo())
	if err != nil {
		return err
	}

	dataStoreFileServerObjectInfo.GetInfo = getInfo.(*datastore.DataStoreReqGetInfo)

	return nil
}

// Bytes encodes the DataStoreFileServerObjectInfo and returns a byte array
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreFileServerObjectInfo.DataID)
	stream.WriteStructure(dataStoreFileServerObjectInfo.GetInfo)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreFileServerObjectInfo
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreFileServerObjectInfo()

	copied.DataID = dataStoreFileServerObjectInfo.DataID
	copied.GetInfo = dataStoreFileServerObjectInfo.GetInfo.Copy().(*datastore.DataStoreReqGetInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreFileServerObjectInfo *DataStoreFileServerObjectInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreFileServerObjectInfo)

	if dataStoreFileServerObjectInfo.DataID != other.DataID {
		return false
	}

	if !dataStoreFileServerObjectInfo.GetInfo.Equals(other.GetInfo) {
		return false
	}

	return true
}

// NewDataStoreFileServerObjectInfo returns a new DataStoreFileServerObjectInfo
func NewDataStoreFileServerObjectInfo() *DataStoreFileServerObjectInfo {
	return &DataStoreFileServerObjectInfo{}
}
