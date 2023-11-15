// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateAndGetTicketInfo(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.UpdateAndGetTicketInfo == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::UpdateAndGetTicketInfo not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	forceRetrieveFromEShop, err := parametersStream.ReadBool()
	if err != nil {
		_, errorCode = protocol.UpdateAndGetTicketInfo(fmt.Errorf("Failed to read forceRetrieveFromEShop from parameters. %s", err.Error()), packet, callID, false)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.UpdateAndGetTicketInfo(nil, packet, callID, forceRetrieveFromEShop)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
