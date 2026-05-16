package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// TimeScope is used by RankingOrderParam.TimeScope to request that scores only be shown from a certain timeframe.
type TimeScope uint8

// WriteTo writes the TimeScope to the given writable
func (ts TimeScope) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(ts))
}

// ExtractFrom extracts the TimeScope value from the given readable
func (ts *TimeScope) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*ts = TimeScope(value)
	return nil
}

// String returns a human-readable representation of the TimeScope.
func (ts TimeScope) String() string {
	switch ts {
	case TimeScopeCustom0:
		return "Custom0"
	case TimeScopeCustom1:
		return "Custom1"
	case TimeScopeAll:
		return "All"
	default:
		return fmt.Sprintf("TimeScope(%d)", int(ts))
	}
}

const (
	// TimeScopeCustom0 requests only scores from game-specific time scope.
	TimeScopeCustom0 TimeScope = iota

	// TimeScopeCustom1 requests only scores from a second, funnier game-specific time scope.
	TimeScopeCustom1

	// TimeScopeAll requests scores fom all time (no filtering).
	TimeScopeAll
)
