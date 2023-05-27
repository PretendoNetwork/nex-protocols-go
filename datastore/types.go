package datastore

import (
	"bytes"
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
)

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

// Copy returns a new copied instance of DataStoreNotificationV1
func (dataStoreNotificationV1 *DataStoreNotificationV1) Copy() nex.StructureInterface {
	copied := NewDataStoreNotificationV1()

	copied.NotificationID = dataStoreNotificationV1.NotificationID
	copied.DataID = dataStoreNotificationV1.DataID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreNotificationV1 *DataStoreNotificationV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreNotificationV1)

	if dataStoreNotificationV1.NotificationID != other.NotificationID {
		return false
	}

	if dataStoreNotificationV1.DataID != other.DataID {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreNotification
func (dataStoreNotification *DataStoreNotification) Copy() nex.StructureInterface {
	copied := NewDataStoreNotification()

	copied.NotificationID = dataStoreNotification.NotificationID
	copied.DataID = dataStoreNotification.DataID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreNotification *DataStoreNotification) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreNotification)

	if dataStoreNotification.NotificationID != other.NotificationID {
		return false
	}

	if dataStoreNotification.DataID != other.DataID {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreGetSpecificMetaParamV1
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) Copy() nex.StructureInterface {
	copied := NewDataStoreGetSpecificMetaParamV1()

	copied.DataIDs = make([]uint32, len(dataStoreGetSpecificMetaParamV1.DataIDs))

	copy(copied.DataIDs, dataStoreGetSpecificMetaParamV1.DataIDs)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetSpecificMetaParamV1 *DataStoreGetSpecificMetaParamV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetSpecificMetaParamV1)

	if len(dataStoreGetSpecificMetaParamV1.DataIDs) != len(other.DataIDs) {
		return false
	}

	for i := 0; i < len(dataStoreGetSpecificMetaParamV1.DataIDs); i++ {
		if dataStoreGetSpecificMetaParamV1.DataIDs[i] != other.DataIDs[i] {
			return false
		}
	}

	return true
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

// Copy returns a new copied instance of DataStoreGetSpecificMetaParam
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetSpecificMetaParam()

	copied.DataIDs = make([]uint64, len(dataStoreGetSpecificMetaParam.DataIDs))

	copy(copied.DataIDs, dataStoreGetSpecificMetaParam.DataIDs)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetSpecificMetaParam *DataStoreGetSpecificMetaParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetSpecificMetaParam)

	if len(dataStoreGetSpecificMetaParam.DataIDs) != len(other.DataIDs) {
		return false
	}

	for i := 0; i < len(dataStoreGetSpecificMetaParam.DataIDs); i++ {
		if dataStoreGetSpecificMetaParam.DataIDs[i] != other.DataIDs[i] {
			return false
		}
	}

	return true
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

// Copy returns a new copied instance of DataStoreSpecificMetaInfoV1
func (dataStoreSpecificMetaInfoV1 *DataStoreSpecificMetaInfoV1) Copy() nex.StructureInterface {
	copied := NewDataStoreSpecificMetaInfoV1()

	copied.DataID = dataStoreSpecificMetaInfoV1.DataID
	copied.OwnerID = dataStoreSpecificMetaInfoV1.OwnerID
	copied.Size = dataStoreSpecificMetaInfoV1.Size
	copied.DataType = dataStoreSpecificMetaInfoV1.DataType
	copied.Version = dataStoreSpecificMetaInfoV1.Version

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSpecificMetaInfoV1 *DataStoreSpecificMetaInfoV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSpecificMetaInfoV1)

	if dataStoreSpecificMetaInfoV1.DataID != other.DataID {
		return false
	}

	if dataStoreSpecificMetaInfoV1.OwnerID != other.OwnerID {
		return false
	}

	if dataStoreSpecificMetaInfoV1.Size != other.Size {
		return false
	}

	if dataStoreSpecificMetaInfoV1.DataType != other.DataType {
		return false
	}

	if dataStoreSpecificMetaInfoV1.Version != other.Version {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreSpecificMetaInfo
func (dataStoreSpecificMetaInfo *DataStoreSpecificMetaInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreSpecificMetaInfo()

	copied.DataID = dataStoreSpecificMetaInfo.DataID
	copied.OwnerID = dataStoreSpecificMetaInfo.OwnerID
	copied.Size = dataStoreSpecificMetaInfo.Size
	copied.DataType = dataStoreSpecificMetaInfo.DataType
	copied.Version = dataStoreSpecificMetaInfo.Version

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSpecificMetaInfo *DataStoreSpecificMetaInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSpecificMetaInfo)

	if dataStoreSpecificMetaInfo.DataID != other.DataID {
		return false
	}

	if dataStoreSpecificMetaInfo.OwnerID != other.OwnerID {
		return false
	}

	if dataStoreSpecificMetaInfo.Size != other.Size {
		return false
	}

	if dataStoreSpecificMetaInfo.DataType != other.DataType {
		return false
	}

	if dataStoreSpecificMetaInfo.Version != other.Version {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreTouchObjectParam
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) Copy() nex.StructureInterface {
	copied := NewDataStoreTouchObjectParam()

	copied.DataID = dataStoreTouchObjectParam.DataID
	copied.LockID = dataStoreTouchObjectParam.LockID
	copied.AccessPassword = dataStoreTouchObjectParam.AccessPassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreTouchObjectParam)

	if dataStoreTouchObjectParam.DataID != other.DataID {
		return false
	}

	if dataStoreTouchObjectParam.LockID != other.LockID {
		return false
	}

	if dataStoreTouchObjectParam.AccessPassword != other.AccessPassword {
		return false
	}

	return true
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
	stream.WriteDateTime(dataStoreRatingLog.LockExpirationTime)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreRatingLog
func (dataStoreRatingLog *DataStoreRatingLog) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingLog()

	copied.IsRated = dataStoreRatingLog.IsRated
	copied.Pid = dataStoreRatingLog.Pid
	copied.RatingValue = dataStoreRatingLog.RatingValue
	copied.LockExpirationTime = dataStoreRatingLog.LockExpirationTime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingLog *DataStoreRatingLog) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingLog)

	if dataStoreRatingLog.IsRated != other.IsRated {
		return false
	}

	if dataStoreRatingLog.Pid != other.Pid {
		return false
	}

	if dataStoreRatingLog.RatingValue != other.RatingValue {
		return false
	}

	if !dataStoreRatingLog.LockExpirationTime.Equals(other.LockExpirationTime) {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStorePersistenceInfo
func (dataStorePersistenceInfo *DataStorePersistenceInfo) Copy() nex.StructureInterface {
	copied := NewDataStorePersistenceInfo()

	copied.OwnerID = dataStorePersistenceInfo.OwnerID
	copied.PersistenceSlotID = dataStorePersistenceInfo.PersistenceSlotID
	copied.DataID = dataStorePersistenceInfo.DataID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePersistenceInfo *DataStorePersistenceInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePersistenceInfo)

	if dataStorePersistenceInfo.OwnerID != other.OwnerID {
		return false
	}

	if dataStorePersistenceInfo.PersistenceSlotID != other.PersistenceSlotID {
		return false
	}

	if dataStorePersistenceInfo.DataID != other.DataID {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStorePasswordInfo
func (dataStorePasswordInfo *DataStorePasswordInfo) Copy() nex.StructureInterface {
	copied := NewDataStorePasswordInfo()

	copied.DataID = dataStorePasswordInfo.DataID
	copied.AccessPassword = dataStorePasswordInfo.AccessPassword
	copied.UpdatePassword = dataStorePasswordInfo.UpdatePassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePasswordInfo *DataStorePasswordInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePasswordInfo)

	if dataStorePasswordInfo.DataID != other.DataID {
		return false
	}

	if dataStorePasswordInfo.AccessPassword != other.AccessPassword {
		return false
	}

	if dataStorePasswordInfo.UpdatePassword != other.UpdatePassword {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreGetNewArrivedNotificationsParam
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetNewArrivedNotificationsParam()

	copied.LastNotificationID = dataStoreGetNewArrivedNotificationsParam.LastNotificationID
	copied.Limit = dataStoreGetNewArrivedNotificationsParam.Limit

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetNewArrivedNotificationsParam *DataStoreGetNewArrivedNotificationsParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetNewArrivedNotificationsParam)

	if dataStoreGetNewArrivedNotificationsParam.LastNotificationID != other.LastNotificationID {
		return false
	}

	if dataStoreGetNewArrivedNotificationsParam.Limit != other.Limit {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreReqGetNotificationUrlInfo
func (dataStoreReqGetNotificationUrlInfo *DataStoreReqGetNotificationUrlInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReqGetNotificationUrlInfo()

	copied.Url = dataStoreReqGetNotificationUrlInfo.Url
	copied.Key = dataStoreReqGetNotificationUrlInfo.Key
	copied.Query = dataStoreReqGetNotificationUrlInfo.Query
	copied.RootCaCert = make([]byte, len(dataStoreReqGetNotificationUrlInfo.RootCaCert))

	copy(copied.RootCaCert, dataStoreReqGetNotificationUrlInfo.RootCaCert)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetNotificationUrlInfo *DataStoreReqGetNotificationUrlInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqGetNotificationUrlInfo)

	if dataStoreReqGetNotificationUrlInfo.Url != other.Url {
		return false
	}

	if dataStoreReqGetNotificationUrlInfo.Key != other.Key {
		return false
	}

	if dataStoreReqGetNotificationUrlInfo.Query != other.Query {
		return false
	}

	if !bytes.Equal(dataStoreReqGetNotificationUrlInfo.RootCaCert, other.RootCaCert) {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreGetNotificationUrlParam
func (dataStoreGetNotificationUrlParam *DataStoreGetNotificationUrlParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetNotificationUrlParam()

	copied.PreviousUrl = dataStoreGetNotificationUrlParam.PreviousUrl

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetNotificationUrlParam *DataStoreGetNotificationUrlParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetNotificationUrlParam)

	return dataStoreGetNotificationUrlParam.PreviousUrl != other.PreviousUrl
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

	result, err := stream.ReadListStructure(NewDataStoreMetaInfo())
	if err != nil {
		return err
	}

	dataStoreSearchResult.Result = result.([]*DataStoreMetaInfo)
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

// Copy returns a new copied instance of DataStoreSearchResult
func (dataStoreSearchResult *DataStoreSearchResult) Copy() nex.StructureInterface {
	copied := NewDataStoreSearchResult()

	copied.TotalCount = dataStoreSearchResult.TotalCount
	copied.Result = make([]*DataStoreMetaInfo, len(dataStoreSearchResult.Result))

	for i := 0; i < len(dataStoreSearchResult.Result); i++ {
		copied.Result[i] = dataStoreSearchResult.Result[i].Copy().(*DataStoreMetaInfo)
	}

	copied.TotalCountType = dataStoreSearchResult.TotalCountType

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchResult *DataStoreSearchResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSearchResult)

	if dataStoreSearchResult.TotalCount != other.TotalCount {
		return false
	}

	if len(dataStoreSearchResult.Result) != len(other.Result) {
		return false
	}

	for i := 0; i < len(dataStoreSearchResult.Result); i++ {
		if dataStoreSearchResult.Result[i] != other.Result[i] {
			return false
		}
	}

	if dataStoreSearchResult.TotalCountType != other.TotalCountType {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreCompleteUpdateParam
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) Copy() nex.StructureInterface {
	copied := NewDataStoreCompleteUpdateParam()

	copied.DataID = dataStoreCompleteUpdateParam.DataID
	copied.Version = dataStoreCompleteUpdateParam.Version
	copied.IsSuccess = dataStoreCompleteUpdateParam.IsSuccess

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompleteUpdateParam)

	if dataStoreCompleteUpdateParam.DataID != other.DataID {
		return false
	}

	if dataStoreCompleteUpdateParam.Version != other.Version {
		return false
	}

	if dataStoreCompleteUpdateParam.IsSuccess != other.IsSuccess {
		return false
	}

	return true
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

	requestHeaders, err := stream.ReadListStructure(NewDataStoreKeyValue())
	if err != nil {
		return err
	}

	dataStoreReqUpdateInfo.RequestHeaders = requestHeaders.([]*DataStoreKeyValue)

	formFields, err := stream.ReadListStructure(NewDataStoreKeyValue())
	if err != nil {
		return err
	}

	dataStoreReqUpdateInfo.FormFields = formFields.([]*DataStoreKeyValue)

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

// Copy returns a new copied instance of DataStoreReqUpdateInfo
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReqUpdateInfo()

	copied.Version = dataStoreReqUpdateInfo.Version
	copied.Url = dataStoreReqUpdateInfo.Url
	copied.RequestHeaders = make([]*DataStoreKeyValue, len(dataStoreReqUpdateInfo.RequestHeaders))

	for i := 0; i < len(dataStoreReqUpdateInfo.RequestHeaders); i++ {
		copied.RequestHeaders[i] = dataStoreReqUpdateInfo.RequestHeaders[i].Copy().(*DataStoreKeyValue)
	}

	copied.FormFields = make([]*DataStoreKeyValue, len(dataStoreReqUpdateInfo.FormFields))

	for i := 0; i < len(dataStoreReqUpdateInfo.FormFields); i++ {
		copied.FormFields[i] = dataStoreReqUpdateInfo.FormFields[i].Copy().(*DataStoreKeyValue)
	}

	copied.RootCaCert = make([]byte, len(dataStoreReqUpdateInfo.RootCaCert))

	copy(copied.RootCaCert, dataStoreReqUpdateInfo.RootCaCert)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqUpdateInfo *DataStoreReqUpdateInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqUpdateInfo)

	if dataStoreReqUpdateInfo.Version != other.Version {
		return false
	}

	if dataStoreReqUpdateInfo.Url != other.Url {
		return false
	}

	if len(dataStoreReqUpdateInfo.RequestHeaders) != len(other.RequestHeaders) {
		return false
	}

	for i := 0; i < len(dataStoreReqUpdateInfo.RequestHeaders); i++ {
		if dataStoreReqUpdateInfo.RequestHeaders[i] != other.RequestHeaders[i] {
			return false
		}
	}

	if len(dataStoreReqUpdateInfo.FormFields) != len(other.FormFields) {
		return false
	}

	for i := 0; i < len(dataStoreReqUpdateInfo.FormFields); i++ {
		if dataStoreReqUpdateInfo.FormFields[i] != other.FormFields[i] {
			return false
		}
	}

	if !bytes.Equal(dataStoreReqUpdateInfo.RootCaCert, other.RootCaCert) {
		return false
	}

	return true
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
	ExtraData      []string // NEX 3.5.0+
}

// ExtractFromStream extracts a DataStorePrepareUpdateParam structure from a stream
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) ExtractFromStream(stream *nex.StreamIn) error {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	dataStorePrepareUpdateParam.DataID = stream.ReadUInt64LE()
	dataStorePrepareUpdateParam.Size = stream.ReadUInt32LE()
	dataStorePrepareUpdateParam.UpdatePassword = stream.ReadUInt64LE()

	if datastoreVersion.Major >= 3 && datastoreVersion.Minor >= 5 {
		dataStorePrepareUpdateParam.ExtraData = stream.ReadListString()
	}

	return nil
}

// Bytes encodes the DataStorePrepareUpdateParam and returns a byte array
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) Bytes(stream *nex.StreamOut) []byte {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	stream.WriteUInt64LE(dataStorePrepareUpdateParam.DataID)
	stream.WriteUInt32LE(dataStorePrepareUpdateParam.Size)
	stream.WriteUInt64LE(dataStorePrepareUpdateParam.UpdatePassword)

	if datastoreVersion.Major >= 3 && datastoreVersion.Minor >= 5 {
		stream.WriteListString(dataStorePrepareUpdateParam.ExtraData)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePrepareUpdateParam
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) Copy() nex.StructureInterface {
	copied := NewDataStorePrepareUpdateParam()

	copied.DataID = dataStorePrepareUpdateParam.DataID
	copied.Size = dataStorePrepareUpdateParam.Size
	copied.UpdatePassword = dataStorePrepareUpdateParam.UpdatePassword
	copied.ExtraData = make([]string, len(dataStorePrepareUpdateParam.ExtraData))

	copy(copied.ExtraData, dataStorePrepareUpdateParam.ExtraData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePrepareUpdateParam)

	if dataStorePrepareUpdateParam.DataID != other.DataID {
		return false
	}

	if dataStorePrepareUpdateParam.Size != other.Size {
		return false
	}

	if dataStorePrepareUpdateParam.UpdatePassword != other.UpdatePassword {
		return false
	}

	if len(dataStorePrepareUpdateParam.ExtraData) != len(other.ExtraData) {
		return false
	}

	for i := 0; i < len(dataStorePrepareUpdateParam.ExtraData); i++ {
		if dataStorePrepareUpdateParam.ExtraData[i] != other.ExtraData[i] {
			return false
		}
	}

	return true
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

// Copy returns a new copied instance of DataStoreChangeMetaParamV1
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) Copy() nex.StructureInterface {
	copied := NewDataStoreChangeMetaParamV1()

	copied.DataID = dataStoreChangeMetaParamV1.DataID
	copied.ModifiesFlag = dataStoreChangeMetaParamV1.ModifiesFlag
	copied.Name = dataStoreChangeMetaParamV1.Name
	copied.Permission = dataStoreChangeMetaParamV1.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStoreChangeMetaParamV1.DelPermission.Copy().(*DataStorePermission)
	copied.Period = dataStoreChangeMetaParamV1.Period
	copied.MetaBinary = make([]byte, len(dataStoreChangeMetaParamV1.MetaBinary))

	copy(copied.MetaBinary, dataStoreChangeMetaParamV1.MetaBinary)

	copied.Tags = make([]string, len(dataStoreChangeMetaParamV1.Tags))

	copy(copied.Tags, dataStoreChangeMetaParamV1.Tags)

	copied.UpdatePassword = dataStoreChangeMetaParamV1.UpdatePassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreChangeMetaParamV1 *DataStoreChangeMetaParamV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreChangeMetaParamV1)

	if dataStoreChangeMetaParamV1.DataID != other.DataID {
		return false
	}

	if dataStoreChangeMetaParamV1.ModifiesFlag != other.ModifiesFlag {
		return false
	}

	if dataStoreChangeMetaParamV1.Name != other.Name {
		return false
	}

	if !dataStoreChangeMetaParamV1.Permission.Equals(other.Permission) {
		return false
	}

	if !dataStoreChangeMetaParamV1.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if dataStoreChangeMetaParamV1.Period != other.Period {
		return false
	}

	if !bytes.Equal(dataStoreChangeMetaParamV1.MetaBinary, other.MetaBinary) {
		return false
	}

	if len(dataStoreChangeMetaParamV1.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStoreChangeMetaParamV1.Tags); i++ {
		if dataStoreChangeMetaParamV1.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if dataStoreChangeMetaParamV1.UpdatePassword != other.UpdatePassword {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreChangeMetaParamV1
func (dataStoreDeleteParam *DataStoreDeleteParam) Copy() nex.StructureInterface {
	copied := NewDataStoreChangeMetaParamV1()

	copied.DataID = dataStoreDeleteParam.DataID
	copied.UpdatePassword = dataStoreDeleteParam.UpdatePassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreDeleteParam *DataStoreDeleteParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreChangeMetaParamV1)

	if dataStoreDeleteParam.DataID != other.DataID {
		return false
	}

	if dataStoreDeleteParam.UpdatePassword != other.UpdatePassword {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreCompletePostParamV1
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) Copy() nex.StructureInterface {
	copied := NewDataStoreCompletePostParamV1()

	copied.DataID = dataStoreCompletePostParamV1.DataID
	copied.IsSuccess = dataStoreCompletePostParamV1.IsSuccess

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostParamV1 *DataStoreCompletePostParamV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompletePostParamV1)

	if dataStoreCompletePostParamV1.DataID != other.DataID {
		return false
	}

	if dataStoreCompletePostParamV1.IsSuccess != other.IsSuccess {
		return false
	}

	return true
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

	requestHeaders, err := stream.ReadListStructure(NewDataStoreKeyValue())
	if err != nil {
		return err
	}

	dataStoreReqPostInfoV1.RequestHeaders = requestHeaders.([]*DataStoreKeyValue)

	formFields, err := stream.ReadListStructure(NewDataStoreKeyValue())
	if err != nil {
		return err
	}

	dataStoreReqPostInfoV1.FormFields = formFields.([]*DataStoreKeyValue)

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

// Copy returns a new copied instance of DataStoreReqPostInfoV1
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) Copy() nex.StructureInterface {
	copied := NewDataStoreReqPostInfoV1()

	copied.DataID = dataStoreReqPostInfoV1.DataID
	copied.Url = dataStoreReqPostInfoV1.Url
	copied.RequestHeaders = make([]*DataStoreKeyValue, len(dataStoreReqPostInfoV1.RequestHeaders))

	for i := 0; i < len(dataStoreReqPostInfoV1.RequestHeaders); i++ {
		copied.RequestHeaders[i] = dataStoreReqPostInfoV1.RequestHeaders[i].Copy().(*DataStoreKeyValue)
	}

	copied.FormFields = make([]*DataStoreKeyValue, len(dataStoreReqPostInfoV1.FormFields))

	for i := 0; i < len(dataStoreReqPostInfoV1.FormFields); i++ {
		copied.FormFields[i] = dataStoreReqPostInfoV1.FormFields[i].Copy().(*DataStoreKeyValue)
	}

	copied.RootCaCert = make([]byte, len(dataStoreReqPostInfoV1.RootCaCert))

	copy(copied.RootCaCert, dataStoreReqPostInfoV1.RootCaCert)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqPostInfoV1 *DataStoreReqPostInfoV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqPostInfoV1)

	if dataStoreReqPostInfoV1.DataID != other.DataID {
		return false
	}

	if dataStoreReqPostInfoV1.Url != other.Url {
		return false
	}

	if len(dataStoreReqPostInfoV1.RequestHeaders) != len(other.RequestHeaders) {
		return false
	}

	for i := 0; i < len(dataStoreReqPostInfoV1.RequestHeaders); i++ {
		if dataStoreReqPostInfoV1.RequestHeaders[i] != other.RequestHeaders[i] {
			return false
		}
	}

	if len(dataStoreReqPostInfoV1.FormFields) != len(other.FormFields) {
		return false
	}

	for i := 0; i < len(dataStoreReqPostInfoV1.FormFields); i++ {
		if dataStoreReqPostInfoV1.FormFields[i] != other.FormFields[i] {
			return false
		}
	}

	if !bytes.Equal(dataStoreReqPostInfoV1.RootCaCert, other.RootCaCert) {
		return false
	}

	return true
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

	ratingInitParams, err := stream.ReadListStructure(NewDataStoreRatingInitParamWithSlot())
	if err != nil {
		return err
	}

	dataStorePreparePostParamV1.RatingInitParams = ratingInitParams.([]*DataStoreRatingInitParamWithSlot)

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

// Copy returns a new copied instance of DataStorePreparePostParamV1
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) Copy() nex.StructureInterface {
	copied := NewDataStorePreparePostParamV1()

	copied.Size = dataStorePreparePostParamV1.Size
	copied.Name = dataStorePreparePostParamV1.Name
	copied.DataType = dataStorePreparePostParamV1.DataType
	copied.MetaBinary = make([]byte, len(dataStorePreparePostParamV1.MetaBinary))

	copy(copied.MetaBinary, dataStorePreparePostParamV1.MetaBinary)

	copied.Permission = dataStorePreparePostParamV1.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStorePreparePostParamV1.DelPermission.Copy().(*DataStorePermission)
	copied.Flag = dataStorePreparePostParamV1.Flag
	copied.Period = dataStorePreparePostParamV1.Period
	copied.ReferDataID = dataStorePreparePostParamV1.ReferDataID
	copied.Tags = make([]string, len(dataStorePreparePostParamV1.Tags))

	copy(copied.Tags, dataStorePreparePostParamV1.Tags)

	copied.RatingInitParams = make([]*DataStoreRatingInitParamWithSlot, len(dataStorePreparePostParamV1.RatingInitParams))

	for i := 0; i < len(dataStorePreparePostParamV1.RatingInitParams); i++ {
		copied.RatingInitParams[i] = dataStorePreparePostParamV1.RatingInitParams[i].Copy().(*DataStoreRatingInitParamWithSlot)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePreparePostParamV1 *DataStorePreparePostParamV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePreparePostParamV1)

	if dataStorePreparePostParamV1.Size != other.Size {
		return false
	}

	if dataStorePreparePostParamV1.Name != other.Name {
		return false
	}

	if dataStorePreparePostParamV1.DataType != other.DataType {
		return false
	}

	if !bytes.Equal(dataStorePreparePostParamV1.MetaBinary, other.MetaBinary) {
		return false
	}

	if !dataStorePreparePostParamV1.Permission.Equals(other.Permission) {
		return false
	}

	if !dataStorePreparePostParamV1.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if dataStorePreparePostParamV1.Flag != other.Flag {
		return false
	}

	if dataStorePreparePostParamV1.Period != other.Period {
		return false
	}

	if dataStorePreparePostParamV1.ReferDataID != other.ReferDataID {
		return false
	}

	if len(dataStorePreparePostParamV1.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostParamV1.Tags); i++ {
		if dataStorePreparePostParamV1.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if len(dataStorePreparePostParamV1.RatingInitParams) != len(other.RatingInitParams) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostParamV1.RatingInitParams); i++ {
		if dataStorePreparePostParamV1.RatingInitParams[i] != other.RatingInitParams[i] {
			return false
		}
	}

	return true
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

	requestHeaders, err := stream.ReadListStructure(NewDataStoreKeyValue())
	if err != nil {
		return err
	}

	dataStoreReqGetInfoV1.RequestHeaders = requestHeaders.([]*DataStoreKeyValue)
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

// Copy returns a new copied instance of DataStoreReqGetInfoV1
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) Copy() nex.StructureInterface {
	copied := NewDataStoreReqGetInfoV1()

	copied.Url = dataStoreReqGetInfoV1.Url
	copied.RequestHeaders = make([]*DataStoreKeyValue, len(dataStoreReqGetInfoV1.RequestHeaders))

	for i := 0; i < len(dataStoreReqGetInfoV1.RequestHeaders); i++ {
		copied.RequestHeaders[i] = dataStoreReqGetInfoV1.RequestHeaders[i].Copy().(*DataStoreKeyValue)
	}

	copied.Size = dataStoreReqGetInfoV1.Size

	copied.RootCaCert = make([]byte, len(dataStoreReqGetInfoV1.RootCaCert))

	copy(copied.RootCaCert, dataStoreReqGetInfoV1.RootCaCert)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetInfoV1 *DataStoreReqGetInfoV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqGetInfoV1)

	if dataStoreReqGetInfoV1.Url != other.Url {
		return false
	}

	if len(dataStoreReqGetInfoV1.RequestHeaders) != len(other.RequestHeaders) {
		return false
	}

	for i := 0; i < len(dataStoreReqGetInfoV1.RequestHeaders); i++ {
		if dataStoreReqGetInfoV1.RequestHeaders[i] != other.RequestHeaders[i] {
			return false
		}
	}

	if dataStoreReqGetInfoV1.Size != other.Size {
		return false
	}

	if !bytes.Equal(dataStoreReqGetInfoV1.RootCaCert, other.RootCaCert) {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStorePrepareGetParamV1
func (dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1) Copy() nex.StructureInterface {
	copied := NewDataStorePrepareGetParamV1()

	copied.DataID = dataStorePrepareGetParamV1.DataID
	copied.LockID = dataStorePrepareGetParamV1.LockID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePrepareGetParamV1)

	if dataStorePrepareGetParamV1.DataID != other.DataID {
		return false
	}

	if dataStorePrepareGetParamV1.LockID != other.LockID {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreRateObjectParam
func (dataStoreRateObjectParam *DataStoreRateObjectParam) Copy() nex.StructureInterface {
	copied := NewDataStoreRateObjectParam()

	copied.RatingValue = dataStoreRateObjectParam.RatingValue
	copied.AccessPassword = dataStoreRateObjectParam.AccessPassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRateObjectParam *DataStoreRateObjectParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRateObjectParam)

	if dataStoreRateObjectParam.RatingValue != other.RatingValue {
		return false
	}

	if dataStoreRateObjectParam.AccessPassword != other.AccessPassword {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreRatingTarget
func (dataStoreRatingTarget *DataStoreRatingTarget) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingTarget()

	copied.DataID = dataStoreRatingTarget.DataID
	copied.Slot = dataStoreRatingTarget.Slot

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingTarget *DataStoreRatingTarget) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingTarget)

	if dataStoreRatingTarget.DataID != other.DataID {
		return false
	}

	if dataStoreRatingTarget.Slot != other.Slot {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreCompletePostParam
func (dataStoreCompletePostParam *DataStoreCompletePostParam) Copy() nex.StructureInterface {
	copied := NewDataStoreCompletePostParam()

	copied.DataID = dataStoreCompletePostParam.DataID
	copied.IsSuccess = dataStoreCompletePostParam.IsSuccess

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompletePostParam *DataStoreCompletePostParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompletePostParam)

	if dataStoreCompletePostParam.DataID != other.DataID {
		return false
	}

	if dataStoreCompletePostParam.IsSuccess != other.IsSuccess {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreReqPostInfo
func (dataStoreReqPostInfo *DataStoreReqPostInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReqPostInfo()

	copied.DataID = dataStoreReqPostInfo.DataID
	copied.URL = dataStoreReqPostInfo.URL
	copied.RequestHeaders = make([]*DataStoreKeyValue, len(dataStoreReqPostInfo.RequestHeaders))

	for i := 0; i < len(dataStoreReqPostInfo.RequestHeaders); i++ {
		copied.RequestHeaders[i] = dataStoreReqPostInfo.RequestHeaders[i].Copy().(*DataStoreKeyValue)
	}

	copied.FormFields = make([]*DataStoreKeyValue, len(dataStoreReqPostInfo.FormFields))

	for i := 0; i < len(dataStoreReqPostInfo.FormFields); i++ {
		copied.FormFields[i] = dataStoreReqPostInfo.FormFields[i].Copy().(*DataStoreKeyValue)
	}

	copied.RootCACert = make([]byte, len(dataStoreReqPostInfo.RootCACert))

	copy(copied.RootCACert, dataStoreReqPostInfo.RootCACert)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqPostInfo *DataStoreReqPostInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqPostInfo)

	if dataStoreReqPostInfo.DataID != other.DataID {
		return false
	}

	if dataStoreReqPostInfo.URL != other.URL {
		return false
	}

	if len(dataStoreReqPostInfo.RequestHeaders) != len(other.RequestHeaders) {
		return false
	}

	for i := 0; i < len(dataStoreReqPostInfo.RequestHeaders); i++ {
		if dataStoreReqPostInfo.RequestHeaders[i] != other.RequestHeaders[i] {
			return false
		}
	}

	if len(dataStoreReqPostInfo.FormFields) != len(other.FormFields) {
		return false
	}

	for i := 0; i < len(dataStoreReqPostInfo.FormFields); i++ {
		if dataStoreReqPostInfo.FormFields[i] != other.FormFields[i] {
			return false
		}
	}

	if !bytes.Equal(dataStoreReqPostInfo.RootCACert, other.RootCACert) {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStorePersistenceInitParam
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) Copy() nex.StructureInterface {
	copied := NewDataStorePersistenceInitParam()

	copied.PersistenceSlotId = dataStorePersistenceInitParam.PersistenceSlotId
	copied.DeleteLastObject = dataStorePersistenceInitParam.DeleteLastObject

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePersistenceInitParam *DataStorePersistenceInitParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePersistenceInitParam)

	if dataStorePersistenceInitParam.PersistenceSlotId != other.PersistenceSlotId {
		return false
	}

	if dataStorePersistenceInitParam.DeleteLastObject != other.DeleteLastObject {
		return false
	}

	return true
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
	InitialValue   int64
	RangeMin       int32
	RangeMax       int32
	PeriodHour     int8
	PeriodDuration int16
}

// ExtractFromStream extracts a DataStoreRatingInitParam structure from a stream
func (dataStoreRatingInitParam *DataStoreRatingInitParam) ExtractFromStream(stream *nex.StreamIn) error {
	dataStoreRatingInitParam.Flag = stream.ReadUInt8()
	dataStoreRatingInitParam.InternalFlag = stream.ReadUInt8()
	dataStoreRatingInitParam.LockType = stream.ReadUInt8()
	dataStoreRatingInitParam.InitialValue = int64(stream.ReadUInt64LE())
	dataStoreRatingInitParam.RangeMin = int32(stream.ReadUInt32LE())
	dataStoreRatingInitParam.RangeMax = int32(stream.ReadUInt32LE())
	dataStoreRatingInitParam.PeriodHour = int8(stream.ReadUInt8())
	dataStoreRatingInitParam.PeriodDuration = int16(stream.ReadUInt16LE())

	return nil
}

// Copy returns a new copied instance of DataStoreRatingInitParam
func (dataStoreRatingInitParam *DataStoreRatingInitParam) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingInitParam()

	copied.Flag = dataStoreRatingInitParam.Flag
	copied.InternalFlag = dataStoreRatingInitParam.InternalFlag
	copied.LockType = dataStoreRatingInitParam.LockType
	copied.InitialValue = dataStoreRatingInitParam.InitialValue
	copied.RangeMin = dataStoreRatingInitParam.RangeMin
	copied.RangeMax = dataStoreRatingInitParam.RangeMax
	copied.PeriodHour = dataStoreRatingInitParam.PeriodHour
	copied.PeriodDuration = dataStoreRatingInitParam.PeriodDuration

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingInitParam *DataStoreRatingInitParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingInitParam)

	if dataStoreRatingInitParam.Flag != other.Flag {
		return false
	}

	if dataStoreRatingInitParam.InternalFlag != other.InternalFlag {
		return false
	}

	if dataStoreRatingInitParam.LockType != other.LockType {
		return false
	}

	if dataStoreRatingInitParam.InitialValue != other.InitialValue {
		return false
	}

	if dataStoreRatingInitParam.RangeMin != other.RangeMin {
		return false
	}

	if dataStoreRatingInitParam.RangeMax != other.RangeMax {
		return false
	}

	if dataStoreRatingInitParam.PeriodHour != other.PeriodHour {
		return false
	}

	if dataStoreRatingInitParam.PeriodDuration != other.PeriodDuration {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreRatingInitParamWithSlot
func (dataStoreRatingInitParamWithSlot *DataStoreRatingInitParamWithSlot) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingInitParamWithSlot()

	copied.Slot = dataStoreRatingInitParamWithSlot.Slot
	copied.Param = dataStoreRatingInitParamWithSlot.Param.Copy().(*DataStoreRatingInitParam)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingInitParamWithSlot *DataStoreRatingInitParamWithSlot) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingInitParamWithSlot)

	if dataStoreRatingInitParamWithSlot.Slot != other.Slot {
		return false
	}

	if !dataStoreRatingInitParamWithSlot.Param.Equals(other.Param) {
		return false
	}

	return true
}

// NewDataStoreRatingInitParamWithSlot returns a new DataStoreRatingInitParamWithSlot
func NewDataStoreRatingInitParamWithSlot() *DataStoreRatingInitParamWithSlot {
	return &DataStoreRatingInitParamWithSlot{}
}

// DataStorePreparePostParam is sent in the PreparePostObject method
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
	ExtraData            []string // NEX 3.5.0+
}

// ExtractFromStream extracts a DataStorePreparePostParam structure from a stream
func (dataStorePreparePostParam *DataStorePreparePostParam) ExtractFromStream(stream *nex.StreamIn) error {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

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

	ratingInitParams, err := stream.ReadListStructure(NewDataStoreRatingInitParamWithSlot())
	if err != nil {
		return err
	}

	dataStorePreparePostParam.RatingInitParams = ratingInitParams.([]*DataStoreRatingInitParamWithSlot)

	persistenceInitParam, err := stream.ReadStructure(NewDataStorePersistenceInitParam())
	if err != nil {
		return err
	}

	dataStorePreparePostParam.PersistenceInitParam = persistenceInitParam.(*DataStorePersistenceInitParam)

	if datastoreVersion.Major >= 3 && datastoreVersion.Minor >= 5 {
		dataStorePreparePostParam.ExtraData = stream.ReadListString()
	}

	return nil
}

// Copy returns a new copied instance of DataStorePreparePostParam
func (dataStorePreparePostParam *DataStorePreparePostParam) Copy() nex.StructureInterface {
	copied := NewDataStorePreparePostParam()

	copied.Size = dataStorePreparePostParam.Size
	copied.Name = dataStorePreparePostParam.Name
	copied.DataType = dataStorePreparePostParam.DataType
	copied.MetaBinary = make([]byte, len(dataStorePreparePostParam.MetaBinary))

	copy(copied.MetaBinary, dataStorePreparePostParam.MetaBinary)

	copied.Permission = dataStorePreparePostParam.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStorePreparePostParam.DelPermission.Copy().(*DataStorePermission)
	copied.Flag = dataStorePreparePostParam.Flag
	copied.Period = dataStorePreparePostParam.Period
	copied.ReferDataId = dataStorePreparePostParam.ReferDataId
	copied.Tags = make([]string, len(dataStorePreparePostParam.Tags))

	copy(copied.Tags, dataStorePreparePostParam.Tags)

	copied.RatingInitParams = make([]*DataStoreRatingInitParamWithSlot, len(dataStorePreparePostParam.RatingInitParams))

	for i := 0; i < len(dataStorePreparePostParam.RatingInitParams); i++ {
		copied.RatingInitParams[i] = dataStorePreparePostParam.RatingInitParams[i].Copy().(*DataStoreRatingInitParamWithSlot)
	}

	copied.PersistenceInitParam = dataStorePreparePostParam.PersistenceInitParam.Copy().(*DataStorePersistenceInitParam)
	copied.ExtraData = make([]string, len(dataStorePreparePostParam.ExtraData))

	copy(copied.ExtraData, dataStorePreparePostParam.ExtraData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePreparePostParam *DataStorePreparePostParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePreparePostParam)

	if dataStorePreparePostParam.Size != other.Size {
		return false
	}

	if dataStorePreparePostParam.Name != other.Name {
		return false
	}

	if dataStorePreparePostParam.DataType != other.DataType {
		return false
	}

	if !bytes.Equal(dataStorePreparePostParam.MetaBinary, other.MetaBinary) {
		return false
	}

	if !dataStorePreparePostParam.Permission.Equals(other.Permission) {
		return false
	}

	if !dataStorePreparePostParam.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if dataStorePreparePostParam.Flag != other.Flag {
		return false
	}

	if dataStorePreparePostParam.Period != other.Period {
		return false
	}

	if dataStorePreparePostParam.ReferDataId != other.ReferDataId {
		return false
	}

	if len(dataStorePreparePostParam.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostParam.Tags); i++ {
		if dataStorePreparePostParam.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if len(dataStorePreparePostParam.RatingInitParams) != len(other.RatingInitParams) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostParam.RatingInitParams); i++ {
		if !dataStorePreparePostParam.RatingInitParams[i].Equals(other.RatingInitParams[i]) {
			return false
		}
	}

	if !dataStorePreparePostParam.PersistenceInitParam.Equals(other.PersistenceInitParam) {
		return false
	}

	if len(dataStorePreparePostParam.ExtraData) != len(other.ExtraData) {
		return false
	}

	for i := 0; i < len(dataStorePreparePostParam.ExtraData); i++ {
		if dataStorePreparePostParam.ExtraData[i] != other.ExtraData[i] {
			return false
		}
	}

	return true
}

// NewDataStorePreparePostParam returns a new DataStorePreparePostParam
func NewDataStorePreparePostParam() *DataStorePreparePostParam {
	return &DataStorePreparePostParam{}
}

// DataStoreSearchParam is sent in DataStore search methods
type DataStoreSearchParam struct {
	nex.Structure
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
	UseCache               bool // NEX 3.5.0+
}

// ExtractFromStream extracts a DataStoreSearchParam structure from a stream
func (dataStoreSearchParam *DataStoreSearchParam) ExtractFromStream(stream *nex.StreamIn) error {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

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

	if datastoreVersion.Major >= 3 && datastoreVersion.Minor >= 5 {
		dataStoreSearchParam.UseCache = stream.ReadBool()
	}

	return nil
}

// Copy returns a new copied instance of DataStoreSearchParam
func (dataStoreSearchParam *DataStoreSearchParam) Copy() nex.StructureInterface {
	copied := NewDataStoreSearchParam()

	copied.SearchTarget = dataStoreSearchParam.SearchTarget
	copied.OwnerIds = make([]uint32, len(dataStoreSearchParam.OwnerIds))

	copy(copied.OwnerIds, dataStoreSearchParam.OwnerIds)

	copied.OwnerType = dataStoreSearchParam.OwnerType
	copied.DestinationIds = make([]uint64, len(dataStoreSearchParam.DestinationIds))

	copy(copied.DestinationIds, dataStoreSearchParam.DestinationIds)

	copied.DataType = dataStoreSearchParam.DataType
	copied.CreatedAfter = dataStoreSearchParam.CreatedAfter.Copy()
	copied.CreatedBefore = dataStoreSearchParam.CreatedBefore.Copy()
	copied.UpdatedAfter = dataStoreSearchParam.UpdatedAfter.Copy()
	copied.UpdatedBefore = dataStoreSearchParam.UpdatedBefore.Copy()
	copied.ReferDataId = dataStoreSearchParam.ReferDataId
	copied.Tags = make([]string, len(dataStoreSearchParam.Tags))

	copy(copied.Tags, dataStoreSearchParam.Tags)

	copied.ResultOrderColumn = dataStoreSearchParam.ResultOrderColumn
	copied.ResultOrder = dataStoreSearchParam.ResultOrder
	copied.ResultRange = dataStoreSearchParam.ResultRange.Copy().(*nex.ResultRange)
	copied.ResultOption = dataStoreSearchParam.ResultOption
	copied.MinimalRatingFrequency = dataStoreSearchParam.MinimalRatingFrequency
	copied.UseCache = dataStoreSearchParam.UseCache

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreSearchParam *DataStoreSearchParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSearchParam)

	if dataStoreSearchParam.SearchTarget != other.SearchTarget {
		return false
	}

	if len(dataStoreSearchParam.OwnerIds) != len(other.OwnerIds) {
		return false
	}

	for i := 0; i < len(dataStoreSearchParam.OwnerIds); i++ {
		if dataStoreSearchParam.OwnerIds[i] != other.OwnerIds[i] {
			return false
		}
	}

	if dataStoreSearchParam.OwnerType != other.OwnerType {
		return false
	}

	if len(dataStoreSearchParam.DestinationIds) != len(other.DestinationIds) {
		return false
	}

	for i := 0; i < len(dataStoreSearchParam.DestinationIds); i++ {
		if dataStoreSearchParam.DestinationIds[i] != other.DestinationIds[i] {
			return false
		}
	}

	if dataStoreSearchParam.DataType != other.DataType {
		return false
	}

	if !dataStoreSearchParam.CreatedAfter.Equals(other.CreatedAfter) {
		return false
	}

	if !dataStoreSearchParam.CreatedBefore.Equals(other.CreatedBefore) {
		return false
	}

	if !dataStoreSearchParam.UpdatedAfter.Equals(other.UpdatedAfter) {
		return false
	}

	if !dataStoreSearchParam.UpdatedBefore.Equals(other.UpdatedBefore) {
		return false
	}

	if dataStoreSearchParam.ReferDataId != other.ReferDataId {
		return false
	}

	if len(dataStoreSearchParam.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStoreSearchParam.Tags); i++ {
		if dataStoreSearchParam.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if dataStoreSearchParam.ResultOrderColumn != other.ResultOrderColumn {
		return false
	}

	if dataStoreSearchParam.ResultOrder != other.ResultOrder {
		return false
	}

	if !dataStoreSearchParam.ResultRange.Equals(other.ResultRange) {
		return false
	}

	if dataStoreSearchParam.ResultOption != other.ResultOption {
		return false
	}

	if dataStoreSearchParam.MinimalRatingFrequency != other.MinimalRatingFrequency {
		return false
	}

	if dataStoreSearchParam.UseCache != other.UseCache {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreGetMetaParam
func (dataStoreGetMetaParam *DataStoreGetMetaParam) Copy() nex.StructureInterface {
	copied := NewDataStoreGetMetaParam()

	copied.DataID = dataStoreGetMetaParam.DataID
	copied.PersistenceTarget = dataStoreGetMetaParam.PersistenceTarget.Copy().(*DataStorePersistenceTarget)
	copied.ResultOption = dataStoreGetMetaParam.ResultOption
	copied.AccessPassword = dataStoreGetMetaParam.AccessPassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreGetMetaParam *DataStoreGetMetaParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreGetMetaParam)

	if dataStoreGetMetaParam.DataID != other.DataID {
		return false
	}

	if !dataStoreGetMetaParam.PersistenceTarget.Equals(other.PersistenceTarget) {
		return false
	}

	if dataStoreGetMetaParam.ResultOption != other.ResultOption {
		return false
	}

	if dataStoreGetMetaParam.AccessPassword != other.AccessPassword {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreChangeMetaParam
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) Copy() nex.StructureInterface {
	copied := NewDataStoreChangeMetaParam()

	copied.DataID = dataStoreChangeMetaParam.DataID
	copied.ModifiesFlag = dataStoreChangeMetaParam.ModifiesFlag
	copied.Name = dataStoreChangeMetaParam.Name
	copied.Permission = dataStoreChangeMetaParam.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStoreChangeMetaParam.DelPermission.Copy().(*DataStorePermission)
	copied.Period = dataStoreChangeMetaParam.Period
	copied.MetaBinary = make([]byte, len(dataStoreChangeMetaParam.MetaBinary))

	copy(copied.MetaBinary, dataStoreChangeMetaParam.MetaBinary)

	copied.Tags = make([]string, len(dataStoreChangeMetaParam.Tags))

	copy(copied.Tags, dataStoreChangeMetaParam.Tags)

	copied.UpdatePassword = dataStoreChangeMetaParam.UpdatePassword
	copied.ReferredCnt = dataStoreChangeMetaParam.ReferredCnt
	copied.DataType = dataStoreChangeMetaParam.DataType
	copied.Status = dataStoreChangeMetaParam.Status
	copied.CompareParam = dataStoreChangeMetaParam.CompareParam.Copy().(*DataStoreChangeMetaCompareParam)
	// * Uncomment when this is added back
	//copied.PersistenceTarget = dataStoreChangeMetaParam.PersistenceTarget.Copy().(*DataStorePersistenceTarget)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreChangeMetaParam *DataStoreChangeMetaParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreChangeMetaParam)

	if dataStoreChangeMetaParam.DataID != other.DataID {
		return false
	}

	if dataStoreChangeMetaParam.ModifiesFlag != other.ModifiesFlag {
		return false
	}

	if dataStoreChangeMetaParam.Name != other.Name {
		return false
	}

	if dataStoreChangeMetaParam.Permission.Equals(other.Permission) {
		return false
	}

	if dataStoreChangeMetaParam.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if dataStoreChangeMetaParam.Period != other.Period {
		return false
	}

	if !bytes.Equal(dataStoreChangeMetaParam.MetaBinary, other.MetaBinary) {
		return false
	}

	if len(dataStoreChangeMetaParam.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStoreChangeMetaParam.Tags); i++ {
		if dataStoreChangeMetaParam.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if dataStoreChangeMetaParam.UpdatePassword != other.UpdatePassword {
		return false
	}

	if dataStoreChangeMetaParam.ReferredCnt != other.ReferredCnt {
		return false
	}

	if dataStoreChangeMetaParam.DataType != other.DataType {
		return false
	}

	if dataStoreChangeMetaParam.Status != other.Status {
		return false
	}

	if dataStoreChangeMetaParam.CompareParam.Equals(other.CompareParam) {
		return false
	}

	// * Uncomment when this is added back
	/*
		if dataStoreChangeMetaParam.PersistenceTarget.Equals(other.PersistenceTarget) {
			return false
		}
	*/

	return true
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

// Copy returns a new copied instance of DataStoreChangeMetaCompareParam
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) Copy() nex.StructureInterface {
	copied := NewDataStoreChangeMetaCompareParam()

	copied.ComparisonFlag = dataStoreChangeMetaCompareParam.ComparisonFlag
	copied.Name = dataStoreChangeMetaCompareParam.Name
	copied.Permission = dataStoreChangeMetaCompareParam.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStoreChangeMetaCompareParam.DelPermission.Copy().(*DataStorePermission)
	copied.Period = dataStoreChangeMetaCompareParam.Period
	copied.MetaBinary = make([]byte, len(dataStoreChangeMetaCompareParam.MetaBinary))

	copy(copied.MetaBinary, dataStoreChangeMetaCompareParam.MetaBinary)

	copied.Tags = make([]string, len(dataStoreChangeMetaCompareParam.Tags))

	copy(copied.Tags, dataStoreChangeMetaCompareParam.Tags)

	copied.ReferredCnt = dataStoreChangeMetaCompareParam.ReferredCnt
	copied.DataType = dataStoreChangeMetaCompareParam.DataType
	copied.Status = dataStoreChangeMetaCompareParam.Status

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreChangeMetaCompareParam *DataStoreChangeMetaCompareParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreChangeMetaCompareParam)

	if dataStoreChangeMetaCompareParam.ComparisonFlag != other.ComparisonFlag {
		return false
	}

	if dataStoreChangeMetaCompareParam.Name != other.Name {
		return false
	}

	if dataStoreChangeMetaCompareParam.Permission.Equals(other.Permission) {
		return false
	}

	if dataStoreChangeMetaCompareParam.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if dataStoreChangeMetaCompareParam.Period != other.Period {
		return false
	}

	if !bytes.Equal(dataStoreChangeMetaCompareParam.MetaBinary, other.MetaBinary) {
		return false
	}

	if len(dataStoreChangeMetaCompareParam.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStoreChangeMetaCompareParam.Tags); i++ {
		if dataStoreChangeMetaCompareParam.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if dataStoreChangeMetaCompareParam.ReferredCnt != other.ReferredCnt {
		return false
	}

	if dataStoreChangeMetaCompareParam.DataType != other.DataType {
		return false
	}

	if dataStoreChangeMetaCompareParam.Status != other.Status {
		return false
	}

	return true
}

// NewDataStoreChangeMetaCompareParam returns a new DataStoreChangeMetaCompareParam
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

// Copy returns a new copied instance of DataStorePermission
func (dataStorePermission *DataStorePermission) Copy() nex.StructureInterface {
	copied := NewDataStorePermission()

	copied.Permission = dataStorePermission.Permission
	copied.RecipientIds = make([]uint32, len(dataStorePermission.RecipientIds))

	copy(copied.RecipientIds, dataStorePermission.RecipientIds)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePermission *DataStorePermission) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePermission)

	if dataStorePermission.Permission != other.Permission {
		return false
	}

	if len(dataStorePermission.RecipientIds) != len(other.RecipientIds) {
		return false
	}

	for i := 0; i < len(dataStorePermission.RecipientIds); i++ {
		if dataStorePermission.RecipientIds[i] != other.RecipientIds[i] {
			return false
		}
	}

	return true
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

// Copy returns a new copied instance of DataStoreRatingInfo
func (dataStoreRatingInfo *DataStoreRatingInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingInfo()

	copied.TotalValue = dataStoreRatingInfo.TotalValue
	copied.Count = dataStoreRatingInfo.Count
	copied.InitialValue = dataStoreRatingInfo.InitialValue

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingInfo *DataStoreRatingInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingInfo)

	if dataStoreRatingInfo.TotalValue != other.TotalValue {
		return false
	}

	if dataStoreRatingInfo.Count != other.Count {
		return false
	}

	if dataStoreRatingInfo.InitialValue != other.InitialValue {
		return false
	}

	return true
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

// Copy returns a new copied instance of DataStoreRatingInfoWithSlot
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) Copy() nex.StructureInterface {
	copied := NewDataStoreRatingInfoWithSlot()

	copied.Slot = dataStoreRatingInfoWithSlot.Slot
	copied.Rating = dataStoreRatingInfoWithSlot.Rating.Copy().(*DataStoreRatingInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreRatingInfoWithSlot *DataStoreRatingInfoWithSlot) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreRatingInfoWithSlot)

	if dataStoreRatingInfoWithSlot.Slot != other.Slot {
		return false
	}

	if !dataStoreRatingInfoWithSlot.Rating.Equals(other.Rating) {
		return false
	}

	return true
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

	ratings, err := stream.ReadListStructure(NewDataStoreRatingInfoWithSlot())
	if err != nil {
		return err
	}

	dataStoreMetaInfo.Ratings = ratings.([]*DataStoreRatingInfoWithSlot)

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
	stream.WriteDateTime(dataStoreMetaInfo.CreatedTime)
	stream.WriteDateTime(dataStoreMetaInfo.UpdatedTime)
	stream.WriteUInt16LE(dataStoreMetaInfo.Period)
	stream.WriteUInt8(dataStoreMetaInfo.Status)
	stream.WriteUInt32LE(dataStoreMetaInfo.ReferredCnt)
	stream.WriteUInt32LE(dataStoreMetaInfo.ReferDataID)
	stream.WriteUInt32LE(dataStoreMetaInfo.Flag)
	stream.WriteDateTime(dataStoreMetaInfo.ReferredTime)
	stream.WriteDateTime(dataStoreMetaInfo.ExpireTime)
	stream.WriteListString(dataStoreMetaInfo.Tags)
	stream.WriteListStructure(dataStoreMetaInfo.Ratings)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreMetaInfo
func (dataStoreMetaInfo *DataStoreMetaInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreMetaInfo()

	copied.DataID = dataStoreMetaInfo.DataID
	copied.OwnerID = dataStoreMetaInfo.OwnerID
	copied.Size = dataStoreMetaInfo.Size
	copied.DataType = dataStoreMetaInfo.DataType
	copied.Name = dataStoreMetaInfo.Name
	copied.MetaBinary = make([]byte, len(dataStoreMetaInfo.MetaBinary))

	copy(copied.MetaBinary, dataStoreMetaInfo.MetaBinary)

	copied.Permission = dataStoreMetaInfo.Permission.Copy().(*DataStorePermission)
	copied.DelPermission = dataStoreMetaInfo.DelPermission.Copy().(*DataStorePermission)
	copied.CreatedTime = dataStoreMetaInfo.CreatedTime.Copy()
	copied.UpdatedTime = dataStoreMetaInfo.UpdatedTime.Copy()
	copied.Period = dataStoreMetaInfo.Period
	copied.Status = dataStoreMetaInfo.Status
	copied.ReferredCnt = dataStoreMetaInfo.ReferredCnt
	copied.ReferDataID = dataStoreMetaInfo.ReferDataID
	copied.Flag = dataStoreMetaInfo.Flag
	copied.ReferredTime = dataStoreMetaInfo.ReferredTime.Copy()
	copied.ExpireTime = dataStoreMetaInfo.ExpireTime.Copy()
	copied.Tags = make([]string, len(dataStoreMetaInfo.Tags))

	copy(copied.Tags, dataStoreMetaInfo.Tags)

	copied.Ratings = make([]*DataStoreRatingInfoWithSlot, len(dataStoreMetaInfo.Ratings))

	for i := 0; i < len(dataStoreMetaInfo.Ratings); i++ {
		copied.Ratings[i] = dataStoreMetaInfo.Ratings[i].Copy().(*DataStoreRatingInfoWithSlot)
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreMetaInfo *DataStoreMetaInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreMetaInfo)

	if dataStoreMetaInfo.DataID != other.DataID {
		return false
	}

	if dataStoreMetaInfo.OwnerID != other.OwnerID {
		return false
	}

	if dataStoreMetaInfo.Size != other.Size {
		return false
	}

	if dataStoreMetaInfo.DataType != other.DataType {
		return false
	}

	if dataStoreMetaInfo.Name != other.Name {
		return false
	}

	if !bytes.Equal(dataStoreMetaInfo.MetaBinary, other.MetaBinary) {
		return false
	}

	if !dataStoreMetaInfo.Permission.Equals(other.Permission) {
		return false
	}

	if !dataStoreMetaInfo.DelPermission.Equals(other.DelPermission) {
		return false
	}

	if !dataStoreMetaInfo.CreatedTime.Equals(other.CreatedTime) {
		return false
	}

	if !dataStoreMetaInfo.UpdatedTime.Equals(other.UpdatedTime) {
		return false
	}

	if dataStoreMetaInfo.Period != other.Period {
		return false
	}

	if dataStoreMetaInfo.Status != other.Status {
		return false
	}

	if dataStoreMetaInfo.ReferredCnt != other.ReferredCnt {
		return false
	}

	if dataStoreMetaInfo.ReferDataID != other.ReferDataID {
		return false
	}

	if dataStoreMetaInfo.Flag != other.Flag {
		return false
	}

	if !dataStoreMetaInfo.ReferredTime.Equals(other.ReferredTime) {
		return false
	}

	if !dataStoreMetaInfo.ExpireTime.Equals(other.ExpireTime) {
		return false
	}

	if len(dataStoreMetaInfo.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStoreMetaInfo.Tags); i++ {
		if dataStoreMetaInfo.Tags[i] != other.Tags[i] {
			return false
		}
	}

	if len(dataStoreMetaInfo.Tags) != len(other.Tags) {
		return false
	}

	for i := 0; i < len(dataStoreMetaInfo.Ratings); i++ {
		if !dataStoreMetaInfo.Ratings[i].Equals(other.Ratings[i]) {
			return false
		}
	}

	return true
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
	ExtraData         []string // NEX 3.5.0+
}

// ExtractFromStream extracts a DataStorePrepareGetParam structure from a stream
func (dataStorePrepareGetParam *DataStorePrepareGetParam) ExtractFromStream(stream *nex.StreamIn) error {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	dataStorePrepareGetParam.DataID = stream.ReadUInt64LE()
	dataStorePrepareGetParam.LockID = stream.ReadUInt32LE()

	persistenceTarget, err := stream.ReadStructure(NewDataStorePersistenceTarget())
	if err != nil {
		return err
	}

	dataStorePrepareGetParam.PersistenceTarget = persistenceTarget.(*DataStorePersistenceTarget)
	dataStorePrepareGetParam.AccessPassword = stream.ReadUInt64LE()

	if datastoreVersion.Major >= 3 && datastoreVersion.Minor >= 5 {
		dataStorePrepareGetParam.ExtraData = stream.ReadListString()
	}

	return nil
}

// Copy returns a new copied instance of DataStorePrepareGetParam
func (dataStorePrepareGetParam *DataStorePrepareGetParam) Copy() nex.StructureInterface {
	copied := NewDataStorePrepareGetParam()

	copied.DataID = dataStorePrepareGetParam.DataID
	copied.LockID = dataStorePrepareGetParam.LockID
	copied.PersistenceTarget = dataStorePrepareGetParam.PersistenceTarget.Copy().(*DataStorePersistenceTarget)
	copied.AccessPassword = dataStorePrepareGetParam.AccessPassword
	copied.ExtraData = make([]string, len(dataStorePrepareGetParam.ExtraData))

	copy(copied.ExtraData, dataStorePrepareGetParam.ExtraData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePrepareGetParam *DataStorePrepareGetParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePrepareGetParam)

	if dataStorePrepareGetParam.DataID != other.DataID {
		return false
	}

	if dataStorePrepareGetParam.LockID != other.LockID {
		return false
	}

	if !dataStorePrepareGetParam.PersistenceTarget.Equals(other.PersistenceTarget) {
		return false
	}

	if dataStorePrepareGetParam.AccessPassword != other.AccessPassword {
		return false
	}

	if len(dataStorePrepareGetParam.ExtraData) != len(other.ExtraData) {
		return false
	}

	for i := 0; i < len(dataStorePrepareGetParam.ExtraData); i++ {
		if dataStorePrepareGetParam.ExtraData[i] != other.ExtraData[i] {
			return false
		}
	}

	return true
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

// Copy returns a new copied instance of DataStoreKeyValue
func (dataStoreKeyValue *DataStoreKeyValue) Copy() nex.StructureInterface {
	copied := NewDataStoreKeyValue()

	copied.Key = dataStoreKeyValue.Key
	copied.Value = dataStoreKeyValue.Value

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreKeyValue *DataStoreKeyValue) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreKeyValue)

	if dataStoreKeyValue.Key != other.Key {
		return false
	}

	if dataStoreKeyValue.Value != other.Value {
		return false
	}

	return true
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
	DataID         uint64 // NEX 3.5.0+
}

// Bytes encodes the DataStoreReqGetInfo and returns a byte array
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Bytes(stream *nex.StreamOut) []byte {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	stream.WriteString(dataStoreReqGetInfo.URL)
	stream.WriteListStructure(dataStoreReqGetInfo.RequestHeaders)
	stream.WriteUInt32LE(dataStoreReqGetInfo.Size)
	stream.WriteBuffer(dataStoreReqGetInfo.RootCA)

	if datastoreVersion.Major >= 3 && datastoreVersion.Minor >= 5 {
		stream.WriteUInt64LE(dataStoreReqGetInfo.DataID)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreReqGetInfo
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreReqGetInfo()

	copied.URL = dataStoreReqGetInfo.URL
	copied.RequestHeaders = make([]*DataStoreKeyValue, len(dataStoreReqGetInfo.RequestHeaders))

	for i := 0; i < len(dataStoreReqGetInfo.RequestHeaders); i++ {
		copied.RequestHeaders[i] = dataStoreReqGetInfo.RequestHeaders[i].Copy().(*DataStoreKeyValue)
	}

	copied.Size = dataStoreReqGetInfo.Size
	copied.RootCA = make([]byte, len(dataStoreReqGetInfo.RootCA))

	copy(copied.RootCA, dataStoreReqGetInfo.RootCA)

	copied.DataID = dataStoreReqGetInfo.DataID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreReqGetInfo *DataStoreReqGetInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreReqGetInfo)

	if dataStoreReqGetInfo.URL != other.URL {
		return false
	}

	if len(dataStoreReqGetInfo.RequestHeaders) != len(other.RequestHeaders) {
		return false
	}

	for i := 0; i < len(dataStoreReqGetInfo.RequestHeaders); i++ {
		if !dataStoreReqGetInfo.RequestHeaders[i].Equals(other.RequestHeaders[i]) {
			return false
		}
	}

	if dataStoreReqGetInfo.Size != other.Size {
		return false
	}

	if !bytes.Equal(dataStoreReqGetInfo.RootCA, other.RootCA) {
		return false
	}

	if dataStoreReqGetInfo.DataID != other.DataID {
		return false
	}

	return true
}

// NewDataStoreReqGetInfo returns a new DataStoreReqGetInfo
func NewDataStoreReqGetInfo() *DataStoreReqGetInfo {
	return &DataStoreReqGetInfo{}
}
