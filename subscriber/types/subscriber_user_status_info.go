// Package types implements all the types used by the Shop protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// SubscriberUserStatusInfo is unknown
type SubscriberUserStatusInfo struct {
	nex.Structure
	PID     *nex.PID
	Unknown [][]byte
}

// ExtractFromStream extracts a SubscriberUserStatusInfo structure from a stream
func (s *SubscriberUserStatusInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	s.PID, err = stream.ReadPID()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberUserStatusInfo.PID from stream. %s", err.Error())
	}

	s.Unknown, err = stream.ReadListQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberUserStatusInfo.Unknown from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the SubscriberUserStatusInfo and returns a byte array
func (s *SubscriberUserStatusInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WritePID(s.PID)
	stream.WriteListQBuffer(s.Unknown)

	return stream.Bytes()
}

// Copy returns a new copied instance of SubscriberUserStatusInfo
func (s *SubscriberUserStatusInfo) Copy() nex.StructureInterface {
	copied := NewSubscriberUserStatusInfo()

	copied.SetStructureVersion(s.StructureVersion())

	copied.PID = s.PID.Copy()
	copied.Unknown = make([][]byte, len(s.Unknown))

	for i := 0; i < len(s.Unknown); i++ {
		copied.Unknown[i] = make([]byte, len(s.Unknown[i]))

		copy(copied.Unknown[i], s.Unknown[i])
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (s *SubscriberUserStatusInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SubscriberUserStatusInfo)

	if s.StructureVersion() != other.StructureVersion() {
		return false
	}

	if !s.PID.Equals(other.PID) {
		return false
	}

	if len(s.Unknown) != len(other.Unknown) {
		return false
	}

	for i := 0; i < len(s.Unknown); i++ {
		if !bytes.Equal(s.Unknown[i], other.Unknown[i]) {
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
