package models

import (
	"database/sql"

	"app.roya_immobile.com/lib"
	tableElements "app.roya_immobile.com/models/annonce"
)

type ModelAnnonce struct {
	DBAdapter lib.DbAdapter
}

func (m *ModelAnnonce) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m.DBAdapter.Connect("ry_annonce", "royaannonse_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *ModelAnnonce) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_annonce", "royaannonse_")
}

// Returns db fiedls for frontend iteration
// []string
func (m *ModelAnnonce) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_ry_annonce
}

// Returns db fiedls which are serachable
// []string
func (m *ModelAnnonce) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_ry_annonce
}

// Globalise DB Fields

// Type		: timestamp
// Null		: NO
// Key		:
// Default	: current_timestamp()
// Extra	: on update current_timestamp()
// Used_on = "zsapiot_used_on"
func (m *ModelAnnonce) GetTableElement_Id() string {
	return tableElements.Id
}

func (m *ModelAnnonce) GetTableElement_Category() string {
	return tableElements.Category
}

func (m *ModelAnnonce) GetTableElement_UserId() string {
	return tableElements.UserId
}

func (m *ModelAnnonce) GetTableElement_Region() string {
	return tableElements.Region
}

func (m *ModelAnnonce) GetTableElement_City() string {
	return tableElements.City
}

func (m *ModelAnnonce) GetTableElement_Transaction() string {
	return tableElements.Transaction
}

func (m *ModelAnnonce) GetTableElement_PropertyType() string {
	return tableElements.PropertyType
}

func (m *ModelAnnonce) GetTableElement_Status() string {
	return tableElements.Status
}

func (m *ModelAnnonce) GetTableElement_Address() string {
	return tableElements.Address
}

func (m *ModelAnnonce) GetTableElement_Quartier() string {
	return tableElements.Quartier
}

func (m *ModelAnnonce) GetTableElement_Area() string {
	return tableElements.Area
}

func (m *ModelAnnonce) GetTableElement_Price() string {
	return tableElements.Price
}

func (m *ModelAnnonce) GetTableElement_Age() string {
	return tableElements.Age
}

func (m *ModelAnnonce) GetTableElement_FloorType() string {
	return tableElements.FloorType
}

func (m *ModelAnnonce) GetTableElement_Floor() string {
	return tableElements.Floor
}

func (m *ModelAnnonce) GetTableElement_Apartment() string {
	return tableElements.Apartment
}

func (m *ModelAnnonce) GetTableElement_Bedrooms() string {
	return tableElements.Bedrooms
}

func (m *ModelAnnonce) GetTableElement_Bathrooms() string {
	return tableElements.Bathrooms
}

func (m *ModelAnnonce) GetTableElement_Kitchens() string {
	return tableElements.Kitchens
}

func (m *ModelAnnonce) GetTableElement_Title() string {
	return tableElements.Title
}

func (m *ModelAnnonce) GetTableElement_Description() string {
	return tableElements.Description
}

func (m *ModelAnnonce) GetTableElement_Phone1() string {
	return tableElements.Phone1
}

func (m *ModelAnnonce) GetTableElement_Phone2() string {
	return tableElements.Phone2
}

func (m *ModelAnnonce) GetTableElement_Phone3() string {
	return tableElements.Phone3
}

func (m *ModelAnnonce) GetAllByCityForView(mycity string, _modelUserapp ModelUserapp) map[int]map[string]string {
	return m.DBAdapter.SelectAll().
		RightJoin(_modelUserapp.DBAdapter.GetTableName(), tableElements.UserId, _modelUserapp.GetTableElement_Id()).
		Where(tableElements.City, mycity, m.DBAdapter.Wildcard.NotWildcard()).
		//Where(tableElements.Is_delete, "0", m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()
}

func (m *ModelAnnonce) GetAnnonceByUserIdForView(userid string, _modelUserapp ModelUserapp) map[int]map[string]string {
	return m.DBAdapter.SelectAll().
		Where(tableElements.UserId, userid, m.DBAdapter.Wildcard.NotWildcard()).
		RightJoin(_modelUserapp.DBAdapter.GetTableName(), tableElements.UserId, _modelUserapp.GetTableElement_Id()).

		//Where(tableElements.Is_delete, "0", m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()
}

func (m *ModelAnnonce) GetAllAnnonceForView() map[int]map[string]string {
	return m.DBAdapter.SelectAll().
		//Where(tableElements.Is_delete, "0", m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()
}

func (m *ModelAnnonce) FindAnnonceByID(annonceId string) map[string]string {
	return m.DBAdapter.SelectAll().
		Where(tableElements.Id, annonceId, m.DBAdapter.Wildcard).
		ExecSelect()[0]
}

func (m *ModelAnnonce) DeleteAnnonce(userId string, annonceId string) int {
	var _timeHandler = lib.TimeHandler{}
	return m.DBAdapter.Update([]string{tableElements.Deleted_by, tableElements.Deleted_on, tableElements.Is_delete}, []string{userId, _timeHandler.Now(), "1"}).
		Where(tableElements.Id, annonceId, m.DBAdapter.Wildcard).
		ExecUpdate()
}

func (m *ModelAnnonce) InserAnnonce(fields []string, values []string) (int64, string) {

	return m.DBAdapter.Insert(fields, values)
}
