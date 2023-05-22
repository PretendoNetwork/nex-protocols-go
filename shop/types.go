package shop

import (
	"bytes"

	"github.com/PretendoNetwork/nex-go"
)

type ShopItem struct {
	nex.Structure
	ItemID      uint32
	ReferenceID []byte
	ServiceName string
	ItemCode    string
}

// ExtractFromStream extracts a ShopItem structure from a stream
func (shopItem *ShopItem) ExtractFromStream(stream *nex.StreamIn) error {
	shopItem.ItemID = stream.ReadUInt32LE()

	referenceID, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	shopItem.ReferenceID = referenceID

	serviceName, err := stream.ReadString()
	if err != nil {
		return err
	}

	shopItem.ServiceName = serviceName

	itemCode, err := stream.ReadString()
	if err != nil {
		return err
	}

	shopItem.ItemCode = itemCode

	return nil
}

// Bytes encodes the ShopItem and returns a byte array
func (shopItem *ShopItem) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(shopItem.ItemID)
	stream.WriteQBuffer(shopItem.ReferenceID)
	stream.WriteString(shopItem.ServiceName)
	stream.WriteString(shopItem.ItemCode)

	return stream.Bytes()
}

// Copy returns a new copied instance of ShopItem
func (shopItem *ShopItem) Copy() nex.StructureInterface {
	copied := NewShopItem()

	copied.ItemID = shopItem.ItemID
	copied.ReferenceID = make([]byte, len(shopItem.ReferenceID))

	copy(copied.ReferenceID, shopItem.ReferenceID)

	copied.ServiceName = shopItem.ServiceName
	copied.ItemCode = shopItem.ItemCode

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (shopItem *ShopItem) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ShopItem)

	if shopItem.ItemID != other.ItemID {
		return false
	}

	if !bytes.Equal(shopItem.ReferenceID, other.ReferenceID) {
		return false
	}

	if shopItem.ServiceName != other.ServiceName {
		return false
	}

	if shopItem.ItemCode != other.ItemCode {
		return false
	}

	return true
}

// NewShopItem returns a new ShopItem
func NewShopItem() *ShopItem {
	return &ShopItem{}
}

type ShopItemRights struct {
	nex.Structure
	ReferenceID []byte
	ItemType    int8
	Attribute   uint32
}

// ExtractFromStream extracts a ShopItemRights structure from a stream
func (shopItemRights *ShopItemRights) ExtractFromStream(stream *nex.StreamIn) error {
	referenceID, err := stream.ReadQBuffer()
	if err != nil {
		return err
	}

	shopItemRights.ReferenceID = referenceID
	shopItemRights.ItemType = int8(stream.ReadUInt8())
	shopItemRights.Attribute = stream.ReadUInt32LE()

	return nil
}

// Bytes encodes the ShopItemRights and returns a byte array
func (shopItemRights *ShopItemRights) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteQBuffer(shopItemRights.ReferenceID)
	stream.WriteUInt8(uint8(shopItemRights.ItemType))
	stream.WriteUInt32LE(shopItemRights.Attribute)

	return stream.Bytes()
}

// Copy returns a new copied instance of ShopItemRights
func (shopItemRights *ShopItemRights) Copy() nex.StructureInterface {
	copied := NewShopItemRights()

	copied.ReferenceID = make([]byte, len(shopItemRights.ReferenceID))

	copy(copied.ReferenceID, shopItemRights.ReferenceID)

	copied.ItemType = shopItemRights.ItemType
	copied.Attribute = shopItemRights.Attribute

	return copied
}

// Equals checks if the passed Structure contains the same data as the current instance
func (shopItemRights *ShopItemRights) Equals(structure nex.StructureInterface) bool {
	other := structure.(*ShopItemRights)

	if !bytes.Equal(shopItemRights.ReferenceID, other.ReferenceID) {
		return false
	}

	if shopItemRights.ItemType != other.ItemType {
		return false
	}

	if shopItemRights.Attribute != other.Attribute {
		return false
	}

	return true
}

// NewShopItemRights returns a new ShopItemRights
func NewShopItemRights() *ShopItemRights {
	return &ShopItemRights{}
}
