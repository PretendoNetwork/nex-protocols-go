// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePrepareGetParam is sent in the PrepareGetObject method
type DataStorePrepareGetParam struct {
	types.Structure
	DataID            *types.PrimitiveU64
	LockID            *types.PrimitiveU32
	PersistenceTarget *DataStorePersistenceTarget
	AccessPassword    *types.PrimitiveU64
	ExtraData         *types.List[*types.String] // NEX 3.5.0+
}

// WriteTo writes the DataStorePersistenceInitParam to the given writable
func (dataStorePrepareGetParam *DataStorePrepareGetParam) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	contentWritable := writable.CopyNew()

	dataStorePrepareGetParam.DataID.WriteTo(contentWritable)
	dataStorePrepareGetParam.LockID.WriteTo(contentWritable)
	dataStorePrepareGetParam.PersistenceTarget.WriteTo(contentWritable)
	dataStorePrepareGetParam.AccessPassword.WriteTo(contentWritable)

	if datastoreVersion.GreaterOrEqual("3.5.0") {
		dataStorePrepareGetParam.ExtraData.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	dataStorePrepareGetParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStorePrepareGetParam from the given readable
func (dataStorePrepareGetParam *DataStorePrepareGetParam) ExtractFrom(readable types.Readable) error {
	stream := readable.(*nex.ByteStreamIn)
	datastoreVersion := stream.Server.DataStoreProtocolVersion()

	var err error

	if err = dataStorePrepareGetParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePrepareGetParam header. %s", err.Error())
	}

	err = dataStorePrepareGetParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam.DataID. %s", err.Error())
	}

	err = dataStorePrepareGetParam.LockID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam.LockID. %s", err.Error())
	}

	err = dataStorePrepareGetParam.PersistenceTarget.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam.PersistenceTarget. %s", err.Error())
	}

	err = dataStorePrepareGetParam.AccessPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParam.AccessPassword. %s", err.Error())
	}

	if datastoreVersion.GreaterOrEqual("3.5.0") {
		err = dataStorePrepareGetParam.ExtraData.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract DataStorePrepareGetParam.ExtraData. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of DataStorePrepareGetParam
func (dataStorePrepareGetParam *DataStorePrepareGetParam) Copy() types.RVType {
	copied := NewDataStorePrepareGetParam()

	copied.StructureVersion = dataStorePrepareGetParam.StructureVersion

	copied.DataID = dataStorePrepareGetParam.DataID.Copy().(*types.PrimitiveU64)
	copied.LockID = dataStorePrepareGetParam.LockID.Copy().(*types.PrimitiveU32)
	copied.PersistenceTarget = dataStorePrepareGetParam.PersistenceTarget.Copy().(*DataStorePersistenceTarget)
	copied.AccessPassword = dataStorePrepareGetParam.AccessPassword.Copy().(*types.PrimitiveU64)
	copied.ExtraData = dataStorePrepareGetParam.ExtraData.Copy().(*types.List[*types.String])

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePrepareGetParam *DataStorePrepareGetParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePrepareGetParam); !ok {
		return false
	}

	other := o.(*DataStorePrepareGetParam)

	if dataStorePrepareGetParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePrepareGetParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStorePrepareGetParam.LockID.Equals(other.LockID) {
		return false
	}

	if !dataStorePrepareGetParam.PersistenceTarget.Equals(other.PersistenceTarget) {
		return false
	}

	if !dataStorePrepareGetParam.AccessPassword.Equals(other.AccessPassword) {
		return false
	}

	if !dataStorePrepareGetParam.ExtraData.Equals(other.ExtraData) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStorePrepareGetParam *DataStorePrepareGetParam) String() string {
	return dataStorePrepareGetParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStorePrepareGetParam *DataStorePrepareGetParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStorePrepareGetParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePrepareGetParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStorePrepareGetParam.DataID))
	b.WriteString(fmt.Sprintf("%sLockID: %s,\n", indentationValues, dataStorePrepareGetParam.LockID))
	b.WriteString(fmt.Sprintf("%sPersistenceTarget: %s,\n", indentationValues, dataStorePrepareGetParam.PersistenceTarget.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %s,\n", indentationValues, dataStorePrepareGetParam.AccessPassword))
	b.WriteString(fmt.Sprintf("%sExtraData: %s\n", indentationValues, dataStorePrepareGetParam.ExtraData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePrepareGetParam returns a new DataStorePrepareGetParam
func NewDataStorePrepareGetParam() *DataStorePrepareGetParam {
	dataStorePrepareGetParam := &DataStorePrepareGetParam{
		DataID:            types.NewPrimitiveU64(0),
		LockID:            types.NewPrimitiveU32(0),
		PersistenceTarget: NewDataStorePersistenceTarget(),
		AccessPassword:    types.NewPrimitiveU64(0),
		ExtraData:         types.NewList[*types.String](),
	}

	dataStorePrepareGetParam.ExtraData.Type = types.NewString("")

	return dataStorePrepareGetParam
}
