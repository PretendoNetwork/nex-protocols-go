// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Mii is a type within the Friends3DS protocol
type Mii struct {
	types.Structure
	types.Data
	Name          types.String
	ProfanityFlag types.Bool
	CharacterSet  types.UInt8
	MiiData       types.Buffer
}

// WriteTo writes the Mii to the given writable
func (m Mii) WriteTo(writable types.Writable) {
	m.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	m.Name.WriteTo(contentWritable)
	m.ProfanityFlag.WriteTo(contentWritable)
	m.CharacterSet.WriteTo(contentWritable)
	m.MiiData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	m.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Mii from the given readable
func (m *Mii) ExtractFrom(readable types.Readable) error {
	var err error

	err = m.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.Data. %s", err.Error())
	}

	err = m.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii header. %s", err.Error())
	}

	err = m.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.Name. %s", err.Error())
	}

	err = m.ProfanityFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.ProfanityFlag. %s", err.Error())
	}

	err = m.CharacterSet.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.CharacterSet. %s", err.Error())
	}

	err = m.MiiData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.MiiData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Mii
func (m Mii) Copy() types.RVType {
	copied := NewMii()

	copied.StructureVersion = m.StructureVersion
	copied.Data = m.Data.Copy().(types.Data)
	copied.Name = m.Name.Copy().(types.String)
	copied.ProfanityFlag = m.ProfanityFlag.Copy().(types.Bool)
	copied.CharacterSet = m.CharacterSet.Copy().(types.UInt8)
	copied.MiiData = m.MiiData.Copy().(types.Buffer)

	return copied
}

// Equals checks if the given Mii contains the same data as the current Mii
func (m Mii) Equals(o types.RVType) bool {
	if _, ok := o.(*Mii); !ok {
		return false
	}

	other := o.(*Mii)

	if m.StructureVersion != other.StructureVersion {
		return false
	}

	if !m.Data.Equals(other.Data) {
		return false
	}

	if !m.Name.Equals(other.Name) {
		return false
	}

	if !m.ProfanityFlag.Equals(other.ProfanityFlag) {
		return false
	}

	if !m.CharacterSet.Equals(other.CharacterSet) {
		return false
	}

	return m.MiiData.Equals(other.MiiData)
}

// CopyRef copies the current value of the Mii
// and returns a pointer to the new copy
func (m Mii) CopyRef() types.RVTypePtr {
	copied := m.Copy().(Mii)
	return &copied
}

// Deref takes a pointer to the Mii
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (m *Mii) Deref() types.RVType {
	return *m
}

// String returns the string representation of the Mii
func (m Mii) String() string {
	return m.FormatToString(0)
}

// FormatToString pretty-prints the Mii using the provided indentation level
func (m Mii) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Mii{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, m.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sName: %s,\n", indentationValues, m.Name))
	b.WriteString(fmt.Sprintf("%sProfanityFlag: %s,\n", indentationValues, m.ProfanityFlag))
	b.WriteString(fmt.Sprintf("%sCharacterSet: %s,\n", indentationValues, m.CharacterSet))
	b.WriteString(fmt.Sprintf("%sMiiData: %s,\n", indentationValues, m.MiiData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMii returns a new Mii
func NewMii() Mii {
	return Mii{
		Data:          types.NewData(),
		Name:          types.NewString(""),
		ProfanityFlag: types.NewBool(false),
		CharacterSet:  types.NewUInt8(0),
		MiiData:       types.NewBuffer(nil),
	}

}
