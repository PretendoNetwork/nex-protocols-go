// Package types implements all the types used by the Friends 3DS protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// FriendPicture is a data structure used by the Friends 3DS protocol to hold information about a friends PictureData
type FriendPicture struct {
	nex.Structure
	*nex.Data
	Unknown1    uint32
	PictureData []byte
	Unknown2    *nex.DateTime
}

// Bytes encodes the FriendPicture and returns a byte array
func (friendPicture *FriendPicture) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(friendPicture.Unknown1)
	stream.WriteBuffer(friendPicture.PictureData)
	stream.WriteDateTime(friendPicture.Unknown2)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendPicture
func (friendPicture *FriendPicture) Copy() nex.StructureInterface {
	copied := NewFriendPicture()

	copied.Data = friendPicture.ParentType().Copy().(*nex.Data)
	copied.SetParentType(copied.Data)

	copied.Unknown1 = friendPicture.Unknown1
	copied.PictureData = make([]byte, len(friendPicture.PictureData))

	copy(copied.PictureData, friendPicture.PictureData)

	copied.Unknown2 = friendPicture.Unknown2.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendPicture *FriendPicture) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendPicture)

	if !friendPicture.ParentType().Equals(other.ParentType()) {
		return false
	}

	if friendPicture.Unknown1 != other.Unknown1 {
		return false
	}

	if !bytes.Equal(friendPicture.PictureData, other.PictureData) {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, friendPicture.StructureVersion()))
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
