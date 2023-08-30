// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Ranking2CommonData holds data for the Ranking 2  protocol
type Ranking2CommonData struct {
	nex.Structure
	UserName   string
	Mii        []byte
	BinaryData []byte
}

// ExtractFromStream extracts a Ranking2CommonData structure from a stream
func (ranking2CommonData *Ranking2CommonData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	ranking2CommonData.UserName, err = stream.ReadString()
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

// Bytes encodes the Ranking2CommonData and returns a byte array
func (ranking2CommonData *Ranking2CommonData) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(ranking2CommonData.UserName)
	stream.WriteQBuffer(ranking2CommonData.Mii)
	stream.WriteQBuffer(ranking2CommonData.BinaryData)

	return stream.Bytes()
}

// Copy returns a new copied instance of Ranking2CommonData
func (ranking2CommonData *Ranking2CommonData) Copy() nex.StructureInterface {
	copied := NewRanking2CommonData()

	copied.SetStructureVersion(ranking2CommonData.StructureVersion())

	copied.UserName = ranking2CommonData.UserName
	copied.Mii = ranking2CommonData.Mii
	copied.BinaryData = ranking2CommonData.BinaryData
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2CommonData *Ranking2CommonData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Ranking2CommonData)

	if ranking2CommonData.StructureVersion() != other.StructureVersion() {
		return false
	}

	if ranking2CommonData.UserName != other.UserName {
		return false
	}

	if !bytes.Equal(ranking2CommonData.Mii, other.Mii) {
		return false
	}

	if !bytes.Equal(ranking2CommonData.BinaryData, other.BinaryData) {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, ranking2CommonData.StructureVersion()))
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
