// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// RankingCachedResult holds the result of a Ranking get request
type RankingCachedResult struct {
	types.Structure
	*RankingResult
	CreatedTime *types.DateTime
	ExpiredTime *types.DateTime
	MaxLength   *types.PrimitiveU8
}

// ExtractFrom extracts the RankingCachedResult from the given readable
func (rankingCachedResult *RankingCachedResult) ExtractFrom(readable types.Readable) error {
	var err error

	if err = rankingCachedResult.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read RankingCachedResult header. %s", err.Error())
	}

	err = rankingCachedResult.CreatedTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingCachedResult.CreatedTime from stream. %s", err.Error())
	}

	err = rankingCachedResult.ExpiredTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingCachedResult.ExpiredTime from stream. %s", err.Error())
	}

	err = rankingCachedResult.MaxLength.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingCachedResult.MaxLength from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the RankingCachedResult to the given writable
func (rankingCachedResult *RankingCachedResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rankingCachedResult.CreatedTime.WriteTo(contentWritable)
	rankingCachedResult.ExpiredTime.WriteTo(contentWritable)
	rankingCachedResult.MaxLength.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rankingCachedResult.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of RankingCachedResult
func (rankingCachedResult *RankingCachedResult) Copy() types.RVType {
	copied := NewRankingCachedResult()

	copied.StructureVersion = rankingCachedResult.StructureVersion

	copied.RankingResult = rankingCachedResult.RankingResult.Copy().(*RankingResult)

	copied.CreatedTime = rankingCachedResult.CreatedTime.Copy()

	copied.ExpiredTime = rankingCachedResult.ExpiredTime.Copy()

	copied.MaxLength = rankingCachedResult.MaxLength

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingCachedResult *RankingCachedResult) Equals(o types.RVType) bool {
	if _, ok := o.(*RankingCachedResult); !ok {
		return false
	}

	other := o.(*RankingCachedResult)

	if rankingCachedResult.StructureVersion != other.StructureVersion {
		return false
	}

	if !rankingCachedResult.CreatedTime.Equals(other.CreatedTime) {
		return false
	}

	if !rankingCachedResult.ExpiredTime.Equals(other.ExpiredTime) {
		return false
	}

	if !rankingCachedResult.MaxLength.Equals(other.MaxLength) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, rankingCachedResult.StructureVersion))

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
