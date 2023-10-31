// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetTransactionParam sets the GetTransactionParam handler function
func (protocol *Protocol) GetTransactionParam(handler func(err error, packet nex.PacketInterface, callID uint32, slotID uint16) uint32) {
	protocol.getTransactionParamHandler = handler
}

func (protocol *Protocol) handleGetTransactionParam(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.getTransactionParamHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::GetTransactionParam not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	slotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		errorCode = protocol.getTransactionParamHandler(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.getTransactionParamHandler(nil, packet, callID, slotID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
