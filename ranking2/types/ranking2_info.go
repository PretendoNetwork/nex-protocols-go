// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Ranking2Info holds data for the Ranking 2  protocol
type Ranking2Info struct {
	types.Structure
	RankDataList []*Ranking2RankData
	LowestRank   *types.PrimitiveU32
	NumRankedIn  *types.PrimitiveU32
	Season       *types.PrimitiveS32
}

// ExtractFrom extracts the Ranking2Info from the given readable
func (ranking2Info *Ranking2Info) ExtractFrom(readable types.Readable) error {
	var err error

	if err = ranking2Info.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Ranking2Info header. %s", err.Error())
	}

	rankDataList, err := nex.StreamReadListStructure(stream, NewRanking2RankData())
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info.RankDataList from stream. %s", err.Error())
	}

	ranking2Info.RankDataList = rankDataList

	err = ranking2Info.LowestRank.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info.LowestRank from stream. %s", err.Error())
	}

	err = ranking2Info.NumRankedIn.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info.NumRankedIn from stream. %s", err.Error())
	}

	err = ranking2Info.Season.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info.Season from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Ranking2Info to the given writable
func (ranking2Info *Ranking2Info) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ranking2Info.RankDataList.WriteTo(contentWritable)
	ranking2Info.LowestRank.WriteTo(contentWritable)
	ranking2Info.NumRankedIn.WriteTo(contentWritable)
	ranking2Info.Season.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ranking2Info.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Ranking2Info
func (ranking2Info *Ranking2Info) Copy() types.RVType {
	copied := NewRanking2Info()

	copied.StructureVersion = ranking2Info.StructureVersion

	copied.RankDataList = make([]*Ranking2RankData, len(ranking2Info.RankDataList))

	for i := 0; i < len(ranking2Info.RankDataList); i++ {
		copied.RankDataList[i] = ranking2Info.RankDataList[i].Copy().(*Ranking2RankData)
	}

	copied.LowestRank = ranking2Info.LowestRank
	copied.NumRankedIn = ranking2Info.NumRankedIn
	copied.Season = ranking2Info.Season
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2Info *Ranking2Info) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2Info); !ok {
		return false
	}

	other := o.(*Ranking2Info)

	if ranking2Info.StructureVersion != other.StructureVersion {
		return false
	}

	if len(ranking2Info.RankDataList) != len(other.RankDataList) {
		return false
	}

	for i := 0; i < len(ranking2Info.RankDataList); i++ {
		if !ranking2Info.RankDataList[i].Equals(other.RankDataList[i]) {
			return false
		}
	}

	if !ranking2Info.LowestRank.Equals(other.LowestRank) {
		return false
	}

	if !ranking2Info.NumRankedIn.Equals(other.NumRankedIn) {
		return false
	}

	if !ranking2Info.Season.Equals(other.Season) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (ranking2Info *Ranking2Info) String() string {
	return ranking2Info.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2Info *Ranking2Info) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2Info{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, ranking2Info.StructureVersion))

	if len(ranking2Info.RankDataList) == 0 {
		b.WriteString(fmt.Sprintf("%sRankDataList: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sRankDataList: [\n", indentationValues))

		for i := 0; i < len(ranking2Info.RankDataList); i++ {
			str := ranking2Info.RankDataList[i].FormatToString(indentationLevel + 2)
			if i == len(ranking2Info.RankDataList)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sLowestRank: %d,\n", indentationValues, ranking2Info.LowestRank))
	b.WriteString(fmt.Sprintf("%sNumRankedIn: %d,\n", indentationValues, ranking2Info.NumRankedIn))
	b.WriteString(fmt.Sprintf("%sSeason: %d,\n", indentationValues, ranking2Info.Season))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2Info returns a new Ranking2Info
func NewRanking2Info() *Ranking2Info {
	return &Ranking2Info{}
}
