// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handleGetPurchaseHistoryRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetPurchaseHistoryRequest == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetPurchaseHistoryRequest not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	getPurchaseHistoryParam, err := nex.StreamReadStructure(parametersStream, service_item_team_kirby_clash_deluxe_types.NewServiceItemGetPurchaseHistoryParam())
	if err != nil {
		_, errorCode = protocol.GetPurchaseHistoryRequest(fmt.Errorf("Failed to read getPurchaseHistoryParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetPurchaseHistoryRequest(nil, packet, callID, getPurchaseHistoryParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
