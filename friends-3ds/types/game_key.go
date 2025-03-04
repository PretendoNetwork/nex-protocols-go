// Package types implements all the types used by the Friends3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// GameKey is a type within the Friends3DS protocol
type GameKey struct {
	types.Structure
	types.Data
	TitleID      types.UInt64
	TitleVersion types.UInt16
}

// ObjectID returns the object identifier of the type
func (gk GameKey) ObjectID() types.RVType {
	return gk.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (gk GameKey) DataObjectID() types.RVType {
	return types.NewString("GameKey")
}

// WriteTo writes the GameKey to the given writable
func (gk GameKey) WriteTo(writable types.Writable) {
	gk.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	gk.TitleID.WriteTo(contentWritable)
	gk.TitleVersion.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gk.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameKey from the given readable
func (gk *GameKey) ExtractFrom(readable types.Readable) error {
	var err error

	err = gk.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GameKey.Data. %s", err.Error())
	}

	err = gk.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GameKey header. %s", err.Error())
	}

	err = gk.TitleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GameKey.TitleID. %s", err.Error())
	}

	err = gk.TitleVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GameKey.TitleVersion. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameKey
func (gk GameKey) Copy() types.RVType {
	copied := NewGameKey()

	copied.StructureVersion = gk.StructureVersion
	copied.Data = gk.Data.Copy().(types.Data)
	copied.TitleID = gk.TitleID.Copy().(types.UInt64)
	copied.TitleVersion = gk.TitleVersion.Copy().(types.UInt16)

	return copied
}

// Equals checks if the given GameKey contains the same data as the current GameKey
func (gk GameKey) Equals(o types.RVType) bool {
	if _, ok := o.(GameKey); !ok {
		return false
	}

	other := o.(GameKey)

	if gk.StructureVersion != other.StructureVersion {
		return false
	}

	if !gk.Data.Equals(other.Data) {
		return false
	}

	if !gk.TitleID.Equals(other.TitleID) {
		return false
	}

	return gk.TitleVersion.Equals(other.TitleVersion)
}

// CopyRef copies the current value of the GameKey
// and returns a pointer to the new copy
func (gk GameKey) CopyRef() types.RVTypePtr {
	copied := gk.Copy().(GameKey)
	return &copied
}

// Deref takes a pointer to the GameKey
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (gk *GameKey) Deref() types.RVType {
	return *gk
}

// String returns the string representation of the GameKey
func (gk GameKey) String() string {
	return gk.FormatToString(0)
}

// FormatToString pretty-prints the GameKey using the provided indentation level
func (gk GameKey) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameKey{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, gk.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sTitleID: %s,\n", indentationValues, gk.TitleID))
	b.WriteString(fmt.Sprintf("%sTitleVersion: %s,\n", indentationValues, gk.TitleVersion))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGameKey returns a new GameKey
func NewGameKey() GameKey {
	return GameKey{
		Data:         types.NewData(),
		TitleID:      types.NewUInt64(0),
		TitleVersion: types.NewUInt16(0),
	}

}
