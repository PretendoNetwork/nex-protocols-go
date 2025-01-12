// Package types implements all the types used by the DataStore protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// DataStoreTouchObjectParam is a type within the DataStore protocol
type DataStoreTouchObjectParam struct {
	types.Structure
	DataID         types.UInt64
	LockID         types.UInt32
	AccessPassword types.UInt64
}

// WriteTo writes the DataStoreTouchObjectParam to the given writable
func (dstop DataStoreTouchObjectParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dstop.DataID.WriteTo(contentWritable)
	dstop.LockID.WriteTo(contentWritable)
	dstop.AccessPassword.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dstop.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreTouchObjectParam from the given readable
func (dstop *DataStoreTouchObjectParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dstop.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam header. %s", err.Error())
	}

	err = dstop.DataID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.DataID. %s", err.Error())
	}

	err = dstop.LockID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.LockID. %s", err.Error())
	}

	err = dstop.AccessPassword.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreTouchObjectParam.AccessPassword. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreTouchObjectParam
func (dstop DataStoreTouchObjectParam) Copy() types.RVType {
	copied := NewDataStoreTouchObjectParam()

	copied.StructureVersion = dstop.StructureVersion
	copied.DataID = dstop.DataID.Copy().(types.UInt64)
	copied.LockID = dstop.LockID.Copy().(types.UInt32)
	copied.AccessPassword = dstop.AccessPassword.Copy().(types.UInt64)

	return copied
}

// Equals checks if the given DataStoreTouchObjectParam contains the same data as the current DataStoreTouchObjectParam
func (dstop DataStoreTouchObjectParam) Equals(o types.RVType) bool {
	if _, ok := o.(DataStoreTouchObjectParam); !ok {
		return false
	}

	other := o.(DataStoreTouchObjectParam)

	if dstop.StructureVersion != other.StructureVersion {
		return false
	}

	if !dstop.DataID.Equals(other.DataID) {
		return false
	}

	if !dstop.LockID.Equals(other.LockID) {
		return false
	}

	return dstop.AccessPassword.Equals(other.AccessPassword)
}

// CopyRef copies the current value of the DataStoreTouchObjectParam
// and returns a pointer to the new copy
func (dstop DataStoreTouchObjectParam) CopyRef() types.RVTypePtr {
	copied := dstop.Copy().(DataStoreTouchObjectParam)
	return &copied
}

// Deref takes a pointer to the DataStoreTouchObjectParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dstop *DataStoreTouchObjectParam) Deref() types.RVType {
	return *dstop
}

// String returns the string representation of the DataStoreTouchObjectParam
func (dstop DataStoreTouchObjectParam) String() string {
	return dstop.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreTouchObjectParam using the provided indentation level
func (dstop DataStoreTouchObjectParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreTouchObjectParam{\n")
	b.WriteString(fmt.Sprintf("%sDataID: %s,\n", indentationValues, dstop.DataID))
	b.WriteString(fmt.Sprintf("%sLockID: %s,\n", indentationValues, dstop.LockID))
	b.WriteString(fmt.Sprintf("%sAccessPassword: %s,\n", indentationValues, dstop.AccessPassword))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreTouchObjectParam returns a new DataStoreTouchObjectParam
func NewDataStoreTouchObjectParam() DataStoreTouchObjectParam {
	return DataStoreTouchObjectParam{
		DataID:         types.NewUInt64(0),
		LockID:         types.NewUInt32(0),
		AccessPassword: types.NewUInt64(0),
	}

}
