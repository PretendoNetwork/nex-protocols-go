// Package protocol implements the Matchmake tension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// JoinMatchmakeSession sets the JoinMatchmakeSession handler function
func (protocol *Protocol) JoinMatchmakeSession(handler func(err error, client *nex.Client, callID uint32, gid uint32, strMessage string)) {
	protocol.joinMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleJoinMatchmakeSession(packet nex.PacketInterface) {
	if protocol.joinMatchmakeSessionExHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::JoinMatchmakeSession not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		go protocol.joinMatchmakeSessionHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, "")
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.joinMatchmakeSessionHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, "")
		return
	}

	go protocol.joinMatchmakeSessionHandler(nil, client, callID, gid, strMessage)
}
