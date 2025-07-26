// Package types implements all the types used by the legacy Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RankingScoreWithLimit is a type within the Ranking protocol
type RankingScoreWithLimit struct {
	types.Structure
	Category types.UInt32
	Score    types.List[types.UInt32]
	Unknown1 types.UInt8
	Unknown2 types.UInt32
	Limit    types.UInt16
}

// WriteTo writes the RankingScoreWithLimit to the given writable
func (rswl RankingScoreWithLimit) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rswl.Category.WriteTo(contentWritable)
	rswl.Score.WriteTo(contentWritable)
	rswl.Unknown1.WriteTo(contentWritable)
	rswl.Unknown2.WriteTo(contentWritable)
	rswl.Limit.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rswl.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RankingScoreWithLimit from the given readable
func (rswl *RankingScoreWithLimit) ExtractFrom(readable types.Readable) error {
	var err error

	err = rswl.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreWithLimit header. %s", err.Error())
	}

	err = rswl.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreWithLimit.Category. %s", err.Error())
	}

	err = rswl.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreWithLimit.Score. %s", err.Error())
	}

	err = rswl.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreWithLimit.Unknown1. %s", err.Error())
	}

	err = rswl.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreWithLimit.Unknown2. %s", err.Error())
	}

	err = rswl.Limit.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreWithLimit.Limit. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RankingScoreWithLimit
func (rswl RankingScoreWithLimit) Copy() types.RVType {
	copied := NewRankingScoreWithLimit()

	copied.StructureVersion = rswl.StructureVersion
	copied.Category = rswl.Category.Copy().(types.UInt32)
	copied.Score = rswl.Score.Copy().(types.List[types.UInt32])
	copied.Unknown1 = rswl.Unknown1.Copy().(types.UInt8)
	copied.Unknown2 = rswl.Unknown2.Copy().(types.UInt32)
	copied.Limit = rswl.Limit.Copy().(types.UInt16)

	return copied
}

// Equals checks if the given RankingScoreWithLimit contains the same data as the current RankingScoreWithLimit
func (rswl RankingScoreWithLimit) Equals(o types.RVType) bool {
	if _, ok := o.(RankingScoreWithLimit); !ok {
		return false
	}

	other := o.(RankingScoreWithLimit)

	if rswl.StructureVersion != other.StructureVersion {
		return false
	}

	if !rswl.Category.Equals(other.Category) {
		return false
	}

	if !rswl.Score.Equals(other.Score) {
		return false
	}

	if !rswl.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !rswl.Unknown2.Equals(other.Unknown2) {
		return false
	}

	return rswl.Limit.Equals(other.Limit)
}

// CopyRef copies the current value of the RankingScoreWithLimit
// and returns a pointer to the new copy
func (rswl RankingScoreWithLimit) CopyRef() types.RVTypePtr {
	copied := rswl.Copy().(RankingScoreWithLimit)
	return &copied
}

// Deref takes a pointer to the RankingScoreWithLimit
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rswl *RankingScoreWithLimit) Deref() types.RVType {
	return *rswl
}

// String returns the string representation of the RankingScoreWithLimit
func (rswl RankingScoreWithLimit) String() string {
	return rswl.FormatToString(0)
}

// FormatToString pretty-prints the RankingScoreWithLimit using the provided indentation level
func (rswl RankingScoreWithLimit) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingScoreWithLimit{\n")
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, rswl.Category))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, rswl.Score))
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, rswl.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, rswl.Unknown2))
	b.WriteString(fmt.Sprintf("%sLimit: %s,\n", indentationValues, rswl.Limit))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingScoreWithLimit returns a new RankingScoreWithLimit
func NewRankingScoreWithLimit() RankingScoreWithLimit {
	return RankingScoreWithLimit{
		Category: types.NewUInt32(0),
		Score:    types.NewList[types.UInt32](),
		Unknown1: types.NewUInt8(0),
		Unknown2: types.NewUInt32(0),
		Limit:    types.NewUInt16(0),
	}

}
