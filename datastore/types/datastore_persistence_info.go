// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePersistenceInfo is a data structure used by the DataStore protocol
type DataStorePersistenceInfo struct {
	nex.Structure
	OwnerID           *nex.PID
	PersistenceSlotID uint16
	DataID            uint64
}

// ExtractFromStream extracts a DataStorePersistenceInfo structure from a stream
func (dataStorePersistenceInfo *DataStorePersistenceInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePersistenceInfo.OwnerID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo.OwnerID. %s", err.Error())
	}

	dataStorePersistenceInfo.PersistenceSlotID, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo.PersistenceSlotID. %s", err.Error())
	}

	dataStorePersistenceInfo.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePersistenceInfo.DataID. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStorePersistenceInfo and returns a byte array
func (dataStorePersistenceInfo *DataStorePersistenceInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WritePID(dataStorePersistenceInfo.OwnerID)
	stream.WriteUInt16LE(dataStorePersistenceInfo.PersistenceSlotID)
	stream.WriteUInt64LE(dataStorePersistenceInfo.DataID)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePersistenceInfo
func (dataStorePersistenceInfo *DataStorePersistenceInfo) Copy() nex.StructureInterface {
	copied := NewDataStorePersistenceInfo()

	copied.SetStructureVersion(dataStorePersistenceInfo.StructureVersion())

	copied.OwnerID = dataStorePersistenceInfo.OwnerID.Copy()
	copied.PersistenceSlotID = dataStorePersistenceInfo.PersistenceSlotID
	copied.DataID = dataStorePersistenceInfo.DataID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePersistenceInfo *DataStorePersistenceInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePersistenceInfo)

	if dataStorePersistenceInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !dataStorePersistenceInfo.OwnerID.Equals(other.OwnerID) {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePersistenceInfo.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sOwnerID: %d,\n", indentationValues, dataStorePersistenceInfo.OwnerID))
	b.WriteString(fmt.Sprintf("%sPersistenceSlotID: %d,\n", indentationValues, dataStorePersistenceInfo.PersistenceSlotID))
	b.WriteString(fmt.Sprintf("%sDataID: %d\n", indentationValues, dataStorePersistenceInfo.DataID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePersistenceInfo returns a new DataStorePersistenceInfo
func NewDataStorePersistenceInfo() *DataStorePersistenceInfo {
	return &DataStorePersistenceInfo{
		OwnerID:           nex.NewPID[uint32](0),
		PersistenceSlotID: 0,
		DataID:            0,
	}
}
