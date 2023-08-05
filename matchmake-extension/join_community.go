// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// JoinCommunity sets the JoinCommunity handler function
func (protocol *Protocol) JoinCommunity(handler func(err error, client *nex.Client, callID uint32, gid uint32, strMessage string, strPassword string) uint32) {
	protocol.joinCommunityHandler = handler
}

func (protocol *Protocol) handleJoinCommunity(packet nex.PacketInterface) {
	var errorCode uint32

	if protocol.joinCommunityHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::JoinCommunity not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid, err := parametersStream.ReadUInt32LE()
	if err != nil {
		errorCode = protocol.joinCommunityHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, "", "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.joinCommunityHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, "", "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	strPassword, err := parametersStream.ReadString()
	if err != nil {
		errorCode = protocol.joinCommunityHandler(fmt.Errorf("Failed to read strPassword from parameters. %s", err.Error()), client, callID, 0, "", "")
		if errorCode != 0 {
			globals.RespondError(packet, ProtocolID, errorCode)
		}

		return
	}

	errorCode = protocol.joinCommunityHandler(nil, client, callID, gid, strMessage, strPassword)
	if errorCode != 0 {
		globals.RespondError(packet, ProtocolID, errorCode)
	}
}
