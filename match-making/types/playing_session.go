// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// PlayingSession holds information for a session
type PlayingSession struct {
	nex.Structure
	PrincipalID uint32
	Gathering   *nex.DataHolder
}

// ExtractFromStream extracts a PlayingSession structure from a stream
func (playingSession *PlayingSession) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	playingSession.PrincipalID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract PlayingSession.PrincipalID. %s", err.Error())
	}

	playingSession.Gathering, err = stream.ReadDataHolder()
	if err != nil {
		return fmt.Errorf("Failed to extract PlayingSession.Gathering. %s", err.Error())
	}

	return nil
}

// Bytes encodes the PlayingSession and returns a byte array
func (playingSession *PlayingSession) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(playingSession.PrincipalID)
	stream.WriteDataHolder(playingSession.Gathering)

	return stream.Bytes()
}

// Copy returns a new copied instance of PlayingSession
func (playingSession *PlayingSession) Copy() nex.StructureInterface {
	copied := NewPlayingSession()

	copied.PrincipalID = playingSession.PrincipalID

	if playingSession.Gathering != nil {
		copied.Gathering = playingSession.Gathering.Copy()
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (playingSession *PlayingSession) Equals(structure nex.StructureInterface) bool {
	other := structure.(*PlayingSession)

	if playingSession.PrincipalID != other.PrincipalID {
		return false
	}

	if playingSession.Gathering != nil && other.Gathering == nil {
		return false
	}

	if playingSession.Gathering == nil && other.Gathering != nil {
		return false
	}

	if playingSession.Gathering != nil && other.Gathering != nil {
		if !playingSession.Gathering.Equals(other.Gathering) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (playingSession *PlayingSession) String() string {
	return playingSession.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (playingSession *PlayingSession) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PlayingSession{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, playingSession.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %d,\n", indentationValues, playingSession.PrincipalID))

	if playingSession.Gathering != nil {
		b.WriteString(fmt.Sprintf("%sGathering: %s\n", indentationValues, playingSession.Gathering.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sGathering: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewPlayingSession returns a new PlayingSession
func NewPlayingSession() *PlayingSession {
	return &PlayingSession{}
}
