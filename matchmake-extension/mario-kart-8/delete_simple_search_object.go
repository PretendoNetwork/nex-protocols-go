// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteSimpleSearchObject sets the DeleteSimpleSearchObject handler function
func (protocol *Protocol) DeleteSimpleSearchObject(handler func(err error, client *nex.Client, callID uint32, objectID uint32)) {
	protocol.deleteSimpleSearchObjectHandler = handler
}

func (protocol *Protocol) handleDeleteSimpleSearchObject(packet nex.PacketInterface) {
	if protocol.deleteSimpleSearchObjectHandler == nil {
		globals.Logger.Warning("MatchmakeExtensionMarioKart8::DeleteSimpleSearchObject not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	objectID, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.deleteSimpleSearchObjectHandler(fmt.Errorf("Failed to read objectID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.deleteSimpleSearchObjectHandler(nil, client, callID, objectID)
}
