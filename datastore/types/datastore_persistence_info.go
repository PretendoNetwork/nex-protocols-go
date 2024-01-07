// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePersistenceInfo is a data structure used by the DataStore protocol
type DataStorePersistenceInfo struct {
	types.Structure
	OwnerID           *types.PID
	PersistenceSlotID *types.PrimitiveU16
	DataID            *types.PrimitiveU64
}

// ExtractFrom extracts the DataStorePersistenceInfo from the given readable
func (dataStorePersistenceInfo *DataStorePersistenceInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStorePersistenceInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePersistenceInfo header. %s", err.Error())
	}

	err = dataStorePersistenceInfo.OwnerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo.OwnerID. %s", err.Error())
	}

	err = dataStorePersistenceInfo.PersistenceSlotID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo.PersistenceSlotID. %s", err.Error())
	}

	err = dataStorePersistenceInfo.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo.DataID. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStorePersistenceInfo to the given writable
func (dataStorePersistenceInfo *DataStorePersistenceInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStorePersistenceInfo.OwnerID.WriteTo(contentWritable)
	dataStorePersistenceInfo.PersistenceSlotID.WriteTo(contentWritable)
	dataStorePersistenceInfo.DataID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStorePersistenceInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStorePersistenceInfo
func (dataStorePersistenceInfo *DataStorePersistenceInfo) Copy() types.RVType {
	copied := NewDataStorePersistenceInfo()

	copied.StructureVersion = dataStorePersistenceInfo.StructureVersion

	copied.OwnerID = dataStorePersistenceInfo.OwnerID.Copy().(*types.PID)
	copied.PersistenceSlotID = dataStorePersistenceInfo.PersistenceSlotID.Copy().(*types.PrimitiveU16)
	copied.DataID = dataStorePersistenceInfo.DataID.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePersistenceInfo *DataStorePersistenceInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePersistenceInfo); !ok {
		return false
	}

	other := o.(*DataStorePersistenceInfo)

	if dataStorePersistenceInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePersistenceInfo.OwnerID.Equals(other.OwnerID) {
		return false
	}

	if !dataStorePersistenceInfo.PersistenceSlotID.Equals(other.PersistenceSlotID) {
		return false
	}

	if !dataStorePersistenceInfo.DataID.Equals(other.DataID) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePersistenceInfo *DataStorePersistenceInfo) String() string {
	return dataStorePersistenceInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePersistenceInfo *DataStorePersistenceInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePersistenceInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePersistenceInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sOwnerID: %s,\n", indentationValues, dataStorePersistenceInfo.OwnerID))
	b.WriteString(fmt.Sprintf("%sPersistenceSlotID: %s,\n", indentationValues, dataStorePersistenceInfo.PersistenceSlotID))
	b.WriteString(fmt.Sprintf("%sDataID: %s\n", indentationValues, dataStorePersistenceInfo.DataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePersistenceInfo returns a new DataStorePersistenceInfo
func NewDataStorePersistenceInfo() *DataStorePersistenceInfo {
	return &DataStorePersistenceInfo{
		OwnerID:           types.NewPID(0),
		PersistenceSlotID: types.NewPrimitiveU16(0),
		DataID:            types.NewPrimitiveU64(0),
	}
}
