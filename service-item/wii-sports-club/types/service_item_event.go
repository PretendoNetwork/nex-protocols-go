// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemEvent holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemEvent struct {
	nex.Structure
	EventID           uint64
	ParamInt          int32
	ParamString       string
	ParamBinary       []byte
	PresentTicketType uint32
	PresentTicketNum  uint32
	TimeBegin         *nex.DateTime
	TimeEnd           *nex.DateTime
}

// ExtractFromStream extracts a ServiceItemEvent structure from a stream
func (serviceItemEvent *ServiceItemEvent) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemEvent.EventID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.EventID from stream. %s", err.Error())
	}

	serviceItemEvent.ParamInt, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.ParamInt from stream. %s", err.Error())
	}

	serviceItemEvent.ParamString, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.ParamString from stream. %s", err.Error())
	}

	serviceItemEvent.ParamBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.ParamBinary from stream. %s", err.Error())
	}

	serviceItemEvent.PresentTicketType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.PresentTicketType from stream. %s", err.Error())
	}

	serviceItemEvent.PresentTicketNum, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.PresentTicketNum from stream. %s", err.Error())
	}

	serviceItemEvent.TimeBegin, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.TimeBegin from stream. %s", err.Error())
	}

	serviceItemEvent.TimeEnd, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEvent.TimeEnd from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemEvent and returns a byte array
func (serviceItemEvent *ServiceItemEvent) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(serviceItemEvent.EventID)
	stream.WriteInt32LE(serviceItemEvent.ParamInt)
	stream.WriteString(serviceItemEvent.ParamString)
	stream.WriteQBuffer(serviceItemEvent.ParamBinary)
	stream.WriteUInt32LE(serviceItemEvent.PresentTicketType)
	stream.WriteUInt32LE(serviceItemEvent.PresentTicketNum)
	stream.WriteDateTime(serviceItemEvent.TimeBegin)
	stream.WriteDateTime(serviceItemEvent.TimeEnd)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemEvent
func (serviceItemEvent *ServiceItemEvent) Copy() nex.StructureInterface {
	copied := NewServiceItemEvent()

	copied.SetStructureVersion(serviceItemEvent.StructureVersion())

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
func (serviceItemEvent *ServiceItemEvent) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemEvent)

	if serviceItemEvent.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemEvent.EventID != other.EventID {
		return false
	}

	if serviceItemEvent.ParamInt != other.ParamInt {
		return false
	}

	if serviceItemEvent.ParamString != other.ParamString {
		return false
	}

	if !bytes.Equal(serviceItemEvent.ParamBinary, other.ParamBinary) {
		return false
	}

	if serviceItemEvent.PresentTicketType != other.PresentTicketType {
		return false
	}

	if serviceItemEvent.PresentTicketNum != other.PresentTicketNum {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemEvent.StructureVersion()))
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
