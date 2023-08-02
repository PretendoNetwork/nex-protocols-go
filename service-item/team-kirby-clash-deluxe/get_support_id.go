// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// GetSupportID sets the GetSupportID handler function
func (protocol *Protocol) GetSupportID(handler func(err error, client *nex.Client, callID uint32, getSuppordIDParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetSupportIDParam)) {
	protocol.getSupportIDHandler = handler
}

func (protocol *Protocol) handleGetSupportID(packet nex.PacketInterface) {
	if protocol.getSupportIDHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetSupportID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getSuppordIDParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemGetSupportIDParam())
	if err != nil {
		go protocol.getSupportIDHandler(fmt.Errorf("Failed to read getSuppordIDParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.getSupportIDHandler(nil, client, callID, getSuppordIDParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemGetSupportIDParam))
}
