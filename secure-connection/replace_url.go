// Package secure_connection implements the Secure Connection NEX protocol
package secure_connection

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReplaceURL sets the ReplaceURL handler function
func (protocol *SecureConnectionProtocol) ReplaceURL(handler func(err error, client *nex.Client, callID uint32, target *nex.StationURL, url *nex.StationURL)) {
	protocol.ReplaceURLHandler = handler
}

func (protocol *SecureConnectionProtocol) handleReplaceURL(packet nex.PacketInterface) {
	if protocol.ReplaceURLHandler == nil {
		globals.Logger.Warning("SecureConnection::ReplaceURL not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStationURL()
	if err != nil {
		go protocol.ReplaceURLHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	url, err := parametersStream.ReadStationURL()
	if err != nil {
		go protocol.ReplaceURLHandler(fmt.Errorf("Failed to read url from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.ReplaceURLHandler(nil, client, callID, target, url)
}
