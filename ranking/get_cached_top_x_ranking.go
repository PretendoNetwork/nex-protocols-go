// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
)

func (protocol *Protocol) handleGetCachedTopXRanking(packet nex.PacketInterface) {
	if protocol.GetCachedTopXRanking == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking::GetCachedTopXRanking not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var category types.UInt32
	orderParam := ranking_types.NewRankingOrderParam()

	var err error

	err = category.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCachedTopXRanking(fmt.Errorf("Failed to read category from parameters. %s", err.Error()), packet, callID, category, orderParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = orderParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCachedTopXRanking(fmt.Errorf("Failed to read orderParam from parameters. %s", err.Error()), packet, callID, category, orderParam)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetCachedTopXRanking(nil, packet, callID, category, orderParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
