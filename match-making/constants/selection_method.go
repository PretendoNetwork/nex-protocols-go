package constants

// SelectionMethod is used to indicate the selection method used when selecting a gathering
type SelectionMethod uint32

const (
	// SelectionMethodRandom indicates a random selection
	SelectionMethodRandom SelectionMethod = iota

	// SelectionMethodNearestNeighbor indicates a selection based on proximity to an attribute
	SelectionMethodNearestNeighbor

	// SelectionMethodBroadenRange indicates a ranked selection
	SelectionMethodBroadenRange

	// SelectionMethodProgressScore indicates a selection based on the progress score
	SelectionMethodProgressScore

	// SelectionMethodBroadenRange indicates a ranked selection based on the progress score
	SelectionMethodBroadenRangeWithProgressScore

	// SelectionMethodScoreBased indicates an unknown selection
	SelectionMethodScoreBased
)
