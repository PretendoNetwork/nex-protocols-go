// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetPrepurchaseInfoResponse sets the GetPrepurchaseInfoResponse handler function
func (protocol *Protocol) GetPrepurchaseInfoResponse(handler func(err error, client *nex.Client, callID uint32, requestID uint32)) {
	protocol.getPrepurchaseInfoResponseHandler = handler
}

func (protocol *Protocol) handleGetPrepurchaseInfoResponse(packet nex.PacketInterface) {
	if protocol.getPrepurchaseInfoResponseHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetPrepurchaseInfoResponse not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	requestID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getPrepurchaseInfoResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getPrepurchaseInfoResponseHandler(nil, client, callID, requestID)
}