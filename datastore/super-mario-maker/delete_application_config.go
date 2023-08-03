// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteApplicationConfig sets the DeleteApplicationConfig handler function
func (protocol *Protocol) DeleteApplicationConfig(handler func(err error, client *nex.Client, callID uint32, applicationID uint32, key uint32)) {
	protocol.deleteApplicationConfigHandler = handler
}

func (protocol *Protocol) handleDeleteApplicationConfig(packet nex.PacketInterface) {
	if protocol.deleteApplicationConfigHandler == nil {
		globals.Logger.Warning("DataStoreSuperMarioMaker::DeleteApplicationConfig not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	applicationID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.deleteApplicationConfigHandler(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	key, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.deleteApplicationConfigHandler(fmt.Errorf("Failed to read key from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.deleteApplicationConfigHandler(nil, client, callID, applicationID, key)
}
