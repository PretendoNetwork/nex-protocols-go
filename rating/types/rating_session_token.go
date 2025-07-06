// Package types implements all the types used by the Rating protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// RatingSessionToken is a type within the Rating protocol
type RatingSessionToken struct {
	types.Structure
	Unknown1 types.UInt64
	Unknown2 types.QBuffer
}

// WriteTo writes the RatingSessionToken to the given writable
func (rst RatingSessionToken) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	rst.Unknown1.WriteTo(contentWritable)
	rst.Unknown2.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	rst.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the RatingSessionToken from the given readable
func (rst *RatingSessionToken) ExtractFrom(readable types.Readable) error {
	var err error

	err = rst.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingSessionToken header. %s", err.Error())
	}

	err = rst.Unknown1.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingSessionToken.Unknown1. %s", err.Error())
	}

	err = rst.Unknown2.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract RatingSessionToken.Unknown2. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of RatingSessionToken
func (rst RatingSessionToken) Copy() types.RVType {
	copied := NewRatingSessionToken()

	copied.StructureVersion = rst.StructureVersion
	copied.Unknown1 = rst.Unknown1.Copy().(types.UInt64)
	copied.Unknown2 = rst.Unknown2.Copy().(types.QBuffer)

	return copied
}

// Equals checks if the given RatingSessionToken contains the same data as the current RatingSessionToken
func (rst RatingSessionToken) Equals(o types.RVType) bool {
	if _, ok := o.(RatingSessionToken); !ok {
		return false
	}

	other := o.(RatingSessionToken)

	if rst.StructureVersion != other.StructureVersion {
		return false
	}

	if !rst.Unknown1.Equals(other.Unknown1) {
		return false
	}

	return rst.Unknown2.Equals(other.Unknown2)
}

// CopyRef copies the current value of the RatingSessionToken
// and returns a pointer to the new copy
func (rst RatingSessionToken) CopyRef() types.RVTypePtr {
	copied := rst.Copy().(RatingSessionToken)
	return &copied
}

// Deref takes a pointer to the RatingSessionToken
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (rst *RatingSessionToken) Deref() types.RVType {
	return *rst
}

// String returns the string representation of the RatingSessionToken
func (rst RatingSessionToken) String() string {
	return rst.FormatToString(0)
}

// FormatToString pretty-prints the RatingSessionToken using the provided indentation level
func (rst RatingSessionToken) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("RatingSessionToken{\n")
	b.WriteString(fmt.Sprintf("%sUnknown1: %s,\n", indentationValues, rst.Unknown1))
	b.WriteString(fmt.Sprintf("%sUnknown2: %s,\n", indentationValues, rst.Unknown2))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewRatingSessionToken returns a new RatingSessionToken
func NewRatingSessionToken() RatingSessionToken {
	return RatingSessionToken{
		Unknown1: types.NewUInt64(0),
		Unknown2: types.NewQBuffer(nil),
	}

}
