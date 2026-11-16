package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplaySectionContentLocalized is a type within the Uplay Win protocol
type UplaySectionContentLocalized struct {
	types.Structure
	m_key      types.String
	m_culture  types.String
	m_text     types.String
	m_url      types.String
	m_duration types.Int32
	m_size     types.String
	m_width    types.String
	m_height   types.String
}

// ObjectID returns the object identifier of the type
func (uscl UplaySectionContentLocalized) ObjectID() types.RVType {
	return uscl.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (uscl UplaySectionContentLocalized) DataObjectID() types.RVType {
	return types.NewString("UplaySectionContentLocalized")
}

// WriteTo writes the UplaySectionContentLocalized to the given writable
func (uscl UplaySectionContentLocalized) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	uscl.m_key.WriteTo(contentWritable)
	uscl.m_culture.WriteTo(contentWritable)
	uscl.m_text.WriteTo(contentWritable)
	uscl.m_url.WriteTo(contentWritable)
	uscl.m_duration.WriteTo(contentWritable)
	uscl.m_size.WriteTo(contentWritable)
	uscl.m_width.WriteTo(contentWritable)
	uscl.m_height.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	uscl.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplaySectionContentLocalized from the given readable
func (uscl *UplaySectionContentLocalized) ExtractFrom(readable types.Readable) error {
	if err := uscl.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized Header. %s", err.Error())
	}

	if err := uscl.m_key.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_key. %s", err.Error())
	}

	if err := uscl.m_culture.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_culture. %s", err.Error())
	}

	if err := uscl.m_text.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_text. %s", err.Error())
	}

	if err := uscl.m_url.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_url. %s", err.Error())
	}

	if err := uscl.m_duration.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_duration. %s", err.Error())
	}

	if err := uscl.m_size.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_size. %s", err.Error())
	}

	if err := uscl.m_width.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_width. %s", err.Error())
	}

	if err := uscl.m_height.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_height. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplaySectionContentLocalized
func (uscl UplaySectionContentLocalized) Copy() types.RVType {
	copied := NewUplaySectionContentLocalized()

	copied.StructureVersion = uscl.StructureVersion
	copied.m_key = uscl.m_key
	copied.m_culture = uscl.m_culture
	copied.m_text = uscl.m_text
	copied.m_url = uscl.m_url
	copied.m_duration = uscl.m_duration
	copied.m_size = uscl.m_size
	copied.m_width = uscl.m_width
	copied.m_height = uscl.m_height

	return copied
}

// Equals checks if the given UplaySectionContentLocalized contains the same data as the current UplaySectionContentLocalized
func (uscl UplaySectionContentLocalized) Equals(o types.RVType) bool {
	if _, ok := o.(UplaySectionContentLocalized); !ok {
		return false
	}

	other := o.(UplaySectionContentLocalized)

	if uscl.StructureVersion != other.StructureVersion {
		return false
	}

	if !uscl.m_key.Equals(other.m_key) {
		return false
	}

	if !uscl.m_culture.Equals(other.m_culture) {
		return false
	}

	if !uscl.m_text.Equals(other.m_text) {
		return false
	}

	if !uscl.m_url.Equals(other.m_url) {
		return false
	}

	if !uscl.m_duration.Equals(other.m_duration) {
		return false
	}

	if !uscl.m_size.Equals(other.m_size) {
		return false
	}

	if !uscl.m_width.Equals(other.m_width) {
		return false
	}

	return uscl.m_height.Equals(other.m_height)
}

// CopyRef copies the current value of the UplaySectionContentLocalized
// and returns a pointer to the new copy
func (uscl UplaySectionContentLocalized) CopyRef() types.RVTypePtr {
	copied := uscl.Copy().(UplaySectionContentLocalized)
	return &copied
}

// Deref takes a pointer to the UplaySectionContentLocalized
// and dereferences it to the raw value.
func (uscl *UplaySectionContentLocalized) Deref() types.RVType {
	return *uscl
}

// String returns the string representation of the UplaySectionContentLocalized
func (uscl UplaySectionContentLocalized) String() string {
	return uscl.FormatToString(0)
}

// FormatToString pretty-prints the UplaySectionContentLocalized using the provided indentation level
func (uscl UplaySectionContentLocalized) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("UplaySectionContentLocalized{\n")
	fmt.Fprintf(&b, "%sm_key: %v,\n", indentationValues, uscl.m_key)
	fmt.Fprintf(&b, "%sm_culture: %v,\n", indentationValues, uscl.m_culture)
	fmt.Fprintf(&b, "%sm_text: %v,\n", indentationValues, uscl.m_text)
	fmt.Fprintf(&b, "%sm_url: %v,\n", indentationValues, uscl.m_url)
	fmt.Fprintf(&b, "%sm_duration: %v,\n", indentationValues, uscl.m_duration)
	fmt.Fprintf(&b, "%sm_size: %v,\n", indentationValues, uscl.m_size)
	fmt.Fprintf(&b, "%sm_width: %v,\n", indentationValues, uscl.m_width)
	fmt.Fprintf(&b, "%sm_height: %v,\n", indentationValues, uscl.m_height)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplaySectionContentLocalized returns a new UplaySectionContentLocalized
func NewUplaySectionContentLocalized() UplaySectionContentLocalized {
	return UplaySectionContentLocalized{
		m_key:      "key",
		m_culture:  "en",
		m_text:     "text",
		m_url:      "url",
		m_duration: 0,
		m_size:     "size",
		m_width:    "width",
		m_height:   "height",
	}
}
