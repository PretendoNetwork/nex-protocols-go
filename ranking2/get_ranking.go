// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/ranking2/types"
)

func (protocol *Protocol) handleGetRanking(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.GetRanking == nil {
		globals.Logger.Warning("Ranking2::GetRanking not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	getParam := ranking2_types.NewRanking2GetParam()
	err = getParam.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.GetRanking(fmt.Errorf("Failed to read getParam from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.GetRanking(nil, packet, callID, getParam)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
