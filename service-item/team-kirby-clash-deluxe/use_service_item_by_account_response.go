// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UseServiceItemByAccountResponse sets the UseServiceItemByAccountResponse handler function
func (protocol *Protocol) UseServiceItemByAccountResponse(handler func(err error, packet nex.PacketInterface, callID uint32, requestID uint32) uint32) {
	protocol.useServiceItemByAccountResponseHandler = handler
}

func (protocol *Protocol) handleUseServiceItemByAccountResponse(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.useServiceItemByAccountResponseHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::UseServiceItemByAccountResponse not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	requestID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.useServiceItemByAccountResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.useServiceItemByAccountResponseHandler(nil, packet, callID, requestID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
