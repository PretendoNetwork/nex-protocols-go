// Package types implements all the types used by the Ranking protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// RankingRankData holds parameters for ordering rankings
type RankingRankData struct {
	nex.Structure
	PrincipalID uint32
	UniqueID    uint64
	Order       uint32
	Category    uint32
	Score       uint32
	Groups      []byte
	Param       uint64
	CommonData  []byte
	UpdateTime  *nex.DateTime // * NEX 3.6.0+
}

// ExtractFromStream extracts a RankingRankData structure from a stream
func (rankingRankData *RankingRankData) ExtractFromStream(stream *nex.StreamIn) error {
	nexVersion := stream.Server.LibraryVersion()

	var err error

	rankingRankData.PrincipalID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.PrincipalID from stream. %s", err.Error())
	}

	rankingRankData.UniqueID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.UniqueID from stream. %s", err.Error())
	}

	rankingRankData.Order, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Order from stream. %s", err.Error())
	}

	rankingRankData.Category, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Category from stream. %s", err.Error())
	}

	rankingRankData.Score, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Score from stream. %s", err.Error())
	}

	rankingRankData.Groups, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Groups from stream. %s", err.Error())
	}

	rankingRankData.Param, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.Param from stream. %s", err.Error())
	}

	rankingRankData.CommonData, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract RankingRankData.CommonData from stream. %s", err.Error())
	}

	if nexVersion.GreaterOrEqual("3.6.0") {
		rankingRankData.UpdateTime, err = stream.ReadDateTime()
		if err != nil {
			return fmt.Errorf("Failed to extract RankingRankData.UpdateTime from stream. %s", err.Error())
		}
	}

	return nil
}

// Bytes encodes the RankingRankData and returns a byte array
func (rankingRankData *RankingRankData) Bytes(stream *nex.StreamOut) []byte {
	nexVersion := stream.Server.LibraryVersion()

	stream.WriteUInt32LE(rankingRankData.PrincipalID)
	stream.WriteUInt64LE(rankingRankData.UniqueID)
	stream.WriteUInt32LE(rankingRankData.Order)
	stream.WriteUInt32LE(rankingRankData.Category)
	stream.WriteUInt32LE(rankingRankData.Score)
	stream.WriteBuffer(rankingRankData.Groups)
	stream.WriteUInt64LE(rankingRankData.Param)
	stream.WriteBuffer(rankingRankData.CommonData)

	if nexVersion.GreaterOrEqual("4.0.0") {
		stream.WriteDateTime(rankingRankData.UpdateTime)
	}

	return stream.Bytes()
}

// Copy returns a new copied instance of RankingRankData
func (rankingRankData *RankingRankData) Copy() nex.StructureInterface {
	copied := NewRankingRankData()

	copied.SetStructureVersion(rankingRankData.StructureVersion())

	copied.PrincipalID = rankingRankData.PrincipalID
	copied.UniqueID = rankingRankData.UniqueID
	copied.Order = rankingRankData.Order
	copied.Category = rankingRankData.Category
	copied.Score = rankingRankData.Score
	copied.Groups = make([]byte, len(rankingRankData.Groups))

	copy(copied.Groups, rankingRankData.Groups)

	copied.Param = rankingRankData.Param
	copied.CommonData = make([]byte, len(rankingRankData.CommonData))

	copy(copied.CommonData, rankingRankData.CommonData)

	if rankingRankData.UpdateTime != nil {
		copied.UpdateTime = rankingRankData.UpdateTime.Copy()
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (rankingRankData *RankingRankData) Equals(structure nex.StructureInterface) bool {
	other := structure.(*RankingRankData)

	if rankingRankData.StructureVersion() != other.StructureVersion() {
		return false
	}

	if rankingRankData.PrincipalID != other.PrincipalID {
		return false
	}

	if rankingRankData.UniqueID != other.UniqueID {
		return false
	}

	if rankingRankData.Order != other.Order {
		return false
	}

	if rankingRankData.Category != other.Category {
		return false
	}

	if rankingRankData.Score != other.Score {
		return false
	}

	if !bytes.Equal(rankingRankData.Groups, other.Groups) {
		return false
	}

	if rankingRankData.Param != other.Param {
		return false
	}

	if !bytes.Equal(rankingRankData.CommonData, other.CommonData) {
		return false
	}

	if rankingRankData.UpdateTime == nil && other.UpdateTime != nil {
		return false
	}

	if rankingRankData.UpdateTime != nil && other.UpdateTime == nil {
		return false
	}

	if rankingRankData.UpdateTime != nil && other.UpdateTime != nil {
		if !rankingRankData.UpdateTime.Equals(other.UpdateTime) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (rankingRankData *RankingRankData) String() string {
	return rankingRankData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (rankingRankData *RankingRankData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RankingRankData{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, rankingRankData.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %d,\n", indentationValues, rankingRankData.PrincipalID))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, rankingRankData.UniqueID))
	b.WriteString(fmt.Sprintf("%sOrder: %d,\n", indentationValues, rankingRankData.Order))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, rankingRankData.Category))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, rankingRankData.Score))
	b.WriteString(fmt.Sprintf("%sGroups: %x,\n", indentationValues, rankingRankData.Groups))
	b.WriteString(fmt.Sprintf("%sParam: %d,\n", indentationValues, rankingRankData.Param))
	b.WriteString(fmt.Sprintf("%sCommonData: %x,\n", indentationValues, rankingRankData.CommonData))

	if rankingRankData.UpdateTime != nil {
		b.WriteString(fmt.Sprintf("%sUpdateTime: %s\n", indentationValues, rankingRankData.UpdateTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUpdateTime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRankingRankData returns a new RankingRankData
func NewRankingRankData() *RankingRankData {
	return &RankingRankData{}
}
