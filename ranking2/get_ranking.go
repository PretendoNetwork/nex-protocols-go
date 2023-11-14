// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

func (protocol *Protocol) handleGetRanking(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.GetRanking == nil {
		globals.Logger.Warning("Ranking2::GetRanking not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	getParam, err := parametersStream.ReadStructure(ranking2_types.NewRanking2GetParam())
	if err != nil {
		errorCode = protocol.GetRanking(fmt.Errorf("Failed to read getParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.GetRanking(nil, packet, callID, getParam.(*ranking2_types.Ranking2GetParam))
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
