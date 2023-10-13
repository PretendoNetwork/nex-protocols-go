// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// FriendComment is a data structure used by the Friends 3DS protocol to hold information about a friends Mii
type FriendComment struct {
	nex.Structure
	*nex.Data
	PID        uint32
	Comment    string
	ModifiedAt *nex.DateTime
}

// Bytes encodes the FriendComment and returns a byte array
func (friendComment *FriendComment) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(friendComment.PID)
	stream.WriteString(friendComment.Comment)
	stream.WriteDateTime(friendComment.ModifiedAt)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendComment
func (friendComment *FriendComment) Copy() nex.StructureInterface {
	copied := NewFriendComment()

	copied.SetStructureVersion(friendComment.StructureVersion())

	if friendComment.ParentType() != nil {
		copied.Data = friendComment.ParentType().Copy().(*nex.Data)
	} else {
		copied.Data = nex.NewData()
	}

	copied.PID = friendComment.PID
	copied.Comment = friendComment.Comment
	copied.ModifiedAt = friendComment.ModifiedAt.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendComment *FriendComment) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendComment)

	if friendComment.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !friendComment.ParentType().Equals(other.ParentType()) {
		return false
	}

	if friendComment.PID != other.PID {
		return false
	}

	if friendComment.Comment != other.Comment {
		return false
	}

	if !friendComment.ModifiedAt.Equals(other.ModifiedAt) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (friendComment *FriendComment) String() string {
	return friendComment.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (friendComment *FriendComment) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendComment{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, friendComment.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, friendComment.PID))
	b.WriteString(fmt.Sprintf("%sComment: %q,\n", indentationValues, friendComment.Comment))

	if friendComment.ModifiedAt != nil {
		b.WriteString(fmt.Sprintf("%sModifiedAt: %s\n", indentationValues, friendComment.ModifiedAt.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sModifiedAt: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendComment returns a new FriendComment
func NewFriendComment() *FriendComment {
	return &FriendComment{}
}
