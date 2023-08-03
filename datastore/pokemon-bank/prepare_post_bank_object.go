// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// PreparePostBankObject sets the PreparePostBankObject handler function
func (protocol *Protocol) PreparePostBankObject(handler func(err error, client *nex.Client, callID uint32, slotID uint16, size uint32)) {
	protocol.preparePostBankObjectHandler = handler
}

func (protocol *Protocol) handlePreparePostBankObject(packet nex.PacketInterface) {
	if protocol.preparePostBankObjectHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::PreparePostBankObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	slotID, err := parametersStream.ReadUInt16LE()
	if err != nil {
		go protocol.preparePostBankObjectHandler(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	size, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.preparePostBankObjectHandler(fmt.Errorf("Failed to read size from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.preparePostBankObjectHandler(nil, client, callID, slotID, size)
}
