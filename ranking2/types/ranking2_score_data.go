// Package types implements all the types used by the Ranking2 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2ScoreData is a type within the Ranking2 protocol
type Ranking2ScoreData struct {
	types.Structure
	Misc     types.UInt64
	Category types.UInt32
	Score    types.UInt32
}

// WriteTo writes the Ranking2ScoreData to the given writable
func (rsd Ranking2ScoreData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rsd.Misc.WriteTo(contentWritable)
	rsd.Category.WriteTo(contentWritable)
	rsd.Score.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rsd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Ranking2ScoreData from the given readable
func (rsd *Ranking2ScoreData) ExtractFrom(readable types.Readable) error {
	var err error

	err = rsd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ScoreData header. %s", err.Error())
	}

	err = rsd.Misc.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ScoreData.Misc. %s", err.Error())
	}

	err = rsd.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ScoreData.Category. %s", err.Error())
	}

	err = rsd.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ScoreData.Score. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Ranking2ScoreData
func (rsd Ranking2ScoreData) Copy() types.RVType {
	copied := NewRanking2ScoreData()

	copied.StructureVersion = rsd.StructureVersion
	copied.Misc = rsd.Misc.Copy().(types.UInt64)
	copied.Category = rsd.Category.Copy().(types.UInt32)
	copied.Score = rsd.Score.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given Ranking2ScoreData contains the same data as the current Ranking2ScoreData
func (rsd Ranking2ScoreData) Equals(o types.RVType) bool {
	if _, ok := o.(Ranking2ScoreData); !ok {
		return false
	}

	other := o.(Ranking2ScoreData)

	if rsd.StructureVersion != other.StructureVersion {
		return false
	}

	if !rsd.Misc.Equals(other.Misc) {
		return false
	}

	if !rsd.Category.Equals(other.Category) {
		return false
	}

	return rsd.Score.Equals(other.Score)
}

// CopyRef copies the current value of the Ranking2ScoreData
// and returns a pointer to the new copy
func (rsd Ranking2ScoreData) CopyRef() types.RVTypePtr {
	copied := rsd.Copy().(Ranking2ScoreData)
	return &copied
}

// Deref takes a pointer to the Ranking2ScoreData
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rsd *Ranking2ScoreData) Deref() types.RVType {
	return *rsd
}

// String returns the string representation of the Ranking2ScoreData
func (rsd Ranking2ScoreData) String() string {
	return rsd.FormatToString(0)
}

// FormatToString pretty-prints the Ranking2ScoreData using the provided indentation level
func (rsd Ranking2ScoreData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2ScoreData{\n")
	b.WriteString(fmt.Sprintf("%sMisc: %s,\n", indentationValues, rsd.Misc))
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, rsd.Category))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, rsd.Score))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2ScoreData returns a new Ranking2ScoreData
func NewRanking2ScoreData() Ranking2ScoreData {
	return Ranking2ScoreData{
		Misc:     types.NewUInt64(0),
		Category: types.NewUInt32(0),
		Score:    types.NewUInt32(0),
	}

}
