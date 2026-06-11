// Package types implements all the types used by the ServiceItem protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2/types"
)

// ServiceItemPurchaseServiceItemParam is a type within the ServiceItem protocol
type ServiceItemPurchaseServiceItemParam struct {
	types.Structure
	ItemCode       types.String
	PriceID        types.String
	ReferenceID    types.String
	Balance        types.String
	ItemName       types.String
	EcServiceToken types.String
	Language       types.String
	TitleID        types.String
}

// WriteTo writes the ServiceItemPurchaseServiceItemParam to the given writable
func (sipsip ServiceItemPurchaseServiceItemParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	sipsip.ItemCode.WriteTo(contentWritable)
	sipsip.PriceID.WriteTo(contentWritable)
	sipsip.ReferenceID.WriteTo(contentWritable)
	sipsip.Balance.WriteTo(contentWritable)
	sipsip.ItemName.WriteTo(contentWritable)
	sipsip.EcServiceToken.WriteTo(contentWritable)
	sipsip.Language.WriteTo(contentWritable)
	sipsip.TitleID.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	sipsip.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// ExtractFrom extracts the ServiceItemPurchaseServiceItemParam from the given readable
func (sipsip *ServiceItemPurchaseServiceItemParam) ExtractFrom(readable types.Readable) error {
	if err := sipsip.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseServiceItemParam header. %s", err.Error())
	}

	if err := sipsip.ItemCode.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseServiceItemParam.ItemCode. %s", err.Error())
	}

	if err := sipsip.PriceID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseServiceItemParam.PriceID. %s", err.Error())
	}

	if err := sipsip.ReferenceID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseServiceItemParam.ReferenceID. %s", err.Error())
	}

	if err := sipsip.Balance.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseServiceItemParam.Balance. %s", err.Error())
	}

	if err := sipsip.ItemName.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseServiceItemParam.ItemName. %s", err.Error())
	}

	if err := sipsip.EcServiceToken.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseServiceItemParam.EcServiceToken. %s", err.Error())
	}

	if err := sipsip.Language.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseServiceItemParam.Language. %s", err.Error())
	}

	if err := sipsip.TitleID.ExtractFrom(readable); err != nil {
		return fmt.Errorf("failed to extract ServiceItemPurchaseServiceItemParam.TitleID. %s", err.Error())
	}

	return nil
}

// Copy returns a new copied instance of ServiceItemPurchaseServiceItemParam
func (sipsip ServiceItemPurchaseServiceItemParam) Copy() types.RVType {
	copied := NewServiceItemPurchaseServiceItemParam()

	copied.StructureVersion = sipsip.StructureVersion
	copied.ItemCode = sipsip.ItemCode.Copy().(types.String)
	copied.PriceID = sipsip.PriceID.Copy().(types.String)
	copied.ReferenceID = sipsip.ReferenceID.Copy().(types.String)
	copied.Balance = sipsip.Balance.Copy().(types.String)
	copied.ItemName = sipsip.ItemName.Copy().(types.String)
	copied.EcServiceToken = sipsip.EcServiceToken.Copy().(types.String)
	copied.Language = sipsip.Language.Copy().(types.String)
	copied.TitleID = sipsip.TitleID.Copy().(types.String)

	return copied
}

// Equals checks if the given ServiceItemPurchaseServiceItemParam contains the same data as the current ServiceItemPurchaseServiceItemParam
func (sipsip ServiceItemPurchaseServiceItemParam) Equals(o types.RVType) bool {
	if _, ok := o.(ServiceItemPurchaseServiceItemParam); !ok {
		return false
	}

	other := o.(ServiceItemPurchaseServiceItemParam)

	if sipsip.StructureVersion != other.StructureVersion {
		return false
	}

	if !sipsip.ItemCode.Equals(other.ItemCode) {
		return false
	}

	if !sipsip.PriceID.Equals(other.PriceID) {
		return false
	}

	if !sipsip.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !sipsip.Balance.Equals(other.Balance) {
		return false
	}

	if !sipsip.ItemName.Equals(other.ItemName) {
		return false
	}

	if !sipsip.EcServiceToken.Equals(other.EcServiceToken) {
		return false
	}

	if !sipsip.Language.Equals(other.Language) {
		return false
	}

	return sipsip.TitleID.Equals(other.TitleID)
}

// CopyRef copies the current value of the ServiceItemPurchaseServiceItemParam
// and returns a pointer to the new copy
func (sipsip ServiceItemPurchaseServiceItemParam) CopyRef() types.RVTypePtr {
	copied := sipsip.Copy().(ServiceItemPurchaseServiceItemParam)
	return &copied
}

// Deref takes a pointer to the ServiceItemPurchaseServiceItemParam
// and dereferences it to the raw value.
// Only useful when working with an instance of RVTypePtr
func (sipsip *ServiceItemPurchaseServiceItemParam) Deref() types.RVType {
	return *sipsip
}

// String returns the string representation of the ServiceItemPurchaseServiceItemParam
func (sipsip ServiceItemPurchaseServiceItemParam) String() string {
	return sipsip.FormatToString(0)
}

// FormatToString pretty-prints the ServiceItemPurchaseServiceItemParam using the provided indentation level
func (sipsip ServiceItemPurchaseServiceItemParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaseServiceItemParam{\n")
	fmt.Fprintf(&b, "%sItemCode: %s,\n", indentationValues, sipsip.ItemCode)
	fmt.Fprintf(&b, "%sPriceID: %s,\n", indentationValues, sipsip.PriceID)
	fmt.Fprintf(&b, "%sReferenceID: %s,\n", indentationValues, sipsip.ReferenceID)
	fmt.Fprintf(&b, "%sBalance: %s,\n", indentationValues, sipsip.Balance)
	fmt.Fprintf(&b, "%sItemName: %s,\n", indentationValues, sipsip.ItemName)
	fmt.Fprintf(&b, "%sEcServiceToken: %s,\n", indentationValues, sipsip.EcServiceToken)
	fmt.Fprintf(&b, "%sLanguage: %s,\n", indentationValues, sipsip.Language)
	fmt.Fprintf(&b, "%sTitleID: %s,\n", indentationValues, sipsip.TitleID)
	fmt.Fprintf(&b, "%s}", indentationEnd)

	return b.String()
}

// NewServiceItemPurchaseServiceItemParam returns a new ServiceItemPurchaseServiceItemParam
func NewServiceItemPurchaseServiceItemParam() ServiceItemPurchaseServiceItemParam {
	return ServiceItemPurchaseServiceItemParam{
		ItemCode:       types.NewString(""),
		PriceID:        types.NewString(""),
		ReferenceID:    types.NewString(""),
		Balance:        types.NewString(""),
		ItemName:       types.NewString(""),
		EcServiceToken: types.NewString(""),
		Language:       types.NewString(""),
		TitleID:        types.NewString(""),
	}

}
