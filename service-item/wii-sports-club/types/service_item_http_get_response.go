// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemHTTPGetResponse holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemHTTPGetResponse struct {
	nex.Structure
	Response []byte
}

// ExtractFromStream extracts a ServiceItemHTTPGetResponse structure from a stream
func (serviceItemHTTPGetResponse *ServiceItemHTTPGetResponse) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemHTTPGetResponse.Response, err = stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemHTTPGetResponse.Response from stream. %s", err.Error())
	}

	return nil
}

// Bytes encodes the ServiceItemHTTPGetResponse and returns a byte array
func (serviceItemHTTPGetResponse *ServiceItemHTTPGetResponse) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteQBuffer(serviceItemHTTPGetResponse.Response)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemHTTPGetResponse
func (serviceItemHTTPGetResponse *ServiceItemHTTPGetResponse) Copy() nex.StructureInterface {
	copied := NewServiceItemHTTPGetResponse()

	copied.SetStructureVersion(serviceItemHTTPGetResponse.StructureVersion())

	copied.Response = serviceItemHTTPGetResponse.Response

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemHTTPGetResponse *ServiceItemHTTPGetResponse) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemHTTPGetResponse)

	if serviceItemHTTPGetResponse.StructureVersion() != other.StructureVersion() {
		return false
	}

	return bytes.Equal(serviceItemHTTPGetResponse.Response, other.Response)
}

// String returns a string representation of the struct
func (serviceItemHTTPGetResponse *ServiceItemHTTPGetResponse) String() string {
	return serviceItemHTTPGetResponse.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemHTTPGetResponse *ServiceItemHTTPGetResponse) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemHTTPGetResponse{\n")
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemHTTPGetResponse.StructureVersion()))
	b.WriteString(fmt.Sprintf("%sResponse: %x,\n", indentationValues, serviceItemHTTPGetResponse.Response))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemHTTPGetResponse returns a new ServiceItemHTTPGetResponse
func NewServiceItemHTTPGetResponse() *ServiceItemHTTPGetResponse {
	return &ServiceItemHTTPGetResponse{}
}
