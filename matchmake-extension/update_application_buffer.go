// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// UpdateApplicationBuffer sets the UpdateApplicationBuffer handler function
func (protocol *Protocol) UpdateApplicationBuffer(handler func(err error, client *nex.Client, callID uint32, gid uint32, applicationBuffer []byte)) {
	protocol.updateApplicationBufferHandler = handler
}

func (protocol *Protocol) handleUpdateApplicationBuffer(packet nex.PacketInterface) {
	if protocol.updateApplicationBufferHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::UpdateApplicationBuffer not implemented")
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
		go protocol.updateApplicationBufferHandler(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	applicationBuffer, err := parametersStream.ReadBuffer()
	if err != nil {
		go protocol.updateApplicationBufferHandler(fmt.Errorf("Failed to read applicationBuffer from parameters. %s", err.Error()), client, callID, 0, nil)
		return
	}

	go protocol.updateApplicationBufferHandler(nil, client, callID, gid, applicationBuffer)
}
