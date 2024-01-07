// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemTicketInfo holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemTicketInfo struct {
	types.Structure
	TicketType *types.PrimitiveU32
	NumTotal   *types.PrimitiveU32
}

// ExtractFrom extracts the ServiceItemTicketInfo from the given readable
func (serviceItemTicketInfo *ServiceItemTicketInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemTicketInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemTicketInfo header. %s", err.Error())
	}

	err = serviceItemTicketInfo.TicketType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTicketInfo.TicketType from stream. %s", err.Error())
	}

	err = serviceItemTicketInfo.NumTotal.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemTicketInfo.NumTotal from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemTicketInfo to the given writable
func (serviceItemTicketInfo *ServiceItemTicketInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemTicketInfo.TicketType.WriteTo(contentWritable)
	serviceItemTicketInfo.NumTotal.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemTicketInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemTicketInfo
func (serviceItemTicketInfo *ServiceItemTicketInfo) Copy() types.RVType {
	copied := NewServiceItemTicketInfo()

	copied.StructureVersion = serviceItemTicketInfo.StructureVersion

	copied.TicketType = serviceItemTicketInfo.TicketType
	copied.NumTotal = serviceItemTicketInfo.NumTotal

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemTicketInfo *ServiceItemTicketInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemTicketInfo); !ok {
		return false
	}

	other := o.(*ServiceItemTicketInfo)

	if serviceItemTicketInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemTicketInfo.TicketType.Equals(other.TicketType) {
		return false
	}

	if !serviceItemTicketInfo.NumTotal.Equals(other.NumTotal) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemTicketInfo *ServiceItemTicketInfo) String() string {
	return serviceItemTicketInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemTicketInfo *ServiceItemTicketInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemTicketInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemTicketInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sTicketType: %d,\n", indentationValues, serviceItemTicketInfo.TicketType))
	b.WriteString(fmt.Sprintf("%sNumTotal: %d,\n", indentationValues, serviceItemTicketInfo.NumTotal))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemTicketInfo returns a new ServiceItemTicketInfo
func NewServiceItemTicketInfo() *ServiceItemTicketInfo {
	return &ServiceItemTicketInfo{}
}
