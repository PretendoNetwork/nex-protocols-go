// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

func (protocol *Protocol) handlePutScore(packet nex.PacketInterface) {
	var err error

	if protocol.PutScore == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking2::PutScore not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	scoreDataList := types.NewList[*ranking2_types.Ranking2ScoreData]()
	scoreDataList.Type = ranking2_types.NewRanking2ScoreData()
	err = scoreDataList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PutScore(fmt.Errorf("Failed to read scoreDataList from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	nexUniqueID := types.NewPrimitiveU64(0)
	err = nexUniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PutScore(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PutScore(nil, packet, callID, scoreDataList, nexUniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
