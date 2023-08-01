// Package types implements all the types used by the Shop protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// SubscriberStatusInfo is unknown
type SubscriberStatusInfo struct {
	nex.Structure
	PID     uint32
	Unknown [][]byte
}

// ExtractFromStream extracts a SubscriberStatusInfo structure from a stream
func (subscriberPostContentParam *SubscriberStatusInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	subscriberPostContentParam.PID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberStatusInfo.PID from stream. %s", err.Error())
	}

	subscriberPostContentParam.Unknown, err = stream.ReadListQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriberStatusInfo.Unknown from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the SubscriberStatusInfo and returns a byte array
func (subscriberPostContentParam *SubscriberStatusInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(subscriberPostContentParam.PID)
	stream.WriteListQBuffer(subscriberPostContentParam.Unknown)

	return stream.Bytes()
}

// Copy returns a new copied instance of SubscriberStatusInfo
func (subscriberPostContentParam *SubscriberStatusInfo) Copy() nex.StructureInterface {
	copied := NewSubscriberStatusInfo()

	copied.PID = subscriberPostContentParam.PID
	copied.Unknown = make([][]byte, len(subscriberPostContentParam.Unknown))

	for i := 0; i < len(subscriberPostContentParam.Unknown); i++ {
		copied.Unknown[i] = make([]byte, len(subscriberPostContentParam.Unknown[i]))

		copy(copied.Unknown[i], subscriberPostContentParam.Unknown[i])
	}

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (subscriberPostContentParam *SubscriberStatusInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*SubscriberStatusInfo)

	if subscriberPostContentParam.PID != other.PID {
		return false
	}

	if len(subscriberPostContentParam.Unknown) != len(other.Unknown) {
		return false
	}

	for i := 0; i < len(subscriberPostContentParam.Unknown); i++ {
		if !bytes.Equal(subscriberPostContentParam.Unknown[i], other.Unknown[i]) {
			return false
		}
	}

	return true
}

// String returns a string representation of the struct
func (subscriberPostContentParam *SubscriberStatusInfo) String() string {
	return subscriberPostContentParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (subscriberPostContentParam *SubscriberStatusInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriberStatusInfo{\n")
	b.WriteString(fmt.Sprintf("%sPID: %d,\n", indentationValues, subscriberPostContentParam.PID))
	b.WriteString(fmt.Sprintf("%sUnknown: %v\n", indentationValues, subscriberPostContentParam.Unknown)) // TODO - Make this a nicer looking log
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriberStatusInfo returns a new SubscriberStatusInfo
func NewSubscriberStatusInfo() *SubscriberStatusInfo {
	return &SubscriberStatusInfo{}
}
