// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

func (protocol *Protocol) handlePutScore(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.PutScore == nil {
		globals.Logger.Warning("Ranking2::PutScore not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	scoreDataList, err := nex.StreamReadListStructure(parametersStream, ranking2_types.NewRanking2ScoreData())
	if err != nil {
		_, errorCode = protocol.PutScore(fmt.Errorf("Failed to read scoreDataList from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	nexUniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		_, errorCode = protocol.PutScore(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.PutScore(nil, packet, callID, scoreDataList, nexUniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
