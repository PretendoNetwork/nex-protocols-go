// Package types implements all the types used by the Shop protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SubscriberContent is a type within the Shop protocol
type SubscriberContent struct {
	types.Structure
	ContentID types.UInt64
	Message   types.String
	Binary    types.QBuffer
	PID       types.PID
	Topics    types.List[types.String]
	PostTime  types.DateTime
}

// WriteTo writes the SubscriberContent to the given writable
func (sc SubscriberContent) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sc.ContentID.WriteTo(contentWritable)
	sc.Message.WriteTo(contentWritable)
	sc.Binary.WriteTo(contentWritable)
	sc.PID.WriteTo(contentWritable)
	sc.Topics.WriteTo(contentWritable)
	sc.PostTime.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sc.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SubscriberContent from the given readable
func (sc *SubscriberContent) ExtractFrom(readable types.Readable) error {
	if err := sc.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent header. %s", err.Error())
	}

	if err := sc.ContentID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.ContentID. %s", err.Error())
	}

	if err := sc.Message.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Message. %s", err.Error())
	}

	if err := sc.Binary.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Binary. %s", err.Error())
	}

	if err := sc.PID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.PID. %s", err.Error())
	}

	if err := sc.Topics.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.Topics. %s", err.Error())
	}

	if err := sc.PostTime.ExtractFrom(readable); err != nil {
		return fmt.Errorf("Failed to extract SubscriberContent.PostTime. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SubscriberContent
func (sc SubscriberContent) Copy() types.RVType {
	copied := NewSubscriberContent()

	copied.StructureVersion = sc.StructureVersion
	copied.ContentID = sc.ContentID.Copy().(types.UInt64)
	copied.Message = sc.Message.Copy().(types.String)
	copied.Binary = sc.Binary.Copy().(types.QBuffer)
	copied.PID = sc.PID.Copy().(types.PID)
	copied.Topics = sc.Topics.Copy().(types.List[types.String])
	copied.PostTime = sc.PostTime.Copy().(types.DateTime)

	return copied
}

// Equals checks if the given SubscriberContent contains the same data as the current SubscriberContent
func (sc SubscriberContent) Equals(o types.RVType) bool {
	if _, ok := o.(SubscriberContent); !ok {
		return false
	}

	other := o.(SubscriberContent)

	if sc.StructureVersion != other.StructureVersion {
		return false
	}

	if !sc.ContentID.Equals(other.ContentID) {
		return false
	}

	if !sc.Message.Equals(other.Message) {
		return false
	}

	if !sc.Binary.Equals(other.Binary) {
		return false
	}

	if !sc.PID.Equals(other.PID) {
		return false
	}

	if !sc.Topics.Equals(other.Topics) {
		return false
	}

	return sc.PostTime.Equals(other.PostTime)
}

// CopyRef copies the current value of the SubscriberContent
// and returns a pointer to the new copy
func (sc SubscriberContent) CopyRef() types.RVTypePtr {
	copied := sc.Copy().(SubscriberContent)
	return &copied
}

// Deref takes a pointer to the SubscriberContent
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sc *SubscriberContent) Deref() types.RVType {
	return *sc
}

// String returns the string representation of the SubscriberContent
func (sc SubscriberContent) String() string {
	return sc.FormatToString(0)
}

// FormatToString pretty-prints the SubscriberContent using the provided indentation level
func (sc SubscriberContent) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberContent{\n")
	b.WriteString(fmt.Sprintf("%sContentID: %s,\n", indentationValues, sc.ContentID))
	b.WriteString(fmt.Sprintf("%sMessage: %s,\n", indentationValues, sc.Message))
	b.WriteString(fmt.Sprintf("%sBinary: %s,\n", indentationValues, sc.Binary))
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, sc.PID))
	b.WriteString(fmt.Sprintf("%sTopics: %s,\n", indentationValues, sc.Topics))
	b.WriteString(fmt.Sprintf("%sPostTime: %s,\n", indentationValues, sc.PostTime.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberContent returns a new SubscriberContent
func NewSubscriberContent() SubscriberContent {
	return SubscriberContent{
		ContentID: types.NewUInt64(0),
		Message:   types.NewString(""),
		Binary:    types.NewQBuffer(nil),
		PID:       types.NewPID(0),
		Topics:    types.NewList[types.String](),
		PostTime:  types.NewDateTime(0),
	}

}
