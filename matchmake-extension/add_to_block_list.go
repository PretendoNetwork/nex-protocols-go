// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// AddToBlockList sets the AddToBlockList handler function
func (protocol *Protocol) AddToBlockList(handler func(err error, client *nex.Client, callID uint32, lstPrincipalID []uint32) uint32) {
	protocol.addToBlockListHandler = handler
}

func (protocol *Protocol) handleAddToBlockList(packet nex.PacketInterface) {
	if protocol.addToBlockListHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::AddToBlockList not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstPrincipalID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.addToBlockListHandler(fmt.Errorf("Failed to read lstPrincipalID from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.addToBlockListHandler(nil, client, callID, lstPrincipalID)
}
