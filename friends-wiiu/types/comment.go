package friends_wiiu_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Comment contains data about a text comment
type Comment struct {
	nex.Structure
	Unknown     uint8
	Contents    string
	LastChanged *nex.DateTime
}

// Bytes encodes the Comment and returns a byte array
func (comment *Comment) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt8(comment.Unknown)
	stream.WriteString(comment.Contents)
	stream.WriteDateTime(comment.LastChanged)

	return stream.Bytes()
}

// ExtractFromStream extracts a Comment structure from a stream
func (comment *Comment) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	comment.Unknown, err = stream.ReadUInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.Unknown. %s", err.Error())
	}

	comment.Contents, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.Contents. %s", err.Error())
	}

	comment.LastChanged, err = stream.ReadDateTime()
	if err != nil {
		return fmt.Errorf("Failed to extract Comment.LastChanged. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Comment
func (comment *Comment) Copy() nex.StructureInterface {
	copied := NewComment()

	copied.Unknown = comment.Unknown
	copied.Contents = comment.Contents
	copied.LastChanged = comment.LastChanged.Copy()

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (comment *Comment) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Comment)

	if comment.Unknown != other.Unknown {
		return false
	}

	if comment.Contents != other.Contents {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, comment.StructureVersion()))
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
