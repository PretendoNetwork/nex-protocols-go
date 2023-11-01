// Package protocol implements the Nintendo Badge Arcade Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	shop_nintendo_badge_arcade_types "github.com/PretendoNetwork/nex-protocols-go/shop/nintendo-badge-arcade/types"
)

// PostPlayLog sets the PostPlayLog function
func (protocol *Protocol) PostPlayLog(handler func(err error, packet nex.PacketInterface, callID uint32, param *shop_nintendo_badge_arcade_types.ShopPostPlayLogParam) uint32) {
	protocol.postPlayLogHandler = handler
}

func (protocol *Protocol) handlePostPlayLog(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.postPlayLogHandler == nil {
		globals.Logger.Warning("ShopNintendoBadgeArcade::PostPlayLog not implemented")
		go globals.RespondErrorCustom(packet, CustomProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(shop_nintendo_badge_arcade_types.NewShopPostPlayLogParam())
	if err != nil {
		errorCode = protocol.postPlayLogHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondErrorCustom(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.postPlayLogHandler(nil, packet, callID, param.(*shop_nintendo_badge_arcade_types.ShopPostPlayLogParam))
	if errorCode != 0 {
		globals.RespondErrorCustom(packet, ProtocolID, errorCode)
	}
}
