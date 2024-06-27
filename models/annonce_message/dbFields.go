package annonce_message

const (
	MessageId    = "ryannmess_id"         // Type: int(20), Null: NO, Key: PRI, Default: <auto_increment>, Extra: AUTO_INCREMENT
	UserId       = "ryannmess_user_id"    // Type: int(20), Null: NO, Key: , Default: , Extra:
	AdminId      = "ryannmess_admin_id"   // Type: int(20), Null: NO, Key: , Default: , Extra:
	AnnonceId    = "ryannmess_annonce_id" // Type: int(20), Null: NO, Key: , Default: , Extra:
	Message      = "ryannmess_message"    // Type: varchar(255), Null: YES, Key: , Default: , Extra:
	CreationDate = "ryannmess_crdate"     // Type: timestamp, Null: , Key: , Default: CURRENT_TIMESTAMP, Extra: DEFAULT_GENERATED
	CreatedBy    = "ryannmess_create_by"  // Type: int(20), Null: NO, Key: , Default: , Extra:
	SeeMessageAt = "ryannmess_seemessage"
	Status       = "ryannmess_status"
)

var KeysForIterationAnnonceMessage = []string{MessageId, UserId, Status, SeeMessageAt, AdminId, AnnonceId, Message, CreationDate, CreatedBy}
var SearchableFieldsAnnonceMessage = []string{}
