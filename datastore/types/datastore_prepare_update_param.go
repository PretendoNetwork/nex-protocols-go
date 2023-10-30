// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStorePrepareUpdateParam is a data structure used by the DataStore protocol
type DataStorePrepareUpdateParam struct {
	nex.Structure
	DataID         uint64
	Size           uint32
	UpdatePassword uint64   // NEX 3.0.0+
	ExtraData      []string // NEX 3.5.0+
}

// ExtractFromStream extracts a DataStorePrepareUpdateParam structure from a stream
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) ExtractFromStream(stream *nex.StreamIn) error {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	var err error

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		dataID, err := stream.ReadUInt64LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.DataID. %s", err.Error())
		}

		dataStorePrepareUpdateParam.DataID = dataID
	} else {
		dataID, err := stream.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.DataID. %s", err.Error())
		}

		dataStorePrepareUpdateParam.DataID = uint64(dataID)
	}

	dataStorePrepareUpdateParam.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.Size. %s", err.Error())
	}

	dataStorePrepareUpdateParam.UpdatePassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.UpdatePassword. %s", err.Error())
	}

	if datastoreVersion.GreaterOrEqual("3.5.0") {
		dataStorePrepareUpdateParam.ExtraData, err = stream.ReadListString()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.ExtraData. %s", err.Error())
		}
	}

	return nil
}

// Bytes encodes the DataStorePrepareUpdateParam and returns a byte array
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) Bytes(stream *nex.StreamOut) []byte {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		stream.WriteUInt64LE(dataStorePrepareUpdateParam.DataID)
	} else {
		stream.WriteUInt32LE(uint32(dataStorePrepareUpdateParam.DataID))
	}

	stream.WriteUInt32LE(dataStorePrepareUpdateParam.Size)

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		stream.WriteUInt64LE(dataStorePrepareUpdateParam.UpdatePassword)
	}

	if datastoreVersion.GreaterOrEqual("3.5.0") {
		stream.WriteListString(dataStorePrepareUpdateParam.ExtraData)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePrepareUpdateParam
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) Copy() nex.StructureInterface {
	copied := NewDataStorePrepareUpdateParam()

	copied.SetStructureVersion(dataStorePrepareUpdateParam.StructureVersion())

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

	if dataStorePrepareUpdateParam.StructureVersion() != other.StructureVersion() {
		return false
	}

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

// String returns a string representation of the struct
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) String() string {
	return dataStorePrepareUpdateParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePrepareUpdateParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStorePrepareUpdateParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStorePrepareUpdateParam.DataID))
	b.WriteString(fmt.Sprintf("%sSize: %d,\n", indentationValues, dataStorePrepareUpdateParam.Size))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %d,\n", indentationValues, dataStorePrepareUpdateParam.UpdatePassword))
	b.WriteString(fmt.Sprintf("%sExtraData: %v\n", indentationValues, dataStorePrepareUpdateParam.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePrepareUpdateParam returns a new DataStorePrepareUpdateParam
func NewDataStorePrepareUpdateParam() *DataStorePrepareUpdateParam {
	return &DataStorePrepareUpdateParam{
		DataID:         0,
		Size:           0,
		UpdatePassword: 0,
		ExtraData:      make([]string, 0),
	}
}
