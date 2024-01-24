// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSetApplicationConfig(packet nex.PacketInterface) {
	var err error
	var errorCode uint32

	if protocol.SetApplicationConfig == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::SetApplicationConfig not implemented")
		globals.RespondError(packet, ProtocolID, nex.ResultCodes.Core.NotImplemented)
		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	applicationID := types.NewPrimitiveU32(0)
	err = applicationID.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SetApplicationConfig(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	key := types.NewPrimitiveU32(0)
	err = key.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SetApplicationConfig(fmt.Errorf("Failed to read key from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	value := types.NewPrimitiveS32(0)
	err = value.ExtractFrom(parametersStream)
	if err != nil {
		_, errorCode = protocol.SetApplicationConfig(fmt.Errorf("Failed to read value from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	rmcMessage, errorCode := protocol.SetApplicationConfig(nil, packet, callID, applicationID, key, value)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
		return
	}

	globals.Respond(packet, rmcMessage)
}
