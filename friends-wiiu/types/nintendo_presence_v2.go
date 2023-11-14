// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// NintendoPresenceV2 contains information about a users online presence
type NintendoPresenceV2 struct {
	nex.Structure
	*nex.Data
	ChangedFlags    uint32
	Online          bool
	GameKey         *GameKey
	Unknown1        uint8
	Message         string
	Unknown2        uint32
	Unknown3        uint8
	GameServerID    uint32
	Unknown4        uint32
	PID             *nex.PID
	GatheringID     uint32
	ApplicationData []byte
	Unknown5        uint8
	Unknown6        uint8
	Unknown7        uint8
}

// Bytes encodes the NintendoPresenceV2 and returns a byte array
func (presence *NintendoPresenceV2) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(presence.ChangedFlags)
	stream.WriteBool(presence.Online)
	stream.WriteStructure(presence.GameKey)
	stream.WriteUInt8(presence.Unknown1)
	stream.WriteString(presence.Message)
	stream.WriteUInt32LE(presence.Unknown2)
	stream.WriteUInt8(presence.Unknown3)
	stream.WriteUInt32LE(presence.GameServerID)
	stream.WriteUInt32LE(presence.Unknown4)
	stream.WritePID(presence.PID)
	stream.WriteUInt32LE(presence.GatheringID)
	stream.WriteBuffer(presence.ApplicationData)
	stream.WriteUInt8(presence.Unknown5)
	stream.WriteUInt8(presence.Unknown6)
	stream.WriteUInt8(presence.Unknown7)

	return stream.Bytes()
}

// ExtractFromStream extracts a NintendoPresenceV2 structure from a stream
func (presence *NintendoPresenceV2) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	presence.ChangedFlags, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.ChangedFlags. %s", err.Error())
	}

	presence.Online, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Online. %s", err.Error())
	}

	gameKey, err := stream.ReadStructure(NewGameKey())
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.GameKey. %s", err.Error())
	}

	presence.GameKey = gameKey.(*GameKey)
	presence.Unknown1, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown1. %s", err.Error())
	}

	presence.Message, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Message. %s", err.Error())
	}

	presence.Unknown2, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown2. %s", err.Error())
	}

	presence.Unknown3, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown3. %s", err.Error())
	}

	presence.GameServerID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.GameServerID. %s", err.Error())
	}

	presence.Unknown4, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown4. %s", err.Error())
	}

	presence.PID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.PID. %s", err.Error())
	}

	presence.GatheringID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.GatheringID. %s", err.Error())
	}

	presence.ApplicationData, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.ApplicationData. %s", err.Error())
	}

	presence.Unknown5, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown5. %s", err.Error())
	}

	presence.Unknown6, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown6. %s", err.Error())
	}

	presence.Unknown7, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown7. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoPresenceV2
func (presence *NintendoPresenceV2) Copy() nex.StructureInterface {
	copied := NewNintendoPresenceV2()

	copied.SetStructureVersion(presence.StructureVersion())

	if presence.ParentType() != nil {
		copied.Data = presence.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

	copied.SetParentType(copied.Data)

	copied.ChangedFlags = presence.ChangedFlags
	copied.Online = presence.Online
	copied.GameKey = presence.GameKey.Copy().(*GameKey)
	copied.Unknown1 = presence.Unknown1
	copied.Message = presence.Message
	copied.Unknown2 = presence.Unknown2
	copied.Unknown3 = presence.Unknown3
	copied.GameServerID = presence.GameServerID
	copied.Unknown4 = presence.Unknown4
	copied.PID = presence.PID.Copy()
	copied.GatheringID = presence.GatheringID
	copied.ApplicationData = make([]byte, len(presence.ApplicationData))

	copy(copied.ApplicationData, presence.ApplicationData)

	copied.Unknown5 = presence.Unknown5
	copied.Unknown6 = presence.Unknown6
	copied.Unknown7 = presence.Unknown7

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (presence *NintendoPresenceV2) Equals(structure nex.StructureInterface) bool {
	other := structure.(*NintendoPresenceV2)

	if presence.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !presence.ParentType().Equals(other.ParentType()) {
		return false
	}

	if presence.ChangedFlags != other.ChangedFlags {
		return false
	}

	if presence.Online != other.Online {
		return false
	}

	if !presence.GameKey.Equals(other.GameKey) {
		return false
	}

	if presence.Unknown1 != other.Unknown1 {
		return false
	}

	if presence.Message != other.Message {
		return false
	}

	if presence.Unknown2 != other.Unknown2 {
		return false
	}

	if presence.Unknown3 != other.Unknown3 {
		return false
	}

	if presence.GameServerID != other.GameServerID {
		return false
	}

	if presence.Unknown4 != other.Unknown4 {
		return false
	}

	if !presence.PID.Equals(other.PID) {
		return false
	}

	if presence.GatheringID != other.GatheringID {
		return false
	}

	if !bytes.Equal(presence.ApplicationData, other.ApplicationData) {
		return false
	}

	if presence.Unknown5 != other.Unknown5 {
		return false
	}

	if presence.Unknown6 != other.Unknown6 {
		return false
	}

	if presence.Unknown7 != other.Unknown7 {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (presence *NintendoPresenceV2) String() string {
	return presence.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (presence *NintendoPresenceV2) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("PrincipalBasicInfo{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, presence.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sChangedFlags: %d,\n", indentationValues, presence.ChangedFlags))
	b.WriteString(fmt.Sprintf("%sOnline: %t,\n", indentationValues, presence.Online))

	if presence.GameKey != nil {
		b.WriteString(fmt.Sprintf("%sGameKey: %s,\n", indentationValues, presence.GameKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sGameKey: nil,\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sUnknown1: %d,\n", indentationValues, presence.Unknown1))
	b.WriteString(fmt.Sprintf("%sMessage: %q,\n", indentationValues, presence.Message))
	b.WriteString(fmt.Sprintf("%sUnknown2: %d,\n", indentationValues, presence.Unknown2))
	b.WriteString(fmt.Sprintf("%sUnknown3: %d,\n", indentationValues, presence.Unknown3))
	b.WriteString(fmt.Sprintf("%sGameServerID: %d,\n", indentationValues, presence.GameServerID))
	b.WriteString(fmt.Sprintf("%sUnknown4: %d,\n", indentationValues, presence.Unknown4))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, presence.PID))
	b.WriteString(fmt.Sprintf("%sGatheringID: %d,\n", indentationValues, presence.GatheringID))
	b.WriteString(fmt.Sprintf("%sApplicationData: %x,\n", indentationValues, presence.ApplicationData))
	b.WriteString(fmt.Sprintf("%sUnknown5: %d,\n", indentationValues, presence.Unknown5))
	b.WriteString(fmt.Sprintf("%sUnknown6: %d,\n", indentationValues, presence.Unknown6))
	b.WriteString(fmt.Sprintf("%sUnknown7: %d,\n", indentationValues, presence.Unknown7))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewNintendoPresenceV2 returns a new NintendoPresenceV2
func NewNintendoPresenceV2() *NintendoPresenceV2 {
	return &NintendoPresenceV2{}
}
