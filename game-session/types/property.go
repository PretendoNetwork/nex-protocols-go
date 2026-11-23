// Package types implements all the types used by the GameSession protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Property is a type with the GameSession protocol
type Property struct {
	types.Structure
	ID    types.UInt32
	Value types.UInt32
}

// WriteTo writes the Property to the given writable
func (p Property) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	p.ID.WriteTo(contentWritable)
	p.Value.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	p.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Property from the given readable
func (p *Property) ExtractFrom(readable types.Readable) error {
	if err := p.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract Property header. %s", err.Error())
	}

	if err := p.ID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract Property.ID. %s", err.Error())
	}

	if err := p.Value.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract Property.Value. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Property
func (p Property) Copy() types.RVType {
	copied := NewProperty()

	copied.StructureVersion = p.StructureVersion
	copied.ID = p.ID
	copied.Value = p.Value
	return copied
}

// Equals checks if the given Property contains the same data as the current Property
func (p Property) Equals(o types.RVType) bool {
	if _, ok := o.(Property); !ok {
		return false
	}

	other := o.(Property)

	if p.StructureVersion != other.StructureVersion {
		return false
	}

	if p.ID != other.ID {
		return false
	}

	return p.Value == other.Value
}

// CopyRef copies the current value of the Property
// and returns a pointer to the new copy
func (p Property) CopyRef() types.RVTypePtr {
	copied := p.Copy().(Property)
	return &copied
}

// Deref takes a pointer of Property
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (p *Property) Deref() types.RVType {
	return *p
}

// String returns the string representation of the Property
func (p Property) String() string {
	return p.FormatToString(0)
}

// FormatToString pretty-prints the Property using the provided indentation level
func (p Property) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Property{\n")
	fmt.Fprintf(&b, "%sID: %s,\n", indentationValues, p.ID)
	fmt.Fprintf(&b, "%sValue: %s,\n", indentationValues, p.Value)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewProperty returns a new Property
func NewProperty() Property {
	return Property{
		ID:    types.NewUInt32(0),
		Value: types.NewUInt32(0),
	}
}
