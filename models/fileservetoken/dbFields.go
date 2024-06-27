package fileservetoken

const (

	// Type		: varchar(32)
	// Null		: NO
	// Key		: PRI
	// Default	: %!s(<nil>)
	// Extra	:
	Token = "zsflsrvtkn_token"

	// Type		: varchar(255)
	// Null		: NO
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	Portal_key = "zsflsrvtkn_portal_key"

	// Type		: tinytext
	// Null		: YES
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	File_name = "zsflsrvtkn_file_name"

	// Type		: tinyint(4)
	// Null		: NO
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	Destination = "zsflsrvtkn_destination"

	// Type		: varchar(20)
	// Null		: YES
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	File_size = "zsflsrvtkn_file_size"

	// Type		: varchar(20)
	// Null		: YES
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	File_mime_type = "zsflsrvtkn_file_mime_type"

	// Type		: int(20)
	// Null		: NO
	// Key		:
	// Default	: %!s(<nil>)
	// Extra	:
	Created_by = "zsflsrvtkn_created_by"

	// Type		: timestamp
	// Null		: NO
	// Key		:
	// Default	: current_timestamp()
	// Extra	:
	Crdate = "zsflsrvtkn_crdate"
)

var KeysForIteration_zs_file_serve_token = []string{}
var SearchableFields_zs_file_serve_token = []string{}
