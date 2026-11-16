package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplayAction is a type within the Uplay Win protocol
type UplayAction struct {
	types.Structure
	m_code        types.String
	m_name        types.String
	m_description types.String
	m_value       types.Int32
	m_gameCode    types.String
	m_platforms   types.List[UplayActionPlatform]
}

// ObjectID returns the object identifier of the type
func (upa UplayAction) ObjectID() types.RVType {
	return upa.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (upa UplayAction) DataObjectID() types.RVType {
	return types.NewString("UplayAction")
}

// WriteTo writes the UplayAction to the given writable
func (upa UplayAction) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	upa.m_code.WriteTo(contentWritable)
	upa.m_name.WriteTo(contentWritable)
	upa.m_description.WriteTo(contentWritable)
	upa.m_value.WriteTo(contentWritable)
	upa.m_gameCode.WriteTo(contentWritable)
	upa.m_platforms.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	upa.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplayAction from the given readable
func (upa *UplayAction) ExtractFrom(readable types.Readable) error {
	if err := upa.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction Header. %s", err.Error())
	}

	if err := upa.m_code.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.m_code. %s", err.Error())
	}

	if err := upa.m_name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.m_name. %s", err.Error())
	}

	if err := upa.m_description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.m_description. %s", err.Error())
	}

	if err := upa.m_value.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.m_value. %s", err.Error())
	}

	if err := upa.m_gameCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.m_gameCode. %s", err.Error())
	}

	if err := upa.m_platforms.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplayAction.m_platforms. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplayAction
func (upa UplayAction) Copy() types.RVType {
	copied := NewUplayAction()

	copied.StructureVersion = upa.StructureVersion
	copied.m_code = upa.m_code
	copied.m_name = upa.m_name
	copied.m_description = upa.m_description
	copied.m_value = upa.m_value
	copied.m_gameCode = upa.m_gameCode
	copied.m_platforms = upa.m_platforms

	return copied
}

// Equals checks if the given UplayAction contains the same data as the current UplayAction
func (upa UplayAction) Equals(o types.RVType) bool {
	if _, ok := o.(UplayAction); !ok {
		return false
	}

	other := o.(UplayAction)

	if upa.StructureVersion != other.StructureVersion {
		return false
	}

	if !upa.m_code.Equals(other.m_code) {
		return false
	}

	if !upa.m_name.Equals(other.m_name) {
		return false
	}

	if !upa.m_description.Equals(other.m_description) {
		return false
	}

	if !upa.m_value.Equals(other.m_value) {
		return false
	}

	if !upa.m_gameCode.Equals(other.m_gameCode) {
		return false
	}

	return upa.m_platforms.Equals(other.m_platforms)
}

// CopyRef copies the current value of the UplayAction
// and returns a pointer to the new copy
func (upa UplayAction) CopyRef() types.RVTypePtr {
	copied := upa.Copy().(UplayAction)
	return &copied
}

// Deref takes a pointer to the UplayAction
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (upa *UplayAction) Deref() types.RVType {
	return *upa
}

// String returns the string representation of the UplayAction
func (upa UplayAction) String() string {
	return upa.FormatToString(0)
}

// FormatToString pretty-prints the UplayAction using the provided indentation level
func (upa UplayAction) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplayAction{\n")
	fmt.Fprintf(&b, "%sm_code: %v,\n", indentationValues, upa.m_code)
	fmt.Fprintf(&b, "%sm_name: %v,\n", indentationValues, upa.m_name)
	fmt.Fprintf(&b, "%sm_description: %v,\n", indentationValues, upa.m_description)
	fmt.Fprintf(&b, "%sm_value: %v,\n", indentationValues, upa.m_value)
	fmt.Fprintf(&b, "%sm_gameCode: %v,\n", indentationValues, upa.m_gameCode)
	fmt.Fprintf(&b, "%sm_platforms: %v,\n", indentationValues, upa.m_platforms)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplayAction returns a new UplayAction
func NewUplayAction() UplayAction {
	return UplayAction{
		m_code:        "code",
		m_name:        "name",
		m_description: "desc",
		m_value:       0,
		m_gameCode:    "Uplay",
		m_platforms:   types.NewList[UplayActionPlatform](),
	}
}
