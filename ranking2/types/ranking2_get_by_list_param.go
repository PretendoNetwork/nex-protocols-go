// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Ranking2GetByListParam holds data for the Ranking 2  protocol
type Ranking2GetByListParam struct {
	nex.Structure
	Category           uint32
	Offset             uint32
	Length             uint32
	SortFlags          uint32
	OptionFlags        uint32
	NumSeasonsToGoBack uint8
}

// ExtractFromStream extracts a Ranking2GetByListParam structure from a stream
func (ranking2GetByListParam *Ranking2GetByListParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	ranking2GetByListParam.Category, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.Category from stream. %s", err.Error())
	}

	ranking2GetByListParam.Offset, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.Offset from stream. %s", err.Error())
	}

	ranking2GetByListParam.Length, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.Length from stream. %s", err.Error())
	}

	ranking2GetByListParam.SortFlags, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.SortFlags from stream. %s", err.Error())
	}

	ranking2GetByListParam.OptionFlags, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.OptionFlags from stream. %s", err.Error())
	}

	ranking2GetByListParam.NumSeasonsToGoBack, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.NumSeasonsToGoBack from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the Ranking2GetByListParam and returns a byte array
func (ranking2GetByListParam *Ranking2GetByListParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(ranking2GetByListParam.Category)
	stream.WriteUInt32LE(ranking2GetByListParam.Offset)
	stream.WriteUInt32LE(ranking2GetByListParam.Length)
	stream.WriteUInt32LE(ranking2GetByListParam.SortFlags)
	stream.WriteUInt32LE(ranking2GetByListParam.OptionFlags)
	stream.WriteUInt8(ranking2GetByListParam.NumSeasonsToGoBack)

	return stream.Bytes()
}

// Copy returns a new copied instance of Ranking2GetByListParam
func (ranking2GetByListParam *Ranking2GetByListParam) Copy() nex.StructureInterface {
	copied := NewRanking2GetByListParam()

	copied.Category = ranking2GetByListParam.Category
	copied.Offset = ranking2GetByListParam.Offset
	copied.Length = ranking2GetByListParam.Length
	copied.SortFlags = ranking2GetByListParam.SortFlags
	copied.OptionFlags = ranking2GetByListParam.OptionFlags
	copied.NumSeasonsToGoBack = ranking2GetByListParam.NumSeasonsToGoBack
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2GetByListParam *Ranking2GetByListParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Ranking2GetByListParam)

	if ranking2GetByListParam.Category != other.Category {
		return false
	}

	if ranking2GetByListParam.Offset != other.Offset {
		return false
	}

	if ranking2GetByListParam.Length != other.Length {
		return false
	}

	if ranking2GetByListParam.SortFlags != other.SortFlags {
		return false
	}

	if ranking2GetByListParam.OptionFlags != other.OptionFlags {
		return false
	}

	if ranking2GetByListParam.NumSeasonsToGoBack != other.NumSeasonsToGoBack {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (ranking2GetByListParam *Ranking2GetByListParam) String() string {
	return ranking2GetByListParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2GetByListParam *Ranking2GetByListParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2GetByListParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, ranking2GetByListParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, ranking2GetByListParam.Category))
	b.WriteString(fmt.Sprintf("%sOffset: %d,\n", indentationValues, ranking2GetByListParam.Offset))
	b.WriteString(fmt.Sprintf("%sLength: %d,\n", indentationValues, ranking2GetByListParam.Length))
	b.WriteString(fmt.Sprintf("%sSortFlags: %d,\n", indentationValues, ranking2GetByListParam.SortFlags))
	b.WriteString(fmt.Sprintf("%sOptionFlags: %d,\n", indentationValues, ranking2GetByListParam.OptionFlags))
	b.WriteString(fmt.Sprintf("%sNumSeasonsToGoBack: %d,\n", indentationValues, ranking2GetByListParam.NumSeasonsToGoBack))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2GetByListParam returns a new Ranking2GetByListParam
func NewRanking2GetByListParam() *Ranking2GetByListParam {
	return &Ranking2GetByListParam{}
}
