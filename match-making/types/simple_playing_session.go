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

// SimplePlayingSession holds simple information for a session
type SimplePlayingSession struct {
	types.Structure
	PrincipalID *types.PID
	GatheringID *types.PrimitiveU32
	GameMode    *types.PrimitiveU32
	Attribute0  *types.PrimitiveU32
}

// ExtractFrom extracts the SimplePlayingSession from the given readable
func (simplePlayingSession *SimplePlayingSession) ExtractFrom(readable types.Readable) error {
	var err error

	if err = simplePlayingSession.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read SimplePlayingSession header. %s", err.Error())
	}

	err = simplePlayingSession.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession.PrincipalID. %s", err.Error())
	}

	err = simplePlayingSession.GatheringID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession.GatheringID. %s", err.Error())
	}

	err = simplePlayingSession.GameMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession.GameMode. %s", err.Error())
	}

	err = simplePlayingSession.Attribute0.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SimplePlayingSession.Attribute0. %s", err.Error())
	}

	return nil
}

// WriteTo writes the SimplePlayingSession to the given writable
func (simplePlayingSession *SimplePlayingSession) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	simplePlayingSession.PrincipalID.WriteTo(contentWritable)
	simplePlayingSession.GatheringID.WriteTo(contentWritable)
	simplePlayingSession.GameMode.WriteTo(contentWritable)
	simplePlayingSession.Attribute0.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	simplePlayingSession.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of SimplePlayingSession
func (simplePlayingSession *SimplePlayingSession) Copy() types.RVType {
	copied := NewSimplePlayingSession()

	copied.StructureVersion = simplePlayingSession.StructureVersion

	copied.PrincipalID = simplePlayingSession.PrincipalID
	copied.GatheringID = simplePlayingSession.GatheringID
	copied.GameMode = simplePlayingSession.GameMode
	copied.Attribute0 = simplePlayingSession.Attribute0

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (simplePlayingSession *SimplePlayingSession) Equals(o types.RVType) bool {
	if _, ok := o.(*SimplePlayingSession); !ok {
		return false
	}

	other := o.(*SimplePlayingSession)

	if simplePlayingSession.StructureVersion != other.StructureVersion {
		return false
	}

	if !simplePlayingSession.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	if !simplePlayingSession.GatheringID.Equals(other.GatheringID) {
		return false
	}

	if !simplePlayingSession.GameMode.Equals(other.GameMode) {
		return false
	}

	if !simplePlayingSession.Attribute0.Equals(other.Attribute0) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, simplePlayingSession.StructureVersion))
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
