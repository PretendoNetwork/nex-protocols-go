// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// RankingResult holds the result of a Ranking get request
type RankingResult struct {
	types.Structure
	RankDataList []*RankingRankData
	TotalCount   *types.PrimitiveU32
	SinceTime    *types.DateTime
}

// ExtractFrom extracts the RankingResult from the given readable
func (rankingResult *RankingResult) ExtractFrom(readable types.Readable) error {
	var err error

	if err = rankingResult.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read RankingResult header. %s", err.Error())
	}

	rankDataList, err := nex.StreamReadListStructure(stream, NewRankingRankData())
	if err != nil {
		return fmt.Errorf("Failed to extract RankingResult.RankDataList from stream. %s", err.Error())
	}

	rankingResult.RankDataList = rankDataList

	err = rankingResult.TotalCount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingResult.TotalCount from stream. %s", err.Error())
	}

	err = rankingResult.SinceTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingResult.SinceTime from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the RankingResult to the given writable
func (rankingResult *RankingResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rankingResult.RankDataList.WriteTo(contentWritable)
	rankingResult.TotalCount.WriteTo(contentWritable)
	rankingResult.SinceTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rankingResult.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of RankingResult
func (rankingResult *RankingResult) Copy() types.RVType {
	copied := NewRankingResult()

	copied.StructureVersion = rankingResult.StructureVersion

	copied.RankDataList = make([]*RankingRankData, len(rankingResult.RankDataList))

	for i := 0; i < len(rankingResult.RankDataList); i++ {
		copied.RankDataList[i] = rankingResult.RankDataList[i].Copy().(*RankingRankData)
	}

	copied.TotalCount = rankingResult.TotalCount

	copied.SinceTime = rankingResult.SinceTime.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingResult *RankingResult) Equals(o types.RVType) bool {
	if _, ok := o.(*RankingResult); !ok {
		return false
	}

	other := o.(*RankingResult)

	if rankingResult.StructureVersion != other.StructureVersion {
		return false
	}

	if len(rankingResult.RankDataList) != len(other.RankDataList) {
		return false
	}

	for i := 0; i < len(rankingResult.RankDataList); i++ {
		if !rankingResult.RankDataList[i].Equals(other.RankDataList[i]) {
			return false
		}
	}

	if !rankingResult.TotalCount.Equals(other.TotalCount) {
		return false
	}

	if !rankingResult.SinceTime.Equals(other.SinceTime) {
		return false
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, rankingResult.StructureVersion))

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
