// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handleAcquireServiceItemByAccount(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.AcquireServiceItemByAccount == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::AcquireServiceItemByAccount not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	acquireServiceItemByAccountParam := service_item_team_kirby_clash_deluxe_types.NewServiceItemAcquireServiceItemByAccountParam()
	err = acquireServiceItemByAccountParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.AcquireServiceItemByAccount(fmt.Errorf("Failed to read acquireServiceItemByAccountParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.AcquireServiceItemByAccount(nil, packet, callID, acquireServiceItemByAccountParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
