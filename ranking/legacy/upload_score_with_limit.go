// Package protocol implements the legacy Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUploadScoreWithLimit(packet nex.PacketInterface) {
	if protocol.UploadScoreWithLimit == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::UploadScoreWithLimit not implemented")

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
	var category types.UInt32
	var scores types.List[types.UInt32]
	var unknown1 types.UInt8
	var unknown2 types.UInt32
	var limit types.UInt16

	var err error

	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UploadScoreWithLimit(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, uniqueID, category, scores, unknown1, unknown2, limit)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = category.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UploadScoreWithLimit(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, uniqueID, category, scores, unknown1, unknown2, limit)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = scores.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UploadScoreWithLimit(fmt.Errorf("Failed to read scores from parameters. %s", err.Error()), packet, callID, uniqueID, category, scores, unknown1, unknown2, limit)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown1.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UploadScoreWithLimit(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, uniqueID, category, scores, unknown1, unknown2, limit)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UploadScoreWithLimit(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, uniqueID, category, scores, unknown1, unknown2, limit)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = limit.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UploadScoreWithLimit(fmt.Errorf("Failed to read limit from parameters. %s", err.Error()), packet, callID, uniqueID, category, scores, unknown1, unknown2, limit)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UploadScoreWithLimit(nil, packet, callID, uniqueID, category, scores, unknown1, unknown2, limit)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
