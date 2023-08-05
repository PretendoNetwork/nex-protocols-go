// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetLawMessageResponse sets the GetLawMessageResponse handler function
func (protocol *Protocol) GetLawMessageResponse(handler func(err error, client *nex.Client, callID uint32, requestID uint32) uint32) {
	protocol.getLawMessageResponseHandler = handler
}

func (protocol *Protocol) handleGetLawMessageResponse(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getLawMessageResponseHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetLawMessageResponse not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	requestID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getLawMessageResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getLawMessageResponseHandler(nil, client, callID, requestID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
