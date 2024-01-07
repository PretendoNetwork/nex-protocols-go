// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// PlayingSession holds information for a session
type PlayingSession struct {
	types.Structure
	PrincipalID *types.PID
	Gathering   *types.AnyDataHolder
}

// ExtractFrom extracts the PlayingSession from the given readable
func (playingSession *PlayingSession) ExtractFrom(readable types.Readable) error {
	var err error

	if err = playingSession.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read PlayingSession header. %s", err.Error())
	}

	err = playingSession.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PlayingSession.PrincipalID. %s", err.Error())
	}

	err = playingSession.Gathering.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract PlayingSession.Gathering. %s", err.Error())
	}

	return nil
}

// WriteTo writes the PlayingSession to the given writable
func (playingSession *PlayingSession) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	playingSession.PrincipalID.WriteTo(contentWritable)
	playingSession.Gathering.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	playingSession.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of PlayingSession
func (playingSession *PlayingSession) Copy() types.RVType {
	copied := NewPlayingSession()

	copied.StructureVersion = playingSession.StructureVersion

	copied.PrincipalID = playingSession.PrincipalID

	copied.Gathering = playingSession.Gathering.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (playingSession *PlayingSession) Equals(o types.RVType) bool {
	if _, ok := o.(*PlayingSession); !ok {
		return false
	}

	other := o.(*PlayingSession)

	if playingSession.StructureVersion != other.StructureVersion {
		return false
	}

	if !playingSession.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !playingSession.Gathering.Equals(other.Gathering) {
		return false
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, playingSession.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, playingSession.PrincipalID.FormatToString(indentationLevel+1)))

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
