// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateAndGetTicketInfo sets the UpdateAndGetTicketInfo handler function
func (protocol *Protocol) UpdateAndGetTicketInfo(handler func(err error, packet nex.PacketInterface, callID uint32, forceRetrieveFromEShop bool) uint32) {
	protocol.updateAndGetTicketInfoHandler = handler
}

func (protocol *Protocol) handleUpdateAndGetTicketInfo(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.updateAndGetTicketInfoHandler == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::UpdateAndGetTicketInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	forceRetrieveFromEShop, err := parametersStream.ReadBool()
	if err != nil {
		errorCode = protocol.updateAndGetTicketInfoHandler(fmt.Errorf("Failed to read forceRetrieveFromEShop from parameters. %s", err.Error()), packet, callID, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.updateAndGetTicketInfoHandler(nil, packet, callID, forceRetrieveFromEShop)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
