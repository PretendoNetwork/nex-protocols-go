// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetBalanceResponse sets the GetBalanceResponse handler function
func (protocol *Protocol) GetBalanceResponse(handler func(err error, client *nex.Client, callID uint32, requestID uint32)) {
	protocol.getBalanceResponseHandler = handler
}

func (protocol *Protocol) handleGetBalanceResponse(packet nex.PacketInterface) {
	if protocol.getBalanceResponseHandler == nil {
		globals.Logger.Warning("ServiceItemTeamKirbyClashDeluxe::GetBalanceResponse not implemented")
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
		go protocol.getBalanceResponseHandler(fmt.Errorf("Failed to read requestID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getBalanceResponseHandler(nil, client, callID, requestID)
}
