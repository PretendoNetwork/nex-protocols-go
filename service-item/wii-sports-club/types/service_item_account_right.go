// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemAccountRight is a type within the ServiceItem protocol
type ServiceItemAccountRight struct {
	types.Structure
	PID        types.PID
	Limitation ServiceItemLimitation
}

// WriteTo writes the ServiceItemAccountRight to the given writable
func (siar ServiceItemAccountRight) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	siar.PID.WriteTo(contentWritable)
	siar.Limitation.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	siar.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemAccountRight from the given readable
func (siar *ServiceItemAccountRight) ExtractFrom(readable types.Readable) error {
	var err error

	err = siar.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight header. %s", err.Error())
	}

	err = siar.PID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.PID. %s", err.Error())
	}

	err = siar.Limitation.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemAccountRight.Limitation. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemAccountRight
func (siar ServiceItemAccountRight) Copy() types.RVType {
	copied := NewServiceItemAccountRight()

	copied.StructureVersion = siar.StructureVersion
	copied.PID = siar.PID.Copy().(types.PID)
	copied.Limitation = siar.Limitation.Copy().(ServiceItemLimitation)

	return copied
}

// Equals checks if the given ServiceItemAccountRight contains the same data as the current ServiceItemAccountRight
func (siar ServiceItemAccountRight) Equals(o types.RVType) bool {
	if _, ok := o.(ServiceItemAccountRight); !ok {
		return false
	}

	other := o.(ServiceItemAccountRight)

	if siar.StructureVersion != other.StructureVersion {
		return false
	}

	if !siar.PID.Equals(other.PID) {
		return false
	}

	return siar.Limitation.Equals(other.Limitation)
}

// CopyRef copies the current value of the ServiceItemAccountRight
// and returns a pointer to the new copy
func (siar ServiceItemAccountRight) CopyRef() types.RVTypePtr {
	copied := siar.Copy().(ServiceItemAccountRight)
	return &copied
}

// Deref takes a pointer to the ServiceItemAccountRight
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (siar *ServiceItemAccountRight) Deref() types.RVType {
	return *siar
}

// String returns the string representation of the ServiceItemAccountRight
func (siar ServiceItemAccountRight) String() string {
	return siar.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemAccountRight using the provided indentation level
func (siar ServiceItemAccountRight) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemAccountRight{\n")
	b.WriteString(fmt.Sprintf("%sPID: %s,\n", indentationValues, siar.PID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sLimitation: %s,\n", indentationValues, siar.Limitation.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemAccountRight returns a new ServiceItemAccountRight
func NewServiceItemAccountRight() ServiceItemAccountRight {
	return ServiceItemAccountRight{
		PID:        types.NewPID(0),
		Limitation: NewServiceItemLimitation(),
	}

}
