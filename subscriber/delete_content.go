// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteContent sets the DeleteContent handler function
func (protocol *Protocol) DeleteContent(handler func(err error, client *nex.Client, callID uint32, unknown1 []string, unknown2 uint64) uint32) {
	protocol.deleteContentHandler = handler
}

func (protocol *Protocol) handleDeleteContent(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.deleteContentHandler == nil {
		globals.Logger.Warning("Subscriber::DeleteContent not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	unknown1, err := parametersStream.ReadListString()
	if err != nil {
		errorCode = protocol.deleteContentHandler(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), client, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	unknown2, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.deleteContentHandler(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), client, callID, nil, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.deleteContentHandler(nil, client, callID, unknown1, unknown2)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
