// Package secure_connection implements the Secure Connection NEX protocol
package secure_connection

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// Register sets the Register handler function
func (protocol *SecureConnectionProtocol) Register(handler func(err error, client *nex.Client, callID uint32, vecMyURLs []*nex.StationURL)) {
	protocol.RegisterHandler = handler
}

func (protocol *SecureConnectionProtocol) handleRegister(packet nex.PacketInterface) {
	if protocol.RegisterHandler == nil {
		globals.Logger.Warning("SecureConnection::Register not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	vecMyURLs, err := parametersStream.ReadListStationURL()
	if err != nil {
		go protocol.RegisterHandler(fmt.Errorf("Failed to read hCustomData from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.RegisterHandler(nil, client, callID, vecMyURLs)
}
