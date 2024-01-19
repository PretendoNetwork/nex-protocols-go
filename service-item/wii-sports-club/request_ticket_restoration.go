// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

func (protocol *Protocol) handleRequestTicketRestoration(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.RequestTicketRestoration == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::RequestTicketRestoration not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	requestTicketRestorationParam := service_item_wii_sports_club_types.NewServiceItemRequestTicketRestorationParam()
	err = requestTicketRestorationParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.RequestTicketRestoration(fmt.Errorf("Failed to read requestTicketRestorationParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.RequestTicketRestoration(nil, packet, callID, requestTicketRestorationParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
