// Package shop_nintendo_badge_arcade implements the Nintendo Badge Arcade Shop NEX protocol
package shop_nintendo_badge_arcade

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	"github.com/PretendoNetwork/nex-protocols-go/shop"
	shop_nintendo_badge_arcade_types "github.com/PretendoNetwork/nex-protocols-go/shop/nintendo-badge-arcade/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Shop (Badge Arcade) protocol
	ProtocolID = 0x7F

	// CustomProtocolID is the Custom ID for the Shop (Badge Arcade) protocol
	CustomProtocolID = 0xC8

	// MethodGetRivToken is the method ID for GetRivToken
	MethodGetRivToken = 0x1

	// MethodPostPlayLog is the method ID for PostPlayLog
	MethodPostPlayLog = 0x2
)

var patchedMethods = []uint32{
	MethodGetRivToken,
	MethodPostPlayLog,
}

// ShopNintendoBadgeArcadeProtocol handles the Shop (Badge Arcade) NEX protocol
type ShopNintendoBadgeArcadeProtocol struct {
	Server *nex.Server
	shop.ShopProtocol
	GetRivTokenHandler func(err error, client *nex.Client, callID uint32, itemCode string, referenceID []byte)
	PostPlayLogHandler func(err error, client *nex.Client, callID uint32, param *shop_nintendo_badge_arcade_types.ShopPostPlayLogParam)
}

// Setup initializes the protocol
func (protocol *ShopNintendoBadgeArcadeProtocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID && request.CustomID() == CustomProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.ShopProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *ShopNintendoBadgeArcadeProtocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetRivToken:
		go protocol.handleGetRivToken(packet)
	case MethodPostPlayLog:
		go protocol.handlePostPlayLog(packet)
	default:
		go globals.RespondNotImplementedCustom(packet, CustomProtocolID)
		fmt.Printf("Unsupported ShopNintendoBadgeArcade method ID: %#v\n", request.MethodID())
	}
}

// NewShopNintendoBadgeArcadeProtocol returns a new ShopNintendoBadgeArcadeProtocol
func NewShopNintendoBadgeArcadeProtocol(server *nex.Server) *ShopNintendoBadgeArcadeProtocol {
	protocol := &ShopNintendoBadgeArcadeProtocol{Server: server}
	protocol.ShopProtocol.Server = server

	protocol.Setup()

	return protocol
}
