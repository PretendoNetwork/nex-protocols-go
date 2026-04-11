// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SubscriberUserStatusInfo is a type within the Shop protocol
type SubscriberUserStatusInfo struct {
	types.Structure
	PID    types.PID
	Values types.List[types.QBuffer]
}

// WriteTo writes the SubscriberUserStatusInfo to the given writable
func (susi SubscriberUserStatusInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	susi.PID.WriteTo(contentWritable)
	susi.Values.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	susi.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SubscriberUserStatusInfo from the given readable
func (susi *SubscriberUserStatusInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = susi.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberUserStatusInfo header. %s", err.Error())
	}

	err = susi.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberUserStatusInfo.PID. %s", err.Error())
	}

	err = susi.Values.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberUserStatusInfo.Values. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SubscriberUserStatusInfo
func (susi SubscriberUserStatusInfo) Copy() types.RVType {
	copied := NewSubscriberUserStatusInfo()

	copied.StructureVersion = susi.StructureVersion
	copied.PID = susi.PID.Copy().(types.PID)
	copied.Values = susi.Values.Copy().(types.List[types.QBuffer])

	return copied
}

// Equals checks if the given SubscriberUserStatusInfo contains the same data as the current SubscriberUserStatusInfo
func (susi SubscriberUserStatusInfo) Equals(o types.RVType) bool {
	if _, ok := o.(SubscriberUserStatusInfo); !ok {
		return false
	}

	other := o.(SubscriberUserStatusInfo)

	if susi.StructureVersion != other.StructureVersion {
		return false
	}

	if !susi.PID.Equals(other.PID) {
		return false
	}

	return susi.Values.Equals(other.Values)
}

// CopyRef copies the current value of the SubscriberUserStatusInfo
// and returns a pointer to the new copy
func (susi SubscriberUserStatusInfo) CopyRef() types.RVTypePtr {
	copied := susi.Copy().(SubscriberUserStatusInfo)
	return &copied
}

// Deref takes a pointer to the SubscriberUserStatusInfo
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (susi *SubscriberUserStatusInfo) Deref() types.RVType {
	return *susi
}

// String returns the string representation of the SubscriberUserStatusInfo
func (susi SubscriberUserStatusInfo) String() string {
	return susi.FormatToString(0)
}

// FormatToString pretty-prints the SubscriberUserStatusInfo using the provided indentation level
func (susi SubscriberUserStatusInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberUserStatusInfo{\n")
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, susi.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sValues: %s,\n", indentationValues, susi.Values))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberUserStatusInfo returns a new SubscriberUserStatusInfo
func NewSubscriberUserStatusInfo() SubscriberUserStatusInfo {
	return SubscriberUserStatusInfo{
		PID:    types.NewPID(0),
		Values: types.NewList[types.QBuffer](),
	}

}
