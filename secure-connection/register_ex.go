// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RegisterEx sets the RegisterEx handler function
func (protocol *Protocol) RegisterEx(handler func(err error, client *nex.Client, callID uint32, vecMyURLs []*nex.StationURL, hCustomData *nex.DataHolder) uint32) {
	protocol.registerExHandler = handler
}

func (protocol *Protocol) handleRegisterEx(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.registerExHandler == nil {
		globals.Logger.Warning("SecureConnection::RegisterEx not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	vecMyURLs, err := parametersStream.ReadListStationURL()
	if err != nil {
		errorCode = protocol.registerExHandler(fmt.Errorf("Failed to read vecMyURLs from parameters. %s", err.Error()), client, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	hCustomData, err := parametersStream.ReadDataHolder()
	if err != nil {
		errorCode = protocol.registerExHandler(fmt.Errorf("Failed to read hCustomData from parameters. %s", err.Error()), client, callID, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.registerExHandler(nil, client, callID, vecMyURLs, hCustomData)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
