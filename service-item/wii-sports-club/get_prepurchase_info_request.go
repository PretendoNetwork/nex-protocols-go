// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

// GetPrepurchaseInfoRequest sets the GetPrepurchaseInfoRequest handler function
func (protocol *Protocol) GetPrepurchaseInfoRequest(handler func(err error, client *nex.Client, callID uint32, getPrepurchaseInfoParam *service_item_wii_sports_club_types.ServiceItemGetPrepurchaseInfoParam) uint32) {
	protocol.getPrepurchaseInfoRequestHandler = handler
}

func (protocol *Protocol) handleGetPrepurchaseInfoRequest(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getPrepurchaseInfoRequestHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::GetPrepurchaseInfoRequest not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getPrepurchaseInfoParam, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemGetPrepurchaseInfoParam())
	if err != nil {
		errorCode = protocol.getPrepurchaseInfoRequestHandler(fmt.Errorf("Failed to read getPrepurchaseInfoParam from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getPrepurchaseInfoRequestHandler(nil, client, callID, getPrepurchaseInfoParam.(*service_item_wii_sports_club_types.ServiceItemGetPrepurchaseInfoParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
