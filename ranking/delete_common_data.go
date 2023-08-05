// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteCommonData sets the DeleteCommonData handler function
func (protocol *Protocol) DeleteCommonData(handler func(err error, client *nex.Client, callID uint32, uniqueID uint64) uint32) {
	protocol.deleteCommonDataHandler = handler
}

func (protocol *Protocol) handleDeleteCommonData(packet nex.PacketInterface) {
	if protocol.deleteCommonDataHandler == nil {
		globals.Logger.Warning("Ranking::DeleteCommonData not implemented")
		go globals.RespondError(packet, ProtocolID, nex.Errors.Core.NotImplemented)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.deleteCommonDataHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.deleteCommonDataHandler(nil, client, callID, uniqueID)
}
