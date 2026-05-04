package constants

import (
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// PresenceChangedFlag indicates what fields of a users NintendoPresenceV2 have been changed
type PresenceChangedFlag uint32

// WriteTo writes the PresenceChangedFlag to the given writable
func (pcf PresenceChangedFlag) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(pcf))
}

// ExtractFrom extracts the PresenceChangedFlag value from the given readable
func (pcf *PresenceChangedFlag) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*pcf = PresenceChangedFlag(value)
	return nil
}

// HasFlag checks if a given flag is set
func (pcf PresenceChangedFlag) HasFlag(flag PresenceChangedFlag) bool {
	return pcf&flag == flag
}

// HasFlag checks if all given flags are set
func (pcf PresenceChangedFlag) HasFlags(flags ...PresenceChangedFlag) bool {
	if len(flags) == 0 {
		return false
	}

	for _, flag := range flags {
		if pcf&flag != flag {
			return false
		}
	}

	return true
}

// String returns a human-readable representation of the PresenceChangedFlag bitmask.
// Multiple flags are joined with "|", e.g. "GameKey|GatheringID".
// Returns "None" if no flags are set.
func (pcf PresenceChangedFlag) String() string {
	if pcf == PresenceChangedFlagNone {
		return "None"
	}

	flags := []struct {
		flag PresenceChangedFlag
		name string
	}{
		{PresenceChangedFlagGameKey, "GameKey"},
		{PresenceChangedFlagGameModeDescription, "GameModeDescription"},
		{PresenceChangedFlagJoinAvailabilityFlag, "JoinAvailabilityFlag"},
		{PresenceChangedFlagMatchmakeSystemType, "MatchmakeSystemType"},
		{PresenceChangedFlagGameServerID, "GameServerID"},
		{PresenceChangedFlagJoinGameMode, "JoinGameMode"},
		{PresenceChangedFlagOwnerPID, "OwnerPID"},
		{PresenceChangedFlagGatheringID, "GatheringID"},
		{PresenceChangedFlagApplicationData, "ApplicationData"},
	}

	var parts []string
	for _, f := range flags {
		if pcf&f.flag != 0 {
			parts = append(parts, f.name)
		}
	}

	return strings.Join(parts, "|")
}

const (
	// PresenceChangedFlagNone means no data has changed
	PresenceChangedFlagNone PresenceChangedFlag = 0x0

	// PresenceChangedFlagGameKey means the users GameKey has changed.
	// Unverified. This is based off the same flag as documented in the 3DS
	PresenceChangedFlagGameKey PresenceChangedFlag = 0x1

	// PresenceChangedFlagGameModeDescription means that the "game mode description" (comment?)
	// has changed.
	// Unverified. This is based off the same flag as documented in the 3DS
	PresenceChangedFlagGameModeDescription PresenceChangedFlag = 0x2

	// PresenceChangedFlagJoinAvailabilityFlag means that the "join availability flag" ("Is online"?)
	// has changed.
	// Unverified. This is based off the same flag as documented in the 3DS
	PresenceChangedFlagJoinAvailabilityFlag PresenceChangedFlag = 0x4

	// PresenceChangedFlagMatchmakeSystemType means the gatherings matchmake system type has changed.
	// Unverified. This is based off the same flag as documented in the 3DS
	PresenceChangedFlagMatchmakeSystemType PresenceChangedFlag = 0x8

	// PresenceChangedFlagGameServerID means that the users game server ID has changed
	PresenceChangedFlagGameServerID PresenceChangedFlag = 0x10

	// PresenceChangedFlagJoinGameMode means that the gatherings "join game mode" (JoinMatchmakeSessionBehavior?)
	// has changed.
	// Unverified. This is based off the same flag as documented in the 3DS
	PresenceChangedFlagJoinGameMode PresenceChangedFlag = 0x20

	// PresenceChangedFlagOwnerPID means that the owner of the gathering has changed.
	// Unverified. This is based off the same flag as documented in the 3DS
	PresenceChangedFlagOwnerPID PresenceChangedFlag = 0x40

	// PresenceChangedFlagGatheringID means that the users gathering ID has changed
	PresenceChangedFlagGatheringID PresenceChangedFlag = 0x80

	// PresenceChangedFlagApplicationData means that the users application data has changed
	PresenceChangedFlagApplicationData PresenceChangedFlag = 0x100
)
