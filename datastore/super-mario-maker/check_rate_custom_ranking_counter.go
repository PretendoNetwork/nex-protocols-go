package datastore_super_mario_maker

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CheckRateCustomRankingCounter sets the CheckRateCustomRankingCounter handler function
func (protocol *DataStoreSuperMarioMakerProtocol) CheckRateCustomRankingCounter(handler func(err error, client *nex.Client, callID uint32, applicationID uint32)) {
	protocol.CheckRateCustomRankingCounterHandler = handler
}

func (protocol *DataStoreSuperMarioMakerProtocol) handleCheckRateCustomRankingCounter(packet nex.PacketInterface) {
	if protocol.CheckRateCustomRankingCounterHandler == nil {
		globals.Logger.Warning("DataStoreSMM::CheckRateCustomRankingCounter not implemented")
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
		go protocol.CheckRateCustomRankingCounterHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.CheckRateCustomRankingCounterHandler(nil, client, callID, applicationID)
}
