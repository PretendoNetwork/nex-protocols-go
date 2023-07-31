// Package protocol implements the Ranking protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetCommonData sets the GetCommonData handler function
func (protocol *Protocol) GetCommonData(handler func(err error, client *nex.Client, callID uint32, uniqueID uint64)) {
	protocol.getCommonDataHandler = handler
}

func (protocol *Protocol) handleGetCommonData(packet nex.PacketInterface) {
	if protocol.getCommonDataHandler == nil {
		globals.Logger.Warning("Ranking::GetCommonData not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	uniqueID, err := parametersStream.ReadUInt64LE()
	if err != nil {
		go protocol.getCommonDataHandler(fmt.Errorf("Failed to read uniqueID from parameters. %s", err.Error()), client, callID, 0)
		return
	}

	go protocol.getCommonDataHandler(nil, client, callID, uniqueID)
}
