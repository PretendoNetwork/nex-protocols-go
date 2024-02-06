// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleUpdateGatheringOwnership(packet nex.PacketInterface) {
	if protocol.UpdateGatheringOwnership == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::UpdateGatheringOwnership not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	gid := types.NewPrimitiveU32(0)
	participantsOnly := types.NewPrimitiveBool(false)

	var err error

	err = gid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateGatheringOwnership(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = participantsOnly.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateGatheringOwnership(fmt.Errorf("Failed to read participantsOnly from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateGatheringOwnership(nil, packet, callID, gid, participantsOnly)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
