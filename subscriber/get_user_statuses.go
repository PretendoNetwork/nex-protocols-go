// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetUserStatuses(packet nex.PacketInterface) {
	if protocol.GetUserStatuses == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Subscriber::GetUserStatuses not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var pids types.List[types.PID]
	var unknown types.List[types.UInt8]

	var err error

	err = pids.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetUserStatuses(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, pids, unknown)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetUserStatuses(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, pids, unknown)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetUserStatuses(nil, packet, callID, pids, unknown)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
