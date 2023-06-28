package friends_wiiu_types

import "github.com/PretendoNetwork/nex-go"

// FriendRequestMessage contains message data for a FriendRequest
type FriendRequestMessage struct {
	nex.Structure
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

// NewFriendRequestMessage returns a new FriendRequestMessage
func NewFriendRequestMessage() *FriendRequestMessage {
	return &FriendRequestMessage{}
}
