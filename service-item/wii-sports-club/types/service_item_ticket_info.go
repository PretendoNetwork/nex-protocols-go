// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemTicketInfo is a type within the ServiceItem protocol
type ServiceItemTicketInfo struct {
	types.Structure
	TicketType types.UInt32
	NumTotal   types.UInt32
}

// WriteTo writes the ServiceItemTicketInfo to the given writable
func (siti ServiceItemTicketInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siti.TicketType.WriteTo(contentWritable)
	siti.NumTotal.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siti.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemTicketInfo from the given readable
func (siti *ServiceItemTicketInfo) ExtractFrom(readable types.Readable) error {
	var err error

	err = siti.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTicketInfo header. %s", err.Error())
	}

	err = siti.TicketType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTicketInfo.TicketType. %s", err.Error())
	}

	err = siti.NumTotal.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTicketInfo.NumTotal. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemTicketInfo
func (siti ServiceItemTicketInfo) Copy() types.RVType {
	copied := NewServiceItemTicketInfo()

	copied.StructureVersion = siti.StructureVersion
	copied.TicketType = siti.TicketType.Copy().(types.UInt32)
	copied.NumTotal = siti.NumTotal.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given ServiceItemTicketInfo contains the same data as the current ServiceItemTicketInfo
func (siti ServiceItemTicketInfo) Equals(o types.RVType) bool {
	if _, ok := o.(ServiceItemTicketInfo); !ok {
		return false
	}

	other := o.(ServiceItemTicketInfo)

	if siti.StructureVersion != other.StructureVersion {
		return false
	}

	if !siti.TicketType.Equals(other.TicketType) {
		return false
	}

	return siti.NumTotal.Equals(other.NumTotal)
}

// CopyRef copies the current value of the ServiceItemTicketInfo
// and returns a pointer to the new copy
func (siti ServiceItemTicketInfo) CopyRef() types.RVTypePtr {
	copied := siti.Copy().(ServiceItemTicketInfo)
	return &copied
}

// Deref takes a pointer to the ServiceItemTicketInfo
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (siti *ServiceItemTicketInfo) Deref() types.RVType {
	return *siti
}

// String returns the string representation of the ServiceItemTicketInfo
func (siti ServiceItemTicketInfo) String() string {
	return siti.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemTicketInfo using the provided indentation level
func (siti ServiceItemTicketInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemTicketInfo{\n")
	b.WriteString(fmt.Sprintf("%sTicketType: %s,\n", indentationValues, siti.TicketType))
	b.WriteString(fmt.Sprintf("%sNumTotal: %s,\n", indentationValues, siti.NumTotal))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemTicketInfo returns a new ServiceItemTicketInfo
func NewServiceItemTicketInfo() ServiceItemTicketInfo {
	return ServiceItemTicketInfo{
		TicketType: types.NewUInt32(0),
		NumTotal:   types.NewUInt32(0),
	}

}
