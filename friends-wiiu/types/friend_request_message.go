// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// FriendRequestMessage contains message data for a FriendRequest
type FriendRequestMessage struct {
	nex.Structure
	*nex.Data
	FriendRequestID uint64
	Received        bool
	Unknown2        uint8
	Message         string
	Unknown3        uint8
	Unknown4        string
	GameKey         *GameKey
	Unknown5        *nex.DateTime
	ExpiresOn       *nex.DateTime
}

// Bytes encodes the FriendRequestMessage and returns a byte array
func (friendRequestMessage *FriendRequestMessage) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt64LE(friendRequestMessage.FriendRequestID)
	stream.WriteBool(friendRequestMessage.Received)
	stream.WriteUInt8(friendRequestMessage.Unknown2)
	stream.WriteString(friendRequestMessage.Message)
	stream.WriteUInt8(friendRequestMessage.Unknown3)
	stream.WriteString(friendRequestMessage.Unknown4)
	stream.WriteStructure(friendRequestMessage.GameKey)
	stream.WriteDateTime(friendRequestMessage.Unknown5)
	stream.WriteDateTime(friendRequestMessage.ExpiresOn)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendRequestMessage
func (friendRequestMessage *FriendRequestMessage) Copy() nex.StructureInterface {
	copied := NewFriendRequestMessage()

	copied.SetStructureVersion(friendRequestMessage.StructureVersion())

	if friendRequestMessage.ParentType() != nil {
		copied.Data = friendRequestMessage.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

	copied.FriendRequestID = friendRequestMessage.FriendRequestID
	copied.Received = friendRequestMessage.Received
	copied.Unknown2 = friendRequestMessage.Unknown2
	copied.Message = friendRequestMessage.Message
	copied.Unknown3 = friendRequestMessage.Unknown3
	copied.Unknown4 = friendRequestMessage.Unknown4
	copied.GameKey = friendRequestMessage.GameKey.Copy().(*GameKey)
	copied.Unknown5 = friendRequestMessage.Unknown5.Copy()
	copied.ExpiresOn = friendRequestMessage.ExpiresOn.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendRequestMessage *FriendRequestMessage) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendRequestMessage)

	if friendRequestMessage.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !friendRequestMessage.ParentType().Equals(other.ParentType()) {
		return false
	}

	if friendRequestMessage.FriendRequestID != other.FriendRequestID {
		return false
	}

	if friendRequestMessage.Received != other.Received {
		return false
	}

	if friendRequestMessage.Unknown2 != other.Unknown2 {
		return false
	}

	if friendRequestMessage.Message != other.Message {
		return false
	}

	if friendRequestMessage.Unknown3 != other.Unknown3 {
		return false
	}

	if friendRequestMessage.Unknown4 != other.Unknown4 {
		return false
	}

	if !friendRequestMessage.GameKey.Equals(other.GameKey) {
		return false
	}

	if !friendRequestMessage.Unknown5.Equals(other.Unknown5) {
		return false
	}

	if !friendRequestMessage.ExpiresOn.Equals(other.ExpiresOn) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (friendRequestMessage *FriendRequestMessage) String() string {
	return friendRequestMessage.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (friendRequestMessage *FriendRequestMessage) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendRequestMessage{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, friendRequestMessage.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sFriendRequestID: %d,\n", indentationValues, friendRequestMessage.FriendRequestID))
	b.WriteString(fmt.Sprintf("%sReceived: %t,\n", indentationValues, friendRequestMessage.Received))
	b.WriteString(fmt.Sprintf("%sUnknown2: %d,\n", indentationValues, friendRequestMessage.Unknown2))
	b.WriteString(fmt.Sprintf("%sMessage: %q,\n", indentationValues, friendRequestMessage.Message))
	b.WriteString(fmt.Sprintf("%sUnknown3: %d,\n", indentationValues, friendRequestMessage.Unknown3))
	b.WriteString(fmt.Sprintf("%sUnknown4: %q,\n", indentationValues, friendRequestMessage.Unknown4))

	if friendRequestMessage.GameKey != nil {
		b.WriteString(fmt.Sprintf("%sGameKey: %s,\n", indentationValues, friendRequestMessage.GameKey.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sGameKey: nil,\n", indentationValues))
	}

	if friendRequestMessage.Unknown5 != nil {
		b.WriteString(fmt.Sprintf("%sUnknown5: %s,\n", indentationValues, friendRequestMessage.Unknown5.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUnknown5: nil,\n", indentationValues))
	}

	if friendRequestMessage.ExpiresOn != nil {
		b.WriteString(fmt.Sprintf("%sExpiresOn: %s\n", indentationValues, friendRequestMessage.ExpiresOn.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sExpiresOn: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendRequestMessage returns a new FriendRequestMessage
func NewFriendRequestMessage() *FriendRequestMessage {
	return &FriendRequestMessage{}
}
