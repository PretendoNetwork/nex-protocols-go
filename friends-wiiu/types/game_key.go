package friends_wiiu_types

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// GameKey contains the title ID and version for a title
type GameKey struct {
	nex.Structure
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

	copied.TitleID = gameKey.TitleID
	copied.TitleVersion = gameKey.TitleVersion

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (gameKey *GameKey) Equals(structure nex.StructureInterface) bool {
	other := structure.(*GameKey)

	if gameKey.TitleID != other.TitleID {
		return false
	}

	if gameKey.TitleVersion != other.TitleVersion {
		return false
	}

	return true
}

// NewGameKey returns a new GameKey
func NewGameKey() *GameKey {
	return &GameKey{}
}
