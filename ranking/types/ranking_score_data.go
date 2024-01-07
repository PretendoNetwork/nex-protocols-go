// Package types implements all the types used by the Ranking protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// RankingScoreData holds general purpose notification data
type RankingScoreData struct {
	types.Structure
	Category   *types.PrimitiveU32
	Score      *types.PrimitiveU32
	OrderBy    *types.PrimitiveU8
	UpdateMode *types.PrimitiveU8
	Groups     []byte
	Param      *types.PrimitiveU64
}

// ExtractFrom extracts the RankingScoreData from the given readable
func (rankingScoreData *RankingScoreData) ExtractFrom(readable types.Readable) error {
	var err error

	if err = rankingScoreData.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read RankingScoreData header. %s", err.Error())
	}

	err = rankingScoreData.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Category from stream. %s", err.Error())
	}

	err = rankingScoreData.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Score from stream. %s", err.Error())
	}

	err = rankingScoreData.OrderBy.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.OrderBy from stream. %s", err.Error())
	}

	err = rankingScoreData.UpdateMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.UpdateMode from stream. %s", err.Error())
	}

	rankingScoreData.Groups, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Groups from stream. %s", err.Error())
	}

	err = rankingScoreData.Param.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Param from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the RankingScoreData to the given writable
func (rankingScoreData *RankingScoreData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rankingScoreData.Category.WriteTo(contentWritable)
	rankingScoreData.Score.WriteTo(contentWritable)
	rankingScoreData.OrderBy.WriteTo(contentWritable)
	rankingScoreData.UpdateMode.WriteTo(contentWritable)
	stream.WriteBuffer(rankingScoreData.Groups)
	rankingScoreData.Param.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rankingScoreData.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of RankingScoreData
func (rankingScoreData *RankingScoreData) Copy() types.RVType {
	copied := NewRankingScoreData()

	copied.StructureVersion = rankingScoreData.StructureVersion

	copied.Category = rankingScoreData.Category
	copied.Score = rankingScoreData.Score
	copied.OrderBy = rankingScoreData.OrderBy
	copied.UpdateMode = rankingScoreData.UpdateMode
	copied.Groups = make([]byte, len(rankingScoreData.Groups))

	copy(copied.Groups, rankingScoreData.Groups)

	copied.Param = rankingScoreData.Param

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingScoreData *RankingScoreData) Equals(o types.RVType) bool {
	if _, ok := o.(*RankingScoreData); !ok {
		return false
	}

	other := o.(*RankingScoreData)

	if rankingScoreData.StructureVersion != other.StructureVersion {
		return false
	}

	if !rankingScoreData.Category.Equals(other.Category) {
		return false
	}

	if !rankingScoreData.Score.Equals(other.Score) {
		return false
	}

	if !rankingScoreData.OrderBy.Equals(other.OrderBy) {
		return false
	}

	if !rankingScoreData.UpdateMode.Equals(other.UpdateMode) {
		return false
	}

	if !rankingScoreData.Groups.Equals(other.Groups) {
		return false
	}

	return rankingScoreData.Param == other.Param
}

// String returns a string representation of the struct
func (rankingScoreData *RankingScoreData) String() string {
	return rankingScoreData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (rankingScoreData *RankingScoreData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingScoreData{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, rankingScoreData.StructureVersion))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, rankingScoreData.Category))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, rankingScoreData.Score))
	b.WriteString(fmt.Sprintf("%sOrderBy: %d,\n", indentationValues, rankingScoreData.OrderBy))
	b.WriteString(fmt.Sprintf("%sUpdateMode: %d,\n", indentationValues, rankingScoreData.UpdateMode))
	b.WriteString(fmt.Sprintf("%sGroups: %x,\n", indentationValues, rankingScoreData.Groups))
	b.WriteString(fmt.Sprintf("%sParam: %d\n", indentationValues, rankingScoreData.Param))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingScoreData returns a new RankingScoreData
func NewRankingScoreData() *RankingScoreData {
	return &RankingScoreData{}
}
