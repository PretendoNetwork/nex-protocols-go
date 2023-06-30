// Package secure_connection implements the Secure Connection NEX protocol
package secure_connection

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RequestURLs sets the RequestURLs handler function
func (protocol *SecureConnectionProtocol) RequestURLs(handler func(err error, client *nex.Client, callID uint32, cidTarget uint32, pidTarget uint32)) {
	protocol.RequestURLsHandler = handler
}

func (protocol *SecureConnectionProtocol) handleRequestURLs(packet nex.PacketInterface) {
	if protocol.RequestURLsHandler == nil {
		globals.Logger.Warning("SecureConnection::RequestURLs not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	cidTarget, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.RequestURLsHandler(fmt.Errorf("Failed to read cidTarget from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	pidTarget, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.RequestURLsHandler(fmt.Errorf("Failed to read pidTarget from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.RequestURLsHandler(nil, client, callID, cidTarget, pidTarget)
}
