// Package types implements all the types used by the GameSession protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionSocialQuery is a type with the GameSession protocol
type GameSessionSocialQuery struct {
	types.Structure
	TypeID         types.UInt32
	QueryID        types.UInt32
	Parameters     types.List[Property]
	ParticipantIDs types.List[types.PID]
}

// WriteTo writes the GameSessionSocialQuery to the given writable
func (gssq GameSessionSocialQuery) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gssq.TypeID.WriteTo(contentWritable)
	gssq.QueryID.WriteTo(contentWritable)
	gssq.Parameters.WriteTo(contentWritable)
	gssq.ParticipantIDs.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gssq.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionSocialQuery from the given readable
func (gssq *GameSessionSocialQuery) ExtractFrom(readable types.Readable) error {
	if err := gssq.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSocialQuery header. %s", err.Error())
	}

	if err := gssq.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSocialQuery.m_typeID. %s", err.Error())
	}

	if err := gssq.QueryID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSocialQuery.m_queryID. %s", err.Error())
	}

	if err := gssq.Parameters.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSocialQuery.m_parameters. %s", err.Error())
	}

	if err := gssq.ParticipantIDs.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionSocialQuery.m_participantIDs. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionSocialQuery
func (gssq GameSessionSocialQuery) Copy() types.RVType {
	copied := NewGameSessionSocialQuery()

	copied.StructureVersion = gssq.StructureVersion
	copied.TypeID = gssq.TypeID
	copied.QueryID = gssq.QueryID
	copied.Parameters = gssq.Parameters.Copy().(types.List[Property])
	copied.ParticipantIDs = gssq.ParticipantIDs.Copy().(types.List[types.PID])

	return copied
}

// Equals checks if the given GameSessionSocialQuery contains the same data as the current GameSessionSocialQuery
func (gssq GameSessionSocialQuery) Equals(o types.RVType) bool {
	if _, ok := o.(GameSessionSocialQuery); !ok {
		return false
	}

	other := o.(GameSessionSocialQuery)

	if gssq.StructureVersion != other.StructureVersion {
		return false
	}

	if gssq.TypeID != other.TypeID {
		return false
	}

	if gssq.QueryID != other.QueryID {
		return false
	}

	if !gssq.Parameters.Equals(other.Parameters) {
		return false
	}

	return gssq.ParticipantIDs.Equals(other.ParticipantIDs)
}

// CopyRef copies the current value of the GameSessionSocialQuery
// and returns a pointer to the new copy
func (gssq GameSessionSocialQuery) CopyRef() types.RVTypePtr {
	copied := gssq.Copy().(GameSessionSocialQuery)
	return &copied
}

// Deref takes a pointer of GameSessionSocialQuery
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gssq *GameSessionSocialQuery) Deref() types.RVType {
	return *gssq
}

// String returns the string representation of the GameSessionSocialQuery
func (gssq GameSessionSocialQuery) String() string {
	return gssq.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionSocialQuery using the provided indentation level
func (gssq GameSessionSocialQuery) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionSocialQuery{\n")
	fmt.Fprintf(&b, "%sm_typeID: %s, \n", indentationValues, gssq.TypeID)
	fmt.Fprintf(&b, "%sm_queryID: %s, \n", indentationValues, gssq.QueryID)
	fmt.Fprintf(&b, "%sm_parameters: %s,\n", indentationValues, gssq.Parameters)
	fmt.Fprintf(&b, "%sm_participantIDs: %s,\n", indentationValues, gssq.ParticipantIDs)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionSocialQuery returns a new GameSessionSocialQuery
func NewGameSessionSocialQuery() GameSessionSocialQuery {
	return GameSessionSocialQuery{
		TypeID:         types.NewUInt32(0),
		QueryID:        types.NewUInt32(0),
		Parameters:     types.NewList[Property](),
		ParticipantIDs: types.NewList[types.PID](),
	}
}
