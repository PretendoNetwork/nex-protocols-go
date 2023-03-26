package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// DataStoreProtocolID is the protocol ID for the DataStore protocol
	DataStoreProtocolID = 0x73

	// DataStoreMethodPrepareGetObjectV1 is the method ID for the method PrepareGetObjectV1
	DataStoreMethodPrepareGetObjectV1 = 0x1

	// DataStoreMethodPreparePostObjectV1 is the method ID for the method PreparePostObjectV1
	DataStoreMethodPreparePostObjectV1 = 0x2

	// DataStoreMethodCompletePostObjectV1 is the method ID for the method CompletePostObjectV1
	DataStoreMethodCompletePostObjectV1 = 0x3

	// DataStoreMethodDeleteObject is the method ID for the method DeleteObject
	DataStoreMethodDeleteObject = 0x4

	// DataStoreMethodDeleteObjects is the method ID for the method DeleteObjects
	DataStoreMethodDeleteObjects = 0x5

	// DataStoreMethodChangeMetaV1 is the method ID for the method ChangeMetaV1
	DataStoreMethodChangeMetaV1 = 0x6

	// DataStoreMethodChangeMetasV1 is the method ID for the method ChangeMetasV1
	DataStoreMethodChangeMetasV1 = 0x7

	// DataStoreMethodGetMeta is the method ID for the method GetMeta
	DataStoreMethodGetMeta = 0x8

	// DataStoreMethodGetMetas is the method ID for the method GetMetas
	DataStoreMethodGetMetas = 0x9

	// DataStoreMethodPrepareUpdateObject is the method ID for the method PrepareUpdateObject
	DataStoreMethodPrepareUpdateObject = 0xA

	// DataStoreMethodCompleteUpdateObject is the method ID for the method CompleteUpdateObject
	DataStoreMethodCompleteUpdateObject = 0xB

	// DataStoreMethodSearchObject is the method ID for the method SearchObject
	DataStoreMethodSearchObject = 0xC

	// DataStoreMethodGetNotificationURL is the method ID for the method GetNotificationUrl
	DataStoreMethodGetNotificationURL = 0xD

	// DataStoreMethodGetNewArrivedNotificationsV1 is the method ID for the method GetNewArrivedNotificationsV1
	DataStoreMethodGetNewArrivedNotificationsV1 = 0xE

	// DataStoreMethodRateObject is the method ID for the method RateObject
	DataStoreMethodRateObject = 0xF

	// DataStoreMethodGetRating is the method ID for the method GetRating
	DataStoreMethodGetRating = 0x10

	// DataStoreMethodGetRatings is the method ID for the method GetRatings
	DataStoreMethodGetRatings = 0x11

	// DataStoreMethodResetRating is the method ID for the method ResetRating
	DataStoreMethodResetRating = 0x12

	// DataStoreMethodResetRatings is the method ID for the method ResetRatings
	DataStoreMethodResetRatings = 0x13

	// DataStoreMethodGetSpecificMetaV1 is the method ID for the method GetSpecificMetaV1
	DataStoreMethodGetSpecificMetaV1 = 0x14

	// DataStoreMethodPostMetaBinary is the method ID for the method PostMetaBinary
	DataStoreMethodPostMetaBinary = 0x15

	// DataStoreMethodTouchObject is the method ID for the method TouchObject
	DataStoreMethodTouchObject = 0x16

	// DataStoreMethodGetRatingWithLog is the method ID for the method GetRatingWithLog
	DataStoreMethodGetRatingWithLog = 0x17

	// DataStoreMethodPreparePostObject is the method ID for the method PreparePostObject
	DataStoreMethodPreparePostObject = 0x18

	// DataStoreMethodPrepareGetObject is the method ID for the method PrepareGetObject
	DataStoreMethodPrepareGetObject = 0x19

	// DataStoreMethodCompletePostObject is the method ID for the method CompletePostObject
	DataStoreMethodCompletePostObject = 0x1A

	// DataStoreMethodGetNewArrivedNotifications is the method ID for the method GetNewArrivedNotifications
	DataStoreMethodGetNewArrivedNotifications = 0x1B

	// DataStoreMethodGetSpecificMeta is the method ID for the method GetSpecificMeta
	DataStoreMethodGetSpecificMeta = 0x1C

	// DataStoreMethodGetPersistenceInfo is the method ID for the method GetPersistenceInfo
	DataStoreMethodGetPersistenceInfo = 0x1D

	// DataStoreMethodGetPersistenceInfos is the method ID for the method GetPersistenceInfos
	DataStoreMethodGetPersistenceInfos = 0x1E

	// DataStoreMethodPerpetuateObject is the method ID for the method PerpetuateObject
	DataStoreMethodPerpetuateObject = 0x1F

	// DataStoreMethodUnperpetuateObject is the method ID for the method UnperpetuateObject
	DataStoreMethodUnperpetuateObject = 0x20

	// DataStoreMethodPrepareGetObjectOrMetaBinary is the method ID for the method PrepareGetObjectOrMetaBinary
	DataStoreMethodPrepareGetObjectOrMetaBinary = 0x21

	// DataStoreMethodGetPasswordInfo is the method ID for the method GetPasswordInfo
	DataStoreMethodGetPasswordInfo = 0x22

	// DataStoreMethodGetPasswordInfos is the method ID for the method GetPasswordInfos
	DataStoreMethodGetPasswordInfos = 0x23

	// DataStoreMethodGetMetasMultipleParam is the method ID for the method GetMetasMultipleParam
	DataStoreMethodGetMetasMultipleParam = 0x24

	// DataStoreMethodCompletePostObjects is the method ID for the method CompletePostObjects
	DataStoreMethodCompletePostObjects = 0x25

	// DataStoreMethodChangeMeta is the method ID for the method ChangeMeta
	DataStoreMethodChangeMeta = 0x26

	// DataStoreMethodChangeMetas is the method ID for the method ChangeMetas
	DataStoreMethodChangeMetas = 0x27

	// DataStoreMethodRateObjects is the method ID for the method RateObjects
	DataStoreMethodRateObjects = 0x28

	// DataStoreMethodPostMetaBinaryWithDataID is the method ID for the method PostMetaBinaryWithDataId
	DataStoreMethodPostMetaBinaryWithDataID = 0x29

	// DataStoreMethodPostMetaBinariesWithDataID is the method ID for the method PostMetaBinariesWithDataId
	DataStoreMethodPostMetaBinariesWithDataID = 0x2A

	// DataStoreMethodRateObjectWithPosting is the method ID for the method RateObjectWithPosting
	DataStoreMethodRateObjectWithPosting = 0x2B

	// DataStoreMethodRateObjectsWithPosting is the method ID for the method RateObjectsWithPosting
	DataStoreMethodRateObjectsWithPosting = 0x2C

	// DataStoreMethodGetObjectInfos is the method ID for the method GetObjectInfos
	DataStoreMethodGetObjectInfos = 0x2D

	// DataStoreMethodSearchObjectLight is the method ID for the method SearchObjectLight
	DataStoreMethodSearchObjectLight = 0x2E
)

// DataStoreProtocol handles the DataStore nex protocol
type DataStoreProtocol struct {
	server                       *nex.Server
	GetMetaHandler               func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParam *DataStoreGetMetaParam)
	PrepareUpdateObjectHandler   func(err error, client *nex.Client, callID uint32, dataStorePrepareUpdateParam *DataStorePrepareUpdateParam)
	CompleteUpdateObjectHandler  func(err error, client *nex.Client, callID uint32, dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam)
	PostMetaBinaryHandler        func(err error, client *nex.Client, callID uint32, dataStorePreparePostParam *DataStorePreparePostParam)
	PreparePostObjectHandler     func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParam *DataStorePreparePostParam)
	PrepareGetObjectHandler      func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParam *DataStorePrepareGetParam)
	CompletePostObjectHandler    func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParam *DataStoreCompletePostParam)
	GetPersistenceInfoHandler    func(err error, client *nex.Client, callID uint32, ownerID uint32, persistenceSlotID uint16)
	GetMetasMultipleParamHandler func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParams []*DataStoreGetMetaParam)
	CompletePostObjectsHandler   func(err error, client *nex.Client, callID uint32, dataIDs []uint64)
	ChangeMetaHandler            func(err error, client *nex.Client, callID uint32, dataStoreChangeMetaParam *DataStoreChangeMetaParam)
	RateObjectsHandler           func(err error, client *nex.Client, callID uint32, targets []*DataStoreRatingTarget, params []*DataStoreRateObjectParam, transactional bool, fetchRatings bool)
}

type DataStoreNotificationV1 struct {
	nex.Structure

	NotificationID uint64
	DataID         uint32
}

// ExtractFromStream extracts a DataStoreNotificationV1 structure from a stream
func (dataStoreNotificationV1 *DataStoreNotificationV1) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreNotificationV1.NotificationID = stream.ReadUInt64LE()
	dataStoreNotificationV1.DataID = stream.ReadUInt32LE()

	return nil
}

// Bytes encodes the DataStoreNotificationV1 and returns a byte array
func (dataStoreNotificationV1 *DataStoreNotificationV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreNotificationV1.NotificationID)
	stream.WriteUInt32LE(dataStoreNotificationV1.DataID)

	return stream.Bytes()
}

// NewDataStoreNotificationV1 returns a new DataStoreNotificationV1
func NewDataStoreNotificationV1() *DataStoreNotificationV1 {
	return &DataStoreNotificationV1{}
}

type DataStoreNotification struct {
	nex.Structure

	NotificationID uint64
	DataID         uint64
}

// ExtractFromStream extracts a DataStoreNotification structure from a stream
func (dataStoreNotification *DataStoreNotification) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreNotification.NotificationID = stream.ReadUInt64LE()
	dataStoreNotification.DataID = stream.ReadUInt64LE()

	return nil
}

// Bytes encodes the DataStoreNotification and returns a byte array
func (dataStoreNotification *DataStoreNotification) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreNotification.NotificationID)
	stream.WriteUInt64LE(dataStoreNotification.DataID)

	return stream.Bytes()
}

// NewDataStoreNotification returns a new DataStoreNotification
func NewDataStoreNotification() *DataStoreNotification {
	return &DataStoreNotification{}
}

type DataStoreGetSpecificMetaParamV1 struct {
	nex.Structure

	DataIDs []uint32
}

// ExtractFromStream extracts a DataStoreGetSpecificMetaParamV1 structure from a stream
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreGetSpecificMetaParamV1.DataIDs = stream.ReadListUInt32LE()

	return nil
}

// Bytes encodes the DataStoreGetSpecificMetaParamV1 and returns a byte array
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListUInt32LE(dataStoreGetSpecificMetaParamV1.DataIDs)

	return stream.Bytes()
}

// NewDataStoreGetSpecificMetaParamV1 returns a new DataStoreGetSpecificMetaParamV1
func NewDataStoreGetSpecificMetaParamV1() *DataStoreGetSpecificMetaParamV1 {
	return &DataStoreGetSpecificMetaParamV1{}
}

type DataStoreGetSpecificMetaParam struct {
	nex.Structure

	DataIDs []uint64
}

// ExtractFromStream extracts a DataStoreGetSpecificMetaParam structure from a stream
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreGetSpecificMetaParam.DataIDs = stream.ReadListUInt64LE()

	return nil
}

// Bytes encodes the DataStoreGetSpecificMetaParam and returns a byte array
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListUInt64LE(dataStoreGetSpecificMetaParam.DataIDs)

	return stream.Bytes()
}

// NewDataStoreGetSpecificMetaParam returns a new DataStoreGetSpecificMetaParam
func NewDataStoreGetSpecificMetaParam() *DataStoreGetSpecificMetaParam {
	return &DataStoreGetSpecificMetaParam{}
}

type DataStoreSpecificMetaInfoV1 struct {
	nex.Structure

	DataID   uint32
	OwnerID  uint32
	Size     uint32
	DataType uint16
	Version  uint16
}

// ExtractFromStream extracts a DataStoreSpecificMetaInfoV1 structure from a stream
func (dataStoreSpecificMetaInfoV1 *DataStoreSpecificMetaInfoV1) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreSpecificMetaInfoV1.DataID = stream.ReadUInt32LE()
	dataStoreSpecificMetaInfoV1.OwnerID = stream.ReadUInt32LE()
	dataStoreSpecificMetaInfoV1.Size = stream.ReadUInt32LE()
	dataStoreSpecificMetaInfoV1.DataType = stream.ReadUInt16LE()
	dataStoreSpecificMetaInfoV1.Version = stream.ReadUInt16LE()

	return nil
}

// Bytes encodes the DataStoreSpecificMetaInfoV1 and returns a byte array
func (dataStoreSpecificMetaInfoV1 *DataStoreSpecificMetaInfoV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreSpecificMetaInfoV1.DataID)
	stream.WriteUInt32LE(dataStoreSpecificMetaInfoV1.OwnerID)
	stream.WriteUInt32LE(dataStoreSpecificMetaInfoV1.Size)
	stream.WriteUInt16LE(dataStoreSpecificMetaInfoV1.DataType)
	stream.WriteUInt16LE(dataStoreSpecificMetaInfoV1.Version)

	return stream.Bytes()
}

// NewDataStoreSpecificMetaInfoV1 returns a new DataStoreSpecificMetaInfoV1
func NewDataStoreSpecificMetaInfoV1() *DataStoreSpecificMetaInfoV1 {
	return &DataStoreSpecificMetaInfoV1{}
}

type DataStoreSpecificMetaInfo struct {
	nex.Structure

	DataID   uint64
	OwnerID  uint32
	Size     uint32
	DataType uint16
	Version  uint32
}

// ExtractFromStream extracts a DataStoreSpecificMetaInfo structure from a stream
func (dataStoreSpecificMetaInfo *DataStoreSpecificMetaInfo) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreSpecificMetaInfo.DataID = stream.ReadUInt64LE()
	dataStoreSpecificMetaInfo.OwnerID = stream.ReadUInt32LE()
	dataStoreSpecificMetaInfo.Size = stream.ReadUInt32LE()
	dataStoreSpecificMetaInfo.DataType = stream.ReadUInt16LE()
	dataStoreSpecificMetaInfo.Version = stream.ReadUInt32LE()

	return nil
}

// Bytes encodes the DataStoreSpecificMetaInfo and returns a byte array
func (dataStoreSpecificMetaInfo *DataStoreSpecificMetaInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreSpecificMetaInfo.DataID)
	stream.WriteUInt32LE(dataStoreSpecificMetaInfo.OwnerID)
	stream.WriteUInt32LE(dataStoreSpecificMetaInfo.Size)
	stream.WriteUInt16LE(dataStoreSpecificMetaInfo.DataType)
	stream.WriteUInt32LE(dataStoreSpecificMetaInfo.Version)

	return stream.Bytes()
}

// NewDataStoreSpecificMetaInfo returns a new DataStoreSpecificMetaInfo
func NewDataStoreSpecificMetaInfo() *DataStoreSpecificMetaInfo {
	return &DataStoreSpecificMetaInfo{}
}

type DataStoreTouchObjectParam struct {
	nex.Structure

	DataID         uint64
	LockID         uint32
	AccessPassword uint64
}

// ExtractFromStream extracts a DataStoreTouchObjectParam structure from a stream
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreTouchObjectParam.DataID = stream.ReadUInt64LE()
	dataStoreTouchObjectParam.LockID = stream.ReadUInt32LE()
	dataStoreTouchObjectParam.AccessPassword = stream.ReadUInt64LE()

	return nil
}

// Bytes encodes the DataStoreTouchObjectParam and returns a byte array
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreTouchObjectParam.DataID)
	stream.WriteUInt32LE(dataStoreTouchObjectParam.LockID)
	stream.WriteUInt64LE(dataStoreTouchObjectParam.AccessPassword)

	return stream.Bytes()
}

// NewDataStoreTouchObjectParam returns a new DataStoreTouchObjectParam
func NewDataStoreTouchObjectParam() *DataStoreTouchObjectParam {
	return &DataStoreTouchObjectParam{}
}

type DataStoreRatingLog struct {
	nex.Structure

	IsRated            bool
	Pid                uint32
	RatingValue        int32
	LockExpirationTime *nex.DateTime
}

// ExtractFromStream extracts a DataStoreRatingLog structure from a stream
func (dataStoreRatingLog *DataStoreRatingLog) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreRatingLog.IsRated = stream.ReadUInt8() == 1
	dataStoreRatingLog.Pid = stream.ReadUInt32LE()
	dataStoreRatingLog.RatingValue = int32(stream.ReadUInt32LE())
	dataStoreRatingLog.LockExpirationTime = nex.NewDateTime(stream.ReadUInt64LE())

	return nil
}

// Bytes encodes the DataStoreRatingLog and returns a byte array
func (dataStoreRatingLog *DataStoreRatingLog) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteBool(dataStoreRatingLog.IsRated)
	stream.WriteUInt32LE(dataStoreRatingLog.Pid)
	stream.WriteInt32LE(dataStoreRatingLog.RatingValue)
	stream.WriteUInt64LE(dataStoreRatingLog.LockExpirationTime.Value())

	return stream.Bytes()
}

// NewDataStoreRatingLog returns a new DataStoreRatingLog
func NewDataStoreRatingLog() *DataStoreRatingLog {
	return &DataStoreRatingLog{}
}

type DataStorePersistenceInfo struct {
	nex.Structure

	OwnerID           uint32
	PersistenceSlotID uint16
	DataID            uint64
}

// ExtractFromStream extracts a DataStorePersistenceInfo structure from a stream
func (dataStorePersistenceInfo *DataStorePersistenceInfo) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePersistenceInfo.OwnerID = stream.ReadUInt32LE()
	dataStorePersistenceInfo.PersistenceSlotID = stream.ReadUInt16LE()
	dataStorePersistenceInfo.DataID = stream.ReadUInt64LE()

	return nil
}

// Bytes encodes the DataStorePersistenceInfo and returns a byte array
func (dataStorePersistenceInfo *DataStorePersistenceInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStorePersistenceInfo.OwnerID)
	stream.WriteUInt16LE(dataStorePersistenceInfo.PersistenceSlotID)
	stream.WriteUInt64LE(dataStorePersistenceInfo.DataID)

	return stream.Bytes()
}

// NewDataStorePersistenceInfo returns a new DataStorePersistenceInfo
func NewDataStorePersistenceInfo() *DataStorePersistenceInfo {
	return &DataStorePersistenceInfo{}
}

type DataStorePasswordInfo struct {
	nex.Structure

	DataID         uint64
	AccessPassword uint64
	UpdatePassword uint64
}

// ExtractFromStream extracts a DataStorePasswordInfo structure from a stream
func (dataStorePasswordInfo *DataStorePasswordInfo) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePasswordInfo.DataID = stream.ReadUInt64LE()
	dataStorePasswordInfo.AccessPassword = stream.ReadUInt64LE()
	dataStorePasswordInfo.UpdatePassword = stream.ReadUInt64LE()

	return nil
}

// Bytes encodes the DataStorePasswordInfo and returns a byte array
func (dataStorePasswordInfo *DataStorePasswordInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStorePasswordInfo.DataID)
	stream.WriteUInt64LE(dataStorePasswordInfo.AccessPassword)
	stream.WriteUInt64LE(dataStorePasswordInfo.UpdatePassword)

	return stream.Bytes()
}

// NewDataStorePasswordInfo returns a new DataStorePasswordInfo
func NewDataStorePasswordInfo() *DataStorePasswordInfo {
	return &DataStorePasswordInfo{}
}

type DataStoreGetNewArrivedNotificationsParam struct {
	nex.Structure

	LastNotificationID uint64
	Limit              uint16
}

// ExtractFromStream extracts a DataStoreGetNewArrivedNotificationsParam structure from a stream
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreGetNewArrivedNotificationsParam.LastNotificationID = stream.ReadUInt64LE()
	dataStoreGetNewArrivedNotificationsParam.Limit = stream.ReadUInt16LE()

	return nil
}

// Bytes encodes the DataStoreGetNewArrivedNotificationsParam and returns a byte array
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreGetNewArrivedNotificationsParam.LastNotificationID)
	stream.WriteUInt16LE(dataStoreGetNewArrivedNotificationsParam.Limit)

	return stream.Bytes()
}

// NewDataStoreGetNewArrivedNotificationsParam returns a new DataStoreGetNewArrivedNotificationsParam
func NewDataStoreGetNewArrivedNotificationsParam() *DataStoreGetNewArrivedNotificationsParam {
	return &DataStoreGetNewArrivedNotificationsParam{}
}

type DataStoreReqGetNotificationUrlInfo struct {
	nex.Structure

	Url        string
	Key        string
	Query      string
	RootCaCert []byte
}

// ExtractFromStream extracts a DataStoreReqGetNotificationUrlInfo structure from a stream
func (dataStoreReqGetNotificationUrlInfo *DataStoreReqGetNotificationUrlInfo) ExtractFromStream(stream *nex.StreamIn) error {

	url, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreReqGetNotificationUrlInfo.Url = url

	key, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreReqGetNotificationUrlInfo.Key = key

	query, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreReqGetNotificationUrlInfo.Query = query

	rootCaCert, err := stream.ReadBuffer()
	if err != nil {
		return err
	}

	dataStoreReqGetNotificationUrlInfo.RootCaCert = rootCaCert

	return nil
}

// Bytes encodes the DataStoreReqGetNotificationUrlInfo and returns a byte array
func (dataStoreReqGetNotificationUrlInfo *DataStoreReqGetNotificationUrlInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreReqGetNotificationUrlInfo.Url)
	stream.WriteString(dataStoreReqGetNotificationUrlInfo.Key)
	stream.WriteString(dataStoreReqGetNotificationUrlInfo.Query)
	stream.WriteBuffer(dataStoreReqGetNotificationUrlInfo.RootCaCert)

	return stream.Bytes()
}

// NewDataStoreReqGetNotificationUrlInfo returns a new DataStoreReqGetNotificationUrlInfo
func NewDataStoreReqGetNotificationUrlInfo() *DataStoreReqGetNotificationUrlInfo {
	return &DataStoreReqGetNotificationUrlInfo{}
}

type DataStoreGetNotificationUrlParam struct {
	nex.Structure

	PreviousUrl string
}

// ExtractFromStream extracts a DataStoreGetNotificationUrlParam structure from a stream
func (dataStoreGetNotificationUrlParam *DataStoreGetNotificationUrlParam) ExtractFromStream(stream *nex.StreamIn) error {

	previousUrl, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreGetNotificationUrlParam.PreviousUrl = previousUrl

	return nil
}

// Bytes encodes the DataStoreGetNotificationUrlParam and returns a byte array
func (dataStoreGetNotificationUrlParam *DataStoreGetNotificationUrlParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreGetNotificationUrlParam.PreviousUrl)

	return stream.Bytes()
}

// NewDataStoreGetNotificationUrlParam returns a new DataStoreGetNotificationUrlParam
func NewDataStoreGetNotificationUrlParam() *DataStoreGetNotificationUrlParam {
	return &DataStoreGetNotificationUrlParam{}
}

type DataStoreSearchResult struct {
	nex.Structure

	TotalCount     uint32
	Result         []*DataStoreMetaInfo
	TotalCountType uint8
}

// ExtractFromStream extracts a DataStoreSearchResult structure from a stream
func (dataStoreSearchResult *DataStoreSearchResult) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreSearchResult.TotalCount = stream.ReadUInt32LE()

	result, err := ReadListDataStoreMetaInfo(stream)
	if err != nil {
		return err
	}

	dataStoreSearchResult.Result = result
	dataStoreSearchResult.TotalCountType = stream.ReadUInt8()

	return nil
}

// Bytes encodes the DataStoreSearchResult and returns a byte array
func (dataStoreSearchResult *DataStoreSearchResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreSearchResult.TotalCount)
	stream.WriteListStructure(dataStoreSearchResult.Result)
	stream.WriteUInt8(dataStoreSearchResult.TotalCountType)

	return stream.Bytes()
}

// NewDataStoreSearchResult returns a new DataStoreSearchResult
func NewDataStoreSearchResult() *DataStoreSearchResult {
	return &DataStoreSearchResult{}
}

type DataStoreCompleteUpdateParam struct {
	nex.Structure

	DataID    uint64
	Version   uint32
	IsSuccess bool
}

// ExtractFromStream extracts a DataStoreCompleteUpdateParam structure from a stream
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreCompleteUpdateParam.DataID = stream.ReadUInt64LE()
	dataStoreCompleteUpdateParam.Version = stream.ReadUInt32LE()
	dataStoreCompleteUpdateParam.IsSuccess = stream.ReadUInt8() == 1

	return nil
}

// Bytes encodes the DataStoreCompleteUpdateParam and returns a byte array
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreCompleteUpdateParam.DataID)
	stream.WriteUInt32LE(dataStoreCompleteUpdateParam.Version)
	stream.WriteBool(dataStoreCompleteUpdateParam.IsSuccess)

	return stream.Bytes()
}

// NewDataStoreCompleteUpdateParam returns a new DataStoreCompleteUpdateParam
func NewDataStoreCompleteUpdateParam() *DataStoreCompleteUpdateParam {
	return &DataStoreCompleteUpdateParam{}
}

type DataStoreReqUpdateInfo struct {
	nex.Structure

	Version        uint32
	Url            string
	RequestHeaders []*DataStoreKeyValue
	FormFields     []*DataStoreKeyValue
	RootCaCert     []byte
}

// ExtractFromStream extracts a DataStoreReqUpdateInfo structure from a stream
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreReqUpdateInfo.Version = stream.ReadUInt32LE()

	url, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreReqUpdateInfo.Url = url

	requestHeaders, err := ReadListDataStoreKeyValue(stream)
	if err != nil {
		return err
	}

	dataStoreReqUpdateInfo.RequestHeaders = requestHeaders

	formFields, err := ReadListDataStoreKeyValue(stream)
	if err != nil {
		return err
	}

	dataStoreReqUpdateInfo.FormFields = formFields

	rootCaCert, err := stream.ReadBuffer()
	if err != nil {
		return err
	}

	dataStoreReqUpdateInfo.RootCaCert = rootCaCert

	return nil
}

// Bytes encodes the DataStoreReqUpdateInfo and returns a byte array
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreReqUpdateInfo.Version)
	stream.WriteString(dataStoreReqUpdateInfo.Url)
	stream.WriteListStructure(dataStoreReqUpdateInfo.RequestHeaders)
	stream.WriteListStructure(dataStoreReqUpdateInfo.FormFields)
	stream.WriteBuffer(dataStoreReqUpdateInfo.RootCaCert)

	return stream.Bytes()
}

// NewDataStoreReqUpdateInfo returns a new DataStoreReqUpdateInfo
func NewDataStoreReqUpdateInfo() *DataStoreReqUpdateInfo {
	return &DataStoreReqUpdateInfo{}
}

type DataStorePrepareUpdateParam struct {
	nex.Structure

	DataID         uint64
	Size           uint32
	UpdatePassword uint64
	ExtraData      []string
}

// ExtractFromStream extracts a DataStorePrepareUpdateParam structure from a stream
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePrepareUpdateParam.DataID = stream.ReadUInt64LE()
	dataStorePrepareUpdateParam.Size = stream.ReadUInt32LE()
	dataStorePrepareUpdateParam.UpdatePassword = stream.ReadUInt64LE()
	dataStorePrepareUpdateParam.ExtraData = stream.ReadListString()

	return nil
}

// Bytes encodes the DataStorePrepareUpdateParam and returns a byte array
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStorePrepareUpdateParam.DataID)
	stream.WriteUInt32LE(dataStorePrepareUpdateParam.Size)
	stream.WriteUInt64LE(dataStorePrepareUpdateParam.UpdatePassword)
	stream.WriteListString(dataStorePrepareUpdateParam.ExtraData)

	return stream.Bytes()
}

// NewDataStorePrepareUpdateParam returns a new DataStorePrepareUpdateParam
func NewDataStorePrepareUpdateParam() *DataStorePrepareUpdateParam {
	return &DataStorePrepareUpdateParam{}
}

type DataStoreChangeMetaParamV1 struct {
	nex.Structure

	DataID         uint64
	ModifiesFlag   uint32
	Name           string
	Permission     *DataStorePermission
	DelPermission  *DataStorePermission
	Period         uint16
	MetaBinary     []byte
	Tags           []string
	UpdatePassword uint64
}

// ExtractFromStream extracts a DataStoreChangeMetaParamV1 structure from a stream
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreChangeMetaParamV1.DataID = stream.ReadUInt64LE()
	dataStoreChangeMetaParamV1.ModifiesFlag = stream.ReadUInt32LE()

	name, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreChangeMetaParamV1.Name = name

	permission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return err
	}

	dataStoreChangeMetaParamV1.Permission = permission.(*DataStorePermission)

	delPermission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return err
	}

	dataStoreChangeMetaParamV1.DelPermission = delPermission.(*DataStorePermission)
	dataStoreChangeMetaParamV1.Period = stream.ReadUInt16LE()

	metaBinary, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	dataStoreChangeMetaParamV1.MetaBinary = metaBinary
	dataStoreChangeMetaParamV1.Tags = stream.ReadListString()
	dataStoreChangeMetaParamV1.UpdatePassword = stream.ReadUInt64LE()

	return nil
}

// Bytes encodes the DataStoreChangeMetaParamV1 and returns a byte array
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreChangeMetaParamV1.DataID)
	stream.WriteUInt32LE(dataStoreChangeMetaParamV1.ModifiesFlag)
	stream.WriteString(dataStoreChangeMetaParamV1.Name)
	stream.WriteStructure(dataStoreChangeMetaParamV1.Permission)
	stream.WriteStructure(dataStoreChangeMetaParamV1.DelPermission)
	stream.WriteUInt16LE(dataStoreChangeMetaParamV1.Period)
	stream.WriteQBuffer(dataStoreChangeMetaParamV1.MetaBinary)
	stream.WriteListString(dataStoreChangeMetaParamV1.Tags)
	stream.WriteUInt64LE(dataStoreChangeMetaParamV1.UpdatePassword)

	return stream.Bytes()
}

// NewDataStoreChangeMetaParamV1 returns a new DataStoreChangeMetaParamV1
func NewDataStoreChangeMetaParamV1() *DataStoreChangeMetaParamV1 {
	return &DataStoreChangeMetaParamV1{}
}

type DataStoreDeleteParam struct {
	nex.Structure

	DataID         uint64
	UpdatePassword uint64
}

// ExtractFromStream extracts a DataStoreDeleteParam structure from a stream
func (dataStoreDeleteParam *DataStoreDeleteParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreDeleteParam.DataID = stream.ReadUInt64LE()
	dataStoreDeleteParam.UpdatePassword = stream.ReadUInt64LE()

	return nil
}

// Bytes encodes the DataStoreDeleteParam and returns a byte array
func (dataStoreDeleteParam *DataStoreDeleteParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreDeleteParam.DataID)
	stream.WriteUInt64LE(dataStoreDeleteParam.UpdatePassword)

	return stream.Bytes()
}

// NewDataStoreDeleteParam returns a new DataStoreDeleteParam
func NewDataStoreDeleteParam() *DataStoreDeleteParam {
	return &DataStoreDeleteParam{}
}

type DataStoreCompletePostParamV1 struct {
	nex.Structure

	DataID    uint32
	IsSuccess bool
}

// ExtractFromStream extracts a DataStoreCompletePostParamV1 structure from a stream
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreCompletePostParamV1.DataID = stream.ReadUInt32LE()
	dataStoreCompletePostParamV1.IsSuccess = stream.ReadUInt8() == 1

	return nil
}

// Bytes encodes the DataStoreCompletePostParamV1 and returns a byte array
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreCompletePostParamV1.DataID)
	stream.WriteBool(dataStoreCompletePostParamV1.IsSuccess)

	return stream.Bytes()
}

// NewDataStoreCompletePostParamV1 returns a new DataStoreCompletePostParamV1
func NewDataStoreCompletePostParamV1() *DataStoreCompletePostParamV1 {
	return &DataStoreCompletePostParamV1{}
}

type DataStoreReqPostInfoV1 struct {
	nex.Structure

	DataID         uint32
	Url            string
	RequestHeaders []*DataStoreKeyValue
	FormFields     []*DataStoreKeyValue
	RootCaCert     []byte
}

// ExtractFromStream extracts a DataStoreReqPostInfoV1 structure from a stream
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreReqPostInfoV1.DataID = stream.ReadUInt32LE()

	url, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreReqPostInfoV1.Url = url

	requestHeaders, err := ReadListDataStoreKeyValue(stream)
	if err != nil {
		return err
	}

	dataStoreReqPostInfoV1.RequestHeaders = requestHeaders

	formFields, err := ReadListDataStoreKeyValue(stream)
	if err != nil {
		return err
	}

	dataStoreReqPostInfoV1.FormFields = formFields

	rootCaCert, err := stream.ReadBuffer()
	if err != nil {
		return err
	}

	dataStoreReqPostInfoV1.RootCaCert = rootCaCert

	return nil
}

// Bytes encodes the DataStoreReqPostInfoV1 and returns a byte array
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStoreReqPostInfoV1.DataID)
	stream.WriteString(dataStoreReqPostInfoV1.Url)
	stream.WriteListStructure(dataStoreReqPostInfoV1.RequestHeaders)
	stream.WriteListStructure(dataStoreReqPostInfoV1.FormFields)
	stream.WriteBuffer(dataStoreReqPostInfoV1.RootCaCert)

	return stream.Bytes()
}

// NewDataStoreReqPostInfoV1 returns a new DataStoreReqPostInfoV1
func NewDataStoreReqPostInfoV1() *DataStoreReqPostInfoV1 {
	return &DataStoreReqPostInfoV1{}
}

type DataStorePreparePostParamV1 struct {
	nex.Structure

	Size             uint32
	Name             string
	DataType         uint16
	MetaBinary       []byte
	Permission       *DataStorePermission
	DelPermission    *DataStorePermission
	Flag             uint32
	Period           uint16
	ReferDataID      uint32
	Tags             []string
	RatingInitParams []*DataStoreRatingInitParamWithSlot
}

// ExtractFromStream extracts a DataStorePreparePostParamV1 structure from a stream
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePreparePostParamV1.Size = stream.ReadUInt32LE()

	name, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStorePreparePostParamV1.Name = name
	dataStorePreparePostParamV1.DataType = stream.ReadUInt16LE()

	metaBinary, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	dataStorePreparePostParamV1.MetaBinary = metaBinary

	permission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return err
	}

	dataStorePreparePostParamV1.Permission = permission.(*DataStorePermission)

	delPermission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return err
	}

	dataStorePreparePostParamV1.DelPermission = delPermission.(*DataStorePermission)
	dataStorePreparePostParamV1.Flag = stream.ReadUInt32LE()
	dataStorePreparePostParamV1.Period = stream.ReadUInt16LE()
	dataStorePreparePostParamV1.ReferDataID = stream.ReadUInt32LE()
	dataStorePreparePostParamV1.Tags = stream.ReadListString()

	ratingInitParams, err := ReadListDataStoreRatingInitParamWithSlot(stream)
	if err != nil {
		return err
	}

	dataStorePreparePostParamV1.RatingInitParams = ratingInitParams

	return nil
}

// Bytes encodes the DataStorePreparePostParamV1 and returns a byte array
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStorePreparePostParamV1.Size)
	stream.WriteString(dataStorePreparePostParamV1.Name)
	stream.WriteUInt16LE(dataStorePreparePostParamV1.DataType)
	stream.WriteQBuffer(dataStorePreparePostParamV1.MetaBinary)
	stream.WriteStructure(dataStorePreparePostParamV1.Permission)
	stream.WriteStructure(dataStorePreparePostParamV1.DelPermission)
	stream.WriteUInt32LE(dataStorePreparePostParamV1.Flag)
	stream.WriteUInt16LE(dataStorePreparePostParamV1.Period)
	stream.WriteUInt32LE(dataStorePreparePostParamV1.ReferDataID)
	stream.WriteListString(dataStorePreparePostParamV1.Tags)
	stream.WriteListStructure(dataStorePreparePostParamV1.RatingInitParams)

	return stream.Bytes()
}

// NewDataStorePreparePostParamV1 returns a new DataStorePreparePostParamV1
func NewDataStorePreparePostParamV1() *DataStorePreparePostParamV1 {
	return &DataStorePreparePostParamV1{}
}

type DataStoreReqGetInfoV1 struct {
	nex.Structure

	Url            string
	RequestHeaders []*DataStoreKeyValue
	Size           uint32
	RootCaCert     []byte
}

// ExtractFromStream extracts a DataStoreReqGetInfoV1 structure from a stream
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) ExtractFromStream(stream *nex.StreamIn) error {

	url, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreReqGetInfoV1.Url = url

	requestHeaders, err := ReadListDataStoreKeyValue(stream)
	if err != nil {
		return err
	}

	dataStoreReqGetInfoV1.RequestHeaders = requestHeaders
	dataStoreReqGetInfoV1.Size = stream.ReadUInt32LE()

	rootCaCert, err := stream.ReadBuffer()
	if err != nil {
		return err
	}

	dataStoreReqGetInfoV1.RootCaCert = rootCaCert

	return nil
}

// Bytes encodes the DataStoreReqGetInfoV1 and returns a byte array
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreReqGetInfoV1.Url)
	stream.WriteListStructure(dataStoreReqGetInfoV1.RequestHeaders)
	stream.WriteUInt32LE(dataStoreReqGetInfoV1.Size)
	stream.WriteBuffer(dataStoreReqGetInfoV1.RootCaCert)

	return stream.Bytes()
}

// NewDataStoreReqGetInfoV1 returns a new DataStoreReqGetInfoV1
func NewDataStoreReqGetInfoV1() *DataStoreReqGetInfoV1 {
	return &DataStoreReqGetInfoV1{}
}

type DataStorePrepareGetParamV1 struct {
	nex.Structure

	DataID uint32
	LockID uint32
}

// ExtractFromStream extracts a DataStorePrepareGetParamV1 structure from a stream
func (dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePrepareGetParamV1.DataID = stream.ReadUInt32LE()
	dataStorePrepareGetParamV1.LockID = stream.ReadUInt32LE()

	return nil
}

// Bytes encodes the DataStorePrepareGetParamV1 and returns a byte array
func (dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(dataStorePrepareGetParamV1.DataID)
	stream.WriteUInt32LE(dataStorePrepareGetParamV1.LockID)

	return stream.Bytes()
}

// NewDataStorePrepareGetParamV1 returns a new DataStorePrepareGetParamV1
func NewDataStorePrepareGetParamV1() *DataStorePrepareGetParamV1 {
	return &DataStorePrepareGetParamV1{}
}

// DataStoreRateObjectParam is sent in the RateObjects method
type DataStoreRateObjectParam struct {
	nex.Structure
	RatingValue    int32
	AccessPassword uint64
}

// ExtractFromStream extracts a DataStoreRateObjectParam structure from a stream
func (dataStoreRateObjectParam *DataStoreRateObjectParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreRateObjectParam.RatingValue = int32(stream.ReadUInt32LE())
	dataStoreRateObjectParam.AccessPassword = stream.ReadUInt64LE()

	return nil
}

// NewDataStoreRateObjectParam returns a new DataStoreRateObjectParam
func NewDataStoreRateObjectParam() *DataStoreRateObjectParam {
	return &DataStoreRateObjectParam{}
}

// DataStoreRatingTarget is sent in the RateObjects method
type DataStoreRatingTarget struct {
	nex.Structure
	DataID uint64
	Slot   uint8
}

// ExtractFromStream extracts a DataStoreRatingTarget structure from a stream
func (dataStoreRatingTarget *DataStoreRatingTarget) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreRatingTarget.DataID = stream.ReadUInt64LE()
	dataStoreRatingTarget.Slot = stream.ReadUInt8()

	return nil
}

// NewDataStoreRatingTarget returns a new DataStoreRatingTarget
func NewDataStoreRatingTarget() *DataStoreRatingTarget {
	return &DataStoreRatingTarget{}
}

// DataStoreCompletePostParam is sent in the CompletePostObject method
type DataStoreCompletePostParam struct {
	nex.Structure
	DataID    uint64
	IsSuccess bool
}

// ExtractFromStream extracts a DataStoreCompletePostParam structure from a stream
func (dataStoreCompletePostParam *DataStoreCompletePostParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreCompletePostParam.DataID = stream.ReadUInt64LE()
	dataStoreCompletePostParam.IsSuccess = (stream.ReadUInt8() == 1)

	return nil
}

// NewDataStoreCompletePostParam returns a new DataStoreCompletePostParam
func NewDataStoreCompletePostParam() *DataStoreCompletePostParam {
	return &DataStoreCompletePostParam{}
}

// DataStoreReqPostInfo is sent in the PreparePostObject method
type DataStoreReqPostInfo struct {
	nex.Structure
	DataID         uint64
	URL            string
	RequestHeaders []*DataStoreKeyValue
	FormFields     []*DataStoreKeyValue
	RootCACert     []byte
}

// Bytes encodes the DataStoreReqPostInfo and returns a byte array
func (dataStoreReqPostInfo *DataStoreReqPostInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreReqPostInfo.DataID)
	stream.WriteString(dataStoreReqPostInfo.URL)
	stream.WriteListStructure(dataStoreReqPostInfo.RequestHeaders)
	stream.WriteListStructure(dataStoreReqPostInfo.FormFields)
	stream.WriteBuffer(dataStoreReqPostInfo.RootCACert)

	return stream.Bytes()
}

// NewDataStoreReqPostInfo returns a new DataStoreReqPostInfo
func NewDataStoreReqPostInfo() *DataStoreReqPostInfo {
	return &DataStoreReqPostInfo{}
}

// DataStorePersistenceInitParam is sent in the PreparePostObject method
type DataStorePersistenceInitParam struct {
	nex.Structure
	PersistenceSlotId uint16
	DeleteLastObject  bool
}

// ExtractFromStream extracts a DataStorePersistenceInitParam structure from a stream
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePersistenceInitParam.PersistenceSlotId = stream.ReadUInt16LE()
	dataStorePersistenceInitParam.DeleteLastObject = (stream.ReadUInt8() == 1)

	return nil
}

// NewDataStorePersistenceInitParam returns a new DataStorePersistenceInitParam
func NewDataStorePersistenceInitParam() *DataStorePersistenceInitParam {
	return &DataStorePersistenceInitParam{}
}

// DataStoreRatingInitParam is sent in the PreparePostObject method
type DataStoreRatingInitParam struct {
	nex.Structure
	Flag           uint8
	InternalFlag   uint8
	LockType       uint8
	initialValue   int64
	rangeMin       int32
	rangeMax       int32
	periodHour     int8
	periodDuration int16
}

// ExtractFromStream extracts a DataStoreRatingInitParam structure from a stream
func (dataStoreRatingInitParam *DataStoreRatingInitParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreRatingInitParam.Flag = stream.ReadUInt8()
	dataStoreRatingInitParam.InternalFlag = stream.ReadUInt8()
	dataStoreRatingInitParam.LockType = stream.ReadUInt8()
	dataStoreRatingInitParam.initialValue = int64(stream.ReadUInt64LE())
	dataStoreRatingInitParam.rangeMin = int32(stream.ReadUInt32LE())
	dataStoreRatingInitParam.rangeMax = int32(stream.ReadUInt32LE())
	dataStoreRatingInitParam.periodHour = int8(stream.ReadUInt8())
	dataStoreRatingInitParam.periodDuration = int16(stream.ReadUInt16LE())

	return nil
}

// NewDataStoreRatingInitParam returns a new DataStoreRatingInitParam
func NewDataStoreRatingInitParam() *DataStoreRatingInitParam {
	return &DataStoreRatingInitParam{}
}

// DataStoreRatingInitParamWithSlot is sent in the PreparePostObject method
type DataStoreRatingInitParamWithSlot struct {
	nex.Structure
	Slot  int8
	Param *DataStoreRatingInitParam
}

// ExtractFromStream extracts a DataStoreRatingInitParamWithSlot structure from a stream
func (dataStoreRatingInitParamWithSlot *DataStoreRatingInitParamWithSlot) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreRatingInitParamWithSlot.Slot = int8(stream.ReadUInt8())

	param, err := stream.ReadStructure(NewDataStoreRatingInitParam())
	if err != nil {
		return err
	}

	dataStoreRatingInitParamWithSlot.Param = param.(*DataStoreRatingInitParam)

	return nil
}

// NewDataStoreRatingInitParamWithSlot returns a new DataStoreRatingInitParamWithSlot
func NewDataStoreRatingInitParamWithSlot() *DataStoreRatingInitParamWithSlot {
	return &DataStoreRatingInitParamWithSlot{}
}

// DataStoreSearchParam is sent in the PreparePostObject method
type DataStorePreparePostParam struct {
	nex.Structure
	Size                 uint32
	Name                 string
	DataType             uint16
	MetaBinary           []byte
	Permission           *DataStorePermission
	DelPermission        *DataStorePermission
	Flag                 uint32
	Period               uint16
	ReferDataId          uint32
	Tags                 []string
	RatingInitParams     []*DataStoreRatingInitParamWithSlot
	PersistenceInitParam *DataStorePersistenceInitParam
	ExtraData            []string
}

// ExtractFromStream extracts a DataStorePreparePostParam structure from a stream
func (dataStorePreparePostParam *DataStorePreparePostParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePreparePostParam.Size = stream.ReadUInt32LE()

	name, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStorePreparePostParam.Name = name
	dataStorePreparePostParam.DataType = stream.ReadUInt16LE()

	metaBinary, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	dataStorePreparePostParam.MetaBinary = metaBinary

	permission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return err
	}

	dataStorePreparePostParam.Permission = permission.(*DataStorePermission)

	delPermission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return err
	}

	dataStorePreparePostParam.DelPermission = delPermission.(*DataStorePermission)
	dataStorePreparePostParam.Flag = stream.ReadUInt32LE()
	dataStorePreparePostParam.Period = stream.ReadUInt16LE()
	dataStorePreparePostParam.ReferDataId = stream.ReadUInt32LE()
	dataStorePreparePostParam.Tags = stream.ReadListString()

	ratingInitParams, err := ReadListDataStoreRatingInitParamWithSlot(stream)
	if err != nil {
		return err
	}

	dataStorePreparePostParam.RatingInitParams = ratingInitParams

	persistenceInitParam, err := stream.ReadStructure(NewDataStorePersistenceInitParam())
	if err != nil {
		return err
	}

	dataStorePreparePostParam.PersistenceInitParam = persistenceInitParam.(*DataStorePersistenceInitParam)
	dataStorePreparePostParam.ExtraData = stream.ReadListString()

	return nil
}

// NewDataStorePreparePostParam returns a new DataStorePreparePostParam
func NewDataStorePreparePostParam() *DataStorePreparePostParam {
	return &DataStorePreparePostParam{}
}

// DataStoreSearchParam is sent in DataStore search methods
type DataStoreSearchParam struct {
	SearchTarget           uint8
	OwnerIds               []uint32
	OwnerType              uint8
	DestinationIds         []uint64
	DataType               uint16
	CreatedAfter           *nex.DateTime
	CreatedBefore          *nex.DateTime
	UpdatedAfter           *nex.DateTime
	UpdatedBefore          *nex.DateTime
	ReferDataId            uint32
	Tags                   []string
	ResultOrderColumn      uint8
	ResultOrder            uint8
	ResultRange            *nex.ResultRange
	ResultOption           uint8
	MinimalRatingFrequency uint32
	UseCache               bool
	nex.Structure
}

// ExtractFromStream extracts a DataStoreSearchParam structure from a stream
func (dataStoreSearchParam *DataStoreSearchParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreSearchParam.SearchTarget = stream.ReadUInt8()
	dataStoreSearchParam.OwnerIds = stream.ReadListUInt32LE()
	dataStoreSearchParam.OwnerType = stream.ReadUInt8()
	dataStoreSearchParam.DestinationIds = stream.ReadListUInt64LE()
	dataStoreSearchParam.DataType = stream.ReadUInt16LE()
	dataStoreSearchParam.CreatedAfter = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreSearchParam.CreatedBefore = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreSearchParam.UpdatedAfter = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreSearchParam.UpdatedBefore = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreSearchParam.ReferDataId = stream.ReadUInt32LE()
	dataStoreSearchParam.Tags = stream.ReadListString()
	dataStoreSearchParam.ResultOrderColumn = stream.ReadUInt8()
	dataStoreSearchParam.ResultOrder = stream.ReadUInt8()

	resultRange, err := stream.ReadStructure(nex.NewResultRange())

	if err != nil {
		return err
	}

	dataStoreSearchParam.ResultRange = resultRange.(*nex.ResultRange)
	dataStoreSearchParam.ResultOption = stream.ReadUInt8()
	dataStoreSearchParam.MinimalRatingFrequency = stream.ReadUInt32LE()
	dataStoreSearchParam.UseCache = (stream.ReadUInt8() == 1)

	return nil
}

// NewDataStoreSearchParam returns a new DataStoreSearchParam
func NewDataStoreSearchParam() *DataStoreSearchParam {
	return &DataStoreSearchParam{}
}

// DataStoreGetMetaParam is sent in the GetMeta method
type DataStoreGetMetaParam struct {
	nex.Structure
	DataID            uint64
	PersistenceTarget *DataStorePersistenceTarget
	ResultOption      uint8
	AccessPassword    uint64
}

// ExtractFromStream extracts a DataStoreGetMetaParam structure from a stream
func (dataStoreGetMetaParam *DataStoreGetMetaParam) ExtractFromStream(stream *nex.StreamIn) error {
	expectedDataSize := 23 // base size not including Structure header

	if len(stream.Bytes()[stream.ByteOffset():]) < expectedDataSize {
		return errors.New("[DataStoreGetMetaParam::ExtractFromStream] Data size too small")
	}

	dataID := stream.ReadUInt64LE()
	persistenceTarget, err := stream.ReadStructure(NewDataStorePersistenceTarget())

	if err != nil {
		return err
	}

	dataStoreGetMetaParam.DataID = dataID
	dataStoreGetMetaParam.PersistenceTarget = persistenceTarget.(*DataStorePersistenceTarget)
	dataStoreGetMetaParam.ResultOption = stream.ReadUInt8()
	dataStoreGetMetaParam.AccessPassword = stream.ReadUInt64LE()

	return nil
}

// NewDataStoreGetMetaParam returns a new DataStoreGetMetaParam
func NewDataStoreGetMetaParam() *DataStoreGetMetaParam {
	return &DataStoreGetMetaParam{}
}

// DataStoreChangeMetaParam is sent in the ChangeMeta method
type DataStoreChangeMetaParam struct {
	nex.Structure
	DataID         uint64
	ModifiesFlag   uint32
	Name           string
	Permission     *DataStorePermission
	DelPermission  *DataStorePermission
	Period         uint16
	MetaBinary     []byte
	Tags           []string
	UpdatePassword uint64
	ReferredCnt    uint32
	DataType       uint16
	Status         uint8
	CompareParam   *DataStoreChangeMetaCompareParam
	//PersistenceTarget *DataStorePersistenceTarget (not seen in SMM1??)
}

// ExtractFromStream extracts a DataStoreChangeMetaParam structure from a stream
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO: Check size

	dataStoreChangeMetaParam.DataID = stream.ReadUInt64LE()
	dataStoreChangeMetaParam.ModifiesFlag = stream.ReadUInt32LE()

	name, err := stream.ReadString()

	if err != nil {
		return err
	}

	dataStoreChangeMetaParam.Name = name

	permission, err := stream.ReadStructure(NewDataStorePermission())

	if err != nil {
		return err
	}

	dataStoreChangeMetaParam.Permission = permission.(*DataStorePermission)

	delPermission, err := stream.ReadStructure(NewDataStorePermission())

	if err != nil {
		return err
	}

	dataStoreChangeMetaParam.DelPermission = delPermission.(*DataStorePermission)
	dataStoreChangeMetaParam.Period = stream.ReadUInt16LE()

	metaBinary, err := stream.ReadQBuffer()

	if err != nil {
		return err
	}

	dataStoreChangeMetaParam.MetaBinary = metaBinary
	dataStoreChangeMetaParam.Tags = stream.ReadListString()
	dataStoreChangeMetaParam.UpdatePassword = stream.ReadUInt64LE()
	dataStoreChangeMetaParam.ReferredCnt = stream.ReadUInt32LE()
	dataStoreChangeMetaParam.DataType = stream.ReadUInt16LE()
	dataStoreChangeMetaParam.Status = stream.ReadUInt8()

	compareParam, err := stream.ReadStructure(NewDataStoreChangeMetaCompareParam())

	if err != nil {
		return err
	}

	dataStoreChangeMetaParam.CompareParam = compareParam.(*DataStoreChangeMetaCompareParam)

	/*
		persistenceTarget, err := stream.ReadStructure(NewDataStorePersistenceTarget())

		if err != nil {
			return err
		}

		dataStoreChangeMetaParam.PersistenceTarget = persistenceTarget.(*DataStorePersistenceTarget)
	*/

	return nil
}

// NewDataStoreChangeMetaParam returns a new DataStoreChangeMetaParam
func NewDataStoreChangeMetaParam() *DataStoreChangeMetaParam {
	return &DataStoreChangeMetaParam{}
}

// DataStoreChangeMetaCompareParam is sent in the ChangeMeta method
type DataStoreChangeMetaCompareParam struct {
	nex.Structure
	ComparisonFlag uint32
	Name           string
	Permission     *DataStorePermission
	DelPermission  *DataStorePermission
	Period         uint16
	MetaBinary     []byte
	Tags           []string
	ReferredCnt    uint32
	DataType       uint16
	Status         uint8
}

// ExtractFromStream extracts a DataStoreChangeMetaCompareParam structure from a stream
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) ExtractFromStream(stream *nex.StreamIn) error {
	// TODO: Check size

	dataStoreChangeMetaCompareParam.ComparisonFlag = stream.ReadUInt32LE()

	name, err := stream.ReadString()

	if err != nil {
		return err
	}

	dataStoreChangeMetaCompareParam.Name = name

	permission, err := stream.ReadStructure(NewDataStorePermission())

	if err != nil {
		return err
	}

	dataStoreChangeMetaCompareParam.Permission = permission.(*DataStorePermission)

	delPermission, err := stream.ReadStructure(NewDataStorePermission())

	if err != nil {
		return err
	}

	dataStoreChangeMetaCompareParam.DelPermission = delPermission.(*DataStorePermission)
	dataStoreChangeMetaCompareParam.Period = stream.ReadUInt16LE()

	metaBinary, err := stream.ReadQBuffer()

	if err != nil {
		return err
	}

	dataStoreChangeMetaCompareParam.MetaBinary = metaBinary
	dataStoreChangeMetaCompareParam.Tags = stream.ReadListString()
	dataStoreChangeMetaCompareParam.ReferredCnt = stream.ReadUInt32LE()
	dataStoreChangeMetaCompareParam.DataType = stream.ReadUInt16LE()
	dataStoreChangeMetaCompareParam.Status = stream.ReadUInt8()

	return nil
}

// NewDataStoreChangeMetaCompareParam returns a new DataStoreChangeMetaParam
func NewDataStoreChangeMetaCompareParam() *DataStoreChangeMetaCompareParam {
	return &DataStoreChangeMetaCompareParam{}
}

// DataStorePermission contains information about a permission for a DataStore object
type DataStorePermission struct {
	nex.Structure
	Permission   uint8
	RecipientIds []uint32
}

// ExtractFromStream extracts a DataStorePermission structure from a stream
func (dataStorePermission *DataStorePermission) ExtractFromStream(stream *nex.StreamIn) error {
	expectedDataSize := 9 // base size not including Structure header

	if len(stream.Bytes()[stream.ByteOffset():]) < expectedDataSize {
		return errors.New("[DataStorePermission::ExtractFromStream] Data size too small")
	}

	dataStorePermission.Permission = stream.ReadUInt8()
	dataStorePermission.RecipientIds = stream.ReadListUInt32LE()

	return nil
}

// Bytes encodes the DataStorePermission and returns a byte array
func (dataStorePermission *DataStorePermission) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(dataStorePermission.Permission)
	stream.WriteListUInt32LE(dataStorePermission.RecipientIds)

	return stream.Bytes()
}

// NewDataStorePermission returns a new DataStorePermission
func NewDataStorePermission() *DataStorePermission {
	return &DataStorePermission{}
}

// DataStorePersistenceTarget contains information about a DataStore target
type DataStorePersistenceTarget struct {
	nex.Structure
	OwnerID           uint32
	PersistenceSlotID uint16
}

// ExtractFromStream extracts a DataStorePersistenceTarget structure from a stream
func (dataStorePersistenceTarget *DataStorePersistenceTarget) ExtractFromStream(stream *nex.StreamIn) error {
	expectedDataSize := 9 // base size not including Structure header

	if len(stream.Bytes()[stream.ByteOffset():]) < expectedDataSize {
		return errors.New("[DataStorePersistenceTarget::ExtractFromStream] Data size too small")
	}

	dataStorePersistenceTarget.OwnerID = stream.ReadUInt32LE()
	dataStorePersistenceTarget.PersistenceSlotID = stream.ReadUInt16LE()

	return nil
}

// NewDataStorePersistenceTarget returns a new DataStorePersistenceTarget
func NewDataStorePersistenceTarget() *DataStorePersistenceTarget {
	return &DataStorePersistenceTarget{}
}

type DataStoreRatingInfo struct {
	nex.Structure
	TotalValue   int64
	Count        uint32
	InitialValue int64
}

// ExtractFromStream extracts a DataStoreRatingInfo structure from a stream
func (dataStoreRatingInfo *DataStoreRatingInfo) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreRatingInfo.TotalValue = int64(stream.ReadUInt64LE())
	dataStoreRatingInfo.Count = stream.ReadUInt32LE()
	dataStoreRatingInfo.InitialValue = int64(stream.ReadUInt64LE())

	return nil
}

// Bytes encodes the DataStoreRatingInfo and returns a byte array
func (dataStoreRatingInfo *DataStoreRatingInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(uint64(dataStoreRatingInfo.TotalValue))
	stream.WriteUInt32LE(dataStoreRatingInfo.Count)
	stream.WriteUInt64LE(uint64(dataStoreRatingInfo.InitialValue))

	return stream.Bytes()
}

// NewDataStoreRatingInfo returns a new DataStoreRatingInfo
func NewDataStoreRatingInfo() *DataStoreRatingInfo {
	return &DataStoreRatingInfo{}
}

type DataStoreRatingInfoWithSlot struct {
	nex.Structure
	Slot   int8
	Rating *DataStoreRatingInfo
}

// ExtractFromStream extracts a DataStoreRatingInfoWithSlot structure from a stream
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreRatingInfoWithSlot.Slot = int8(stream.ReadUInt8())

	rating, err := stream.ReadStructure(NewDataStoreRatingInfo())

	if err != nil {
		return err
	}

	dataStoreRatingInfoWithSlot.Rating = rating.(*DataStoreRatingInfo)

	return nil
}

// Bytes encodes the DataStoreRatingInfoWithSlot and returns a byte array
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(uint8(dataStoreRatingInfoWithSlot.Slot))
	stream.WriteStructure(dataStoreRatingInfoWithSlot.Rating)

	return stream.Bytes()
}

// NewDataStoreRatingInfoWithSlot returns a new DataStoreRatingInfoWithSlot
func NewDataStoreRatingInfoWithSlot() *DataStoreRatingInfoWithSlot {
	return &DataStoreRatingInfoWithSlot{}
}

// DataStoreMetaInfo contains DataStore meta information
type DataStoreMetaInfo struct {
	nex.Structure
	DataID        uint64
	OwnerID       uint32
	Size          uint32
	DataType      uint16
	Name          string
	MetaBinary    []byte
	Permission    *DataStorePermission
	DelPermission *DataStorePermission
	CreatedTime   *nex.DateTime
	UpdatedTime   *nex.DateTime
	Period        uint16
	Status        uint8
	ReferredCnt   uint32
	ReferDataID   uint32
	Flag          uint32
	ReferredTime  *nex.DateTime
	ExpireTime    *nex.DateTime
	Tags          []string
	Ratings       []*DataStoreRatingInfoWithSlot
}

// ExtractFromStream extracts a DataStoreMetaInfo structure from a stream
func (dataStoreMetaInfo *DataStoreMetaInfo) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreMetaInfo.DataID = stream.ReadUInt64LE()
	dataStoreMetaInfo.OwnerID = stream.ReadUInt32LE()
	dataStoreMetaInfo.Size = stream.ReadUInt32LE()

	name, err := stream.ReadString()
	if err != nil {
		return err
	}

	dataStoreMetaInfo.Name = name
	dataStoreMetaInfo.DataType = stream.ReadUInt16LE()

	metaBinary, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	dataStoreMetaInfo.MetaBinary = metaBinary

	permission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return err
	}

	dataStoreMetaInfo.Permission = permission.(*DataStorePermission)

	delPermission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return err
	}

	dataStoreMetaInfo.DelPermission = delPermission.(*DataStorePermission)
	dataStoreMetaInfo.CreatedTime = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreMetaInfo.UpdatedTime = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreMetaInfo.Period = stream.ReadUInt16LE()
	dataStoreMetaInfo.Status = stream.ReadUInt8()
	dataStoreMetaInfo.ReferredCnt = stream.ReadUInt32LE()
	dataStoreMetaInfo.ReferDataID = stream.ReadUInt32LE()
	dataStoreMetaInfo.Flag = stream.ReadUInt32LE()
	dataStoreMetaInfo.ReferredTime = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreMetaInfo.ExpireTime = nex.NewDateTime(stream.ReadUInt64LE())
	dataStoreMetaInfo.Tags = stream.ReadListString()

	ratings, err := ReadListDataStoreRatingInfoWithSlot(stream)
	if err != nil {
		return err
	}

	dataStoreMetaInfo.Ratings = ratings

	return nil
}

// Bytes encodes the DataStoreMetaInfo and returns a byte array
func (dataStoreMetaInfo *DataStoreMetaInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(dataStoreMetaInfo.DataID)
	stream.WriteUInt32LE(dataStoreMetaInfo.OwnerID)
	stream.WriteUInt32LE(dataStoreMetaInfo.Size)
	stream.WriteString(dataStoreMetaInfo.Name)
	stream.WriteUInt16LE(dataStoreMetaInfo.DataType)
	stream.WriteQBuffer(dataStoreMetaInfo.MetaBinary)
	stream.WriteStructure(dataStoreMetaInfo.Permission)
	stream.WriteStructure(dataStoreMetaInfo.DelPermission)
	stream.WriteUInt64LE(dataStoreMetaInfo.CreatedTime.Value())
	stream.WriteUInt64LE(dataStoreMetaInfo.UpdatedTime.Value())
	stream.WriteUInt16LE(dataStoreMetaInfo.Period)
	stream.WriteUInt8(dataStoreMetaInfo.Status)
	stream.WriteUInt32LE(dataStoreMetaInfo.ReferredCnt)
	stream.WriteUInt32LE(dataStoreMetaInfo.ReferDataID)
	stream.WriteUInt32LE(dataStoreMetaInfo.Flag)
	stream.WriteUInt64LE(dataStoreMetaInfo.ReferredTime.Value())
	stream.WriteUInt64LE(dataStoreMetaInfo.ExpireTime.Value())
	stream.WriteListString(dataStoreMetaInfo.Tags)
	stream.WriteListStructure(dataStoreMetaInfo.Ratings)

	return stream.Bytes()
}

// NewDataStoreMetaInfo returns a new DataStoreMetaInfo
func NewDataStoreMetaInfo() *DataStoreMetaInfo {
	return &DataStoreMetaInfo{}
}

// DataStorePrepareGetParam is sent in the PrepareGetObject method
type DataStorePrepareGetParam struct {
	nex.Structure
	DataID            uint64
	LockID            uint32
	PersistenceTarget *DataStorePersistenceTarget
	AccessPassword    uint64
	ExtraData         []string
}

// ExtractFromStream extracts a DataStorePrepareGetParam structure from a stream
func (dataStorePrepareGetParam *DataStorePrepareGetParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStorePrepareGetParam.DataID = stream.ReadUInt64LE()
	dataStorePrepareGetParam.LockID = stream.ReadUInt32LE()

	persistenceTarget, err := stream.ReadStructure(NewDataStorePersistenceTarget())
	if err != nil {
		return err
	}

	dataStorePrepareGetParam.PersistenceTarget = persistenceTarget.(*DataStorePersistenceTarget)
	dataStorePrepareGetParam.AccessPassword = stream.ReadUInt64LE()
	dataStorePrepareGetParam.ExtraData = stream.ReadListString()

	return nil
}

// NewDataStorePrepareGetParam returns a new DataStorePrepareGetParam
func NewDataStorePrepareGetParam() *DataStorePrepareGetParam {
	return &DataStorePrepareGetParam{}
}

// DataStoreKeyValue is sent in the PrepareGetObject method
type DataStoreKeyValue struct {
	nex.Structure
	Key   string
	Value string
}

// Bytes encodes the DataStoreKeyValue and returns a byte array
func (dataStoreKeyValue *DataStoreKeyValue) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreKeyValue.Key)
	stream.WriteString(dataStoreKeyValue.Value)

	return stream.Bytes()
}

// NewDataStoreKeyValue returns a new DataStoreKeyValue
func NewDataStoreKeyValue() *DataStoreKeyValue {
	return &DataStoreKeyValue{}
}

// DataStoreReqGetInfo is sent in the PrepareGetObject method
type DataStoreReqGetInfo struct {
	nex.Structure
	URL            string
	RequestHeaders []*DataStoreKeyValue
	Size           uint32
	RootCA         []byte
	DataID         uint64
}

// Bytes encodes the DataStoreReqGetInfo and returns a byte array
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(dataStoreReqGetInfo.URL)
	stream.WriteListStructure(dataStoreReqGetInfo.RequestHeaders)
	stream.WriteUInt32LE(dataStoreReqGetInfo.Size)
	stream.WriteBuffer(dataStoreReqGetInfo.RootCA)
	stream.WriteUInt64LE(dataStoreReqGetInfo.DataID)

	return stream.Bytes()
}

// NewDataStoreReqGetInfo returns a new DataStoreReqGetInfo
func NewDataStoreReqGetInfo() *DataStoreReqGetInfo {
	return &DataStoreReqGetInfo{}
}

// Setup initializes the protocol
func (dataStoreProtocol *DataStoreProtocol) Setup() {
	nexServer := dataStoreProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if DataStoreProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case DataStoreMethodGetMeta:
				go dataStoreProtocol.handleGetMeta(packet)
			case DataStoreMethodPrepareUpdateObject:
				go dataStoreProtocol.handlePrepareUpdateObject(packet)
			case DataStoreMethodCompleteUpdateObject:
				go dataStoreProtocol.handleCompleteUpdateObject(packet)
			case DataStoreMethodPostMetaBinary:
				go dataStoreProtocol.handlePostMetaBinary(packet)
			case DataStoreMethodPreparePostObject:
				go dataStoreProtocol.handlePreparePostObject(packet)
			case DataStoreMethodPrepareGetObject:
				go dataStoreProtocol.handlePrepareGetObject(packet)
			case DataStoreMethodCompletePostObject:
				go dataStoreProtocol.handleCompletePostObject(packet)
			case DataStoreMethodGetPersistenceInfo:
				go dataStoreProtocol.handleGetPersistenceInfo(packet)
			case DataStoreMethodGetMetasMultipleParam:
				go dataStoreProtocol.handleGetMetasMultipleParam(packet)
			case DataStoreMethodCompletePostObjects:
				go dataStoreProtocol.handleCompletePostObjects(packet)
			case DataStoreMethodChangeMeta:
				go dataStoreProtocol.handleChangeMeta(packet)
			case DataStoreMethodRateObjects:
				go dataStoreProtocol.handleRateObjects(packet)
			default:
				go respondNotImplemented(packet, DataStoreProtocolID)
				fmt.Printf("Unsupported DataStore method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// GetMeta sets the GetMeta handler function
func (dataStoreProtocol *DataStoreProtocol) GetMeta(handler func(err error, client *nex.Client, callID uint32, dataStoreGetMetaParam *DataStoreGetMetaParam)) {
	dataStoreProtocol.GetMetaHandler = handler
}

// PrepareUpdateObject sets the PrepareUpdateObject handler function
func (dataStoreProtocol *DataStoreProtocol) PrepareUpdateObject(handler func(err error, client *nex.Client, callID uint32, dataStorePrepareUpdateParam *DataStorePrepareUpdateParam)) {
	dataStoreProtocol.PrepareUpdateObjectHandler = handler
}

// CompleteUpdateObject sets the CompleteUpdateObject handler function
func (dataStoreProtocol *DataStoreProtocol) CompleteUpdateObject(handler func(err error, client *nex.Client, callID uint32, dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam)) {
	dataStoreProtocol.CompleteUpdateObjectHandler = handler
}

// PostMetaBinary sets the PostMetaBinary handler function
func (dataStoreProtocol *DataStoreProtocol) PostMetaBinary(handler func(err error, client *nex.Client, callID uint32, dataStorePreparePostParam *DataStorePreparePostParam)) {
	dataStoreProtocol.PostMetaBinaryHandler = handler
}

// PreparePostObject sets the PreparePostObject handler function
func (dataStoreProtocol *DataStoreProtocol) PreparePostObject(handler func(err error, client *nex.Client, callID uint32, dataStorePreparePostParam *DataStorePreparePostParam)) {
	dataStoreProtocol.PreparePostObjectHandler = handler
}

// PrepareGetObject sets the PrepareGetObject handler function
func (dataStoreProtocol *DataStoreProtocol) PrepareGetObject(handler func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParam *DataStorePrepareGetParam)) {
	dataStoreProtocol.PrepareGetObjectHandler = handler
}

// CompletePostObject sets the CompletePostObject handler function
func (dataStoreProtocol *DataStoreProtocol) CompletePostObject(handler func(err error, client *nex.Client, callID uint32, dataStoreCompletePostParam *DataStoreCompletePostParam)) {
	dataStoreProtocol.CompletePostObjectHandler = handler
}

// GetPersistenceInfo sets the GetPersistenceInfo handler function
func (dataStoreProtocol *DataStoreProtocol) GetPersistenceInfo(handler func(err error, client *nex.Client, callID uint32, ownerID uint32, persistenceSlotID uint16)) {
	dataStoreProtocol.GetPersistenceInfoHandler = handler
}

// GetMetasMultipleParam sets the GetMetasMultipleParam handler function
func (dataStoreProtocol *DataStoreProtocol) GetMetasMultipleParam(handler func(err error, client *nex.Client, callID uint32, dataStorePrepareGetParams []*DataStoreGetMetaParam)) {
	dataStoreProtocol.GetMetasMultipleParamHandler = handler
}

// ChangeMeta sets the ChangeMeta handler function
func (dataStoreProtocol *DataStoreProtocol) CompletePostObjects(handler func(err error, client *nex.Client, callID uint32, dataIDs []uint64)) {
	dataStoreProtocol.CompletePostObjectsHandler = handler
}

// ChangeMeta sets the ChangeMeta handler function
func (dataStoreProtocol *DataStoreProtocol) ChangeMeta(handler func(err error, client *nex.Client, callID uint32, dataStoreChangeMetaParam *DataStoreChangeMetaParam)) {
	dataStoreProtocol.ChangeMetaHandler = handler
}

// RateObjects sets the RateObjects handler function
func (dataStoreProtocol *DataStoreProtocol) RateObjects(handler func(err error, client *nex.Client, callID uint32, targets []*DataStoreRatingTarget, params []*DataStoreRateObjectParam, transactional bool, fetchRatings bool)) {
	dataStoreProtocol.RateObjectsHandler = handler
}

func (dataStoreProtocol *DataStoreProtocol) handleGetMeta(packet nex.PacketInterface) {
	if dataStoreProtocol.GetMetaHandler == nil {
		logger.Warning("DataStoreProtocol::GetMeta not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStoreGetMetaParam, err := parametersStream.ReadStructure(NewDataStoreGetMetaParam())

	if err != nil {
		go dataStoreProtocol.GetMetaHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.GetMetaHandler(nil, client, callID, dataStoreGetMetaParam.(*DataStoreGetMetaParam))
}

func (dataStoreProtocol *DataStoreProtocol) handlePrepareUpdateObject(packet nex.PacketInterface) {
	if dataStoreProtocol.PrepareUpdateObjectHandler == nil {
		logger.Warning("DataStoreProtocol::PrepareUpdateObject not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStorePrepareUpdateParam, err := parametersStream.ReadStructure(NewDataStorePrepareUpdateParam())
	if err != nil {
		go dataStoreProtocol.PrepareUpdateObjectHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.PrepareUpdateObjectHandler(nil, client, callID, dataStorePrepareUpdateParam.(*DataStorePrepareUpdateParam))
}

func (dataStoreProtocol *DataStoreProtocol) handleCompleteUpdateObject(packet nex.PacketInterface) {
	if dataStoreProtocol.CompleteUpdateObjectHandler == nil {
		logger.Warning("DataStoreProtocol::CompleteUpdateObject not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStoreCompleteUpdateParam, err := parametersStream.ReadStructure(NewDataStoreCompleteUpdateParam())
	if err != nil {
		go dataStoreProtocol.CompleteUpdateObjectHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.CompleteUpdateObjectHandler(nil, client, callID, dataStoreCompleteUpdateParam.(*DataStoreCompleteUpdateParam))
}

func (dataStoreProtocol *DataStoreProtocol) handlePostMetaBinary(packet nex.PacketInterface) {
	if dataStoreProtocol.PostMetaBinaryHandler == nil {
		logger.Warning("DataStoreProtocol::PostMetaBinary not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStorePreparePostParam, err := parametersStream.ReadStructure(NewDataStorePreparePostParam())
	if err != nil {
		go dataStoreProtocol.PostMetaBinaryHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.PostMetaBinaryHandler(nil, client, callID, dataStorePreparePostParam.(*DataStorePreparePostParam))
}

func (dataStoreProtocol *DataStoreProtocol) handlePreparePostObject(packet nex.PacketInterface) {
	if dataStoreProtocol.PreparePostObjectHandler == nil {
		logger.Warning("DataStoreProtocol::PreparePostObject not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStorePreparePostParam, err := parametersStream.ReadStructure(NewDataStorePreparePostParam())
	if err != nil {
		go dataStoreProtocol.PreparePostObjectHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.PreparePostObjectHandler(nil, client, callID, dataStorePreparePostParam.(*DataStorePreparePostParam))
}

func (dataStoreProtocol *DataStoreProtocol) handlePrepareGetObject(packet nex.PacketInterface) {
	if dataStoreProtocol.PrepareGetObjectHandler == nil {
		logger.Warning("DataStoreProtocol::PrepareGetObject not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStorePrepareGetParam, err := parametersStream.ReadStructure(NewDataStorePrepareGetParam())

	if err != nil {
		go dataStoreProtocol.PrepareGetObjectHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.PrepareGetObjectHandler(nil, client, callID, dataStorePrepareGetParam.(*DataStorePrepareGetParam))
}

func (dataStoreProtocol *DataStoreProtocol) handleCompletePostObject(packet nex.PacketInterface) {
	if dataStoreProtocol.CompletePostObjectHandler == nil {
		logger.Warning("DataStoreProtocol::CompletePostObject not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStoreCompletePostParam, err := parametersStream.ReadStructure(NewDataStoreCompletePostParam())
	if err != nil {
		go dataStoreProtocol.CompletePostObjectHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.CompletePostObjectHandler(nil, client, callID, dataStoreCompletePostParam.(*DataStoreCompletePostParam))
}

func (dataStoreProtocol *DataStoreProtocol) handleGetPersistenceInfo(packet nex.PacketInterface) {
	if dataStoreProtocol.GetPersistenceInfoHandler == nil {
		logger.Warning("DataStorePotocol::GetPersistenceInfo not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	if len(parametersStream.Bytes()[parametersStream.ByteOffset():]) < 6 {
		err := errors.New("[DataStoreProtocol::GetPersistenceInfo] Data length too small")
		go dataStoreProtocol.GetPersistenceInfoHandler(err, client, callID, 0, 0)
		return
	}

	ownerID := parametersStream.ReadUInt32LE()
	persistenceSlotID := parametersStream.ReadUInt16LE()

	go dataStoreProtocol.GetPersistenceInfoHandler(nil, client, callID, ownerID, persistenceSlotID)
}

func (dataStoreProtocol *DataStoreProtocol) handleGetMetasMultipleParam(packet nex.PacketInterface) {
	if dataStoreProtocol.GetMetasMultipleParamHandler == nil {
		logger.Warning("DataStoreProtocol::GetMetasMultipleParam not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStoreGetMetaParams, err := ReaListDataStoreGetMetaParam(parametersStream)

	if err != nil {
		go dataStoreProtocol.GetMetasMultipleParamHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.GetMetasMultipleParamHandler(nil, client, callID, dataStoreGetMetaParams)
}

func (dataStoreProtocol *DataStoreProtocol) handleCompletePostObjects(packet nex.PacketInterface) {
	if dataStoreProtocol.CompletePostObjectsHandler == nil {
		logger.Warning("DataStoreProtocol::CompletePostObjects not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataIDs := parametersStream.ReadListUInt64LE()

	go dataStoreProtocol.CompletePostObjectsHandler(nil, client, callID, dataIDs)
}

func (dataStoreProtocol *DataStoreProtocol) handleChangeMeta(packet nex.PacketInterface) {
	if dataStoreProtocol.ChangeMetaHandler == nil {
		logger.Warning("DataStoreProtocol::ChangeMeta not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	dataStoreChangeMetaParam, err := parametersStream.ReadStructure(NewDataStoreChangeMetaParam())

	if err != nil {
		go dataStoreProtocol.ChangeMetaHandler(err, client, callID, nil)
		return
	}

	go dataStoreProtocol.ChangeMetaHandler(nil, client, callID, dataStoreChangeMetaParam.(*DataStoreChangeMetaParam))
}

func (dataStoreProtocol *DataStoreProtocol) handleRateObjects(packet nex.PacketInterface) {
	if dataStoreProtocol.RateObjectsHandler == nil {
		logger.Warning("DataStoreProtocol::RateObjects not implemented")
		go respondNotImplemented(packet, DataStoreProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, dataStoreProtocol.server)

	targets, err := ReadListDataStoreRatingTarget(parametersStream)
	if err != nil {
		go dataStoreProtocol.RateObjectsHandler(err, client, callID, nil, nil, false, false)
		return
	}

	params, err := ReadListDataStoreRateObjectParam(parametersStream)
	if err != nil {
		go dataStoreProtocol.RateObjectsHandler(err, client, callID, nil, nil, false, false)
		return
	}

	transactional := (parametersStream.ReadUInt8() == 1)
	fetchRatings := (parametersStream.ReadUInt8() == 1)

	go dataStoreProtocol.RateObjectsHandler(nil, client, callID, targets, params, transactional, fetchRatings)
}

// NewDataStoreProtocol returns a new DataStoreProtocol
func NewDataStoreProtocol(server *nex.Server) *DataStoreProtocol {
	dataStoreProtocol := &DataStoreProtocol{server: server}

	dataStoreProtocol.Setup()

	return dataStoreProtocol
}
