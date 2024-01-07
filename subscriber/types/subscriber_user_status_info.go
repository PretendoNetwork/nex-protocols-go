// Package types implements all the types used by the Shop protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// SubscriberUserStatusInfo is unknown
type SubscriberUserStatusInfo struct {
	types.Structure
	PID     *types.PID
	Unknown [][]byte
}

// ExtractFrom extracts the SubscriberUserStatusInfo from the given readable
func (s *SubscriberUserStatusInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = s.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read SubscriberUserStatusInfo header. %s", err.Error())
	}

	err = s.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberUserStatusInfo.PID from stream. %s", err.Error())
	}

	s.Unknown, err = stream.ReadListQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberUserStatusInfo.Unknown from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the SubscriberUserStatusInfo to the given writable
func (s *SubscriberUserStatusInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	s.PID.WriteTo(contentWritable)
	stream.WriteListQBuffer(s.Unknown)

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of SubscriberUserStatusInfo
func (s *SubscriberUserStatusInfo) Copy() types.RVType {
	copied := NewSubscriberUserStatusInfo()

	copied.StructureVersion = s.StructureVersion

	copied.PID = s.PID.Copy()
	copied.Unknown = make([][]byte, len(s.Unknown))

	for i := 0; i < len(s.Unknown); i++ {
		copied.Unknown[i] = make([]byte, len(s.Unknown[i]))

		copy(copied.Unknown[i], s.Unknown[i])
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (s *SubscriberUserStatusInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*SubscriberUserStatusInfo); !ok {
		return false
	}

	other := o.(*SubscriberUserStatusInfo)

	if s.StructureVersion != other.StructureVersion {
		return false
	}

	if !s.PID.Equals(other.PID) {
		return false
	}

	if len(s.Unknown) != len(other.Unknown) {
		return false
	}

	for i := 0; i < len(s.Unknown); i++ {
		if !s.Unknown[i].Equals(other.Unknown[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (s *SubscriberUserStatusInfo) String() string {
	return s.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (s *SubscriberUserStatusInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberUserStatusInfo{\n")
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, s.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown: %v\n", indentationValues, s.Unknown)) // TODO - Make this a nicer looking log
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberUserStatusInfo returns a new SubscriberUserStatusInfo
func NewSubscriberUserStatusInfo() *SubscriberUserStatusInfo {
	return &SubscriberUserStatusInfo{}
}
