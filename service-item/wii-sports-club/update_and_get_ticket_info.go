// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateAndGetTicketInfo(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.UpdateAndGetTicketInfo == nil {
		globals.Logger.Warning("ServiceItemWiiSportsClub::UpdateAndGetTicketInfo not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	forceRetrieveFromEShop := types.NewPrimitiveBool(false)
	err = forceRetrieveFromEShop.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.UpdateAndGetTicketInfo(fmt.Errorf("Failed to read forceRetrieveFromEShop from parameters. %s", err.Error()), packet, callID, nil)
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
