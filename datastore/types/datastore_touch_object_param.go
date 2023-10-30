// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreTouchObjectParam is a data structure used by the DataStore protocol
type DataStoreTouchObjectParam struct {
	nex.Structure
	DataID         uint64
	LockID         uint32
	AccessPassword uint64
}

// ExtractFromStream extracts a DataStoreTouchObjectParam structure from a stream
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	dataStoreTouchObjectParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.DataID. %s", err.Error())
	}

	dataStoreTouchObjectParam.LockID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.LockID. %s", err.Error())
	}

	dataStoreTouchObjectParam.AccessPassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.AccessPassword. %s", err.Error())
	}

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

	copied.SetStructureVersion(dataStoreTouchObjectParam.StructureVersion())

	copied.DataID = dataStoreTouchObjectParam.DataID
	copied.LockID = dataStoreTouchObjectParam.LockID
	copied.AccessPassword = dataStoreTouchObjectParam.AccessPassword

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreTouchObjectParam)

	if dataStoreTouchObjectParam.StructureVersion() != other.StructureVersion() {
		return false
	}

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

// String returns a string representation of the struct
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) String() string {
	return dataStoreTouchObjectParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreTouchObjectParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreTouchObjectParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreTouchObjectParam.DataID))
	b.WriteString(fmt.Sprintf("%sLockID: %d,\n", indentationValues, dataStoreTouchObjectParam.LockID))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %d\n", indentationValues, dataStoreTouchObjectParam.AccessPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreTouchObjectParam returns a new DataStoreTouchObjectParam
func NewDataStoreTouchObjectParam() *DataStoreTouchObjectParam {
	return &DataStoreTouchObjectParam{
		DataID:         0,
		LockID:         0,
		AccessPassword: 0,
	}
}
