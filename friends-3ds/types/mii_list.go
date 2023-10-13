// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// MiiList is a data structure used by the Friends 3DS protocol to hold information about a MiiList
type MiiList struct {
	nex.Structure
	*nex.Data
	Unknown1    string
	Unknown2    bool
	Unknown3    uint8
	MiiDataList [][]byte
}

// Bytes encodes the MiiList and returns a byte array
func (miiList *MiiList) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(miiList.Unknown1)
	stream.WriteBool(miiList.Unknown2)
	stream.WriteUInt8(miiList.Unknown3)
	stream.WriteListBuffer(miiList.MiiDataList)

	return stream.Bytes()
}

// ExtractFromStream extracts a MiiList from a stream
func (miiList *MiiList) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	miiList.Unknown1, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiList.Unknown1. %s", err.Error())
	}

	miiList.Unknown2, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract MiiList.Unknown2. %s", err.Error())
	}

	miiList.Unknown3, err = stream.ReadUInt8()
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
func (miiList *MiiList) Copy() nex.StructureInterface {
	copied := NewMiiList()

	copied.SetStructureVersion(miiList.StructureVersion())

	if miiList.ParentType() != nil {
		copied.Data = miiList.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

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
func (miiList *MiiList) Equals(structure nex.StructureInterface) bool {
	other := structure.(*MiiList)

	if miiList.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !miiList.ParentType().Equals(other.ParentType()) {
		return false
	}

	if miiList.Unknown1 != other.Unknown1 {
		return false
	}

	if miiList.Unknown2 != other.Unknown2 {
		return false
	}

	if miiList.Unknown3 != other.Unknown3 {
		return false
	}

	if len(miiList.MiiDataList) != len(other.MiiDataList) {
		return false
	}

	for i := 0; i < len(miiList.MiiDataList); i++ {
		if !bytes.Equal(miiList.MiiDataList[i], other.MiiDataList[i]) {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, miiList.StructureVersion()))
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
