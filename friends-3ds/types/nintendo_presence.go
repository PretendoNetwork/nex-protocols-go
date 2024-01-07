// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// NintendoPresence contains information about a users online presence
type NintendoPresence struct {
	types.Structure
	*types.Data
	ChangedFlags      *types.PrimitiveU32
	GameKey           *GameKey
	Message           string
	JoinAvailableFlag *types.PrimitiveU32
	MatchmakeType     *types.PrimitiveU8
	JoinGameID        *types.PrimitiveU32
	JoinGameMode      *types.PrimitiveU32
	OwnerPID          *types.PID
	JoinGroupID       *types.PrimitiveU32
	ApplicationArg    []byte
}

// WriteTo writes the NintendoPresence to the given writable
func (presence *NintendoPresence) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	presence.ChangedFlags.WriteTo(contentWritable)
	presence.GameKey.WriteTo(contentWritable)
	presence.Message.WriteTo(contentWritable)
	presence.JoinAvailableFlag.WriteTo(contentWritable)
	presence.MatchmakeType.WriteTo(contentWritable)
	presence.JoinGameID.WriteTo(contentWritable)
	presence.JoinGameMode.WriteTo(contentWritable)
	presence.OwnerPID.WriteTo(contentWritable)
	presence.JoinGroupID.WriteTo(contentWritable)
	stream.WriteBuffer(presence.ApplicationArg)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NintendoPresence from the given readable
func (presence *NintendoPresence) ExtractFrom(readable types.Readable) error {
	var err error

	if err = presence.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read NintendoPresence header. %s", err.Error())
	}

	err = presence.ChangedFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.ChangedFlags. %s", err.Error())
	}

	err = presence.GameKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.GameKey. %s", err.Error())
	}

	err = presence.Message.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.Message. %s", err.Error())
	}

	err = presence.JoinAvailableFlag.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.JoinAvailableFlag. %s", err.Error())
	}

	err = presence.MatchmakeType.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.MatchmakeType. %s", err.Error())
	}

	err = presence.JoinGameID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.JoinGameID. %s", err.Error())
	}

	err = presence.JoinGameMode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.JoinGameMode. %s", err.Error())
	}

	err = presence.OwnerPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.OwnerPID. %s", err.Error())
	}

	err = presence.JoinGroupID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.JoinGroupID. %s", err.Error())
	}

	presence.ApplicationArg, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.ApplicationArg. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoPresence
func (presence *NintendoPresence) Copy() types.RVType {
	copied := NewNintendoPresence()

	copied.StructureVersion = presence.StructureVersion

	copied.Data = presence.Data.Copy().(*types.Data)

	copied.ChangedFlags = presence.ChangedFlags
	copied.GameKey = presence.GameKey.Copy().(*GameKey)
	copied.Message = presence.Message
	copied.JoinAvailableFlag = presence.JoinAvailableFlag
	copied.MatchmakeType = presence.MatchmakeType
	copied.JoinGameID = presence.JoinGameID
	copied.JoinGameMode = presence.JoinGameMode
	copied.OwnerPID = presence.OwnerPID.Copy()
	copied.JoinGroupID = presence.JoinGroupID
	copied.ApplicationArg = make([]byte, len(presence.ApplicationArg))

	copy(copied.ApplicationArg, presence.ApplicationArg)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (presence *NintendoPresence) Equals(o types.RVType) bool {
	if _, ok := o.(*NintendoPresence); !ok {
		return false
	}

	other := o.(*NintendoPresence)

	if presence.StructureVersion != other.StructureVersion {
		return false
	}

	if !presence.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !presence.ChangedFlags.Equals(other.ChangedFlags) {
		return false
	}

	if !presence.GameKey.Equals(other.GameKey) {
		return false
	}

	if !presence.Message.Equals(other.Message) {
		return false
	}

	if !presence.JoinAvailableFlag.Equals(other.JoinAvailableFlag) {
		return false
	}

	if !presence.MatchmakeType.Equals(other.MatchmakeType) {
		return false
	}

	if !presence.JoinGameID.Equals(other.JoinGameID) {
		return false
	}

	if !presence.JoinGameMode.Equals(other.JoinGameMode) {
		return false
	}

	if !presence.OwnerPID.Equals(other.OwnerPID) {
		return false
	}

	if !presence.JoinGroupID.Equals(other.JoinGroupID) {
		return false
	}

	if !presence.ApplicationArg.Equals(other.ApplicationArg) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (presence *NintendoPresence) String() string {
	return presence.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (presence *NintendoPresence) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("NintendoPresence{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, presence.StructureVersion))
	b.WriteString(fmt.Sprintf("%sChangedFlags: %d,\n", indentationValues, presence.ChangedFlags))

	if presence.GameKey != nil {
		b.WriteString(fmt.Sprintf("%sGameKey: %s,\n", indentationValues, presence.GameKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sGameKey: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sMessage: %q,\n", indentationValues, presence.Message))
	b.WriteString(fmt.Sprintf("%sJoinAvailableFlag: %d,\n", indentationValues, presence.JoinAvailableFlag))
	b.WriteString(fmt.Sprintf("%sMatchmakeType: %d,\n", indentationValues, presence.MatchmakeType))
	b.WriteString(fmt.Sprintf("%sJoinGameID: %d,\n", indentationValues, presence.JoinGameID))
	b.WriteString(fmt.Sprintf("%sJoinGameMode: %d,\n", indentationValues, presence.JoinGameMode))
	b.WriteString(fmt.Sprintf("%sOwnerPID: %s,\n", indentationValues, presence.OwnerPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sJoinGroupID: %d,\n", indentationValues, presence.JoinGroupID))
	b.WriteString(fmt.Sprintf("%sApplicationArg: %x\n", indentationValues, presence.ApplicationArg))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoPresence returns a new NintendoPresence
func NewNintendoPresence() *NintendoPresence {
	return &NintendoPresence{}
}
