// Package shop_nintendo_badge_arcade implements the Nintendo Badge Arcade Shop NEX protocol
package shop_nintendo_badge_arcade

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetRivToken sets the GetRivToken function
func (protocol *ShopNintendoBadgeArcadeProtocol) GetRivToken(handler func(err error, client *nex.Client, callID uint32, itemCode string, referenceID []byte)) {
	protocol.GetRivTokenHandler = handler
}

func (protocol *ShopNintendoBadgeArcadeProtocol) handleGetRivToken(packet nex.PacketInterface) {
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
		go protocol.GetRivTokenHandler(fmt.Errorf("Failed to read itemCode from parameters. %s", err.Error()), client, callID, "", nil)
		return
	}

	referenceID, err := parametersStream.ReadQBuffer()
	if err != nil {
		go protocol.GetRivTokenHandler(fmt.Errorf("Failed to read referenceID from parameters. %s", err.Error()), client, callID, "", nil)
		return
	}

	go protocol.GetRivTokenHandler(nil, client, callID, itemCode, referenceID)
}
