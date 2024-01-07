// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ServiceItemPrepurchaseInfo holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemPrepurchaseInfo struct {
	types.Structure
	ItemCode       string
	PriceID        string
	RegularPrice   *ServiceItemAmount
	IsTaxAvailable *types.PrimitiveBool
	TaxAmount      *ServiceItemAmount
	TotalAmount    *ServiceItemAmount
	CurrentBalance *ServiceItemAmount
	PostBalance    *ServiceItemAmount
}

// ExtractFrom extracts the ServiceItemPrepurchaseInfo from the given readable
func (serviceItemPrepurchaseInfo *ServiceItemPrepurchaseInfo) ExtractFrom(readable types.Readable) error {
	var err error

	if err = serviceItemPrepurchaseInfo.ExtractHeaderFrom(readable); err != nil {
		return fmt.Errorf("Failed to read ServiceItemPrepurchaseInfo header. %s", err.Error())
	}

	err = serviceItemPrepurchaseInfo.ItemCode.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.ItemCode from stream. %s", err.Error())
	}

	err = serviceItemPrepurchaseInfo.PriceID.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.PriceID from stream. %s", err.Error())
	}

	err = serviceItemPrepurchaseInfo.RegularPrice.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.RegularPrice from stream. %s", err.Error())
	}

	err = serviceItemPrepurchaseInfo.IsTaxAvailable.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.IsTaxAvailable from stream. %s", err.Error())
	}

	err = serviceItemPrepurchaseInfo.TaxAmount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.TaxAmount from stream. %s", err.Error())
	}

	err = serviceItemPrepurchaseInfo.TotalAmount.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.TotalAmount from stream. %s", err.Error())
	}

	err = serviceItemPrepurchaseInfo.CurrentBalance.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.CurrentBalance from stream. %s", err.Error())
	}

	err = serviceItemPrepurchaseInfo.PostBalance.ExtractFrom(readable)
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.PostBalance from stream. %s", err.Error())
	}

	return nil
}

// WriteTo writes the ServiceItemPrepurchaseInfo to the given writable
func (serviceItemPrepurchaseInfo *ServiceItemPrepurchaseInfo) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	serviceItemPrepurchaseInfo.ItemCode.WriteTo(contentWritable)
	serviceItemPrepurchaseInfo.PriceID.WriteTo(contentWritable)
	serviceItemPrepurchaseInfo.RegularPrice.WriteTo(contentWritable)
	serviceItemPrepurchaseInfo.IsTaxAvailable.WriteTo(contentWritable)
	serviceItemPrepurchaseInfo.TaxAmount.WriteTo(contentWritable)
	serviceItemPrepurchaseInfo.TotalAmount.WriteTo(contentWritable)
	serviceItemPrepurchaseInfo.CurrentBalance.WriteTo(contentWritable)
	serviceItemPrepurchaseInfo.PostBalance.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	serviceItemPrepurchaseInfo.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ServiceItemPrepurchaseInfo
func (serviceItemPrepurchaseInfo *ServiceItemPrepurchaseInfo) Copy() types.RVType {
	copied := NewServiceItemPrepurchaseInfo()

	copied.StructureVersion = serviceItemPrepurchaseInfo.StructureVersion

	copied.ItemCode = serviceItemPrepurchaseInfo.ItemCode
	copied.PriceID = serviceItemPrepurchaseInfo.PriceID
	copied.RegularPrice = serviceItemPrepurchaseInfo.RegularPrice.Copy().(*ServiceItemAmount)
	copied.IsTaxAvailable = serviceItemPrepurchaseInfo.IsTaxAvailable
	copied.TaxAmount = serviceItemPrepurchaseInfo.TaxAmount.Copy().(*ServiceItemAmount)
	copied.TotalAmount = serviceItemPrepurchaseInfo.TotalAmount.Copy().(*ServiceItemAmount)
	copied.CurrentBalance = serviceItemPrepurchaseInfo.CurrentBalance.Copy().(*ServiceItemAmount)
	copied.PostBalance = serviceItemPrepurchaseInfo.PostBalance.Copy().(*ServiceItemAmount)

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (serviceItemPrepurchaseInfo *ServiceItemPrepurchaseInfo) Equals(o types.RVType) bool {
	if _, ok := o.(*ServiceItemPrepurchaseInfo); !ok {
		return false
	}

	other := o.(*ServiceItemPrepurchaseInfo)

	if serviceItemPrepurchaseInfo.StructureVersion != other.StructureVersion {
		return false
	}

	if !serviceItemPrepurchaseInfo.ItemCode.Equals(other.ItemCode) {
		return false
	}

	if !serviceItemPrepurchaseInfo.PriceID.Equals(other.PriceID) {
		return false
	}

	if !serviceItemPrepurchaseInfo.RegularPrice.Equals(other.RegularPrice) {
		return false
	}

	if !serviceItemPrepurchaseInfo.IsTaxAvailable.Equals(other.IsTaxAvailable) {
		return false
	}

	if !serviceItemPrepurchaseInfo.TaxAmount.Equals(other.TaxAmount) {
		return false
	}

	if !serviceItemPrepurchaseInfo.TotalAmount.Equals(other.TotalAmount) {
		return false
	}

	if !serviceItemPrepurchaseInfo.CurrentBalance.Equals(other.CurrentBalance) {
		return false
	}

	if !serviceItemPrepurchaseInfo.PostBalance.Equals(other.PostBalance) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (serviceItemPrepurchaseInfo *ServiceItemPrepurchaseInfo) String() string {
	return serviceItemPrepurchaseInfo.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (serviceItemPrepurchaseInfo *ServiceItemPrepurchaseInfo) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ServiceItemPrepurchaseInfo{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, serviceItemPrepurchaseInfo.StructureVersion))
	b.WriteString(fmt.Sprintf("%sItemCode: %q,\n", indentationValues, serviceItemPrepurchaseInfo.ItemCode))
	b.WriteString(fmt.Sprintf("%sPriceID: %q,\n", indentationValues, serviceItemPrepurchaseInfo.PriceID))

	if serviceItemPrepurchaseInfo.RegularPrice != nil {
		b.WriteString(fmt.Sprintf("%sRegularPrice: %s\n", indentationValues, serviceItemPrepurchaseInfo.RegularPrice.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sRegularPrice: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%sIsTaxAvailable: %t,\n", indentationValues, serviceItemPrepurchaseInfo.IsTaxAvailable))

	if serviceItemPrepurchaseInfo.TaxAmount != nil {
		b.WriteString(fmt.Sprintf("%sTaxAmount: %s\n", indentationValues, serviceItemPrepurchaseInfo.TaxAmount.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sTaxAmount: nil\n", indentationValues))
	}

	if serviceItemPrepurchaseInfo.TotalAmount != nil {
		b.WriteString(fmt.Sprintf("%sTotalAmount: %s\n", indentationValues, serviceItemPrepurchaseInfo.TotalAmount.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sTotalAmount: nil\n", indentationValues))
	}

	if serviceItemPrepurchaseInfo.CurrentBalance != nil {
		b.WriteString(fmt.Sprintf("%sCurrentBalance: %s\n", indentationValues, serviceItemPrepurchaseInfo.CurrentBalance.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sCurrentBalance: nil\n", indentationValues))
	}

	if serviceItemPrepurchaseInfo.PostBalance != nil {
		b.WriteString(fmt.Sprintf("%sPostBalance: %s\n", indentationValues, serviceItemPrepurchaseInfo.PostBalance.FormatToString(indentationLevel+1)))
	} else {
		b.WriteString(fmt.Sprintf("%sPostBalance: nil\n", indentationValues))
	}

	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewServiceItemPrepurchaseInfo returns a new ServiceItemPrepurchaseInfo
func NewServiceItemPrepurchaseInfo() *ServiceItemPrepurchaseInfo {
	return &ServiceItemPrepurchaseInfo{}
}
