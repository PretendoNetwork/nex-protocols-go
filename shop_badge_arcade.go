package nexproto

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// ShopBadgeArcadeProtocolID is the Protocol ID for the Shop (Badge Arcade) protocol
	ShopBadgeArcadeProtocolID = 0x7F

	// ShopBadgeArcadeCustomID is the Custom ID for the Shop (Badge Arcade) protocol
	ShopBadgeArcadeCustomID = 0xC8

	// ShopBadgeArcadeMethodGetRivToken is the method ID for GetRivToken
	ShopBadgeArcadeMethodGetRivToken = 0x1

	// ShopBadgeArcadeMethodPostPlayLog is the method ID for PostPlayLog
	ShopBadgeArcadeMethodPostPlayLog = 0x2
)

// ShopBadgeArcadeProtocol handles the Shop (Badge Arcade) nex protocol
type ShopBadgeArcadeProtocol struct {
	server *nex.Server
	GetRivTokenHandler func(err error, client *nex.Client, callID uint32, itemCode string, referenceID []byte)
	PostPlayLogHandler func(err error, client *nex.Client, callID uint32, param *ShopPostPlayLogParam)
}

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
	stream.WriteUInt64LE(shopPostPlayLogParam.Timestamp.Value())
	stream.WriteString(shopPostPlayLogParam.Unknown2)

	return stream.Bytes()
}

// NewShopPostPlayLogParam returns a new ShopPostPlayLogParam
func NewShopPostPlayLogParam() *ShopPostPlayLogParam {
	return &ShopPostPlayLogParam{}
}

// Setup initializes the protocol
func (shopBadgeArcadeProtocol *ShopBadgeArcadeProtocol) Setup() {
	nexServer := shopBadgeArcadeProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if ShopBadgeArcadeProtocolID == request.ProtocolID() && ShopBadgeArcadeCustomID == request.CustomID() {
			switch request.MethodID() {
			case ShopBadgeArcadeMethodPostPlayLog:
				go shopBadgeArcadeProtocol.handlePostPlayLog(packet)
			default:
				go respondNotImplementedCustom(packet, ShopBadgeArcadeCustomID)
				fmt.Printf("Unsupported ShopBadgeArcade method ID: %#v\n", request.MethodID())
			}
		}
	})
}

// GetRivToken sets the GetRivToken function
func (shopBadgeArcadeProtocol *ShopBadgeArcadeProtocol) GetRivToken(handler func(err error, client *nex.Client, callID uint32, itemCode string, referenceID []byte)) {
	shopBadgeArcadeProtocol.GetRivTokenHandler = handler
}

// PostPlayLog sets the PostPlayLog function
func (shopBadgeArcadeProtocol *ShopBadgeArcadeProtocol) PostPlayLog(handler func(err error, client *nex.Client, callID uint32, param *ShopPostPlayLogParam)) {
	shopBadgeArcadeProtocol.PostPlayLogHandler = handler
}

func (shopBadgeArcadeProtocol *ShopBadgeArcadeProtocol) handleGetRivToken(packet nex.PacketInterface) {
	if shopBadgeArcadeProtocol.GetRivTokenHandler == nil {
		logger.Warning("ShopBadgeArcadeProtocol::GetRivToken not implemented")
		go respondNotImplementedCustom(packet, ShopBadgeArcadeCustomID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, shopBadgeArcadeProtocol.server)

	itemCode, err := parametersStream.ReadString()
	if err != nil {
		go shopBadgeArcadeProtocol.GetRivTokenHandler(err, client, callID, "", []byte{})
		return
	}

	referenceID, err := parametersStream.ReadQBuffer()
	if err != nil {
		go shopBadgeArcadeProtocol.GetRivTokenHandler(err, client, callID, "", []byte{})
		return
	}

	go shopBadgeArcadeProtocol.GetRivTokenHandler(nil, client, callID, itemCode, referenceID)
}

func (shopBadgeArcadeProtocol *ShopBadgeArcadeProtocol) handlePostPlayLog(packet nex.PacketInterface) {
	if shopBadgeArcadeProtocol.PostPlayLogHandler == nil {
		logger.Warning("ShopBadgeArcadeProtocol::PostPlayLog not implemented")
		go respondNotImplementedCustom(packet, ShopBadgeArcadeCustomID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, shopBadgeArcadeProtocol.server)

	param, err := parametersStream.ReadStructure(NewShopPostPlayLogParam())
	if err != nil {
		go shopBadgeArcadeProtocol.PostPlayLogHandler(err, client, callID, nil)
		return
	}

	go shopBadgeArcadeProtocol.PostPlayLogHandler(nil, client, callID, param.(*ShopPostPlayLogParam))
}

// NewShopBadgeArcadeProtocol returns a new ShopBadgeArcadeProtocol
func NewShopBadgeArcadeProtocol(server *nex.Server) *ShopBadgeArcadeProtocol {
	shopBadgeArcadeProtocol := &ShopBadgeArcadeProtocol{server: server}

	shopBadgeArcadeProtocol.Setup()

	return shopBadgeArcadeProtocol
}
