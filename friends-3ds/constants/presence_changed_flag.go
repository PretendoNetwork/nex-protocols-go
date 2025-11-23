package constants

import "github.com/PretendoNetwork/nex-go/v2/types"

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

const (
	// PresenceChangedFlagNone means no data has changed
	PresenceChangedFlagNone PresenceChangedFlag = 0x0

	// PresenceChangedFlagGameKey means that `m_gameKey` has changed
	PresenceChangedFlagGameKey PresenceChangedFlag = 0x1

	// PresenceChangedFlagGameModeDescription means that `m_gameModeDescription` has changed
	PresenceChangedFlagGameModeDescription PresenceChangedFlag = 0x2

	// PresenceChangedFlagJoinAvailabilityFlag means that `m_joinAvailabilityFlag` has changed
	PresenceChangedFlagJoinAvailabilityFlag PresenceChangedFlag = 0x4

	// PresenceChangedFlagMatchmakeSystemType means that `m_matchmakeSystemType` has changed
	PresenceChangedFlagMatchmakeSystemType PresenceChangedFlag = 0x8

	// PresenceChangedFlagJoinGameID means that `m_joinGameID` (game server ID) has changed
	PresenceChangedFlagJoinGameID PresenceChangedFlag = 0x10

	// PresenceChangedFlagJoinGameMode means that `m_joinGameMode` (JoinMatchmakeSessionBehavior?)
	// has changed
	PresenceChangedFlagJoinGameMode PresenceChangedFlag = 0x20

	// PresenceChangedFlagOwnerPID means that `m_ownerPrincipalID` has changed
	PresenceChangedFlagOwnerPID PresenceChangedFlag = 0x40

	// PresenceChangedFlagGatheringID means that `m_joinGroupID` (gathering ID) has changed
	PresenceChangedFlagGatheringID PresenceChangedFlag = 0x80

	// PresenceChangedFlagApplicationData means that `m_applicationArg` (application data) has changed
	PresenceChangedFlagApplicationData PresenceChangedFlag = 0x100
)
