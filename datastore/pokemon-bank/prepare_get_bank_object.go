// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PrepareGetBankObject sets the PrepareGetBankObject handler function
func (protocol *Protocol) PrepareGetBankObject(handler func(err error, client *nex.Client, callID uint32, slotID uint16, applicationID uint16) uint32) {
	protocol.prepareGetBankObjectHandler = handler
}

func (protocol *Protocol) handlePrepareGetBankObject(packet nex.PacketInterface) {
	if protocol.prepareGetBankObjectHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::PrepareGetBankObject not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	slotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		go protocol.prepareGetBankObjectHandler(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	applicationID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		go protocol.prepareGetBankObjectHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.prepareGetBankObjectHandler(nil, client, callID, slotID, applicationID)
}
