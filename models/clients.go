package models

import (
	"database/sql"

	"app.roya_immobile.com/lib"
	tableElements "app.roya_immobile.com/models/clients"
)

type ModelClients struct {
	DBAdapter lib.DbAdapter
}

func (m *ModelClients) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m.DBAdapter.Connect("ry_clients", "ryclien_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *ModelClients) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_clients", "ryclien_")
}

// Returns db fiedls for frontend iteration
// []string
func (m *ModelClients) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_roya_clients
}

// Returns db fiedls which are serachable
// []string
func (m *ModelClients) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_roya_clients
}

// Type		: timestamp
// Null		: NO
// Key		:
// Default	: current_timestamp()
// Extra	:
// Crdate = "zsapiot_crdate"
func (m *ModelClients) GetTableElement_Crdate() string {
	return tableElements.Crdate
}

// Type		: varchar(64)
// Null		: NO
// Key		: PRI
// Default	: %!s(<nil>)
// Extra	:
// Id = "zsapiot_id"
func (m *ModelClients) GetTableElement_Id() string {
	return tableElements.Id
}

// Type		: bigint(20)
// Null		: NO
// Key		: PRI
// Default	: %!s(<nil>)
// Extra	:
// Device_id = "zsapiot_device_id"
func (m *ModelClients) GetTableElement_UserId() string {
	return tableElements.UserId
}

// Type		: varchar(8)
// Null		: NO
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Api_ott = "zsapiot_api_ott"
func (m *ModelClients) GetTableElement_IsBlocked() string {
	return tableElements.IsBlocked
}

// Type		: varchar(15)
// Null		: YES
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Ip = "zsapiot_ip"
func (m *ModelClients) GetTableElement_Cin() string {
	return tableElements.Cin
}

func (m *ModelClients) GetTableElement_AnnonceId() string {
	return tableElements.AnnonceId
}

func (m *ModelClients) GetTableElement_Statue() string {
	return tableElements.Statue
}

func (m *ModelClients) GetTableElement_Prix() string {
	return tableElements.Prix
}

func (m *ModelClients) InseertCliens(fields []string, values []string) (int64, string) {
	return m.DBAdapter.Insert(fields, values)
}

func (m *ModelClients) GetAllCliens(_modelUserapp ModelUserapp) map[int]map[string]string {
	return m.DBAdapter.SelectAll().
		RightJoin(_modelUserapp.DBAdapter.GetTableName(), _modelUserapp.GetTableElement_Id(), tableElements.UserId).
		Where(tableElements.IsBlocked, "0", m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()
}

func (m *ModelClients) UpdateCliens(clienid string, fildes []string, values []string) {
	m.DBAdapter.Update(fildes, values).Where(tableElements.UserId, clienid, m.DBAdapter.Wildcard).ExecUpdate()
}
