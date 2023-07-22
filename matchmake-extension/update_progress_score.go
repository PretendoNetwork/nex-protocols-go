// Package matchmake_extension implements the Matchmake Extension NEX protocol
package matchmake_extension

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateProgressScore sets the UpdateProgressScore handler function
func (protocol *MatchmakeExtensionProtocol) UpdateProgressScore(handler func(err error, client *nex.Client, callID uint32, gid uint32, progressScore uint8)) {
	protocol.updateProgressScoreHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) handleUpdateProgressScore(packet nex.PacketInterface) {
	if protocol.updateProgressScoreHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateProgressScore not implemented")
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
		go protocol.updateProgressScoreHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	progressScore, err := parametersStream.ReadUInt8()
	if err != nil {
		go protocol.updateProgressScoreHandler(fmt.Errorf("Failed to read progressScore from parameters. %s", err.Error()), client, callID, 0, 0)
		return
	}

	go protocol.updateProgressScoreHandler(nil, client, callID, gid, progressScore)
}
