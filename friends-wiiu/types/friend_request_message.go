// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendRequestMessage contains message data for a FriendRequest
type FriendRequestMessage struct {
	types.Structure
	*types.Data
	FriendRequestID *types.PrimitiveU64
	Received        *types.PrimitiveBool
	Unknown2        *types.PrimitiveU8
	Message         string
	Unknown3        *types.PrimitiveU8
	Unknown4        string
	GameKey         *GameKey
	Unknown5        *types.DateTime
	ExpiresOn       *types.DateTime
}

// WriteTo writes the FriendRequestMessage to the given writable
func (friendRequestMessage *FriendRequestMessage) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	friendRequestMessage.FriendRequestID.WriteTo(contentWritable)
	friendRequestMessage.Received.WriteTo(contentWritable)
	friendRequestMessage.Unknown2.WriteTo(contentWritable)
	friendRequestMessage.Message.WriteTo(contentWritable)
	friendRequestMessage.Unknown3.WriteTo(contentWritable)
	friendRequestMessage.Unknown4.WriteTo(contentWritable)
	friendRequestMessage.GameKey.WriteTo(contentWritable)
	friendRequestMessage.Unknown5.WriteTo(contentWritable)
	friendRequestMessage.ExpiresOn.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	friendRequestMessage.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of FriendRequestMessage
func (friendRequestMessage *FriendRequestMessage) Copy() types.RVType {
	copied := NewFriendRequestMessage()

	copied.StructureVersion = friendRequestMessage.StructureVersion

	copied.Data = friendRequestMessage.Data.Copy().(*types.Data)

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
func (friendRequestMessage *FriendRequestMessage) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendRequestMessage); !ok {
		return false
	}

	other := o.(*FriendRequestMessage)

	if friendRequestMessage.StructureVersion != other.StructureVersion {
		return false
	}

	if !friendRequestMessage.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !friendRequestMessage.FriendRequestID.Equals(other.FriendRequestID) {
		return false
	}

	if !friendRequestMessage.Received.Equals(other.Received) {
		return false
	}

	if !friendRequestMessage.Unknown2.Equals(other.Unknown2) {
		return false
	}

	if !friendRequestMessage.Message.Equals(other.Message) {
		return false
	}

	if !friendRequestMessage.Unknown3.Equals(other.Unknown3) {
		return false
	}

	if !friendRequestMessage.Unknown4.Equals(other.Unknown4) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, friendRequestMessage.StructureVersion))
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
