// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePrepareUpdateParam is a data structure used by the DataStore protocol
type DataStorePrepareUpdateParam struct {
	types.Structure
	DataID         *types.PrimitiveU64
	Size           *types.PrimitiveU32
	UpdatePassword *types.PrimitiveU64        // NEX 3.0.0+
	ExtraData      *types.List[*types.String] // NEX 3.5.0+
}

// ExtractFrom extracts the DataStorePrepareUpdateParam from the given readable
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	var err error

	if err = dataStorePrepareUpdateParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePrepareUpdateParam header. %s", err.Error())
	}

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		dataID, err := readable.ReadPrimitiveUInt64LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.DataID. %s", err.Error())
		}

		dataStorePrepareUpdateParam.DataID.Value = dataID
	} else {
		dataID, err := readable.ReadPrimitiveUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.DataID. %s", err.Error())
		}

		dataStorePrepareUpdateParam.DataID.Value = *types.PrimitiveU64(dataID)
	}

	err = dataStorePrepareUpdateParam.Size.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.Size. %s", err.Error())
	}

	err = dataStorePrepareUpdateParam.UpdatePassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.UpdatePassword. %s", err.Error())
	}

	if datastoreVersion.GreaterOrEqual("3.5.0") {
		err = dataStorePrepareUpdateParam.ExtraData.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePrepareUpdateParam.ExtraData. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the DataStorePrepareUpdateParam to the given writable
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	contentWritable := writable.CopyNew()

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		contentWritable.WritePrimitiveUInt64LE(dataStorePrepareUpdateParam.DataID.Value)
	} else {
		contentWritable.WritePrimitiveUInt32LE(*types.PrimitiveU32(dataStorePrepareUpdateParam.DataID.Value))
	}

	dataStorePrepareUpdateParam.Size.WriteTo(contentWritable)

	if datastoreVersion.GreaterOrEqual("3.0.0") {
		dataStorePrepareUpdateParam.UpdatePassword.WriteTo(contentWritable)
	}

	if datastoreVersion.GreaterOrEqual("3.5.0") {
		dataStorePrepareUpdateParam.ExtraData.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	dataStorePrepareUpdateParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStorePrepareUpdateParam
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) Copy() types.RVType {
	copied := NewDataStorePrepareUpdateParam()

	copied.StructureVersion = dataStorePrepareUpdateParam.StructureVersion

	copied.DataID = dataStorePrepareUpdateParam.DataID.Copy().(*types.PrimitiveU64)
	copied.Size = dataStorePrepareUpdateParam.Size.Copy().(*types.PrimitiveU32)
	copied.UpdatePassword = dataStorePrepareUpdateParam.UpdatePassword.Copy().(*types.PrimitiveU64)
	copied.ExtraData = dataStorePrepareUpdateParam.ExtraData.Copy().(*types.List[*types.String])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePrepareUpdateParam *DataStorePrepareUpdateParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePrepareUpdateParam); !ok {
		return false
	}

	other := o.(*DataStorePrepareUpdateParam)

	if dataStorePrepareUpdateParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePrepareUpdateParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStorePrepareUpdateParam.Size.Equals(other.Size) {
		return false
	}

	if !dataStorePrepareUpdateParam.UpdatePassword.Equals(other.UpdatePassword) {
		return false
	}

	if !dataStorePrepareUpdateParam.ExtraData.Equals(other.ExtraData) {
		return false
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePrepareUpdateParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStorePrepareUpdateParam.DataID))
	b.WriteString(fmt.Sprintf("%sSize: %s,\n", indentationValues, dataStorePrepareUpdateParam.Size))
	b.WriteString(fmt.Sprintf("%sUpdatePassword: %s,\n", indentationValues, dataStorePrepareUpdateParam.UpdatePassword))
	b.WriteString(fmt.Sprintf("%sExtraData: %s\n", indentationValues, dataStorePrepareUpdateParam.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePrepareUpdateParam returns a new DataStorePrepareUpdateParam
func NewDataStorePrepareUpdateParam() *DataStorePrepareUpdateParam {
	dataStorePrepareUpdateParam := &DataStorePrepareUpdateParam{
		DataID:         types.NewPrimitiveU64(0),
		Size:           types.NewPrimitiveU32(0),
		UpdatePassword: types.NewPrimitiveU64(0),
		ExtraData:      types.NewList[*types.String](),
	}

	dataStorePrepareUpdateParam.ExtraData.Type = types.NewString("")

	return dataStorePrepareUpdateParam
}
