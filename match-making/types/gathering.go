package match_making_types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// Gathering holds information about a matchmake gathering
type Gathering struct {
	nex.Structure
	ID                  uint32
	OwnerPID            uint32
	HostPID             uint32
	MinimumParticipants uint16
	MaximumParticipants uint16
	ParticipationPolicy uint32
	PolicyArgument      uint32
	Flags               uint32
	State               uint32
	Description         string
}

// ExtractFromStream extracts a Gathering structure from a stream
func (gathering *Gathering) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	gathering.ID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.ID. %s", err.Error())
	}

	gathering.OwnerPID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.OwnerPID. %s", err.Error())
	}

	gathering.HostPID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.HostPID. %s", err.Error())
	}

	gathering.MinimumParticipants, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.MinimumParticipants. %s", err.Error())
	}

	gathering.MaximumParticipants, err = stream.ReadUInt16LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.MaximumParticipants. %s", err.Error())
	}

	gathering.ParticipationPolicy, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.ParticipationPolicy. %s", err.Error())
	}

	gathering.PolicyArgument, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.PolicyArgument. %s", err.Error())
	}

	gathering.Flags, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.Flags. %s", err.Error())
	}

	gathering.State, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.State. %s", err.Error())
	}

	gathering.Description, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract Gathering.Description. %s", err.Error())
	}

	return nil
}

// Bytes encodes the Gathering and returns a byte array
func (gathering *Gathering) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(gathering.ID)
	stream.WriteUInt32LE(gathering.OwnerPID)
	stream.WriteUInt32LE(gathering.HostPID)
	stream.WriteUInt16LE(gathering.MinimumParticipants)
	stream.WriteUInt16LE(gathering.MaximumParticipants)
	stream.WriteUInt32LE(gathering.ParticipationPolicy)
	stream.WriteUInt32LE(gathering.PolicyArgument)
	stream.WriteUInt32LE(gathering.Flags)
	stream.WriteUInt32LE(gathering.State)
	stream.WriteString(gathering.Description)

	return stream.Bytes()
}

// Copy returns a new copied instance of Gathering
func (gathering *Gathering) Copy() nex.StructureInterface {
	copied := NewGathering()

	copied.ID = gathering.ID
	copied.OwnerPID = gathering.OwnerPID
	copied.HostPID = gathering.HostPID
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
func (gathering *Gathering) Equals(structure nex.StructureInterface) bool {
	other := structure.(*Gathering)

	if gathering.ID != other.ID {
		return false
	}

	if gathering.OwnerPID != other.OwnerPID {
		return false
	}

	if gathering.HostPID != other.HostPID {
		return false
	}

	if gathering.MinimumParticipants != other.MinimumParticipants {
		return false
	}

	if gathering.MaximumParticipants != other.MaximumParticipants {
		return false
	}

	if gathering.ParticipationPolicy != other.ParticipationPolicy {
		return false
	}

	if gathering.PolicyArgument != other.PolicyArgument {
		return false
	}

	if gathering.Flags != other.Flags {
		return false
	}

	if gathering.State != other.State {
		return false
	}

	if gathering.Description != other.Description {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, gathering.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sID: %d,\n", indentationValues, gathering.ID))
	b.WriteString(fmt.Sprintf("%sOwnerPID: %d,\n", indentationValues, gathering.OwnerPID))
	b.WriteString(fmt.Sprintf("%sHostPID: %d,\n", indentationValues, gathering.HostPID))
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
