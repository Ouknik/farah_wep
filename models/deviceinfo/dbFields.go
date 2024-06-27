package deviceinfo

const (

	// Type		: timestamp
	// Null		: NO
	// Key		:
	// Default	: current_timestamp()
	// Extra	:
	Crdate = "royadvcinf_crdate"

	// Type		: bigint(20)
	// Null		: NO
	// Key		: PRI
	// Default	: %!s(<nil>)
	// Extra	: auto_increment
	Id = "royadvcinf_id"

	// Type		: varchar(255)
	// Null		: NO
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	Uuid = "royadvcinf_uuid"

	// Type		: tinytext
	// Null		: NO
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	Hardware = "royadvcinf_hardware"

	// Type		: tinytext
	// Null		: NO
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	Fcm_token = "royadvcinf_fcm_token"

	// Type		: varchar(255)
	// Null		: NO
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	Model = "royadvcinf_model"

	// Type		: varchar(255)
	// Null		: NO
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	Os = "royadvcinf_os"

	// Type		: varchar(255)
	// Null		: NO
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	Os_version = "royadvcinf_os_version"

	// Type		: varchar(15)
	// Null		: NO
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	Ip = "royadvcinf_ip"

	User_id = "royadvcinf_userid"
)

var KeysForIteration_ry_device_info = []string{Crdate, User_id, Hardware, Id, Ip, Model, Os, Os_version, Uuid, Fcm_token}
var SearchableFields_ry_device_info = []string{}
