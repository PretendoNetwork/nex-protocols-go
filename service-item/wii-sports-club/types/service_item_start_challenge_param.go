// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemStartChallengeParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemStartChallengeParam struct {
	nex.Structure
	ChallengeScheduleID uint32
	TicketType          uint32
	NumTicket           uint32
}

// ExtractFromStream extracts a ServiceItemStartChallengeParam structure from a stream
func (serviceItemStartChallengeParam *ServiceItemStartChallengeParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemStartChallengeParam.ChallengeScheduleID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemStartChallengeParam.ChallengeScheduleID from stream. %s", err.Error())
	}

	serviceItemStartChallengeParam.TicketType, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemStartChallengeParam.TicketType from stream. %s", err.Error())
	}

	serviceItemStartChallengeParam.NumTicket, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemStartChallengeParam.NumTicket from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemStartChallengeParam and returns a byte array
func (serviceItemStartChallengeParam *ServiceItemStartChallengeParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemStartChallengeParam.ChallengeScheduleID)
	stream.WriteUInt32LE(serviceItemStartChallengeParam.TicketType)
	stream.WriteUInt32LE(serviceItemStartChallengeParam.NumTicket)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemStartChallengeParam
func (serviceItemStartChallengeParam *ServiceItemStartChallengeParam) Copy() nex.StructureInterface {
	copied := NewServiceItemStartChallengeParam()

	copied.SetStructureVersion(serviceItemStartChallengeParam.StructureVersion())

	copied.ChallengeScheduleID = serviceItemStartChallengeParam.ChallengeScheduleID
	copied.TicketType = serviceItemStartChallengeParam.TicketType
	copied.NumTicket = serviceItemStartChallengeParam.NumTicket

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemStartChallengeParam *ServiceItemStartChallengeParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemStartChallengeParam)

	if serviceItemStartChallengeParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemStartChallengeParam.ChallengeScheduleID != other.ChallengeScheduleID {
		return false
	}

	if serviceItemStartChallengeParam.TicketType != other.TicketType {
		return false
	}

	if serviceItemStartChallengeParam.NumTicket != other.NumTicket {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemStartChallengeParam *ServiceItemStartChallengeParam) String() string {
	return serviceItemStartChallengeParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemStartChallengeParam *ServiceItemStartChallengeParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemStartChallengeParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemStartChallengeParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sChallengeScheduleID: %d,\n", indentationValues, serviceItemStartChallengeParam.ChallengeScheduleID))
	b.WriteString(fmt.Sprintf("%sTicketType: %d,\n", indentationValues, serviceItemStartChallengeParam.TicketType))
	b.WriteString(fmt.Sprintf("%sNumTicket: %d,\n", indentationValues, serviceItemStartChallengeParam.NumTicket))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemStartChallengeParam returns a new ServiceItemStartChallengeParam
func NewServiceItemStartChallengeParam() *ServiceItemStartChallengeParam {
	return &ServiceItemStartChallengeParam{}
}
