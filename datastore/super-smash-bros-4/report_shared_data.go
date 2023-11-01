// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReportSharedData sets the ReportSharedData handler function
func (protocol *Protocol) ReportSharedData(handler func(err error, packet nex.PacketInterface, callID uint32, dataID uint64) uint32) {
	protocol.reportSharedDataHandler = handler
}

func (protocol *Protocol) handleReportSharedData(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.reportSharedDataHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::ReportSharedData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.reportSharedDataHandler(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.reportSharedDataHandler(nil, packet, callID, dataID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
