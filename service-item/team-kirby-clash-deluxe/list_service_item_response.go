// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ListServiceItemResponse sets the ListServiceItemResponse handler function
func (protocol *Protocol) ListServiceItemResponse(handler func(err error, client *nex.Client, callID uint32, requestID uint32)) {
	protocol.listServiceItemResponseHandler = handler
}

func (protocol *Protocol) handleListServiceItemResponse(packet nex.PacketInterface) {
	if protocol.listServiceItemResponseHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::ListServiceItemResponse not implemented")
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
		go protocol.listServiceItemResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.listServiceItemResponseHandler(nil, client, callID, requestID)
}