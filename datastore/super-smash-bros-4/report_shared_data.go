// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleReportSharedData(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.ReportSharedData == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::ReportSharedData not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	dataID := types.NewPrimitiveU64(0)
	err = dataID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.ReportSharedData(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), packet, callID, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.ReportSharedData(nil, packet, callID, dataID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
