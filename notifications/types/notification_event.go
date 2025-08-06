// Package types implements all the types used by the Notifications protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// NotificationEvent is a type within the Notifications protocol
type NotificationEvent struct {
	types.Structure
	PIDSource types.PID
	Type      types.UInt32
	Param1    types.UInt64 // * In NEX 3 this field is a UInt32. Storing as a UInt64
	Param2    types.UInt64 // * In NEX 3 this field is a UInt32. Storing as a UInt64
	StrParam  types.String
	Param3    types.UInt64                           // * NEX 3.4+
	MapParam  types.Map[types.String, types.Variant] // * NEX 4.0+ revision 1
}

// WriteTo writes the NotificationEvent to the given writable
func (ne NotificationEvent) WriteTo(writable types.Writable) {
	stream := writable.(*nex.ByteStreamOut)
	libraryVersion := stream.LibraryVersions.Main

	contentWritable := writable.CopyNew()

	ne.PIDSource.WriteTo(contentWritable)
	ne.Type.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("4.0.0") {
		ne.Param1.WriteTo(contentWritable)
	} else {
		contentWritable.WriteUInt32LE(uint32(ne.Param1))
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		ne.Param2.WriteTo(contentWritable)
	} else {
		contentWritable.WriteUInt32LE(uint32(ne.Param2))
	}

	ne.StrParam.WriteTo(contentWritable)

	if libraryVersion.GreaterOrEqual("4.0.0") {
		ne.Param3.WriteTo(contentWritable)
	} else if libraryVersion.GreaterOrEqual("3.4.0") {
		contentWritable.WriteUInt32LE(uint32(ne.Param3))
	}

	if libraryVersion.GreaterOrEqual("4.0.0") && ne.StructureVersion >= 1 {
		ne.MapParam.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	ne.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NotificationEvent from the given readable
func (ne *NotificationEvent) ExtractFrom(readable types.Readable) error {
	var err error

	stream := readable.(*nex.ByteStreamIn)
	libraryVersion := stream.LibraryVersions.Main

	err = ne.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NotificationEvent header. %s", err.Error())
	}

	err = ne.PIDSource.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NotificationEvent.PIDSource. %s", err.Error())
	}

	err = ne.Type.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NotificationEvent.Type. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		err = ne.Param1.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract NotificationEvent.Param1. %s", err.Error())
		}
	} else {
		param1, err := readable.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract NotificationEvent.Param1. %s", err.Error())
		}

		ne.Param1 = types.UInt64(param1)
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		err = ne.Param2.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract NotificationEvent.Param2. %s", err.Error())
		}
	} else {
		param2, err := readable.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract NotificationEvent.Param2. %s", err.Error())
		}

		ne.Param2 = types.UInt64(param2)
	}

	err = ne.StrParam.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NotificationEvent.StrParam. %s", err.Error())
	}

	if libraryVersion.GreaterOrEqual("4.0.0") {
		err = ne.Param3.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract NotificationEvent.Param3. %s", err.Error())
		}
	} else if libraryVersion.GreaterOrEqual("3.4.0") {
		param3, err := readable.ReadUInt32LE()
		if err != nil {
			return fmt.Errorf("Failed to extract NotificationEvent.Param3. %s", err.Error())
		}

		ne.Param3 = types.UInt64(param3)
	}

	if libraryVersion.GreaterOrEqual("4.0.0") && ne.StructureVersion >= 1 {
		err = ne.MapParam.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract NotificationEvent.MapParam. %s", err.Error())
		}
	}

	return nil
}

// Copy returns a new copied instance of NotificationEvent
func (ne NotificationEvent) Copy() types.RVType {
	copied := NewNotificationEvent()

	copied.StructureVersion = ne.StructureVersion
	copied.PIDSource = ne.PIDSource.Copy().(types.PID)
	copied.Type = ne.Type.Copy().(types.UInt32)
	copied.Param1 = ne.Param1.Copy().(types.UInt64)
	copied.Param2 = ne.Param2.Copy().(types.UInt64)
	copied.StrParam = ne.StrParam.Copy().(types.String)
	copied.Param3 = ne.Param3.Copy().(types.UInt64)
	copied.MapParam = ne.MapParam.Copy().(types.Map[types.String, types.Variant])

	return copied
}

// Equals checks if the given NotificationEvent contains the same data as the current NotificationEvent
func (ne NotificationEvent) Equals(o types.RVType) bool {
	if _, ok := o.(NotificationEvent); !ok {
		return false
	}

	other := o.(NotificationEvent)

	if ne.StructureVersion != other.StructureVersion {
		return false
	}

	if !ne.PIDSource.Equals(other.PIDSource) {
		return false
	}

	if !ne.Type.Equals(other.Type) {
		return false
	}

	if !ne.Param1.Equals(other.Param1) {
		return false
	}

	if !ne.Param2.Equals(other.Param2) {
		return false
	}

	if !ne.StrParam.Equals(other.StrParam) {
		return false
	}

	if !ne.Param3.Equals(other.Param3) {
		return false
	}

	return ne.MapParam.Equals(other.MapParam)
}

// CopyRef copies the current value of the NotificationEvent
// and returns a pointer to the new copy
func (ne NotificationEvent) CopyRef() types.RVTypePtr {
	copied := ne.Copy().(NotificationEvent)
	return &copied
}

// Deref takes a pointer to the NotificationEvent
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (ne *NotificationEvent) Deref() types.RVType {
	return *ne
}

// String returns the string representation of the NotificationEvent
func (ne NotificationEvent) String() string {
	return ne.FormatToString(0)
}

// FormatToString pretty-prints the NotificationEvent using the provided indentation level
func (ne NotificationEvent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NotificationEvent{\n")
	b.WriteString(fmt.Sprintf("%sPIDSource: %s,\n", indentationValues, ne.PIDSource.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sType: %s,\n", indentationValues, ne.Type))
	b.WriteString(fmt.Sprintf("%sParam1: %s,\n", indentationValues, ne.Param1))
	b.WriteString(fmt.Sprintf("%sParam2: %s,\n", indentationValues, ne.Param2))
	b.WriteString(fmt.Sprintf("%sStrParam: %s,\n", indentationValues, ne.StrParam))
	b.WriteString(fmt.Sprintf("%sParam3: %s,\n", indentationValues, ne.Param3))
	b.WriteString(fmt.Sprintf("%sMapParam: %s,\n", indentationValues, ne.MapParam))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNotificationEvent returns a new NotificationEvent
func NewNotificationEvent() NotificationEvent {
	return NotificationEvent{
		PIDSource: types.NewPID(0),
		Type:      types.NewUInt32(0),
		Param1:    types.NewUInt64(0),
		Param2:    types.NewUInt64(0),
		StrParam:  types.NewString(""),
		Param3:    types.NewUInt64(0),
		MapParam:  types.NewMap[types.String, types.Variant](),
	}

}
