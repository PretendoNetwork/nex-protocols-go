package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplaySectionContent is a type within the Uplay Win protocol
type UplaySectionContent struct {
	types.Structure
	m_key           types.String
	m_name          types.String
	m_order         types.Int16
	m_typeName      types.String
	m_localizedInfo UplaySectionContentLocalized
}

// ObjectID returns the object identifier of the type
func (usc UplaySectionContent) ObjectID() types.RVType {
	return usc.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (usc UplaySectionContent) DataObjectID() types.RVType {
	return types.NewString("UplaySectionContent")
}

// WriteTo writes the UplaySectionContent to the given writable
func (usc UplaySectionContent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	usc.m_key.WriteTo(contentWritable)
	usc.m_name.WriteTo(contentWritable)
	usc.m_order.WriteTo(contentWritable)
	usc.m_typeName.WriteTo(contentWritable)
	usc.m_localizedInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	usc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplaySectionContent from the given readable
func (usc *UplaySectionContent) ExtractFrom(readable types.Readable) error {
	if err := usc.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent Header. %s", err.Error())
	}

	if err := usc.m_key.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.m_key. %s", err.Error())
	}

	if err := usc.m_name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.m_name. %s", err.Error())
	}

	if err := usc.m_order.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.m_order. %s", err.Error())
	}

	if err := usc.m_typeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.m_typeName. %s", err.Error())
	}

	if err := usc.m_localizedInfo.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.m_localizedInfo. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplaySectionContent
func (usc UplaySectionContent) Copy() types.RVType {
	copied := NewUplaySectionContent()

	copied.StructureVersion = usc.StructureVersion
	copied.m_key = usc.m_key
	copied.m_name = usc.m_name
	copied.m_order = usc.m_order
	copied.m_typeName = usc.m_typeName
	copied.m_localizedInfo = usc.m_localizedInfo

	return copied
}

// Equals checks if the given UplaySectionContent contains the same data as the current UplaySectionContent
func (usc UplaySectionContent) Equals(o types.RVType) bool {
	if _, ok := o.(UplaySectionContent); !ok {
		return false
	}

	other := o.(UplaySectionContent)

	if usc.StructureVersion != other.StructureVersion {
		return false
	}

	if !usc.m_key.Equals(other.m_key) {
		return false
	}

	if !usc.m_name.Equals(other.m_name) {
		return false
	}

	if !usc.m_order.Equals(other.m_order) {
		return false
	}

	if !usc.m_typeName.Equals(other.m_typeName) {
		return false
	}

	return usc.m_localizedInfo.Equals(other.m_localizedInfo)
}

// CopyRef copies the current value of the UplaySectionContent
// and returns a pointer to the new copy
func (usc UplaySectionContent) CopyRef() types.RVTypePtr {
	copied := usc.Copy().(UplaySectionContent)
	return &copied
}

// Deref takes a pointer to the UplaySectionContent
// and dereferences it to the raw value.
func (usc *UplaySectionContent) Deref() types.RVType {
	return *usc
}

// String returns the string representation of the UplaySectionContent
func (usc UplaySectionContent) String() string {
	return usc.FormatToString(0)
}

// FormatToString pretty-prints the UplaySectionContent using the provided indentation level
func (usc UplaySectionContent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplaySectionContent{\n")
	fmt.Fprintf(&b, "%sm_key: %v,\n", indentationValues, usc.m_key)
	fmt.Fprintf(&b, "%sm_name: %v,\n", indentationValues, usc.m_name)
	fmt.Fprintf(&b, "%sm_order: %v,\n", indentationValues, usc.m_order)
	fmt.Fprintf(&b, "%sm_typeName: %v,\n", indentationValues, usc.m_typeName)
	fmt.Fprintf(&b, "%sm_localizedInfo: %v,\n", indentationValues, usc.m_localizedInfo)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplaySectionContent returns a new UplaySectionContent
func NewUplaySectionContent() UplaySectionContent {
	return UplaySectionContent{
		m_key:           "key",
		m_name:          "name",
		m_order:         0,
		m_typeName:      "type",
		m_localizedInfo: NewUplaySectionContentLocalized(),
	}
}
