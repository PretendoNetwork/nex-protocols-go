// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

func (protocol *Protocol) handleGetBalanceRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetBalanceRequest == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::GetBalanceRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getBalanceParam, err := nex.StreamReadStructure(parametersStream, service_item_wii_sports_club_types.NewServiceItemGetBalanceParam())
	if err != nil {
		_, errorCode = protocol.GetBalanceRequest(fmt.Errorf("Failed to read getBalanceParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetBalanceRequest(nil, packet, callID, getBalanceParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
