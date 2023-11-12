// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetApplicationConfigString sets the GetApplicationConfigString handler function
func (protocol *Protocol) GetApplicationConfigString(handler func(err error, packet nex.PacketInterface, callID uint32, applicationID uint32) uint32) {
	protocol.getApplicationConfigStringHandler = handler
}

func (protocol *Protocol) handleGetApplicationConfigString(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getApplicationConfigStringHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::GetApplicationConfigString not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.getApplicationConfigStringHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getApplicationConfigStringHandler(nil, packet, callID, applicationID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
