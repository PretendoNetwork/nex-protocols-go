// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handleGetServiceItemRightRequest(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetServiceItemRightRequest == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetServiceItemRightRequest not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	getServiceItemRightParam := service_item_team_kirby_clash_deluxe_types.NewServiceItemGetServiceItemRightParam()
	err = getServiceItemRightParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetServiceItemRightRequest(fmt.Errorf("Failed to read getServiceItemRightParam from parameters. %s", err.Error()), packet, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	withoutRightBinary := types.NewPrimitiveBool(false)
	err = withoutRightBinary.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetServiceItemRightRequest(fmt.Errorf("Failed to read withoutRightBinary from parameters. %s", err.Error()), packet, callID, nil, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetServiceItemRightRequest(nil, packet, callID, getServiceItemRightParam, withoutRightBinary)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
