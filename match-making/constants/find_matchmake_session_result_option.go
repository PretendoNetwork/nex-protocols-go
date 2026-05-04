package constants

import (
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// FindMatchmakeSessionResultOption indicates how to populate the
// responses from FindMatchmakeSessionByParticipant
type FindMatchmakeSessionResultOption uint32

// WriteTo writes the FindMatchmakeSessionResultOption to the given writable
func (fmsro FindMatchmakeSessionResultOption) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(fmsro))
}

// ExtractFrom extracts the FindMatchmakeSessionResultOption value from the given readable
func (fmsro *FindMatchmakeSessionResultOption) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*fmsro = FindMatchmakeSessionResultOption(value)
	return nil
}

func (fmsro FindMatchmakeSessionResultOption) HasFlag(flag FindMatchmakeSessionResultOption) bool {
	return fmsro&flag == flag
}

func (fmsro FindMatchmakeSessionResultOption) HasFlags(flags ...FindMatchmakeSessionResultOption) bool {
	if len(flags) == 0 {
		return false
	}

	for _, flag := range flags {
		if fmsro&flag != flag {
			return false
		}
	}

	return true
}

// String returns a human-readable representation of the FindMatchmakeSessionResultOption bitmask.
// Multiple flags are joined with "|", e.g. "ApplicationBuffer|MatchmakeParam".
// Returns "None" if no flags are set.
func (fmsro FindMatchmakeSessionResultOption) String() string {
	if fmsro == FindMatchmakeSessionResultOptionNone {
		return "None"
	}

	flags := []struct {
		flag FindMatchmakeSessionResultOption
		name string
	}{
		{FindMatchmakeSessionResultOptionApplicationBuffer, "ApplicationBuffer"},
		{FindMatchmakeSessionResultOptionMatchmakeParam, "MatchmakeParam"},
	}

	var parts []string
	for _, f := range flags {
		if fmsro&f.flag != 0 {
			parts = append(parts, f.name)
		}
	}

	return strings.Join(parts, "|")
}

const (
	// FindMatchmakeSessionResultOptionNone indicates no options
	FindMatchmakeSessionResultOptionNone FindMatchmakeSessionResultOption = 0

	// FindMatchmakeSessionResultOptionApplicationBuffer populates m_ApplicationBuffer in the results
	FindMatchmakeSessionResultOptionApplicationBuffer FindMatchmakeSessionResultOption = 1

	// FindMatchmakeSessionResultOptionMatchmakeParam populates m_MatchmakeParam in the results
	FindMatchmakeSessionResultOptionMatchmakeParam FindMatchmakeSessionResultOption = 2
)
