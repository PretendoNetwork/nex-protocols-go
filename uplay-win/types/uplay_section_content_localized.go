package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplaySectionContentLocalized is a type within the Uplay Win protocol
type UplaySectionContentLocalized struct {
	types.Structure
	Key      types.String
	Culture  types.String
	Text     types.String
	URL      types.String
	Duration types.Int32
	Size     types.String
	Width    types.String
	Height   types.String
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

	uscl.Key.WriteTo(contentWritable)
	uscl.Culture.WriteTo(contentWritable)
	uscl.Text.WriteTo(contentWritable)
	uscl.URL.WriteTo(contentWritable)
	uscl.Duration.WriteTo(contentWritable)
	uscl.Size.WriteTo(contentWritable)
	uscl.Width.WriteTo(contentWritable)
	uscl.Height.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	uscl.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplaySectionContentLocalized from the given readable
func (uscl *UplaySectionContentLocalized) ExtractFrom(readable types.Readable) error {
	if err := uscl.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized Header. %s", err.Error())
	}

	if err := uscl.Key.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_key. %s", err.Error())
	}

	if err := uscl.Culture.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_culture. %s", err.Error())
	}

	if err := uscl.Text.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_text. %s", err.Error())
	}

	if err := uscl.URL.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_url. %s", err.Error())
	}

	if err := uscl.Duration.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_duration. %s", err.Error())
	}

	if err := uscl.Size.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_size. %s", err.Error())
	}

	if err := uscl.Width.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_width. %s", err.Error())
	}

	if err := uscl.Height.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.m_height. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplaySectionContentLocalized
func (uscl UplaySectionContentLocalized) Copy() types.RVType {
	copied := NewUplaySectionContentLocalized()

	copied.StructureVersion = uscl.StructureVersion
	copied.Key = uscl.Key
	copied.Culture = uscl.Culture
	copied.Text = uscl.Text
	copied.URL = uscl.URL
	copied.Duration = uscl.Duration
	copied.Size = uscl.Size
	copied.Width = uscl.Width
	copied.Height = uscl.Height

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

	if !uscl.Key.Equals(other.Key) {
		return false
	}

	if !uscl.Culture.Equals(other.Culture) {
		return false
	}

	if !uscl.Text.Equals(other.Text) {
		return false
	}

	if !uscl.URL.Equals(other.URL) {
		return false
	}

	if !uscl.Duration.Equals(other.Duration) {
		return false
	}

	if !uscl.Size.Equals(other.Size) {
		return false
	}

	if !uscl.Width.Equals(other.Width) {
		return false
	}

	return uscl.Height.Equals(other.Height)
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
	fmt.Fprintf(&b, "%sKey: %v,\n", indentationValues, uscl.Key)
	fmt.Fprintf(&b, "%sCulture: %v,\n", indentationValues, uscl.Culture)
	fmt.Fprintf(&b, "%sText: %v,\n", indentationValues, uscl.Text)
	fmt.Fprintf(&b, "%sURL: %v,\n", indentationValues, uscl.URL)
	fmt.Fprintf(&b, "%sDuration: %v,\n", indentationValues, uscl.Duration)
	fmt.Fprintf(&b, "%sSize: %v,\n", indentationValues, uscl.Size)
	fmt.Fprintf(&b, "%sWidth: %v,\n", indentationValues, uscl.Width)
	fmt.Fprintf(&b, "%sHeight: %v,\n", indentationValues, uscl.Height)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplaySectionContentLocalized returns a new UplaySectionContentLocalized
func NewUplaySectionContentLocalized() UplaySectionContentLocalized {
	return UplaySectionContentLocalized{
		Key:      types.NewString(""),
		Culture:  types.NewString(""),
		Text:     types.NewString(""),
		URL:      types.NewString(""),
		Duration: types.NewInt32(0),
		Size:     types.NewString(""),
		Width:    types.NewString(""),
		Height:   types.NewString(""),
	}
}
