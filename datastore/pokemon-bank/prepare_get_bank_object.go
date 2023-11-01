// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareGetBankObject sets the PrepareGetBankObject handler function
func (protocol *Protocol) PrepareGetBankObject(handler func(err error, packet nex.PacketInterface, callID uint32, slotID uint16, applicationID uint16) uint32) {
	protocol.prepareGetBankObjectHandler = handler
}

func (protocol *Protocol) handlePrepareGetBankObject(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.prepareGetBankObjectHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::PrepareGetBankObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	slotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		errorCode = protocol.prepareGetBankObjectHandler(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	applicationID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		errorCode = protocol.prepareGetBankObjectHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, 0, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.prepareGetBankObjectHandler(nil, packet, callID, slotID, applicationID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
