// Package types implements all the types used by the Matchmake Extension (Monster Hunter XX) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// FriendUserParam holds data for the Matchmake Extension (Monster Hunter XX) protocol
type FriendUserParam struct {
	types.Structure
	Name string
}

// ExtractFrom extracts the FriendUserParam from the given readable
func (friendUserParam *FriendUserParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = friendUserParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read FriendUserParam header. %s", err.Error())
	}

	err = friendUserParam.Name.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract FriendUserParam.Name from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the FriendUserParam to the given writable
func (friendUserParam *FriendUserParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	friendUserParam.Name.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	friendUserParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of FriendUserParam
func (friendUserParam *FriendUserParam) Copy() types.RVType {
	copied := NewFriendUserParam()

	copied.StructureVersion = friendUserParam.StructureVersion

	copied.Name = friendUserParam.Name

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (friendUserParam *FriendUserParam) Equals(o types.RVType) bool {
	if _, ok := o.(*FriendUserParam); !ok {
		return false
	}

	other := o.(*FriendUserParam)

	if friendUserParam.StructureVersion != other.StructureVersion {
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
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, friendUserParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sName: %q,\n", indentationValues, friendUserParam.Name))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewFriendUserParam returns a new FriendUserParam
func NewFriendUserParam() *FriendUserParam {
	return &FriendUserParam{}
}
