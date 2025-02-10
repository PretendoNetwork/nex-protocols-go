package constants

// SearchType represents the type of user who can access an object.
// This is stored in DataStoreSearchParam.searchTarget.
// Used to narrow search results based on access rights
type SearchType uint8

const (

	// SearchTypePublic selects objects whose access permission
	// is set to PermissionPublic
	SearchTypePublic SearchType = iota + 1

	// SearchTypeSendFriend selects objects owned by the current
	// user and whose access permission is set to PermissionFriend
	SearchTypeSendFriend

	// SearchTypeSendSpecified selects objects owned by the current
	// user and whose access permission is set to PermissionSpecified
	SearchTypeSendSpecified

	// SearchTypeSendSpecifiedFriend selects objects owned by the current
	// user and whose access permission is set to PermissionSpecifiedFriend
	SearchTypeSendSpecifiedFriend

	// SearchTypeSend selects objects owned by the current
	// user and whose access permission is set to one of
	// PermissionFriend, PermissionSpecified or
	// PermissionSpecifiedFriend
	SearchTypeSend

	// SearchTypeFriend selects objects owned by the friends of the
	// current user whose access permission is set to PermissionFriend
	SearchTypeFriend

	// SearchTypeReceivedSpecified selects objects whose access permission
	// is set to either PermissionSpecified or PermissionSpecifiedFriend
	// and which the current user is in the recipient IDs list
	SearchTypeReceivedSpecified

	// SearchTypeReceived selects objects who match the criteria of either
	// (or both) SearchTypeFriend/SearchTypeReceivedSpecified
	SearchTypeReceived

	// SearchTypePrivate selects objects owned by the current user and  whose
	// access permission is set to PermissionPrivate
	SearchTypePrivate

	// SearchTypeOwn selects objects owned by the current user and whose
	// status is set to DataStatusNone
	SearchTypeOwn

	// SearchTypePublicExcludeOwnAndFriends selects objects whose access
	// permission is set to PermissionPublic but excludes objects owned
	// by the current user and their friends
	SearchTypePublicExcludeOwnAndFriends

	// SearchTypeOwnPending selects objects owned by the current user and whose
	// status is set to DataStatusPending
	SearchTypeOwnPending

	// SearchTypeOwnRejected selects objects owned by the current user and whose
	// status is set to DataStatusRejected
	SearchTypeOwnRejected

	// SearchTypeOwnAll selects objects owned by the current user regardless of
	// status or access permission
	SearchTypeOwnAll
)
