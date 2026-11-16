package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplayActionPlatform is a type within the Uplay Win protocol
type UplayActionPlatform struct {
	types.Structure
	PlatformCode types.String
	Completed    types.Bool
	SpecificKey  types.String
}

// ObjectID returns the object identifier of the type
func (uap UplayActionPlatform) ObjectID() types.RVType {
	return uap.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (uap UplayActionPlatform) DataObjectID() types.RVType {
	return types.NewString("UplayActionPlatform")
}

// WriteTo writes the UplayActionPlatform to the given writable
func (uap UplayActionPlatform) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	uap.PlatformCode.WriteTo(contentWritable)
	uap.Completed.WriteTo(contentWritable)
	uap.SpecificKey.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	uap.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplayActionPlatform from the given readable
func (uap *UplayActionPlatform) ExtractFrom(readable types.Readable) error {
	if err := uap.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayActionPlatform Header. %s", err.Error())
	}

	if err := uap.PlatformCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayActionPlatform.PlatformCode. %s", err.Error())
	}

	if err := uap.Completed.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayActionPlatform.Completed. %s", err.Error())
	}

	if err := uap.SpecificKey.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayActionPlatform.SpecificKey. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplayActionPlatform
func (uap UplayActionPlatform) Copy() types.RVType {
	copied := NewUplayActionPlatform()

	copied.StructureVersion = uap.StructureVersion
	copied.PlatformCode = uap.PlatformCode
	copied.Completed = uap.Completed
	copied.SpecificKey = uap.SpecificKey

	return copied
}

// Equals checks if the given UplayActionPlatform contains the same data as the current UplayActionPlatform
func (uap UplayActionPlatform) Equals(o types.RVType) bool {
	if _, ok := o.(UplayActionPlatform); !ok {
		return false
	}

	other := o.(UplayActionPlatform)

	if uap.StructureVersion != other.StructureVersion {
		return false
	}

	if !uap.PlatformCode.Equals(other.PlatformCode) {
		return false
	}

	if !uap.Completed.Equals(other.Completed) {
		return false
	}

	return uap.SpecificKey.Equals(other.SpecificKey)
}

// CopyRef copies the current value of the UplayActionPlatform
// and returns a pointer to the new copy
func (uap UplayActionPlatform) CopyRef() types.RVTypePtr {
	copied := uap.Copy().(UplayActionPlatform)
	return &copied
}

// Deref takes a pointer to the UplayActionPlatform
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (uap *UplayActionPlatform) Deref() types.RVType {
	return *uap
}

// String returns the string representation of the UplayActionPlatform
func (uap UplayActionPlatform) String() string {
	return uap.FormatToString(0)
}

// FormatToString pretty-prints the UplayActionPlatform using the provided indentation level
func (uap UplayActionPlatform) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplayActionPlatform{\n")
	fmt.Fprintf(&b, "%sm_platformCode: %v,\n", indentationValues, uap.PlatformCode)
	fmt.Fprintf(&b, "%sm_completed: %v,\n", indentationValues, uap.Completed)
	fmt.Fprintf(&b, "%sm_specificKey: %v,\n", indentationValues, uap.SpecificKey)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplayActionPlatform returns a new UplayActionPlatform
func NewUplayActionPlatform() UplayActionPlatform {
	return UplayActionPlatform{
		PlatformCode: "PC",
		Completed:    false,
		SpecificKey:  "",
	}
}
