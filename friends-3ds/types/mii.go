// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Mii is a data structure used by the Friends 3DS protocol to hold information about a Mii
type Mii struct {
	nex.Structure
	*nex.Data
	Name     string
	Unknown2 bool
	Unknown3 uint8
	MiiData  []byte
}

// Bytes encodes the Mii and returns a byte array
func (mii *Mii) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(mii.Name)
	stream.WriteBool(mii.Unknown2)
	stream.WriteUInt8(mii.Unknown3)
	stream.WriteBuffer(mii.MiiData)

	return stream.Bytes()
}

// ExtractFromStream extracts a Mii from a stream
func (mii *Mii) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	mii.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.Name. %s", err.Error())
	}

	mii.Unknown2, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract Mii.Unknown2. %s", err.Error())
	}

	mii.Unknown3, err = stream.ReadUInt8()
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
func (mii *Mii) Copy() nex.StructureInterface {
	copied := NewMii()

	copied.SetStructureVersion(mii.StructureVersion())

	if mii.ParentType() != nil {
		copied.Data = mii.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

	copied.SetParentType(copied.Data)

	copied.Name = mii.Name
	copied.Unknown2 = mii.Unknown2
	copied.Unknown3 = mii.Unknown3
	copied.MiiData = make([]byte, len(mii.MiiData))

	copy(copied.MiiData, mii.MiiData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (mii *Mii) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Mii)

	if mii.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !mii.ParentType().Equals(other.ParentType()) {
		return false
	}

	if mii.Name != other.Name {
		return false
	}

	if mii.Unknown2 != other.Unknown2 {
		return false
	}

	if mii.Unknown3 != other.Unknown3 {
		return false
	}

	if !bytes.Equal(mii.MiiData, other.MiiData) {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, mii.StructureVersion()))
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
