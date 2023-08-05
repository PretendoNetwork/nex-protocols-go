// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/service-item/wii-sports-club/types"
)

// RequestTicketRestoration sets the RequestTicketRestoration handler function
func (protocol *Protocol) RequestTicketRestoration(handler func(err error, client *nex.Client, callID uint32, requestTicketRestorationParam *service_item_wii_sports_club_types.ServiceItemRequestTicketRestorationParam) uint32) {
	protocol.requestTicketRestorationHandler = handler
}

func (protocol *Protocol) handleRequestTicketRestoration(packet nex.PacketInterface) {
	if protocol.requestTicketRestorationHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::RequestTicketRestoration not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	requestTicketRestorationParam, err := parametersStream.ReadStructure(service_item_wii_sports_club_types.NewServiceItemRequestTicketRestorationParam())
	if err != nil {
		go protocol.requestTicketRestorationHandler(fmt.Errorf("Failed to read requestTicketRestorationParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.requestTicketRestorationHandler(nil, client, callID, requestTicketRestorationParam.(*service_item_wii_sports_club_types.ServiceItemRequestTicketRestorationParam))
}
