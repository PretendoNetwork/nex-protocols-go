package constants

// MatchmakeSelectionMethod is used to indicate the selection method used when selecting a gathering
type MatchmakeSelectionMethod uint32

const (
	// MatchmakeSelectionMethodRandom indicates a random selection
	MatchmakeSelectionMethodRandom MatchmakeSelectionMethod = iota

	// MatchmakeSelectionMethodNearestNeighbor indicates a selection based on proximity to an attribute
	MatchmakeSelectionMethodNearestNeighbor

	// MatchmakeSelectionMethodBroadenRange indicates a ranked selection
	MatchmakeSelectionMethodBroadenRange

	// MatchmakeSelectionMethodProgressScore indicates a selection based on the progress score
	MatchmakeSelectionMethodProgressScore

	// MatchmakeSelectionMethodBroadenRange indicates a ranked selection based on the progress score
	MatchmakeSelectionMethodBroadenRangeWithProgressScore

	// MatchmakeSelectionMethodScoreBased indicates an unknown selection
	MatchmakeSelectionMethodScoreBased
)
