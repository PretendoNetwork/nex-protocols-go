// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// MiiV2 contains data about a Mii
type MiiV2 struct {
	types.Structure
	*types.Data
	Name     string
	Unknown1 *types.PrimitiveU8
	Unknown2 *types.PrimitiveU8
	MiiData  []byte
	Datetime *types.DateTime
}

// WriteTo writes the MiiV2 to the given writable
func (mii *MiiV2) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	mii.Name.WriteTo(contentWritable)
	mii.Unknown1.WriteTo(contentWritable)
	mii.Unknown2.WriteTo(contentWritable)
	stream.WriteBuffer(mii.MiiData)
	mii.Datetime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mii.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MiiV2 from the given readable
func (mii *MiiV2) ExtractFrom(readable types.Readable) error {
	var err error

	if err = mii.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read MiiV2 header. %s", err.Error())
	}

	err = mii.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Name. %s", err.Error())
	}

	err = mii.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Unknown1. %s", err.Error())
	}

	err = mii.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Unknown2. %s", err.Error())
	}

	mii.MiiData, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.MiiData. %s", err.Error())
	}

	err = mii.Datetime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiV2.Datetime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MiiV2
func (mii *MiiV2) Copy() types.RVType {
	copied := NewMiiV2()

	copied.StructureVersion = mii.StructureVersion

	copied.Data = mii.Data.Copy().(*types.Data)

	copied.Name = mii.Name
	copied.Unknown1 = mii.Unknown1
	copied.Unknown2 = mii.Unknown2
	copied.MiiData = make([]byte, len(mii.MiiData))

	copy(copied.MiiData, mii.MiiData)

	copied.Datetime = mii.Datetime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (mii *MiiV2) Equals(o types.RVType) bool {
	if _, ok := o.(*MiiV2); !ok {
		return false
	}

	other := o.(*MiiV2)

	if mii.StructureVersion != other.StructureVersion {
		return false
	}

	if !mii.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !mii.Name.Equals(other.Name) {
		return false
	}

	if !mii.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !mii.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !mii.MiiData.Equals(other.MiiData) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, mii.StructureVersion))
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
