// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetUnlockKey sets the GetUnlockKey handler function
func (protocol *Protocol) GetUnlockKey(handler func(err error, client *nex.Client, callID uint32, challengeValue uint32) uint32) {
	protocol.getUnlockKeyHandler = handler
}

func (protocol *Protocol) handleGetUnlockKey(packet nex.PacketInterface) {
	if protocol.getUnlockKeyHandler == nil {
		globals.Logger.Warning("DataStorePokemonBank::GetUnlockKey not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	challengeValue, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.getUnlockKeyHandler(fmt.Errorf("Failed to read challengeValue from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getUnlockKeyHandler(nil, client, callID, challengeValue)
}
