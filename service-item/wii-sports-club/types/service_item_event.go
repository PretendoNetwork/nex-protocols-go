// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemEvent holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemEvent struct {
	types.Structure
	EventID           *types.PrimitiveU64
	ParamInt          *types.PrimitiveS32
	ParamString       string
	ParamBinary       []byte
	PresentTicketType *types.PrimitiveU32
	PresentTicketNum  *types.PrimitiveU32
	TimeBegin         *types.DateTime
	TimeEnd           *types.DateTime
}

// ExtractFrom extracts the ServiceItemEvent from the given readable
func (serviceItemEvent *ServiceItemEvent) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemEvent.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemEvent header. %s", err.Error())
	}

	err = serviceItemEvent.EventID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.EventID from stream. %s", err.Error())
	}

	err = serviceItemEvent.ParamInt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.ParamInt from stream. %s", err.Error())
	}

	err = serviceItemEvent.ParamString.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.ParamString from stream. %s", err.Error())
	}

	serviceItemEvent.ParamBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.ParamBinary from stream. %s", err.Error())
	}

	err = serviceItemEvent.PresentTicketType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.PresentTicketType from stream. %s", err.Error())
	}

	err = serviceItemEvent.PresentTicketNum.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.PresentTicketNum from stream. %s", err.Error())
	}

	err = serviceItemEvent.TimeBegin.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.TimeBegin from stream. %s", err.Error())
	}

	err = serviceItemEvent.TimeEnd.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.TimeEnd from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemEvent to the given writable
func (serviceItemEvent *ServiceItemEvent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemEvent.EventID.WriteTo(contentWritable)
	serviceItemEvent.ParamInt.WriteTo(contentWritable)
	serviceItemEvent.ParamString.WriteTo(contentWritable)
	stream.WriteQBuffer(serviceItemEvent.ParamBinary)
	serviceItemEvent.PresentTicketType.WriteTo(contentWritable)
	serviceItemEvent.PresentTicketNum.WriteTo(contentWritable)
	serviceItemEvent.TimeBegin.WriteTo(contentWritable)
	serviceItemEvent.TimeEnd.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemEvent.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemEvent
func (serviceItemEvent *ServiceItemEvent) Copy() types.RVType {
	copied := NewServiceItemEvent()

	copied.StructureVersion = serviceItemEvent.StructureVersion

	copied.EventID = serviceItemEvent.EventID
	copied.ParamInt = serviceItemEvent.ParamInt
	copied.ParamString = serviceItemEvent.ParamString
	copied.ParamBinary = serviceItemEvent.ParamBinary
	copied.PresentTicketType = serviceItemEvent.PresentTicketType
	copied.PresentTicketNum = serviceItemEvent.PresentTicketNum
	copied.TimeBegin = serviceItemEvent.TimeBegin.Copy()
	copied.TimeEnd = serviceItemEvent.TimeEnd.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemEvent *ServiceItemEvent) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemEvent); !ok {
		return false
	}

	other := o.(*ServiceItemEvent)

	if serviceItemEvent.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemEvent.EventID.Equals(other.EventID) {
		return false
	}

	if !serviceItemEvent.ParamInt.Equals(other.ParamInt) {
		return false
	}

	if !serviceItemEvent.ParamString.Equals(other.ParamString) {
		return false
	}

	if !serviceItemEvent.ParamBinary.Equals(other.ParamBinary) {
		return false
	}

	if !serviceItemEvent.PresentTicketType.Equals(other.PresentTicketType) {
		return false
	}

	if !serviceItemEvent.PresentTicketNum.Equals(other.PresentTicketNum) {
		return false
	}

	if !serviceItemEvent.TimeBegin.Equals(other.TimeBegin) {
		return false
	}

	if !serviceItemEvent.TimeEnd.Equals(other.TimeEnd) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemEvent *ServiceItemEvent) String() string {
	return serviceItemEvent.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemEvent *ServiceItemEvent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemEvent{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemEvent.StructureVersion))
	b.WriteString(fmt.Sprintf("%sEventID: %d,\n", indentationValues, serviceItemEvent.EventID))
	b.WriteString(fmt.Sprintf("%sParamInt: %d,\n", indentationValues, serviceItemEvent.ParamInt))
	b.WriteString(fmt.Sprintf("%sParamString: %q,\n", indentationValues, serviceItemEvent.ParamString))
	b.WriteString(fmt.Sprintf("%sParamBinary: %x,\n", indentationValues, serviceItemEvent.ParamBinary))
	b.WriteString(fmt.Sprintf("%sPresentTicketType: %d,\n", indentationValues, serviceItemEvent.PresentTicketType))
	b.WriteString(fmt.Sprintf("%sPresentTicketNum: %d,\n", indentationValues, serviceItemEvent.PresentTicketNum))

	if serviceItemEvent.TimeBegin != nil {
		b.WriteString(fmt.Sprintf("%sTimeBegin: %s\n", indentationValues, serviceItemEvent.TimeBegin.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sTimeBegin: nil\n", indentationValues))
	}

	if serviceItemEvent.TimeEnd != nil {
		b.WriteString(fmt.Sprintf("%sTimeEnd: %s\n", indentationValues, serviceItemEvent.TimeEnd.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sTimeEnd: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemEvent returns a new ServiceItemEvent
func NewServiceItemEvent() *ServiceItemEvent {
	return &ServiceItemEvent{}
}
