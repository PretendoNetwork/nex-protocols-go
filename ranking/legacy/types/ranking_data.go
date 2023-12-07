// Package types implements all the types used by the Ranking (Legacy) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// RankingData holds information about a rank
type RankingData struct {
	nex.Structure
	UniqueID   uint32
	PID        *nex.PID
	Order      uint32
	Category   uint32
	Scores     []uint32
	Unknown1   uint8
	Unknown2   uint32
	CommonData []byte
}

// ExtractFromStream extracts a RankingData structure from a stream
func (rankingData *RankingData) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	rankingData.UniqueID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.UniqueID from stream. %s", err.Error())
	}

	rankingData.PID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.PID from stream. %s", err.Error())
	}

	rankingData.Order, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.Order from stream. %s", err.Error())
	}

	rankingData.Category, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.Category from stream. %s", err.Error())
	}

	rankingData.Scores, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.Scores from stream. %s", err.Error())
	}

	rankingData.Unknown1, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.Unknown1 from stream. %s", err.Error())
	}

	rankingData.Unknown2, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.Unknown2 from stream. %s", err.Error())
	}

	rankingData.CommonData, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingData.CommonData from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the RankingData and returns a byte array
func (rankingData *RankingData) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(rankingData.UniqueID)
	stream.WritePID(rankingData.PID)
	stream.WriteUInt32LE(rankingData.Order)
	stream.WriteUInt32LE(rankingData.Category)
	stream.WriteListUInt32LE(rankingData.Scores)
	stream.WriteUInt8(rankingData.Unknown1)
	stream.WriteUInt32LE(rankingData.Unknown2)
	stream.WriteBuffer(rankingData.CommonData)

	return stream.Bytes()
}

// Copy returns a new copied instance of RankingData
func (rankingData *RankingData) Copy() nex.StructureInterface {
	copied := NewRankingData()

	copied.SetStructureVersion(rankingData.StructureVersion())

	copied.UniqueID = rankingData.UniqueID
	copied.PID = rankingData.PID.Copy()
	copied.Order = rankingData.Order
	copied.Category = rankingData.Category
	copied.Scores = make([]uint32, len(rankingData.Scores))

	copy(copied.Scores, rankingData.Scores)

	copied.Unknown1 = rankingData.Unknown1
	copied.Unknown2 = rankingData.Unknown2
	copied.CommonData = make([]byte, len(rankingData.CommonData))

	copy(copied.CommonData, rankingData.CommonData)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingData *RankingData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*RankingData)

	if rankingData.StructureVersion() != other.StructureVersion() {
		return false
	}

	if rankingData.UniqueID != other.UniqueID {
		return false
	}

	if !rankingData.PID.Equals(other.PID) {
		return false
	}

	if rankingData.Order != other.Order {
		return false
	}

	if rankingData.Category != other.Category {
		return false
	}

	if len(rankingData.Scores) != len(other.Scores) {
		return false
	}

	for i := 0; i < len(rankingData.Scores); i++ {
		if rankingData.Scores[i] != other.Scores[i] {
			return false
		}
	}

	if rankingData.Unknown1 != other.Unknown1 {
		return false
	}

	if rankingData.Unknown2 != other.Unknown2 {
		return false
	}

	return bytes.Equal(rankingData.CommonData, other.CommonData)
}

// String returns a string representation of the struct
func (rankingData *RankingData) String() string {
	return rankingData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (rankingData *RankingData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingData{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, rankingData.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, rankingData.UniqueID))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, rankingData.PID))
	b.WriteString(fmt.Sprintf("%sOrder: %d,\n", indentationValues, rankingData.Order))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, rankingData.Category))
	b.WriteString(fmt.Sprintf("%sScores: %v,\n", indentationValues, rankingData.Scores))
	b.WriteString(fmt.Sprintf("%sUnknown1: %d,\n", indentationValues, rankingData.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %d,\n", indentationValues, rankingData.Unknown2))
	b.WriteString(fmt.Sprintf("%sCommonData: %x,\n", indentationValues, rankingData.CommonData))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingData returns a new RankingData
func NewRankingData() *RankingData {
	return &RankingData{}
}
