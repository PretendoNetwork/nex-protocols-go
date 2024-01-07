// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// GameKey contains the title ID and version for a title
type GameKey struct {
	types.Structure
	*types.Data
	TitleID      *types.PrimitiveU64
	TitleVersion *types.PrimitiveU16
}

// WriteTo writes the GameKey to the given writable
func (gameKey *GameKey) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gameKey.TitleID.WriteTo(contentWritable)
	gameKey.TitleVersion.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gameKey.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the GameKey from the given readable
func (gameKey *GameKey) ExtractFrom(readable types.Readable) error {
	var err error

	if err = gameKey.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read GameKey header. %s", err.Error())
	}

	err = gameKey.TitleID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GameKey.TitleID. %s", err.Error())
	}

	err = gameKey.TitleVersion.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract GameKey.TitleVersion. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameKey
func (gameKey *GameKey) Copy() types.RVType {
	copied := NewGameKey()

	copied.StructureVersion = gameKey.StructureVersion

	copied.Data = gameKey.Data.Copy().(*types.Data)

	copied.TitleID = gameKey.TitleID
	copied.TitleVersion = gameKey.TitleVersion

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (gameKey *GameKey) Equals(o types.RVType) bool {
	if _, ok := o.(*GameKey); !ok {
		return false
	}

	other := o.(*GameKey)

	if gameKey.StructureVersion != other.StructureVersion {
		return false
	}

	if !gameKey.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !gameKey.TitleID.Equals(other.TitleID) {
		return false
	}

	if !gameKey.TitleVersion.Equals(other.TitleVersion) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (gameKey *GameKey) String() string {
	return gameKey.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (gameKey *GameKey) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("GameKey{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, gameKey.StructureVersion))
	b.WriteString(fmt.Sprintf("%sTitleID: %d,\n", indentationValues, gameKey.TitleID))
	b.WriteString(fmt.Sprintf("%sTitleVersion: %d\n", indentationValues, gameKey.TitleVersion))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGameKey returns a new GameKey
func NewGameKey() *GameKey {
	return &GameKey{}
}
