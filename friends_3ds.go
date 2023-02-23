package nexproto

import (
	"errors"
	"fmt"

	nex "github.com/PretendoNetwork/nex-go"
)

const (
	// Friends3DSProtocolID is the protocol ID for the Friends (3DS) protocol
	Friends3DSProtocolID = 0x65

	// Friends3DSMethodUpdateProfile is the method ID for method UpdateProfile
	Friends3DSMethodUpdateProfile = 0x1

	// Friends3DSMethodUpdateMii is the method ID for method UpdateMii
	Friends3DSMethodUpdateMii = 0x2

	// Friends3DSMethodUpdatePreference is the method ID for method UpdatePreference
	Friends3DSMethodUpdatePreference = 0x5

	// Friends3DSMethodGetFriendMii is the method ID for method GetFriendMii
	Friends3DSMethodGetFriendMii = 0x6

	// Friends3DSMethodAddFriendByPID is the method ID for method AddFriendByPID
	Friends3DSMethodAddFriendByPID = 0xb

	// Friends3DSMethodGetPIDByLocalFriendCode is the method ID for method GetPIDByLocalFriendCode
	Friends3DSMethodGetPIDByLocalFriendCode = 0x9

	// Friends3DSMethodRemoveFriendByLocalFriendCode is the method ID for method RemoveFriendByLocalFriendCode
	Friends3DSMethodRemoveFriendByLocalFriendCode = 0xd

	// Friends3DSMethodRemoveFriendByPID is the method ID for method RemoveFriendByPID
	Friends3DSMethodRemoveFriendByPID = 0xe

	// Friends3DSMethodRemoveFriendByPID is the method ID for method RemoveFriendByPID
	Friends3DSMethodGetAllFriends = 0xf

	// Friends3DSMethodSyncFriend is the method ID for method SyncFriend
	Friends3DSMethodSyncFriend = 0x11

	// Friends3DSMethodUpdatePresence is the method ID for method UpdatePresence
	Friends3DSMethodUpdatePresence = 0x12

	// Friends3DSMethodUpdateFavoriteGameKey is the method ID for method UpdateFavoriteGameKey
	Friends3DSMethodUpdateFavoriteGameKey = 0x13

	// Friends3DSMethodUpdateComment is the method ID for method UpdateComment
	Friends3DSMethodUpdateComment = 0x14

	// Friends3DSMethodGetFriendPresence is the method ID for method GetFriendPresence
	Friends3DSMethodGetFriendPresence = 0x16

	// Friends3DSMethodGetFriendPersistentInfo is the method ID for method GetFriendPersistentInfo
	Friends3DSMethodGetFriendPersistentInfo = 0x19
)

// Friends3DSProtocol handles the Friends (3DS) nex protocol
type Friends3DSProtocol struct {
	server                               *nex.Server
	UpdateProfileHandler                 func(err error, client *nex.Client, callID uint32, profileData *MyProfile)
	UpdateMiiHandler                     func(err error, client *nex.Client, callID uint32, mii *Mii)
	UpdatePreferenceHandler              func(err error, client *nex.Client, callID uint32, publicMode bool, showGame bool, showPlayedGame bool)
	UpdatePresenceHandler                func(err error, client *nex.Client, callID uint32, presence *NintendoPresence, showGame bool)
	UpdateFavoriteGameKeyHandler         func(err error, client *nex.Client, callID uint32, gameKey *GameKey)
	UpdateCommentHandler                 func(err error, client *nex.Client, callID uint32, comment string)
	SyncFriendHandler                    func(err error, client *nex.Client, callID uint32, lfc uint64, pids []uint32, lfcList []uint64)
	GetPIDByLocalFriendCodeHandler       func(err error, client *nex.Client, callID uint32, lfc uint64, lfcList []uint64)
	AddFriendByPIDHandler                func(err error, client *nex.Client, callID uint32, lfc uint64, pid uint32)
	RemoveFriendByLocalFriendCodeHandler func(err error, client *nex.Client, callID uint32, lfc uint64)
	RemoveFriendByPIDHandler             func(err error, client *nex.Client, callID uint32, pid uint32)
	GetAllFriendsHandler                 func(err error, client *nex.Client, callID uint32)
	GetFriendPersistentInfoHandler       func(err error, client *nex.Client, callID uint32, pidList []uint32)
	GetFriendMiiHandler                  func(err error, client *nex.Client, callID uint32, pidList []uint32)
	GetFriendPresenceHandler             func(err error, client *nex.Client, callID uint32, pidList []uint32)
}

type Mii struct {
	Name     string
	Unknown2 bool
	Unknown3 uint8
	MiiData  []byte

	nex.Structure
}

// Bytes encodes the Mii and returns a byte array
func (mii *Mii) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteString(mii.Name)
	stream.WriteBool(mii.Unknown2)
	stream.WriteUInt8(mii.Unknown3)
	stream.WriteBuffer(mii.MiiData)

	return stream.Bytes()
}

// ExtractFromStream extracts a Mii from a stream
func (mii *Mii) ExtractFromStream(stream *nex.StreamIn) error {
	mii.Name, _ = stream.ReadString()
	mii.Unknown2 = (stream.ReadUInt8() == 1)
	mii.Unknown3 = stream.ReadUInt8()
	mii.MiiData, _ = stream.ReadBuffer()

	return nil
}

// NewMii returns a new Mii
func NewMii() *Mii {
	return &Mii{}
}

type FriendMii struct {
	PID        uint32
	Mii        *Mii
	ModifiedAt *nex.DateTime

	nex.Structure
}

// Bytes encodes the Mii and returns a byte array
func (friendMii *FriendMii) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(friendMii.PID)
	stream.WriteStructure(friendMii.Mii)
	stream.WriteUInt64LE(friendMii.ModifiedAt.Value())

	return stream.Bytes()
}

// NewMii returns a new Mii
func NewFriendMii() *FriendMii {
	return &FriendMii{}
}

type MyProfile struct {
	Region   uint8
	Country  uint8
	Area     uint8
	Language uint8
	Platform uint8
	Unknown1 uint64
	Unknown2 string
	Unknown3 string
}

// ExtractFromStream extracts a MyProfile from a stream
func (myProfile *MyProfile) ExtractFromStream(stream *nex.StreamIn) error {
	myProfile.Region = stream.ReadUInt8()
	myProfile.Country = stream.ReadUInt8()
	myProfile.Area = stream.ReadUInt8()
	myProfile.Language = stream.ReadUInt8()
	myProfile.Platform = stream.ReadUInt8()
	myProfile.Unknown1 = stream.ReadUInt64LE()
	myProfile.Unknown2, _ = stream.ReadString()
	myProfile.Unknown3, _ = stream.ReadString()

	return nil
}

// NewMyProfile returns a new MyProfile
func NewMyProfile() *MyProfile {
	return &MyProfile{}
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
	OwnerPID          uint32
	JoinGroupid       uint32
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
	stream.WriteUInt32LE(presence.OwnerPID)
	stream.WriteUInt32LE(presence.JoinGroupid)
	stream.WriteBuffer(presence.ApplicationArg)

	return stream.Bytes()
}

// ExtractFromStream extracts a NintendoPresence structure from a stream
func (presence *NintendoPresence) ExtractFromStream(stream *nex.StreamIn) error {
	if len(stream.Bytes()[stream.ByteOffset():]) < 25 {
		// length check for the following fixed-size data
		// changedFlags + JoinAvailableFlag + MatchmakeType + JoinGameID + JoinGameMode + OwnerPID + JoinGroupid
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
	OwnerPID := stream.ReadUInt32LE()
	JoinGroupid := stream.ReadUInt32LE()
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
	presence.OwnerPID = OwnerPID
	presence.JoinGroupid = JoinGroupid
	presence.ApplicationArg = ApplicationArg

	return nil
}

// NewNintendoPresence returns a new NintendoPresence
func NewNintendoPresence() *NintendoPresence {
	return &NintendoPresence{}
}

// FriendPresence contains information about a users online presence
type FriendPresence struct {
	PID      uint32
	Presence *NintendoPresence

	nex.Structure
}

// Bytes encodes the FriendPresence and returns a byte array
func (presence *FriendPresence) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(presence.PID)
	stream.WriteStructure(presence.Presence)

	return stream.Bytes()
}

// NewFriendPresence returns a new FriendPresence
func NewFriendPresence() *FriendPresence {
	return &FriendPresence{}
}

// FriendRelationship contains information about a users relationship with another PID
type FriendRelationship struct {
	PID              uint32
	LFC              uint64
	RelationshipType uint8

	nex.Structure
}

// Bytes encodes the FriendRelationship and returns a byte array
func (relationship *FriendRelationship) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(relationship.PID)
	stream.WriteUInt64LE(relationship.LFC)
	stream.WriteUInt8(relationship.RelationshipType)

	return stream.Bytes()
}

// NewFriendRelationship returns a new FriendRelationship
func NewFriendRelationship() *FriendRelationship {
	return &FriendRelationship{}
}

// FriendPersistentInfo contains user settings
type FriendPersistentInfo struct {
	PID              uint32
	Region           uint8
	Country          uint8
	Area             uint8
	Language         uint8
	Platform         uint8
	GameKey          *GameKey
	Message          string
	MessageUpdatedAt *nex.DateTime //appears to be correct, but not 100% sure.
	FriendedAt       *nex.DateTime
	LastOnline       *nex.DateTime

	nex.Structure
}

// Bytes encodes the FriendPersistentInfo and returns a byte array
func (friendPersistentInfo *FriendPersistentInfo) Bytes(stream *nex.StreamOut) []byte {
	stream.WriteUInt32LE(friendPersistentInfo.PID)
	stream.WriteUInt8(friendPersistentInfo.Region)
	stream.WriteUInt8(friendPersistentInfo.Country)
	stream.WriteUInt8(friendPersistentInfo.Area)
	stream.WriteUInt8(friendPersistentInfo.Language)
	stream.WriteUInt8(friendPersistentInfo.Platform)
	stream.WriteStructure(friendPersistentInfo.GameKey)
	stream.WriteString(friendPersistentInfo.Message)
	stream.WriteUInt64LE(friendPersistentInfo.MessageUpdatedAt.Value())
	stream.WriteUInt64LE(friendPersistentInfo.FriendedAt.Value())
	stream.WriteUInt64LE(friendPersistentInfo.LastOnline.Value())

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
			case Friends3DSMethodUpdateProfile:
				go friends3DSProtocol.handleUpdateProfile(packet)
			case Friends3DSMethodUpdateMii:
				go friends3DSProtocol.handleUpdateMii(packet)
			case Friends3DSMethodUpdatePreference:
				go friends3DSProtocol.handleUpdatePreference(packet)
			case Friends3DSMethodSyncFriend:
				go friends3DSProtocol.handleSyncFriend(packet)
			case Friends3DSMethodUpdatePresence:
				go friends3DSProtocol.handleUpdatePresence(packet)
			case Friends3DSMethodUpdateFavoriteGameKey:
				go friends3DSProtocol.handleUpdateFavoriteGameKey(packet)
			case Friends3DSMethodUpdateComment:
				go friends3DSProtocol.handleUpdateComment(packet)
			case Friends3DSMethodGetPIDByLocalFriendCode:
				go friends3DSProtocol.handleGetPIDByLocalFriendCode(packet)
			case Friends3DSMethodAddFriendByPID:
				go friends3DSProtocol.handleAddFriendByPID(packet)
			case Friends3DSMethodRemoveFriendByLocalFriendCode:
				go friends3DSProtocol.handleRemoveFriendByLocalFriendCode(packet)
			case Friends3DSMethodRemoveFriendByPID:
				go friends3DSProtocol.handleRemoveFriendByPID(packet)
			case Friends3DSMethodGetAllFriends:
				go friends3DSProtocol.handleGetAllFriends(packet)
			case Friends3DSMethodGetFriendPersistentInfo:
				go friends3DSProtocol.handleGetFriendPersistentInfo(packet)
			case Friends3DSMethodGetFriendMii:
				go friends3DSProtocol.handleGetFriendMii(packet)
			case Friends3DSMethodGetFriendPresence:
				go friends3DSProtocol.handleGetFriendPresence(packet)
			default:
				go respondNotImplemented(packet, Friends3DSProtocolID)
				fmt.Printf("Unsupported Friends (3DS) method ID: %#v\n", request.MethodID())
			}
		}
	})
}

func (friends3DSProtocol *Friends3DSProtocol) handleUpdateMii(packet nex.PacketInterface) {
	if friends3DSProtocol.UpdateMiiHandler == nil {
		logger.Warning("Friends3DSProtocol::UpdateMii not implemented")
		go respondNotImplemented(packet, Friends3DSProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	mii := NewMii()
	mii.ExtractFromStream(parametersStream)

	go friends3DSProtocol.UpdateMiiHandler(nil, client, callID, mii)
}

func (friends3DSProtocol *Friends3DSProtocol) handleUpdateProfile(packet nex.PacketInterface) {
	if friends3DSProtocol.UpdateProfileHandler == nil {
		logger.Warning("Friends3DSProtocol::UpdateProfile not implemented")
		go respondNotImplemented(packet, Friends3DSProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	profileData := NewMyProfile()
	profileData.ExtractFromStream(parametersStream)

	go friends3DSProtocol.UpdateProfileHandler(nil, client, callID, profileData)
}

func (friends3DSProtocol *Friends3DSProtocol) handleUpdatePreference(packet nex.PacketInterface) {
	if friends3DSProtocol.UpdatePreferenceHandler == nil {
		logger.Warning("Friends3DSProtocol::UpdatePreference not implemented")
		go respondNotImplemented(packet, Friends3DSProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	publicMode := (parametersStream.ReadUInt8() == 1)
	showGame := (parametersStream.ReadUInt8() == 1)
	showPlayedGame := (parametersStream.ReadUInt8() == 1)

	go friends3DSProtocol.UpdatePreferenceHandler(nil, client, callID, publicMode, showGame, showPlayedGame)
}

func (friends3DSProtocol *Friends3DSProtocol) handleSyncFriend(packet nex.PacketInterface) {
	if friends3DSProtocol.SyncFriendHandler == nil {
		logger.Warning("Friends3DSProtocol::SyncFriend not implemented")
		go respondNotImplemented(packet, Friends3DSProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	lfc := parametersStream.ReadUInt64LE()
	pids := parametersStream.ReadListUInt32LE()
	lfcList := parametersStream.ReadListUInt64LE()

	go friends3DSProtocol.SyncFriendHandler(nil, client, callID, lfc, pids, lfcList)
}

func (friends3DSProtocol *Friends3DSProtocol) handleGetFriendPersistentInfo(packet nex.PacketInterface) {
	if friends3DSProtocol.GetFriendPersistentInfoHandler == nil {
		logger.Warning("Friends3DSProtocol::GetFriendPersistentInfo not implemented")
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
		logger.Warning("Friends3DSProtocol::UpdatePresence not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	nintendoPresence := NewNintendoPresence()
	nintendoPresence.ExtractFromStream(parametersStream)

	showGame := (parametersStream.ReadUInt8() == 1)

	go friends3DSProtocol.UpdatePresenceHandler(nil, client, callID, nintendoPresence, showGame)
}

func (friends3DSProtocol *Friends3DSProtocol) handleUpdateFavoriteGameKey(packet nex.PacketInterface) {
	if friends3DSProtocol.UpdateFavoriteGameKeyHandler == nil {
		logger.Warning("Friends3DSProtocol::UpdateFavoriteGameKey not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	gameKey := NewGameKey()
	gameKey.ExtractFromStream(parametersStream)

	go friends3DSProtocol.UpdateFavoriteGameKeyHandler(nil, client, callID, gameKey)
}

func (friends3DSProtocol *Friends3DSProtocol) handleUpdateComment(packet nex.PacketInterface) {
	if friends3DSProtocol.UpdateCommentHandler == nil {
		logger.Warning("Friends3DSProtocol::UpdateComment not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	comment, _ := parametersStream.ReadString()

	go friends3DSProtocol.UpdateCommentHandler(nil, client, callID, comment)
}

func (friends3DSProtocol *Friends3DSProtocol) handleGetPIDByLocalFriendCode(packet nex.PacketInterface) {
	if friends3DSProtocol.GetPIDByLocalFriendCodeHandler == nil {
		logger.Warning("Friends3DSProtocol::GetPIDByLocalFriendCode not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	lfc := parametersStream.ReadUInt64LE()
	lfcList := parametersStream.ReadListUInt64LE()

	go friends3DSProtocol.GetPIDByLocalFriendCodeHandler(nil, client, callID, lfc, lfcList)
}

func (friends3DSProtocol *Friends3DSProtocol) handleAddFriendByPID(packet nex.PacketInterface) {
	if friends3DSProtocol.AddFriendByPIDHandler == nil {
		logger.Warning("Friends3DSProtocol::AddFriendByPID not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	lfc := parametersStream.ReadUInt64LE()
	PID := parametersStream.ReadUInt32LE()

	go friends3DSProtocol.AddFriendByPIDHandler(nil, client, callID, lfc, PID)
}

func (friends3DSProtocol *Friends3DSProtocol) handleRemoveFriendByLocalFriendCode(packet nex.PacketInterface) {
	if friends3DSProtocol.RemoveFriendByLocalFriendCodeHandler == nil {
		logger.Warning("Friends3DSProtocol::RemoveFriendByLocalFriendCode not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	lfc := parametersStream.ReadUInt64LE()

	go friends3DSProtocol.RemoveFriendByLocalFriendCodeHandler(nil, client, callID, lfc)
}

func (friends3DSProtocol *Friends3DSProtocol) handleRemoveFriendByPID(packet nex.PacketInterface) {
	if friends3DSProtocol.RemoveFriendByPIDHandler == nil {
		logger.Warning("Friends3DSProtocol::RemoveFriendByPID not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()
	parameters := request.Parameters()

	parametersStream := nex.NewStreamIn(parameters, friends3DSProtocol.server)

	PID := parametersStream.ReadUInt32LE()

	go friends3DSProtocol.RemoveFriendByPIDHandler(nil, client, callID, PID)
}

func (friends3DSProtocol *Friends3DSProtocol) handleGetAllFriends(packet nex.PacketInterface) {
	if friends3DSProtocol.GetAllFriendsHandler == nil {
		logger.Warning("Friends3DSProtocol::GetAllFriends not implemented")
		go respondNotImplemented(packet, AuthenticationProtocolID)
		return
	}

	client := packet.Sender()
	request := packet.RMCRequest()

	callID := request.CallID()

	go friends3DSProtocol.GetAllFriendsHandler(nil, client, callID)
}

func (friends3DSProtocol *Friends3DSProtocol) handleGetFriendMii(packet nex.PacketInterface) {
	if friends3DSProtocol.GetFriendMiiHandler == nil {
		logger.Warning("Friends3DSProtocol::GetFriendMiiHandler not implemented")
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
		logger.Warning("Friends3DSProtocol::GetFriendPresenceHandler not implemented")
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

// UpdateProfile sets the UpdateProfile handler function
func (friends3DSProtocol *Friends3DSProtocol) UpdateProfile(handler func(err error, client *nex.Client, callID uint32, profileData *MyProfile)) {
	friends3DSProtocol.UpdateProfileHandler = handler
}

// UpdateMii sets the UpdateMii handler function
func (friends3DSProtocol *Friends3DSProtocol) UpdateMii(handler func(err error, client *nex.Client, callID uint32, mii *Mii)) {
	friends3DSProtocol.UpdateMiiHandler = handler
}

// UpdatePreference sets the UpdatePreference handler function
func (friends3DSProtocol *Friends3DSProtocol) UpdatePreference(handler func(err error, client *nex.Client, callID uint32, publicMode bool, showGame bool, showPlayedGame bool)) {
	friends3DSProtocol.UpdatePreferenceHandler = handler
}

// GetPIDByLocalFriendCode sets the GetPIDByLocalFriendCode handler function
func (friends3DSProtocol *Friends3DSProtocol) GetPIDByLocalFriendCode(handler func(err error, client *nex.Client, callID uint32, lfc uint64, lfcList []uint64)) {
	friends3DSProtocol.GetPIDByLocalFriendCodeHandler = handler
}

// AddFriendByPID sets the AddFriendByPID handler function
func (friends3DSProtocol *Friends3DSProtocol) AddFriendByPID(handler func(err error, client *nex.Client, callID uint32, lfc uint64, pid uint32)) {
	friends3DSProtocol.AddFriendByPIDHandler = handler
}

// RemoveFriendByLocalFriendCode sets the RemoveFriendByLocalFriendCode handler function
func (friends3DSProtocol *Friends3DSProtocol) RemoveFriendByLocalFriendCode(handler func(err error, client *nex.Client, callID uint32, lfc uint64)) {
	friends3DSProtocol.RemoveFriendByLocalFriendCodeHandler = handler
}

// RemoveFriendByPID sets the RemoveFriendByPID handler function
func (friends3DSProtocol *Friends3DSProtocol) RemoveFriendByPID(handler func(err error, client *nex.Client, callID uint32, pid uint32)) {
	friends3DSProtocol.RemoveFriendByPIDHandler = handler
}

// GetAllFriends sets the GetAllFriends handler function
func (friends3DSProtocol *Friends3DSProtocol) GetAllFriends(handler func(err error, client *nex.Client, callID uint32)) {
	friends3DSProtocol.GetAllFriendsHandler = handler
}

// SyncFriend sets the SyncFriend handler function
func (friends3DSProtocol *Friends3DSProtocol) SyncFriend(handler func(err error, client *nex.Client, callID uint32, lfc uint64, pids []uint32, lfcList []uint64)) {
	friends3DSProtocol.SyncFriendHandler = handler
}

// UpdatePresence sets the UpdatePresence handler function
func (friends3DSProtocol *Friends3DSProtocol) UpdatePresence(handler func(err error, client *nex.Client, callID uint32, presence *NintendoPresence, showGame bool)) {
	friends3DSProtocol.UpdatePresenceHandler = handler
}

// UpdateFavoriteGameKey sets the UpdateFavoriteGameKey handler function
func (friends3DSProtocol *Friends3DSProtocol) UpdateFavoriteGameKey(handler func(err error, client *nex.Client, callID uint32, gameKey *GameKey)) {
	friends3DSProtocol.UpdateFavoriteGameKeyHandler = handler
}

// UpdateComment sets the UpdateComment handler function
func (friends3DSProtocol *Friends3DSProtocol) UpdateComment(handler func(err error, client *nex.Client, callID uint32, comment string)) {
	friends3DSProtocol.UpdateCommentHandler = handler
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

// NewFriends3DSProtocol returns a new Friends3DSProtocol
func NewFriends3DSProtocol(server *nex.Server) *Friends3DSProtocol {
	Friends3DSProtocol := &Friends3DSProtocol{server: server}

	Friends3DSProtocol.Setup()

	return Friends3DSProtocol
}
