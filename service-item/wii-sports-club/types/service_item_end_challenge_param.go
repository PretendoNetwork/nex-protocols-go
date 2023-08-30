// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemEndChallengeParam holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemEndChallengeParam struct {
	nex.Structure
	ChallengeScheduleID uint32
	UserInfo            *ServiceItemUserInfo
}

// ExtractFromStream extracts a ServiceItemEndChallengeParam structure from a stream
func (serviceItemEndChallengeParam *ServiceItemEndChallengeParam) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemEndChallengeParam.ChallengeScheduleID, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEndChallengeParam.ChallengeScheduleID from stream. %s", err.Error())
	}

	userInfo, err := stream.ReadStructure(NewServiceItemUserInfo())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEndChallengeParam.UserInfo from stream. %s", err.Error())
	}

	serviceItemEndChallengeParam.UserInfo = userInfo.(*ServiceItemUserInfo)

	return nil
}

// Bytes encodes the ServiceItemEndChallengeParam and returns a byte array
func (serviceItemEndChallengeParam *ServiceItemEndChallengeParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemEndChallengeParam.ChallengeScheduleID)
	stream.WriteStructure(serviceItemEndChallengeParam.UserInfo)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemEndChallengeParam
func (serviceItemEndChallengeParam *ServiceItemEndChallengeParam) Copy() nex.StructureInterface {
	copied := NewServiceItemEndChallengeParam()

	copied.SetStructureVersion(serviceItemEndChallengeParam.StructureVersion())

	copied.ChallengeScheduleID = serviceItemEndChallengeParam.ChallengeScheduleID
	copied.UserInfo = serviceItemEndChallengeParam.UserInfo.Copy().(*ServiceItemUserInfo)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemEndChallengeParam *ServiceItemEndChallengeParam) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemEndChallengeParam)

	if serviceItemEndChallengeParam.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemEndChallengeParam.ChallengeScheduleID != other.ChallengeScheduleID {
		return false
	}

	if !serviceItemEndChallengeParam.UserInfo.Equals(other.UserInfo) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemEndChallengeParam *ServiceItemEndChallengeParam) String() string {
	return serviceItemEndChallengeParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemEndChallengeParam *ServiceItemEndChallengeParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemEndChallengeParam{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemEndChallengeParam.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sChallengeScheduleID: %d,\n", indentationValues, serviceItemEndChallengeParam.ChallengeScheduleID))

	if serviceItemEndChallengeParam.UserInfo != nil {
		b.WriteString(fmt.Sprintf("%sUserInfo: %s\n", indentationValues, serviceItemEndChallengeParam.UserInfo.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sUserInfo: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemEndChallengeParam returns a new ServiceItemEndChallengeParam
func NewServiceItemEndChallengeParam() *ServiceItemEndChallengeParam {
	return &ServiceItemEndChallengeParam{}
}
