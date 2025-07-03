// Package types implements all the types used by the legacy Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RankingScore is a type within the Ranking protocol
type RankingScore struct {
	types.Structure
	Category types.UInt32
	Score    types.List[types.UInt32]
	Unknown1 types.UInt8
	Unknown2 types.UInt32
}

// WriteTo writes the RankingScore to the given writable
func (rs RankingScore) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rs.Category.WriteTo(contentWritable)
	rs.Score.WriteTo(contentWritable)
	rs.Unknown1.WriteTo(contentWritable)
	rs.Unknown2.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rs.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RankingScore from the given readable
func (rs *RankingScore) ExtractFrom(readable types.Readable) error {
	var err error

	err = rs.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScore header. %s", err.Error())
	}

	err = rs.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScore.Category. %s", err.Error())
	}

	err = rs.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScore.Score. %s", err.Error())
	}

	err = rs.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScore.Unknown1. %s", err.Error())
	}

	err = rs.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScore.Unknown2. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RankingScore
func (rs RankingScore) Copy() types.RVType {
	copied := NewRankingScore()

	copied.StructureVersion = rs.StructureVersion
	copied.Category = rs.Category.Copy().(types.UInt32)
	copied.Score = rs.Score.Copy().(types.List[types.UInt32])
	copied.Unknown1 = rs.Unknown1.Copy().(types.UInt8)
	copied.Unknown2 = rs.Unknown2.Copy().(types.UInt32)

	return copied
}

// Equals checks if the given RankingScore contains the same data as the current RankingScore
func (rs RankingScore) Equals(o types.RVType) bool {
	if _, ok := o.(RankingScore); !ok {
		return false
	}

	other := o.(RankingScore)

	if rs.StructureVersion != other.StructureVersion {
		return false
	}

	if !rs.Category.Equals(other.Category) {
		return false
	}

	if !rs.Score.Equals(other.Score) {
		return false
	}

	if !rs.Unknown1.Equals(other.Unknown1) {
		return false
	}

	return rs.Unknown2.Equals(other.Unknown2)
}

// CopyRef copies the current value of the RankingScore
// and returns a pointer to the new copy
func (rs RankingScore) CopyRef() types.RVTypePtr {
	copied := rs.Copy().(RankingScore)
	return &copied
}

// Deref takes a pointer to the RankingScore
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rs *RankingScore) Deref() types.RVType {
	return *rs
}

// String returns the string representation of the RankingScore
func (rs RankingScore) String() string {
	return rs.FormatToString(0)
}

// FormatToString pretty-prints the RankingScore using the provided indentation level
func (rs RankingScore) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingScore{\n")
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, rs.Category))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, rs.Score))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, rs.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, rs.Unknown2))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingScore returns a new RankingScore
func NewRankingScore() RankingScore {
	return RankingScore{
		Category: types.NewUInt32(0),
		Score:    types.NewList[types.UInt32](),
		Unknown1: types.NewUInt8(0),
		Unknown2: types.NewUInt32(0),
	}

}
