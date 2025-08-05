// Package types implements all the types used by the Matchmaking protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/match-making/constants"
)

// Gathering is a type within the Matchmaking protocol
type Gathering struct {
	types.Structure
	ID                  types.UInt32
	OwnerPID            types.PID
	HostPID             types.PID
	MinimumParticipants types.UInt16
	MaximumParticipants types.UInt16
	ParticipationPolicy constants.ParticipationPolicy
	PolicyArgument      constants.PolicyArgument
	Flags               constants.GatheringFlags
	State               constants.GatheringState
	Description         types.String
}

// ObjectID returns the object identifier of the type
func (g Gathering) ObjectID() types.RVType {
	return g.GatheringObjectID()
}

// DataObjectID returns the object identifier of the type embedding Gathering
func (g Gathering) GatheringObjectID() types.RVType {
	return types.NewString("Gathering")
}

// WriteTo writes the Gathering to the given writable
func (g Gathering) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	g.ID.WriteTo(contentWritable)
	g.OwnerPID.WriteTo(contentWritable)
	g.HostPID.WriteTo(contentWritable)
	g.MinimumParticipants.WriteTo(contentWritable)
	g.MaximumParticipants.WriteTo(contentWritable)
	types.UInt32(g.ParticipationPolicy).WriteTo(contentWritable)
	types.UInt32(g.PolicyArgument).WriteTo(contentWritable)
	types.UInt16(g.Flags).WriteTo(contentWritable)
	types.UInt32(g.State).WriteTo(contentWritable)
	g.Description.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	g.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the Gathering from the given readable
func (g *Gathering) ExtractFrom(readable types.Readable) error {
	var err error

	err = g.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering header. %s", err.Error())
	}

	err = g.ID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.ID. %s", err.Error())
	}

	err = g.OwnerPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.OwnerPID. %s", err.Error())
	}

	err = g.HostPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.HostPID. %s", err.Error())
	}

	err = g.MinimumParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.MinimumParticipants. %s", err.Error())
	}

	err = g.MaximumParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.MaximumParticipants. %s", err.Error())
	}

	participationPolicy, err := readable.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.ParticipationPolicy. %s", err.Error())
	}

	g.ParticipationPolicy = constants.ParticipationPolicy(participationPolicy)
	if !g.ParticipationPolicy.IsValid() {
		return fmt.Errorf("Gathering.ParticipationPolicy is invalid. Value %d is out of range", participationPolicy)
	}

	policyArgument, err := readable.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.PolicyArgument. %s", err.Error())
	}

	g.PolicyArgument = constants.PolicyArgument(policyArgument)

	flags, err := readable.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.Flags. %s", err.Error())
	}

	g.Flags = constants.GatheringFlags(flags)

	state, err := readable.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.State. %s", err.Error())
	}

	g.State = constants.GatheringState(state)
	if !g.State.IsValid() {
		return fmt.Errorf("Gathering.State is invalid. Value %d is out of range", state)
	}

	err = g.Description.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.Description. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of Gathering
func (g Gathering) Copy() types.RVType {
	copied := NewGathering()

	copied.StructureVersion = g.StructureVersion
	copied.ID = g.ID.Copy().(types.UInt32)
	copied.OwnerPID = g.OwnerPID.Copy().(types.PID)
	copied.HostPID = g.HostPID.Copy().(types.PID)
	copied.MinimumParticipants = g.MinimumParticipants.Copy().(types.UInt16)
	copied.MaximumParticipants = g.MaximumParticipants.Copy().(types.UInt16)
	copied.ParticipationPolicy = g.ParticipationPolicy
	copied.PolicyArgument = g.PolicyArgument
	copied.Flags = g.Flags
	copied.State = g.State
	copied.Description = g.Description.Copy().(types.String)

	return copied
}

// Equals checks if the given Gathering contains the same data as the current Gathering
func (g Gathering) Equals(o types.RVType) bool {
	if _, ok := o.(Gathering); !ok {
		return false
	}

	other := o.(Gathering)

	if g.StructureVersion != other.StructureVersion {
		return false
	}

	if !g.ID.Equals(other.ID) {
		return false
	}

	if !g.OwnerPID.Equals(other.OwnerPID) {
		return false
	}

	if !g.HostPID.Equals(other.HostPID) {
		return false
	}

	if !g.MinimumParticipants.Equals(other.MinimumParticipants) {
		return false
	}

	if !g.MaximumParticipants.Equals(other.MaximumParticipants) {
		return false
	}

	if g.ParticipationPolicy != other.ParticipationPolicy {
		return false
	}

	if g.PolicyArgument != other.PolicyArgument {
		return false
	}

	if g.Flags != other.Flags {
		return false
	}

	if g.State != other.State {
		return false
	}

	return g.Description.Equals(other.Description)
}

// CopyRef copies the current value of the Gathering
// and returns a pointer to the new copy
func (g Gathering) CopyRef() types.RVTypePtr {
	copied := g.Copy().(Gathering)
	return &copied
}

// Deref takes a pointer to the Gathering
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (g *Gathering) Deref() types.RVType {
	return *g
}

// String returns the string representation of the Gathering
func (g Gathering) String() string {
	return g.FormatToString(0)
}

// FormatToString pretty-prints the Gathering using the provided indentation level
func (g Gathering) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Gathering{\n")
	b.WriteString(fmt.Sprintf("%sID: %s,\n", indentationValues, g.ID))
	b.WriteString(fmt.Sprintf("%sOwnerPID: %s,\n", indentationValues, g.OwnerPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sHostPID: %s,\n", indentationValues, g.HostPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sMinimumParticipants: %s,\n", indentationValues, g.MinimumParticipants))
	b.WriteString(fmt.Sprintf("%sMaximumParticipants: %s,\n", indentationValues, g.MaximumParticipants))
	b.WriteString(fmt.Sprintf("%sParticipationPolicy: %s,\n", indentationValues, g.ParticipationPolicy))
	b.WriteString(fmt.Sprintf("%sPolicyArgument: %s,\n", indentationValues, g.PolicyArgument))
	b.WriteString(fmt.Sprintf("%sFlags: %s,\n", indentationValues, g.Flags))
	b.WriteString(fmt.Sprintf("%sState: %s,\n", indentationValues, g.State))
	b.WriteString(fmt.Sprintf("%sDescription: %s,\n", indentationValues, g.Description))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGathering returns a new Gathering
func NewGathering() Gathering {
	return Gathering{
		ID:                  types.NewUInt32(0),
		OwnerPID:            types.NewPID(0),
		HostPID:             types.NewPID(0),
		MinimumParticipants: types.NewUInt16(0),
		MaximumParticipants: types.NewUInt16(0),
		ParticipationPolicy: constants.ParticipationPolicy(1), // TODO - This is what the nn::nex::Gathering::Reset() function sets this to in Xenoblade. I have no idea what this value actually means
		PolicyArgument:      constants.PolicyArgument(0),      // TODO - This is what the nn::nex::Gathering::Reset() function sets this to in Xenoblade. I have no idea what this value actually means
		Flags:               constants.GatheringFlagNone,
		State:               constants.GatheringState(0), // TODO - This is what the nn::nex::Gathering::Reset() function sets this to in Xenoblade. I have no idea what this value actually means
		Description:         types.NewString(""),
	}

}
