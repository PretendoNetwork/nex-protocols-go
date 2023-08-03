// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CheckRateCustomRankingCounter sets the CheckRateCustomRankingCounter handler function
func (protocol *Protocol) CheckRateCustomRankingCounter(handler func(err error, client *nex.Client, callID uint32, applicationID uint32)) {
	protocol.checkRateCustomRankingCounterHandler = handler
}

func (protocol *Protocol) handleCheckRateCustomRankingCounter(packet nex.PacketInterface) {
	if protocol.checkRateCustomRankingCounterHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::CheckRateCustomRankingCounter not implemented")
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
		go protocol.checkRateCustomRankingCounterHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.checkRateCustomRankingCounterHandler(nil, client, callID, applicationID)
}
