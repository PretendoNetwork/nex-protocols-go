// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReplaceURL sets the ReplaceURL handler function
func (protocol *Protocol) ReplaceURL(handler func(err error, client *nex.Client, callID uint32, target *nex.StationURL, url *nex.StationURL) uint32) {
	protocol.replaceURLHandler = handler
}

func (protocol *Protocol) handleReplaceURL(packet nex.PacketInterface) {
	if protocol.replaceURLHandler == nil {
		globals.Logger.Warning("SecureConnection::ReplaceURL not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStationURL()
	if err != nil {
		go protocol.replaceURLHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	url, err := parametersStream.ReadStationURL()
	if err != nil {
		go protocol.replaceURLHandler(fmt.Errorf("Failed to read url from parameters. %s", err.Error()), client, callID, nil, nil)
		return
	}

	go protocol.replaceURLHandler(nil, client, callID, target, url)
}
