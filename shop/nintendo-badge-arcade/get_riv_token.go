package shop_nintendo_badge_arcade

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRivToken sets the GetRivToken function
func (protocol *ShopNintendoBadgeArcadeProtocol) GetRivToken(handler func(err error, client *nex.Client, callID uint32, itemCode string, referenceID []byte)) {
	protocol.GetRivTokenHandler = handler
}

func (protocol *ShopNintendoBadgeArcadeProtocol) HandleGetRivToken(packet nex.PacketInterface) {
	if protocol.GetRivTokenHandler == nil {
		globals.Logger.Warning("ShopNintendoBadgeArcade::GetRivToken not implemented")
		go globals.RespondNotImplementedCustom(packet, CustomProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	itemCode, err := parametersStream.ReadString()
	if err != nil {
		go protocol.GetRivTokenHandler(err, client, callID, "", []byte{})
		return
	}

	referenceID, err := parametersStream.ReadQBuffer()
	if err != nil {
		go protocol.GetRivTokenHandler(err, client, callID, "", []byte{})
		return
	}

	go protocol.GetRivTokenHandler(nil, client, callID, itemCode, referenceID)
}