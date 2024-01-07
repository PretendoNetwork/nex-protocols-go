// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreTouchObjectParam is a data structure used by the DataStore protocol
type DataStoreTouchObjectParam struct {
	types.Structure
	DataID         *types.PrimitiveU64
	LockID         *types.PrimitiveU32
	AccessPassword *types.PrimitiveU64
}

// ExtractFrom extracts the DataStoreTouchObjectParam from the given readable
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreTouchObjectParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreTouchObjectParam header. %s", err.Error())
	}

	err = dataStoreTouchObjectParam.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.DataID. %s", err.Error())
	}

	err = dataStoreTouchObjectParam.LockID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.LockID. %s", err.Error())
	}

	err = dataStoreTouchObjectParam.AccessPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.AccessPassword. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreTouchObjectParam to the given writable
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreTouchObjectParam.DataID.WriteTo(contentWritable)
	dataStoreTouchObjectParam.LockID.WriteTo(contentWritable)
	dataStoreTouchObjectParam.AccessPassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreTouchObjectParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreTouchObjectParam
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) Copy() types.RVType {
	copied := NewDataStoreTouchObjectParam()

	copied.StructureVersion = dataStoreTouchObjectParam.StructureVersion

	copied.DataID = dataStoreTouchObjectParam.DataID.Copy().(*types.PrimitiveU64)
	copied.LockID = dataStoreTouchObjectParam.LockID.Copy().(*types.PrimitiveU32)
	copied.AccessPassword = dataStoreTouchObjectParam.AccessPassword.Copy().(*types.PrimitiveU64)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreTouchObjectParam *DataStoreTouchObjectParam) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreTouchObjectParam); !ok {
		return false
	}

	other := o.(*DataStoreTouchObjectParam)

	if dataStoreTouchObjectParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreTouchObjectParam.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStoreTouchObjectParam.LockID.Equals(other.LockID) {
		return false
	}

	if !dataStoreTouchObjectParam.AccessPassword.Equals(other.AccessPassword) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreTouchObjectParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStoreTouchObjectParam.DataID))
	b.WriteString(fmt.Sprintf("%sLockID: %s,\n", indentationValues, dataStoreTouchObjectParam.LockID))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %s\n", indentationValues, dataStoreTouchObjectParam.AccessPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreTouchObjectParam returns a new DataStoreTouchObjectParam
func NewDataStoreTouchObjectParam() *DataStoreTouchObjectParam {
	return &DataStoreTouchObjectParam{
		DataID:         types.NewPrimitiveU64(0),
		LockID:         types.NewPrimitiveU32(0),
		AccessPassword: types.NewPrimitiveU64(0),
	}
}
