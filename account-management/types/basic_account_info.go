// Package types implements all the types used by the Account Management protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/types"
)

// BasicAccountInfo contains data for creating a new NNID on the network
type BasicAccountInfo struct {
	types.Structure
	PIDOwner *types.PID
	StrName  *types.String
}

// WriteTo writes the BasicAccountInfo to the given writable
func (basicAccountInfo *BasicAccountInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	basicAccountInfo.PIDOwner.WriteTo(contentWritable)
	basicAccountInfo.StrName.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	basicAccountInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the BasicAccountInfo from the given readable
func (basicAccountInfo *BasicAccountInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = basicAccountInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read BasicAccountInfo header. %s", err.Error())
	}

	err = basicAccountInfo.PIDOwner.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BasicAccountInfo.PIDOwner. %s", err.Error())
	}

	err = basicAccountInfo.StrName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract BasicAccountInfo.StrName. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of BasicAccountInfo
func (basicAccountInfo *BasicAccountInfo) Copy() types.RVType {
	copied := NewBasicAccountInfo()

	copied.StructureVersion = basicAccountInfo.StructureVersion

	copied.PIDOwner = basicAccountInfo.PIDOwner.Copy().(*types.PID)
	copied.StrName = basicAccountInfo.StrName.Copy().(*types.String)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (basicAccountInfo *BasicAccountInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*BasicAccountInfo); !ok {
		return false
	}

	other := o.(*BasicAccountInfo)

	if basicAccountInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !basicAccountInfo.PIDOwner.Equals(other.PIDOwner) {
		return false
	}

	if !basicAccountInfo.StrName.Equals(other.StrName) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (basicAccountInfo *BasicAccountInfo) String() string {
	return basicAccountInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (basicAccountInfo *BasicAccountInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("BasicAccountInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, basicAccountInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPIDOwner: %s,\n", indentationValues, basicAccountInfo.PIDOwner.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sStrName: %s\n", indentationValues, basicAccountInfo.StrName))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewBasicAccountInfo returns a new BasicAccountInfo
func NewBasicAccountInfo() *BasicAccountInfo {
	return &BasicAccountInfo{
		PIDOwner: types.NewPID(0),
		StrName: types.NewString(""),
	}
}
