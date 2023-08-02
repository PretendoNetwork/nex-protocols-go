// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteCommonData sets the DeleteCommonData handler function
func (protocol *Protocol) DeleteCommonData(handler func(err error, client *nex.Client, callID uint32, nexUniqueID uint64)) {
	protocol.deleteCommonDataHandler = handler
}

func (protocol *Protocol) handleDeleteCommonData(packet nex.PacketInterface) {
	if protocol.deleteCommonDataHandler == nil {
		globals.Logger.Warning("Ranking2::DeleteCommonData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}
	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	nexUniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.deleteCommonDataHandler(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.deleteCommonDataHandler(nil, client, callID, nexUniqueID)
}