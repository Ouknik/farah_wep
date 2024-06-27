package models

import (
	"database/sql"

	tableElements "app.roya_immobile.com/models/abilities"

	"app.roya_immobile.com/lib"
)

type ModelAbilities struct {
	DBAdapter lib.DbAdapter
}

func (m *ModelAbilities) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m.DBAdapter.Connect("ry_abilities", "ryabilit_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *ModelAbilities) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_abilities", "ryabilit_")
}

// Returns db fiedls for frontend iteration
// []string
func (m *ModelAbilities) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_roya_abilities
}

// Returns db fiedls which are serachable
// []string
func (m *ModelAbilities) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_roya_abilities
}

// Globalise DB Fields

// Type		: varchar(255)
// Null		: YES
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Address_additional = "zscatavlb_id"
func (m *ModelAbilities) GetTableElement_Id() string {
	return tableElements.Id
}

// Type		: timestamp
// Null		: NO
// Key		:
// Default	: current_timestamp()
// Extra	:
// Crdate = "zscatavlb_category"
func (m *ModelAbilities) GetTableElement_NameAbili() string {
	return tableElements.NameAbili
}

// Type		: int(1)
// Null		: YES
// Key		:
// Default	: 0
// Extra	:
// Is_deleted = "zscatavlb_priority"
func (m *ModelAbilities) GetTableElement_Icon() string {
	return tableElements.Icon
}

func (m *ModelAbilities) GetTableElement_Type() string {
	return tableElements.Type
}

func (m *ModelAbilities) GetAllAbilities() map[int]map[string]string {

	m.DBAdapter.SetOrderByField(tableElements.Id, m.DBAdapter.OrderBy.Asc())
	return m.DBAdapter.SelectAll().ExecSelect()
}

func (m *ModelAbilities) FindByAbilitiesVaribaleId(abilites_id string) map[string]string {
	return m.DBAdapter.SelectAll().
		Where(tableElements.Id, abilites_id, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()[0]
}
