// Package types implements all the types used by the MatchmakeReferee protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeRefereeRound is a type within the MatchmakeReferee protocol
type MatchmakeRefereeRound struct {
	types.Structure
	types.Data
	RoundID                        types.UInt64
	GID                            types.UInt32
	State                          types.UInt32
	PersonalDataCategory           types.UInt32
	NormalizedPersonalRoundResults types.List[MatchmakeRefereePersonalRoundResult]
}

// WriteTo writes the MatchmakeRefereeRound to the given writable
func (mrr MatchmakeRefereeRound) WriteTo(writable types.Writable) {
	mrr.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	mrr.RoundID.WriteTo(contentWritable)
	mrr.GID.WriteTo(contentWritable)
	mrr.State.WriteTo(contentWritable)
	mrr.PersonalDataCategory.WriteTo(contentWritable)
	mrr.NormalizedPersonalRoundResults.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	mrr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the MatchmakeRefereeRound from the given readable
func (mrr *MatchmakeRefereeRound) ExtractFrom(readable types.Readable) error {
	var err error

	err = mrr.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.Data. %s", err.Error())
	}

	err = mrr.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound header. %s", err.Error())
	}

	err = mrr.RoundID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.RoundID. %s", err.Error())
	}

	err = mrr.GID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.GID. %s", err.Error())
	}

	err = mrr.State.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.State. %s", err.Error())
	}

	err = mrr.PersonalDataCategory.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.PersonalDataCategory. %s", err.Error())
	}

	err = mrr.NormalizedPersonalRoundResults.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract MatchmakeRefereeRound.NormalizedPersonalRoundResults. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of MatchmakeRefereeRound
func (mrr MatchmakeRefereeRound) Copy() types.RVType {
	copied := NewMatchmakeRefereeRound()

	copied.StructureVersion = mrr.StructureVersion
	copied.Data = mrr.Data.Copy().(types.Data)
	copied.RoundID = mrr.RoundID.Copy().(types.UInt64)
	copied.GID = mrr.GID.Copy().(types.UInt32)
	copied.State = mrr.State.Copy().(types.UInt32)
	copied.PersonalDataCategory = mrr.PersonalDataCategory.Copy().(types.UInt32)
	copied.NormalizedPersonalRoundResults = mrr.NormalizedPersonalRoundResults.Copy().(types.List[MatchmakeRefereePersonalRoundResult])

	return copied
}

// Equals checks if the given MatchmakeRefereeRound contains the same data as the current MatchmakeRefereeRound
func (mrr MatchmakeRefereeRound) Equals(o types.RVType) bool {
	if _, ok := o.(MatchmakeRefereeRound); !ok {
		return false
	}

	other := o.(MatchmakeRefereeRound)

	if mrr.StructureVersion != other.StructureVersion {
		return false
	}

	if !mrr.Data.Equals(other.Data) {
		return false
	}

	if !mrr.RoundID.Equals(other.RoundID) {
		return false
	}

	if !mrr.GID.Equals(other.GID) {
		return false
	}

	if !mrr.State.Equals(other.State) {
		return false
	}

	if !mrr.PersonalDataCategory.Equals(other.PersonalDataCategory) {
		return false
	}

	return mrr.NormalizedPersonalRoundResults.Equals(other.NormalizedPersonalRoundResults)
}

// CopyRef copies the current value of the MatchmakeRefereeRound
// and returns a pointer to the new copy
func (mrr MatchmakeRefereeRound) CopyRef() types.RVTypePtr {
	copied := mrr.Copy().(MatchmakeRefereeRound)
	return &copied
}

// Deref takes a pointer to the MatchmakeRefereeRound
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (mrr *MatchmakeRefereeRound) Deref() types.RVType {
	return *mrr
}

// String returns the string representation of the MatchmakeRefereeRound
func (mrr MatchmakeRefereeRound) String() string {
	return mrr.FormatToString(0)
}

// FormatToString pretty-prints the MatchmakeRefereeRound using the provided indentation level
func (mrr MatchmakeRefereeRound) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("MatchmakeRefereeRound{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, mrr.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sRoundID: %s,\n", indentationValues, mrr.RoundID))
	b.WriteString(fmt.Sprintf("%sGID: %s,\n", indentationValues, mrr.GID))
	b.WriteString(fmt.Sprintf("%sState: %s,\n", indentationValues, mrr.State))
	b.WriteString(fmt.Sprintf("%sPersonalDataCategory: %s,\n", indentationValues, mrr.PersonalDataCategory))
	b.WriteString(fmt.Sprintf("%sNormalizedPersonalRoundResults: %s,\n", indentationValues, mrr.NormalizedPersonalRoundResults))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewMatchmakeRefereeRound returns a new MatchmakeRefereeRound
func NewMatchmakeRefereeRound() MatchmakeRefereeRound {
	return MatchmakeRefereeRound{
		Data:                           types.NewData(),
		RoundID:                        types.NewUInt64(0),
		GID:                            types.NewUInt32(0),
		State:                          types.NewUInt32(0),
		PersonalDataCategory:           types.NewUInt32(0),
		NormalizedPersonalRoundResults: types.NewList[MatchmakeRefereePersonalRoundResult](),
	}

}
