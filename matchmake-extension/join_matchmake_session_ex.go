// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// JoinMatchmakeSessionEx sets the JoinMatchmakeSessionEx handler function
func (protocol *MatchmakeExtensionProtocol) JoinMatchmakeSessionEx(handler func(err error, client *nex.Client, callID uint32, gid uint32, strMessage string, dontCareMyBlockList bool, participationCount uint16)) {
	protocol.joinMatchmakeSessionExHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleJoinMatchmakeSessionEx(packet nex.PacketInterface) {
	if protocol.joinMatchmakeSessionExHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::JoinMatchmakeSessionEx not implemented")
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
		go protocol.joinMatchmakeSessionExHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, "", false, 0)
		return
	}

	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.joinMatchmakeSessionExHandler(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), client, callID, 0, "", false, 0)
		return
	}

	dontCareMyBlockList, err := parametersStream.ReadBool()
	if err != nil {
		go protocol.joinMatchmakeSessionExHandler(fmt.Errorf("Failed to read dontCareMyBlockList from parameters. %s", err.Error()), client, callID, 0, "", false, 0)
		return
	}

	participationCount, err := parametersStream.ReadUInt16LE()
	if err != nil {
		go protocol.joinMatchmakeSessionExHandler(fmt.Errorf("Failed to read participationCount from parameters. %s", err.Error()), client, callID, 0, "", false, 0)
		return
	}

	go protocol.joinMatchmakeSessionExHandler(nil, client, callID, gid, strMessage, dontCareMyBlockList, participationCount)
}
