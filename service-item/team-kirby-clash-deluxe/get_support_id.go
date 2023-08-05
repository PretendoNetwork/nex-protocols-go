// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// GetSupportID sets the GetSupportID handler function
func (protocol *Protocol) GetSupportID(handler func(err error, client *nex.Client, callID uint32, getSuppordIDParam *service_item_team_kirby_clash_deluxe_types.ServiceItemGetSupportIDParam) uint32) {
	protocol.getSupportIDHandler = handler
}

func (protocol *Protocol) handleGetSupportID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getSupportIDHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetSupportID not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getSuppordIDParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemGetSupportIDParam())
	if err != nil {
		errorCode = protocol.getSupportIDHandler(fmt.Errorf("Failed to read getSuppordIDParam from parameters. %s", err.Error()), client, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getSupportIDHandler(nil, client, callID, getSuppordIDParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemGetSupportIDParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
