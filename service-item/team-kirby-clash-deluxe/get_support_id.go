// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handleGetSupportID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetSupportID == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetSupportID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getSuppordIDParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemGetSupportIDParam())
	if err != nil {
		errorCode = protocol.GetSupportID(fmt.Errorf("Failed to read getSuppordIDParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.GetSupportID(nil, packet, callID, getSuppordIDParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemGetSupportIDParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
