package model

// Services
const (
	// PARTY
	PARTY = "PARTY"
)

// Commands
const (
	// COMMAND CONST
	COMMAND = "COMMAND"

	// HEARTBEAT CONST
	HEARTBEAT = "HEARTBEAT"

	// JOIN
	JOIN = "JOIN"

	// DISBAND
	DISBAND = "DISBAND"
)

// States
const (
	// OPENED State
	OPENED = "OPENED"

	// SUSPENDED State
	SUSPENDED = "SUSPENDED"

	// STALE State
	STALE = "STALE"

	// INGAME State
	INGAME = "IN-GAME"

	// CLOSED State
	CLOSED = "CLOSED"

	// CONNECTED State
	CONNECTED = "CONNECTED"

	// DISCONNECTED State
	DISCONNECTED = "DISCONNECTED"

	// JOINED State
	JOINED = "PLAYER_JOINED"

	// DISBANDED State
	DISBANDED = "PLAYER_DISBANDED"

	// NOTAVAILABLE State
	NOTAVAILABLE = "NOTAVAILABLE"
)

// Actions
const (
	// ERROR STATE
	ERROR = "ERROR"

	// CONNECT State
	CONNECT = "CONNECT"

	// DISCONENECT State
	DISCONNECT = "DISCONNECT"
)