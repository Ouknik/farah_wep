package models

import (
	"database/sql"

	tableElements "app.roya_immobile.com/models/imageannonce"

	"app.roya_immobile.com/lib"
)

type ModelImageAnnonce struct {
	DBAdapter lib.DbAdapter
}

func (m *ModelImageAnnonce) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m.DBAdapter.Connect("ry_annose_images", "royaimg_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *ModelImageAnnonce) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_annose_images", "royaimg_")
}

// Returns db fiedls for frontend iteration
// []string
func (m *ModelImageAnnonce) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_roya_imageannoce
}

// Returns db fiedls which are serachable
// []string
func (m *ModelImageAnnonce) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_roya_imageannoce
}

// Globalise DB Fields

// Type		: varchar(255)
// Null		: YES
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Address_additional = "zscatavlb_id"
func (m *ModelImageAnnonce) GetTableElement_Id() string {
	return tableElements.Id
}

// Type		: timestamp
// Null		: NO
// Key		:
// Default	: current_timestamp()
// Extra	:
// Crdate = "zscatavlb_category"
func (m *ModelImageAnnonce) GetTableElement_ImageAnnonce() string {
	return tableElements.ImageAnnonce
}

func (m *ModelImageAnnonce) FindByImageVaribaleId(id string) map[int]map[string]string {
	return m.DBAdapter.SelectAll().
		Where(tableElements.Id, id, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()

}
