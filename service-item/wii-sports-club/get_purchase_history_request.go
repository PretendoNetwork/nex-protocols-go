// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPurchaseHistoryRequest sets the GetPurchaseHistoryRequest handler function
func (protocol *Protocol) GetPurchaseHistoryRequest(handler func(err error, client *nex.Client, callID uint32, getPurchaseHistoryParam *service_item_wii_sports_club_types.ServiceItemGetPurchaseHistoryParam)) {
	protocol.getPurchaseHistoryRequestHandler = handler
}

func (protocol *Protocol) handleGetPurchaseHistoryRequest(packet nex.PacketInterface) {
	if protocol.getPurchaseHistoryRequestHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::GetPurchaseHistoryRequest not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getPurchaseHistoryParam, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemGetPurchaseHistoryParam())
	if err != nil {
		go protocol.getPurchaseHistoryRequestHandler(fmt.Errorf("Failed to read getPurchaseHistoryParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getPurchaseHistoryRequestHandler(nil, client, callID, getPurchaseHistoryParam.(*service_item_wii_sports_club_types.ServiceItemGetPurchaseHistoryParam))
}