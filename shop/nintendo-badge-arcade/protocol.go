// Package protocol implements the Nintendo Badge Arcade Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	shop "github.com/PretendoNetwork/nex-protocols-go/shop"
	shop_nintendo_badge_arcade_types "github.com/PretendoNetwork/nex-protocols-go/shop/nintendo-badge-arcade/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Shop (Nintendo Badge Arcade) protocol
	ProtocolID = 0x7F

	// CustomProtocolID is the Custom ID for the Shop (Nintendo Badge Arcade) protocol
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

type shopProtocol = shop.Protocol

// Protocol stores all the RMC method handlers for the Shop (Nintendo Badge Arcade) protocol and listens for requests
// Embeds the Shop protocol
type Protocol struct {
	Server *nex.Server
	shopProtocol
	getRivTokenHandler func(err error, packet nex.PacketInterface, callID uint32, itemCode string, referenceID []byte) uint32
	postPlayLogHandler func(err error, packet nex.PacketInterface, callID uint32, param *shop_nintendo_badge_arcade_types.ShopPostPlayLogParam) uint32
}

// Setup initializes the protocol
func (protocol *Protocol) Setup() {
	protocol.Server.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if request.ProtocolID() == ProtocolID && request.CustomID() == CustomProtocolID {
			if slices.Contains(patchedMethods, request.MethodID()) {
				protocol.HandlePacket(packet)
			} else {
				protocol.shopProtocol.HandlePacket(packet)
			}
		}
	})
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	request := packet.RMCRequest()

	switch request.MethodID() {
	case MethodGetRivToken:
		go protocol.handleGetRivToken(packet)
	case MethodPostPlayLog:
		go protocol.handlePostPlayLog(packet)
	default:
		go globals.RespondErrorCustom(packet, CustomProtocolID, nex.Errors.Core.NotImplemented)
		fmt.Printf("Unsupported ShopNintendoBadgeArcade method ID: %#v\n", request.MethodID())
	}
}

// NewProtocol returns a new Shop (Nintendo Badge Arcade)
func NewProtocol(server *nex.Server) *Protocol {
	protocol := &Protocol{Server: server}
	protocol.shopProtocol.Server = server

	protocol.Setup()

	return protocol
}
