// Package types implements all the types used by the Friends WiiU protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Comment contains data about a text comment
type Comment struct {
	types.Structure
	*types.Data
	Unknown     *types.PrimitiveU8
	Contents    string
	LastChanged *types.DateTime
}

// WriteTo writes the Comment to the given writable
func (comment *Comment) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	comment.Unknown.WriteTo(contentWritable)
	comment.Contents.WriteTo(contentWritable)
	comment.LastChanged.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	comment.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Comment from the given readable
func (comment *Comment) ExtractFrom(readable types.Readable) error {
	var err error

	if err = comment.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Comment header. %s", err.Error())
	}

	err = comment.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.Unknown. %s", err.Error())
	}

	err = comment.Contents.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.Contents. %s", err.Error())
	}

	err = comment.LastChanged.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.LastChanged. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Comment
func (comment *Comment) Copy() types.RVType {
	copied := NewComment()

	copied.StructureVersion = comment.StructureVersion

	copied.Data = comment.Data.Copy().(*types.Data)

	copied.Unknown = comment.Unknown
	copied.Contents = comment.Contents
	copied.LastChanged = comment.LastChanged.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (comment *Comment) Equals(o types.RVType) bool {
	if _, ok := o.(*Comment); !ok {
		return false
	}

	other := o.(*Comment)

	if comment.StructureVersion != other.StructureVersion {
		return false
	}

	if !comment.ParentType().Equals(other.ParentType()) {
		return false
	}

	if !comment.Unknown.Equals(other.Unknown) {
		return false
	}

	if !comment.Contents.Equals(other.Contents) {
		return false
	}

	if !comment.LastChanged.Equals(other.LastChanged) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (comment *Comment) String() string {
	return comment.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (comment *Comment) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Comment{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, comment.StructureVersion))
	b.WriteString(fmt.Sprintf("%sUnknown: %d\n", indentationValues, comment.Unknown))
	b.WriteString(fmt.Sprintf("%sContents: %q\n", indentationValues, comment.Contents))

	if comment.LastChanged != nil {
		b.WriteString(fmt.Sprintf("%sLastChanged: %s\n", indentationValues, comment.LastChanged.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sLastChanged: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewComment returns a new Comment
func NewComment() *Comment {
	return &Comment{}
}
