// Package types implements all the types used by the Ranking protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// RankingScoreData holds general purpose notification data
type RankingScoreData struct {
	nex.Structure
	Category   uint32
	Score      uint32
	OrderBy    uint8
	UpdateMode uint8
	Groups     []byte
	Param      uint64
}

// ExtractFromStream extracts a RankingScoreData structure from a stream
func (rankingScoreData *RankingScoreData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	rankingScoreData.Category, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Category from stream. %s", err.Error())
	}

	rankingScoreData.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Score from stream. %s", err.Error())
	}

	rankingScoreData.OrderBy, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.OrderBy from stream. %s", err.Error())
	}

	rankingScoreData.UpdateMode, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.UpdateMode from stream. %s", err.Error())
	}

	rankingScoreData.Groups, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Groups from stream. %s", err.Error())
	}

	rankingScoreData.Param, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingScoreData.Param from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the RankingScoreData and returns a byte array
func (rankingScoreData *RankingScoreData) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(rankingScoreData.Category)
	stream.WriteUInt32LE(rankingScoreData.Score)
	stream.WriteUInt8(rankingScoreData.OrderBy)
	stream.WriteUInt8(rankingScoreData.UpdateMode)
	stream.WriteBuffer(rankingScoreData.Groups)
	stream.WriteUInt64LE(rankingScoreData.Param)

	return stream.Bytes()
}

// Copy returns a new copied instance of RankingScoreData
func (rankingScoreData *RankingScoreData) Copy() nex.StructureInterface {
	copied := NewRankingScoreData()

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
func (rankingScoreData *RankingScoreData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*RankingScoreData)

	if rankingScoreData.Category != other.Category {
		return false
	}

	if rankingScoreData.Score != other.Score {
		return false
	}

	if rankingScoreData.OrderBy != other.OrderBy {
		return false
	}

	if rankingScoreData.UpdateMode != other.UpdateMode {
		return false
	}

	if !bytes.Equal(rankingScoreData.Groups, other.Groups) {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, rankingScoreData.StructureVersion()))
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
