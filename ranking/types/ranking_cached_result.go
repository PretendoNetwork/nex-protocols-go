// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// RankingCachedResult holds the result of a Ranking get request
type RankingCachedResult struct {
	nex.Structure
	*RankingResult
	CreatedTime *nex.DateTime
	ExpiredTime *nex.DateTime
	MaxLength   uint8
}

// ExtractFromStream extracts a RankingCachedResult structure from a stream
func (rankingCachedResult *RankingCachedResult) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	rankingCachedResult.CreatedTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingCachedResult.CreatedTime from stream. %s", err.Error())
	}

	rankingCachedResult.ExpiredTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingCachedResult.ExpiredTime from stream. %s", err.Error())
	}

	rankingCachedResult.MaxLength, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingCachedResult.MaxLength from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the RankingCachedResult and returns a byte array
func (rankingCachedResult *RankingCachedResult) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteDateTime(rankingCachedResult.CreatedTime)
	stream.WriteDateTime(rankingCachedResult.ExpiredTime)
	stream.WriteUInt8(rankingCachedResult.MaxLength)

	return stream.Bytes()
}

// Copy returns a new copied instance of RankingCachedResult
func (rankingCachedResult *RankingCachedResult) Copy() nex.StructureInterface {
	copied := NewRankingCachedResult()

	copied.RankingResult = rankingCachedResult.RankingResult.Copy().(*RankingResult)
	copied.SetParentType(copied.RankingResult)

	if rankingCachedResult.CreatedTime != nil {
		copied.CreatedTime = rankingCachedResult.CreatedTime.Copy()
	}

	if rankingCachedResult.ExpiredTime != nil {
		copied.ExpiredTime = rankingCachedResult.ExpiredTime.Copy()
	}

	copied.MaxLength = rankingCachedResult.MaxLength

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingCachedResult *RankingCachedResult) Equals(structure nex.StructureInterface) bool {
	other := structure.(*RankingCachedResult)

	if rankingCachedResult.CreatedTime == nil && other.CreatedTime != nil {
		return false
	}

	if rankingCachedResult.CreatedTime != nil && other.CreatedTime == nil {
		return false
	}

	if rankingCachedResult.CreatedTime != nil && other.CreatedTime != nil {
		if !rankingCachedResult.CreatedTime.Equals(other.CreatedTime) {
			return false
		}
	}

	if rankingCachedResult.ExpiredTime == nil && other.ExpiredTime != nil {
		return false
	}

	if rankingCachedResult.ExpiredTime != nil && other.ExpiredTime == nil {
		return false
	}

	if rankingCachedResult.ExpiredTime != nil && other.SinceTime != nil {
		if !rankingCachedResult.ExpiredTime.Equals(other.ExpiredTime) {
			return false
		}
	}

	if rankingCachedResult.MaxLength != other.MaxLength {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (rankingCachedResult *RankingCachedResult) String() string {
	return rankingCachedResult.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (rankingCachedResult *RankingCachedResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingCachedResult{\n")
	b.WriteString(fmt.Sprintf("%sParentType: %s,\n", indentationValues, rankingCachedResult.ParentType().FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, rankingCachedResult.StructureVersion()))

	if rankingCachedResult.CreatedTime != nil {
		b.WriteString(fmt.Sprintf("%sCreatedTime: %s\n", indentationValues, rankingCachedResult.CreatedTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCreatedTime: nil\n", indentationValues))
	}

	if rankingCachedResult.ExpiredTime != nil {
		b.WriteString(fmt.Sprintf("%sExpiredTime: %s\n", indentationValues, rankingCachedResult.ExpiredTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sExpiredTime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sMaxLength: %d\n", indentationValues, rankingCachedResult.MaxLength))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingCachedResult returns a new RankingCachedResult
func NewRankingCachedResult() *RankingCachedResult {
	return &RankingCachedResult{}
}
