package datastore_types

import (
	"bytes"
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

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
	var err error

	dataStoreMetaInfo.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.DataID. %s", err.Error())
	}

	dataStoreMetaInfo.OwnerID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.OwnerID. %s", err.Error())
	}

	dataStoreMetaInfo.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Size. %s", err.Error())
	}

	dataStoreMetaInfo.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Name. %s", err.Error())
	}

	dataStoreMetaInfo.DataType, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.DataType. %s", err.Error())
	}

	dataStoreMetaInfo.MetaBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.MetaBinary. %s", err.Error())
	}

	permission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Permission. %s", err.Error())
	}

	dataStoreMetaInfo.Permission = permission.(*DataStorePermission)
	delPermission, err := stream.ReadStructure(NewDataStorePermission())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.DelPermission. %s", err.Error())
	}

	dataStoreMetaInfo.DelPermission = delPermission.(*DataStorePermission)
	dataStoreMetaInfo.CreatedTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.CreatedTime. %s", err.Error())
	}

	dataStoreMetaInfo.UpdatedTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.UpdatedTime. %s", err.Error())
	}

	dataStoreMetaInfo.Period, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Period. %s", err.Error())
	}

	dataStoreMetaInfo.Status, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Status. %s", err.Error())
	}

	dataStoreMetaInfo.ReferredCnt, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.ReferredCnt. %s", err.Error())
	}

	dataStoreMetaInfo.ReferDataID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.ReferDataID. %s", err.Error())
	}

	dataStoreMetaInfo.Flag, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Flag. %s", err.Error())
	}

	dataStoreMetaInfo.ReferredTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.ReferredTime. %s", err.Error())
	}

	dataStoreMetaInfo.ExpireTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.ExpireTime. %s", err.Error())
	}

	dataStoreMetaInfo.Tags, err = stream.ReadListString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Tags. %s", err.Error())
	}

	ratings, err := stream.ReadListStructure(NewDataStoreRatingInfoWithSlot())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreMetaInfo.Ratings. %s", err.Error())
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
