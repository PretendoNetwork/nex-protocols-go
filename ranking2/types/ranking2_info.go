// Package types implements all the types used by the Ranking2 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2Info is a type within the Ranking2 protocol
type Ranking2Info struct {
	types.Structure
	RankDataList types.List[Ranking2RankData]
	LowestRank   types.UInt32
	NumRankedIn  types.UInt32
	Season       types.Int32
}

// WriteTo writes the Ranking2Info to the given writable
func (ri Ranking2Info) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ri.RankDataList.WriteTo(contentWritable)
	ri.LowestRank.WriteTo(contentWritable)
	ri.NumRankedIn.WriteTo(contentWritable)
	ri.Season.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ri.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Ranking2Info from the given readable
func (ri *Ranking2Info) ExtractFrom(readable types.Readable) error {
	var err error

	err = ri.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info header. %s", err.Error())
	}

	err = ri.RankDataList.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info.RankDataList. %s", err.Error())
	}

	err = ri.LowestRank.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info.LowestRank. %s", err.Error())
	}

	err = ri.NumRankedIn.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info.NumRankedIn. %s", err.Error())
	}

	err = ri.Season.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info.Season. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Ranking2Info
func (ri Ranking2Info) Copy() types.RVType {
	copied := NewRanking2Info()

	copied.StructureVersion = ri.StructureVersion
	copied.RankDataList = ri.RankDataList.Copy().(types.List[Ranking2RankData])
	copied.LowestRank = ri.LowestRank.Copy().(types.UInt32)
	copied.NumRankedIn = ri.NumRankedIn.Copy().(types.UInt32)
	copied.Season = ri.Season.Copy().(types.Int32)

	return copied
}

// Equals checks if the given Ranking2Info contains the same data as the current Ranking2Info
func (ri Ranking2Info) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2Info); !ok {
		return false
	}

	other := o.(*Ranking2Info)

	if ri.StructureVersion != other.StructureVersion {
		return false
	}

	if !ri.RankDataList.Equals(other.RankDataList) {
		return false
	}

	if !ri.LowestRank.Equals(other.LowestRank) {
		return false
	}

	if !ri.NumRankedIn.Equals(other.NumRankedIn) {
		return false
	}

	return ri.Season.Equals(other.Season)
}

// CopyRef copies the current value of the Ranking2Info
// and returns a pointer to the new copy
func (ri Ranking2Info) CopyRef() types.RVTypePtr {
	copied := ri.Copy().(Ranking2Info)
	return &copied
}

// Deref takes a pointer to the Ranking2Info
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (ri *Ranking2Info) Deref() types.RVType {
	return *ri
}

// String returns the string representation of the Ranking2Info
func (ri Ranking2Info) String() string {
	return ri.FormatToString(0)
}

// FormatToString pretty-prints the Ranking2Info using the provided indentation level
func (ri Ranking2Info) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2Info{\n")
	b.WriteString(fmt.Sprintf("%sRankDataList: %s,\n", indentationValues, ri.RankDataList))
	b.WriteString(fmt.Sprintf("%sLowestRank: %s,\n", indentationValues, ri.LowestRank))
	b.WriteString(fmt.Sprintf("%sNumRankedIn: %s,\n", indentationValues, ri.NumRankedIn))
	b.WriteString(fmt.Sprintf("%sSeason: %s,\n", indentationValues, ri.Season))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2Info returns a new Ranking2Info
func NewRanking2Info() Ranking2Info {
	return Ranking2Info{
		RankDataList: types.NewList[Ranking2RankData](),
		LowestRank:   types.NewUInt32(0),
		NumRankedIn:  types.NewUInt32(0),
		Season:       types.NewInt32(0),
	}

}
