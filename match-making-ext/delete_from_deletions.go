package match_making_ext

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// DeleteFromDeletions sets the DeleteFromDeletions handler function
func (protocol *MatchMakingExtProtocol) DeleteFromDeletions(handler func(err error, client *nex.Client, callID uint32, lstDeletions []uint32, pid uint32)) {
	protocol.DeleteFromDeletionsHandler = handler
}

func (protocol *MatchMakingExtProtocol) HandleDeleteFromDeletions(packet nex.PacketInterface) {
	if protocol.DeleteFromDeletionsHandler == nil {
		globals.Logger.Warning("MatchMakingExt::DeleteFromDeletions not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	lstDeletionsCount := parametersStream.ReadUInt32LE()
	lstDeletions := make([]uint32, lstDeletionsCount)
	for i := 0; uint32(i) < lstDeletionsCount; i++ {
		lstDeletions[i] = parametersStream.ReadUInt32LE()
	}

	pid := parametersStream.ReadUInt32LE()

	go protocol.DeleteFromDeletionsHandler(nil, client, callID, lstDeletions, pid)
}
