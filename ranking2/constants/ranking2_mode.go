package constants

import (
	"fmt"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// Ranking2Mode determines what rankings to return and how to order them.
//
// Modes refer to a "selected user". This user can be either the caller, OR
// a different user, based on the input parameters. The NEX unique ID and PID
// fields determine what user is the "selected user", defaulting to the caller
// if neither are used. Only up to 100 rankings can be selected at a time.
type Ranking2Mode uint8

// WriteTo writes the Ranking2Mode to the given writable
func (r2m Ranking2Mode) WriteTo(writable types.Writable) {
	writable.WriteUInt8(uint8(r2m))
}

// ExtractFrom extracts the Ranking2Mode value from the given readable
func (r2m *Ranking2Mode) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt8()
	if err != nil {
		return err
	}

	*r2m = Ranking2Mode(value)
	return nil
}

// String returns a human-readable representation of the Ranking2Mode.
func (r2m Ranking2Mode) String() string {
	switch r2m {
	case Ranking2ModeUserRanking:
		return "UserRanking"
	case Ranking2ModeNearRanking:
		return "NearRanking"
	case Ranking2ModeRangeRanking:
		return "RangeRanking"
	case Ranking2ModeFriendRanking:
		return "FriendRanking"
	default:
		return fmt.Sprintf("Ranking2Mode(%d)", int(r2m))
	}
}

const (
	// Ranking2ModeMin is the min value for the Ranking2Mode enum.
	Ranking2ModeMin Ranking2Mode = 0

	// Ranking2ModeUserRanking limits the returned rankings to those belonging
	// only to the selected user. If a unique ID is not set in the paramater
	// and the selected user has associated unique IDs, get rankings from all
	// associated unique IDs instead.
	Ranking2ModeUserRanking Ranking2Mode = 0

	// Ranking2ModeNearRanking returns rankings around the selected user. "Around"
	// means attempting to place the selected user in the middle of the results
	// as best as the server can.
	Ranking2ModeNearRanking Ranking2Mode = 1

	// Ranking2ModeRangeRanking returns global rankings using the given offset/length.
	Ranking2ModeRangeRanking Ranking2Mode = 2

	// Ranking2ModeFriendRanking returns rankings for the caller and all friends. This
	// includes rankings from associated NEX unique IDs.
	Ranking2ModeFriendRanking Ranking2Mode = 3

	// Ranking2ModeMax is the max value for the Ranking2Mode enum.
	Ranking2ModeMax Ranking2Mode = 3
)
