// Package types implements all the types used by the DataStore protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

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

	copied.SetStructureVersion(dataStoreMetaInfo.StructureVersion())

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

	if dataStoreMetaInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

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

// String returns a string representation of the struct
func (dataStoreMetaInfo *DataStoreMetaInfo) String() string {
	return dataStoreMetaInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreMetaInfo *DataStoreMetaInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreMetaInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreMetaInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreMetaInfo.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %d,\n", indentationValues, dataStoreMetaInfo.OwnerID))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, dataStoreMetaInfo.Size))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStoreMetaInfo.DataType))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, dataStoreMetaInfo.Name))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %x,\n", indentationValues, dataStoreMetaInfo.MetaBinary))

	if dataStoreMetaInfo.Permission != nil {
		b.WriteString(fmt.Sprintf("%sPermission: %s,\n", indentationValues, dataStoreMetaInfo.Permission.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPermission: nil,\n", indentationValues))
	}

	if dataStoreMetaInfo.DelPermission != nil {
		b.WriteString(fmt.Sprintf("%sDelPermission: %s,\n", indentationValues, dataStoreMetaInfo.DelPermission.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDelPermission: nil,\n", indentationValues))
	}

	if dataStoreMetaInfo.CreatedTime != nil {
		b.WriteString(fmt.Sprintf("%sCreatedTime: %s,\n", indentationValues, dataStoreMetaInfo.CreatedTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCreatedTime: nil,\n", indentationValues))
	}

	if dataStoreMetaInfo.UpdatedTime != nil {
		b.WriteString(fmt.Sprintf("%sUpdatedTime: %s,\n", indentationValues, dataStoreMetaInfo.UpdatedTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUpdatedTime: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sPeriod: %d,\n", indentationValues, dataStoreMetaInfo.Period))
	b.WriteString(fmt.Sprintf("%sStatus: %d,\n", indentationValues, dataStoreMetaInfo.Status))
	b.WriteString(fmt.Sprintf("%sReferredCnt: %d,\n", indentationValues, dataStoreMetaInfo.ReferredCnt))
	b.WriteString(fmt.Sprintf("%sReferDataID: %d,\n", indentationValues, dataStoreMetaInfo.ReferDataID))
	b.WriteString(fmt.Sprintf("%sFlag: %d,\n", indentationValues, dataStoreMetaInfo.Flag))

	if dataStoreMetaInfo.ReferredTime != nil {
		b.WriteString(fmt.Sprintf("%sReferredTime: %s,\n", indentationValues, dataStoreMetaInfo.ReferredTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sReferredTime: nil,\n", indentationValues))
	}

	if dataStoreMetaInfo.ExpireTime != nil {
		b.WriteString(fmt.Sprintf("%sExpireTime: %s,\n", indentationValues, dataStoreMetaInfo.ExpireTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sExpireTime: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sTags: %v,\n", indentationValues, dataStoreMetaInfo.Tags))

	if len(dataStoreMetaInfo.Ratings) == 0 {
		b.WriteString(fmt.Sprintf("%sRatings: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sRatings: [\n", indentationValues))

		for i := 0; i < len(dataStoreMetaInfo.Ratings); i++ {
			str := dataStoreMetaInfo.Ratings[i].FormatToString(indentationLevel + 2)
			if i == len(dataStoreMetaInfo.Ratings)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s]\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// FilterPropertiesByResultOption zeroes out certain struct properties based on the input flags
func (dataStoreMetaInfo *DataStoreMetaInfo) FilterPropertiesByResultOption(resultOption uint8) {
	// * This is kind of backwards
	// *
	// * This method assumes all struct data exists
	// * by default. This is done in order to simplify
	// * database calls by just querying for all fields
	// * at once. Therefore, instead of the ResultOption
	// * flags being used to conditionally ADD properties,
	// * it's used to conditionally REMOVE them

	if resultOption&0x1 == 0 {
		dataStoreMetaInfo.Tags = make([]string, 0)
	}

	if resultOption&0x2 == 0 {
		dataStoreMetaInfo.Ratings = make([]*DataStoreRatingInfoWithSlot, 0)
	}

	if resultOption&0x4 == 0 {
		dataStoreMetaInfo.MetaBinary = make([]byte, 0)
	}
}

// NewDataStoreMetaInfo returns a new DataStoreMetaInfo
func NewDataStoreMetaInfo() *DataStoreMetaInfo {
	return &DataStoreMetaInfo{
		DataID:        0,
		OwnerID:       0,
		Size:          0,
		DataType:      0,
		Name:          "",
		MetaBinary:    make([]byte, 0),
		Permission:    NewDataStorePermission(),
		DelPermission: NewDataStorePermission(),
		CreatedTime:   nex.NewDateTime(0),
		UpdatedTime:   nex.NewDateTime(0),
		Period:        0,
		Status:        0,
		ReferredCnt:   0,
		ReferDataID:   0,
		Flag:          0,
		ReferredTime:  nex.NewDateTime(0),
		ExpireTime:    nex.NewDateTime(0),
		Tags:          make([]string, 0),
		Ratings:       make([]*DataStoreRatingInfoWithSlot, 0),
	}
}
