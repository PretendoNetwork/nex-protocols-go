// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Ranking2ChartInfo holds data for the Ranking 2  protocol
type Ranking2ChartInfo struct {
	types.Structure
	CreateTime           *types.DateTime
	Index                *types.PrimitiveU32
	Category             *types.PrimitiveU32
	Season               *types.PrimitiveS32
	BinsSize             *types.PrimitiveU8
	SamplingRate         *types.PrimitiveU8
	ScoreOrder           *types.PrimitiveBool
	EstimateLength       *types.PrimitiveU32
	EstimateHighestScore *types.PrimitiveU32
	EstimateLowestScore  *types.PrimitiveU32
	EstimateMedianScore  *types.PrimitiveU32
	EstimateAverageScore *types.PrimitiveF64
	HighestBinsScore     *types.PrimitiveU32
	LowestBinsScore      *types.PrimitiveU32
	BinsWidth            *types.PrimitiveU32
	Attribute1           *types.PrimitiveU32
	Attribute2           *types.PrimitiveU32
	Quantities           *types.List[*types.PrimitiveU32]
}

// ExtractFrom extracts the Ranking2ChartInfo from the given readable
func (ranking2ChartInfo *Ranking2ChartInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = ranking2ChartInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Ranking2ChartInfo header. %s", err.Error())
	}

	err = ranking2ChartInfo.CreateTime.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.CreateTime from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.Index.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Index from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Category from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.Season.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Season from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.BinsSize.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.BinsSize from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.SamplingRate.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.SamplingRate from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.ScoreOrder.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.ScoreOrder from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.EstimateLength.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateLength from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.EstimateHighestScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateHighestScore from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.EstimateLowestScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateLowestScore from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.EstimateMedianScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateMedianScore from stream. %s", err.Error())
	}

	ranking2ChartInfo.EstimateAverageScore, err = stream.ReadFloat64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.EstimateAverageScore from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.HighestBinsScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.HighestBinsScore from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.LowestBinsScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.LowestBinsScore from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.BinsWidth.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.BinsWidth from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.Attribute1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Attribute1 from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.Attribute2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Attribute2 from stream. %s", err.Error())
	}

	err = ranking2ChartInfo.Quantities.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2ChartInfo.Quantities from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Ranking2ChartInfo to the given writable
func (ranking2ChartInfo *Ranking2ChartInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ranking2ChartInfo.CreateTime.WriteTo(contentWritable)
	ranking2ChartInfo.Index.WriteTo(contentWritable)
	ranking2ChartInfo.Category.WriteTo(contentWritable)
	ranking2ChartInfo.Season.WriteTo(contentWritable)
	ranking2ChartInfo.BinsSize.WriteTo(contentWritable)
	ranking2ChartInfo.SamplingRate.WriteTo(contentWritable)
	ranking2ChartInfo.ScoreOrder.WriteTo(contentWritable)
	ranking2ChartInfo.EstimateLength.WriteTo(contentWritable)
	ranking2ChartInfo.EstimateHighestScore.WriteTo(contentWritable)
	ranking2ChartInfo.EstimateLowestScore.WriteTo(contentWritable)
	ranking2ChartInfo.EstimateMedianScore.WriteTo(contentWritable)
	ranking2ChartInfo.EstimateAverageScore.WriteTo(contentWritable)
	ranking2ChartInfo.HighestBinsScore.WriteTo(contentWritable)
	ranking2ChartInfo.LowestBinsScore.WriteTo(contentWritable)
	ranking2ChartInfo.BinsWidth.WriteTo(contentWritable)
	ranking2ChartInfo.Attribute1.WriteTo(contentWritable)
	ranking2ChartInfo.Attribute2.WriteTo(contentWritable)
	ranking2ChartInfo.Quantities.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ranking2ChartInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Ranking2ChartInfo
func (ranking2ChartInfo *Ranking2ChartInfo) Copy() types.RVType {
	copied := NewRanking2ChartInfo()

	copied.StructureVersion = ranking2ChartInfo.StructureVersion

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
	copied.Quantities = make(*types.List[*types.PrimitiveU32], len(ranking2ChartInfo.Quantities))

	copy(copied.Quantities, ranking2ChartInfo.Quantities)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2ChartInfo *Ranking2ChartInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2ChartInfo); !ok {
		return false
	}

	other := o.(*Ranking2ChartInfo)

	if ranking2ChartInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !ranking2ChartInfo.CreateTime.Equals(other.CreateTime) {
		return false
	}

	if !ranking2ChartInfo.Index.Equals(other.Index) {
		return false
	}

	if !ranking2ChartInfo.Category.Equals(other.Category) {
		return false
	}

	if !ranking2ChartInfo.Season.Equals(other.Season) {
		return false
	}

	if !ranking2ChartInfo.BinsSize.Equals(other.BinsSize) {
		return false
	}

	if !ranking2ChartInfo.SamplingRate.Equals(other.SamplingRate) {
		return false
	}

	if !ranking2ChartInfo.ScoreOrder.Equals(other.ScoreOrder) {
		return false
	}

	if !ranking2ChartInfo.EstimateLength.Equals(other.EstimateLength) {
		return false
	}

	if !ranking2ChartInfo.EstimateHighestScore.Equals(other.EstimateHighestScore) {
		return false
	}

	if !ranking2ChartInfo.EstimateLowestScore.Equals(other.EstimateLowestScore) {
		return false
	}

	if !ranking2ChartInfo.EstimateMedianScore.Equals(other.EstimateMedianScore) {
		return false
	}

	if !ranking2ChartInfo.EstimateAverageScore.Equals(other.EstimateAverageScore) {
		return false
	}

	if !ranking2ChartInfo.HighestBinsScore.Equals(other.HighestBinsScore) {
		return false
	}

	if !ranking2ChartInfo.LowestBinsScore.Equals(other.LowestBinsScore) {
		return false
	}

	if !ranking2ChartInfo.BinsWidth.Equals(other.BinsWidth) {
		return false
	}

	if !ranking2ChartInfo.Attribute1.Equals(other.Attribute1) {
		return false
	}

	if !ranking2ChartInfo.Attribute2.Equals(other.Attribute2) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, ranking2ChartInfo.StructureVersion))

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
