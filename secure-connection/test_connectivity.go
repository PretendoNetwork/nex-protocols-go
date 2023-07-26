// Package protocol implements the Secure Connection protocol
package protocol

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// TestConnectivity sets the TestConnectivity handler function
func (protocol *Protocol) TestConnectivity(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.TestConnectivityHandler = handler
}

func (protocol *Protocol) handleTestConnectivity(packet nex.PacketInterface) {
	if protocol.TestConnectivityHandler == nil {
		globals.Logger.Warning("SecureConnection::TestConnectivity not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go protocol.TestConnectivityHandler(nil, client, callID)
}
