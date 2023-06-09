package datastore_super_mario_maker

import (
	"fmt"

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
	var err error

	dataStoreUploadCourseRecordParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam.DataID. %s", err.Error())
	}

	dataStoreUploadCourseRecordParam.Slot, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam.Slot. %s", err.Error())
	}

	dataStoreUploadCourseRecordParam.Score, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreUploadCourseRecordParam.Score. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreUploadCourseRecordParam
func (dataStoreUploadCourseRecordParam *DataStoreUploadCourseRecordParam) Copy() nex.StructureInterface {
	copied := NewDataStoreUploadCourseRecordParam()

	copied.DataID = dataStoreUploadCourseRecordParam.DataID
	copied.Slot = dataStoreUploadCourseRecordParam.Slot
	copied.Score = dataStoreUploadCourseRecordParam.Score

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreUploadCourseRecordParam *DataStoreUploadCourseRecordParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreUploadCourseRecordParam)

	if dataStoreUploadCourseRecordParam.DataID != other.DataID {
		return false
	}

	if dataStoreUploadCourseRecordParam.Slot != other.Slot {
		return false
	}

	if dataStoreUploadCourseRecordParam.Score != other.Score {
		return false
	}

	return true
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
	stream.WriteInt32LE(dataStoreGetCourseRecordResult.BestScore)
	stream.WriteDateTime(dataStoreGetCourseRecordResult.CreatedTime)
	stream.WriteDateTime(dataStoreGetCourseRecordResult.UpdatedTime)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreGetCourseRecordResult
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) Copy() nex.StructureInterface {
	copied := NewDataStoreGetCourseRecordResult()

	copied.DataID = dataStoreGetCourseRecordResult.DataID
	copied.Slot = dataStoreGetCourseRecordResult.Slot
	copied.FirstPID = dataStoreGetCourseRecordResult.FirstPID
	copied.BestPID = dataStoreGetCourseRecordResult.BestPID
	copied.BestScore = dataStoreGetCourseRecordResult.BestScore
	copied.CreatedTime = dataStoreGetCourseRecordResult.CreatedTime.Copy()
	copied.UpdatedTime = dataStoreGetCourseRecordResult.UpdatedTime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCourseRecordResult *DataStoreGetCourseRecordResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetCourseRecordResult)

	if dataStoreGetCourseRecordResult.DataID != other.DataID {
		return false
	}

	if dataStoreGetCourseRecordResult.Slot != other.Slot {
		return false
	}

	if dataStoreGetCourseRecordResult.FirstPID != other.FirstPID {
		return false
	}

	if dataStoreGetCourseRecordResult.BestPID != other.BestPID {
		return false
	}

	if dataStoreGetCourseRecordResult.BestScore != other.BestScore {
		return false
	}

	if !dataStoreGetCourseRecordResult.CreatedTime.Equals(other.CreatedTime) {
		return false
	}

	if !dataStoreGetCourseRecordResult.UpdatedTime.Equals(other.UpdatedTime) {
		return false
	}

	return true
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

// DataStoreAttachFileParam is sent in the PrepareAttachFile method
type DataStoreAttachFileParam struct {
	nex.Structure
	PostParam   *datastore.DataStorePreparePostParam
	ReferDataID uint64
	ContentType string
}

// ExtractFromStream extracts a DataStoreAttachFileParam structure from a stream
func (dataStoreAttachFileParam *DataStoreAttachFileParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	postParam, err := stream.ReadStructure(datastore.NewDataStorePreparePostParam())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam.PostParam. %s", err.Error())
	}

	dataStoreAttachFileParam.PostParam = postParam.(*datastore.DataStorePreparePostParam)
	dataStoreAttachFileParam.ReferDataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam.ReferDataID. %s", err.Error())
	}

	dataStoreAttachFileParam.ContentType, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreAttachFileParam.ContentType. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreAttachFileParam
func (dataStoreAttachFileParam *DataStoreAttachFileParam) Copy() nex.StructureInterface {
	copied := NewDataStoreAttachFileParam()

	copied.PostParam = dataStoreAttachFileParam.PostParam.Copy().(*datastore.DataStorePreparePostParam)
	copied.ReferDataID = dataStoreAttachFileParam.ReferDataID
	copied.ContentType = dataStoreAttachFileParam.ContentType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreAttachFileParam *DataStoreAttachFileParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreAttachFileParam)

	if !dataStoreAttachFileParam.PostParam.Equals(other.PostParam) {
		return false
	}

	if dataStoreAttachFileParam.ReferDataID != other.ReferDataID {
		return false
	}

	if dataStoreAttachFileParam.ContentType != other.ContentType {
		return false
	}

	return true
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
	var err error

	dataStoreGetCourseRecordParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordParam.DataID. %s", err.Error())
	}

	dataStoreGetCourseRecordParam.Slot, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCourseRecordParam.Slot. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetCourseRecordParam
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetCourseRecordParam()

	copied.DataID = dataStoreGetCourseRecordParam.DataID
	copied.Slot = dataStoreGetCourseRecordParam.Slot

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCourseRecordParam *DataStoreGetCourseRecordParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetCourseRecordParam)

	if dataStoreGetCourseRecordParam.DataID != other.DataID {
		return false
	}

	if dataStoreGetCourseRecordParam.Slot != other.Slot {
		return false
	}

	return true
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
	var err error

	bufferQueueParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract BufferQueueParam.DataID. %s", err.Error())
	}

	bufferQueueParam.Slot, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract BufferQueueParam.Slot. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BufferQueueParam
func (bufferQueueParam *BufferQueueParam) Copy() nex.StructureInterface {
	copied := NewBufferQueueParam()

	copied.DataID = bufferQueueParam.DataID
	copied.Slot = bufferQueueParam.Slot

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (bufferQueueParam *BufferQueueParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*BufferQueueParam)

	if bufferQueueParam.DataID != other.DataID {
		return false
	}

	if bufferQueueParam.Slot != other.Slot {
		return false
	}

	return true
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
	var err error

	dataStoreRateCustomRankingParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.DataID. %s", err.Error())
	}

	dataStoreRateCustomRankingParam.ApplicationId, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.ApplicationId. %s", err.Error())
	}

	dataStoreRateCustomRankingParam.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.Score. %s", err.Error())
	}

	dataStoreRateCustomRankingParam.Period, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreRateCustomRankingParam.Period. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreRateCustomRankingParam
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) Copy() nex.StructureInterface {
	copied := NewDataStoreRateCustomRankingParam()

	copied.DataID = dataStoreRateCustomRankingParam.DataID
	copied.ApplicationId = dataStoreRateCustomRankingParam.ApplicationId
	copied.Score = dataStoreRateCustomRankingParam.Score
	copied.Period = dataStoreRateCustomRankingParam.Period

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRateCustomRankingParam *DataStoreRateCustomRankingParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRateCustomRankingParam)

	if dataStoreRateCustomRankingParam.DataID != other.DataID {
		return false
	}

	if dataStoreRateCustomRankingParam.ApplicationId != other.ApplicationId {
		return false
	}

	if dataStoreRateCustomRankingParam.Score != other.Score {
		return false
	}

	if dataStoreRateCustomRankingParam.Period != other.Period {
		return false
	}

	return true
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
	var err error

	dataStoreGetCustomRankingByDataIdParam.ApplicationId, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIdParam.ApplicationId. %s", err.Error())
	}

	dataStoreGetCustomRankingByDataIdParam.DataIdList, err = stream.ReadListUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIdParam.DataIdList. %s", err.Error())
	}

	dataStoreGetCustomRankingByDataIdParam.ResultOption, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreGetCustomRankingByDataIdParam.ResultOption. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreGetCustomRankingByDataIdParam
func (dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetCustomRankingByDataIdParam()

	copied.ApplicationId = dataStoreGetCustomRankingByDataIdParam.ApplicationId
	copied.DataIdList = make([]uint64, len(dataStoreGetCustomRankingByDataIdParam.DataIdList))

	copy(copied.DataIdList, dataStoreGetCustomRankingByDataIdParam.DataIdList)

	copied.ResultOption = dataStoreGetCustomRankingByDataIdParam.ResultOption

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetCustomRankingByDataIdParam *DataStoreGetCustomRankingByDataIdParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetCustomRankingByDataIdParam)

	if dataStoreGetCustomRankingByDataIdParam.ApplicationId != other.ApplicationId {
		return false
	}

	if len(dataStoreGetCustomRankingByDataIdParam.DataIdList) != len(other.DataIdList) {
		return false
	}

	for i := 0; i < len(dataStoreGetCustomRankingByDataIdParam.DataIdList); i++ {
		if dataStoreGetCustomRankingByDataIdParam.DataIdList[i] != other.DataIdList[i] {
			return false
		}
	}

	if dataStoreGetCustomRankingByDataIdParam.ResultOption != other.ResultOption {
		return false
	}

	return true
}

// NewDataStoreGetCustomRankingByDataIdParam returns a new DataStoreGetCustomRankingByDataIdParam
func NewDataStoreGetCustomRankingByDataIdParam() *DataStoreGetCustomRankingByDataIdParam {
	return &DataStoreGetCustomRankingByDataIdParam{}
}

// DataStoreCustomRankingResult is sent in the FollowingsLatestCourseSearchObject method
type DataStoreCustomRankingResult struct {
	nex.Structure
	Order    uint32
	Score    uint32
	MetaInfo *datastore.DataStoreMetaInfo
}

// ExtractFromStream extracts a DataStoreCustomRankingResult structure from a stream
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreCustomRankingResult.Order, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.Order. %s", err.Error())
	}

	dataStoreCustomRankingResult.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.Score. %s", err.Error())
	}

	metaInfo, err := stream.ReadStructure(datastore.NewDataStoreMetaInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCustomRankingResult.MetaInfo. %s", err.Error())
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

// Copy returns a new copied instance of DataStoreCustomRankingResult
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) Copy() nex.StructureInterface {
	copied := NewDataStoreCustomRankingResult()

	copied.Order = dataStoreCustomRankingResult.Order
	copied.Score = dataStoreCustomRankingResult.Score
	copied.MetaInfo = dataStoreCustomRankingResult.MetaInfo.Copy().(*datastore.DataStoreMetaInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCustomRankingResult *DataStoreCustomRankingResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCustomRankingResult)

	if dataStoreCustomRankingResult.Order != other.Order {
		return false
	}

	if dataStoreCustomRankingResult.Score != other.Score {
		return false
	}

	if !dataStoreCustomRankingResult.MetaInfo.Equals(other.MetaInfo) {
		return false
	}

	return true
}

// NewDataStoreCustomRankingResult returns a new DataStoreCustomRankingResult
func NewDataStoreCustomRankingResult() *DataStoreCustomRankingResult {
	return &DataStoreCustomRankingResult{}
}
