// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// CreateMatchmakeSession sets the CreateMatchmakeSession handler function
func (protocol *Protocol) CreateMatchmakeSession(handler func(err error, client *nex.Client, callID uint32, anyGathering *nex.DataHolder, message string, participationCount uint16)) {
	protocol.createMatchmakeSessionHandler = handler
}

func (protocol *Protocol) handleCreateMatchmakeSession(packet nex.PacketInterface) {
	matchmakingVersion := protocol.Server.MatchMakingProtocolVersion()

	if protocol.createMatchmakeSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::CreateMatchmakeSession not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	anyGathering, err := parametersStream.ReadDataHolder()
	if err != nil {
		go protocol.createMatchmakeSessionHandler(fmt.Errorf("Failed to read anyGathering from parameters. %s", err.Error()), client, callID, nil, "", 0)
		return
	}

	message, err := parametersStream.ReadString()
	if err != nil {
		go protocol.createMatchmakeSessionHandler(fmt.Errorf("Failed to read message from parameters. %s", err.Error()), client, callID, nil, "", 0)
		return
	}

	var participationCount uint16 = 0

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 4 {
		participationCount, err = parametersStream.ReadUInt16LE()
		if err != nil {
			go protocol.createMatchmakeSessionHandler(fmt.Errorf("Failed to read message from participationCount. %s", err.Error()), client, callID, nil, "", 0)
			return
		}
	}

	go protocol.createMatchmakeSessionHandler(nil, client, callID, anyGathering, message, participationCount)
}
