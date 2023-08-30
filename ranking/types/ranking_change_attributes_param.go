// Package types implements all the types used by the Ranking protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// RankingChangeAttributesParam holds parameters for ordering rankings
type RankingChangeAttributesParam struct {
	nex.Structure
	ModificationFlag uint8
	Groups           []uint8
	Param            uint64
}

// ExtractFromStream extracts a RankingChangeAttributesParam structure from a stream
func (rankingChangeAttributesParam *RankingChangeAttributesParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	rankingChangeAttributesParam.ModificationFlag, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingChangeAttributesParam.ModificationFlag from stream. %s", err.Error())
	}

	rankingChangeAttributesParam.Groups, err = stream.ReadListUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingChangeAttributesParam.Groups from stream. %s", err.Error())
	}

	rankingChangeAttributesParam.Param, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingChangeAttributesParam.Param from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the RankingChangeAttributesParam and returns a byte array
func (rankingChangeAttributesParam *RankingChangeAttributesParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(rankingChangeAttributesParam.ModificationFlag)
	stream.WriteListUInt8(rankingChangeAttributesParam.Groups)
	stream.WriteUInt64LE(rankingChangeAttributesParam.Param)

	return stream.Bytes()
}

// Copy returns a new copied instance of RankingChangeAttributesParam
func (rankingChangeAttributesParam *RankingChangeAttributesParam) Copy() nex.StructureInterface {
	copied := NewRankingChangeAttributesParam()

	copied.SetStructureVersion(rankingChangeAttributesParam.StructureVersion())

	copied.ModificationFlag = rankingChangeAttributesParam.ModificationFlag
	copied.Groups = make([]uint8, len(rankingChangeAttributesParam.Groups))

	copy(copied.Groups, rankingChangeAttributesParam.Groups)

	copied.Param = rankingChangeAttributesParam.Param

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingChangeAttributesParam *RankingChangeAttributesParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*RankingChangeAttributesParam)

	if rankingChangeAttributesParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if rankingChangeAttributesParam.ModificationFlag != other.ModificationFlag {
		return false
	}

	if !bytes.Equal(rankingChangeAttributesParam.Groups, other.Groups) {
		return false
	}

	if rankingChangeAttributesParam.Param != other.Param {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, rankingChangeAttributesParam.StructureVersion()))
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
