// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveFromBlockList sets the RemoveFromBlockList handler function
func (protocol *Protocol) RemoveFromBlockList(handler func(err error, client *nex.Client, callID uint32, lstPrincipalID []uint32)) {
	protocol.removeFromBlockListHandler = handler
}

func (protocol *Protocol) handleRemoveFromBlockList(packet nex.PacketInterface) {
	if protocol.removeFromBlockListHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::RemoveFromBlockList not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstPrincipalID, err := parametersStream.ReadListUInt32LE()
	if err != nil {
		go protocol.removeFromBlockListHandler(fmt.Errorf("Failed to read lstPrincipalID from parameters. %s", err.Error()), client, callID, nil)
		return
	}

	go protocol.removeFromBlockListHandler(nil, client, callID, lstPrincipalID)
}
