// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreSpecificMetaInfo is a data structure used by the DataStore protocol
type DataStoreSpecificMetaInfo struct {
	nex.Structure
	DataID   uint64
	OwnerID  *nex.PID
	Size     uint32
	DataType uint16
	Version  uint32
}

// ExtractFromStream extracts a DataStoreSpecificMetaInfo structure from a stream
func (d *DataStoreSpecificMetaInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	d.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.DataID. %s", err.Error())
	}

	d.OwnerID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.OwnerID. %s", err.Error())
	}

	d.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.Size. %s", err.Error())
	}

	d.DataType, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.DataType. %s", err.Error())
	}

	d.Version, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreSpecificMetaInfo.Version. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreSpecificMetaInfo and returns a byte array
func (d *DataStoreSpecificMetaInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(d.DataID)
	stream.WritePID(d.OwnerID)
	stream.WriteUInt32LE(d.Size)
	stream.WriteUInt16LE(d.DataType)
	stream.WriteUInt32LE(d.Version)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreSpecificMetaInfo
func (d *DataStoreSpecificMetaInfo) Copy() nex.StructureInterface {
	copied := NewDataStoreSpecificMetaInfo()

	copied.SetStructureVersion(d.StructureVersion())

	copied.DataID = d.DataID
	copied.OwnerID = d.OwnerID.Copy()
	copied.Size = d.Size
	copied.DataType = d.DataType
	copied.Version = d.Version

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (d *DataStoreSpecificMetaInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreSpecificMetaInfo)

	if d.StructureVersion() != other.StructureVersion() {
		return false
	}

	if d.DataID != other.DataID {
		return false
	}

	if !d.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if d.Size != other.Size {
		return false
	}

	if d.DataType != other.DataType {
		return false
	}

	if d.Version != other.Version {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (d *DataStoreSpecificMetaInfo) String() string {
	return d.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (d *DataStoreSpecificMetaInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreSpecificMetaInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, d.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, d.DataID))
	b.WriteString(fmt.Sprintf("%sOwnerID: %d,\n", indentationValues, d.OwnerID))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, d.Size))
	b.WriteString(fmt.Sprintf("%sDataType: %d,\n", indentationValues, d.DataType))
	b.WriteString(fmt.Sprintf("%sVersion: %d\n", indentationValues, d.Version))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreSpecificMetaInfo returns a new DataStoreSpecificMetaInfo
func NewDataStoreSpecificMetaInfo() *DataStoreSpecificMetaInfo {
	return &DataStoreSpecificMetaInfo{
		DataID:   0,
		OwnerID:  nex.NewPID[uint32](0),
		Size:     0,
		DataType: 0,
		Version:  0,
	}
}
