// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePrepareGetParamV1 is a data structure used by the DataStore protocol
type DataStorePrepareGetParamV1 struct {
	nex.Structure
	DataID uint32
	LockID uint32
}

// ExtractFromStream extracts a DataStorePrepareGetParamV1 structure from a stream
func (dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStorePrepareGetParamV1.DataID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParamV1.DataID. %s", err.Error())
	}

	dataStorePrepareGetParamV1.LockID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParamV1.LockID. %s", err.Error())
	}

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

	copied.SetStructureVersion(dataStorePrepareGetParamV1.StructureVersion())

	copied.DataID = dataStorePrepareGetParamV1.DataID
	copied.LockID = dataStorePrepareGetParamV1.LockID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStorePrepareGetParamV1)

	if dataStorePrepareGetParamV1.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStorePrepareGetParamV1.DataID != other.DataID {
		return false
	}

	if dataStorePrepareGetParamV1.LockID != other.LockID {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1) String() string {
	return dataStorePrepareGetParamV1.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePrepareGetParamV1{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePrepareGetParamV1.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStorePrepareGetParamV1.DataID))
	b.WriteString(fmt.Sprintf("%sLockID: %d\n", indentationValues, dataStorePrepareGetParamV1.LockID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePrepareGetParamV1 returns a new DataStorePrepareGetParamV1
func NewDataStorePrepareGetParamV1() *DataStorePrepareGetParamV1 {
	return &DataStorePrepareGetParamV1{
		DataID: 0,
		LockID: 0,
	}
}
