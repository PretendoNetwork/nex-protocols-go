// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendPicture is a data structure used by the Friends 3DS protocol to hold information about a friends PictureData
type FriendPicture struct {
	types.Structure
	*types.Data
	Unknown1    *types.PrimitiveU32
	PictureData []byte
	Unknown2    *types.DateTime
}

// WriteTo writes the FriendPicture to the given writable
func (friendPicture *FriendPicture) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	friendPicture.Unknown1.WriteTo(contentWritable)
	stream.WriteBuffer(friendPicture.PictureData)
	friendPicture.Unknown2.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	friendPicture.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of FriendPicture
func (friendPicture *FriendPicture) Copy() types.RVType {
	copied := NewFriendPicture()

	copied.StructureVersion = friendPicture.StructureVersion

	copied.Data = friendPicture.Data.Copy().(*types.Data)

	copied.Unknown1 = friendPicture.Unknown1
	copied.PictureData = make([]byte, len(friendPicture.PictureData))

	copy(copied.PictureData, friendPicture.PictureData)

	copied.Unknown2 = friendPicture.Unknown2.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendPicture *FriendPicture) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendPicture); !ok {
		return false
	}

	other := o.(*FriendPicture)

	if friendPicture.StructureVersion != other.StructureVersion {
		return false
	}

	if !friendPicture.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !friendPicture.Unknown1.Equals(other.Unknown1) {
		return false
	}

	if !friendPicture.PictureData.Equals(other.PictureData) {
		return false
	}

	if !friendPicture.Unknown2.Equals(other.Unknown2) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (friendPicture *FriendPicture) String() string {
	return friendPicture.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (friendPicture *FriendPicture) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendPicture{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, friendPicture.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUnknown1: %d,\n", indentationValues, friendPicture.Unknown1))
	b.WriteString(fmt.Sprintf("%sPictureData: %x,\n", indentationValues, friendPicture.PictureData))

	if friendPicture.Unknown2 != nil {
		b.WriteString(fmt.Sprintf("%sUnknown2: %s\n", indentationValues, friendPicture.Unknown2.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUnknown2: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendPicture returns a new FriendPicture
func NewFriendPicture() *FriendPicture {
	return &FriendPicture{}
}
