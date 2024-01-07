// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemNotice holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemNotice struct {
	types.Structure
	ScheduleID   *types.PrimitiveU64
	ScheduleType *types.PrimitiveU32
	ParamInt     *types.PrimitiveS32
	ParamString  string
	ParamBinary  []byte
	TimeBegin    *types.DateTime
	TimeEnd      *types.DateTime
}

// ExtractFrom extracts the ServiceItemNotice from the given readable
func (serviceItemNotice *ServiceItemNotice) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemNotice.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemNotice header. %s", err.Error())
	}

	err = serviceItemNotice.ScheduleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ScheduleID from stream. %s", err.Error())
	}

	err = serviceItemNotice.ScheduleType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ScheduleType from stream. %s", err.Error())
	}

	err = serviceItemNotice.ParamInt.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ParamInt from stream. %s", err.Error())
	}

	err = serviceItemNotice.ParamString.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ParamString from stream. %s", err.Error())
	}

	serviceItemNotice.ParamBinary, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.ParamBinary from stream. %s", err.Error())
	}

	err = serviceItemNotice.TimeBegin.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.TimeBegin from stream. %s", err.Error())
	}

	err = serviceItemNotice.TimeEnd.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemNotice.TimeEnd from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemNotice to the given writable
func (serviceItemNotice *ServiceItemNotice) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemNotice.ScheduleID.WriteTo(contentWritable)
	serviceItemNotice.ScheduleType.WriteTo(contentWritable)
	serviceItemNotice.ParamInt.WriteTo(contentWritable)
	serviceItemNotice.ParamString.WriteTo(contentWritable)
	stream.WriteQBuffer(serviceItemNotice.ParamBinary)
	serviceItemNotice.TimeBegin.WriteTo(contentWritable)
	serviceItemNotice.TimeEnd.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemNotice.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemNotice
func (serviceItemNotice *ServiceItemNotice) Copy() types.RVType {
	copied := NewServiceItemNotice()

	copied.StructureVersion = serviceItemNotice.StructureVersion

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
func (serviceItemNotice *ServiceItemNotice) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemNotice); !ok {
		return false
	}

	other := o.(*ServiceItemNotice)

	if serviceItemNotice.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemNotice.ScheduleID.Equals(other.ScheduleID) {
		return false
	}

	if !serviceItemNotice.ScheduleType.Equals(other.ScheduleType) {
		return false
	}

	if !serviceItemNotice.ParamInt.Equals(other.ParamInt) {
		return false
	}

	if !serviceItemNotice.ParamString.Equals(other.ParamString) {
		return false
	}

	if !serviceItemNotice.ParamBinary.Equals(other.ParamBinary) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemNotice.StructureVersion))
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
