// Package protocol implements the legacy Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/legacy/types"
)

// TODO - Find name if possible
func (protocol *Protocol) handleUnk0xD(packet nex.PacketInterface) {
	if protocol.Unk0xD == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::Unk0xD not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	endpoint := packet.Sender().Endpoint()
	rankingVersion := endpoint.LibraryVersions().Ranking

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var uniqueID types.UInt32
	var category types.UInt32
	var orderParam ranking_types.RankingOrderParam

	var err error

	err = uniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Unk0xD(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, uniqueID, category, orderParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	if rankingVersion.GreaterOrEqual("2.0.0") {
		err = category.ExtractFrom(parametersStream)
		if err != nil {
			_, rmcError := protocol.Unk0xD(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, uniqueID, category, orderParam)
			if rmcError != nil {
				globals.RespondError(packet, ProtocolID, rmcError)
			}

			return
		}
	} else {
		var categories types.List[types.UInt16]

		err = categories.ExtractFrom(parametersStream)
		if err != nil {
			_, rmcError := protocol.Unk0xD(fmt.Errorf("Failed to read categories from parameters. %s", err.Error()), packet, callID, uniqueID, category, orderParam)
			if rmcError != nil {
				globals.RespondError(packet, ProtocolID, rmcError)
			}

			return
		}

		if len(categories) != 1 {
			_, rmcError := protocol.Unk0xD(fmt.Errorf("Failed to read categories from parameters. Expected length of 1, got %d", len(categories)), packet, callID, uniqueID, category, orderParam)
			if rmcError != nil {
				globals.RespondError(packet, ProtocolID, rmcError)
			}

			return
		}

		category = types.UInt32(categories[0])
	}

	err = orderParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Unk0xD(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, uniqueID, category, orderParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.Unk0xD(nil, packet, callID, uniqueID, category, orderParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
