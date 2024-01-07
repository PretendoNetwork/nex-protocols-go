// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Ranking2GetParam holds data for the Ranking 2  protocol
type Ranking2GetParam struct {
	types.Structure
	NexUniqueID        *types.PrimitiveU64
	PrincipalID        *types.PID
	Category           *types.PrimitiveU32
	Offset             *types.PrimitiveU32
	Length             *types.PrimitiveU32
	SortFlags          *types.PrimitiveU32
	OptionFlags        *types.PrimitiveU32
	Mode               *types.PrimitiveU8
	NumSeasonsToGoBack *types.PrimitiveU8
}

// ExtractFrom extracts the Ranking2GetParam from the given readable
func (ranking2GetParam *Ranking2GetParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = ranking2GetParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Ranking2GetParam header. %s", err.Error())
	}

	err = ranking2GetParam.NexUniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.NexUniqueID from stream. %s", err.Error())
	}

	err = ranking2GetParam.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.PrincipalID from stream. %s", err.Error())
	}

	err = ranking2GetParam.Category.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.Category from stream. %s", err.Error())
	}

	err = ranking2GetParam.Offset.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.Offset from stream. %s", err.Error())
	}

	err = ranking2GetParam.Length.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.Length from stream. %s", err.Error())
	}

	err = ranking2GetParam.SortFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.SortFlags from stream. %s", err.Error())
	}

	err = ranking2GetParam.OptionFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.OptionFlags from stream. %s", err.Error())
	}

	err = ranking2GetParam.Mode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.Mode from stream. %s", err.Error())
	}

	err = ranking2GetParam.NumSeasonsToGoBack.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.NumSeasonsToGoBack from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Ranking2GetParam to the given writable
func (ranking2GetParam *Ranking2GetParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ranking2GetParam.NexUniqueID.WriteTo(contentWritable)
	ranking2GetParam.PrincipalID.WriteTo(contentWritable)
	ranking2GetParam.Category.WriteTo(contentWritable)
	ranking2GetParam.Offset.WriteTo(contentWritable)
	ranking2GetParam.Length.WriteTo(contentWritable)
	ranking2GetParam.SortFlags.WriteTo(contentWritable)
	ranking2GetParam.OptionFlags.WriteTo(contentWritable)
	ranking2GetParam.Mode.WriteTo(contentWritable)
	ranking2GetParam.NumSeasonsToGoBack.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ranking2GetParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Ranking2GetParam
func (ranking2GetParam *Ranking2GetParam) Copy() types.RVType {
	copied := NewRanking2GetParam()

	copied.StructureVersion = ranking2GetParam.StructureVersion

	copied.NexUniqueID = ranking2GetParam.NexUniqueID
	copied.PrincipalID = ranking2GetParam.PrincipalID.Copy()
	copied.Category = ranking2GetParam.Category
	copied.Offset = ranking2GetParam.Offset
	copied.Length = ranking2GetParam.Length
	copied.SortFlags = ranking2GetParam.SortFlags
	copied.OptionFlags = ranking2GetParam.OptionFlags
	copied.Mode = ranking2GetParam.Mode
	copied.NumSeasonsToGoBack = ranking2GetParam.NumSeasonsToGoBack
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2GetParam *Ranking2GetParam) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2GetParam); !ok {
		return false
	}

	other := o.(*Ranking2GetParam)

	if ranking2GetParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !ranking2GetParam.NexUniqueID.Equals(other.NexUniqueID) {
		return false
	}

	if !ranking2GetParam.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !ranking2GetParam.Category.Equals(other.Category) {
		return false
	}

	if !ranking2GetParam.Offset.Equals(other.Offset) {
		return false
	}

	if !ranking2GetParam.Length.Equals(other.Length) {
		return false
	}

	if !ranking2GetParam.SortFlags.Equals(other.SortFlags) {
		return false
	}

	if !ranking2GetParam.OptionFlags.Equals(other.OptionFlags) {
		return false
	}

	if !ranking2GetParam.Mode.Equals(other.Mode) {
		return false
	}

	if !ranking2GetParam.NumSeasonsToGoBack.Equals(other.NumSeasonsToGoBack) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (ranking2GetParam *Ranking2GetParam) String() string {
	return ranking2GetParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2GetParam *Ranking2GetParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2GetParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, ranking2GetParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sNexUniqueID: %d,\n", indentationValues, ranking2GetParam.NexUniqueID))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, ranking2GetParam.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sCategory: %d,\n", indentationValues, ranking2GetParam.Category))
	b.WriteString(fmt.Sprintf("%sOffset: %d,\n", indentationValues, ranking2GetParam.Offset))
	b.WriteString(fmt.Sprintf("%sLength: %d,\n", indentationValues, ranking2GetParam.Length))
	b.WriteString(fmt.Sprintf("%sSortFlags: %d,\n", indentationValues, ranking2GetParam.SortFlags))
	b.WriteString(fmt.Sprintf("%sOptionFlags: %d,\n", indentationValues, ranking2GetParam.OptionFlags))
	b.WriteString(fmt.Sprintf("%sMode: %d,\n", indentationValues, ranking2GetParam.Mode))
	b.WriteString(fmt.Sprintf("%sNumSeasonsToGoBack: %d,\n", indentationValues, ranking2GetParam.NumSeasonsToGoBack))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2GetParam returns a new Ranking2GetParam
func NewRanking2GetParam() *Ranking2GetParam {
	return &Ranking2GetParam{}
}
