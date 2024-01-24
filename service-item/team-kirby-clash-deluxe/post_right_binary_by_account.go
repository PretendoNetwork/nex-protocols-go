// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handlePostRightBinaryByAccount(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.PostRightBinaryByAccount == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::PostRightBinaryByAccount not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	postRightBinaryByAccountParam := service_item_team_kirby_clash_deluxe_types.NewServiceItemPostRightBinaryByAccountParam()
	err = postRightBinaryByAccountParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.PostRightBinaryByAccount(fmt.Errorf("Failed to read postRightBinaryByAccountParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.PostRightBinaryByAccount(nil, packet, callID, postRightBinaryByAccountParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
