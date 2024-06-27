package models

import (
	"database/sql"

	tableElements "app.roya_immobile.com/models/verificationsms"

	"app.roya_immobile.com/persistence"

	"app.roya_immobile.com/lib"
)

type ModelVerificationSms struct {
	DBAdapter               lib.DbAdapter
	userDetails             persistence.UserDetails
	_handleRedirectAndPanic *lib.HandleRedirectAndPanic
}

func (m *ModelVerificationSms) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m._handleRedirectAndPanic = _handleRedirectAndPanic
	m.DBAdapter.Connect("ry_sms_verification", "ryusrvs_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *ModelVerificationSms) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_sms_verification", "ryusrvs_")
	m._handleRedirectAndPanic = _handleRedirectAndPanic

}

// Returns db fiedls for frontend iteration
// []string
func (m *ModelVerificationSms) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_ry_verification_sms
}

// Returns db fiedls which are serachable
// []string
func (m *ModelVerificationSms) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_ry_verification_sms
}

func (m *ModelVerificationSms) GetTableElement_Verified_on() string {
	return tableElements.Verified_on
}

func (m *ModelVerificationSms) GetTableElement_Verification_code() string {
	return tableElements.Verification_code
}

func (m *ModelVerificationSms) GetTableElement_Crdate() string {
	return tableElements.Crdate
}
func (m *ModelVerificationSms) GetTableElement_Mobile_number() string {
	return tableElements.Mobile_number
}

func (m *ModelVerificationSms) GetTableElement_Statue() string {
	return tableElements.Statue
}

func (m *ModelVerificationSms) InsertVerificationCodeWithDeviceId(mobileNumber string, verificationCode string) (int64, string) {
	return m.DBAdapter.Insert(
		[]string{
			tableElements.Mobile_number,

			tableElements.Verification_code,
		},
		[]string{
			mobileNumber,

			verificationCode,
		})
}

func (m *ModelVerificationSms) FindUserByMobileNumberAndVerificationCode(mobileNumber, verificationCode string) map[int]map[string]string {
	return m.DBAdapter.SelectAll().
		//Where(tableElements.Statue, "N", m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.Mobile_number, mobileNumber, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.Verification_code, verificationCode, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()
}

func (m *ModelVerificationSms) ValidateSmsByDeviceIdMobileNumberAndSms(mobileNumber, verificationCode string) {
	var _time = lib.TimeHandler{}
	var now = _time.Now()
	m.DBAdapter.Update(
		[]string{
			tableElements.Statue,
			tableElements.Verified_on,
		},
		[]string{
			"1",
			now,
		}).
		Where(tableElements.Mobile_number, mobileNumber, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.Verification_code, verificationCode, m.DBAdapter.Wildcard.NotWildcard()).
		ExecUpdate()
}

func (m *ModelVerificationSms) FindUserByMobileNumberDeviceIdAndVerificationCode(mobileNumber, verificationCode string) map[int]map[string]string {
	return m.DBAdapter.SelectAll().
		Where(tableElements.Statue, "0", m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.Mobile_number, mobileNumber, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.Verification_code, verificationCode, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()
}
