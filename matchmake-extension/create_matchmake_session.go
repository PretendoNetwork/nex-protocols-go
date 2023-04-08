package matchmake_extension

import (
	"errors"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/nex-protocols-go/globals"
	match_making "github.com/PretendoNetwork/nex-protocols-go/match-making"
)

// CreateMatchmakeSession sets the CreateMatchmakeSession handler function
func (protocol *MatchmakeExtensionProtocol) CreateMatchmakeSession(handler func(err error, client *nex.Client, callID uint32, matchmakeSession *match_making.MatchmakeSession, message string, participationCount uint16)) {
	protocol.CreateMatchmakeSessionHandler = handler
}

func (protocol *MatchmakeExtensionProtocol) HandleCreateMatchmakeSession(packet nex.PacketInterface) {
	matchmakingVersion := protocol.Server.MatchMakingProtocolVersion()

	if protocol.CreateMatchmakeSessionHandler == nil {
		globals.Logger.Warning("MatchmakeExtension::CreateMatchmakeSession not implemented")
		go globals.RespondNotImplemented(packet, ProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, protocol.Server)

	dataHolderType, err := parametersStream.ReadString()

	if err != nil {
		go protocol.CreateMatchmakeSessionHandler(err, client, callID, nil, "", 0)
		return
	}

	if dataHolderType != "MatchmakeSession" {
		err := errors.New("[MatchmakeExtension::CreateMatchmakeSession] Data holder name does not match")
		go protocol.CreateMatchmakeSessionHandler(err, client, callID, nil, "", 0)
		return
	}

	if (parametersStream.ByteCapacity() - parametersStream.ByteOffset()) < 8 {
		err := errors.New("[MatchmakeExtension::CreateMatchmakeSession] Data holder missing lengths")
		go protocol.CreateMatchmakeSessionHandler(err, client, callID, nil, "", 0)
		return
	}

	parametersStream.SeekByte(4, true) // Skip length including next buffer length field
	dataHolderContent, err := parametersStream.ReadBuffer()

	if err != nil {
		go protocol.CreateMatchmakeSessionHandler(err, client, callID, nil, "", 0)
		return
	}

	dataHolderContentStream := nex.NewStreamIn(dataHolderContent, protocol.Server)

	matchmakeSession, err := dataHolderContentStream.ReadStructure(match_making.NewMatchmakeSession())

	if err != nil {
		go protocol.CreateMatchmakeSessionHandler(err, client, callID, nil, "", 0)
		return
	}

	message, err := dataHolderContentStream.ReadString()

	if err != nil {
		go protocol.CreateMatchmakeSessionHandler(err, client, callID, nil, "", 0)
		return
	}

	var participationCount uint16 = 0

	if matchmakingVersion.Major >= 3 && matchmakingVersion.Minor >= 5 {
		participationCount = dataHolderContentStream.ReadUInt16LE()
	}

	go protocol.CreateMatchmakeSessionHandler(nil, client, callID, matchmakeSession.(*match_making.MatchmakeSession), message, participationCount)
}
