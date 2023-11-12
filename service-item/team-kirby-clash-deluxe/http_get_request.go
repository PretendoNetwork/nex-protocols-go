// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// HTTPGetRequest sets the HTTPGetRequest handler function
func (protocol *Protocol) HTTPGetRequest(handler func(err error, packet nex.PacketInterface, callID uint32, url *service_item_team_kirby_clash_deluxe_types.ServiceItemHTTPGetParam) uint32) {
	protocol.httpGetRequestHandler = handler
}

func (protocol *Protocol) handleHTTPGetRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.httpGetRequestHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::HTTPGetRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	url, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemHTTPGetParam())
	if err != nil {
		errorCode = protocol.httpGetRequestHandler(fmt.Errorf("Failed to read url from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.httpGetRequestHandler(nil, packet, callID, url.(*service_item_team_kirby_clash_deluxe_types.ServiceItemHTTPGetParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
