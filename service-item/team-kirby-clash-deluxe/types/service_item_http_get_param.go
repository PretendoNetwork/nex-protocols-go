// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemHTTPGetParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemHTTPGetParam struct {
	types.Structure
	URL string
}

// ExtractFrom extracts the ServiceItemHTTPGetParam from the given readable
func (serviceItemHTTPGetParam *ServiceItemHTTPGetParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemHTTPGetParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemHTTPGetParam header. %s", err.Error())
	}

	err = serviceItemHTTPGetParam.URL.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemHTTPGetParam.URL from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemHTTPGetParam to the given writable
func (serviceItemHTTPGetParam *ServiceItemHTTPGetParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemHTTPGetParam.URL.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemHTTPGetParam.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemHTTPGetParam
func (serviceItemHTTPGetParam *ServiceItemHTTPGetParam) Copy() types.RVType {
	copied := NewServiceItemHTTPGetParam()

	copied.StructureVersion = serviceItemHTTPGetParam.StructureVersion

	copied.URL = serviceItemHTTPGetParam.URL

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemHTTPGetParam *ServiceItemHTTPGetParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemHTTPGetParam); !ok {
		return false
	}

	other := o.(*ServiceItemHTTPGetParam)

	if serviceItemHTTPGetParam.StructureVersion != other.StructureVersion {
		return false
	}

	return serviceItemHTTPGetParam.URL == other.URL
}

// String returns a string representation of the struct
func (serviceItemHTTPGetParam *ServiceItemHTTPGetParam) String() string {
	return serviceItemHTTPGetParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemHTTPGetParam *ServiceItemHTTPGetParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemHTTPGetParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemHTTPGetParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sURL: %q,\n", indentationValues, serviceItemHTTPGetParam.URL))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemHTTPGetParam returns a new ServiceItemHTTPGetParam
func NewServiceItemHTTPGetParam() *ServiceItemHTTPGetParam {
	return &ServiceItemHTTPGetParam{}
}
