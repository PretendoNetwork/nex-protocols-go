// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Ranking2Info holds data for the Ranking 2  protocol
type Ranking2Info struct {
	nex.Structure
	RankDataList []*Ranking2RankData
	LowestRank   uint32
	NumRankedIn  uint32
	Season       int32
}

// ExtractFromStream extracts a Ranking2Info structure from a stream
func (ranking2Info *Ranking2Info) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	rankDataList, err := stream.ReadListStructure(NewRanking2RankData())
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info.RankDataList from stream. %s", err.Error())
	}

	ranking2Info.RankDataList = rankDataList.([]*Ranking2RankData)

	ranking2Info.LowestRank, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info.LowestRank from stream. %s", err.Error())
	}

	ranking2Info.NumRankedIn, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info.NumRankedIn from stream. %s", err.Error())
	}

	ranking2Info.Season, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2Info.Season from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the Ranking2Info and returns a byte array
func (ranking2Info *Ranking2Info) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListStructure(ranking2Info.RankDataList)
	stream.WriteUInt32LE(ranking2Info.LowestRank)
	stream.WriteUInt32LE(ranking2Info.NumRankedIn)
	stream.WriteInt32LE(ranking2Info.Season)

	return stream.Bytes()
}

// Copy returns a new copied instance of Ranking2Info
func (ranking2Info *Ranking2Info) Copy() nex.StructureInterface {
	copied := NewRanking2Info()

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
func (ranking2Info *Ranking2Info) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Ranking2Info)

	if len(ranking2Info.RankDataList) != len(other.RankDataList) {
		return false
	}

	for i := 0; i < len(ranking2Info.RankDataList); i++ {
		if !ranking2Info.RankDataList[i].Equals(other.RankDataList[i]) {
			return false
		}
	}

	if ranking2Info.LowestRank != other.LowestRank {
		return false
	}

	if ranking2Info.NumRankedIn != other.NumRankedIn {
		return false
	}

	if ranking2Info.Season != other.Season {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, ranking2Info.StructureVersion()))

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
