// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Ranking2CategorySetting holds data for the Ranking 2  protocol
type Ranking2CategorySetting struct {
	nex.Structure
	MinScore           uint32
	MaxScore           uint32
	LowestRank         uint32
	ResetMonth         uint16
	ResetDay           uint8
	ResetHour          uint8
	ResetMode          uint8
	MaxSeasonsToGoBack uint8
	ScoreOrder         bool
}

// ExtractFromStream extracts a Ranking2CategorySetting structure from a stream
func (ranking2CategorySetting *Ranking2CategorySetting) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	ranking2CategorySetting.MinScore, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.MinScore from stream. %s", err.Error())
	}

	ranking2CategorySetting.MaxScore, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.MaxScore from stream. %s", err.Error())
	}

	ranking2CategorySetting.LowestRank, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.LowestRank from stream. %s", err.Error())
	}

	ranking2CategorySetting.ResetMonth, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ResetMonth from stream. %s", err.Error())
	}

	ranking2CategorySetting.ResetDay, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ResetDay from stream. %s", err.Error())
	}

	ranking2CategorySetting.ResetHour, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ResetHour from stream. %s", err.Error())
	}

	ranking2CategorySetting.ResetMode, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ResetMode from stream. %s", err.Error())
	}

	ranking2CategorySetting.MaxSeasonsToGoBack, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.MaxSeasonsToGoBack from stream. %s", err.Error())
	}

	ranking2CategorySetting.ScoreOrder, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2CategorySetting.ScoreOrder from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the Ranking2CategorySetting and returns a byte array
func (ranking2CategorySetting *Ranking2CategorySetting) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(ranking2CategorySetting.MinScore)
	stream.WriteUInt32LE(ranking2CategorySetting.MaxScore)
	stream.WriteUInt32LE(ranking2CategorySetting.LowestRank)
	stream.WriteUInt16LE(ranking2CategorySetting.ResetMonth)
	stream.WriteUInt8(ranking2CategorySetting.ResetDay)
	stream.WriteUInt8(ranking2CategorySetting.ResetHour)
	stream.WriteUInt8(ranking2CategorySetting.ResetMode)
	stream.WriteUInt8(ranking2CategorySetting.MaxSeasonsToGoBack)
	stream.WriteBool(ranking2CategorySetting.ScoreOrder)

	return stream.Bytes()
}

// Copy returns a new copied instance of Ranking2CategorySetting
func (ranking2CategorySetting *Ranking2CategorySetting) Copy() nex.StructureInterface {
	copied := NewRanking2CategorySetting()

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
func (ranking2CategorySetting *Ranking2CategorySetting) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Ranking2CategorySetting)

	if ranking2CategorySetting.MinScore != other.MinScore {
		return false
	}

	if ranking2CategorySetting.MaxScore != other.MaxScore {
		return false
	}

	if ranking2CategorySetting.LowestRank != other.LowestRank {
		return false
	}

	if ranking2CategorySetting.ResetMonth != other.ResetMonth {
		return false
	}

	if ranking2CategorySetting.ResetDay != other.ResetDay {
		return false
	}

	if ranking2CategorySetting.ResetHour != other.ResetHour {
		return false
	}

	if ranking2CategorySetting.ResetMode != other.ResetMode {
		return false
	}

	if ranking2CategorySetting.MaxSeasonsToGoBack != other.MaxSeasonsToGoBack {
		return false
	}

	if ranking2CategorySetting.ScoreOrder != other.ScoreOrder {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, ranking2CategorySetting.StructureVersion()))
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
