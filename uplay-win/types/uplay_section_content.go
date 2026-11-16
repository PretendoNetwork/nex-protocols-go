package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// UplaySectionContent is a type within the Uplay Win protocol
type UplaySectionContent struct {
	types.Structure
	Key           types.String
	Name          types.String
	Order         types.Int16
	TypeName      types.String
	LocalizedInfo UplaySectionContentLocalized
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

	usc.Key.WriteTo(contentWritable)
	usc.Name.WriteTo(contentWritable)
	usc.Order.WriteTo(contentWritable)
	usc.TypeName.WriteTo(contentWritable)
	usc.LocalizedInfo.WriteTo(contentWritable)

	content := contentWritable.Bytes()
	usc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the UplaySectionContent from the given readable
func (usc *UplaySectionContent) ExtractFrom(readable types.Readable) error {
	if err := usc.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent Header. %s", err.Error())
	}

	if err := usc.Key.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.Key. %s", err.Error())
	}

	if err := usc.Name.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.Name. %s", err.Error())
	}

	if err := usc.Order.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.Order. %s", err.Error())
	}

	if err := usc.TypeName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.TypeName. %s", err.Error())
	}

	if err := usc.LocalizedInfo.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract UplaySectionContent.LocalizedInfo. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of UplaySectionContent
func (usc UplaySectionContent) Copy() types.RVType {
	copied := NewUplaySectionContent()

	copied.StructureVersion = usc.StructureVersion
	copied.Key = usc.Key
	copied.Name = usc.Name
	copied.Order = usc.Order
	copied.TypeName = usc.TypeName
	copied.LocalizedInfo = usc.LocalizedInfo

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

	if !usc.Key.Equals(other.Key) {
		return false
	}

	if !usc.Name.Equals(other.Name) {
		return false
	}

	if !usc.Order.Equals(other.Order) {
		return false
	}

	if !usc.TypeName.Equals(other.TypeName) {
		return false
	}

	return usc.LocalizedInfo.Equals(other.LocalizedInfo)
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
	fmt.Fprintf(&b, "%sm_key: %v,\n", indentationValues, usc.Key)
	fmt.Fprintf(&b, "%sm_name: %v,\n", indentationValues, usc.Name)
	fmt.Fprintf(&b, "%sm_order: %v,\n", indentationValues, usc.Order)
	fmt.Fprintf(&b, "%sm_typeName: %v,\n", indentationValues, usc.TypeName)
	fmt.Fprintf(&b, "%sm_localizedInfo: %v,\n", indentationValues, usc.LocalizedInfo)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewUplaySectionContent returns a new UplaySectionContent
func NewUplaySectionContent() UplaySectionContent {
	return UplaySectionContent{
		Key:           "key",
		Name:          "name",
		Order:         0,
		TypeName:      "type",
		LocalizedInfo: NewUplaySectionContentLocalized(),
	}
}
