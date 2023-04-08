package debug

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// GetApiCalls sets the GetApiCalls handler function
func (protocol *DebugProtocol) GetApiCalls(handler func(err error, client *nex.Client, callID uint32, pids []uint32, dateUnk1 uint64, dateUnk2 uint64)) {
	protocol.GetApiCallsHandler = handler
}

func (protocol *DebugProtocol) HandleGetApiCalls(packet nex.PacketInterface) {
	if protocol.GetApiCallsHandler == nil {
		globals.Logger.Warning("Debug::GetApiCalls not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	pidsCount := parametersStream.ReadUInt32LE()
	pids := make([]uint32, pidsCount)
	for i := 0; uint32(i) < pidsCount; i++ {
		pids[i] = parametersStream.ReadUInt32LE()
	}

	dateUnk1 := parametersStream.ReadUInt64LE()

	dateUnk2 := parametersStream.ReadUInt64LE()

	go protocol.GetApiCallsHandler(nil, client, callID, pids, dateUnk1, dateUnk2)
}
