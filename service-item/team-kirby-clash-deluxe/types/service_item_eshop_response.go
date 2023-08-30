// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemEShopResponse holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemEShopResponse struct {
	nex.Structure
	HTTPStatus    uint32
	ErrorCode     uint32
	CorrelationID string
}

// ExtractFromStream extracts a ServiceItemEShopResponse structure from a stream
func (serviceItemEShopResponse *ServiceItemEShopResponse) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemEShopResponse.HTTPStatus, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEShopResponse.HTTPStatus from stream. %s", err.Error())
	}

	serviceItemEShopResponse.ErrorCode, err = stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEShopResponse.ErrorCode from stream. %s", err.Error())
	}

	serviceItemEShopResponse.CorrelationID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemEShopResponse.CorrelationID from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemEShopResponse and returns a byte array
func (serviceItemEShopResponse *ServiceItemEShopResponse) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(serviceItemEShopResponse.HTTPStatus)
	stream.WriteUInt32LE(serviceItemEShopResponse.ErrorCode)
	stream.WriteString(serviceItemEShopResponse.CorrelationID)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemEShopResponse
func (serviceItemEShopResponse *ServiceItemEShopResponse) Copy() nex.StructureInterface {
	copied := NewServiceItemEShopResponse()

	copied.SetStructureVersion(serviceItemEShopResponse.StructureVersion())

	copied.HTTPStatus = serviceItemEShopResponse.HTTPStatus
	copied.ErrorCode = serviceItemEShopResponse.ErrorCode
	copied.CorrelationID = serviceItemEShopResponse.CorrelationID

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemEShopResponse *ServiceItemEShopResponse) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemEShopResponse)

	if serviceItemEShopResponse.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemEShopResponse.HTTPStatus != other.HTTPStatus {
		return false
	}

	if serviceItemEShopResponse.ErrorCode != other.ErrorCode {
		return false
	}

	if serviceItemEShopResponse.CorrelationID != other.CorrelationID {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemEShopResponse *ServiceItemEShopResponse) String() string {
	return serviceItemEShopResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemEShopResponse *ServiceItemEShopResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemEShopResponse{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemEShopResponse.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sHTTPStatus: %d,\n", indentationValues, serviceItemEShopResponse.HTTPStatus))
	b.WriteString(fmt.Sprintf("%sErrorCode: %d,\n", indentationValues, serviceItemEShopResponse.ErrorCode))
	b.WriteString(fmt.Sprintf("%sCorrelationID: %q,\n", indentationValues, serviceItemEShopResponse.CorrelationID))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemEShopResponse returns a new ServiceItemEShopResponse
func NewServiceItemEShopResponse() *ServiceItemEShopResponse {
	return &ServiceItemEShopResponse{}
}
