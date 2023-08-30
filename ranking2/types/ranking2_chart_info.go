// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Ranking2ChartInfo holds data for the Ranking 2  protocol
type Ranking2ChartInfo struct {
	nex.Structure
	CreateTime           *nex.DateTime
	Index                uint32
	Category             uint32
	Season               int32
	BinsSize             uint8
	SamplingRate         uint8
	ScoreOrder           bool
	EstimateLength       uint32
	EstimateHighestScore uint32
	EstimateLowestScore  uint32
	EstimateMedianScore  uint32
	EstimateAverageScore float64
	HighestBinsScore     uint32
	LowestBinsScore      uint32
	BinsWidth            uint32
	Attribute1           uint32
	Attribute2           uint32
	Quantities           []uint32
}

// ExtractFromStream extracts a Ranking2ChartInfo structure from a stream
func (ranking2ChartInfo *Ranking2ChartInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	ranking2ChartInfo.CreateTime, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.CreateTime from stream. %s", err.Error())
	}

	ranking2ChartInfo.Index, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Index from stream. %s", err.Error())
	}

	ranking2ChartInfo.Category, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Category from stream. %s", err.Error())
	}

	ranking2ChartInfo.Season, err = stream.ReadInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Season from stream. %s", err.Error())
	}

	ranking2ChartInfo.BinsSize, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.BinsSize from stream. %s", err.Error())
	}

	ranking2ChartInfo.SamplingRate, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.SamplingRate from stream. %s", err.Error())
	}

	ranking2ChartInfo.ScoreOrder, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.ScoreOrder from stream. %s", err.Error())
	}

	ranking2ChartInfo.EstimateLength, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateLength from stream. %s", err.Error())
	}

	ranking2ChartInfo.EstimateHighestScore, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateHighestScore from stream. %s", err.Error())
	}

	ranking2ChartInfo.EstimateLowestScore, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateLowestScore from stream. %s", err.Error())
	}

	ranking2ChartInfo.EstimateMedianScore, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateMedianScore from stream. %s", err.Error())
	}

	ranking2ChartInfo.EstimateAverageScore, err = stream.ReadFloat64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateAverageScore from stream. %s", err.Error())
	}

	ranking2ChartInfo.HighestBinsScore, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.HighestBinsScore from stream. %s", err.Error())
	}

	ranking2ChartInfo.LowestBinsScore, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.LowestBinsScore from stream. %s", err.Error())
	}

	ranking2ChartInfo.BinsWidth, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.BinsWidth from stream. %s", err.Error())
	}

	ranking2ChartInfo.Attribute1, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Attribute1 from stream. %s", err.Error())
	}

	ranking2ChartInfo.Attribute2, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Attribute2 from stream. %s", err.Error())
	}

	ranking2ChartInfo.Quantities, err = stream.ReadListUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Quantities from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the Ranking2ChartInfo and returns a byte array
func (ranking2ChartInfo *Ranking2ChartInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteDateTime(ranking2ChartInfo.CreateTime)
	stream.WriteUInt32LE(ranking2ChartInfo.Index)
	stream.WriteUInt32LE(ranking2ChartInfo.Category)
	stream.WriteInt32LE(ranking2ChartInfo.Season)
	stream.WriteUInt8(ranking2ChartInfo.BinsSize)
	stream.WriteUInt8(ranking2ChartInfo.SamplingRate)
	stream.WriteBool(ranking2ChartInfo.ScoreOrder)
	stream.WriteUInt32LE(ranking2ChartInfo.EstimateLength)
	stream.WriteUInt32LE(ranking2ChartInfo.EstimateHighestScore)
	stream.WriteUInt32LE(ranking2ChartInfo.EstimateLowestScore)
	stream.WriteUInt32LE(ranking2ChartInfo.EstimateMedianScore)
	stream.WriteFloat64LE(ranking2ChartInfo.EstimateAverageScore)
	stream.WriteUInt32LE(ranking2ChartInfo.HighestBinsScore)
	stream.WriteUInt32LE(ranking2ChartInfo.LowestBinsScore)
	stream.WriteUInt32LE(ranking2ChartInfo.BinsWidth)
	stream.WriteUInt32LE(ranking2ChartInfo.Attribute1)
	stream.WriteUInt32LE(ranking2ChartInfo.Attribute2)
	stream.WriteListUInt32LE(ranking2ChartInfo.Quantities)

	return stream.Bytes()
}

// Copy returns a new copied instance of Ranking2ChartInfo
func (ranking2ChartInfo *Ranking2ChartInfo) Copy() nex.StructureInterface {
	copied := NewRanking2ChartInfo()

	copied.SetStructureVersion(ranking2ChartInfo.StructureVersion())

	copied.CreateTime = ranking2ChartInfo.CreateTime.Copy()
	copied.Index = ranking2ChartInfo.Index
	copied.Category = ranking2ChartInfo.Category
	copied.Season = ranking2ChartInfo.Season
	copied.BinsSize = ranking2ChartInfo.BinsSize
	copied.SamplingRate = ranking2ChartInfo.SamplingRate
	copied.ScoreOrder = ranking2ChartInfo.ScoreOrder
	copied.EstimateLength = ranking2ChartInfo.EstimateLength
	copied.EstimateHighestScore = ranking2ChartInfo.EstimateHighestScore
	copied.EstimateLowestScore = ranking2ChartInfo.EstimateLowestScore
	copied.EstimateMedianScore = ranking2ChartInfo.EstimateMedianScore
	copied.EstimateAverageScore = ranking2ChartInfo.EstimateAverageScore
	copied.HighestBinsScore = ranking2ChartInfo.HighestBinsScore
	copied.LowestBinsScore = ranking2ChartInfo.LowestBinsScore
	copied.BinsWidth = ranking2ChartInfo.BinsWidth
	copied.Attribute1 = ranking2ChartInfo.Attribute1
	copied.Attribute2 = ranking2ChartInfo.Attribute2
	copied.Quantities = make([]uint32, len(ranking2ChartInfo.Quantities))

	copy(copied.Quantities, ranking2ChartInfo.Quantities)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2ChartInfo *Ranking2ChartInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Ranking2ChartInfo)

	if ranking2ChartInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !ranking2ChartInfo.CreateTime.Equals(other.CreateTime) {
		return false
	}

	if ranking2ChartInfo.Index != other.Index {
		return false
	}

	if ranking2ChartInfo.Category != other.Category {
		return false
	}

	if ranking2ChartInfo.Season != other.Season {
		return false
	}

	if ranking2ChartInfo.BinsSize != other.BinsSize {
		return false
	}

	if ranking2ChartInfo.SamplingRate != other.SamplingRate {
		return false
	}

	if ranking2ChartInfo.ScoreOrder != other.ScoreOrder {
		return false
	}

	if ranking2ChartInfo.EstimateLength != other.EstimateLength {
		return false
	}

	if ranking2ChartInfo.EstimateHighestScore != other.EstimateHighestScore {
		return false
	}

	if ranking2ChartInfo.EstimateLowestScore != other.EstimateLowestScore {
		return false
	}

	if ranking2ChartInfo.EstimateMedianScore != other.EstimateMedianScore {
		return false
	}

	if ranking2ChartInfo.EstimateAverageScore != other.EstimateAverageScore {
		return false
	}

	if ranking2ChartInfo.HighestBinsScore != other.HighestBinsScore {
		return false
	}

	if ranking2ChartInfo.LowestBinsScore != other.LowestBinsScore {
		return false
	}

	if ranking2ChartInfo.BinsWidth != other.BinsWidth {
		return false
	}

	if ranking2ChartInfo.Attribute1 != other.Attribute1 {
		return false
	}

	if ranking2ChartInfo.Attribute2 != other.Attribute2 {
		return false
	}

	if len(ranking2ChartInfo.Quantities) != len(other.Quantities) {
		return false
	}

	for i := 0; i < len(ranking2ChartInfo.Quantities); i++ {
		if ranking2ChartInfo.Quantities[i] != other.Quantities[i] {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (ranking2ChartInfo *Ranking2ChartInfo) String() string {
	return ranking2ChartInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2ChartInfo *Ranking2ChartInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2ChartInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, ranking2ChartInfo.StructureVersion()))

	if ranking2ChartInfo.CreateTime != nil {
		b.WriteString(fmt.Sprintf("%sCreateTime: %s\n", indentationValues, ranking2ChartInfo.CreateTime.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCreateTime: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sIndex: %d,\n", indentationValues, ranking2ChartInfo.Index))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, ranking2ChartInfo.Category))
	b.WriteString(fmt.Sprintf("%sSeason: %d,\n", indentationValues, ranking2ChartInfo.Season))
	b.WriteString(fmt.Sprintf("%sBinsSize: %d,\n", indentationValues, ranking2ChartInfo.BinsSize))
	b.WriteString(fmt.Sprintf("%sSamplingRate: %d,\n", indentationValues, ranking2ChartInfo.SamplingRate))
	b.WriteString(fmt.Sprintf("%sScoreOrder: %t,\n", indentationValues, ranking2ChartInfo.ScoreOrder))
	b.WriteString(fmt.Sprintf("%sEstimateLength: %d,\n", indentationValues, ranking2ChartInfo.EstimateLength))
	b.WriteString(fmt.Sprintf("%sEstimateHighestScore: %d,\n", indentationValues, ranking2ChartInfo.EstimateHighestScore))
	b.WriteString(fmt.Sprintf("%sEstimateLowestScore: %d,\n", indentationValues, ranking2ChartInfo.EstimateLowestScore))
	b.WriteString(fmt.Sprintf("%sEstimateMedianScore: %d,\n", indentationValues, ranking2ChartInfo.EstimateMedianScore))
	b.WriteString(fmt.Sprintf("%sEstimateAverageScore: %f,\n", indentationValues, ranking2ChartInfo.EstimateAverageScore))
	b.WriteString(fmt.Sprintf("%sHighestBinsScore: %d,\n", indentationValues, ranking2ChartInfo.HighestBinsScore))
	b.WriteString(fmt.Sprintf("%sLowestBinsScore: %d,\n", indentationValues, ranking2ChartInfo.LowestBinsScore))
	b.WriteString(fmt.Sprintf("%sBinsWidth: %d,\n", indentationValues, ranking2ChartInfo.BinsWidth))
	b.WriteString(fmt.Sprintf("%sAttribute1: %d,\n", indentationValues, ranking2ChartInfo.Attribute1))
	b.WriteString(fmt.Sprintf("%sAttribute2: %d,\n", indentationValues, ranking2ChartInfo.Attribute2))
	b.WriteString(fmt.Sprintf("%sQuantities: %v,\n", indentationValues, ranking2ChartInfo.Quantities))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2ChartInfo returns a new Ranking2ChartInfo
func NewRanking2ChartInfo() *Ranking2ChartInfo {
	return &Ranking2ChartInfo{}
}
