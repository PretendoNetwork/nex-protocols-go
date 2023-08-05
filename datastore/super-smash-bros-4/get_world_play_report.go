// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetWorldPlayReport sets the GetWorldPlayReport handler function
func (protocol *Protocol) GetWorldPlayReport(handler func(err error, client *nex.Client, callID uint32) uint32) {
	protocol.getWorldPlayReportHandler = handler
}

func (protocol *Protocol) handleGetWorldPlayReport(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getWorldPlayReportHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetWorldPlayReport not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	errorCode = protocol.getWorldPlayReportHandler(nil, client, callID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
