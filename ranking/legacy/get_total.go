// Package protocol implements the legacy Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetTotal(packet nex.PacketInterface) {
	if protocol.GetTotal == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::GetTotal not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	endpoint := packet.Sender().Endpoint()

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var uniqueID types.UInt32
	var unknown1 types.UInt8
	var unknown2 types.UInt8
	var unknown3 types.UInt8
	var unknown4 types.UInt32

	var err error

	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetTotal(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, uniqueID, unknown1, unknown2, unknown3, unknown4)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown1.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetTotal(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, uniqueID, unknown1, unknown2, unknown3, unknown4)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetTotal(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, uniqueID, unknown1, unknown2, unknown3, unknown4)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown3.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetTotal(fmt.Errorf("Failed to read unknown3 from parameters. %s", err.Error()), packet, callID, uniqueID, unknown1, unknown2, unknown3, unknown4)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown4.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetTotal(fmt.Errorf("Failed to read unknown4 from parameters. %s", err.Error()), packet, callID, uniqueID, unknown1, unknown2, unknown3, unknown4)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetTotal(nil, packet, callID, uniqueID, unknown1, unknown2, unknown3, unknown4)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
