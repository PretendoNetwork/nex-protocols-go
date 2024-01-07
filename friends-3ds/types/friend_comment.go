// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendComment is a data structure used by the Friends 3DS protocol to hold information about a friends Mii
type FriendComment struct {
	types.Structure
	*types.Data
	PID        *types.PID
	Comment    string
	ModifiedAt *types.DateTime
}

// WriteTo writes the FriendComment to the given writable
func (friendComment *FriendComment) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	friendComment.PID.WriteTo(contentWritable)
	friendComment.Comment.WriteTo(contentWritable)
	friendComment.ModifiedAt.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	friendComment.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of FriendComment
func (friendComment *FriendComment) Copy() types.RVType {
	copied := NewFriendComment()

	copied.StructureVersion = friendComment.StructureVersion

	copied.Data = friendComment.Data.Copy().(*types.Data)

	copied.PID = friendComment.PID.Copy()
	copied.Comment = friendComment.Comment
	copied.ModifiedAt = friendComment.ModifiedAt.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendComment *FriendComment) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendComment); !ok {
		return false
	}

	other := o.(*FriendComment)

	if friendComment.StructureVersion != other.StructureVersion {
		return false
	}

	if !friendComment.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !friendComment.PID.Equals(other.PID) {
		return false
	}

	if !friendComment.Comment.Equals(other.Comment) {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, friendComment.StructureVersion))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, friendComment.PID.FormatToString(indentationLevel+1)))
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
