// Package types implements all the types used by the Subscription protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// SubscriptionData is a type within the Subscription protocol
type SubscriptionData struct {
	types.Structure
	types.Data
	PrincipalID types.PID
	Unknown     types.QBuffer
}

// ObjectID returns the object identifier of the type
func (sd SubscriptionData) ObjectID() types.RVType {
	return sd.DataObjectID()
}

// DataObjectID returns the object identifier of the type embedding Data
func (sd SubscriptionData) DataObjectID() types.RVType {
	return types.NewString("SubscriptionData")
}

// WriteTo writes the SubscriptionData to the given writable
func (sd SubscriptionData) WriteTo(writable types.Writable) {
	sd.Data.WriteTo(writable)

	contentWritable := writable.CopyNew()

	sd.PrincipalID.WriteTo(contentWritable)
	sd.Unknown.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the SubscriptionData from the given readable
func (sd *SubscriptionData) ExtractFrom(readable types.Readable) error {
	var err error

	err = sd.Data.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriptionData.Data. %s", err.Error())
	}

	err = sd.ExtractHeaderFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriptionData header. %s", err.Error())
	}

	err = sd.PrincipalID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriptionData.PrincipalID. %s", err.Error())
	}

	err = sd.Unknown.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract SubscriptionData.Unknown. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of SubscriptionData
func (sd SubscriptionData) Copy() types.RVType {
	copied := NewSubscriptionData()

	copied.StructureVersion = sd.StructureVersion
	copied.Data = sd.Data.Copy().(types.Data)
	copied.PrincipalID = sd.PrincipalID.Copy().(types.PID)
	copied.Unknown = sd.Unknown.Copy().(types.QBuffer)

	return copied
}

// Equals checks if the given SubscriptionData contains the same data as the current SubscriptionData
func (sd SubscriptionData) Equals(o types.RVType) bool {
	if _, ok := o.(SubscriptionData); !ok {
		return false
	}

	other := o.(SubscriptionData)

	if sd.StructureVersion != other.StructureVersion {
		return false
	}

	if !sd.Data.Equals(other.Data) {
		return false
	}

	if !sd.PrincipalID.Equals(other.PrincipalID) {
		return false
	}

	return sd.Unknown.Equals(other.Unknown)
}

// CopyRef copies the current value of the SubscriptionData
// and returns a pointer to the new copy
func (sd SubscriptionData) CopyRef() types.RVTypePtr {
	copied := sd.Copy().(SubscriptionData)
	return &copied
}

// Deref takes a pointer to the SubscriptionData
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sd *SubscriptionData) Deref() types.RVType {
	return *sd
}

// String returns the string representation of the SubscriptionData
func (sd SubscriptionData) String() string {
	return sd.FormatToString(0)
}

// FormatToString pretty-prints the SubscriptionData using the provided indentation level
func (sd SubscriptionData) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("SubscriptionData{\n")
	b.WriteString(fmt.Sprintf("%sData (parent): %s,\n", indentationValues, sd.Data.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sPrincipalID: %s,\n", indentationValues, sd.PrincipalID.FormatToString(indentationLevel+1)))
	b.WriteString(fmt.Sprintf("%sUnknown: %s,\n", indentationValues, sd.Unknown))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewSubscriptionData returns a new SubscriptionData
func NewSubscriptionData() SubscriptionData {
	return SubscriptionData{
		PrincipalID: types.NewPID(0),
		Unknown:     types.NewQBuffer(nil),
	}

}
