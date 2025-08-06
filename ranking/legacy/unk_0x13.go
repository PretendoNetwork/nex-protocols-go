// Package protocol implements the legacy Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

// TODO - Find name if possible
func (protocol *Protocol) handleUnk0x13(packet nex.PacketInterface) {
	if protocol.Unk0x13 == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::Unk0x13 not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	endpoint := packet.Sender().Endpoint()

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var unknown1 types.UInt32
	var unknown2 types.UInt32
	var unknown3 types.List[types.UInt32]
	var unknown4 types.List[types.UInt8]
	var unknown5 types.Bool
	var unknown6 types.UInt16

	var err error

	err = unknown1.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Unk0x13(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, unknown1, unknown2, unknown3, unknown4, unknown5, unknown6)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Unk0x13(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, unknown1, unknown2, unknown3, unknown4, unknown5, unknown6)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown3.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Unk0x13(fmt.Errorf("Failed to read unknown3 from parameters. %s", err.Error()), packet, callID, unknown1, unknown2, unknown3, unknown4, unknown5, unknown6)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown4.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Unk0x13(fmt.Errorf("Failed to read unknown4 from parameters. %s", err.Error()), packet, callID, unknown1, unknown2, unknown3, unknown4, unknown5, unknown6)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown5.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Unk0x13(fmt.Errorf("Failed to read unknown5 from parameters. %s", err.Error()), packet, callID, unknown1, unknown2, unknown3, unknown4, unknown5, unknown6)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown6.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Unk0x13(fmt.Errorf("Failed to read unknown6 from parameters. %s", err.Error()), packet, callID, unknown1, unknown2, unknown3, unknown4, unknown5, unknown6)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.Unk0x13(nil, packet, callID, unknown1, unknown2, unknown3, unknown4, unknown5, unknown6)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
