// Package types implements all the types used by the Matchmaking protocols.
//
// Since there are multiple match making related protocols, and they all share types
// all types used by all match making protocols is defined here
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// Gathering holds information about a matchmake gathering
type Gathering struct {
	types.Structure
	ID                  *types.PrimitiveU32
	OwnerPID            *types.PID
	HostPID             *types.PID
	MinimumParticipants *types.PrimitiveU16
	MaximumParticipants *types.PrimitiveU16
	ParticipationPolicy *types.PrimitiveU32
	PolicyArgument      *types.PrimitiveU32
	Flags               *types.PrimitiveU32
	State               *types.PrimitiveU32
	Description         string
}

// ExtractFrom extracts the Gathering from the given readable
func (gathering *Gathering) ExtractFrom(readable types.Readable) error {
	var err error

	if err = gathering.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read Gathering header. %s", err.Error())
	}

	err = gathering.ID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.ID. %s", err.Error())
	}

	err = gathering.OwnerPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.OwnerPID. %s", err.Error())
	}

	err = gathering.HostPID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.HostPID. %s", err.Error())
	}

	err = gathering.MinimumParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.MinimumParticipants. %s", err.Error())
	}

	err = gathering.MaximumParticipants.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.MaximumParticipants. %s", err.Error())
	}

	err = gathering.ParticipationPolicy.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.ParticipationPolicy. %s", err.Error())
	}

	err = gathering.PolicyArgument.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.PolicyArgument. %s", err.Error())
	}

	err = gathering.Flags.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.Flags. %s", err.Error())
	}

	err = gathering.State.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.State. %s", err.Error())
	}

	err = gathering.Description.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.Description. %s", err.Error())
	}

	return nil
}

// WriteTo writes the Gathering to the given writable
func (gathering *Gathering) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	gathering.ID.WriteTo(contentWritable)
	gathering.OwnerPID.WriteTo(contentWritable)
	gathering.HostPID.WriteTo(contentWritable)
	gathering.MinimumParticipants.WriteTo(contentWritable)
	gathering.MaximumParticipants.WriteTo(contentWritable)
	gathering.ParticipationPolicy.WriteTo(contentWritable)
	gathering.PolicyArgument.WriteTo(contentWritable)
	gathering.Flags.WriteTo(contentWritable)
	gathering.State.WriteTo(contentWritable)
	gathering.Description.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	gathering.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of Gathering
func (gathering *Gathering) Copy() types.RVType {
	copied := NewGathering()

	copied.StructureVersion = gathering.StructureVersion

	copied.ID = gathering.ID

	copied.OwnerPID = gathering.OwnerPID.Copy()

	copied.HostPID = gathering.HostPID.Copy()

	copied.MinimumParticipants = gathering.MinimumParticipants
	copied.MaximumParticipants = gathering.MaximumParticipants
	copied.ParticipationPolicy = gathering.ParticipationPolicy
	copied.PolicyArgument = gathering.PolicyArgument
	copied.Flags = gathering.Flags
	copied.State = gathering.State
	copied.Description = gathering.Description

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (gathering *Gathering) Equals(o types.RVType) bool {
	if _, ok := o.(*Gathering); !ok {
		return false
	}

	other := o.(*Gathering)

	if gathering.StructureVersion != other.StructureVersion {
		return false
	}

	if !gathering.ID.Equals(other.ID) {
		return false
	}

	if !gathering.OwnerPID.Equals(other.OwnerPID) {
		return false
	}

	if !gathering.HostPID.Equals(other.HostPID) {
		return false
	}

	if !gathering.MinimumParticipants.Equals(other.MinimumParticipants) {
		return false
	}

	if !gathering.MaximumParticipants.Equals(other.MaximumParticipants) {
		return false
	}

	if !gathering.ParticipationPolicy.Equals(other.ParticipationPolicy) {
		return false
	}

	if !gathering.PolicyArgument.Equals(other.PolicyArgument) {
		return false
	}

	if !gathering.Flags.Equals(other.Flags) {
		return false
	}

	if !gathering.State.Equals(other.State) {
		return false
	}

	if !gathering.Description.Equals(other.Description) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (gathering *Gathering) String() string {
	return gathering.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (gathering *Gathering) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("Gathering{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, gathering.StructureVersion))
	b.WriteString(fmt.Sprintf("%sID: %d,\n", indentationValues, gathering.ID))
	b.WriteString(fmt.Sprintf("%sOwnerPID: %s,\n", indentationValues, gathering.OwnerPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sHostPID: %s,\n", indentationValues, gathering.HostPID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sMinimumParticipants: %d,\n", indentationValues, gathering.MinimumParticipants))
	b.WriteString(fmt.Sprintf("%sMaximumParticipants: %d,\n", indentationValues, gathering.MaximumParticipants))
	b.WriteString(fmt.Sprintf("%sParticipationPolicy: %d,\n", indentationValues, gathering.ParticipationPolicy))
	b.WriteString(fmt.Sprintf("%sPolicyArgument: %d,\n", indentationValues, gathering.PolicyArgument))
	b.WriteString(fmt.Sprintf("%sFlags: %d,\n", indentationValues, gathering.Flags))
	b.WriteString(fmt.Sprintf("%sState: %d,\n", indentationValues, gathering.State))
	b.WriteString(fmt.Sprintf("%sDescription: %q\n", indentationValues, gathering.Description))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewGathering returns a new Gathering
func NewGathering() *Gathering {
	return &Gathering{}
}
