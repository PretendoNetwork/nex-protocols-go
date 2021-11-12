package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// DataStoreSmash4ProtocolID is the protocol ID for the DataStore (Smash4) protocol. ID is the same as the DataStore protocol
	DataStoreSmash4ProtocolID = 0x73

	// DataStoreSmash4MethodPostProfile is the method ID for the method PostProfile
	DataStoreSmash4MethodPostProfile = 0x2D

	// DataStoreSmash4MethodGetProfiles is the method ID for the method GetProfiles
	DataStoreSmash4MethodGetProfiles = 0x2E

	// DataStoreSmash4MethodSendPlayReport is the method ID for the method SendPlayReport
	DataStoreSmash4MethodSendPlayReport = 0x2F

	// DataStoreSmash4MethodGetWorldPlayReport is the method ID for the method GetWorldPlayReport
	DataStoreSmash4MethodGetWorldPlayReport = 0x30

	// DataStoreSmash4MethodGetReplayMeta is the method ID for the method GetReplayMeta
	DataStoreSmash4MethodGetReplayMeta = 0x31

	// DataStoreSmash4MethodPrepareGetReplay is the method ID for the method PrepareGetReplay
	DataStoreSmash4MethodPrepareGetReplay = 0x32

	// DataStoreSmash4MethodPreparePostReplay is the method ID for the method PreparePostReplay
	DataStoreSmash4MethodPreparePostReplay = 0x33

	// DataStoreSmash4MethodCompletePostReplay is the method ID for the method CompletePostReplay
	DataStoreSmash4MethodCompletePostReplay = 0x34

	// DataStoreSmash4MethodCheckPostReplay is the method ID for the method CheckPostReplay
	DataStoreSmash4MethodCheckPostReplay = 0x35

	// DataStoreSmash4MethodGetNextReplay is the method ID for the method GetNextReplay
	DataStoreSmash4MethodGetNextReplay = 0x36

	// DataStoreSmash4MethodPreparePostSharedData is the method ID for the method PreparePostSharedData
	DataStoreSmash4MethodPreparePostSharedData = 0x37

	// DataStoreSmash4MethodCompletePostSharedData is the method ID for the method CompletePostSharedData
	DataStoreSmash4MethodCompletePostSharedData = 0x38

	// DataStoreSmash4MethodSearchSharedData is the method ID for the method SearchSharedData
	DataStoreSmash4MethodSearchSharedData = 0x39

	// DataStoreSmash4MethodGetApplicationConfig is the method ID for the method GetApplicationConfig
	DataStoreSmash4MethodGetApplicationConfig = 0x3A

	// DataStoreSmash4MethodSearchReplay is the method ID for the method SearchReplay
	DataStoreSmash4MethodSearchReplay = 0x3B

	// DataStoreSmash4MethodPostFightingPowerScore is the method ID for the method PostFightingPowerScore
	DataStoreSmash4MethodPostFightingPowerScore = 0x3C

	// DataStoreSmash4MethodGetFightingPowerChart is the method ID for the method GetFightingPowerChart
	DataStoreSmash4MethodGetFightingPowerChart = 0x3D

	// DataStoreSmash4MethodGetFightingPowerChartAll is the method ID for the method GetFightingPowerChartAll
	DataStoreSmash4MethodGetFightingPowerChartAll = 0x3E

	// DataStoreSmash4MethodReportSharedData is the method ID for the method ReportSharedData
	DataStoreSmash4MethodReportSharedData = 0x3F
)

// DataStoreSmash4Protocol handles the DataStore (Smash4) nex protocol. Embeds DataStoreProtocol
type DataStoreSmash4Protocol struct {
	server *nex.Server
	DataStoreProtocol

	PostProfileHandler              func(err error, client *nex.Client, callID uint32, param *DataStorePostProfileParam)
	GetProfilesHandler              func(err error, client *nex.Client, callID uint32, pidList []uint32)
	SendPlayReportHandler           func(err error, client *nex.Client, callID uint32, playReport []int32)
	GetWorldPlayReportHandler       func(err error, client *nex.Client, callID uint32)
	GetReplayMetaHandler            func(err error, client *nex.Client, callID uint32, param *DataStoreGetReplayMetaParam)
	PrepareGetReplayHandler         func(err error, client *nex.Client, callID uint32, param *DataStorePrepareGetReplayParam)
	PreparePostReplayHandler        func(err error, client *nex.Client, callID uint32, param *DataStorePreparePostReplayParam)
	CompletePostReplayHandler       func(err error, client *nex.Client, callID uint32, param *DataStoreCompletePostReplayParam)
	CheckPostReplayHandler          func(err error, client *nex.Client, callID uint32, param *DataStorePreparePostReplayParam)
	GetNextReplayHandler            func(err error, client *nex.Client, callID uint32)
	PreparePostSharedDataHandler    func(err error, client *nex.Client, callID uint32, param *DataStorePreparePostSharedDataParam)
	CompletePostSharedDataHandler   func(err error, client *nex.Client, callID uint32, param *DataStoreCompletePostSharedDataParam)
	SearchSharedDataHandler         func(err error, client *nex.Client, callID uint32, param *DataStoreSearchSharedDataParam)
	GetApplicationConfigHandler     func(err error, client *nex.Client, callID uint32, applicationID uint32)
	SearchReplayHandler             func(err error, client *nex.Client, callID uint32, param *DataStoreSearchReplayParam)
	PostFightingPowerScoreHandler   func(err error, client *nex.Client, callID uint32, params []*DataStorePostFightingPowerScoreParam)
	GetFightingPowerChartHandler    func(err error, client *nex.Client, callID uint32, mode uint8)
	GetFightingPowerChartAllHandler func(err error, client *nex.Client, callID uint32)
	ReportSharedDataHandler         func(err error, client *nex.Client, callID uint32, dataID uint64)
}

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

	players, err := ReadListDataStoreReplayPlayer(stream)
	if err != nil {
		return err
	}

	dataStoreReplayMetaInfo.Players = players
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

	players, err := ReadListDataStoreReplayPlayer(stream)
	if err != nil {
		return err
	}

	dataStorePreparePostReplayParam.Players = players

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
	CompleteParam *DataStoreCompletePostParam
	PrepareParam  *DataStorePreparePostReplayParam
}

// ExtractFromStream extracts a DataStoreCompletePostReplayParam structure from a stream
func (dataStoreCompletePostReplayParam *DataStoreCompletePostReplayParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreCompletePostReplayParam.ReplayID = stream.ReadUInt64LE()

	completeParam, err := stream.ReadStructure(NewDataStoreCompletePostParam())
	if err != nil {
		return err
	}

	dataStoreCompletePostReplayParam.CompleteParam = completeParam.(*DataStoreCompletePostParam)

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
	CompleteParam *DataStoreCompletePostParam
	PrepareParam  *DataStorePreparePostSharedDataParam
}

// ExtractFromStream extracts a DataStoreCompletePostSharedDataParam structure from a stream
func (dataStoreCompletePostSharedDataParam *DataStoreCompletePostSharedDataParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreCompletePostSharedDataParam.DataID = stream.ReadUInt64LE()

	completeParam, err := stream.ReadStructure(NewDataStoreCompletePostParam())
	if err != nil {
		return err
	}

	dataStoreCompletePostSharedDataParam.CompleteParam = completeParam.(*DataStoreCompletePostParam)

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
	stream.WriteUInt64LE(dataStoreSharedDataInfo.CreatedTime.Value())
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

	chart, err := ReadListDataStoreFightingPowerScore(stream)
	if err != nil {
		return err
	}

	dataStoreFightingPowerChart.Chart = chart

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

// Setup initializes the protocol
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) Setup() {
	nexServer := dataStoreSmash4Protocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if DataStoreSmash4ProtocolID == request.ProtocolID() {
			// Commented out the ones which aren't implemented in DataStore yet
			switch request.MethodID() {
			case DataStoreMethodPrepareGetObjectV1:
				//go dataStoreSmash4Protocol.handlePrepareGetObjectV1(packet)
			case DataStoreMethodPreparePostObjectV1:
				//go dataStoreSmash4Protocol.handlePreparePostObjectV1(packet)
			case DataStoreMethodCompletePostObjectV1:
				//go dataStoreSmash4Protocol.handleCompletePostObjectV1(packet)
			case DataStoreMethodDeleteObject:
				//go dataStoreSmash4Protocol.handleDeleteObject(packet)
			case DataStoreMethodDeleteObjects:
				//go dataStoreSmash4Protocol.handleDeleteObjects(packet)
			case DataStoreMethodChangeMetaV1:
				//go dataStoreSmash4Protocol.handleChangeMetaV1(packet)
			case DataStoreMethodChangeMetasV1:
				//go dataStoreSmash4Protocol.handleChangeMetasV1(packet)
			case DataStoreMethodGetMeta:
				go dataStoreSmash4Protocol.handleGetMeta(packet)
			case DataStoreMethodGetMetas:
				//go dataStoreSmash4Protocol.handleGetMetas(packet)
			case DataStoreMethodPrepareUpdateObject:
				//go dataStoreSmash4Protocol.handlePrepareUpdateObject(packet)
			case DataStoreMethodCompleteUpdateObject:
				//go dataStoreSmash4Protocol.handleCompleteUpdateObject(packet)
			case DataStoreMethodSearchObject:
				//go dataStoreSmash4Protocol.handleSearchObject(packet)
			case DataStoreMethodGetNotificationURL:
				//go dataStoreSmash4Protocol.handleGetNotificationURL(packet)
			case DataStoreMethodGetNewArrivedNotificationsV1:
				//go dataStoreSmash4Protocol.handleGetNewArrivedNotificationsV1(packet)
			case DataStoreMethodRateObject:
				//go dataStoreSmash4Protocol.handleRateObject(packet)
			case DataStoreMethodGetRating:
				//go dataStoreSmash4Protocol.handleGetRating(packet)
			case DataStoreMethodGetRatings:
				//go dataStoreSmash4Protocol.handleGetRatings(packet)
			case DataStoreMethodResetRating:
				//go dataStoreSmash4Protocol.handleResetRating(packet)
			case DataStoreMethodResetRatings:
				//go dataStoreSmash4Protocol.handleResetRatings(packet)
			case DataStoreMethodGetSpecificMetaV1:
				//go dataStoreSmash4Protocol.handleGetSpecificMetaV1(packet)
			case DataStoreMethodPostMetaBinary:
				//go dataStoreSmash4Protocol.handlePostMetaBinary(packet)
			case DataStoreMethodTouchObject:
				//go dataStoreSmash4Protocol.handleTouchObject(packet)
			case DataStoreMethodGetRatingWithLog:
				//go dataStoreSmash4Protocol.handleGetRatingWithLog(packet)
			case DataStoreMethodPreparePostObject:
				go dataStoreSmash4Protocol.handlePreparePostObject(packet)
			case DataStoreMethodPrepareGetObject:
				go dataStoreSmash4Protocol.handlePrepareGetObject(packet)
			case DataStoreMethodCompletePostObject:
				go dataStoreSmash4Protocol.handleCompletePostObject(packet)
			case DataStoreMethodGetNewArrivedNotifications:
				//go dataStoreSmash4Protocol.handleGetNewArrivedNotifications(packet)
			case DataStoreMethodGetSpecificMeta:
				//go dataStoreSmash4Protocol.handleGetSpecificMeta(packet)
			case DataStoreMethodGetPersistenceInfo:
				//go dataStoreSmash4Protocol.handleGetPersistenceInfo(packet)
			case DataStoreMethodGetPersistenceInfos:
				//go dataStoreSmash4Protocol.handleGetPersistenceInfos(packet)
			case DataStoreMethodPerpetuateObject:
				//go dataStoreSmash4Protocol.handlePerpetuateObject(packet)
			case DataStoreMethodUnperpetuateObject:
				//go dataStoreSmash4Protocol.handleUnperpetuateObject(packet)
			case DataStoreMethodPrepareGetObjectOrMetaBinary:
				//go dataStoreSmash4Protocol.handlePrepareGetObjectOrMetaBinary(packet)
			case DataStoreMethodGetPasswordInfo:
				//go dataStoreSmash4Protocol.handleGetPasswordInfo(packet)
			case DataStoreMethodGetPasswordInfos:
				//go dataStoreSmash4Protocol.handleGetPasswordInfos(packet)
			case DataStoreMethodGetMetasMultipleParam:
				go dataStoreSmash4Protocol.handleGetMetasMultipleParam(packet)
			case DataStoreMethodCompletePostObjects:
				//go dataStoreSmash4Protocol.handleCompletePostObjects(packet)
			case DataStoreMethodChangeMeta:
				go dataStoreSmash4Protocol.handleChangeMeta(packet)
			case DataStoreMethodChangeMetas:
				//go dataStoreSmash4Protocol.handleChangeMetas(packet)
			case DataStoreMethodRateObjects:
				go dataStoreSmash4Protocol.handleRateObjects(packet)
			case DataStoreMethodPostMetaBinaryWithDataID:
				//go dataStoreSmash4Protocol.handlePostMetaBinaryWithDataId(packet)
			case DataStoreMethodPostMetaBinariesWithDataID:
				//go dataStoreSmash4Protocol.handlePostMetaBinariesWithDataId(packet)
			case DataStoreMethodRateObjectWithPosting:
				//go dataStoreSmash4Protocol.handleRateObjectWithPosting(packet)
			case DataStoreMethodRateObjectsWithPosting:
				//go dataStoreSmash4Protocol.handleRateObjectsWithPosting(packet)
			case DataStoreSmash4MethodPostProfile:
				go dataStoreSmash4Protocol.handlePostProfile(packet)
			case DataStoreSmash4MethodGetProfiles:
				go dataStoreSmash4Protocol.handleGetProfiles(packet)
			case DataStoreSmash4MethodSendPlayReport:
				go dataStoreSmash4Protocol.handleSendPlayReport(packet)
			case DataStoreSmash4MethodGetWorldPlayReport:
				go dataStoreSmash4Protocol.handleGetWorldPlayReport(packet)
			case DataStoreSmash4MethodGetReplayMeta:
				go dataStoreSmash4Protocol.handleGetReplayMeta(packet)
			case DataStoreSmash4MethodPrepareGetReplay:
				go dataStoreSmash4Protocol.handlePrepareGetReplay(packet)
			case DataStoreSmash4MethodPreparePostReplay:
				go dataStoreSmash4Protocol.handlePreparePostReplay(packet)
			case DataStoreSmash4MethodCompletePostReplay:
				go dataStoreSmash4Protocol.handleCompletePostReplay(packet)
			case DataStoreSmash4MethodCheckPostReplay:
				go dataStoreSmash4Protocol.handleCheckPostReplay(packet)
			case DataStoreSmash4MethodGetNextReplay:
				go dataStoreSmash4Protocol.handleGetNextReplay(packet)
			case DataStoreSmash4MethodPreparePostSharedData:
				go dataStoreSmash4Protocol.handlePreparePostSharedData(packet)
			case DataStoreSmash4MethodCompletePostSharedData:
				go dataStoreSmash4Protocol.handleCompletePostSharedData(packet)
			case DataStoreSmash4MethodSearchSharedData:
				go dataStoreSmash4Protocol.handleSearchSharedData(packet)
			case DataStoreSmash4MethodGetApplicationConfig:
				go dataStoreSmash4Protocol.handleGetApplicationConfig(packet)
			case DataStoreSmash4MethodSearchReplay:
				go dataStoreSmash4Protocol.handleSearchReplay(packet)
			case DataStoreSmash4MethodPostFightingPowerScore:
				go dataStoreSmash4Protocol.handlePostFightingPowerScore(packet)
			case DataStoreSmash4MethodGetFightingPowerChart:
				go dataStoreSmash4Protocol.handleGetFightingPowerChart(packet)
			case DataStoreSmash4MethodGetFightingPowerChartAll:
				go dataStoreSmash4Protocol.handleGetFightingPowerChartAll(packet)
			case DataStoreSmash4MethodReportSharedData:
				go dataStoreSmash4Protocol.handleReportSharedData(packet)
			default:
				go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
				fmt.Printf("Unsupported DataStoreSmash4 method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// PostProfile sets the PostProfile handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) PostProfile(handler func(err error, client *nex.Client, callID uint32, param *DataStorePostProfileParam)) {
	dataStoreSmash4Protocol.PostProfileHandler = handler
}

// GetProfiles sets the GetProfiles handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) GetProfiles(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	dataStoreSmash4Protocol.GetProfilesHandler = handler
}

// SendPlayReport sets the SendPlayReport handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) SendPlayReport(handler func(err error, client *nex.Client, callID uint32, playReport []int32)) {
	dataStoreSmash4Protocol.SendPlayReportHandler = handler
}

// GetWorldPlayReport sets the GetWorldPlayReport handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) GetWorldPlayReport(handler func(err error, client *nex.Client, callID uint32)) {
	dataStoreSmash4Protocol.GetWorldPlayReportHandler = handler
}

// GetReplayMeta sets the GetReplayMeta handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) GetReplayMeta(handler func(err error, client *nex.Client, callID uint32, param *DataStoreGetReplayMetaParam)) {
	dataStoreSmash4Protocol.GetReplayMetaHandler = handler
}

// PrepareGetReplay sets the PrepareGetReplay handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) PrepareGetReplay(handler func(err error, client *nex.Client, callID uint32, param *DataStorePrepareGetReplayParam)) {
	dataStoreSmash4Protocol.PrepareGetReplayHandler = handler
}

// PreparePostReplay sets the PreparePostReplay handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) PreparePostReplay(handler func(err error, client *nex.Client, callID uint32, param *DataStorePreparePostReplayParam)) {
	dataStoreSmash4Protocol.PreparePostReplayHandler = handler
}

// CompletePostReplay sets the CompletePostReplay handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) CompletePostReplay(handler func(err error, client *nex.Client, callID uint32, param *DataStoreCompletePostReplayParam)) {
	dataStoreSmash4Protocol.CompletePostReplayHandler = handler
}

// CheckPostReplay sets the CheckPostReplay handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) CheckPostReplay(handler func(err error, client *nex.Client, callID uint32, param *DataStorePreparePostReplayParam)) {
	dataStoreSmash4Protocol.CheckPostReplayHandler = handler
}

// GetNextReplay sets the GetNextReplay handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) GetNextReplay(handler func(err error, client *nex.Client, callID uint32)) {
	dataStoreSmash4Protocol.GetNextReplayHandler = handler
}

// PreparePostSharedData sets the PreparePostSharedData handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) PreparePostSharedData(handler func(err error, client *nex.Client, callID uint32, param *DataStorePreparePostSharedDataParam)) {
	dataStoreSmash4Protocol.PreparePostSharedDataHandler = handler
}

// CompletePostSharedData sets the CompletePostSharedData handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) CompletePostSharedData(handler func(err error, client *nex.Client, callID uint32, param *DataStoreCompletePostSharedDataParam)) {
	dataStoreSmash4Protocol.CompletePostSharedDataHandler = handler
}

// SearchSharedData sets the SearchSharedData handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) SearchSharedData(handler func(err error, client *nex.Client, callID uint32, param *DataStoreSearchSharedDataParam)) {
	dataStoreSmash4Protocol.SearchSharedDataHandler = handler
}

// GetApplicationConfig sets the GetApplicationConfig handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) GetApplicationConfig(handler func(err error, client *nex.Client, callID uint32, applicationID uint32)) {
	dataStoreSmash4Protocol.GetApplicationConfigHandler = handler
}

// SearchReplay sets the SearchReplay handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) SearchReplay(handler func(err error, client *nex.Client, callID uint32, param *DataStoreSearchReplayParam)) {
	dataStoreSmash4Protocol.SearchReplayHandler = handler
}

// PostFightingPowerScore sets the PostFightingPowerScore handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) PostFightingPowerScore(handler func(err error, client *nex.Client, callID uint32, params []*DataStorePostFightingPowerScoreParam)) {
	dataStoreSmash4Protocol.PostFightingPowerScoreHandler = handler
}

// GetFightingPowerChart sets the GetFightingPowerChart handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) GetFightingPowerChart(handler func(err error, client *nex.Client, callID uint32, mode uint8)) {
	dataStoreSmash4Protocol.GetFightingPowerChartHandler = handler
}

// GetFightingPowerChartAll sets the GetFightingPowerChartAll handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) GetFightingPowerChartAll(handler func(err error, client *nex.Client, callID uint32)) {
	dataStoreSmash4Protocol.GetFightingPowerChartAllHandler = handler
}

// ReportSharedData sets the ReportSharedData handler function
func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) ReportSharedData(handler func(err error, client *nex.Client, callID uint32, dataID uint64)) {
	dataStoreSmash4Protocol.ReportSharedDataHandler = handler
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handlePostProfile(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.PostProfileHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::PostProfile not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	param, err := parametersStream.ReadStructure(NewDataStorePostProfileParam())
	if err != nil {
		go dataStoreSmash4Protocol.PostProfileHandler(err, client, callID, nil)
		return
	}

	go dataStoreSmash4Protocol.PostProfileHandler(nil, client, callID, param.(*DataStorePostProfileParam))
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleGetProfiles(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.GetProfilesHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::GetProfiles not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	pidList := parametersStream.ReadListUInt32LE()

	go dataStoreSmash4Protocol.GetProfilesHandler(nil, client, callID, pidList)
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleSendPlayReport(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.SendPlayReportHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::SendPlayReport not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	playReport := parametersStream.ReadListInt32LE()

	go dataStoreSmash4Protocol.SendPlayReportHandler(nil, client, callID, playReport)
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleGetWorldPlayReport(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.GetWorldPlayReportHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::GetWorldPlayReport not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go dataStoreSmash4Protocol.GetWorldPlayReportHandler(nil, client, callID)
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleGetReplayMeta(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.GetReplayMetaHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::GetReplayMeta not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	param, err := parametersStream.ReadStructure(NewDataStoreGetReplayMetaParam())
	if err != nil {
		go dataStoreSmash4Protocol.GetReplayMetaHandler(err, client, callID, nil)
		return
	}

	go dataStoreSmash4Protocol.GetReplayMetaHandler(nil, client, callID, param.(*DataStoreGetReplayMetaParam))
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handlePrepareGetReplay(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.PrepareGetReplayHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::PrepareGetReplay not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	param, err := parametersStream.ReadStructure(NewDataStorePrepareGetReplayParam())
	if err != nil {
		go dataStoreSmash4Protocol.PrepareGetReplayHandler(err, client, callID, nil)
		return
	}

	go dataStoreSmash4Protocol.PrepareGetReplayHandler(nil, client, callID, param.(*DataStorePrepareGetReplayParam))
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handlePreparePostReplay(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.PreparePostReplayHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::PreparePostReplay not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	param, err := parametersStream.ReadStructure(NewDataStorePreparePostReplayParam())
	if err != nil {
		go dataStoreSmash4Protocol.PreparePostReplayHandler(err, client, callID, nil)
		return
	}

	go dataStoreSmash4Protocol.PreparePostReplayHandler(nil, client, callID, param.(*DataStorePreparePostReplayParam))
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleCompletePostReplay(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.CompletePostReplayHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::CompletePostReplay not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	param, err := parametersStream.ReadStructure(NewDataStoreCompletePostReplayParam())
	if err != nil {
		go dataStoreSmash4Protocol.CompletePostReplayHandler(err, client, callID, nil)
		return
	}

	go dataStoreSmash4Protocol.CompletePostReplayHandler(nil, client, callID, param.(*DataStoreCompletePostReplayParam))
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleCheckPostReplay(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.CheckPostReplayHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::CheckPostReplay not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	param, err := parametersStream.ReadStructure(NewDataStorePreparePostReplayParam())
	if err != nil {
		go dataStoreSmash4Protocol.CheckPostReplayHandler(err, client, callID, nil)
		return
	}

	go dataStoreSmash4Protocol.CheckPostReplayHandler(nil, client, callID, param.(*DataStorePreparePostReplayParam))
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleGetNextReplay(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.GetNextReplayHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::GetNextReplay not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go dataStoreSmash4Protocol.GetNextReplayHandler(nil, client, callID)
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handlePreparePostSharedData(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.PreparePostSharedDataHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::PreparePostSharedData not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	param, err := parametersStream.ReadStructure(NewDataStorePreparePostSharedDataParam())
	if err != nil {
		go dataStoreSmash4Protocol.PreparePostSharedDataHandler(err, client, callID, nil)
		return
	}

	go dataStoreSmash4Protocol.PreparePostSharedDataHandler(nil, client, callID, param.(*DataStorePreparePostSharedDataParam))
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleCompletePostSharedData(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.CompletePostSharedDataHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::CompletePostSharedData not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	param, err := parametersStream.ReadStructure(NewDataStoreCompletePostSharedDataParam())
	if err != nil {
		go dataStoreSmash4Protocol.CompletePostSharedDataHandler(err, client, callID, nil)
		return
	}

	go dataStoreSmash4Protocol.CompletePostSharedDataHandler(nil, client, callID, param.(*DataStoreCompletePostSharedDataParam))
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleSearchSharedData(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.SearchSharedDataHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::SearchSharedData not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	param, err := parametersStream.ReadStructure(NewDataStoreSearchSharedDataParam())
	if err != nil {
		go dataStoreSmash4Protocol.SearchSharedDataHandler(err, client, callID, nil)
		return
	}

	go dataStoreSmash4Protocol.SearchSharedDataHandler(nil, client, callID, param.(*DataStoreSearchSharedDataParam))
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleGetApplicationConfig(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.GetApplicationConfigHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::GetApplicationConfig not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	applicationID := parametersStream.ReadUInt32LE()

	go dataStoreSmash4Protocol.GetApplicationConfigHandler(nil, client, callID, applicationID)
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleSearchReplay(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.SearchReplayHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::SearchReplay not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	param, err := parametersStream.ReadStructure(NewDataStoreSearchReplayParam())
	if err != nil {
		go dataStoreSmash4Protocol.SearchReplayHandler(err, client, callID, nil)
		return
	}

	go dataStoreSmash4Protocol.SearchReplayHandler(nil, client, callID, param.(*DataStoreSearchReplayParam))
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handlePostFightingPowerScore(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.PostFightingPowerScoreHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::PostFightingPowerScore not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	params, err := ReadListDataStorePostFightingPowerScoreParam(parametersStream)
	if err != nil {
		go dataStoreSmash4Protocol.PostFightingPowerScoreHandler(err, client, callID, nil)
		return
	}

	go dataStoreSmash4Protocol.PostFightingPowerScoreHandler(nil, client, callID, params)
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleGetFightingPowerChart(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.GetFightingPowerChartHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::GetFightingPowerChart not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	mode := parametersStream.ReadUInt8()

	go dataStoreSmash4Protocol.GetFightingPowerChartHandler(nil, client, callID, mode)
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleGetFightingPowerChartAll(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.GetFightingPowerChartAllHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::GetFightingPowerChartAll not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go dataStoreSmash4Protocol.GetFightingPowerChartAllHandler(nil, client, callID)
}

func (dataStoreSmash4Protocol *DataStoreSmash4Protocol) handleReportSharedData(packet nex.PacketInterface) {
	if dataStoreSmash4Protocol.ReportSharedDataHandler == nil {
		fmt.Println("[Warning] DataStoreSmash4Protocol::ReportSharedData not implemented")
		go respondNotImplemented(packet, DataStoreSmash4ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreSmash4Protocol.server)

	dataID := parametersStream.ReadUInt64LE()

	go dataStoreSmash4Protocol.ReportSharedDataHandler(nil, client, callID, dataID)
}

// NewDataStoreSmash4Protocol returns a new DataStoreSmash4Protocol
func NewDataStoreSmash4Protocol(server *nex.Server) *DataStoreSmash4Protocol {
	dataStoreSmash4Protocol := &DataStoreSmash4Protocol{server: server}
	dataStoreSmash4Protocol.DataStoreProtocol.server = server

	dataStoreSmash4Protocol.Setup()

	return dataStoreSmash4Protocol
}
