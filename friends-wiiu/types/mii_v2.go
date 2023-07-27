// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// MiiV2 contains data about a Mii
type MiiV2 struct {
	nex.Structure
	*nex.Data
	Name     string
	Unknown1 uint8
	Unknown2 uint8
	MiiData  []byte
	Datetime *nex.DateTime
}

// Bytes encodes the MiiV2 and returns a byte array
func (mii *MiiV2) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(mii.Name)
	stream.WriteUInt8(mii.Unknown1)
	stream.WriteUInt8(mii.Unknown2)
	stream.WriteBuffer(mii.MiiData)
	stream.WriteDateTime(mii.Datetime)

	return stream.Bytes()
}

// ExtractFromStream extracts a MiiV2 structure from a stream
func (mii *MiiV2) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	mii.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Name. %s", err.Error())
	}

	mii.Unknown1, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Unknown1. %s", err.Error())
	}

	mii.Unknown2, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Unknown2. %s", err.Error())
	}

	mii.MiiData, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.MiiData. %s", err.Error())
	}

	mii.Datetime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Datetime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MiiV2
func (mii *MiiV2) Copy() nex.StructureInterface {
	copied := NewMiiV2()

	copied.Data = mii.ParentType().Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

	copied.Name = mii.Name
	copied.Unknown1 = mii.Unknown1
	copied.Unknown2 = mii.Unknown2
	copied.MiiData = make([]byte, len(mii.MiiData))

	copy(copied.MiiData, mii.MiiData)

	copied.Datetime = mii.Datetime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (mii *MiiV2) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MiiV2)

	if !mii.ParentType().Equals(other.ParentType()) {
		return false
	}

	if mii.Name != other.Name {
		return false
	}

	if mii.Unknown1 != other.Unknown1 {
		return false
	}

	if mii.Unknown2 != other.Unknown2 {
		return false
	}

	if !bytes.Equal(mii.MiiData, other.MiiData) {
		return false
	}

	if !mii.Datetime.Equals(other.Datetime) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (mii *MiiV2) String() string {
	return mii.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (mii *MiiV2) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MiiV2{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, mii.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, mii.Name))
	b.WriteString(fmt.Sprintf("%sUnknown1: %d,\n", indentationValues, mii.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %d,\n", indentationValues, mii.Unknown2))
	b.WriteString(fmt.Sprintf("%sMiiData: %x,\n", indentationValues, mii.MiiData))

	if mii.Datetime != nil {
		b.WriteString(fmt.Sprintf("%sDatetime: %s\n", indentationValues, mii.Datetime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sDatetime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMiiV2 returns a new MiiV2
func NewMiiV2() *MiiV2 {
	return &MiiV2{}
}
