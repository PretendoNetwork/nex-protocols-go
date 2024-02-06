// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleGetBasicInfo(packet nex.PacketInterface) {
	if protocol.GetBasicInfo == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "FriendsWiiU::GetBasicInfo not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	pids := types.NewList[*types.PID]()
	pids.Type = types.NewPID(0)

	err := pids.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetBasicInfo(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetBasicInfo(nil, packet, callID, pids)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
