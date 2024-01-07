// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// PlayedGame is a data structure used by the Friends 3DS protocol to hold information about a friends Mii
type PlayedGame struct {
	types.Structure
	GameKey *GameKey
	Unknown *types.DateTime
}

// WriteTo writes the PlayedGame to the given writable
func (playedGame *PlayedGame) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	playedGame.GameKey.WriteTo(contentWritable)
	playedGame.Unknown.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	playedGame.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of PlayedGame
func (playedGame *PlayedGame) Copy() types.RVType {
	copied := NewPlayedGame()

	copied.StructureVersion = playedGame.StructureVersion

	copied.GameKey = playedGame.GameKey.Copy().(*GameKey)
	copied.Unknown = playedGame.Unknown.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (playedGame *PlayedGame) Equals(o types.RVType) bool {
	if _, ok := o.(*PlayedGame); !ok {
		return false
	}

	other := o.(*PlayedGame)

	if playedGame.StructureVersion != other.StructureVersion {
		return false
	}

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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, playedGame.StructureVersion))

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
