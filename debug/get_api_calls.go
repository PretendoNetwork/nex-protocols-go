// Package protocol implements the Debug protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetAPICalls(packet nex.PacketInterface) {
	if protocol.GetAPICalls == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Debug::GetAPICalls not implemented")

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
	var unknown types.DateTime
	var unknown2 types.DateTime

	var err error

	err = pids.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetAPICalls(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, pids, unknown, unknown2)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetAPICalls(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, pids, unknown, unknown2)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetAPICalls(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, pids, unknown, unknown2)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetAPICalls(nil, packet, callID, pids, unknown, unknown2)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
