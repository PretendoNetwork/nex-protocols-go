package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// Friends3DSProtocolID is the protocol ID for the Friends (3DS) protocol
	Friends3DSProtocolID = 0x65

	// Friends3DSMethodUpdateMii is the method ID for method UpdateMii
	Friends3DSMethodUpdateMii = 0x2

	// Friends3DSMethodGetFriendMii is the method ID for method GetFriendMii
	Friends3DSMethodGetFriendMii = 0x6

	// Friends3DSMethodAddFriendByPrincipalID is the method ID for method AddFriendByPrincipalID
	Friends3DSMethodAddFriendByPrincipalID = 0xb

	// Friends3DSMethodDeleteFriendByPrincipalID is the method ID for method DeleteFriendByPrincipalID
	Friends3DSMethodDeleteFriendByPrincipalID = 0xb

	// Friends3DSMethodSyncFriend is the method ID for method SyncFriend
	Friends3DSMethodSyncFriend = 0x11

	// Friends3DSMethodUpdatePresence is the method ID for method UpdatePresence
	Friends3DSMethodUpdatePresence = 0x12

	// Friends3DSMethodUpdateFavoriteGame is the method ID for method UpdateFavoriteGame
	Friends3DSMethodUpdateFavoriteGame = 0x13

	// Friends3DSMethodUpdateMessage is the method ID for method UpdateMessage
	Friends3DSMethodUpdateMessage = 0x14

	// Friends3DSMethodGetFriendPresence is the method ID for method GetFriendPresence
	Friends3DSMethodGetFriendPresence = 0x16

	// Friends3DSMethodGetFriendPersistentInfo is the method ID for method GetFriendPersistentInfo
	Friends3DSMethodGetFriendPersistentInfo = 0x19
)

// Friends3DSProtocol handles the Friends (3DS) nex protocol
type Friends3DSProtocol struct {
	server                           *nex.Server
	UpdatePresenceHandler            func(err error, client *nex.Client, callID uint32, presence *NintendoPresence)
	SyncFriendHandler                func(err error, client *nex.Client, callID uint32, unknown1 uint64, unknown2 []uint32, unknown3 []uint64)
	AddFriendByPrincipalIDHandler    func(err error, client *nex.Client, callID uint32, unknown1 uint64, principalID uint32)
	GetFriendPersistentInfoHandler   func(err error, client *nex.Client, callID uint32, pidList []uint32)
	GetFriendMiiHandler              func(err error, client *nex.Client, callID uint32, pidList []uint32)
	GetFriendPresenceHandler         func(err error, client *nex.Client, callID uint32, pidList []uint32)
	UpdateMiiHandler                 func(err error, client *nex.Client, callID uint32)
	DeleteFriendByPrincipalIDHandler func(err error, client *nex.Client, callID uint32)
	UpdateFavoriteGameHandler        func(err error, client *nex.Client, callID uint32)
	UpdateMessageHandler             func(err error, client *nex.Client, callID uint32)
}

// NintendoPresence contains information about a users online presence
type NintendoPresence struct {
	ChangedFlags      uint32
	GameKey           *GameKey
	Message           string
	JoinAvailableFlag uint32
	MatchmakeType     uint8
	JoinGameID        uint32
	JoinGameMode      uint32
	OwnerPrincipalID  uint32
	JoinGroupID       uint32
	ApplicationArg    []byte

	nex.Structure
}

// Bytes encodes the NintendoPresence and returns a byte array
func (presence *NintendoPresence) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(presence.ChangedFlags)
	stream.WriteStructure(presence.GameKey)
	stream.WriteString(presence.Message)
	stream.WriteUInt32LE(presence.JoinAvailableFlag)
	stream.WriteUInt8(presence.MatchmakeType)
	stream.WriteUInt32LE(presence.JoinGameID)
	stream.WriteUInt32LE(presence.JoinGameMode)
	stream.WriteUInt32LE(presence.OwnerPrincipalID)
	stream.WriteUInt32LE(presence.JoinGroupID)
	stream.WriteBuffer(presence.ApplicationArg)

	return stream.Bytes()
}

// ExtractFromStream extracts a NintendoPresence structure from a stream
func (presence *NintendoPresence) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 25 {
		// length check for the following fixed-size data
		// changedFlags + JoinAvailableFlag + MatchmakeType + JoinGameID + JoinGameMode + OwnerPrincipalID + JoinGroupID
		return errors.New("[NintendoPresence::ExtractFromStream] Data size too small")
	}

	changedFlags := stream.ReadUInt32LE()
	gameKeyStructureInterface, err := stream.ReadStructure(NewGameKey())
	if err != nil {
		return err
	}
	gameKey := gameKeyStructureInterface.(*GameKey)
	message, err := stream.ReadString()
	if err != nil {
		return err
	}
	JoinAvailableFlag := stream.ReadUInt32LE()
	MatchmakeType := stream.ReadUInt8()
	JoinGameID := stream.ReadUInt32LE()
	JoinGameMode := stream.ReadUInt32LE()
	OwnerPrincipalID := stream.ReadUInt32LE()
	JoinGroupID := stream.ReadUInt32LE()
	ApplicationArg, err := stream.ReadBuffer()
	if err != nil {
		return err
	}

	presence.ChangedFlags = changedFlags
	presence.GameKey = gameKey
	presence.Message = message
	presence.JoinAvailableFlag = JoinAvailableFlag
	presence.MatchmakeType = MatchmakeType
	presence.JoinGameID = JoinGameID
	presence.JoinGameMode = JoinGameMode
	presence.OwnerPrincipalID = OwnerPrincipalID
	presence.JoinGroupID = JoinGroupID
	presence.ApplicationArg = ApplicationArg

	return nil
}

// NewNintendoPresence returns a new NintendoPresence
func NewNintendoPresence() *NintendoPresence {
	return &NintendoPresence{}
}

// FriendRelationship contains information about a users relationship with another PID
type FriendRelationship struct {
	PrincipalID  uint32
	Unknown1     uint64
	RelationType uint8 // guess

	nex.Structure
}

// Bytes encodes the FriendRelationship and returns a byte array
func (relationship *FriendRelationship) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(relationship.PrincipalID)
	stream.WriteUInt64LE(relationship.Unknown1)
	stream.WriteUInt8(relationship.RelationType)

	return stream.Bytes()
}

// NewFriendRelationship returns a new FriendRelationship
func NewFriendRelationship() *FriendRelationship {
	return &FriendRelationship{}
}

// FriendPersistentInfo contains user settings
type FriendPersistentInfo struct {
	PrincipalID  uint32
	Region       uint8
	Country      uint8
	Area         uint8
	Language     uint8
	Platform     uint8
	GameKey      *GameKey
	Message      string
	MsgUpdatedAt *nex.DateTime //appears to be correct, but not 100% sure.
	FriendedAt   *nex.DateTime
	DateTime3    *nex.DateTime

	nex.Structure
}

// Bytes encodes the FriendPersistentInfo and returns a byte array
func (friendPersistentInfo *FriendPersistentInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(friendPersistentInfo.PrincipalID)
	stream.WriteUInt8(friendPersistentInfo.Region)
	stream.WriteUInt8(friendPersistentInfo.Country)
	stream.WriteUInt8(friendPersistentInfo.Area)
	stream.WriteUInt8(friendPersistentInfo.Language)
	stream.WriteUInt8(friendPersistentInfo.Platform)
	stream.WriteStructure(friendPersistentInfo.GameKey)
	stream.WriteString(friendPersistentInfo.Message)
	stream.WriteUInt64LE(friendPersistentInfo.MsgUpdatedAt.Value())
	stream.WriteUInt64LE(friendPersistentInfo.FriendedAt.Value())
	stream.WriteUInt64LE(friendPersistentInfo.DateTime3.Value())

	return stream.Bytes()
}

// NewFriendPersistentInfo returns a new FriendPersistentInfo
func NewFriendPersistentInfo() *FriendPersistentInfo {
	return &FriendPersistentInfo{}
}

// Setup initializes the protocol
func (friends3DSProtocol *Friends3DSProtocol) Setup() {
	nexServer := friends3DSProtocol.server

	nexServer.On("Data", func(packet nex.PacketInterface) {
		request := packet.RMCRequest()

		if Friends3DSProtocolID == request.ProtocolID() {
			switch request.MethodID() {
			case Friends3DSMethodSyncFriend:
				go friends3DSProtocol.handleSyncFriend(packet)
			case Friends3DSMethodUpdatePresence:
				go friends3DSProtocol.handleUpdatePresence(packet)
			case Friends3DSMethodAddFriendByPrincipalID:
				go friends3DSProtocol.handleAddFriendByPrincipalID(packet)
			case Friends3DSMethodGetFriendPersistentInfo:
				go friends3DSProtocol.handleGetFriendPersistentInfo(packet)
			case Friends3DSMethodGetFriendMii:
				go friends3DSProtocol.handleGetFriendMii(packet)
			case Friends3DSMethodGetFriendPresence:
				go friends3DSProtocol.handleGetFriendPresence(packet)
			case Friends3DSMethodUpdateMii:
				go friends3DSProtocol.handleUpdateMii(packet)
			case Friends3DSMethodDeleteFriendByPrincipalID:
				go friends3DSProtocol.handleDeleteFriendByPrincipalID(packet)
			case Friends3DSMethodUpdateFavoriteGame:
				go friends3DSProtocol.handleUpdateFavoriteGame(packet)
			case Friends3DSMethodUpdateMessage:
				go friends3DSProtocol.handleUpdateMessage(packet)
			default:
				fmt.Printf("Unsupported Friends (3DS) method ID: %#v\n", request.MethodID())
			}
		}
	})
}

func (friends3DSProtocol *Friends3DSProtocol) handleSyncFriend(packet nex.PacketInterface) {
	if friends3DSProtocol.SyncFriendHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::SyncFriend not implemented")
		go respondNotImplemented(packet, Friends3DSProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	Unknown1 := parametersStream.ReadUInt64LE()
	Unknown2 := parametersStream.ReadListUInt32LE()
	Unknown3 := parametersStream.ReadListUInt64LE()

	go friends3DSProtocol.SyncFriendHandler(nil, client, callID, Unknown1, Unknown2, Unknown3)
}

func (friends3DSProtocol *Friends3DSProtocol) handleGetFriendPersistentInfo(packet nex.PacketInterface) {
	if friends3DSProtocol.GetFriendPersistentInfoHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::GetFriendPersistentInfo not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	PidList := parametersStream.ReadListUInt32LE()

	go friends3DSProtocol.GetFriendPersistentInfoHandler(nil, client, callID, PidList)
}

func (friends3DSProtocol *Friends3DSProtocol) handleUpdatePresence(packet nex.PacketInterface) {
	if friends3DSProtocol.UpdatePresenceHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::UpdatePresence not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	nintendoPresenceStructureInterface, err := parametersStream.ReadStructure(NewNintendoPresence())
	if err != nil {
		go friends3DSProtocol.UpdatePresenceHandler(err, client, callID, nil)
		return
	}

	nintendoPresence := nintendoPresenceStructureInterface.(*NintendoPresence)

	go friends3DSProtocol.UpdatePresenceHandler(nil, client, callID, nintendoPresence)
}

func (friends3DSProtocol *Friends3DSProtocol) handleAddFriendByPrincipalID(packet nex.PacketInterface) {
	if friends3DSProtocol.AddFriendByPrincipalIDHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::AddFriendByPrincipalID not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	Unknown1 := parametersStream.ReadUInt64LE()
	PrincipalID := parametersStream.ReadUInt32LE()

	go friends3DSProtocol.AddFriendByPrincipalIDHandler(nil, client, callID, Unknown1, PrincipalID)
}

func (friends3DSProtocol *Friends3DSProtocol) handleGetFriendMii(packet nex.PacketInterface) {
	if friends3DSProtocol.GetFriendMiiHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::GetFriendMiiHandler not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	PidList := parametersStream.ReadListUInt32LE()

	go friends3DSProtocol.GetFriendMiiHandler(nil, client, callID, PidList)
}

func (friends3DSProtocol *Friends3DSProtocol) handleGetFriendPresence(packet nex.PacketInterface) {
	if friends3DSProtocol.GetFriendPresenceHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::GetFriendPresenceHandler not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	PidList := parametersStream.ReadListUInt32LE()

	go friends3DSProtocol.GetFriendPresenceHandler(nil, client, callID, PidList)
}

func (friends3DSProtocol *Friends3DSProtocol) handleUpdateMii(packet nex.PacketInterface) {
	if friends3DSProtocol.UpdateMiiHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::UpdateMiiHandler not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	go friends3DSProtocol.UpdateMiiHandler(nil, client, callID)
}

func (friends3DSProtocol *Friends3DSProtocol) handleDeleteFriendByPrincipalID(packet nex.PacketInterface) {
	if friends3DSProtocol.DeleteFriendByPrincipalIDHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::DeleteFriendByPrincipalIDHandler not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	principalID := parametersStream.ReadUInt32LE()

	go friends3DSProtocol.DeleteFriendByPrincipalIDHandler(nil, client, callID, principalID)
}

func (friends3DSProtocol *Friends3DSProtocol) handleUpdateFavoriteGame(packet nex.PacketInterface) {
	if friends3DSProtocol.UpdateFavoriteGameHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::UpdateFavoriteGameHandler not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	gameKeyStructureInterface, err := stream.ReadStructure(NewGameKey())

	if err != nil {
		return err
	}

	gameKey := gameKeyStructureInterface.(*GameKey)

	go friends3DSProtocol.UpdateFavoriteGameHandler(nil, client, callID, gameKey)
}

func (friends3DSProtocol *Friends3DSProtocol) handleUpdateMessage(packet nex.PacketInterface) {
	if friends3DSProtocol.UpdateMessageHandler == nil {
		fmt.Println("[Warning] Friends3DSProtocol::UpdateMessageHandler not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	message, err := stream.ReadString()

	if err != nil {
		return err
	}

	go friends3DSProtocol.UpdateMessageHandler(nil, client, callID, message)
}

// AddFriendByPrincipalID sets the AddFriendByPrincipalID handler function
func (friends3DSProtocol *Friends3DSProtocol) AddFriendByPrincipalID(handler func(err error, client *nex.Client, callID uint32, unknown1 uint64, principalID uint32)) {
	friends3DSProtocol.AddFriendByPrincipalIDHandler = handler
}

// SyncFriend sets the SyncFriend handler function
func (friends3DSProtocol *Friends3DSProtocol) SyncFriend(handler func(err error, client *nex.Client, callID uint32, unknown1 uint64, unknown2 []uint32, unknown3 []uint64)) {
	friends3DSProtocol.SyncFriendHandler = handler
}

// UpdatePresence sets the UpdatePresence handler function
func (friends3DSProtocol *Friends3DSProtocol) UpdatePresence(handler func(err error, client *nex.Client, callID uint32, presence *NintendoPresence)) {
	friends3DSProtocol.UpdatePresenceHandler = handler
}

// GetFriendPersistentInfo sets the GetFriendPersistentInfo handler function
func (friends3DSProtocol *Friends3DSProtocol) GetFriendPersistentInfo(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	friends3DSProtocol.GetFriendPersistentInfoHandler = handler
}

// GetFriendMii sets the GetFriendMii handler function
func (friends3DSProtocol *Friends3DSProtocol) GetFriendMii(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	friends3DSProtocol.GetFriendMiiHandler = handler
}

// GetFriendPresence sets the GetFriendPresence handler function
func (friends3DSProtocol *Friends3DSProtocol) GetFriendPresence(handler func(err error, client *nex.Client, callID uint32, pidList []uint32)) {
	friends3DSProtocol.GetFriendPresenceHandler = handler
}

// UpdateMii sets the UpdateMii handler function
func (friends3DSProtocol *Friends3DSProtocol) UpdateMii(handler func(err error, client *nex.Client, callID uint32)) {
	friends3DSProtocol.UpdateMiiHandler = handler
}

// DeleteFriendByPrincipalID sets the DeleteFriendByPrincipalID handler function
func (friends3DSProtocol *Friends3DSProtocol) DeleteFriendByPrincipalID(handler func(err error, client *nex.Client, callID uint32, principalID uint32)) {
	friends3DSProtocol.DeleteFriendByPrincipalIDHandler = handler
}

// UpdateFavoriteGame sets the UpdateFavoriteGame handler function
func (friends3DSProtocol *Friends3DSProtocol) UpdateFavoriteGame(handler func(err error, client *nex.Client, callID uint32, gameKey *GameKey)) {
	friends3DSProtocol.UpdateFavoriteGameHandler = handler
}

// UpdateMessage sets the UpdateMessage handler function
func (friends3DSProtocol *Friends3DSProtocol) UpdateMessage(handler func(err error, client *nex.Client, callID uint32, message string)) {
	friends3DSProtocol.UpdateMessageHandler = handler
}

// NewFriends3DSProtocol returns a new Friends3DSProtocol
func NewFriends3DSProtocol(server *nex.Server) *Friends3DSProtocol {
	Friends3DSProtocol := &Friends3DSProtocol{server: server}

	Friends3DSProtocol.Setup()

	return Friends3DSProtocol
}
