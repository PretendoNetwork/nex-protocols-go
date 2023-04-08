package shop_nintendo_badge_arcade

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PostPlayLog sets the PostPlayLog function
func (protocol *ShopNintendoBadgeArcadeProtocol) PostPlayLog(handler func(err error, client *nex.Client, callID uint32, param *ShopPostPlayLogParam)) {
	protocol.PostPlayLogHandler = handler
}

func (protocol *ShopNintendoBadgeArcadeProtocol) HandlePostPlayLog(packet nex.PacketInterface) {
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

	param, err := parametersStream.ReadStructure(NewShopPostPlayLogParam())
	if err != nil {
		go protocol.PostPlayLogHandler(err, client, callID, nil)
		return
	}

	go protocol.PostPlayLogHandler(nil, client, callID, param.(*ShopPostPlayLogParam))
}
