// Package protocol implements the Nintendo Badge Arcade Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	shop_nintendo_badge_arcade_types "github.com/PretendoNetwork/nex-protocols-go/shop/nintendo-badge-arcade/types"
)

// PostPlayLog sets the PostPlayLog function
func (protocol *Protocol) PostPlayLog(handler func(err error, client *nex.Client, callID uint32, param *shop_nintendo_badge_arcade_types.ShopPostPlayLogParam)) {
	protocol.PostPlayLogHandler = handler
}

func (protocol *Protocol) handlePostPlayLog(packet nex.PacketInterface) {
	if protocol.PostPlayLogHandler == nil {
		globals.Logger.Warning("ShopNintendoBadgeArcade::PostPlayLog not implemented")
		go globals.RespondNotImplementedCustom(packet, CustomProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	param, err := parametersStream.ReadStructure(shop_nintendo_badge_arcade_types.NewShopPostPlayLogParam())
	if err != nil {
		go protocol.PostPlayLogHandler(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.PostPlayLogHandler(nil, client, callID, param.(*shop_nintendo_badge_arcade_types.ShopPostPlayLogParam))
}
