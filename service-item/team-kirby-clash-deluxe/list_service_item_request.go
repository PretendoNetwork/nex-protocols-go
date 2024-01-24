// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handleListServiceItemRequest(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.ListServiceItemRequest == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::ListServiceItemRequest not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	listServiceItemParam := service_item_team_kirby_clash_deluxe_types.NewServiceItemListServiceItemParam()
	err = listServiceItemParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ListServiceItemRequest(fmt.Errorf("Failed to read listServiceItemParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ListServiceItemRequest(nil, packet, callID, listServiceItemParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
