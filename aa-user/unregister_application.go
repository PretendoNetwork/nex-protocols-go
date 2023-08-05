// Package protocol implements the AAUser protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnregisterApplication sets the UnregisterApplication handler function
func (protocol *Protocol) UnregisterApplication(handler func(err error, client *nex.Client, callID uint32, titleID uint64) uint32) {
	protocol.unregisterApplicationHandler = handler
}

func (protocol *Protocol) handleUnregisterApplication(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.unregisterApplicationHandler == nil {
		globals.Logger.Warning("AAUser::UnregisterApplication not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	titleID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		errorCode = protocol.unregisterApplicationHandler(fmt.Errorf("Failed to read titleID from parameters. %s", err.Error()), client, callID, 0)
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.unregisterApplicationHandler(nil, client, callID, titleID)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
