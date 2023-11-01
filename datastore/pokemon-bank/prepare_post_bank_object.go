// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PreparePostBankObject sets the PreparePostBankObject handler function
func (protocol *Protocol) PreparePostBankObject(handler func(err error, packet nex.PacketInterface, callID uint32, slotID uint16, size uint32) uint32) {
	protocol.preparePostBankObjectHandler = handler
}

func (protocol *Protocol) handlePreparePostBankObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.preparePostBankObjectHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::PreparePostBankObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	slotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		errorCode = protocol.preparePostBankObjectHandler(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	size, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.preparePostBankObjectHandler(fmt.Errorf("Failed to read size from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.preparePostBankObjectHandler(nil, packet, callID, slotID, size)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
