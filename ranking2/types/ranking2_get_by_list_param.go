// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Ranking2GetByListParam holds data for the Ranking 2  protocol
type Ranking2GetByListParam struct {
	types.Structure
	Category           *types.PrimitiveU32
	Offset             *types.PrimitiveU32
	Length             *types.PrimitiveU32
	SortFlags          *types.PrimitiveU32
	OptionFlags        *types.PrimitiveU32
	NumSeasonsToGoBack *types.PrimitiveU8
}

// ExtractFrom extracts the Ranking2GetByListParam from the given readable
func (ranking2GetByListParam *Ranking2GetByListParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = ranking2GetByListParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Ranking2GetByListParam header. %s", err.Error())
	}

	err = ranking2GetByListParam.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.Category from stream. %s", err.Error())
	}

	err = ranking2GetByListParam.Offset.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.Offset from stream. %s", err.Error())
	}

	err = ranking2GetByListParam.Length.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.Length from stream. %s", err.Error())
	}

	err = ranking2GetByListParam.SortFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.SortFlags from stream. %s", err.Error())
	}

	err = ranking2GetByListParam.OptionFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.OptionFlags from stream. %s", err.Error())
	}

	err = ranking2GetByListParam.NumSeasonsToGoBack.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetByListParam.NumSeasonsToGoBack from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Ranking2GetByListParam to the given writable
func (ranking2GetByListParam *Ranking2GetByListParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ranking2GetByListParam.Category.WriteTo(contentWritable)
	ranking2GetByListParam.Offset.WriteTo(contentWritable)
	ranking2GetByListParam.Length.WriteTo(contentWritable)
	ranking2GetByListParam.SortFlags.WriteTo(contentWritable)
	ranking2GetByListParam.OptionFlags.WriteTo(contentWritable)
	ranking2GetByListParam.NumSeasonsToGoBack.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ranking2GetByListParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Ranking2GetByListParam
func (ranking2GetByListParam *Ranking2GetByListParam) Copy() types.RVType {
	copied := NewRanking2GetByListParam()

	copied.StructureVersion = ranking2GetByListParam.StructureVersion

	copied.Category = ranking2GetByListParam.Category
	copied.Offset = ranking2GetByListParam.Offset
	copied.Length = ranking2GetByListParam.Length
	copied.SortFlags = ranking2GetByListParam.SortFlags
	copied.OptionFlags = ranking2GetByListParam.OptionFlags
	copied.NumSeasonsToGoBack = ranking2GetByListParam.NumSeasonsToGoBack
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2GetByListParam *Ranking2GetByListParam) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2GetByListParam); !ok {
		return false
	}

	other := o.(*Ranking2GetByListParam)

	if ranking2GetByListParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !ranking2GetByListParam.Category.Equals(other.Category) {
		return false
	}

	if !ranking2GetByListParam.Offset.Equals(other.Offset) {
		return false
	}

	if !ranking2GetByListParam.Length.Equals(other.Length) {
		return false
	}

	if !ranking2GetByListParam.SortFlags.Equals(other.SortFlags) {
		return false
	}

	if !ranking2GetByListParam.OptionFlags.Equals(other.OptionFlags) {
		return false
	}

	if !ranking2GetByListParam.NumSeasonsToGoBack.Equals(other.NumSeasonsToGoBack) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, ranking2GetByListParam.StructureVersion))
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
