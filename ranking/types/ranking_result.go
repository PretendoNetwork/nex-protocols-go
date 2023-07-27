// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// RankingResult holds the result of a Ranking get request
type RankingResult struct {
	nex.Structure
	RankDataList []*RankingRankData
	TotalCount   uint32
	SinceTime    *nex.DateTime
}

// ExtractFromStream extracts a RankingResult structure from a stream
func (rankingResult *RankingResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	rankDataList, err := stream.ReadListStructure(NewRankingRankData())
	if err != nil {
		return fmt.Errorf("Failed to extract RankingResult.RankDataList from stream. %s", err.Error())
	}

	rankingResult.RankDataList = rankDataList.([]*RankingRankData)

	rankingResult.TotalCount, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingResult.TotalCount from stream. %s", err.Error())
	}

	rankingResult.SinceTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingResult.SinceTime from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the RankingResult and returns a byte array
func (rankingResult *RankingResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListStructure(rankingResult.RankDataList)
	stream.WriteUInt32LE(rankingResult.TotalCount)
	stream.WriteDateTime(rankingResult.SinceTime)

	return stream.Bytes()
}

// Copy returns a new copied instance of RankingResult
func (rankingResult *RankingResult) Copy() nex.StructureInterface {
	copied := NewRankingResult()

	copied.RankDataList = make([]*RankingRankData, len(rankingResult.RankDataList))

	for i := 0; i < len(rankingResult.RankDataList); i++ {
		copied.RankDataList[i] = rankingResult.RankDataList[i].Copy().(*RankingRankData)
	}

	copied.TotalCount = rankingResult.TotalCount

	if rankingResult.SinceTime != nil {
		copied.SinceTime = rankingResult.SinceTime.Copy()
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingResult *RankingResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*RankingResult)

	if len(rankingResult.RankDataList) != len(other.RankDataList) {
		return false
	}

	for i := 0; i < len(rankingResult.RankDataList); i++ {
		if !rankingResult.RankDataList[i].Equals(other.RankDataList[i]) {
			return false
		}
	}

	if rankingResult.TotalCount != other.TotalCount {
		return false
	}

	if rankingResult.SinceTime == nil && other.SinceTime != nil {
		return false
	}

	if rankingResult.SinceTime != nil && other.SinceTime == nil {
		return false
	}

	if rankingResult.SinceTime != nil && other.SinceTime != nil {
		if !rankingResult.SinceTime.Equals(other.SinceTime) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (rankingResult *RankingResult) String() string {
	return rankingResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (rankingResult *RankingResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationListValues := strings.Repeat("\t", indentationLevel+2)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingResult{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, rankingResult.StructureVersion()))

	if len(rankingResult.RankDataList) == 0 {
		b.WriteString(fmt.Sprintf("%sRankDataList: [],\n", indentationValues))
	} else {
		b.WriteString(fmt.Sprintf("%sRankDataList: [\n", indentationValues))

		for i := 0; i < len(rankingResult.RankDataList); i++ {
			str := rankingResult.RankDataList[i].FormatToString(indentationLevel + 2)
			if i == len(rankingResult.RankDataList)-1 {
				b.WriteString(fmt.Sprintf("%s%s\n", indentationListValues, str))
			} else {
				b.WriteString(fmt.Sprintf("%s%s,\n", indentationListValues, str))
			}
		}

		b.WriteString(fmt.Sprintf("%s],\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sTotalCount: %d,\n", indentationValues, rankingResult.TotalCount))

	if rankingResult.SinceTime != nil {
		b.WriteString(fmt.Sprintf("%sSinceTime: %s\n", indentationValues, rankingResult.SinceTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sSinceTime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingResult returns a new RankingResult
func NewRankingResult() *RankingResult {
	return &RankingResult{}
}
