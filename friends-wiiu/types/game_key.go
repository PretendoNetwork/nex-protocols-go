// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// GameKey contains the title ID and version for a title
type GameKey struct {
	nex.Structure
	*nex.Data
	TitleID      uint64
	TitleVersion uint16
}

// Bytes encodes the GameKey and returns a byte array
func (gameKey *GameKey) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(gameKey.TitleID)
	stream.WriteUInt16LE(gameKey.TitleVersion)

	return stream.Bytes()
}

// ExtractFromStream extracts a GameKey structure from a stream
func (gameKey *GameKey) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	gameKey.TitleID, err = stream.ReadUInt64LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GameKey.TitleID. %s", err.Error())
	}

	gameKey.TitleVersion, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract GameKey.TitleVersion. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of GameKey
func (gameKey *GameKey) Copy() nex.StructureInterface {
	copied := NewGameKey()

	copied.SetStructureVersion(gameKey.StructureVersion())

	if gameKey.ParentType() != nil {
		copied.Data = gameKey.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

	copied.SetParentType(copied.Data)

	copied.TitleID = gameKey.TitleID
	copied.TitleVersion = gameKey.TitleVersion

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (gameKey *GameKey) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GameKey)

	if gameKey.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !gameKey.ParentType().Equals(other.ParentType()) {
		return false
	}

	if gameKey.TitleID != other.TitleID {
		return false
	}

	if gameKey.TitleVersion != other.TitleVersion {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, gameKey.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sTitleID: %d,\n", indentationValues, gameKey.TitleID))
	b.WriteString(fmt.Sprintf("%sTitleVersion: %d\n", indentationValues, gameKey.TitleVersion))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGameKey returns a new GameKey
func NewGameKey() *GameKey {
	return &GameKey{}
}
