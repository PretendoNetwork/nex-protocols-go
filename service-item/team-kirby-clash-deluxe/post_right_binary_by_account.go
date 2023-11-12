// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// PostRightBinaryByAccount sets the PostRightBinaryByAccount handler function
func (protocol *Protocol) PostRightBinaryByAccount(handler func(err error, packet nex.PacketInterface, callID uint32, postRightBinaryByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemPostRightBinaryByAccountParam) uint32) {
	protocol.postRightBinaryByAccountHandler = handler
}

func (protocol *Protocol) handlePostRightBinaryByAccount(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.postRightBinaryByAccountHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::PostRightBinaryByAccount not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	postRightBinaryByAccountParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemPostRightBinaryByAccountParam())
	if err != nil {
		errorCode = protocol.postRightBinaryByAccountHandler(fmt.Errorf("Failed to read postRightBinaryByAccountParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.postRightBinaryByAccountHandler(nil, packet, callID, postRightBinaryByAccountParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemPostRightBinaryByAccountParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
