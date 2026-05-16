package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MessageFlag determines how the message is delivered to the recipient.
type MessageFlag uint32

// WriteTo writes the MessageFlag to the given writable
func (mf MessageFlag) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(mf))
}

// ExtractFrom extracts the MessageFlag value from the given readable
func (mf *MessageFlag) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*mf = MessageFlag(value)
	return nil
}

// HasFlag checks if a given flag is set
func (mf MessageFlag) HasFlag(flag MessageFlag) bool {
	return mf&flag == flag
}

// HasFlag checks if all given flags are set
func (mf MessageFlag) HasFlags(flags ...MessageFlag) bool {
	if len(flags) == 0 {
		return false
	}

	for _, flag := range flags {
		if mf&flag != flag {
			return false
		}
	}

	return true
}

// String returns a human-readable representation of the MessageFlag.
func (mf MessageFlag) String() string {
	switch mf {
	case MessageFlagPersistentMessage:
		return "PersistentMessage"
	case MessageFlagInstantMessage:
		return "InstantMessage"
	default:
		return fmt.Sprintf("MessageFlag(%d)", int(mf))
	}
}

const (
	// MessageFlagPersistentMessage means that message is stored in the database
	// and not sent immediately.
	//
	// Note: This name is a guess based on the naming conventions of other protocols
	// and how they store data in databases.
	MessageFlagPersistentMessage MessageFlag = iota

	// MessageFlagInstantMessage means that message is not stored in the database
	// and is instead sent immediately to the recipient.
	MessageFlagInstantMessage
)
