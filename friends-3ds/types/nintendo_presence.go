package friends_3ds_types

import (
	"bytes"
	"fmt"

	"github.com/PretendoNetwork/nex-go"
)

// NintendoPresence contains information about a users online presence
type NintendoPresence struct {
	nex.Structure
	ChangedFlags      uint32
	GameKey           *GameKey
	Message           string
	JoinAvailableFlag uint32
	MatchmakeType     uint8
	JoinGameID        uint32
	JoinGameMode      uint32
	OwnerPID          uint32
	JoinGroupID       uint32
	ApplicationArg    []byte
}

// Bytes encodes the NintendoPresence and returns a byte array
func (presence *NintendoPresence) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(presence.ChangedFlags)
	stream.WriteStructure(presence.GameKey)
	stream.WriteString(presence.Message)
	stream.WriteUInt32LE(presence.JoinAvailableFlag)
	stream.WriteUInt8(presence.MatchmakeType)
	stream.WriteUInt32LE(presence.JoinGameID)
	stream.WriteUInt32LE(presence.JoinGameMode)
	stream.WriteUInt32LE(presence.OwnerPID)
	stream.WriteUInt32LE(presence.JoinGroupID)
	stream.WriteBuffer(presence.ApplicationArg)

	return stream.Bytes()
}

// ExtractFromStream extracts a NintendoPresence structure from a stream
func (presence *NintendoPresence) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	presence.ChangedFlags, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.ChangedFlags. %s", err.Error())
	}

	gameKey, err := stream.ReadStructure(NewGameKey())
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.GameKey. %s", err.Error())
	}

	presence.GameKey = gameKey.(*GameKey)
	presence.Message, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.Message. %s", err.Error())
	}

	presence.JoinAvailableFlag, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.JoinAvailableFlag. %s", err.Error())
	}

	presence.MatchmakeType, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.MatchmakeType. %s", err.Error())
	}

	presence.JoinGameID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.JoinGameID. %s", err.Error())
	}

	presence.JoinGameMode, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.JoinGameMode. %s", err.Error())
	}

	presence.OwnerPID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresence.OwnerPID. %s", err.Error())
	}

	presence.JoinGroupID, err = stream.ReadUInt32LE()
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
func (presence *NintendoPresence) Copy() nex.StructureInterface {
	copied := NewNintendoPresence()

	copied.ChangedFlags = presence.ChangedFlags
	copied.GameKey = presence.GameKey.Copy().(*GameKey)
	copied.Message = presence.Message
	copied.JoinAvailableFlag = presence.JoinAvailableFlag
	copied.MatchmakeType = presence.MatchmakeType
	copied.JoinGameID = presence.JoinGameID
	copied.JoinGameMode = presence.JoinGameMode
	copied.OwnerPID = presence.OwnerPID
	copied.JoinGroupID = presence.JoinGroupID
	copied.ApplicationArg = make([]byte, len(presence.ApplicationArg))

	copy(copied.ApplicationArg, presence.ApplicationArg)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (presence *NintendoPresence) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NintendoPresence)

	if presence.ChangedFlags != other.ChangedFlags {
		return false
	}

	if !presence.GameKey.Equals(other.GameKey) {
		return false
	}

	if presence.Message != other.Message {
		return false
	}

	if presence.JoinAvailableFlag != other.JoinAvailableFlag {
		return false
	}

	if presence.MatchmakeType != other.MatchmakeType {
		return false
	}

	if presence.JoinGameID != other.JoinGameID {
		return false
	}

	if presence.JoinGameMode != other.JoinGameMode {
		return false
	}

	if presence.OwnerPID != other.OwnerPID {
		return false
	}

	if presence.JoinGroupID != other.JoinGroupID {
		return false
	}

	if !bytes.Equal(presence.ApplicationArg, other.ApplicationArg) {
		return false
	}

	return true
}

// NewNintendoPresence returns a new NintendoPresence
func NewNintendoPresence() *NintendoPresence {
	return &NintendoPresence{}
}
