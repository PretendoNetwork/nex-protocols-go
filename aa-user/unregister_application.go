// Package protocol implements the AAUser protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UnregisterApplication sets the UnregisterApplication handler function
func (protocol *Protocol) UnregisterApplication(handler func(err error, client *nex.Client, callID uint32, titleID uint64)) {
	protocol.unregisterApplicationHandler = handler
}

func (protocol *Protocol) handleUnregisterApplication(packet nex.PacketInterface) {
	if protocol.unregisterApplicationHandler == nil {
		globals.Logger.Warning("AAUser::UnregisterApplication not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	titleID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.unregisterApplicationHandler(fmt.Errorf("Failed to read titleID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.unregisterApplicationHandler(nil, client, callID, titleID)
}
