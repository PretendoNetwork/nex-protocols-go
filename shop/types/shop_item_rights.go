// Package types implements all the types used by the Shop protocol
package types

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
)

// ShopItemRights is a data structure used by the Shop protocol
type ShopItemRights struct {
	types.Structure
	ReferenceID []byte
	ItemType    *types.PrimitiveS8
	Attribute   *types.PrimitiveU32
}

// ExtractFrom extracts the ShopItemRights from the given readable
func (shopItemRights *ShopItemRights) ExtractFrom(readable types.Readable) error {
	referenceID, err := stream.ReadQBuffer()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItemRights.ReferenceID from stream. %s", err.Error())
	}

	itemType, err := stream.ReadInt8()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItemRights.ItemType from stream. %s", err.Error())
	}

	attribute, err := stream.ReadUInt32LE()
	if err != nil {
		return fmt.Errorf("Failed to extract ShopItemRights.Attribute from stream. %s", err.Error())
	}

	shopItemRights.ReferenceID = referenceID
	shopItemRights.ItemType = itemType
	shopItemRights.Attribute = attribute

	return nil
}

// WriteTo writes the ShopItemRights to the given writable
func (shopItemRights *ShopItemRights) WriteTo(writable types.Writable) {
	contentWritable := writable.CopyNew()

	stream.WriteQBuffer(shopItemRights.ReferenceID)
	*types.PrimitiveU8(shopItemRights.ItemType).WriteTo(contentWritable)
	shopItemRights.Attribute.WriteTo(contentWritable)

	content := contentWritable.Bytes()

	shopItemRights.WriteHeaderTo(writable, uint32(len(content)))

	writable.Write(content)
}

// Copy returns a new copied instance of ShopItemRights
func (shopItemRights *ShopItemRights) Copy() types.RVType {
	copied := NewShopItemRights()

	copied.StructureVersion = shopItemRights.StructureVersion

	copied.ReferenceID = make([]byte, len(shopItemRights.ReferenceID))

	copy(copied.ReferenceID, shopItemRights.ReferenceID)

	copied.ItemType = shopItemRights.ItemType
	copied.Attribute = shopItemRights.Attribute

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (shopItemRights *ShopItemRights) Equals(o types.RVType) bool {
	if _, ok := o.(*ShopItemRights); !ok {
		return false
	}

	other := o.(*ShopItemRights)

	if shopItemRights.StructureVersion != other.StructureVersion {
		return false
	}

	if !shopItemRights.ReferenceID.Equals(other.ReferenceID) {
		return false
	}

	if !shopItemRights.ItemType.Equals(other.ItemType) {
		return false
	}

	if !shopItemRights.Attribute.Equals(other.Attribute) {
		return false
	}

	return true
}

// String returns a string representation of the struct
func (shopItemRights *ShopItemRights) String() string {
	return shopItemRights.FormatToString(0)
}

// FormatToString pretty-prints the struct data using the provided indentation level
func (shopItemRights *ShopItemRights) FormatToString(indentationLevel int) string {
	indentationValues := strings.Repeat("\t", indentationLevel+1)
	indentationEnd := strings.Repeat("\t", indentationLevel)

	var b strings.Builder

	b.WriteString("ShopItem{\n")
	b.WriteString(fmt.Sprintf("%sStructureVersion: %d,\n", indentationValues, shopItemRights.StructureVersion))
	b.WriteString(fmt.Sprintf("%sReferenceID: %x,\n", indentationValues, shopItemRights.ReferenceID))
	b.WriteString(fmt.Sprintf("%sItemType: %d,\n", indentationValues, shopItemRights.ItemType))
	b.WriteString(fmt.Sprintf("%sAttribute: %d\n", indentationValues, shopItemRights.Attribute))
	b.WriteString(fmt.Sprintf("%s}", indentationEnd))

	return b.String()
}

// NewShopItemRights returns a new ShopItemRights
func NewShopItemRights() *ShopItemRights {
	return &ShopItemRights{}
}
