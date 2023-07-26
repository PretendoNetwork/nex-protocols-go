// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// RemoveFriendByPrincipalID sets the RemoveFriendByPrincipalID handler function
func (protocol *Protocol) RemoveFriendByPrincipalID(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	protocol.removeFriendByPrincipalIDHandler = handler
}

func (protocol *Protocol) handleRemoveFriendByPrincipalID(packet nex.PacketInterface) {
	if protocol.removeFriendByPrincipalIDHandler == nil {
		globals.Logger.Warning("Friends3DS::RemoveFriendByPrincipalID not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.removeFriendByPrincipalIDHandler(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.removeFriendByPrincipalIDHandler(nil, client, callID, pid)
}
