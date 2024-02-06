// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-go/types"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

func (protocol *Protocol) handleSetApplicationConfigString(packet nex.PacketInterface) {
	var err error

	if protocol.SetApplicationConfigString == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::SetApplicationConfigString not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()

	callID := request.CallID
	parameters := request.Parameters

	parametersStream := nex.NewByteStreamIn(parameters, protocol.server)

	applicationID := types.NewPrimitiveU32(0)
	err = applicationID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetApplicationConfigString(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	key := types.NewPrimitiveU32(0)
	err = key.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetApplicationConfigString(fmt.Errorf("Failed to read key from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	value := types.NewString("")
	err = value.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetApplicationConfigString(fmt.Errorf("Failed to read value from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SetApplicationConfigString(nil, packet, callID, applicationID, key, value)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
