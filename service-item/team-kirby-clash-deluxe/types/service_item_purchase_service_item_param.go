// Package types implements all the types used by the Service Item (Team Kirby Clash Deluxe) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemPurchaseServiceItemParam holds data for the Service Item (Team Kirby Clash Deluxe) protocol
type ServiceItemPurchaseServiceItemParam struct {
	types.Structure
	ItemCode       string
	PriceID        string
	ReferenceID    string
	Balance        string
	ItemName       string
	EcServiceToken string
	Language       string
	UniqueID       *types.PrimitiveU32
	Platform       *types.PrimitiveU8 // * Revision 1
}

// ExtractFrom extracts the ServiceItemPurchaseServiceItemParam from the given readable
func (serviceItemPurchaseServiceItemParam *ServiceItemPurchaseServiceItemParam) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemPurchaseServiceItemParam.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemPurchaseServiceItemParam header. %s", err.Error())
	}

	err = serviceItemPurchaseServiceItemParam.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.ItemCode from stream. %s", err.Error())
	}

	err = serviceItemPurchaseServiceItemParam.PriceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.PriceID from stream. %s", err.Error())
	}

	err = serviceItemPurchaseServiceItemParam.ReferenceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.ReferenceID from stream. %s", err.Error())
	}

	err = serviceItemPurchaseServiceItemParam.Balance.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.Balance from stream. %s", err.Error())
	}

	err = serviceItemPurchaseServiceItemParam.ItemName.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.ItemName from stream. %s", err.Error())
	}

	err = serviceItemPurchaseServiceItemParam.EcServiceToken.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.EcServiceToken from stream. %s", err.Error())
	}

	err = serviceItemPurchaseServiceItemParam.Language.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.Language from stream. %s", err.Error())
	}

	err = serviceItemPurchaseServiceItemParam.UniqueID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.UniqueID from stream. %s", err.Error())
	}

	if serviceItemPurchaseServiceItemParam.StructureVersion >= 1 {
	err = 	serviceItemPurchaseServiceItemParam.Platform.ExtractFrom(readable)
		if err != nil {
			return fmt.Errorf("Failed to extract ServiceItemPurchaseServiceItemParam.Platform from stream. %s", err.Error())
		}
	}

	return nil
}

// WriteTo writes the ServiceItemPurchaseServiceItemParam to the given writable
func (serviceItemPurchaseServiceItemParam *ServiceItemPurchaseServiceItemParam) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemPurchaseServiceItemParam.ItemCode.WriteTo(contentWritable)
	serviceItemPurchaseServiceItemParam.PriceID.WriteTo(contentWritable)
	serviceItemPurchaseServiceItemParam.ReferenceID.WriteTo(contentWritable)
	serviceItemPurchaseServiceItemParam.Balance.WriteTo(contentWritable)
	serviceItemPurchaseServiceItemParam.ItemName.WriteTo(contentWritable)
	serviceItemPurchaseServiceItemParam.EcServiceToken.WriteTo(contentWritable)
	serviceItemPurchaseServiceItemParam.Language.WriteTo(contentWritable)
	serviceItemPurchaseServiceItemParam.UniqueID.WriteTo(contentWritable)

	if serviceItemPurchaseServiceItemParam.StructureVersion >= 1 {
		serviceItemPurchaseServiceItemParam.Platform.WriteTo(contentWritable)
	}

	content := contentWritable.Bytes()

	rvcd.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemPurchaseServiceItemParam
func (serviceItemPurchaseServiceItemParam *ServiceItemPurchaseServiceItemParam) Copy() types.RVType {
	copied := NewServiceItemPurchaseServiceItemParam()

	copied.StructureVersion = serviceItemPurchaseServiceItemParam.StructureVersion

	copied.ItemCode = serviceItemPurchaseServiceItemParam.ItemCode
	copied.PriceID = serviceItemPurchaseServiceItemParam.PriceID
	copied.ReferenceID = serviceItemPurchaseServiceItemParam.ReferenceID
	copied.Balance = serviceItemPurchaseServiceItemParam.Balance
	copied.ItemName = serviceItemPurchaseServiceItemParam.ItemName
	copied.EcServiceToken = serviceItemPurchaseServiceItemParam.EcServiceToken
	copied.Language = serviceItemPurchaseServiceItemParam.Language
	copied.UniqueID = serviceItemPurchaseServiceItemParam.UniqueID
	copied.Platform = serviceItemPurchaseServiceItemParam.Platform

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPurchaseServiceItemParam *ServiceItemPurchaseServiceItemParam) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPurchaseServiceItemParam); !ok {
		return false
	}

	other := o.(*ServiceItemPurchaseServiceItemParam)

	if serviceItemPurchaseServiceItemParam.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemPurchaseServiceItemParam.ItemCode.Equals(other.ItemCode) {
		return false
	}

	if !serviceItemPurchaseServiceItemParam.PriceID.Equals(other.PriceID) {
		return false
	}

	if !serviceItemPurchaseServiceItemParam.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !serviceItemPurchaseServiceItemParam.Balance.Equals(other.Balance) {
		return false
	}

	if !serviceItemPurchaseServiceItemParam.ItemName.Equals(other.ItemName) {
		return false
	}

	if !serviceItemPurchaseServiceItemParam.EcServiceToken.Equals(other.EcServiceToken) {
		return false
	}

	if !serviceItemPurchaseServiceItemParam.Language.Equals(other.Language) {
		return false
	}

	if !serviceItemPurchaseServiceItemParam.UniqueID.Equals(other.UniqueID) {
		return false
	}

	if !serviceItemPurchaseServiceItemParam.Platform.Equals(other.Platform) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemPurchaseServiceItemParam *ServiceItemPurchaseServiceItemParam) String() string {
	return serviceItemPurchaseServiceItemParam.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemPurchaseServiceItemParam *ServiceItemPurchaseServiceItemParam) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPurchaseServiceItemParam{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemPurchaseServiceItemParam.StructureVersion))
	b.WriteString(fmt.Sprintf("%sItemCode: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.ItemCode))
	b.WriteString(fmt.Sprintf("%sPriceID: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.PriceID))
	b.WriteString(fmt.Sprintf("%sReferenceID: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.ReferenceID))
	b.WriteString(fmt.Sprintf("%sBalance: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.Balance))
	b.WriteString(fmt.Sprintf("%sItemName: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.ItemName))
	b.WriteString(fmt.Sprintf("%sEcServiceToken: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.EcServiceToken))
	b.WriteString(fmt.Sprintf("%sLanguage: %q,\n", indentationValues, serviceItemPurchaseServiceItemParam.Language))
	b.WriteString(fmt.Sprintf("%sUniqueID: %d,\n", indentationValues, serviceItemPurchaseServiceItemParam.UniqueID))

	if serviceItemPurchaseServiceItemParam.StructureVersion >= 1 {
		b.WriteString(fmt.Sprintf("%sPlatform: %d,\n", indentationValues, serviceItemPurchaseServiceItemParam.Platform))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPurchaseServiceItemParam returns a new ServiceItemPurchaseServiceItemParam
func NewServiceItemPurchaseServiceItemParam() *ServiceItemPurchaseServiceItemParam {
	return &ServiceItemPurchaseServiceItemParam{}
}
