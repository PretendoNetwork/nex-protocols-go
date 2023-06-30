// Package secure_connection implements the Secure Connection NEX protocol
package secure_connection

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// TestConnectivity sets the TestConnectivity handler function
func (protocol *SecureConnectionProtocol) TestConnectivity(handler func(err error, client *nex.Client, callID uint32)) {
	protocol.TestConnectivityHandler = handler
}

func (protocol *SecureConnectionProtocol) handleTestConnectivity(packet nex.PacketInterface) {
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
