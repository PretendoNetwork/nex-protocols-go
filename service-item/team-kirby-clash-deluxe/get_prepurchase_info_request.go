// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handleGetPrepurchaseInfoRequest(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetPrepurchaseInfoRequest == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetPrepurchaseInfoRequest not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	getPrepurchaseInfoParam := service_item_team_kirby_clash_deluxe_types.NewServiceItemGetPrepurchaseInfoParam()
	err = getPrepurchaseInfoParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetPrepurchaseInfoRequest(fmt.Errorf("Failed to read getPrepurchaseInfoParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetPrepurchaseInfoRequest(nil, packet, callID, getPrepurchaseInfoParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
