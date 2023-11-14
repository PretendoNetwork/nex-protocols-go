// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/ranking/types"
)

func (protocol *Protocol) handleChangeAllAttributes(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.ChangeAllAttributes == nil {
		globals.Logger.Warning("Ranking::ChangeAllAttributes not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	changeParam, err := parametersStream.ReadStructure(ranking_types.NewRankingChangeAttributesParam())
	if err != nil {
		errorCode = protocol.ChangeAllAttributes(fmt.Errorf("Failed to read changeParam from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.ChangeAllAttributes(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), packet, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.ChangeAllAttributes(nil, packet, callID, changeParam.(*ranking_types.RankingChangeAttributesParam), uniqueID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
