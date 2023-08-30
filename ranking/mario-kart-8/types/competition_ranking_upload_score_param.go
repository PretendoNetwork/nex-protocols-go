// Package types implements all the types used by the Ranking (Mario Kart 8) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// CompetitionRankingUploadScoreParam holds data for the Ranking (Mario Kart 8) protocol
type CompetitionRankingUploadScoreParam struct {
	nex.Structure
	Unknown  uint32
	Unknown2 uint32
	Unknown3 uint32
	Unknown4 uint32
	Unknown5 uint8
	Unknown6 uint32
	Unknown7 bool
	Metadata []byte
}

// ExtractFromStream extracts a CompetitionRankingUploadScoreParam structure from a stream
func (competitionRankingUploadScoreParam *CompetitionRankingUploadScoreParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	competitionRankingUploadScoreParam.Unknown, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown from stream. %s", err.Error())
	}

	competitionRankingUploadScoreParam.Unknown2, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown2 from stream. %s", err.Error())
	}

	competitionRankingUploadScoreParam.Unknown3, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown3 from stream. %s", err.Error())
	}

	competitionRankingUploadScoreParam.Unknown4, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown4 from stream. %s", err.Error())
	}

	competitionRankingUploadScoreParam.Unknown5, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown5 from stream. %s", err.Error())
	}

	competitionRankingUploadScoreParam.Unknown6, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown6 from stream. %s", err.Error())
	}

	competitionRankingUploadScoreParam.Unknown7, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Unknown7 from stream. %s", err.Error())
	}

	competitionRankingUploadScoreParam.Metadata, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract CompetitionRankingUploadScoreParam.Metadata from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the CompetitionRankingUploadScoreParam and returns a byte array
func (competitionRankingUploadScoreParam *CompetitionRankingUploadScoreParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(competitionRankingUploadScoreParam.Unknown)
	stream.WriteUInt32LE(competitionRankingUploadScoreParam.Unknown2)
	stream.WriteUInt32LE(competitionRankingUploadScoreParam.Unknown3)
	stream.WriteUInt32LE(competitionRankingUploadScoreParam.Unknown4)
	stream.WriteUInt8(competitionRankingUploadScoreParam.Unknown5)
	stream.WriteUInt32LE(competitionRankingUploadScoreParam.Unknown6)
	stream.WriteBool(competitionRankingUploadScoreParam.Unknown7)
	stream.WriteQBuffer(competitionRankingUploadScoreParam.Metadata)

	return stream.Bytes()
}

// Copy returns a new copied instance of CompetitionRankingUploadScoreParam
func (competitionRankingUploadScoreParam *CompetitionRankingUploadScoreParam) Copy() nex.StructureInterface {
	copied := NewCompetitionRankingUploadScoreParam()

	copied.SetStructureVersion(competitionRankingUploadScoreParam.StructureVersion())

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
func (competitionRankingUploadScoreParam *CompetitionRankingUploadScoreParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*CompetitionRankingUploadScoreParam)

	if competitionRankingUploadScoreParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if competitionRankingUploadScoreParam.Unknown != other.Unknown {
		return false
	}

	if competitionRankingUploadScoreParam.Unknown2 != other.Unknown2 {
		return false
	}

	if competitionRankingUploadScoreParam.Unknown3 != other.Unknown3 {
		return false
	}

	if competitionRankingUploadScoreParam.Unknown4 != other.Unknown4 {
		return false
	}

	if competitionRankingUploadScoreParam.Unknown5 != other.Unknown5 {
		return false
	}

	if competitionRankingUploadScoreParam.Unknown6 != other.Unknown6 {
		return false
	}

	if competitionRankingUploadScoreParam.Unknown7 != other.Unknown7 {
		return false
	}

	if !bytes.Equal(competitionRankingUploadScoreParam.Metadata, other.Metadata) {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, competitionRankingUploadScoreParam.StructureVersion()))
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
