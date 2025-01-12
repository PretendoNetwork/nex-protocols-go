// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RankingScoreData is a type within the Ranking protocol
type RankingScoreData struct {
	types.Structure
	Category   types.UInt32
	Score      types.UInt32
	OrderBy    types.UInt8
	UpdateMode types.UInt8
	Groups     types.Buffer
	Param      types.UInt64
}

// WriteTo writes the RankingScoreData to the given writable
func (rsd RankingScoreData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rsd.Category.WriteTo(contentWritable)
	rsd.Score.WriteTo(contentWritable)
	rsd.OrderBy.WriteTo(contentWritable)
	rsd.UpdateMode.WriteTo(contentWritable)
	rsd.Groups.WriteTo(contentWritable)
	rsd.Param.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rsd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RankingScoreData from the given readable
func (rsd *RankingScoreData) ExtractFrom(readable types.Readable) error {
	var err error

	err = rsd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData header. %s", err.Error())
	}

	err = rsd.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Category. %s", err.Error())
	}

	err = rsd.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Score. %s", err.Error())
	}

	err = rsd.OrderBy.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.OrderBy. %s", err.Error())
	}

	err = rsd.UpdateMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.UpdateMode. %s", err.Error())
	}

	err = rsd.Groups.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Groups. %s", err.Error())
	}

	err = rsd.Param.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Param. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RankingScoreData
func (rsd RankingScoreData) Copy() types.RVType {
	copied := NewRankingScoreData()

	copied.StructureVersion = rsd.StructureVersion
	copied.Category = rsd.Category.Copy().(types.UInt32)
	copied.Score = rsd.Score.Copy().(types.UInt32)
	copied.OrderBy = rsd.OrderBy.Copy().(types.UInt8)
	copied.UpdateMode = rsd.UpdateMode.Copy().(types.UInt8)
	copied.Groups = rsd.Groups.Copy().(types.Buffer)
	copied.Param = rsd.Param.Copy().(types.UInt64)

	return copied
}

// Equals checks if the given RankingScoreData contains the same data as the current RankingScoreData
func (rsd RankingScoreData) Equals(o types.RVType) bool {
	if _, ok := o.(RankingScoreData); !ok {
		return false
	}

	other := o.(RankingScoreData)

	if rsd.StructureVersion != other.StructureVersion {
		return false
	}

	if !rsd.Category.Equals(other.Category) {
		return false
	}

	if !rsd.Score.Equals(other.Score) {
		return false
	}

	if !rsd.OrderBy.Equals(other.OrderBy) {
		return false
	}

	if !rsd.UpdateMode.Equals(other.UpdateMode) {
		return false
	}

	if !rsd.Groups.Equals(other.Groups) {
		return false
	}

	return rsd.Param.Equals(other.Param)
}

// CopyRef copies the current value of the RankingScoreData
// and returns a pointer to the new copy
func (rsd RankingScoreData) CopyRef() types.RVTypePtr {
	copied := rsd.Copy().(RankingScoreData)
	return &copied
}

// Deref takes a pointer to the RankingScoreData
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rsd *RankingScoreData) Deref() types.RVType {
	return *rsd
}

// String returns the string representation of the RankingScoreData
func (rsd RankingScoreData) String() string {
	return rsd.FormatToString(0)
}

// FormatToString pretty-prints the RankingScoreData using the provided indentation level
func (rsd RankingScoreData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingScoreData{\n")
	b.WriteString(fmt.Sprintf("%sCategory: %s,\n", indentationValues, rsd.Category))
	b.WriteString(fmt.Sprintf("%sScore: %s,\n", indentationValues, rsd.Score))
	b.WriteString(fmt.Sprintf("%sOrderBy: %s,\n", indentationValues, rsd.OrderBy))
	b.WriteString(fmt.Sprintf("%sUpdateMode: %s,\n", indentationValues, rsd.UpdateMode))
	b.WriteString(fmt.Sprintf("%sGroups: %s,\n", indentationValues, rsd.Groups))
	b.WriteString(fmt.Sprintf("%sParam: %s,\n", indentationValues, rsd.Param))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingScoreData returns a new RankingScoreData
func NewRankingScoreData() RankingScoreData {
	return RankingScoreData{
		Category:   types.NewUInt32(0),
		Score:      types.NewUInt32(0),
		OrderBy:    types.NewUInt8(0),
		UpdateMode: types.NewUInt8(0),
		Groups:     types.NewBuffer(nil),
		Param:      types.NewUInt64(0),
	}

}
