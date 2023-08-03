// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/service-item/team-kirby-clash-deluxe/types"
)

// PostRightBinaryByAccount sets the PostRightBinaryByAccount handler function
func (protocol *Protocol) PostRightBinaryByAccount(handler func(err error, client *nex.Client, callID uint32, postRightBinaryByAccountParam *service_item_team_kirby_clash_deluxe_types.ServiceItemPostRightBinaryByAccountParam)) {
	protocol.postRightBinaryByAccountHandler = handler
}

func (protocol *Protocol) handlePostRightBinaryByAccount(packet nex.PacketInterface) {
	if protocol.postRightBinaryByAccountHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::PostRightBinaryByAccount not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	postRightBinaryByAccountParam, err := parametersStream.ReadStructure(service_item_team_kirby_clash_deluxe_types.NewServiceItemPostRightBinaryByAccountParam())
	if err != nil {
		go protocol.postRightBinaryByAccountHandler(fmt.Errorf("Failed to read postRightBinaryByAccountParam from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.postRightBinaryByAccountHandler(nil, client, callID, postRightBinaryByAccountParam.(*service_item_team_kirby_clash_deluxe_types.ServiceItemPostRightBinaryByAccountParam))
}
