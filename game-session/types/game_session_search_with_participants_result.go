// Package types implements all the types used by the GameSession protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionSearchWithParticipantsResult is a type with the GameSession protocol
type GameSessionSearchWithParticipantsResult struct {
	types.Structure
	ParticipantIDs types.List[types.PID]
}

// WriteTo writes the GameSessionSearchWithParticipantsResult to the given writable
func (gsswpr GameSessionSearchWithParticipantsResult) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsswpr.ParticipantIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsswpr.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionSearchWithParticipantsResult from the given readable
func (gsswpr *GameSessionSearchWithParticipantsResult) ExtractFrom(readable types.Readable) error {
	if err := gsswpr.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSearchWithParticipantsResult header. %s", err.Error())
	}

	if err := gsswpr.ParticipantIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSearchWithParticipantsResult.m_participantIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionSearchWithParticipantsResult
func (gsswpr GameSessionSearchWithParticipantsResult) Copy() types.RVType {
	copied := NewGameSessionSearchWithParticipantsResult()

	copied.StructureVersion = gsswpr.StructureVersion
	copied.ParticipantIDs = gsswpr.ParticipantIDs.Copy().(types.List[types.PID])

	return copied
}

// Equals checks if the given GameSessionSearchWithParticipantsResult contains the same data as the current GameSessionSearchWithParticipantsResult
func (gsswpr GameSessionSearchWithParticipantsResult) Equals(o types.RVType) bool {
	if _, ok := o.(GameSessionSearchWithParticipantsResult); !ok {
		return false
	}

	other := o.(GameSessionSearchWithParticipantsResult)

	if gsswpr.StructureVersion != other.StructureVersion {
		return false
	}

	return gsswpr.ParticipantIDs.Equals(other.ParticipantIDs)
}

// CopyRef copies the current value of the GameSessionSearchWithParticipantsResult
// and returns a pointer to the new copy
func (gsswpr GameSessionSearchWithParticipantsResult) CopyRef() types.RVTypePtr {
	copied := gsswpr.Copy().(GameSessionSearchWithParticipantsResult)
	return &copied
}

// Deref takes a pointer of GameSessionSearchWithParticipantsResult
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsswpr *GameSessionSearchWithParticipantsResult) Deref() types.RVType {
	return *gsswpr
}

// String returns the string representation of the GameSessionSearchWithParticipantsResult
func (gsswpr GameSessionSearchWithParticipantsResult) String() string {
	return gsswpr.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionSearchWithParticipantsResult using the provided indentation level
func (gsswpr GameSessionSearchWithParticipantsResult) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionSearchWithParticipantsResult{\n")
	fmt.Fprintf(&b, "%sm_participantIDs: %s,\n", indentationValues, gsswpr.ParticipantIDs)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionSearchWithParticipantsResult returns a new GameSessionSearchWithParticipantsResult
func NewGameSessionSearchWithParticipantsResult() GameSessionSearchWithParticipantsResult {
	return GameSessionSearchWithParticipantsResult{
		ParticipantIDs: types.NewList[types.PID](),
	}
}
