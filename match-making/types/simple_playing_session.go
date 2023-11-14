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

// SimplePlayingSession holds simple information for a session
type SimplePlayingSession struct {
	nex.Structure
	PrincipalID *nex.PID
	GatheringID uint32
	GameMode    uint32
	Attribute0  uint32
}

// ExtractFromStream extracts a SimplePlayingSession structure from a stream
func (simplePlayingSession *SimplePlayingSession) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	simplePlayingSession.PrincipalID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession.PrincipalID. %s", err.Error())
	}

	simplePlayingSession.GatheringID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession.GatheringID. %s", err.Error())
	}

	simplePlayingSession.GameMode, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession.GameMode. %s", err.Error())
	}

	simplePlayingSession.Attribute0, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession.Attribute0. %s", err.Error())
	}

	return nil
}

// Bytes encodes the SimplePlayingSession and returns a byte array
func (simplePlayingSession *SimplePlayingSession) Bytes(stream *nex.StreamOut) []byte {
	stream.WritePID(simplePlayingSession.PrincipalID)
	stream.WriteUInt32LE(simplePlayingSession.GatheringID)
	stream.WriteUInt32LE(simplePlayingSession.GameMode)
	stream.WriteUInt32LE(simplePlayingSession.Attribute0)

	return stream.Bytes()
}

// Copy returns a new copied instance of SimplePlayingSession
func (simplePlayingSession *SimplePlayingSession) Copy() nex.StructureInterface {
	copied := NewSimplePlayingSession()

	copied.SetStructureVersion(simplePlayingSession.StructureVersion())

	copied.PrincipalID = simplePlayingSession.PrincipalID
	copied.GatheringID = simplePlayingSession.GatheringID
	copied.GameMode = simplePlayingSession.GameMode
	copied.Attribute0 = simplePlayingSession.Attribute0

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simplePlayingSession *SimplePlayingSession) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SimplePlayingSession)

	if simplePlayingSession.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !simplePlayingSession.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if simplePlayingSession.GatheringID != other.GatheringID {
		return false
	}

	if simplePlayingSession.GameMode != other.GameMode {
		return false
	}

	if simplePlayingSession.Attribute0 != other.Attribute0 {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (simplePlayingSession *SimplePlayingSession) String() string {
	return simplePlayingSession.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (simplePlayingSession *SimplePlayingSession) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SimplePlayingSession{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, simplePlayingSession.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, simplePlayingSession.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sGatheringID: %d,\n", indentationValues, simplePlayingSession.GatheringID))
	b.WriteString(fmt.Sprintf("%sGameMode: %d,\n", indentationValues, simplePlayingSession.GameMode))
	b.WriteString(fmt.Sprintf("%sAttribute0: %d\n", indentationValues, simplePlayingSession.Attribute0))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSimplePlayingSession returns a new SimplePlayingSession
func NewSimplePlayingSession() *SimplePlayingSession {
	return &SimplePlayingSession{}
}
