// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Ranking2ChartInfoInput holds data for the Ranking 2  protocol
type Ranking2ChartInfoInput struct {
	nex.Structure
	ChartIndex         uint32
	NumSeasonsToGoBack uint8
}

// ExtractFromStream extracts a Ranking2ChartInfoInput structure from a stream
func (ranking2ChartInfoInput *Ranking2ChartInfoInput) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	ranking2ChartInfoInput.ChartIndex, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfoInput.ChartIndex from stream. %s", err.Error())
	}

	ranking2ChartInfoInput.NumSeasonsToGoBack, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfoInput.NumSeasonsToGoBack from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the Ranking2ChartInfoInput and returns a byte array
func (ranking2ChartInfoInput *Ranking2ChartInfoInput) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(ranking2ChartInfoInput.ChartIndex)
	stream.WriteUInt8(ranking2ChartInfoInput.NumSeasonsToGoBack)

	return stream.Bytes()
}

// Copy returns a new copied instance of Ranking2ChartInfoInput
func (ranking2ChartInfoInput *Ranking2ChartInfoInput) Copy() nex.StructureInterface {
	copied := NewRanking2ChartInfoInput()

	copied.SetStructureVersion(ranking2ChartInfoInput.StructureVersion())

	copied.ChartIndex = ranking2ChartInfoInput.ChartIndex
	copied.NumSeasonsToGoBack = ranking2ChartInfoInput.NumSeasonsToGoBack
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2ChartInfoInput *Ranking2ChartInfoInput) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Ranking2ChartInfoInput)

	if ranking2ChartInfoInput.StructureVersion() != other.StructureVersion() {
		return false
	}

	if ranking2ChartInfoInput.ChartIndex != other.ChartIndex {
		return false
	}

	if ranking2ChartInfoInput.NumSeasonsToGoBack != other.NumSeasonsToGoBack {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (ranking2ChartInfoInput *Ranking2ChartInfoInput) String() string {
	return ranking2ChartInfoInput.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2ChartInfoInput *Ranking2ChartInfoInput) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2ChartInfoInput{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, ranking2ChartInfoInput.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sChartIndex: %d,\n", indentationValues, ranking2ChartInfoInput.ChartIndex))
	b.WriteString(fmt.Sprintf("%sNumSeasonsToGoBack: %d,\n", indentationValues, ranking2ChartInfoInput.NumSeasonsToGoBack))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2ChartInfoInput returns a new Ranking2ChartInfoInput
func NewRanking2ChartInfoInput() *Ranking2ChartInfoInput {
	return &Ranking2ChartInfoInput{}
}
