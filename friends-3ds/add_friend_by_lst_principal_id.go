// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleAddFriendBylstPrincipalID(packet nex.PacketInterface) {
	if protocol.AddFriendBylstPrincipalID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends3DS::AddFriendBylstPrincipalID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	lfc := types.NewPrimitiveU64(0)
	pids := types.NewList[*types.PID]()
	pids.Type = types.NewPID(0)

	var err error

	err = lfc.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddFriendBylstPrincipalID(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = pids.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.AddFriendBylstPrincipalID(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.AddFriendBylstPrincipalID(nil, packet, callID, lfc, pids)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
