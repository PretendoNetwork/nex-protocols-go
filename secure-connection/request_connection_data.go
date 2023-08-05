// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestConnectionData sets the RequestConnectionData handler function
func (protocol *Protocol) RequestConnectionData(handler func(err error, client *nex.Client, callID uint32, cidTarget uint32, pidTarget uint32) uint32) {
	protocol.requestConnectionDataHandler = handler
}

func (protocol *Protocol) handleRequestConnectionData(packet nex.PacketInterface) {
	if protocol.requestConnectionDataHandler == nil {
		globals.Logger.Warning("SecureConnection::RequestConnectionData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	cidTarget, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.requestConnectionDataHandler(fmt.Errorf("Failed to read cidTarget from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	pidTarget, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.requestConnectionDataHandler(fmt.Errorf("Failed to read pidTarget from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.requestConnectionDataHandler(nil, client, callID, cidTarget, pidTarget)
}
