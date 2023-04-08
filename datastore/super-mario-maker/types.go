package datastore_super_mario_maker

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/datastore"
)

type DataStoreUploadCourseRecordParam struct {
	nex.Structure
	DataID uint64
	Slot   uint8
	Score  int32
}

// ExtractFromStream extracts a DataStoreUploadCourseRecordParam structure from a stream
func (dataStoreUploadCourseRecordParam *DataStoreUploadCourseRecordParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size
	dataStoreUploadCourseRecordParam.DataID = stream.ReadUInt64LE()
	dataStoreUploadCourseRecordParam.Slot = stream.ReadUInt8()
	dataStoreUploadCourseRecordParam.Score = int32(stream.ReadUInt32LE())

	return nil
}

// NewDataStoreUploadCourseRecordParam returns a new DataStoreUploadCourseRecordParam
func NewDataStoreUploadCourseRecordParam() *DataStoreUploadCourseRecordParam {
	return &DataStoreUploadCourseRecordParam{}
}

// DataStoreGetCourseRecordResult is used to send data about a courses world record
type DataStoreGetCourseRecordResult struct {
	nex.Structure
	DataID      uint64
	Slot        uint8
	FirstPID    uint32
	BestPID     uint32
	BestScore   int32
	CreatedTime *nex.DateTime
	UpdatedTime *nex.DateTime
}

// Bytes encodes the DataStoreGetCourseRecordResult and returns a byte array
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreGetCourseRecordResult.DataID)
	stream.WriteUInt8(dataStoreGetCourseRecordResult.Slot)
	stream.WriteUInt32LE(dataStoreGetCourseRecordResult.FirstPID)
	stream.WriteUInt32LE(dataStoreGetCourseRecordResult.BestPID)
	stream.WriteUInt32LE(uint32(dataStoreGetCourseRecordResult.BestScore))
	stream.WriteDateTime(dataStoreGetCourseRecordResult.CreatedTime)
	stream.WriteDateTime(dataStoreGetCourseRecordResult.UpdatedTime)

	return stream.Bytes()
}

// NewDataStoreGetCourseRecordResult returns a new DataStoreGetCourseRecordResult
func NewDataStoreGetCourseRecordResult() *DataStoreGetCourseRecordResult {
	return &DataStoreGetCourseRecordResult{}
}

// DataStoreFileServerObjectInfo is sent in the GetObjectInfos method
type DataStoreFileServerObjectInfo struct {
	nex.Structure
	DataID  uint64
	GetInfo *datastore.DataStoreReqGetInfo
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

// DataStoreAttachFileParam is sent in the PrepareAttachFile method
type DataStoreAttachFileParam struct {
	nex.Structure
	PostParam   *datastore.DataStorePreparePostParam
	ReferDataID uint64
	ContentType string
}

// ExtractFromStream extracts a DataStoreAttachFileParam structure from a stream
func (dataStoreAttachFileParam *DataStoreAttachFileParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	postParam, err := stream.ReadStructure(datastore.NewDataStorePreparePostParam())
	if err != nil {
		return err
	}

	dataStoreAttachFileParam.PostParam = postParam.(*datastore.DataStorePreparePostParam)
	dataStoreAttachFileParam.ReferDataID = stream.ReadUInt64LE()

	contentType, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreAttachFileParam.ContentType = contentType

	return nil
}

// NewDataStoreAttachFileParam returns a new DataStoreAttachFileParam
func NewDataStoreAttachFileParam() *DataStoreAttachFileParam {
	return &DataStoreAttachFileParam{}
}

// DataStoreGetCourseRecordParam is sent in the GetMetasWithCourseRecord method
type DataStoreGetCourseRecordParam struct {
	nex.Structure
	DataID uint64
	Slot   uint8
}

// ExtractFromStream extracts a DataStoreGetCourseRecordParam structure from a stream
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	dataStoreGetCourseRecordParam.DataID = stream.ReadUInt64LE()
	dataStoreGetCourseRecordParam.Slot = stream.ReadUInt8()

	return nil
}

// NewDataStoreGetCourseRecordParamreturns a new DataStoreGetCourseRecordParam
func NewDataStoreGetCourseRecordParam() *DataStoreGetCourseRecordParam {
	return &DataStoreGetCourseRecordParam{}
}

// BufferQueueParam is sent in the GetBufferQueue method
type BufferQueueParam struct {
	nex.Structure
	DataID uint64
	Slot   uint32
}

// ExtractFromStream extracts a BufferQueueParam structure from a stream
func (bufferQueueParam *BufferQueueParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	bufferQueueParam.DataID = stream.ReadUInt64LE()
	bufferQueueParam.Slot = stream.ReadUInt32LE()

	return nil
}

// NewBufferQueueParam returns a new BufferQueueParam
func NewBufferQueueParam() *BufferQueueParam {
	return &BufferQueueParam{}
}

// DataStoreRateCustomRankingParam is sent in the RateCustomRanking method
type DataStoreRateCustomRankingParam struct {
	nex.Structure
	DataID        uint64
	ApplicationId uint32
	Score         uint32
	Period        uint16
}

// ExtractFromStream extracts a DataStoreRateCustomRankingParam structure from a stream
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	dataStoreRateCustomRankingParam.DataID = stream.ReadUInt64LE()
	dataStoreRateCustomRankingParam.ApplicationId = stream.ReadUInt32LE()
	dataStoreRateCustomRankingParam.Score = stream.ReadUInt32LE()
	dataStoreRateCustomRankingParam.Period = stream.ReadUInt16LE()

	return nil
}

// NewDataStoreRateCustomRankingParam returns a new DataStoreRateCustomRankingParam
func NewDataStoreRateCustomRankingParam() *DataStoreRateCustomRankingParam {
	return &DataStoreRateCustomRankingParam{}
}

// DataStoreGetCustomRankingByDataIdParam is sent in the GetCustomRankingByDataId method
type DataStoreGetCustomRankingByDataIdParam struct {
	nex.Structure
	ApplicationId uint32
	DataIdList    []uint64
	ResultOption  uint8
}

// ExtractFromStream extracts a DataStoreGetCustomRankingByDataIdParam structure from a stream
func (dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	dataStoreGetCustomRankingByDataIdParam.ApplicationId = stream.ReadUInt32LE()
	dataStoreGetCustomRankingByDataIdParam.DataIdList = stream.ReadListUInt64LE()
	dataStoreGetCustomRankingByDataIdParam.ResultOption = stream.ReadUInt8()

	return nil
}

// NewDataStoreGetCustomRankingByDataIdParam returns a new DataStoreGetCustomRankingByDataIdParam
func NewDataStoreGetCustomRankingByDataIdParam() *DataStoreGetCustomRankingByDataIdParam {
	return &DataStoreGetCustomRankingByDataIdParam{}
}

// DataStoreCustomRankingResult is sent in the FollowingsLatestCourseSearchObject method
type DataStoreCustomRankingResult struct {
	Order    uint32
	Score    uint32
	MetaInfo *datastore.DataStoreMetaInfo

	nex.Structure
}

// ExtractFromStream extracts a DataStoreCustomRankingResult structure from a stream
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO check size

	dataStoreCustomRankingResult.Order = stream.ReadUInt32LE()
	dataStoreCustomRankingResult.Score = stream.ReadUInt32LE()

	metaInfo, err := stream.ReadStructure(datastore.NewDataStoreMetaInfo())
	if err != nil {
		return err
	}

	dataStoreCustomRankingResult.MetaInfo = metaInfo.(*datastore.DataStoreMetaInfo)

	return nil
}

// Bytes encodes the DataStoreCustomRankingResult and returns a byte array
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreCustomRankingResult.Order)
	stream.WriteUInt32LE(dataStoreCustomRankingResult.Score)
	stream.WriteStructure(dataStoreCustomRankingResult.MetaInfo)

	return stream.Bytes()
}

// NewDataStoreCustomRankingResult returns a new DataStoreCustomRankingResult
func NewDataStoreCustomRankingResult() *DataStoreCustomRankingResult {
	return &DataStoreCustomRankingResult{}
}
