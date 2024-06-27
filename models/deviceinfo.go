package models

import (
	"database/sql"

	"app.roya_immobile.com/lib"
	tableElements "app.roya_immobile.com/models/deviceinfo"
)

type ModelDeviceinfo struct {
	DBAdapter lib.DbAdapter
}

func (m *ModelDeviceinfo) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m.DBAdapter.Connect("ry_device_info", "royadvcinf_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *ModelDeviceinfo) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_device_info", "royadvcinf_")
}

// Returns db fiedls for frontend iteration
// []string
func (m *ModelDeviceinfo) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_ry_device_info
}

// Returns db fiedls which are serachable
// []string
func (m *ModelDeviceinfo) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_ry_device_info
}

// Globalise DB Fields

// Type		: varchar(255)
// Null		: NO
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Model = "zsdvcinf_model"
func (m *ModelDeviceinfo) GetTableElement_Model() string {
	return tableElements.Model
}

// Type		: varchar(255)
// Null		: NO
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Os = "zsdvcinf_os"
func (m *ModelDeviceinfo) GetTableElement_Os() string {
	return tableElements.Os
}

// Type		: varchar(255)
// Null		: NO
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Os_version = "zsdvcinf_os_version"
func (m *ModelDeviceinfo) GetTableElement_Os_version() string {
	return tableElements.Os_version
}

func (m *ModelDeviceinfo) GetTableElement_UserID() string {
	return tableElements.User_id
}

// Type		: varchar(15)
// Null		: NO
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Ip = "zsdvcinf_ip"
func (m *ModelDeviceinfo) GetTableElement_Ip() string {
	return tableElements.Ip
}

// Type		: timestamp
// Null		: NO
// Key		:
// Default	: current_timestamp()
// Extra	:
// Crdate = "zsdvcinf_crdate"
func (m *ModelDeviceinfo) GetTableElement_Crdate() string {
	return tableElements.Crdate
}

// Type		: bigint(20)
// Null		: NO
// Key		: PRI
// Default	: %!s(<nil>)
// Extra	: auto_increment
// Id = "zsdvcinf_id"
func (m *ModelDeviceinfo) GetTableElement_Id() string {
	return tableElements.Id
}

// Type		: varchar(255)
// Null		: NO
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Uuid = "zsdvcinf_uuid"
func (m *ModelDeviceinfo) GetTableElement_Uuid() string {
	return tableElements.Uuid
}

// Type		: tinytext
// Null		: NO
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Hardware = "zsdvcinf_hardware"
func (m *ModelDeviceinfo) GetTableElement_Hardware() string {
	return tableElements.Hardware
}

// Type		: tinytext
// Null		: NO
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Token_fcm = "zsdvcinf_fcm_token"
func (m *ModelDeviceinfo) GetTableElement_token() string {
	return tableElements.Fcm_token
}

func (m *ModelDeviceinfo) InsertDeviceinfo(fields, values []string) (int64, string) {
	return m.DBAdapter.Insert(fields, values)
}

func (m *ModelDeviceinfo) FindById(id string) map[string]string {
	return m.DBAdapter.SelectOne().
		WhereMathComparision(tableElements.Id, id, m.DBAdapter.MathComparision.EqalTo()).
		ExecSelect()[0]
}

func (m *ModelDeviceinfo) FindByUuid(uuid string) map[string]string {
	return m.DBAdapter.SelectOne().
		WhereMathComparision(tableElements.Uuid, uuid, m.DBAdapter.MathComparision.EqalTo()).
		ExecSelect()[0]
}

func (m *ModelDeviceinfo) UpdateById(id string, fields, values []string) {
	m.DBAdapter.Update(fields, values).
		Where(tableElements.Id, id, m.DBAdapter.Wildcard.NotWildcard()).
		ExecUpdate()
}

func (m *ModelDeviceinfo) FindDviceInfoByUserID(userId string) map[string]string {
	return m.DBAdapter.SelectOne().
		WhereMathComparision(tableElements.User_id, userId, m.DBAdapter.MathComparision.EqalTo()).
		ExecSelect()[0]
}
