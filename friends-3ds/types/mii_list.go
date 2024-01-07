// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// MiiList is a data structure used by the Friends 3DS protocol to hold information about a MiiList
type MiiList struct {
	types.Structure
	*types.Data
	Unknown1    string
	Unknown2    *types.PrimitiveBool
	Unknown3    *types.PrimitiveU8
	MiiDataList [][]byte
}

// WriteTo writes the MiiList to the given writable
func (miiList *MiiList) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	miiList.Unknown1.WriteTo(contentWritable)
	miiList.Unknown2.WriteTo(contentWritable)
	miiList.Unknown3.WriteTo(contentWritable)
	stream.WriteListBuffer(miiList.MiiDataList)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFromStream extracts a MiiList from a stream
func (miiList *MiiList) ExtractFrom(readable types.Readable) error {
	var err error

	if err = miiList.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read MiiList header. %s", err.Error())
	}

	err = miiList.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiList.Unknown1. %s", err.Error())
	}

	err = miiList.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiList.Unknown2. %s", err.Error())
	}

	err = miiList.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MiiList.Unknown3. %s", err.Error())
	}

	miiList.MiiDataList, err = stream.ReadListBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiList.MiiDataList. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MiiList
func (miiList *MiiList) Copy() types.RVType {
	copied := NewMiiList()

	copied.StructureVersion = miiList.StructureVersion

	copied.Data = miiList.Data.Copy().(*types.Data)

	copied.Unknown1 = miiList.Unknown1
	copied.Unknown2 = miiList.Unknown2
	copied.Unknown3 = miiList.Unknown3
	copied.MiiDataList = make([][]byte, len(miiList.MiiDataList))

	for i := 0; i < len(miiList.MiiDataList); i++ {
		copied.MiiDataList[i] = make([]byte, len(miiList.MiiDataList[i]))

		copy(copied.MiiDataList[i], miiList.MiiDataList[i])
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (miiList *MiiList) Equals(o types.RVType) bool {
	if _, ok := o.(*MiiList); !ok {
		return false
	}

	other := o.(*MiiList)

	if miiList.StructureVersion != other.StructureVersion {
		return false
	}

	if !miiList.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !miiList.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !miiList.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !miiList.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if len(miiList.MiiDataList) != len(other.MiiDataList) {
		return false
	}

	for i := 0; i < len(miiList.MiiDataList); i++ {
		if !miiList.MiiDataList[i].Equals(other.MiiDataList[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (miiList *MiiList) String() string {
	return miiList.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (miiList *MiiList) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MiiList{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, miiList.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUnknown1: %q,\n", indentationValues, miiList.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %t,\n", indentationValues, miiList.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %d,\n", indentationValues, miiList.Unknown3))
	b.WriteString(fmt.Sprintf("%sMiiDataList: %v\n", indentationValues, miiList.MiiDataList)) // TODO - Make this a nicer looking log
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMiiList returns a new MiiList
func NewMiiList() *MiiList {
	return &MiiList{}
}
