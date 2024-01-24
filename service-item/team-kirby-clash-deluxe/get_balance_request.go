// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handleGetBalanceRequest(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetBalanceRequest == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetBalanceRequest not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	getBalanceParam := service_item_team_kirby_clash_deluxe_types.NewServiceItemGetBalanceParam()
	err = getBalanceParam.ExtractFrom(parametersStream)
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
