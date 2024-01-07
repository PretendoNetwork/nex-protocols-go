// Package types implements all the types used by the DataStore Super Smash Bros. 4 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// DataStoreProfileInfo is a data structure used by the DataStore Super Smash Bros. 4 protocol
type DataStoreProfileInfo struct {
	types.Structure
	PID     *types.PID
	Profile *types.QBuffer
}

// ExtractFrom extracts the DataStoreProfileInfo from the given readable
func (dataStoreProfileInfo *DataStoreProfileInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = dataStoreProfileInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read DataStoreProfileInfo header. %s", err.Error())
	}

	err = dataStoreProfileInfo.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreProfileInfo.PID. %s", err.Error())
	}

	err = dataStoreProfileInfo.Profile.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract DataStoreProfileInfo.Profile. %s", err.Error())
	}

	return nil
}

// WriteTo writes the DataStoreProfileInfo to the given writable
func (dataStoreProfileInfo *DataStoreProfileInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	dataStoreProfileInfo.PID.WriteTo(contentWritable)
	dataStoreProfileInfo.Profile.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	dataStoreProfileInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of DataStoreProfileInfo
func (dataStoreProfileInfo *DataStoreProfileInfo) Copy() types.RVType {
	copied := NewDataStoreProfileInfo()

	copied.StructureVersion = dataStoreProfileInfo.StructureVersion

	copied.PID = dataStoreProfileInfo.PID.Copy().(*types.PID)
	copied.Profile = dataStoreProfileInfo.Profile.Copy().(*types.QBuffer)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (dataStoreProfileInfo *DataStoreProfileInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*DataStoreProfileInfo); !ok {
		return false
	}

	other := o.(*DataStoreProfileInfo)

	if dataStoreProfileInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !dataStoreProfileInfo.PID.Equals(other.PID) {
		return false
	}

	if !dataStoreProfileInfo.Profile.Equals(other.Profile) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (dataStoreProfileInfo *DataStoreProfileInfo) String() string {
	return dataStoreProfileInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (dataStoreProfileInfo *DataStoreProfileInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("DataStoreProfileInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, dataStoreProfileInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, dataStoreProfileInfo.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sProfile: %s\n", indentationValues, dataStoreProfileInfo.Profile))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewDataStoreProfileInfo returns a new DataStoreProfileInfo
func NewDataStoreProfileInfo() *DataStoreProfileInfo {
	return &DataStoreProfileInfo{
		PID: types.NewPID(0),
		Profile: types.NewQBuffer(nil),
	}
}
