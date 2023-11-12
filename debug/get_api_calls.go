// Package protocol implements the Debug protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetAPICalls sets the GetAPICalls handler function
func (protocol *Protocol) GetAPICalls(handler func(err error, packet nex.PacketInterface, callID uint32, pids []uint32, unknown *nex.DateTime, unknown2 *nex.DateTime) uint32) {
	protocol.getAPICallsHandler = handler
}

func (protocol *Protocol) handleGetAPICalls(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getAPICallsHandler == nil {
		globals.Logger.Warning("Debug::GetAPICalls not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pids, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		errorCode = protocol.getAPICallsHandler(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown, err := parametersStream.ReadDateTime()
	if err != nil {
		errorCode = protocol.getAPICallsHandler(fmt.Errorf("Failed to read unknown from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown2, err := parametersStream.ReadDateTime()
	if err != nil {
		errorCode = protocol.getAPICallsHandler(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getAPICallsHandler(nil, packet, callID, pids, unknown, unknown2)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
