// Package types implements all the types used by the GameSession protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSessionQuery is a type with the GameSession protocol
type GameSessionQuery struct {
	types.Structure
	TypeID     types.UInt32
	QueryID    types.UInt32
	Parameters types.List[Property]
}

// WriteTo writes the GameSessionQuery to the given writable
func (gsq GameSessionQuery) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gsq.TypeID.WriteTo(contentWritable)
	gsq.QueryID.WriteTo(contentWritable)
	gsq.Parameters.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gsq.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSessionQuery from the given readable
func (gsq *GameSessionQuery) ExtractFrom(readable types.Readable) error {
	if err := gsq.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionQuery header. %s", err.Error())
	}

	if err := gsq.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionQuery.m_typeID. %s", err.Error())
	}

	if err := gsq.QueryID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionQuery.m_queryID. %s", err.Error())
	}

	if err := gsq.Parameters.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSessionQuery.m_parameters. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSessionQuery
func (gsq GameSessionQuery) Copy() types.RVType {
	copied := NewGameSessionQuery()

	copied.StructureVersion = gsq.StructureVersion
	copied.TypeID = gsq.TypeID
	copied.QueryID = gsq.QueryID
	copied.Parameters = gsq.Parameters.Copy().(types.List[Property])

	return copied
}

// Equals checks if the given GameSessionQuery contains the same data as the current GameSessionQuery
func (gsq GameSessionQuery) Equals(o types.RVType) bool {
	if _, ok := o.(GameSessionQuery); !ok {
		return false
	}

	other := o.(GameSessionQuery)

	if gsq.StructureVersion != other.StructureVersion {
		return false
	}

	if gsq.TypeID != other.TypeID {
		return false
	}

	if gsq.QueryID != other.QueryID {
		return false
	}

	return gsq.Parameters.Equals(other.Parameters)
}

// CopyRef copies the current value of the GameSessionQuery
// and returns a pointer to the new copy
func (gsq GameSessionQuery) CopyRef() types.RVTypePtr {
	copied := gsq.Copy().(GameSessionQuery)
	return &copied
}

// Deref takes a pointer of GameSessionQuery
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gsq *GameSessionQuery) Deref() types.RVType {
	return *gsq
}

// String returns the string representation of the GameSessionQuery
func (gsq GameSessionQuery) String() string {
	return gsq.FormatToString(0)
}

// FormatToString pretty-prints the GameSessionQuery using the provided indentation level
func (gsq GameSessionQuery) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSessionQuery{\n")
	fmt.Fprintf(&b, "%sm_typeID: %s, \n", indentationValues, gsq.TypeID)
	fmt.Fprintf(&b, "%sm_queryID: %s, \n", indentationValues, gsq.QueryID)
	fmt.Fprintf(&b, "%sm_parameters: %s,\n", indentationValues, gsq.Parameters)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSessionQuery returns a new GameSessionQuery
func NewGameSessionQuery() GameSessionQuery {
	return GameSessionQuery{
		TypeID:     types.NewUInt32(0),
		QueryID:    types.NewUInt32(0),
		Parameters: types.NewList[Property](),
	}
}
