package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplaySection is a type within the Uplay Win protocol
type UplaySection struct {
	types.Structure
	m_key          types.String
	m_name         types.String
	m_typeName     types.String
	m_menuTypeName types.String
	m_contentList  types.List[UplaySectionContent]
	m_gameCode     types.String
	m_platformCode types.String
}

// ObjectID returns the object identifier of the type
func (us UplaySection) ObjectID() types.RVType {
	return us.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (us UplaySection) DataObjectID() types.RVType {
	return types.NewString("UplaySection")
}

// WriteTo writes the UplaySection to the given writable
func (us UplaySection) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	us.m_key.WriteTo(contentWritable)
	us.m_name.WriteTo(contentWritable)
	us.m_typeName.WriteTo(contentWritable)
	us.m_menuTypeName.WriteTo(contentWritable)
	us.m_contentList.WriteTo(contentWritable)
	us.m_gameCode.WriteTo(contentWritable)
	us.m_platformCode.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	us.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplaySection from the given readable
func (us *UplaySection) ExtractFrom(readable types.Readable) error {
	if err := us.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection Header. %s", err.Error())
	}

	if err := us.m_key.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.m_key. %s", err.Error())
	}

	if err := us.m_name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.m_name. %s", err.Error())
	}

	if err := us.m_typeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.m_typeName. %s", err.Error())
	}

	if err := us.m_menuTypeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.m_menuTypeName. %s", err.Error())
	}

	if err := us.m_contentList.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.m_contentList. %s", err.Error())
	}

	if err := us.m_gameCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.m_gameCode. %s", err.Error())
	}

	if err := us.m_platformCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySection.m_platformCode. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplaySection
func (us UplaySection) Copy() types.RVType {
	copied := NewUplaySection()

	copied.StructureVersion = us.StructureVersion
	copied.m_key = us.m_key
	copied.m_name = us.m_name
	copied.m_typeName = us.m_typeName
	copied.m_menuTypeName = us.m_menuTypeName
	copied.m_contentList = us.m_contentList
	copied.m_gameCode = us.m_gameCode
	copied.m_platformCode = us.m_platformCode

	return copied
}

// Equals checks if the given UplaySection contains the same data as the current UplaySection
func (us UplaySection) Equals(o types.RVType) bool {
	if _, ok := o.(UplaySection); !ok {
		return false
	}

	other := o.(UplaySection)

	if us.StructureVersion != other.StructureVersion {
		return false
	}

	if !us.m_key.Equals(other.m_key) {
		return false
	}

	if !us.m_name.Equals(other.m_name) {
		return false
	}

	if !us.m_typeName.Equals(other.m_typeName) {
		return false
	}

	if !us.m_menuTypeName.Equals(other.m_menuTypeName) {
		return false
	}

	if !us.m_contentList.Equals(other.m_contentList) {
		return false
	}

	if !us.m_gameCode.Equals(other.m_gameCode) {
		return false
	}

	return us.m_platformCode.Equals(other.m_platformCode)
}

// CopyRef copies the current value of the UplaySection
// and returns a pointer to the new copy
func (us UplaySection) CopyRef() types.RVTypePtr {
	copied := us.Copy().(UplaySection)
	return &copied
}

// Deref takes a pointer to the UplaySection
// and dereferences it to the raw value.
func (us *UplaySection) Deref() types.RVType {
	return *us
}

// String returns the string representation of the UplaySection
func (us UplaySection) String() string {
	return us.FormatToString(0)
}

// FormatToString pretty-prints the UplaySection using the provided indentation level
func (us UplaySection) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplaySection{\n")
	fmt.Fprintf(&b, "%sm_key: %v,\n", indentationValues, us.m_key)
	fmt.Fprintf(&b, "%sm_name: %v,\n", indentationValues, us.m_name)
	fmt.Fprintf(&b, "%sm_typeName: %v,\n", indentationValues, us.m_typeName)
	fmt.Fprintf(&b, "%sm_menuTypeName: %v,\n", indentationValues, us.m_menuTypeName)
	fmt.Fprintf(&b, "%sm_contentList: %v,\n", indentationValues, us.m_contentList)
	fmt.Fprintf(&b, "%sm_gameCode: %v,\n", indentationValues, us.m_gameCode)
	fmt.Fprintf(&b, "%sm_platformCode: %v,\n", indentationValues, us.m_platformCode)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplaySection returns a new UplaySection
func NewUplaySection() UplaySection {
	return UplaySection{
		m_key:          "key",
		m_name:         "name",
		m_typeName:     "type",
		m_menuTypeName: "menu",
		m_contentList:  types.NewList[UplaySectionContent](),
		m_gameCode:     "Uplay",
		m_platformCode: "PC",
	}
}
