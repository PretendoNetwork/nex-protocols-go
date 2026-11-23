// Package types implements all the types used by the GameSession protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameSession is a type with the GameSession protocol
type GameSession struct {
	types.Structure
	TypeID     types.UInt32
	Attributes types.List[Property]
}

// WriteTo writes the GameSession to the given writable
func (gs GameSession) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gs.TypeID.WriteTo(contentWritable)
	gs.Attributes.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gs.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameSession from the given readable
func (gs *GameSession) ExtractFrom(readable types.Readable) error {
	if err := gs.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSession header. %s", err.Error())
	}

	if err := gs.TypeID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSession.m_typeID. %s", err.Error())
	}

	if err := gs.Attributes.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract GameSession.m_attributes. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameSession
func (gs GameSession) Copy() types.RVType {
	copied := NewGameSession()

	copied.StructureVersion = gs.StructureVersion
	copied.TypeID = gs.TypeID
	copied.Attributes = gs.Attributes.Copy().(types.List[Property])

	return copied
}

// Equals checks if the given GameSession contains the same data as the current GameSession
func (gs GameSession) Equals(o types.RVType) bool {
	if _, ok := o.(GameSession); !ok {
		return false
	}

	other := o.(GameSession)

	if gs.StructureVersion != other.StructureVersion {
		return false
	}

	if gs.TypeID != other.TypeID {
		return false
	}

	return gs.Attributes.Equals(other.Attributes)
}

// CopyRef copies the current value of the GameSession
// and returns a pointer to the new copy
func (gs GameSession) CopyRef() types.RVTypePtr {
	copied := gs.Copy().(GameSession)
	return &copied
}

// Deref takes a pointer of GameSession
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gs *GameSession) Deref() types.RVType {
	return *gs
}

// String returns the string representation of the GameSession
func (gs GameSession) String() string {
	return gs.FormatToString(0)
}

// FormatToString pretty-prints the GameSession using the provided indentation level
func (gs GameSession) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameSession{\n")
	fmt.Fprintf(&b, "%sm_typeID: %s, \n", indentationValues, gs.TypeID)
	fmt.Fprintf(&b, "%sm_attributes: %s,\n", indentationValues, gs.Attributes)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewGameSession returns a new GameSession
func NewGameSession() GameSession {
	return GameSession{
		TypeID:     types.NewUInt32(0),
		Attributes: types.NewList[Property](),
	}
}
