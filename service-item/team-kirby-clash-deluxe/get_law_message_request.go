// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handleGetLawMessageRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetLawMessageRequest == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetLawMessageRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getLawMessageParam, err := nex.StreamReadStructure(parametersStream, service_item_team_kirby_clash_deluxe_types.NewServiceItemGetLawMessageParam())
	if err != nil {
		_, errorCode = protocol.GetLawMessageRequest(fmt.Errorf("Failed to read getLawMessageParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetLawMessageRequest(nil, packet, callID, getLawMessageParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
