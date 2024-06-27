package models

import (
	"database/sql"

	tableElements "app.roya_immobile.com/models/category"

	"app.roya_immobile.com/lib"
)

type ModelCategory struct {
	DBAdapter lib.DbAdapter
}

func (m *ModelCategory) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m.DBAdapter.Connect("ry_categories", "royacateg_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *ModelCategory) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_categories", "royacateg_")
}

// Returns db fiedls for frontend iteration
// []string
func (m *ModelCategory) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_roya_categories
}

// Returns db fiedls which are serachable
// []string
func (m *ModelCategory) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_roya_categories
}

// Globalise DB Fields

// Type		: varchar(255)
// Null		: YES
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Address_additional = "zscatavlb_id"
func (m *ModelCategory) GetTableElement_Id() string {
	return tableElements.Id
}

// Type		: timestamp
// Null		: NO
// Key		:
// Default	: current_timestamp()
// Extra	:
// Crdate = "zscatavlb_category"
func (m *ModelCategory) GetTableElement_Category() string {
	return tableElements.Category
}

// Type		: int(1)
// Null		: YES
// Key		:
// Default	: 0
// Extra	:
// Is_deleted = "zscatavlb_priority"
func (m *ModelCategory) GetTableElement_Icon() string {
	return tableElements.Icon
}

func (m *ModelCategory) GetAllCategories() map[int]map[string]string {

	m.DBAdapter.SetOrderByField(tableElements.Id, m.DBAdapter.OrderBy.Asc())
	return m.DBAdapter.SelectAll().ExecSelect()
}

func (m *ModelCategory) FindByCategoryName(categoryName string) map[string]string {
	return m.DBAdapter.SelectOne().
		Where(tableElements.Category, categoryName, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()[0]
}
