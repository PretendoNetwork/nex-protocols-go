// Package protocol implements the Super Smash Bros. 4 DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetApplicationConfig sets the GetApplicationConfig handler function
func (protocol *Protocol) GetApplicationConfig(handler func(err error, client *nex.Client, callID uint32, applicationID uint32) uint32) {
	protocol.getApplicationConfigHandler = handler
}

func (protocol *Protocol) handleGetApplicationConfig(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getApplicationConfigHandler == nil {
		globals.Logger.Warning("DataStoreSuperSmashBros4::GetApplicationConfig not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getApplicationConfigHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getApplicationConfigHandler(nil, client, callID, applicationID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
