// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleInvite(packet nex.PacketInterface) {
	var err error

	if protocol.Invite == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::Invite not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	idGathering := types.NewPrimitiveU32(0)
	err = idGathering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Invite(fmt.Errorf("Failed to read idGathering from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	lstPrincipals := types.NewList[*types.PID]()
	lstPrincipals.Type = types.NewPID(0)
	err = lstPrincipals.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Invite(fmt.Errorf("Failed to read lstPrincipals from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	strMessage := types.NewString("")
	err = strMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Invite(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.Invite(nil, packet, callID, idGathering, lstPrincipals, strMessage)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
