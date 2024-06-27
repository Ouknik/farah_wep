package models

import (
	"database/sql"

	"app.roya_immobile.com/lib"
	tableElements "app.roya_immobile.com/models/apiotts"
)

type ModelApiotts struct {
	DBAdapter lib.DbAdapter
}

func (m *ModelApiotts) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m.DBAdapter.Connect("ry_api_otts", "royaapiot_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *ModelApiotts) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_api_otts", "royaapiot_")
}

// Returns db fiedls for frontend iteration
// []string
func (m *ModelApiotts) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_ry_api_otts
}

// Returns db fiedls which are serachable
// []string
func (m *ModelApiotts) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_ry_api_otts
}

// Globalise DB Fields

// Type		: timestamp
// Null		: NO
// Key		:
// Default	: current_timestamp()
// Extra	: on update current_timestamp()
// Used_on = "zsapiot_used_on"
func (m *ModelApiotts) GetTableElement_Used_on() string {
	return tableElements.Used_on
}

// Type		: timestamp
// Null		: NO
// Key		:
// Default	: current_timestamp()
// Extra	:
// Crdate = "zsapiot_crdate"
func (m *ModelApiotts) GetTableElement_Crdate() string {
	return tableElements.Crdate
}

// Type		: varchar(64)
// Null		: NO
// Key		: PRI
// Default	: %!s(<nil>)
// Extra	:
// Id = "zsapiot_id"
func (m *ModelApiotts) GetTableElement_Id() string {
	return tableElements.Id
}

// Type		: bigint(20)
// Null		: NO
// Key		: PRI
// Default	: %!s(<nil>)
// Extra	:
// Device_id = "zsapiot_device_id"
func (m *ModelApiotts) GetTableElement_Device_id() string {
	return tableElements.Device_id
}

// Type		: varchar(8)
// Null		: NO
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Api_ott = "zsapiot_api_ott"
func (m *ModelApiotts) GetTableElement_Api_ott() string {
	return tableElements.Api_ott
}

// Type		: varchar(15)
// Null		: YES
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Ip = "zsapiot_ip"
func (m *ModelApiotts) GetTableElement_Ip() string {
	return tableElements.Ip
}

func (m *ModelApiotts) GenerateApiOtt(deviceId string) string {
	var apiOtt = lib.GenerateRandomAlphaNumericString(8)
	m.DBAdapter.Insert([]string{
		m.GetTableElement_Id(),
		m.GetTableElement_Api_ott(),
		m.GetTableElement_Device_id(),
	}, []string{
		lib.GenerateRandomAlphaNumericString(64),
		apiOtt,
		deviceId,
	})
	return apiOtt
}

// func (m *ModelApiotts) FindById(id string) map[string]string {
// 	return m.DBAdapter.SelectOne().
// 		WhereMathComparision(tableElements.Id, id, m.DBAdapter.MathComparision.EqalTo()).
// 		ExecSelect()[0]
// }

func (m *ModelApiotts) FindByApiOttAndDeviceId(apiOtt, deviceId string) map[string]string {
	return m.DBAdapter.SelectOne().
		Where(tableElements.Api_ott, apiOtt, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.Device_id, deviceId, m.DBAdapter.Wildcard.NotWildcard()).
		WhereMathComparision(tableElements.Used_on, "", m.DBAdapter.MathComparision.IsNull()).
		ExecSelect()[0]
}

func (m *ModelApiotts) MarkOttAsUsed(apiOtt, ip string) {
	var timeHandler = lib.TimeHandler{}
	m.DBAdapter.Update([]string{
		tableElements.Ip,
		tableElements.Used_on,
	}, []string{
		ip,
		timeHandler.Now(),
	}).
		Where(tableElements.Api_ott, apiOtt, m.DBAdapter.Wildcard.NotWildcard()).
		ExecUpdate()
}
