package datastore_super_smash_bros_4

import (
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

// NewDataStoreFileServerObjectInfo returns a new DataStoreFileServerObjectInfo
func NewDataStoreFileServerObjectInfo() *DataStoreFileServerObjectInfo {
	return &DataStoreFileServerObjectInfo{}
}
