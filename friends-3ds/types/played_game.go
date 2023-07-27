// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// PlayedGame is a data structure used by the Friends 3DS protocol to hold information about a friends Mii
type PlayedGame struct {
	nex.Structure
	GameKey *GameKey
	Unknown *nex.DateTime
}

// Bytes encodes the PlayedGame and returns a byte array
func (playedGame *PlayedGame) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(playedGame.GameKey)
	stream.WriteDateTime(playedGame.Unknown)

	return stream.Bytes()
}

// Copy returns a new copied instance of PlayedGame
func (playedGame *PlayedGame) Copy() nex.StructureInterface {
	copied := NewPlayedGame()

	copied.GameKey = playedGame.GameKey.Copy().(*GameKey)
	copied.Unknown = playedGame.Unknown.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (playedGame *PlayedGame) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PlayedGame)

	if !playedGame.GameKey.Equals(other.GameKey) {
		return false
	}

	if !playedGame.Unknown.Equals(other.Unknown) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (playedGame *PlayedGame) String() string {
	return playedGame.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (playedGame *PlayedGame) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PlayedGame{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, playedGame.StructureVersion()))

	if playedGame.Unknown != nil {
		b.WriteString(fmt.Sprintf("%sGameKey: %s\n", indentationValues, playedGame.GameKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sGameKey: nil\n", indentationValues))
	}

	if playedGame.Unknown != nil {
		b.WriteString(fmt.Sprintf("%sUnknown: %s\n", indentationValues, playedGame.Unknown.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUnknown: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPlayedGame returns a new PlayedGame
func NewPlayedGame() *PlayedGame {
	return &PlayedGame{}
}
