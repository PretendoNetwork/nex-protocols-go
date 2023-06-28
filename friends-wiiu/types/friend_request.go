package friends_wiiu_types

import "github.com/PretendoNetwork/nex-go"

// FriendRequest contains information about a friend request
type FriendRequest struct {
	nex.Structure
	PrincipalInfo *PrincipalBasicInfo
	Message       *FriendRequestMessage
	SentOn        *nex.DateTime
}

// Bytes encodes the FriendRequest and returns a byte array
func (friendRequest *FriendRequest) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteStructure(friendRequest.PrincipalInfo)
	stream.WriteStructure(friendRequest.Message)
	stream.WriteDateTime(friendRequest.SentOn)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendRequest
func (friendRequest *FriendRequest) Copy() nex.StructureInterface {
	copied := NewFriendRequest()

	copied.PrincipalInfo = friendRequest.PrincipalInfo.Copy().(*PrincipalBasicInfo)
	copied.Message = friendRequest.Message.Copy().(*FriendRequestMessage)
	copied.SentOn = friendRequest.SentOn.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendRequest *FriendRequest) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendRequest)

	if !friendRequest.PrincipalInfo.Equals(other.PrincipalInfo) {
		return false
	}

	if !friendRequest.Message.Equals(other.Message) {
		return false
	}

	if !friendRequest.SentOn.Equals(other.SentOn) {
		return false
	}

	return true
}

// NewFriendRequest returns a new FriendRequest
func NewFriendRequest() *FriendRequest {
	return &FriendRequest{}
}
