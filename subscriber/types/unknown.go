// Package types implements all the types used by the Shop protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Unknown is unknown
type Unknown struct {
	nex.Structure
	Unknown []byte
}

// ExtractFromStream extracts a Unknown structure from a stream
func (subscriberPostContentParam *Unknown) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	subscriberPostContentParam.Unknown, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract Unknown.Unknown from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the Unknown and returns a byte array
func (subscriberPostContentParam *Unknown) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteQBuffer(subscriberPostContentParam.Unknown)

	return stream.Bytes()
}

// Copy returns a new copied instance of Unknown
func (subscriberPostContentParam *Unknown) Copy() nex.StructureInterface {
	copied := NewUnknown()

	copied.Unknown = make([]byte, len(subscriberPostContentParam.Unknown))

	copy(copied.Unknown, subscriberPostContentParam.Unknown)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (subscriberPostContentParam *Unknown) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Unknown)

	return bytes.Equal(subscriberPostContentParam.Unknown, other.Unknown)
}

// String returns a string representation of the struct
func (subscriberPostContentParam *Unknown) String() string {
	return subscriberPostContentParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (subscriberPostContentParam *Unknown) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Unknown{\n")
	b.WriteString(fmt.Sprintf("%sUnknown: %x\n", indentationValues, subscriberPostContentParam.Unknown))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewUnknown returns a new Unknown
func NewUnknown() *Unknown {
	return &Unknown{}
}
