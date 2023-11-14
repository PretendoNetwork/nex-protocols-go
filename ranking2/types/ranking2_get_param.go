// Package types implements all the types used by the Ranking 2  protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Ranking2GetParam holds data for the Ranking 2  protocol
type Ranking2GetParam struct {
	nex.Structure
	NexUniqueID        uint64
	PrincipalID        *nex.PID
	Category           uint32
	Offset             uint32
	Length             uint32
	SortFlags          uint32
	OptionFlags        uint32
	Mode               uint8
	NumSeasonsToGoBack uint8
}

// ExtractFromStream extracts a Ranking2GetParam structure from a stream
func (ranking2GetParam *Ranking2GetParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	ranking2GetParam.NexUniqueID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.NexUniqueID from stream. %s", err.Error())
	}

	ranking2GetParam.PrincipalID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.PrincipalID from stream. %s", err.Error())
	}

	ranking2GetParam.Category, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.Category from stream. %s", err.Error())
	}

	ranking2GetParam.Offset, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.Offset from stream. %s", err.Error())
	}

	ranking2GetParam.Length, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.Length from stream. %s", err.Error())
	}

	ranking2GetParam.SortFlags, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.SortFlags from stream. %s", err.Error())
	}

	ranking2GetParam.OptionFlags, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.OptionFlags from stream. %s", err.Error())
	}

	ranking2GetParam.Mode, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.Mode from stream. %s", err.Error())
	}

	ranking2GetParam.NumSeasonsToGoBack, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Ranking2GetParam.NumSeasonsToGoBack from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the Ranking2GetParam and returns a byte array
func (ranking2GetParam *Ranking2GetParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(ranking2GetParam.NexUniqueID)
	stream.WritePID(ranking2GetParam.PrincipalID)
	stream.WriteUInt32LE(ranking2GetParam.Category)
	stream.WriteUInt32LE(ranking2GetParam.Offset)
	stream.WriteUInt32LE(ranking2GetParam.Length)
	stream.WriteUInt32LE(ranking2GetParam.SortFlags)
	stream.WriteUInt32LE(ranking2GetParam.OptionFlags)
	stream.WriteUInt8(ranking2GetParam.Mode)
	stream.WriteUInt8(ranking2GetParam.NumSeasonsToGoBack)

	return stream.Bytes()
}

// Copy returns a new copied instance of Ranking2GetParam
func (ranking2GetParam *Ranking2GetParam) Copy() nex.StructureInterface {
	copied := NewRanking2GetParam()

	copied.SetStructureVersion(ranking2GetParam.StructureVersion())

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
func (ranking2GetParam *Ranking2GetParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Ranking2GetParam)

	if ranking2GetParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if ranking2GetParam.NexUniqueID != other.NexUniqueID {
		return false
	}

	if !ranking2GetParam.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if ranking2GetParam.Category != other.Category {
		return false
	}

	if ranking2GetParam.Offset != other.Offset {
		return false
	}

	if ranking2GetParam.Length != other.Length {
		return false
	}

	if ranking2GetParam.SortFlags != other.SortFlags {
		return false
	}

	if ranking2GetParam.OptionFlags != other.OptionFlags {
		return false
	}

	if ranking2GetParam.Mode != other.Mode {
		return false
	}

	if ranking2GetParam.NumSeasonsToGoBack != other.NumSeasonsToGoBack {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, ranking2GetParam.StructureVersion()))
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
