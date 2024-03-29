// Package types implements all the types used by the Service Item (Wii Sports Club) protocol
package types

import (
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
)

// ServiceItemPrepurchaseInfo holds data for the Service Item (Wii Sports Club) protocol
type ServiceItemPrepurchaseInfo struct {
	nex.Structure
	ItemCode       string
	PriceID        string
	RegularPrice   *ServiceItemAmount
	IsTaxAvailable bool
	TaxAmount      *ServiceItemAmount
	TotalAmount    *ServiceItemAmount
	CurrentBalance *ServiceItemAmount
	PostBalance    *ServiceItemAmount
}

// ExtractFromStream extracts a ServiceItemPrepurchaseInfo structure from a stream
func (serviceItemPrepurchaseInfo *ServiceItemPrepurchaseInfo) ExtractFromStream(stream *nex.StreamIn) error {
	var err error

	serviceItemPrepurchaseInfo.ItemCode, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.ItemCode from stream. %s", err.Error())
	}

	serviceItemPrepurchaseInfo.PriceID, err = stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.PriceID from stream. %s", err.Error())
	}

	regularPrice, err := stream.ReadStructure(NewServiceItemAmount())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.RegularPrice from stream. %s", err.Error())
	}

	serviceItemPrepurchaseInfo.RegularPrice = regularPrice.(*ServiceItemAmount)

	serviceItemPrepurchaseInfo.IsTaxAvailable, err = stream.ReadBool()
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.IsTaxAvailable from stream. %s", err.Error())
	}

	taxAmount, err := stream.ReadStructure(NewServiceItemAmount())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.TaxAmount from stream. %s", err.Error())
	}

	serviceItemPrepurchaseInfo.TaxAmount = taxAmount.(*ServiceItemAmount)

	totalAmount, err := stream.ReadStructure(NewServiceItemAmount())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.TotalAmount from stream. %s", err.Error())
	}

	serviceItemPrepurchaseInfo.TotalAmount = totalAmount.(*ServiceItemAmount)

	currentBalance, err := stream.ReadStructure(NewServiceItemAmount())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.CurrentBalance from stream. %s", err.Error())
	}

	serviceItemPrepurchaseInfo.CurrentBalance = currentBalance.(*ServiceItemAmount)

	postBalance, err := stream.ReadStructure(NewServiceItemAmount())
	if err != nil {
		return fmt.Errorf("Failed to extract ServiceItemPrepurchaseInfo.PostBalance from stream. %s", err.Error())
	}

	serviceItemPrepurchaseInfo.PostBalance = postBalance.(*ServiceItemAmount)

	return nil
}

// Bytes encodes the ServiceItemPrepurchaseInfo and returns a byte array
func (serviceItemPrepurchaseInfo *ServiceItemPrepurchaseInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(serviceItemPrepurchaseInfo.ItemCode)
	stream.WriteString(serviceItemPrepurchaseInfo.PriceID)
	stream.WriteStructure(serviceItemPrepurchaseInfo.RegularPrice)
	stream.WriteBool(serviceItemPrepurchaseInfo.IsTaxAvailable)
	stream.WriteStructure(serviceItemPrepurchaseInfo.TaxAmount)
	stream.WriteStructure(serviceItemPrepurchaseInfo.TotalAmount)
	stream.WriteStructure(serviceItemPrepurchaseInfo.CurrentBalance)
	stream.WriteStructure(serviceItemPrepurchaseInfo.PostBalance)

	return stream.Bytes()
}

// Copy returns a new copied instance of ServiceItemPrepurchaseInfo
func (serviceItemPrepurchaseInfo *ServiceItemPrepurchaseInfo) Copy() nex.StructureInterface {
	copied := NewServiceItemPrepurchaseInfo()

	copied.SetStructureVersion(serviceItemPrepurchaseInfo.StructureVersion())

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
func (serviceItemPrepurchaseInfo *ServiceItemPrepurchaseInfo) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ServiceItemPrepurchaseInfo)

	if serviceItemPrepurchaseInfo.StructureVersion() != other.StructureVersion() {
		return false
	}

	if serviceItemPrepurchaseInfo.ItemCode != other.ItemCode {
		return false
	}

	if serviceItemPrepurchaseInfo.PriceID != other.PriceID {
		return false
	}

	if !serviceItemPrepurchaseInfo.RegularPrice.Equals(other.RegularPrice) {
		return false
	}

	if serviceItemPrepurchaseInfo.IsTaxAvailable != other.IsTaxAvailable {
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
	b.WriteString(fmt.Sprintf("%sstructureVersion: %d,\n", indentationValues, serviceItemPrepurchaseInfo.StructureVersion()))
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
