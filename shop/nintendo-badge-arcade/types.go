package shop_nintendo_badge_arcade

import nex "github.com/PretendoNetwork/nex-go"

type ShopPostPlayLogParam struct {
	nex.Structure
	Unknown1  []uint32
	Timestamp *nex.DateTime
	Unknown2  string
}

// ExtractFromStream extracts a ShopPostPlayLogParam structure from a stream
func (shopPostPlayLogParam *ShopPostPlayLogParam) ExtractFromStream(stream *nex.StreamIn) error {
	shopPostPlayLogParam.Unknown1 = stream.ReadListUInt32LE()
	shopPostPlayLogParam.Timestamp = stream.ReadDateTime()

	unknown2, err := stream.ReadString()
	if err != nil {
		return err
	}

	shopPostPlayLogParam.Unknown2 = unknown2

	return nil
}

// Bytes encodes the ShopPostPlayLogParam and returns a byte array
func (shopPostPlayLogParam *ShopPostPlayLogParam) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteListUInt32LE(shopPostPlayLogParam.Unknown1)
	stream.WriteDateTime(shopPostPlayLogParam.Timestamp)
	stream.WriteString(shopPostPlayLogParam.Unknown2)

	return stream.Bytes()
}

// NewShopPostPlayLogParam returns a new ShopPostPlayLogParam
func NewShopPostPlayLogParam() *ShopPostPlayLogParam {
	return &ShopPostPlayLogParam{}
}
