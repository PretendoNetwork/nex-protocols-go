package datastore_super_smash_bros_4_types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreSharedDataInfo is a data structure used by the DataStore Super Smash Bros. 4 protocol
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
	var err error

	dataStoreSharedDataInfo.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.DataID. %s", err.Error())
	}

	dataStoreSharedDataInfo.OwnerID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.OwnerID. %s", err.Error())
	}

	dataStoreSharedDataInfo.DataType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.DataType. %s", err.Error())
	}

	dataStoreSharedDataInfo.Comment, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.Comment. %s", err.Error())
	}

	dataStoreSharedDataInfo.MetaBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.MetaBinary. %s", err.Error())
	}

	dataStoreSharedDataInfo.Profile, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.MetaBinary. %s", err.Error())
	}

	dataStoreSharedDataInfo.Rating, err = stream.ReadInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.Rating. %s", err.Error())
	}

	dataStoreSharedDataInfo.CreatedTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.CreatedTime. %s", err.Error())
	}

	info, err := stream.ReadStructure(NewDataStoreFileServerObjectInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.Info. %s", err.Error())
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

// String returns a string representation of the struct
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) String() string {
	return dataStoreSharedDataInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSharedDataInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreSharedDataInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreSharedDataInfo.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %d,\n", indentationValues, dataStoreSharedDataInfo.OwnerID))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, dataStoreSharedDataInfo.DataType))
	b.WriteString(fmt.Sprintf("%sComment: %q,\n", indentationValues, dataStoreSharedDataInfo.Comment))
	b.WriteString(fmt.Sprintf("%sMetaBinary: %x,\n", indentationValues, dataStoreSharedDataInfo.MetaBinary))
	b.WriteString(fmt.Sprintf("%sProfile: %x,\n", indentationValues, dataStoreSharedDataInfo.Profile))
	b.WriteString(fmt.Sprintf("%sRating: %d,\n", indentationValues, dataStoreSharedDataInfo.Rating))

	if dataStoreSharedDataInfo.CreatedTime != nil {
		b.WriteString(fmt.Sprintf("%sCreatedTime: %s\n", indentationValues, dataStoreSharedDataInfo.CreatedTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCreatedTime: nil\n", indentationValues))
	}

	if dataStoreSharedDataInfo.Info != nil {
		b.WriteString(fmt.Sprintf("%sInfo: %s\n", indentationValues, dataStoreSharedDataInfo.Info.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sInfo: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSharedDataInfo returns a new DataStoreSharedDataInfo
func NewDataStoreSharedDataInfo() *DataStoreSharedDataInfo {
	return &DataStoreSharedDataInfo{}
}
