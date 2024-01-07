// Package types implements all the types used by the Shop protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Unknown is unknown
type Unknown struct {
	types.Structure
	Unknown []byte
}

// ExtractFrom extracts the Unknown from the given readable
func (subscriberPostContentParam *Unknown) ExtractFrom(readable types.Readable) error {
	var err error

	if err = subscriberPostContentParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Unknown header. %s", err.Error())
	}

	subscriberPostContentParam.Unknown, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract Unknown.Unknown from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Unknown to the given writable
func (subscriberPostContentParam *Unknown) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	stream.WriteQBuffer(subscriberPostContentParam.Unknown)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Unknown
func (subscriberPostContentParam *Unknown) Copy() types.RVType {
	copied := NewUnknown()

	copied.StructureVersion = subscriberPostContentParam.StructureVersion

	copied.Unknown = make([]byte, len(subscriberPostContentParam.Unknown))

	copy(copied.Unknown, subscriberPostContentParam.Unknown)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (subscriberPostContentParam *Unknown) Equals(o types.RVType) bool {
	if _, ok := o.(*Unknown); !ok {
		return false
	}

	other := o.(*Unknown)

	if subscriberPostContentParam.StructureVersion != other.StructureVersion {
		return false
	}

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
