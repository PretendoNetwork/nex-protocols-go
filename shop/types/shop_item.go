// Package types implements all the types used by the Shop protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ShopItem is a data structure used by the Shop protocol
type ShopItem struct {
	types.Structure
	ItemID      *types.PrimitiveU32
	ReferenceID []byte
	ServiceName string
	ItemCode    string
}

// ExtractFrom extracts the ShopItem from the given readable
func (shopItem *ShopItem) ExtractFrom(readable types.Readable) error {
	itemID, err := stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem.ItemID from stream. %s", err.Error())
	}

	referenceID, err := stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem.ReferenceID from stream. %s", err.Error())
	}

	serviceName, err := stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem.ServiceName from stream. %s", err.Error())
	}

	itemCode, err := stream.ReadString()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItem.ItemCode from stream. %s", err.Error())
	}

	shopItem.ItemID = itemID
	shopItem.ReferenceID = referenceID
	shopItem.ServiceName = serviceName
	shopItem.ItemCode = itemCode

	return nil
}

// WriteTo writes the ShopItem to the given writable
func (shopItem *ShopItem) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	shopItem.ItemID.WriteTo(contentWritable)
	stream.WriteQBuffer(shopItem.ReferenceID)
	shopItem.ServiceName.WriteTo(contentWritable)
	shopItem.ItemCode.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	shopItem.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ShopItem
func (shopItem *ShopItem) Copy() types.RVType {
	copied := NewShopItem()

	copied.StructureVersion = shopItem.StructureVersion

	copied.ItemID = shopItem.ItemID
	copied.ReferenceID = make([]byte, len(shopItem.ReferenceID))

	copy(copied.ReferenceID, shopItem.ReferenceID)

	copied.ServiceName = shopItem.ServiceName
	copied.ItemCode = shopItem.ItemCode

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (shopItem *ShopItem) Equals(o types.RVType) bool {
	if _, ok := o.(*ShopItem); !ok {
		return false
	}

	other := o.(*ShopItem)

	if shopItem.StructureVersion != other.StructureVersion {
		return false
	}

	if !shopItem.ItemID.Equals(other.ItemID) {
		return false
	}

	if !shopItem.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !shopItem.ServiceName.Equals(other.ServiceName) {
		return false
	}

	if !shopItem.ItemCode.Equals(other.ItemCode) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (shopItem *ShopItem) String() string {
	return shopItem.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (shopItem *ShopItem) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ShopItem{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, shopItem.StructureVersion))
	b.WriteString(fmt.Sprintf("%sItemID: %d,\n", indentationValues, shopItem.ItemID))
	b.WriteString(fmt.Sprintf("%sReferenceID: %x,\n", indentationValues, shopItem.ReferenceID))
	b.WriteString(fmt.Sprintf("%sServiceName: %q,\n", indentationValues, shopItem.ServiceName))
	b.WriteString(fmt.Sprintf("%sItemCode: %q\n", indentationValues, shopItem.ItemCode))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewShopItem returns a new ShopItem
func NewShopItem() *ShopItem {
	return &ShopItem{}
}
