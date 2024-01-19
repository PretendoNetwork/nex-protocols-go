// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleCheckRateCustomRankingCounter(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.CheckRateCustomRankingCounter == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::CheckRateCustomRankingCounter not implemented")
		globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	applicationID := types.NewPrimitiveU32(0)
	err = applicationID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.CheckRateCustomRankingCounter(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.CheckRateCustomRankingCounter(nil, packet, callID, applicationID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
