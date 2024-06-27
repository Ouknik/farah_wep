package models

import (
	"database/sql"

	tableElements "app.roya_immobile.com/models/userapp"

	"app.roya_immobile.com/persistence"

	"app.roya_immobile.com/lib"
)

type ModelUserapp struct {
	DBAdapter               lib.DbAdapter
	userDetails             persistence.UserDetails
	_handleRedirectAndPanic *lib.HandleRedirectAndPanic
}

func (m *ModelUserapp) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m._handleRedirectAndPanic = _handleRedirectAndPanic
	m.DBAdapter.Connect("ry_user", "ryusr_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *ModelUserapp) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_user", "ryusr_")
	m._handleRedirectAndPanic = _handleRedirectAndPanic

}

// Returns db fiedls for frontend iteration
// []string
func (m *ModelUserapp) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_ry_user_app
}

// Returns db fiedls which are serachable
// []string
func (m *ModelUserapp) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_ry_user_app
}

func (m *ModelUserapp) GetTableElement_Id() string {
	return tableElements.Id
}

func (m *ModelUserapp) GetTableElement_Role() string {
	return tableElements.Role
}

func (m *ModelUserapp) GetTableElement_Email() string {
	return tableElements.Email
}

func (m *ModelUserapp) GetTableElement_Blocked() string {
	return tableElements.Blocked
}

func (m *ModelUserapp) GetTableElement_First_name() string {
	return tableElements.First_name
}

func (m *ModelUserapp) GetTableElement_Last_name() string {
	return tableElements.Last_name
}

func (m *ModelUserapp) GetTableElement_Mobile_number() string {
	return tableElements.Mobile_number
}

func (m *ModelUserapp) GetTableElement_Discription() string {
	return tableElements.Discription
}

// Type		: varchar(255)
// Null		: YES
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Avatar_filename = "zsusrapp_avatar_filename"
func (m *ModelUserapp) GetTableElement_Avatar() string {
	return tableElements.Avatar
}

func (m *ModelUserapp) FindUserByMobileNumber(mobileNumber string) map[string]string {
	m.DBAdapter.SetLimit(0, 1)
	return m.DBAdapter.SelectAll().
		Where(tableElements.Mobile_number, mobileNumber, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.Blocked, "0", m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()[0]
}

func (m *ModelUserapp) Insert(fields, values []string) (int64, string) {
	return m.DBAdapter.Insert(fields, values)
}

func (m *ModelUserapp) InsertUserDetails(userDetails persistence.UserDetails) (int64, string) {

	return m.DBAdapter.Insert([]string{
		//		tableElements.Salutation,
		tableElements.First_name,
		tableElements.Last_name,
		tableElements.Email,
		tableElements.Role,
		tableElements.Mobile_number,
	}, []string{
		//		string(userDetails.Salutation),
		userDetails.First_name,
		userDetails.Last_name,
		userDetails.Email,
		userDetails.Telephone,
		userDetails.Mobile_phone,
		userDetails.CountryCode,
		lib.GetUserRealIP(m._handleRedirectAndPanic.GetHttpRequest())})
}

func (m *ModelUserapp) UpdateUserById(userId string, fields, values []string) {
	m.DBAdapter.Update(fields, values).
		Where(tableElements.Id, userId, m.DBAdapter.Wildcard.NotWildcard()).
		ExecUpdate()
}

func (m *ModelUserapp) FindUserByEmail(email string) map[int]map[string]string {
	return m.DBAdapter.SelectAll().
		Where(tableElements.Email, email, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.Blocked, "0", m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()
}

func (m *ModelUserapp) FindUserById(id string) map[string]string {
	return m.DBAdapter.SelectFieldsAs(map[string]string{
		tableElements.Id:            "Id",
		tableElements.First_name:    "First_name",
		tableElements.Last_name:     "Last_name",
		tableElements.Email:         "Email",
		tableElements.Avatar:        "avatar",
		tableElements.Role:          "role",
		tableElements.Mobile_number: "Mobile_phone",
	}).WhereMathComparision(tableElements.Id, id, m.DBAdapter.MathComparision.EqalTo()).
		Where(tableElements.Blocked, "0", m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()[0]
}

func (m *ModelUserapp) FindUserByID(userid string) map[string]string {
	return m.DBAdapter.SelectAll().
		//WhereMathComparision(tableElements.Id, userid, persistence.MathComparision(m.DBAdapter.Wildcard)).
		Where(tableElements.Id, userid, m.DBAdapter.Wildcard.NotWildcard()).
		//Where(tableElements.Blocked, "0", m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()[0]
}

func (m *ModelUserapp) FindUserByIdIgnoreDeleteStatus(id string) map[string]string {
	return m.DBAdapter.SelectAll().
		//WhereMathComparision(tableElements.Id, id, m.DBAdapter.MathComparision.EqalTo()).
		Where(tableElements.Id, id, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()[0]
}

func (m *ModelUserapp) FindUserByMobilePhone(mobile_phone string) map[string]string {
	return m.DBAdapter.SelectAll().
		Where(tableElements.Mobile_number, mobile_phone, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()[0]
}

func (m *ModelUserapp) GetUserCredentialsByEmail(mobile_number string) map[string]string {

	return m.DBAdapter.SelectAll().
		RightJoin("zs_user_pwd", "zsusrapp_id", "zsusrpwd_uid").
		RightJoin("zs_user_salt", "MD5(zsusrapp_id)", "zsusrslt_uid").
		Where(tableElements.Mobile_number, mobile_number, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.Blocked, "0", m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()[0]
}

func (m *ModelUserapp) UpdateAvatar(userId string, avatarFileName string) {
	m.DBAdapter.Update([]string{
		tableElements.Avatar,
	}, []string{
		avatarFileName,
	}).Where(tableElements.Id, userId, m.DBAdapter.Wildcard.NotWildcard()).
		ExecUpdate()
}

func (m *ModelUserapp) UpdateGivenFieldsById(fieldsToUpdate, values []string, id string) {

	m.DBAdapter.Update(fieldsToUpdate, values).
		Where(tableElements.Id, id, m.DBAdapter.Wildcard.NotWildcard()).
		ExecUpdate()

}

func (m *ModelUserapp) GetAllUsers() map[int]map[string]string {
	return m.DBAdapter.SelectAll().
		Where(tableElements.Blocked, "0", m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()
}

func (m *ModelUserapp) UpdateUserInfo(fieldsForDB []string, valuesForDB []string, iduser string) {
	m.DBAdapter.Update(fieldsForDB, valuesForDB).
		Where(tableElements.Id, iduser, m.DBAdapter.Wildcard.NotWildcard()).
		ExecUpdate()
}
