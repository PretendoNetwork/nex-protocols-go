// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreCompleteUpdateParam is a data structure used by the DataStore protocol
type DataStoreCompleteUpdateParam struct {
	types.Structure
	DataID    *types.PrimitiveU64
	Version   *types.PrimitiveU32
	IsSuccess *types.PrimitiveBool
}

// ExtractFrom extracts the DataStoreCompleteUpdateParam from the given readable
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	var err error

	if err = dataStoreCompleteUpdateParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreCompleteUpdateParam header. %s", err.Error())
	}

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		dataID, err := readable.ReadPrimitiveUInt64LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.DataID. %s", err.Error())
		}

		dataStoreCompleteUpdateParam.DataID.Value = dataID
	} else {
		dataID, err := readable.ReadPrimitiveUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.DataID. %s", err.Error())
		}

		dataStoreCompleteUpdateParam.DataID.Value = *types.PrimitiveU64(dataID)
	}

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		version, err := readable.ReadPrimitiveUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.Version. %s", err.Error())
		}

		dataStoreCompleteUpdateParam.Version.Value = version
	} else {
		version, err := readable.ReadPrimitiveUInt16LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStoreReqUpdateInfo.Version. %s", err.Error())
		}

		dataStoreCompleteUpdateParam.Version.Value = *types.PrimitiveU32(version)
	}

	err = dataStoreCompleteUpdateParam.IsSuccess.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompleteUpdateParam.IsSuccess. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreCompleteUpdateParam to the given writable
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	contentWritable := writable.CopyNew()

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		contentWritable.WritePrimitiveUInt64LE(dataStoreCompleteUpdateParam.DataID.Value)
	} else {
		contentWritable.WritePrimitiveUInt32LE(*types.PrimitiveU32(dataStoreCompleteUpdateParam.DataID.Value))
	}

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		contentWritable.WritePrimitiveUInt32LE(dataStoreCompleteUpdateParam.Version.Value)
	} else {
		contentWritable.WritePrimitiveUInt16LE(*types.PrimitiveU16(dataStoreCompleteUpdateParam.Version.Value))
	}

	dataStoreCompleteUpdateParam.IsSuccess.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreCompleteUpdateParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreCompleteUpdateParam
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) Copy() types.RVType {
	copied := NewDataStoreCompleteUpdateParam()

	copied.StructureVersion = dataStoreCompleteUpdateParam.StructureVersion

	copied.DataID = dataStoreCompleteUpdateParam.DataID.Copy().(*types.PrimitiveU64)
	copied.Version = dataStoreCompleteUpdateParam.Version.Copy().(*types.PrimitiveU32)
	copied.IsSuccess = dataStoreCompleteUpdateParam.IsSuccess.Copy().(*types.PrimitiveBool)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreCompleteUpdateParam *DataStoreCompleteUpdateParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreCompleteUpdateParam); !ok {
		return false
	}

	other := o.(*DataStoreCompleteUpdateParam)

	if dataStoreCompleteUpdateParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreCompleteUpdateParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreCompleteUpdateParam.Version.Equals(other.Version) {
		return false
	}

	if !dataStoreCompleteUpdateParam.IsSuccess.Equals(other.IsSuccess) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreCompleteUpdateParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreCompleteUpdateParam.DataID))
	b.WriteString(fmt.Sprintf("%sVersion: %s,\n", indentationValues, dataStoreCompleteUpdateParam.Version))
	b.WriteString(fmt.Sprintf("%sIsSuccess: %s\n", indentationValues, dataStoreCompleteUpdateParam.IsSuccess))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompleteUpdateParam returns a new DataStoreCompleteUpdateParam
func NewDataStoreCompleteUpdateParam() *DataStoreCompleteUpdateParam {
	return &DataStoreCompleteUpdateParam{
		DataID:    types.NewPrimitiveU64(0),
		Version:   types.NewPrimitiveU32(0),
		IsSuccess: types.NewPrimitiveBool(false),
	}
}
