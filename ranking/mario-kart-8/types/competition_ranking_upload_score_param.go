// Package types implements all the types used by the Ranking (Mario Kart 8) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// CompetitionRankingUploadScoreParam holds data for the Ranking (Mario Kart 8) protocol
type CompetitionRankingUploadScoreParam struct {
	types.Structure
	Unknown  *types.PrimitiveU32
	Unknown2 *types.PrimitiveU32
	Unknown3 *types.PrimitiveU32
	Unknown4 *types.PrimitiveU32
	Unknown5 *types.PrimitiveU8
	Unknown6 *types.PrimitiveU32
	Unknown7 *types.PrimitiveBool
	Metadata []byte
}

// ExtractFrom extracts the CompetitionRankingUploadScoreParam from the given readable
func (competitionRankingUploadScoreParam *CompetitionRankingUploadScoreParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = competitionRankingUploadScoreParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read CompetitionRankingUploadScoreParam header. %s", err.Error())
	}

	err = competitionRankingUploadScoreParam.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown from stream. %s", err.Error())
	}

	err = competitionRankingUploadScoreParam.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown2 from stream. %s", err.Error())
	}

	err = competitionRankingUploadScoreParam.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown3 from stream. %s", err.Error())
	}

	err = competitionRankingUploadScoreParam.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown4 from stream. %s", err.Error())
	}

	err = competitionRankingUploadScoreParam.Unknown5.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown5 from stream. %s", err.Error())
	}

	err = competitionRankingUploadScoreParam.Unknown6.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown6 from stream. %s", err.Error())
	}

	err = competitionRankingUploadScoreParam.Unknown7.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown7 from stream. %s", err.Error())
	}

	competitionRankingUploadScoreParam.Metadata, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Metadata from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the CompetitionRankingUploadScoreParam to the given writable
func (competitionRankingUploadScoreParam *CompetitionRankingUploadScoreParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	competitionRankingUploadScoreParam.Unknown.WriteTo(contentWritable)
	competitionRankingUploadScoreParam.Unknown2.WriteTo(contentWritable)
	competitionRankingUploadScoreParam.Unknown3.WriteTo(contentWritable)
	competitionRankingUploadScoreParam.Unknown4.WriteTo(contentWritable)
	competitionRankingUploadScoreParam.Unknown5.WriteTo(contentWritable)
	competitionRankingUploadScoreParam.Unknown6.WriteTo(contentWritable)
	competitionRankingUploadScoreParam.Unknown7.WriteTo(contentWritable)
	stream.WriteQBuffer(competitionRankingUploadScoreParam.Metadata)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of CompetitionRankingUploadScoreParam
func (competitionRankingUploadScoreParam *CompetitionRankingUploadScoreParam) Copy() types.RVType {
	copied := NewCompetitionRankingUploadScoreParam()

	copied.StructureVersion = competitionRankingUploadScoreParam.StructureVersion

	copied.Unknown = competitionRankingUploadScoreParam.Unknown
	copied.Unknown2 = competitionRankingUploadScoreParam.Unknown2
	copied.Unknown3 = competitionRankingUploadScoreParam.Unknown3
	copied.Unknown4 = competitionRankingUploadScoreParam.Unknown4
	copied.Unknown5 = competitionRankingUploadScoreParam.Unknown5
	copied.Unknown6 = competitionRankingUploadScoreParam.Unknown6
	copied.Unknown7 = competitionRankingUploadScoreParam.Unknown7
	copied.Metadata = competitionRankingUploadScoreParam.Metadata

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (competitionRankingUploadScoreParam *CompetitionRankingUploadScoreParam) Equals(o types.RVType) bool {
	if _, ok := o.(*CompetitionRankingUploadScoreParam); !ok {
		return false
	}

	other := o.(*CompetitionRankingUploadScoreParam)

	if competitionRankingUploadScoreParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !competitionRankingUploadScoreParam.Unknown.Equals(other.Unknown) {
		return false
	}

	if !competitionRankingUploadScoreParam.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !competitionRankingUploadScoreParam.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !competitionRankingUploadScoreParam.Unknown4.Equals(other.Unknown4) {
		return false
	}

	if !competitionRankingUploadScoreParam.Unknown5.Equals(other.Unknown5) {
		return false
	}

	if !competitionRankingUploadScoreParam.Unknown6.Equals(other.Unknown6) {
		return false
	}

	if !competitionRankingUploadScoreParam.Unknown7.Equals(other.Unknown7) {
		return false
	}

	if !competitionRankingUploadScoreParam.Metadata.Equals(other.Metadata) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (competitionRankingUploadScoreParam *CompetitionRankingUploadScoreParam) String() string {
	return competitionRankingUploadScoreParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (competitionRankingUploadScoreParam *CompetitionRankingUploadScoreParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("CompetitionRankingUploadScoreParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, competitionRankingUploadScoreParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUnknown: %d,\n", indentationValues, competitionRankingUploadScoreParam.Unknown))
	b.WriteString(fmt.Sprintf("%sUnknown2: %d,\n", indentationValues, competitionRankingUploadScoreParam.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %d,\n", indentationValues, competitionRankingUploadScoreParam.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %d,\n", indentationValues, competitionRankingUploadScoreParam.Unknown4))
	b.WriteString(fmt.Sprintf("%sUnknown5: %d,\n", indentationValues, competitionRankingUploadScoreParam.Unknown5))
	b.WriteString(fmt.Sprintf("%sUnknown6: %d,\n", indentationValues, competitionRankingUploadScoreParam.Unknown6))
	b.WriteString(fmt.Sprintf("%sUnknown7: %t,\n", indentationValues, competitionRankingUploadScoreParam.Unknown7))
	b.WriteString(fmt.Sprintf("%sMetadata: %x,\n", indentationValues, competitionRankingUploadScoreParam.Metadata))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewCompetitionRankingUploadScoreParam returns a new CompetitionRankingUploadScoreParam
func NewCompetitionRankingUploadScoreParam() *CompetitionRankingUploadScoreParam {
	return &CompetitionRankingUploadScoreParam{}
}
