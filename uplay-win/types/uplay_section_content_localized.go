package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplaySectionContentLocalized is a type within the Uplay Win protocol
type UplaySectionContentLocalized struct {
	types.Structure
	CultureName types.String
	Name        types.String
	Description types.String
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

	uscl.CultureName.WriteTo(contentWritable)
	uscl.Name.WriteTo(contentWritable)
	uscl.Description.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	uscl.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplaySectionContentLocalized from the given readable
func (uscl *UplaySectionContentLocalized) ExtractFrom(readable types.Readable) error {
	if err := uscl.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized Header. %s", err.Error())
	}

	if err := uscl.CultureName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.CultureName. %s", err.Error())
	}

	if err := uscl.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.Name. %s", err.Error())
	}

	if err := uscl.Description.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContentLocalized.Description. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplaySectionContentLocalized
func (uscl UplaySectionContentLocalized) Copy() types.RVType {
	copied := NewUplaySectionContentLocalized()

	copied.StructureVersion = uscl.StructureVersion
	copied.CultureName = uscl.CultureName
	copied.Name = uscl.Name
	copied.Description = uscl.Description

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

	if !uscl.CultureName.Equals(other.CultureName) {
		return false
	}

	if !uscl.Name.Equals(other.Name) {
		return false
	}

	return uscl.Description.Equals(other.Description)
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
	fmt.Fprintf(&b, "%sm_cultureName: %v,\n", indentationValues, uscl.CultureName)
	fmt.Fprintf(&b, "%sm_name: %v,\n", indentationValues, uscl.Name)
	fmt.Fprintf(&b, "%sm_description: %v,\n", indentationValues, uscl.Description)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplaySectionContentLocalized returns a new UplaySectionContentLocalized
func NewUplaySectionContentLocalized() UplaySectionContentLocalized {
	return UplaySectionContentLocalized{
		CultureName: "en",
		Name:        "name",
		Description: "desc",
	}
}
