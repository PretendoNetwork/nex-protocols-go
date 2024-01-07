// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreSharedDataInfo is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreSharedDataInfo struct {
	types.Structure
	DataID      *types.PrimitiveU64
	OwnerID     *types.PrimitiveU32
	DataType    *types.PrimitiveU8
	Comment     string
	MetaBinary  []byte
	Profile     []byte
	Rating      *types.PrimitiveS64
	CreatedTime *types.DateTime
	Info        *DataStoreFileServerObjectInfo
}

// ExtractFrom extracts the DataStoreSharedDataInfo from the given readable
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreSharedDataInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreSharedDataInfo header. %s", err.Error())
	}

	err = dataStoreSharedDataInfo.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.DataID. %s", err.Error())
	}

	err = dataStoreSharedDataInfo.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.OwnerID. %s", err.Error())
	}

	err = dataStoreSharedDataInfo.DataType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.DataType. %s", err.Error())
	}

	err = dataStoreSharedDataInfo.Comment.ExtractFrom(readable)
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

	err = dataStoreSharedDataInfo.Rating.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.Rating. %s", err.Error())
	}

	err = dataStoreSharedDataInfo.CreatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.CreatedTime. %s", err.Error())
	}

	err = dataStoreSharedDataInfo.Info.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSharedDataInfo.Info. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreSharedDataInfo to the given writable
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreSharedDataInfo.DataID.WriteTo(contentWritable)
	dataStoreSharedDataInfo.OwnerID.WriteTo(contentWritable)
	dataStoreSharedDataInfo.DataType.WriteTo(contentWritable)
	dataStoreSharedDataInfo.Comment.WriteTo(contentWritable)
	stream.WriteQBuffer(dataStoreSharedDataInfo.MetaBinary)
	stream.WriteQBuffer(dataStoreSharedDataInfo.Profile)
	dataStoreSharedDataInfo.Rating.WriteTo(contentWritable)
	dataStoreSharedDataInfo.CreatedTime.WriteTo(contentWritable)
	dataStoreSharedDataInfo.Info.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreSharedDataInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreSharedDataInfo
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) Copy() types.RVType {
	copied := NewDataStoreSharedDataInfo()

	copied.StructureVersion = dataStoreSharedDataInfo.StructureVersion

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
func (dataStoreSharedDataInfo *DataStoreSharedDataInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreSharedDataInfo); !ok {
		return false
	}

	other := o.(*DataStoreSharedDataInfo)

	if dataStoreSharedDataInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreSharedDataInfo.DataType.Equals(other.DataType) {
		return false
	}

	if !dataStoreSharedDataInfo.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreSharedDataInfo.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !dataStoreSharedDataInfo.DataType.Equals(other.DataType) {
		return false
	}

	if !dataStoreSharedDataInfo.Comment.Equals(other.Comment) {
		return false
	}

	if !dataStoreSharedDataInfo.MetaBinary.Equals(other.MetaBinary) {
		return false
	}

	if !dataStoreSharedDataInfo.Profile.Equals(other.Profile) {
		return false
	}

	if !dataStoreSharedDataInfo.Rating.Equals(other.Rating) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreSharedDataInfo.StructureVersion))
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
