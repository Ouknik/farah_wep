package apiotts

const (

	// Type		: timestamp
	// Null		: NO
	// Key		:
	// Default	: current_timestamp()
	// Extra	:
	Crdate = "royaapiot_crdate"

	// Type		: varchar(64)
	// Null		: NO
	// Key		: PRI
	// Default	: %!s(<nil>)
	// Extra	:
	Id = "royaapiot_id"

	// Type		: bigint(20)
	// Null		: NO
	// Key		: PRI
	// Default	: %!s(<nil>)
	// Extra	:
	Device_id = "royaapiot_device_id"

	// Type		: varchar(8)
	// Null		: NO
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	Api_ott = "royaapiot_api_ott"

	// Type		: varchar(15)
	// Null		: YES
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	Ip = "royaapiot_ip"

	// Type		: timestamp
	// Null		: NO
	// Key		:
	// Default	: current_timestamp()
	// Extra	: on update current_timestamp()
	Used_on = "royaapiot_used_on"
)

var KeysForIteration_ry_api_otts = []string{Api_ott, Crdate, Device_id, Id, Ip, Used_on}
var SearchableFields_ry_api_otts = []string{}
