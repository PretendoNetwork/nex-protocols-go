// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Ranking2CategorySetting holds data for the Ranking 2  protocol
type Ranking2CategorySetting struct {
	types.Structure
	MinScore           *types.PrimitiveU32
	MaxScore           *types.PrimitiveU32
	LowestRank         *types.PrimitiveU32
	ResetMonth         *types.PrimitiveU16
	ResetDay           *types.PrimitiveU8
	ResetHour          *types.PrimitiveU8
	ResetMode          *types.PrimitiveU8
	MaxSeasonsToGoBack *types.PrimitiveU8
	ScoreOrder         *types.PrimitiveBool
}

// ExtractFrom extracts the Ranking2CategorySetting from the given readable
func (ranking2CategorySetting *Ranking2CategorySetting) ExtractFrom(readable types.Readable) error {
	var err error

	if err = ranking2CategorySetting.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Ranking2CategorySetting header. %s", err.Error())
	}

	err = ranking2CategorySetting.MinScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.MinScore from stream. %s", err.Error())
	}

	err = ranking2CategorySetting.MaxScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.MaxScore from stream. %s", err.Error())
	}

	err = ranking2CategorySetting.LowestRank.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.LowestRank from stream. %s", err.Error())
	}

	err = ranking2CategorySetting.ResetMonth.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ResetMonth from stream. %s", err.Error())
	}

	err = ranking2CategorySetting.ResetDay.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ResetDay from stream. %s", err.Error())
	}

	err = ranking2CategorySetting.ResetHour.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ResetHour from stream. %s", err.Error())
	}

	err = ranking2CategorySetting.ResetMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ResetMode from stream. %s", err.Error())
	}

	err = ranking2CategorySetting.MaxSeasonsToGoBack.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.MaxSeasonsToGoBack from stream. %s", err.Error())
	}

	err = ranking2CategorySetting.ScoreOrder.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ScoreOrder from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Ranking2CategorySetting to the given writable
func (ranking2CategorySetting *Ranking2CategorySetting) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ranking2CategorySetting.MinScore.WriteTo(contentWritable)
	ranking2CategorySetting.MaxScore.WriteTo(contentWritable)
	ranking2CategorySetting.LowestRank.WriteTo(contentWritable)
	ranking2CategorySetting.ResetMonth.WriteTo(contentWritable)
	ranking2CategorySetting.ResetDay.WriteTo(contentWritable)
	ranking2CategorySetting.ResetHour.WriteTo(contentWritable)
	ranking2CategorySetting.ResetMode.WriteTo(contentWritable)
	ranking2CategorySetting.MaxSeasonsToGoBack.WriteTo(contentWritable)
	ranking2CategorySetting.ScoreOrder.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ranking2CategorySetting.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Ranking2CategorySetting
func (ranking2CategorySetting *Ranking2CategorySetting) Copy() types.RVType {
	copied := NewRanking2CategorySetting()

	copied.StructureVersion = ranking2CategorySetting.StructureVersion

	copied.MinScore = ranking2CategorySetting.MinScore
	copied.MaxScore = ranking2CategorySetting.MaxScore
	copied.LowestRank = ranking2CategorySetting.LowestRank
	copied.ResetMonth = ranking2CategorySetting.ResetMonth
	copied.ResetDay = ranking2CategorySetting.ResetDay
	copied.ResetHour = ranking2CategorySetting.ResetHour
	copied.ResetMode = ranking2CategorySetting.ResetMode
	copied.MaxSeasonsToGoBack = ranking2CategorySetting.MaxSeasonsToGoBack
	copied.ScoreOrder = ranking2CategorySetting.ScoreOrder
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2CategorySetting *Ranking2CategorySetting) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2CategorySetting); !ok {
		return false
	}

	other := o.(*Ranking2CategorySetting)

	if ranking2CategorySetting.StructureVersion != other.StructureVersion {
		return false
	}

	if !ranking2CategorySetting.MinScore.Equals(other.MinScore) {
		return false
	}

	if !ranking2CategorySetting.MaxScore.Equals(other.MaxScore) {
		return false
	}

	if !ranking2CategorySetting.LowestRank.Equals(other.LowestRank) {
		return false
	}

	if !ranking2CategorySetting.ResetMonth.Equals(other.ResetMonth) {
		return false
	}

	if !ranking2CategorySetting.ResetDay.Equals(other.ResetDay) {
		return false
	}

	if !ranking2CategorySetting.ResetHour.Equals(other.ResetHour) {
		return false
	}

	if !ranking2CategorySetting.ResetMode.Equals(other.ResetMode) {
		return false
	}

	if !ranking2CategorySetting.MaxSeasonsToGoBack.Equals(other.MaxSeasonsToGoBack) {
		return false
	}

	if !ranking2CategorySetting.ScoreOrder.Equals(other.ScoreOrder) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (ranking2CategorySetting *Ranking2CategorySetting) String() string {
	return ranking2CategorySetting.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2CategorySetting *Ranking2CategorySetting) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2CategorySetting{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, ranking2CategorySetting.StructureVersion))
	b.WriteString(fmt.Sprintf("%sMinScore: %d,\n", indentationValues, ranking2CategorySetting.MinScore))
	b.WriteString(fmt.Sprintf("%sMaxScore: %d,\n", indentationValues, ranking2CategorySetting.MaxScore))
	b.WriteString(fmt.Sprintf("%sLowestRank: %d,\n", indentationValues, ranking2CategorySetting.LowestRank))
	b.WriteString(fmt.Sprintf("%sResetMonth: %d,\n", indentationValues, ranking2CategorySetting.ResetMonth))
	b.WriteString(fmt.Sprintf("%sResetDay: %d,\n", indentationValues, ranking2CategorySetting.ResetDay))
	b.WriteString(fmt.Sprintf("%sResetHour: %d,\n", indentationValues, ranking2CategorySetting.ResetHour))
	b.WriteString(fmt.Sprintf("%sResetMode: %d,\n", indentationValues, ranking2CategorySetting.ResetMode))
	b.WriteString(fmt.Sprintf("%sMaxSeasonsToGoBack: %d,\n", indentationValues, ranking2CategorySetting.MaxSeasonsToGoBack))
	b.WriteString(fmt.Sprintf("%sScoreOrder: %t,\n", indentationValues, ranking2CategorySetting.ScoreOrder))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2CategorySetting returns a new Ranking2CategorySetting
func NewRanking2CategorySetting() *Ranking2CategorySetting {
	return &Ranking2CategorySetting{}
}
