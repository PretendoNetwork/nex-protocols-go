// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Ranking2CommonData holds data for the Ranking 2  protocol
type Ranking2CommonData struct {
	types.Structure
	UserName   string
	Mii        []byte
	BinaryData []byte
}

// ExtractFrom extracts the Ranking2CommonData from the given readable
func (ranking2CommonData *Ranking2CommonData) ExtractFrom(readable types.Readable) error {
	var err error

	if err = ranking2CommonData.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Ranking2CommonData header. %s", err.Error())
	}

	err = ranking2CommonData.UserName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CommonData.UserName from stream. %s", err.Error())
	}

	ranking2CommonData.Mii, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CommonData.Mii from stream. %s", err.Error())
	}

	ranking2CommonData.BinaryData, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CommonData.BinaryData from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Ranking2CommonData to the given writable
func (ranking2CommonData *Ranking2CommonData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ranking2CommonData.UserName.WriteTo(contentWritable)
	stream.WriteQBuffer(ranking2CommonData.Mii)
	stream.WriteQBuffer(ranking2CommonData.BinaryData)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Ranking2CommonData
func (ranking2CommonData *Ranking2CommonData) Copy() types.RVType {
	copied := NewRanking2CommonData()

	copied.StructureVersion = ranking2CommonData.StructureVersion

	copied.UserName = ranking2CommonData.UserName
	copied.Mii = ranking2CommonData.Mii
	copied.BinaryData = ranking2CommonData.BinaryData
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2CommonData *Ranking2CommonData) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2CommonData); !ok {
		return false
	}

	other := o.(*Ranking2CommonData)

	if ranking2CommonData.StructureVersion != other.StructureVersion {
		return false
	}

	if !ranking2CommonData.UserName.Equals(other.UserName) {
		return false
	}

	if !ranking2CommonData.Mii.Equals(other.Mii) {
		return false
	}

	if !ranking2CommonData.BinaryData.Equals(other.BinaryData) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (ranking2CommonData *Ranking2CommonData) String() string {
	return ranking2CommonData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2CommonData *Ranking2CommonData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2CommonData{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, ranking2CommonData.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUserName: %q,\n", indentationValues, ranking2CommonData.UserName))
	b.WriteString(fmt.Sprintf("%sMii: %x,\n", indentationValues, ranking2CommonData.Mii))
	b.WriteString(fmt.Sprintf("%sBinaryData: %x,\n", indentationValues, ranking2CommonData.BinaryData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2CommonData returns a new Ranking2CommonData
func NewRanking2CommonData() *Ranking2CommonData {
	return &Ranking2CommonData{}
}
