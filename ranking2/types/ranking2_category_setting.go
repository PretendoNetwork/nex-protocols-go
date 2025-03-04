// Package types implements all the types used by the Ranking2 protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2CategorySetting is a type within the Ranking2 protocol
type Ranking2CategorySetting struct {
	types.Structure
	MinScore           types.UInt32
	MaxScore           types.UInt32
	LowestRank         types.UInt32
	ResetMonth         types.UInt16
	ResetDay           types.UInt8
	ResetHour          types.UInt8
	ResetMode          types.UInt8
	MaxSeasonsToGoBack types.UInt8
	ScoreOrder         types.Bool
}

// WriteTo writes the Ranking2CategorySetting to the given writable
func (rcs Ranking2CategorySetting) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rcs.MinScore.WriteTo(contentWritable)
	rcs.MaxScore.WriteTo(contentWritable)
	rcs.LowestRank.WriteTo(contentWritable)
	rcs.ResetMonth.WriteTo(contentWritable)
	rcs.ResetDay.WriteTo(contentWritable)
	rcs.ResetHour.WriteTo(contentWritable)
	rcs.ResetMode.WriteTo(contentWritable)
	rcs.MaxSeasonsToGoBack.WriteTo(contentWritable)
	rcs.ScoreOrder.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rcs.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Ranking2CategorySetting from the given readable
func (rcs *Ranking2CategorySetting) ExtractFrom(readable types.Readable) error {
	var err error

	err = rcs.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting header. %s", err.Error())
	}

	err = rcs.MinScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.MinScore. %s", err.Error())
	}

	err = rcs.MaxScore.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.MaxScore. %s", err.Error())
	}

	err = rcs.LowestRank.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.LowestRank. %s", err.Error())
	}

	err = rcs.ResetMonth.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ResetMonth. %s", err.Error())
	}

	err = rcs.ResetDay.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ResetDay. %s", err.Error())
	}

	err = rcs.ResetHour.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ResetHour. %s", err.Error())
	}

	err = rcs.ResetMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ResetMode. %s", err.Error())
	}

	err = rcs.MaxSeasonsToGoBack.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.MaxSeasonsToGoBack. %s", err.Error())
	}

	err = rcs.ScoreOrder.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ScoreOrder. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Ranking2CategorySetting
func (rcs Ranking2CategorySetting) Copy() types.RVType {
	copied := NewRanking2CategorySetting()

	copied.StructureVersion = rcs.StructureVersion
	copied.MinScore = rcs.MinScore.Copy().(types.UInt32)
	copied.MaxScore = rcs.MaxScore.Copy().(types.UInt32)
	copied.LowestRank = rcs.LowestRank.Copy().(types.UInt32)
	copied.ResetMonth = rcs.ResetMonth.Copy().(types.UInt16)
	copied.ResetDay = rcs.ResetDay.Copy().(types.UInt8)
	copied.ResetHour = rcs.ResetHour.Copy().(types.UInt8)
	copied.ResetMode = rcs.ResetMode.Copy().(types.UInt8)
	copied.MaxSeasonsToGoBack = rcs.MaxSeasonsToGoBack.Copy().(types.UInt8)
	copied.ScoreOrder = rcs.ScoreOrder.Copy().(types.Bool)

	return copied
}

// Equals checks if the given Ranking2CategorySetting contains the same data as the current Ranking2CategorySetting
func (rcs Ranking2CategorySetting) Equals(o types.RVType) bool {
	if _, ok := o.(Ranking2CategorySetting); !ok {
		return false
	}

	other := o.(Ranking2CategorySetting)

	if rcs.StructureVersion != other.StructureVersion {
		return false
	}

	if !rcs.MinScore.Equals(other.MinScore) {
		return false
	}

	if !rcs.MaxScore.Equals(other.MaxScore) {
		return false
	}

	if !rcs.LowestRank.Equals(other.LowestRank) {
		return false
	}

	if !rcs.ResetMonth.Equals(other.ResetMonth) {
		return false
	}

	if !rcs.ResetDay.Equals(other.ResetDay) {
		return false
	}

	if !rcs.ResetHour.Equals(other.ResetHour) {
		return false
	}

	if !rcs.ResetMode.Equals(other.ResetMode) {
		return false
	}

	if !rcs.MaxSeasonsToGoBack.Equals(other.MaxSeasonsToGoBack) {
		return false
	}

	return rcs.ScoreOrder.Equals(other.ScoreOrder)
}

// CopyRef copies the current value of the Ranking2CategorySetting
// and returns a pointer to the new copy
func (rcs Ranking2CategorySetting) CopyRef() types.RVTypePtr {
	copied := rcs.Copy().(Ranking2CategorySetting)
	return &copied
}

// Deref takes a pointer to the Ranking2CategorySetting
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rcs *Ranking2CategorySetting) Deref() types.RVType {
	return *rcs
}

// String returns the string representation of the Ranking2CategorySetting
func (rcs Ranking2CategorySetting) String() string {
	return rcs.FormatToString(0)
}

// FormatToString pretty-prints the Ranking2CategorySetting using the provided indentation level
func (rcs Ranking2CategorySetting) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2CategorySetting{\n")
	b.WriteString(fmt.Sprintf("%sMinScore: %s,\n", indentationValues, rcs.MinScore))
	b.WriteString(fmt.Sprintf("%sMaxScore: %s,\n", indentationValues, rcs.MaxScore))
	b.WriteString(fmt.Sprintf("%sLowestRank: %s,\n", indentationValues, rcs.LowestRank))
	b.WriteString(fmt.Sprintf("%sResetMonth: %s,\n", indentationValues, rcs.ResetMonth))
	b.WriteString(fmt.Sprintf("%sResetDay: %s,\n", indentationValues, rcs.ResetDay))
	b.WriteString(fmt.Sprintf("%sResetHour: %s,\n", indentationValues, rcs.ResetHour))
	b.WriteString(fmt.Sprintf("%sResetMode: %s,\n", indentationValues, rcs.ResetMode))
	b.WriteString(fmt.Sprintf("%sMaxSeasonsToGoBack: %s,\n", indentationValues, rcs.MaxSeasonsToGoBack))
	b.WriteString(fmt.Sprintf("%sScoreOrder: %s,\n", indentationValues, rcs.ScoreOrder))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2CategorySetting returns a new Ranking2CategorySetting
func NewRanking2CategorySetting() Ranking2CategorySetting {
	return Ranking2CategorySetting{
		MinScore:           types.NewUInt32(0),
		MaxScore:           types.NewUInt32(0),
		LowestRank:         types.NewUInt32(0),
		ResetMonth:         types.NewUInt16(0),
		ResetDay:           types.NewUInt8(0),
		ResetHour:          types.NewUInt8(0),
		ResetMode:          types.NewUInt8(0),
		MaxSeasonsToGoBack: types.NewUInt8(0),
		ScoreOrder:         types.NewBool(false),
	}

}
