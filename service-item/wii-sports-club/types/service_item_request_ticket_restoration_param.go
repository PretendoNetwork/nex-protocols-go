// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemRequestTicketRestorationParam is a type within the ServiceItem protocol
type ServiceItemRequestTicketRestorationParam struct {
	types.Structure
	TicketType types.UInt32
	NumTicket  types.UInt32
}

// WriteTo writes the ServiceItemRequestTicketRestorationParam to the given writable
func (sirtrp ServiceItemRequestTicketRestorationParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sirtrp.TicketType.WriteTo(contentWritable)
	sirtrp.NumTicket.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sirtrp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemRequestTicketRestorationParam from the given readable
func (sirtrp *ServiceItemRequestTicketRestorationParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = sirtrp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRequestTicketRestorationParam header. %s", err.Error())
	}

	err = sirtrp.TicketType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRequestTicketRestorationParam.TicketType. %s", err.Error())
	}

	err = sirtrp.NumTicket.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemRequestTicketRestorationParam.NumTicket. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemRequestTicketRestorationParam
func (sirtrp ServiceItemRequestTicketRestorationParam) Copy() types.RVType {
	copied := NewServiceItemRequestTicketRestorationParam()

	copied.StructureVersion = sirtrp.StructureVersion
	copied.TicketType = sirtrp.TicketType.Copy().(types.UInt32)
	copied.NumTicket = sirtrp.NumTicket.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given ServiceItemRequestTicketRestorationParam contains the same data as the current ServiceItemRequestTicketRestorationParam
func (sirtrp ServiceItemRequestTicketRestorationParam) Equals(o types.RVType) bool {
	if _, ok := o.(ServiceItemRequestTicketRestorationParam); !ok {
		return false
	}

	other := o.(ServiceItemRequestTicketRestorationParam)

	if sirtrp.StructureVersion != other.StructureVersion {
		return false
	}

	if !sirtrp.TicketType.Equals(other.TicketType) {
		return false
	}

	return sirtrp.NumTicket.Equals(other.NumTicket)
}

// CopyRef copies the current value of the ServiceItemRequestTicketRestorationParam
// and returns a pointer to the new copy
func (sirtrp ServiceItemRequestTicketRestorationParam) CopyRef() types.RVTypePtr {
	copied := sirtrp.Copy().(ServiceItemRequestTicketRestorationParam)
	return &copied
}

// Deref takes a pointer to the ServiceItemRequestTicketRestorationParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sirtrp *ServiceItemRequestTicketRestorationParam) Deref() types.RVType {
	return *sirtrp
}

// String returns the string representation of the ServiceItemRequestTicketRestorationParam
func (sirtrp ServiceItemRequestTicketRestorationParam) String() string {
	return sirtrp.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemRequestTicketRestorationParam using the provided indentation level
func (sirtrp ServiceItemRequestTicketRestorationParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemRequestTicketRestorationParam{\n")
	b.WriteString(fmt.Sprintf("%sTicketType: %s,\n", indentationValues, sirtrp.TicketType))
	b.WriteString(fmt.Sprintf("%sNumTicket: %s,\n", indentationValues, sirtrp.NumTicket))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemRequestTicketRestorationParam returns a new ServiceItemRequestTicketRestorationParam
func NewServiceItemRequestTicketRestorationParam() ServiceItemRequestTicketRestorationParam {
	return ServiceItemRequestTicketRestorationParam{
		TicketType: types.NewUInt32(0),
		NumTicket:  types.NewUInt32(0),
	}

}
