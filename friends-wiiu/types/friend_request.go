// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// FriendRequest contains information about a friend request
type FriendRequest struct {
	nex.Structure
	*nex.Data
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

	copied.SetStructureVersion(friendRequest.StructureVersion())

	if friendRequest.ParentType() != nil {
		copied.Data = friendRequest.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

	copied.SetParentType(copied.Data)

	copied.PrincipalInfo = friendRequest.PrincipalInfo.Copy().(*PrincipalBasicInfo)
	copied.Message = friendRequest.Message.Copy().(*FriendRequestMessage)
	copied.SentOn = friendRequest.SentOn.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendRequest *FriendRequest) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendRequest)

	if friendRequest.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !friendRequest.ParentType().Equals(other.ParentType()) {
		return false
	}

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

// String returns a string representation of the struct
func (friendRequest *FriendRequest) String() string {
	return friendRequest.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (friendRequest *FriendRequest) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendRequest{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, friendRequest.StructureVersion()))

	if friendRequest.PrincipalInfo != nil {
		b.WriteString(fmt.Sprintf("%sPrincipalInfo: %s,\n", indentationValues, friendRequest.PrincipalInfo.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPrincipalInfo: nil,\n", indentationValues))
	}

	if friendRequest.Message != nil {
		b.WriteString(fmt.Sprintf("%sMessage: %s,\n", indentationValues, friendRequest.Message.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sMessage: nil,\n", indentationValues))
	}

	if friendRequest.SentOn != nil {
		b.WriteString(fmt.Sprintf("%sSentOn: %s\n", indentationValues, friendRequest.SentOn.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sSentOn: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendRequest returns a new FriendRequest
func NewFriendRequest() *FriendRequest {
	return &FriendRequest{}
}
