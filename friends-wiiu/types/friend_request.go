// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendRequest contains information about a friend request
type FriendRequest struct {
	types.Structure
	*types.Data
	PrincipalInfo *PrincipalBasicInfo
	Message       *FriendRequestMessage
	SentOn        *types.DateTime
}

// WriteTo writes the FriendRequest to the given writable
func (friendRequest *FriendRequest) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	friendRequest.PrincipalInfo.WriteTo(contentWritable)
	friendRequest.Message.WriteTo(contentWritable)
	friendRequest.SentOn.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	friendRequest.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of FriendRequest
func (friendRequest *FriendRequest) Copy() types.RVType {
	copied := NewFriendRequest()

	copied.StructureVersion = friendRequest.StructureVersion

	copied.Data = friendRequest.Data.Copy().(*types.Data)

	copied.PrincipalInfo = friendRequest.PrincipalInfo.Copy().(*PrincipalBasicInfo)
	copied.Message = friendRequest.Message.Copy().(*FriendRequestMessage)
	copied.SentOn = friendRequest.SentOn.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendRequest *FriendRequest) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendRequest); !ok {
		return false
	}

	other := o.(*FriendRequest)

	if friendRequest.StructureVersion != other.StructureVersion {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, friendRequest.StructureVersion))

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
