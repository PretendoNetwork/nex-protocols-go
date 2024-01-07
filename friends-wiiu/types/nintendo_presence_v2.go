// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// NintendoPresenceV2 contains information about a users online presence
type NintendoPresenceV2 struct {
	types.Structure
	*types.Data
	ChangedFlags    *types.PrimitiveU32
	Online          *types.PrimitiveBool
	GameKey         *GameKey
	Unknown1        *types.PrimitiveU8
	Message         string
	Unknown2        *types.PrimitiveU32
	Unknown3        *types.PrimitiveU8
	GameServerID    *types.PrimitiveU32
	Unknown4        *types.PrimitiveU32
	PID             *types.PID
	GatheringID     *types.PrimitiveU32
	ApplicationData []byte
	Unknown5        *types.PrimitiveU8
	Unknown6        *types.PrimitiveU8
	Unknown7        *types.PrimitiveU8
}

// WriteTo writes the NintendoPresenceV2 to the given writable
func (presence *NintendoPresenceV2) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	presence.ChangedFlags.WriteTo(contentWritable)
	presence.Online.WriteTo(contentWritable)
	presence.GameKey.WriteTo(contentWritable)
	presence.Unknown1.WriteTo(contentWritable)
	presence.Message.WriteTo(contentWritable)
	presence.Unknown2.WriteTo(contentWritable)
	presence.Unknown3.WriteTo(contentWritable)
	presence.GameServerID.WriteTo(contentWritable)
	presence.Unknown4.WriteTo(contentWritable)
	presence.PID.WriteTo(contentWritable)
	presence.GatheringID.WriteTo(contentWritable)
	stream.WriteBuffer(presence.ApplicationData)
	presence.Unknown5.WriteTo(contentWritable)
	presence.Unknown6.WriteTo(contentWritable)
	presence.Unknown7.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	presence.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the NintendoPresenceV2 from the given readable
func (presence *NintendoPresenceV2) ExtractFrom(readable types.Readable) error {
	var err error

	if err = presence.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read NintendoPresenceV2 header. %s", err.Error())
	}

	err = presence.ChangedFlags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.ChangedFlags. %s", err.Error())
	}

	err = presence.Online.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Online. %s", err.Error())
	}

	err = presence.GameKey.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.GameKey. %s", err.Error())
	}

	err = presence.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown1. %s", err.Error())
	}

	err = presence.Message.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Message. %s", err.Error())
	}

	err = presence.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown2. %s", err.Error())
	}

	err = presence.Unknown3.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown3. %s", err.Error())
	}

	err = presence.GameServerID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.GameServerID. %s", err.Error())
	}

	err = presence.Unknown4.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown4. %s", err.Error())
	}

	err = presence.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.PID. %s", err.Error())
	}

	err = presence.GatheringID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.GatheringID. %s", err.Error())
	}

	presence.ApplicationData, err = stream.ReadBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.ApplicationData. %s", err.Error())
	}

	err = presence.Unknown5.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown5. %s", err.Error())
	}

	err = presence.Unknown6.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown6. %s", err.Error())
	}

	err = presence.Unknown7.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract NintendoPresenceV2.Unknown7. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of NintendoPresenceV2
func (presence *NintendoPresenceV2) Copy() types.RVType {
	copied := NewNintendoPresenceV2()

	copied.StructureVersion = presence.StructureVersion

	copied.Data = presence.Data.Copy().(*types.Data)

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
func (presence *NintendoPresenceV2) Equals(o types.RVType) bool {
	if _, ok := o.(*NintendoPresenceV2); !ok {
		return false
	}

	other := o.(*NintendoPresenceV2)

	if presence.StructureVersion != other.StructureVersion {
		return false
	}

	if !presence.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !presence.ChangedFlags.Equals(other.ChangedFlags) {
		return false
	}

	if !presence.Online.Equals(other.Online) {
		return false
	}

	if !presence.GameKey.Equals(other.GameKey) {
		return false
	}

	if !presence.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !presence.Message.Equals(other.Message) {
		return false
	}

	if !presence.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !presence.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !presence.GameServerID.Equals(other.GameServerID) {
		return false
	}

	if !presence.Unknown4.Equals(other.Unknown4) {
		return false
	}

	if !presence.PID.Equals(other.PID) {
		return false
	}

	if !presence.GatheringID.Equals(other.GatheringID) {
		return false
	}

	if !presence.ApplicationData.Equals(other.ApplicationData) {
		return false
	}

	if !presence.Unknown5.Equals(other.Unknown5) {
		return false
	}

	if !presence.Unknown6.Equals(other.Unknown6) {
		return false
	}

	if !presence.Unknown7.Equals(other.Unknown7) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, presence.StructureVersion))
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
