package matchmake_extension

import (
	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
)

// JoinMatchmakeSessionEx sets the JoinMatchmakeSessionEx handler function
func (protocol *MatchmakeExtensionProtocol) JoinMatchmakeSessionEx(handler func(err error, client *nex.Client, callID uint32, gid uint32, strMessage string, dontCareMyBlockList bool, participationCount uint16)) {
	protocol.JoinMatchmakeSessionExHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleJoinMatchmakeSessionEx(packet nex.PacketInterface) {
	if protocol.JoinMatchmakeSessionExHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::JoinMatchmakeSessionEx not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	gid := parametersStream.ReadUInt32LE()
	strMessage, err := parametersStream.ReadString()
	if err != nil {
		go protocol.JoinMatchmakeSessionExHandler(err, client, callID, 0, "", false, 0)
		return
	}

	dontCareMyBlockList := parametersStream.ReadBool()
	participationCount := parametersStream.ReadUInt16LE()

	go protocol.JoinMatchmakeSessionExHandler(nil, client, callID, gid, strMessage, dontCareMyBlockList, participationCount)
}
