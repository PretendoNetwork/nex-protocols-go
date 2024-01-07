// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Ranking2RankData holds data for the Ranking 2  protocol
type Ranking2RankData struct {
	types.Structure
	Misc        *types.PrimitiveU64
	NexUniqueID *types.PrimitiveU64
	PrincipalID *types.PID
	Rank        *types.PrimitiveU32
	Score       *types.PrimitiveU32
	CommonData  *Ranking2CommonData
}

// ExtractFrom extracts the Ranking2RankData from the given readable
func (ranking2RankData *Ranking2RankData) ExtractFrom(readable types.Readable) error {
	var err error

	if err = ranking2RankData.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Ranking2RankData header. %s", err.Error())
	}

	err = ranking2RankData.Misc.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.Misc from stream. %s", err.Error())
	}

	err = ranking2RankData.NexUniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.NexUniqueID from stream. %s", err.Error())
	}

	err = ranking2RankData.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.PrincipalID from stream. %s", err.Error())
	}

	err = ranking2RankData.Rank.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.Rank from stream. %s", err.Error())
	}

	err = ranking2RankData.Score.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.Score from stream. %s", err.Error())
	}

	err = ranking2RankData.CommonData.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2RankData.CommonData from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Ranking2RankData to the given writable
func (ranking2RankData *Ranking2RankData) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	ranking2RankData.Misc.WriteTo(contentWritable)
	ranking2RankData.NexUniqueID.WriteTo(contentWritable)
	ranking2RankData.PrincipalID.WriteTo(contentWritable)
	ranking2RankData.Rank.WriteTo(contentWritable)
	ranking2RankData.Score.WriteTo(contentWritable)
	ranking2RankData.CommonData.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	ranking2RankData.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Ranking2RankData
func (ranking2RankData *Ranking2RankData) Copy() types.RVType {
	copied := NewRanking2RankData()

	copied.StructureVersion = ranking2RankData.StructureVersion

	copied.Misc = ranking2RankData.Misc
	copied.NexUniqueID = ranking2RankData.NexUniqueID
	copied.PrincipalID = ranking2RankData.PrincipalID.Copy()
	copied.Rank = ranking2RankData.Rank
	copied.Score = ranking2RankData.Score
	copied.CommonData = ranking2RankData.CommonData.Copy().(*Ranking2CommonData)
	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (ranking2RankData *Ranking2RankData) Equals(o types.RVType) bool {
	if _, ok := o.(*Ranking2RankData); !ok {
		return false
	}

	other := o.(*Ranking2RankData)

	if ranking2RankData.StructureVersion != other.StructureVersion {
		return false
	}

	if !ranking2RankData.Misc.Equals(other.Misc) {
		return false
	}

	if !ranking2RankData.NexUniqueID.Equals(other.NexUniqueID) {
		return false
	}

	if !ranking2RankData.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !ranking2RankData.Rank.Equals(other.Rank) {
		return false
	}

	if !ranking2RankData.Score.Equals(other.Score) {
		return false
	}

	if !ranking2RankData.CommonData.Equals(other.CommonData) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (ranking2RankData *Ranking2RankData) String() string {
	return ranking2RankData.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (ranking2RankData *Ranking2RankData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Ranking2RankData{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, ranking2RankData.StructureVersion))
	b.WriteString(fmt.Sprintf("%sMisc: %d,\n", indentationValues, ranking2RankData.Misc))
	b.WriteString(fmt.Sprintf("%sNexUniqueID: %d,\n", indentationValues, ranking2RankData.NexUniqueID))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, ranking2RankData.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sRank: %d,\n", indentationValues, ranking2RankData.Rank))
	b.WriteString(fmt.Sprintf("%sScore: %d,\n", indentationValues, ranking2RankData.Score))

	if ranking2RankData.CommonData != nil {
		b.WriteString(fmt.Sprintf("%sCommonData: %s\n", indentationValues, ranking2RankData.CommonData.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCommonData: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRanking2RankData returns a new Ranking2RankData
func NewRanking2RankData() *Ranking2RankData {
	return &Ranking2RankData{}
}
