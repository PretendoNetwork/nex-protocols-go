// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// SetApplicationConfig sets the SetApplicationConfig handler function
func (protocol *Protocol) SetApplicationConfig(handler func(err error, packet nex.PacketInterface, callID uint32, applicationID uint32, key uint32, value int32) uint32) {
	protocol.setApplicationConfigHandler = handler
}

func (protocol *Protocol) handleSetApplicationConfig(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.setApplicationConfigHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SetApplicationConfig not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.setApplicationConfigHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	key, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.setApplicationConfigHandler(fmt.Errorf("Failed to read key from parameters. %s", err.Error()), packet, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	value, err := parametersStream.ReadInt32LE()
	if err != nil {
		errorCode = protocol.setApplicationConfigHandler(fmt.Errorf("Failed to read value from parameters. %s", err.Error()), packet, callID, 0, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.setApplicationConfigHandler(nil, packet, callID, applicationID, key, value)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
