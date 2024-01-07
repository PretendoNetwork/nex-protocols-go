// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Mii is a data structure used by the Friends 3DS protocol to hold information about a Mii
type Mii struct {
	types.Structure
	*types.Data
	Name     string
	Unknown2 *types.PrimitiveBool
	Unknown3 *types.PrimitiveU8
	MiiData  []byte
}

// WriteTo writes the Mii to the given writable
func (mii *Mii) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	mii.Name.WriteTo(contentWritable)
	mii.Unknown2.WriteTo(contentWritable)
	mii.Unknown3.WriteTo(contentWritable)
	stream.WriteBuffer(mii.MiiData)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFromStream extracts a Mii from a stream
func (mii *Mii) ExtractFrom(readable types.Readable) error {
	var err error

	if err = mii.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Mii header. %s", err.Error())
	}

	err = mii.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.Name. %s", err.Error())
	}

	err = mii.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.Unknown2. %s", err.Error())
	}

	err = mii.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.Unknown3. %s", err.Error())
	}

	mii.MiiData, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.MiiData. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Mii
func (mii *Mii) Copy() types.RVType {
	copied := NewMii()

	copied.StructureVersion = mii.StructureVersion

	copied.Data = mii.Data.Copy().(*types.Data)

	copied.Name = mii.Name
	copied.Unknown2 = mii.Unknown2
	copied.Unknown3 = mii.Unknown3
	copied.MiiData = make([]byte, len(mii.MiiData))

	copy(copied.MiiData, mii.MiiData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (mii *Mii) Equals(o types.RVType) bool {
	if _, ok := o.(*Mii); !ok {
		return false
	}

	other := o.(*Mii)

	if mii.StructureVersion != other.StructureVersion {
		return false
	}

	if !mii.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !mii.Name.Equals(other.Name) {
		return false
	}

	if !mii.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !mii.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !mii.MiiData.Equals(other.MiiData) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (mii *Mii) String() string {
	return mii.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (mii *Mii) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Mii{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, mii.StructureVersion))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, mii.Name))
	b.WriteString(fmt.Sprintf("%sUnknown2: %t,\n", indentationValues, mii.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %d,\n", indentationValues, mii.Unknown3))
	b.WriteString(fmt.Sprintf("%sMiiData: %x\n", indentationValues, mii.MiiData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMii returns a new Mii
func NewMii() *Mii {
	return &Mii{}
}
