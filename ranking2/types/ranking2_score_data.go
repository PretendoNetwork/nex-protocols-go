// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Ranking2ScoreData holds data for the Ranking 2  protocol
type Ranking2ScoreData struct {
	types.Structure
	Misc     *types.PrimitiveU64
	Category *types.PrimitiveU32
	Score    *types.PrimitiveU32
}

// ExtractFrom extracts the Ranking2ScoreData from the given readable
func (ranking2ScoreData *Ranking2ScoreData) ExtractFrom(readable types.Readable) error {
	var err error

	if err = ranking2ScoreData.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Ranking2ScoreData header. %s", err.Error())
	}

	err = ranking2ScoreData.Misc.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ScoreData.Misc from stream. %s", err.Error())
	}

	err = ranking2ScoreData.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ScoreData.Category from stream. %s", err.Error())
	}

	err = ranking2ScoreData.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ScoreData.Score from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Ranking2ScoreData to the given writable
func (ranking2ScoreData *Ranking2ScoreData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ranking2ScoreData.Misc.WriteTo(contentWritable)
	ranking2ScoreData.Category.WriteTo(contentWritable)
	ranking2ScoreData.Score.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ranking2ScoreData.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Ranking2ScoreData
func (ranking2ScoreData *Ranking2ScoreData) Copy() types.RVType {
	copied := NewRanking2ScoreData()

	copied.StructureVersion = ranking2ScoreData.StructureVersion

	copied.Misc = ranking2ScoreData.Misc
	copied.Category = ranking2ScoreData.Category
	copied.Score = ranking2ScoreData.Score
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2ScoreData *Ranking2ScoreData) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2ScoreData); !ok {
		return false
	}

	other := o.(*Ranking2ScoreData)

	if ranking2ScoreData.StructureVersion != other.StructureVersion {
		return false
	}

	if !ranking2ScoreData.Misc.Equals(other.Misc) {
		return false
	}

	if !ranking2ScoreData.Category.Equals(other.Category) {
		return false
	}

	if !ranking2ScoreData.Score.Equals(other.Score) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (ranking2ScoreData *Ranking2ScoreData) String() string {
	return ranking2ScoreData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2ScoreData *Ranking2ScoreData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2ScoreData{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, ranking2ScoreData.StructureVersion))
	b.WriteString(fmt.Sprintf("%sMisc: %d,\n", indentationValues, ranking2ScoreData.Misc))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, ranking2ScoreData.Category))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, ranking2ScoreData.Score))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2ScoreData returns a new Ranking2ScoreData
func NewRanking2ScoreData() *Ranking2ScoreData {
	return &Ranking2ScoreData{}
}
