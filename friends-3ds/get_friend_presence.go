// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetFriendPresence(packet nex.PacketInterface) {
	if protocol.GetFriendPresence == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends3DS::GetFriendPresence not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var pidList types.List[types.PID]

	err := pidList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetFriendPresence(fmt.Errorf("Failed to read pidList from parameters. %s", err.Error()), packet, callID, pidList)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetFriendPresence(nil, packet, callID, pidList)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
