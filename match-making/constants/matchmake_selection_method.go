package constants

import (
	"github.com/PretendoNetwork/nex-go/v2/types"
)

// MatchmakeSelectionMethod is used to indicate the selection method used when selecting a gathering
type MatchmakeSelectionMethod uint32

// WriteTo writes the MatchmakeSelectionMethod to the given writable
func (msm MatchmakeSelectionMethod) WriteTo(writable types.Writable) {
	writable.WriteUInt32LE(uint32(msm))
}

// ExtractFrom extracts the MatchmakeSelectionMethod value from the given readable
func (msm *MatchmakeSelectionMethod) ExtractFrom(readable types.Readable) error {
	value, err := readable.ReadUInt32LE()
	if err != nil {
		return err
	}

	*msm = MatchmakeSelectionMethod(value)
	return nil
}

const (
	// MatchmakeSelectionMethodRandom indicates a random selection
	MatchmakeSelectionMethodRandom MatchmakeSelectionMethod = iota

	// MatchmakeSelectionMethodNearestNeighbor selects the session with the closest
	// matching Attributes[MatchmakeSessionNearestNeighborAttributeIndex] value. If
	// multiple are found, pick the one with the highest gathering ID.
	MatchmakeSelectionMethodNearestNeighbor

	// MatchmakeSelectionMethodBroadenRange uses an ELO-based selection method. The
	// value of Attributes[MatchmakeSessionBroadenRangeAttributeIndex] is used to
	// select the users current ELO rating pool. We do not know the real values of
	// the ELO pools, and likely never will, but we can make "good enough" educated
	// guesses.
	//
	// Nintendo seems to have used 11/12 rating pools, with values starting from 0,
	// based on tests done in Splatoon. 11 is an awkward number and not likely used,
	// so we can assume 12. Based on observations of various ranks these ranges
	// should be "good enough" (though may need adjusting on a per-game basis):
	//
	// Pool 1	<1100
	// Pool 2	1100-1199
	// Pool 3	1200-1299
	// Pool 4	1300-1399
	// Pool 5	1400-1499
	// Pool 6	1500-1599
	// Pool 7	1600-1699
	// Pool 8	1700-1799
	// Pool 9	1800-1899
	// Pool 10	1900-1999
	// Pool 11	2000-2199
	// Pool 12	2200+
	//
	// When this method is used, sessions are selected by the search criteria, then
	// filtered by the ELO rating, and a random valid session is returned. This way
	// users are pooled into similar skill levels.
	//
	// If no valid session is found, widen the ELO rating window to those outside the
	// users pool. The exact widening method is unknown, but it should be safe to assume
	// it widens to other ELO pools before/after the users pool slowly (such as 1 range
	// at a time) until one is found.
	MatchmakeSelectionMethodBroadenRange

	// MatchmakeSelectionMethodProgressScore first selects session by the search criteria,
	// then select the one with the highest progress score. If multiple valid sessions are
	// found, pick the one with the highest gathering ID.
	MatchmakeSelectionMethodProgressScore

	// MatchmakeSelectionMethodBroadenRange selects sessions using the same method as
	// MatchmakeSelectionMethodNearestNeighbor, but instead of selecting a valid session
	// at random select the valid session using logic from MatchmakeSelectionMethodProgressScore.
	MatchmakeSelectionMethodBroadenRangeWithProgressScore

	// MatchmakeSelectionMethodScoreBased uses unknown logic to select the session.
	MatchmakeSelectionMethodScoreBased
)
