// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// DataStoreCompleteUpdateParam is a data structure used by the DataStore protocol
type DataStoreCompleteUpdateParam struct {
	nex.Structure
	DataID    uint64
	Version   uint32
	IsSuccess bool
}

// ExtractFromStream extracts a DataStoreCompleteUpdateParam structure from a stream
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) ExtractFromStream(stream *nex.StreamIn) error {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	var err error

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		dataID, err := stream.ReadUInt64LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.DataID. %s", err.Error())
		}

		dataStoreCompleteUpdateParam.DataID = dataID
	} else {
		dataID, err := stream.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.DataID. %s", err.Error())
		}

		dataStoreCompleteUpdateParam.DataID = uint64(dataID)
	}

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		version, err := stream.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.Version. %s", err.Error())
		}

		dataStoreCompleteUpdateParam.Version = version
	} else {
		version, err := stream.ReadUInt16LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.Version. %s", err.Error())
		}

		dataStoreCompleteUpdateParam.Version = uint32(version)
	}

	dataStoreCompleteUpdateParam.IsSuccess, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.IsSuccess. %s", err.Error())
	}

	return nil
}

// Bytes encodes the DataStoreCompleteUpdateParam and returns a byte array
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) Bytes(stream *nex.StreamOut) []byte {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		stream.WriteUInt64LE(dataStoreCompleteUpdateParam.DataID)
	} else {
		stream.WriteUInt32LE(uint32(dataStoreCompleteUpdateParam.DataID))
	}

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		stream.WriteUInt32LE(dataStoreCompleteUpdateParam.Version)
	} else {
		stream.WriteUInt16LE(uint16(dataStoreCompleteUpdateParam.Version))
	}

	stream.WriteBool(dataStoreCompleteUpdateParam.IsSuccess)

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStoreCompleteUpdateParam
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) Copy() nex.StructureInterface {
	copied := NewDataStoreCompleteUpdateParam()

	copied.SetStructureVersion(dataStoreCompleteUpdateParam.StructureVersion())

	copied.DataID = dataStoreCompleteUpdateParam.DataID
	copied.Version = dataStoreCompleteUpdateParam.Version
	copied.IsSuccess = dataStoreCompleteUpdateParam.IsSuccess

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*DataStoreCompleteUpdateParam)

	if dataStoreCompleteUpdateParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if dataStoreCompleteUpdateParam.DataID != other.DataID {
		return false
	}

	if dataStoreCompleteUpdateParam.Version != other.Version {
		return false
	}

	if dataStoreCompleteUpdateParam.IsSuccess != other.IsSuccess {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) String() string {
	return dataStoreCompleteUpdateParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompleteUpdateParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, dataStoreCompleteUpdateParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sDataID: %d,\n", indentationValues, dataStoreCompleteUpdateParam.DataID))
	b.WriteString(fmt.Sprintf("%sVersion: %d,\n", indentationValues, dataStoreCompleteUpdateParam.Version))
	b.WriteString(fmt.Sprintf("%sIsSuccess: %t\n", indentationValues, dataStoreCompleteUpdateParam.IsSuccess))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompleteUpdateParam returns a new DataStoreCompleteUpdateParam
func NewDataStoreCompleteUpdateParam() *DataStoreCompleteUpdateParam {
	return &DataStoreCompleteUpdateParam{}
}
