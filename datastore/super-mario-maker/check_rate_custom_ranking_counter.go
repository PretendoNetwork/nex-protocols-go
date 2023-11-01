// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CheckRateCustomRankingCounter sets the CheckRateCustomRankingCounter handler function
func (protocol *Protocol) CheckRateCustomRankingCounter(handler func(err error, packet nex.PacketInterface, callID uint32, applicationID uint32) uint32) {
	protocol.checkRateCustomRankingCounterHandler = handler
}

func (protocol *Protocol) handleCheckRateCustomRankingCounter(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.checkRateCustomRankingCounterHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::CheckRateCustomRankingCounter not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.checkRateCustomRankingCounterHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.checkRateCustomRankingCounterHandler(nil, packet, callID, applicationID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
