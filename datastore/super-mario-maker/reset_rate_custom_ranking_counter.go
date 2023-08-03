// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ResetRateCustomRankingCounter sets the ResetRateCustomRankingCounter handler function
func (protocol *Protocol) ResetRateCustomRankingCounter(handler func(err error, client *nex.Client, callID uint32, applicationID uint32)) {
	protocol.resetRateCustomRankingCounterHandler = handler
}

func (protocol *Protocol) handleResetRateCustomRankingCounter(packet nex.PacketInterface) {
	if protocol.resetRateCustomRankingCounterHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::ResetRateCustomRankingCounter not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.resetRateCustomRankingCounterHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.resetRateCustomRankingCounterHandler(nil, client, callID, applicationID)
}
