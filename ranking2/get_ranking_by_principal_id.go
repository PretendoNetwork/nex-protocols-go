// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

func (protocol *Protocol) handleGetRankingByPrincipalID(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetRankingByPrincipalID == nil {
		globals.Logger.Warning("Ranking2::GetRankingByPrincipalID not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.server)

	getParam, err := nex.StreamReadStructure(parametersStream, ranking2_types.NewRanking2GetParam())
	if err != nil {
		_, errorCode = protocol.GetRankingByPrincipalID(fmt.Errorf("Failed to read getParam from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	principalIDList, err := parametersStream.ReadListPID()
	if err != nil {
		_, errorCode = protocol.GetRankingByPrincipalID(fmt.Errorf("Failed to read principalIDList from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetRankingByPrincipalID(nil, packet, callID, getParam, principalIDList)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
