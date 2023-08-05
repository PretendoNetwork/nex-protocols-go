// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// AcquireServiceItemByAccount sets the AcquireServiceItemByAccount handler function
func (protocol *Protocol) AcquireServiceItemByAccount(handler func(err error, client *nex.Client, callID uint32, acquireServiceItemByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemAcquireServiceItemByAccountParam) uint32) {
	protocol.acquireServiceItemByAccountHandler = handler
}

func (protocol *Protocol) handleAcquireServiceItemByAccount(packet nex.PacketInterface) {
	if protocol.acquireServiceItemByAccountHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::AcquireServiceItemByAccount not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	acquireServiceItemByAccountParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemAcquireServiceItemByAccountParam())
	if err != nil {
		go protocol.acquireServiceItemByAccountHandler(fmt.Errorf("Failed to read acquireServiceItemByAccountParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.acquireServiceItemByAccountHandler(nil, client, callID, acquireServiceItemByAccountParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemAcquireServiceItemByAccountParam))
}
