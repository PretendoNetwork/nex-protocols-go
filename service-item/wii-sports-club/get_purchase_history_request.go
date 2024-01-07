// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

func (protocol *Protocol) handleGetPurchaseHistoryRequest(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetPurchaseHistoryRequest == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::GetPurchaseHistoryRequest not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	getPurchaseHistoryParam := service_item_wii_sports_club_types.NewServiceItemGetPurchaseHistoryParam()
	err = getPurchaseHistoryParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetPurchaseHistoryRequest(fmt.Errorf("Failed to read getPurchaseHistoryParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetPurchaseHistoryRequest(nil, packet, callID, getPurchaseHistoryParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
