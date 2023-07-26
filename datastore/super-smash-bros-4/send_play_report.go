// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SendPlayReport sets the SendPlayReport handler function
func (protocol *Protocol) SendPlayReport(handler func(err error, client *nex.Client, callID uint32, playReport []int32)) {
	protocol.SendPlayReportHandler = handler
}

func (protocol *Protocol) handleSendPlayReport(packet nex.PacketInterface) {
	if protocol.SendPlayReportHandler == nil {
		globals.Logger.Warning("DataStoreSmash4::SendPlayReport not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	playReport, err := parametersStream.ReadListInt32LE()
	if err != nil {
		go protocol.SendPlayReportHandler(fmt.Errorf("Failed to read playReport from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.SendPlayReportHandler(nil, client, callID, playReport)
}
