// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemNotice holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemNotice struct {
	nex.Structure
	ScheduleID   uint64
	ScheduleType uint32
	ParamInt     int32
	ParamString  string
	ParamBinary  []byte
	TimeBegin    *nex.DateTime
	TimeEnd      *nex.DateTime
}

// ExtractFromStream extracts a ServiceItemNotice structure from a stream
func (serviceItemNotice *ServiceItemNotice) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemNotice.ScheduleID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ScheduleID from stream. %s", err.Error())
	}

	serviceItemNotice.ScheduleType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ScheduleType from stream. %s", err.Error())
	}

	serviceItemNotice.ParamInt, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ParamInt from stream. %s", err.Error())
	}

	serviceItemNotice.ParamString, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ParamString from stream. %s", err.Error())
	}

	serviceItemNotice.ParamBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ParamBinary from stream. %s", err.Error())
	}

	serviceItemNotice.TimeBegin, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.TimeBegin from stream. %s", err.Error())
	}

	serviceItemNotice.TimeEnd, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.TimeEnd from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemNotice and returns a byte array
func (serviceItemNotice *ServiceItemNotice) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(serviceItemNotice.ScheduleID)
	stream.WriteUInt32LE(serviceItemNotice.ScheduleType)
	stream.WriteInt32LE(serviceItemNotice.ParamInt)
	stream.WriteString(serviceItemNotice.ParamString)
	stream.WriteQBuffer(serviceItemNotice.ParamBinary)
	stream.WriteDateTime(serviceItemNotice.TimeBegin)
	stream.WriteDateTime(serviceItemNotice.TimeEnd)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemNotice
func (serviceItemNotice *ServiceItemNotice) Copy() nex.StructureInterface {
	copied := NewServiceItemNotice()

	copied.SetStructureVersion(serviceItemNotice.StructureVersion())

	copied.ScheduleID = serviceItemNotice.ScheduleID
	copied.ScheduleType = serviceItemNotice.ScheduleType
	copied.ParamInt = serviceItemNotice.ParamInt
	copied.ParamString = serviceItemNotice.ParamString
	copied.ParamBinary = serviceItemNotice.ParamBinary
	copied.TimeBegin = serviceItemNotice.TimeBegin.Copy()
	copied.TimeEnd = serviceItemNotice.TimeEnd.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemNotice *ServiceItemNotice) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemNotice)

	if serviceItemNotice.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemNotice.ScheduleID != other.ScheduleID {
		return false
	}

	if serviceItemNotice.ScheduleType != other.ScheduleType {
		return false
	}

	if serviceItemNotice.ParamInt != other.ParamInt {
		return false
	}

	if serviceItemNotice.ParamString != other.ParamString {
		return false
	}

	if !bytes.Equal(serviceItemNotice.ParamBinary, other.ParamBinary) {
		return false
	}

	if !serviceItemNotice.TimeBegin.Equals(other.TimeBegin) {
		return false
	}

	if !serviceItemNotice.TimeEnd.Equals(other.TimeEnd) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemNotice *ServiceItemNotice) String() string {
	return serviceItemNotice.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemNotice *ServiceItemNotice) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemNotice{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemNotice.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sScheduleID: %d,\n", indentationValues, serviceItemNotice.ScheduleID))
	b.WriteString(fmt.Sprintf("%sScheduleType: %d,\n", indentationValues, serviceItemNotice.ScheduleType))
	b.WriteString(fmt.Sprintf("%sParamInt: %d,\n", indentationValues, serviceItemNotice.ParamInt))
	b.WriteString(fmt.Sprintf("%sParamString: %q,\n", indentationValues, serviceItemNotice.ParamString))
	b.WriteString(fmt.Sprintf("%sParamBinary: %x,\n", indentationValues, serviceItemNotice.ParamBinary))

	if serviceItemNotice.TimeBegin != nil {
		b.WriteString(fmt.Sprintf("%sTimeBegin: %s\n", indentationValues, serviceItemNotice.TimeBegin.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sTimeBegin: nil\n", indentationValues))
	}

	if serviceItemNotice.TimeEnd != nil {
		b.WriteString(fmt.Sprintf("%sTimeEnd: %s\n", indentationValues, serviceItemNotice.TimeEnd.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sTimeEnd: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemNotice returns a new ServiceItemNotice
func NewServiceItemNotice() *ServiceItemNotice {
	return &ServiceItemNotice{}
}
