// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// ReplaceURL sets the ReplaceURL handler function
func (protocol *Protocol) ReplaceURL(handler func(err error, packet nex.PacketInterface, callID uint32, target *nex.StationURL, url *nex.StationURL) uint32) {
	protocol.replaceURLHandler = handler
}

func (protocol *Protocol) handleReplaceURL(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.replaceURLHandler == nil {
		globals.Logger.Warning("SecureConnection::ReplaceURL not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	target, err := parametersStream.ReadStationURL()
	if err != nil {
		errorCode = protocol.replaceURLHandler(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	url, err := parametersStream.ReadStationURL()
	if err != nil {
		errorCode = protocol.replaceURLHandler(fmt.Errorf("Failed to read url from parameters. %s", err.Error()), packet, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.replaceURLHandler(nil, packet, callID, target, url)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
