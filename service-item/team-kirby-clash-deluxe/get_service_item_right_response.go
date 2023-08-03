// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetServiceItemRightResponse sets the GetServiceItemRightResponse handler function
func (protocol *Protocol) GetServiceItemRightResponse(handler func(err error, client *nex.Client, callID uint32, requestID uint32)) {
	protocol.getServiceItemRightResponseHandler = handler
}

func (protocol *Protocol) handleGetServiceItemRightResponse(packet nex.PacketInterface) {
	if protocol.getServiceItemRightResponseHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetServiceItemRightResponse not implemented")
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
		go protocol.getServiceItemRightResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getServiceItemRightResponseHandler(nil, client, callID, requestID)
}
