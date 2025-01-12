// Package types implements all the types used by the DataStoreSuperSmashBros.4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
)

// DataStoreCompletePostReplayParam is a type within the DataStoreSuperSmashBros.4 protocol
type DataStoreCompletePostReplayParam struct {
	types.Structure
	ReplayID      types.UInt64
	CompleteParam datastore_types.DataStoreCompletePostParam
	PrepareParam  DataStorePreparePostReplayParam
}

// WriteTo writes the DataStoreCompletePostReplayParam to the given writable
func (dscprp DataStoreCompletePostReplayParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dscprp.ReplayID.WriteTo(contentWritable)
	dscprp.CompleteParam.WriteTo(contentWritable)
	dscprp.PrepareParam.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dscprp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the DataStoreCompletePostReplayParam from the given readable
func (dscprp *DataStoreCompletePostReplayParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = dscprp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostReplayParam header. %s", err.Error())
	}

	err = dscprp.ReplayID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostReplayParam.ReplayID. %s", err.Error())
	}

	err = dscprp.CompleteParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostReplayParam.CompleteParam. %s", err.Error())
	}

	err = dscprp.PrepareParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreCompletePostReplayParam.PrepareParam. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of DataStoreCompletePostReplayParam
func (dscprp DataStoreCompletePostReplayParam) Copy() types.RVType {
	copied := NewDataStoreCompletePostReplayParam()

	copied.StructureVersion = dscprp.StructureVersion
	copied.ReplayID = dscprp.ReplayID.Copy().(types.UInt64)
	copied.CompleteParam = dscprp.CompleteParam.Copy().(datastore_types.DataStoreCompletePostParam)
	copied.PrepareParam = dscprp.PrepareParam.Copy().(DataStorePreparePostReplayParam)

	return copied
}

// Equals checks if the given DataStoreCompletePostReplayParam contains the same data as the current DataStoreCompletePostReplayParam
func (dscprp DataStoreCompletePostReplayParam) Equals(o types.RVType) bool {
	if _, ok := o.(DataStoreCompletePostReplayParam); !ok {
		return false
	}

	other := o.(DataStoreCompletePostReplayParam)

	if dscprp.StructureVersion != other.StructureVersion {
		return false
	}

	if !dscprp.ReplayID.Equals(other.ReplayID) {
		return false
	}

	if !dscprp.CompleteParam.Equals(other.CompleteParam) {
		return false
	}

	return dscprp.PrepareParam.Equals(other.PrepareParam)
}

// CopyRef copies the current value of the DataStoreCompletePostReplayParam
// and returns a pointer to the new copy
func (dscprp DataStoreCompletePostReplayParam) CopyRef() types.RVTypePtr {
	copied := dscprp.Copy().(DataStoreCompletePostReplayParam)
	return &copied
}

// Deref takes a pointer to the DataStoreCompletePostReplayParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (dscprp *DataStoreCompletePostReplayParam) Deref() types.RVType {
	return *dscprp
}

// String returns the string representation of the DataStoreCompletePostReplayParam
func (dscprp DataStoreCompletePostReplayParam) String() string {
	return dscprp.FormatToString(0)
}

// FormatToString pretty-prints the DataStoreCompletePostReplayParam using the provided indentation level
func (dscprp DataStoreCompletePostReplayParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreCompletePostReplayParam{\n")
	b.WriteString(fmt.Sprintf("%sReplayID: %s,\n", indentationValues, dscprp.ReplayID))
	b.WriteString(fmt.Sprintf("%sCompleteParam: %s,\n", indentationValues, dscprp.CompleteParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrepareParam: %s,\n", indentationValues, dscprp.PrepareParam.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreCompletePostReplayParam returns a new DataStoreCompletePostReplayParam
func NewDataStoreCompletePostReplayParam() DataStoreCompletePostReplayParam {
	return DataStoreCompletePostReplayParam{
		ReplayID:      types.NewUInt64(0),
		CompleteParam: datastore_types.NewDataStoreCompletePostParam(),
		PrepareParam:  NewDataStorePreparePostReplayParam(),
	}

}
