// Package types implements all the types used by the Ranking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// CompetitionRankingInfoGetParam is a type within the Ranking protocol
type CompetitionRankingInfoGetParam struct {
	types.Structure
	Unknown types.UInt8
	Result  types.ResultRange
}

// WriteTo writes the CompetitionRankingInfoGetParam to the given writable
func (crigp CompetitionRankingInfoGetParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	crigp.Unknown.WriteTo(contentWritable)
	crigp.Result.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	crigp.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the CompetitionRankingInfoGetParam from the given readable
func (crigp *CompetitionRankingInfoGetParam) ExtractFrom(readable types.Readable) error {
	var err error

	err = crigp.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfoGetParam header. %s", err.Error())
	}

	err = crigp.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfoGetParam.Unknown. %s", err.Error())
	}

	err = crigp.Result.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingInfoGetParam.Result. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of CompetitionRankingInfoGetParam
func (crigp CompetitionRankingInfoGetParam) Copy() types.RVType {
	copied := NewCompetitionRankingInfoGetParam()

	copied.StructureVersion = crigp.StructureVersion
	copied.Unknown = crigp.Unknown.Copy().(types.UInt8)
	copied.Result = crigp.Result.Copy().(types.ResultRange)

	return copied
}

// Equals checks if the given CompetitionRankingInfoGetParam contains the same data as the current CompetitionRankingInfoGetParam
func (crigp CompetitionRankingInfoGetParam) Equals(o types.RVType) bool {
	if _, ok := o.(CompetitionRankingInfoGetParam); !ok {
		return false
	}

	other := o.(CompetitionRankingInfoGetParam)

	if crigp.StructureVersion != other.StructureVersion {
		return false
	}

	if !crigp.Unknown.Equals(other.Unknown) {
		return false
	}

	return crigp.Result.Equals(other.Result)
}

// CopyRef copies the current value of the CompetitionRankingInfoGetParam
// and returns a pointer to the new copy
func (crigp CompetitionRankingInfoGetParam) CopyRef() types.RVTypePtr {
	copied := crigp.Copy().(CompetitionRankingInfoGetParam)
	return &copied
}

// Deref takes a pointer to the CompetitionRankingInfoGetParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (crigp *CompetitionRankingInfoGetParam) Deref() types.RVType {
	return *crigp
}

// String returns the string representation of the CompetitionRankingInfoGetParam
func (crigp CompetitionRankingInfoGetParam) String() string {
	return crigp.FormatToString(0)
}

// FormatToString pretty-prints the CompetitionRankingInfoGetParam using the provided indentation level
func (crigp CompetitionRankingInfoGetParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CompetitionRankingInfoGetParam{\n")
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, crigp.Unknown))
	b.WriteString(fmt.Sprintf("%sResult: %s,\n", indentationValues, crigp.Result.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewCompetitionRankingInfoGetParam returns a new CompetitionRankingInfoGetParam
func NewCompetitionRankingInfoGetParam() CompetitionRankingInfoGetParam {
	return CompetitionRankingInfoGetParam{
		Unknown: types.NewUInt8(0),
		Result:  types.NewResultRange(),
	}

}
