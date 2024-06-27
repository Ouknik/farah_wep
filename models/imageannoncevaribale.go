package models

import (
	"database/sql"

	tableElements "app.roya_immobile.com/models/imageannoncevaribale"

	"app.roya_immobile.com/lib"
)

type ModelImageAnnonceVaribale struct {
	DBAdapter lib.DbAdapter
}

func (m *ModelImageAnnonceVaribale) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m.DBAdapter.Connect("ry_annonce_images_avariable", "royaannimgava_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *ModelImageAnnonceVaribale) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_annonce_images_avariable", "royaannimgava_")
}

// Returns db fiedls for frontend iteration
// []string
func (m *ModelImageAnnonceVaribale) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_roya_imageannoce
}

// Returns db fiedls which are serachable
// []string
func (m *ModelImageAnnonceVaribale) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_roya_imageannoce
}

// Globalise DB Fields

// Type		: varchar(255)
// Null		: YES
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Address_additional = "zscatavlb_id"
func (m *ModelImageAnnonceVaribale) GetTableElement_Id() string {
	return tableElements.Id
}

// Type		: timestamp
// Null		: NO
// Key		:
// Default	: current_timestamp()
// Extra	:
// Crdate = "zscatavlb_category"
func (m *ModelImageAnnonceVaribale) GetTableElement_Image_Id() string {
	return tableElements.ImageId
}

// Type		: timestamp
// Null		: NO
// Key		:
// Default	: current_timestamp()
// Extra	:
// Crdate = "zscatavlb_category"
func (m *ModelImageAnnonceVaribale) GetTableElement_AnnonceId() string {
	return tableElements.AnnonceId
}

func (m *ModelImageAnnonceVaribale) FindByAnnonceId(id string) map[int]map[string]string {
	return m.DBAdapter.SelectAll().
		WhereMathComparision(tableElements.AnnonceId, id, m.DBAdapter.MathComparision.EqalTo()).
		ExecSelect()

}
