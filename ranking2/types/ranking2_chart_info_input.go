// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Ranking2ChartInfoInput holds data for the Ranking 2  protocol
type Ranking2ChartInfoInput struct {
	types.Structure
	ChartIndex         *types.PrimitiveU32
	NumSeasonsToGoBack *types.PrimitiveU8
}

// ExtractFrom extracts the Ranking2ChartInfoInput from the given readable
func (ranking2ChartInfoInput *Ranking2ChartInfoInput) ExtractFrom(readable types.Readable) error {
	var err error

	if err = ranking2ChartInfoInput.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Ranking2ChartInfoInput header. %s", err.Error())
	}

	err = ranking2ChartInfoInput.ChartIndex.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfoInput.ChartIndex from stream. %s", err.Error())
	}

	err = ranking2ChartInfoInput.NumSeasonsToGoBack.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfoInput.NumSeasonsToGoBack from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Ranking2ChartInfoInput to the given writable
func (ranking2ChartInfoInput *Ranking2ChartInfoInput) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ranking2ChartInfoInput.ChartIndex.WriteTo(contentWritable)
	ranking2ChartInfoInput.NumSeasonsToGoBack.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ranking2ChartInfoInput.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Ranking2ChartInfoInput
func (ranking2ChartInfoInput *Ranking2ChartInfoInput) Copy() types.RVType {
	copied := NewRanking2ChartInfoInput()

	copied.StructureVersion = ranking2ChartInfoInput.StructureVersion

	copied.ChartIndex = ranking2ChartInfoInput.ChartIndex
	copied.NumSeasonsToGoBack = ranking2ChartInfoInput.NumSeasonsToGoBack
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2ChartInfoInput *Ranking2ChartInfoInput) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2ChartInfoInput); !ok {
		return false
	}

	other := o.(*Ranking2ChartInfoInput)

	if ranking2ChartInfoInput.StructureVersion != other.StructureVersion {
		return false
	}

	if !ranking2ChartInfoInput.ChartIndex.Equals(other.ChartIndex) {
		return false
	}

	if !ranking2ChartInfoInput.NumSeasonsToGoBack.Equals(other.NumSeasonsToGoBack) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, ranking2ChartInfoInput.StructureVersion))
	b.WriteString(fmt.Sprintf("%sChartIndex: %d,\n", indentationValues, ranking2ChartInfoInput.ChartIndex))
	b.WriteString(fmt.Sprintf("%sNumSeasonsToGoBack: %d,\n", indentationValues, ranking2ChartInfoInput.NumSeasonsToGoBack))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2ChartInfoInput returns a new Ranking2ChartInfoInput
func NewRanking2ChartInfoInput() *Ranking2ChartInfoInput {
	return &Ranking2ChartInfoInput{}
}
