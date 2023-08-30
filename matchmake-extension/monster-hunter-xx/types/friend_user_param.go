// Package types implements all the types used by the Matchmake Extension (Monster Hunter XX) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// FriendUserParam holds data for the Matchmake Extension (Monster Hunter XX) protocol
type FriendUserParam struct {
	nex.Structure
	Name string
}

// ExtractFromStream extracts a FriendUserParam structure from a stream
func (friendUserParam *FriendUserParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	friendUserParam.Name, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserParam.Name from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the FriendUserParam and returns a byte array
func (friendUserParam *FriendUserParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(friendUserParam.Name)

	return stream.Bytes()
}

// Copy returns a new copied instance of FriendUserParam
func (friendUserParam *FriendUserParam) Copy() nex.StructureInterface {
	copied := NewFriendUserParam()

	copied.SetStructureVersion(friendUserParam.StructureVersion())

	copied.Name = friendUserParam.Name

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendUserParam *FriendUserParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*FriendUserParam)

	if friendUserParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	return friendUserParam.Name == other.Name
}

// String returns a string representation of the struct
func (friendUserParam *FriendUserParam) String() string {
	return friendUserParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (friendUserParam *FriendUserParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("FriendUserParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, friendUserParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, friendUserParam.Name))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendUserParam returns a new FriendUserParam
func NewFriendUserParam() *FriendUserParam {
	return &FriendUserParam{}
}
