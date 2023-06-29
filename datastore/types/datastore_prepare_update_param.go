package datastore_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

type DataStorePrepareUpdateParam struct {
	nex.Structure
	DataID         uint64
	Size           uint32
	UpdatePassword uint64
	ExtraData      []string // NEX 3.5.0+
}

// ExtractFromStream extracts a DataStorePrepareUpdateParam structure from a stream
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) ExtractFromStream(stream *nex.StreamIn) error {
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	var err error

	dataStorePrepareUpdateParam.DataID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.DataID. %s", err.Error())
	}

	dataStorePrepareUpdateParam.Size, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.Size. %s", err.Error())
	}

	dataStorePrepareUpdateParam.UpdatePassword, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.UpdatePassword. %s", err.Error())
	}

	if datastoreVersion.Major >= 3 && datastoreVersion.Minor >= 5 {
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

	stream.WriteUInt64LE(dataStorePrepareUpdateParam.DataID)
	stream.WriteUInt32LE(dataStorePrepareUpdateParam.Size)
	stream.WriteUInt64LE(dataStorePrepareUpdateParam.UpdatePassword)

	if datastoreVersion.Major >= 3 && datastoreVersion.Minor >= 5 {
		stream.WriteListString(dataStorePrepareUpdateParam.ExtraData)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of DataStorePrepareUpdateParam
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) Copy() nex.StructureInterface {
	copied := NewDataStorePrepareUpdateParam()

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
	return &DataStorePrepareUpdateParam{}
}
