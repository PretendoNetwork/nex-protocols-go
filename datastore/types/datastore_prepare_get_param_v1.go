// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStorePrepareGetParamV1 is a data structure used by the DataStore protocol
type DataStorePrepareGetParamV1 struct {
	types.Structure
	DataID *types.PrimitiveU32
	LockID *types.PrimitiveU32
}

// ExtractFrom extracts the DataStorePrepareGetParamV1 from the given readable
func (dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStorePrepareGetParamV1.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStorePrepareGetParamV1 header. %s", err.Error())
	}

	err = dataStorePrepareGetParamV1.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParamV1.DataID. %s", err.Error())
	}

	err = dataStorePrepareGetParamV1.LockID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStorePrepareGetParamV1.LockID. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStorePrepareGetParamV1 to the given writable
func (dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStorePrepareGetParamV1.DataID.WriteTo(contentWritable)
	dataStorePrepareGetParamV1.LockID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStorePrepareGetParamV1.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStorePrepareGetParamV1
func (dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1) Copy() types.RVType {
	copied := NewDataStorePrepareGetParamV1()

	copied.StructureVersion = dataStorePrepareGetParamV1.StructureVersion

	copied.DataID = dataStorePrepareGetParamV1.DataID.Copy().(*types.PrimitiveU32)
	copied.LockID = dataStorePrepareGetParamV1.LockID.Copy().(*types.PrimitiveU32)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStorePrepareGetParamV1 *DataStorePrepareGetParamV1) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStorePrepareGetParamV1); !ok {
		return false
	}

	other := o.(*DataStorePrepareGetParamV1)

	if dataStorePrepareGetParamV1.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStorePrepareGetParamV1.DataID.Equals(other.DataID) {
		return false
	}

	if !dataStorePrepareGetParamV1.LockID.Equals(other.LockID) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStorePrepareGetParamV1.StructureVersion))
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dataStorePrepareGetParamV1.DataID))
	b.WriteString(fmt.Sprintf("%sLockID: %s\n", indentationValues, dataStorePrepareGetParamV1.LockID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStorePrepareGetParamV1 returns a new DataStorePrepareGetParamV1
func NewDataStorePrepareGetParamV1() *DataStorePrepareGetParamV1 {
	return &DataStorePrepareGetParamV1{
		DataID: types.NewPrimitiveU32(0),
		LockID: types.NewPrimitiveU32(0),
	}
}
