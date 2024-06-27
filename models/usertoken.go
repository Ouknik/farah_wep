package models

import (
	"database/sql"

	tableElements "app.roya_immobile.com/models/usertoken"

	"app.roya_immobile.com/persistence"

	"app.roya_immobile.com/lib"
)

type ModelUserToken struct {
	DBAdapter               lib.DbAdapter
	userDetails             persistence.UserDetails
	_handleRedirectAndPanic *lib.HandleRedirectAndPanic
}

func (m *ModelUserToken) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m._handleRedirectAndPanic = _handleRedirectAndPanic
	m.DBAdapter.Connect("ry_user_token", "ryustkn_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *ModelUserToken) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_user_token", "ryustkn_")
	m._handleRedirectAndPanic = _handleRedirectAndPanic

}

// Returns db fiedls for frontend iteration
// []string
func (m *ModelUserToken) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_ry_user_token
}

// Returns db fiedls which are serachable
// []string
func (m *ModelUserToken) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_ry_user_token
}

func (m *ModelUserToken) GetTableElement_Id() string {
	return tableElements.Id
}

func (m *ModelUserToken) GetTableElement_UserID() string {
	return tableElements.UserId
}

func (m *ModelUserToken) GetTableElement_Token() string {
	return tableElements.Token
}

func (m *ModelUserToken) GetUnusedQrCodeByUserId(userId string) map[int]map[string]string {
	return m.DBAdapter.SelectAll().
		WhereMathComparision(tableElements.UserId, userId, m.DBAdapter.MathComparision.EqalTo()).
		//	WhereMathComparision(tableElements.Statue, "0", m.DBAdapter.MathComparision.EqalTo()).
		ExecSelect()
}

func (m *ModelUserToken) InsertUserToken(userId, userToken string) (int64, string) {
	return m.DBAdapter.Insert(
		[]string{
			tableElements.UserId,
			tableElements.Token,
		},
		[]string{
			userId,
			userToken,
		})
}

func (m *ModelUserToken) GetUserByQrcodeAndMobileNumber(token string, mobile_number string, _modelUser *ModelUserapp) map[string]string {
	return m.DBAdapter.SelectFieldsAs(map[string]string{
		_modelUser.GetTableElement_Id():            "Id",
		_modelUser.GetTableElement_First_name():    "First_name",
		_modelUser.GetTableElement_Last_name():     "Last_name",
		_modelUser.GetTableElement_Email():         "Email",
		_modelUser.GetTableElement_Avatar():        "avatar",
		_modelUser.GetTableElement_Role():          "Role",
		_modelUser.GetTableElement_Mobile_number(): "Mobile_phone",
		_modelUser.GetTableElement_Discription():   "Discription",
	}).
		RightJoin(_modelUser.DBAdapter.GetTableName(), tableElements.UserId, _modelUser.GetTableElement_Id()).
		Where(_modelUser.GetTableElement_Mobile_number(), mobile_number, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.Token, token, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()[0]
}
