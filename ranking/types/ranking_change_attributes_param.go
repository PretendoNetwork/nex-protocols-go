// Package types implements all the types used by the Ranking protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// RankingChangeAttributesParam holds parameters for ordering rankings
type RankingChangeAttributesParam struct {
	types.Structure
	ModificationFlag *types.PrimitiveU8
	Groups           *types.List[*types.PrimitiveU8]
	Param            *types.PrimitiveU64
}

// ExtractFrom extracts the RankingChangeAttributesParam from the given readable
func (rankingChangeAttributesParam *RankingChangeAttributesParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = rankingChangeAttributesParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read RankingChangeAttributesParam header. %s", err.Error())
	}

	err = rankingChangeAttributesParam.ModificationFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingChangeAttributesParam.ModificationFlag from stream. %s", err.Error())
	}

	err = rankingChangeAttributesParam.Groups.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingChangeAttributesParam.Groups from stream. %s", err.Error())
	}

	err = rankingChangeAttributesParam.Param.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RankingChangeAttributesParam.Param from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the RankingChangeAttributesParam to the given writable
func (rankingChangeAttributesParam *RankingChangeAttributesParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rankingChangeAttributesParam.ModificationFlag.WriteTo(contentWritable)
	rankingChangeAttributesParam.Groups.WriteTo(contentWritable)
	rankingChangeAttributesParam.Param.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rankingChangeAttributesParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of RankingChangeAttributesParam
func (rankingChangeAttributesParam *RankingChangeAttributesParam) Copy() types.RVType {
	copied := NewRankingChangeAttributesParam()

	copied.StructureVersion = rankingChangeAttributesParam.StructureVersion

	copied.ModificationFlag = rankingChangeAttributesParam.ModificationFlag
	copied.Groups = make(*types.List[*types.PrimitiveU8], len(rankingChangeAttributesParam.Groups))

	copy(copied.Groups, rankingChangeAttributesParam.Groups)

	copied.Param = rankingChangeAttributesParam.Param

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingChangeAttributesParam *RankingChangeAttributesParam) Equals(o types.RVType) bool {
	if _, ok := o.(*RankingChangeAttributesParam); !ok {
		return false
	}

	other := o.(*RankingChangeAttributesParam)

	if rankingChangeAttributesParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !rankingChangeAttributesParam.ModificationFlag.Equals(other.ModificationFlag) {
		return false
	}

	if !rankingChangeAttributesParam.Groups.Equals(other.Groups) {
		return false
	}

	if !rankingChangeAttributesParam.Param.Equals(other.Param) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (rankingChangeAttributesParam *RankingChangeAttributesParam) String() string {
	return rankingChangeAttributesParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (rankingChangeAttributesParam *RankingChangeAttributesParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingChangeAttributesParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, rankingChangeAttributesParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sModificationFlag: %d,\n", indentationValues, rankingChangeAttributesParam.ModificationFlag))
	b.WriteString(fmt.Sprintf("%sGroups: %v,\n", indentationValues, rankingChangeAttributesParam.Groups))
	b.WriteString(fmt.Sprintf("%sParam: %d\n", indentationValues, rankingChangeAttributesParam.Param))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingChangeAttributesParam returns a new RankingChangeAttributesParam
func NewRankingChangeAttributesParam() *RankingChangeAttributesParam {
	return &RankingChangeAttributesParam{}
}
