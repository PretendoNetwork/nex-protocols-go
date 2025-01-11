// Package types implements all the types used by the Matchmaking protocol
package types

import "github.com/PretendoNetwork/nex-go/v2/types"

// DataInterface defines an interface to track types which have Gathering anywhere
// in their parent tree.
type GatheringInterface interface {
	types.HoldableObject
	GatheringObjectID() types.RVType // Returns the object identifier of the type embedding Gathering
}

// GatheringHolder is an AnyObjectHolder for types which embed Gathering
type GatheringHolder = types.AnyObjectHolder[GatheringInterface]

// NewGatheringHolder returns a new GatheringHolder
func NewGatheringHolder() GatheringHolder {
	return GatheringHolder{}
}
