package models

import (
	"database/sql"

	tableElements "app.roya_immobile.com/models/annonce_message" // تأكد من أن المسار صحيح

	"app.roya_immobile.com/lib"
)

type ModelAnnnceMessage struct {
	DBAdapter lib.DbAdapter
}

func (m *ModelAnnnceMessage) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m.DBAdapter.Connect("ry_annonce_message", "ryannmess_", _handleRedirectAndPanic)
}

func (m *ModelAnnnceMessage) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_annonce_message", "ryannmess_")
}

func (m *ModelAnnnceMessage) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIterationAnnonceMessage // تصحيح اسم المتغير
}

func (m *ModelAnnnceMessage) GetSearchableDbFields() []string {
	return tableElements.SearchableFieldsAnnonceMessage // تصحيح اسم المتغير
}

func (m *ModelAnnnceMessage) GetTableElement_Id() string {
	return tableElements.MessageId
}

func (m *ModelAnnnceMessage) GetTableElement_UserID() string {
	return tableElements.UserId
}

func (m *ModelAnnnceMessage) GetTableElement_AdminID() string {
	return tableElements.AdminId
}

func (m *ModelAnnnceMessage) GetTableElement_AnnounceID() string {
	return tableElements.AnnonceId
}

func (m *ModelAnnnceMessage) GetTableElement_Message() string {
	return tableElements.Message
}

func (m *ModelAnnnceMessage) GetTableElement_CreationDate() string {
	return tableElements.CreationDate
}

func (m *ModelAnnnceMessage) GetTableElement_CreatedBy() string {
	return tableElements.CreatedBy
}

func (m *ModelAnnnceMessage) GetTableElement_SeeMessageAt() string {
	return tableElements.SeeMessageAt
}
func (m *ModelAnnnceMessage) GetTableElement_Status() string {
	return tableElements.Status
}

func (m *ModelAnnnceMessage) GetAllAnnonceMessages(userId string, adminId string) map[int]map[string]string {
	m.DBAdapter.SetOrderByField(tableElements.MessageId, m.DBAdapter.OrderBy.Asc())
	return m.DBAdapter.SelectAll().
		Where(tableElements.UserId, userId, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.AdminId, adminId, m.DBAdapter.Wildcard.NotWildcard()).
		//Where(tableElements.AnnonceId, annonceId, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()
}

func (m *ModelAnnnceMessage) SeeMessage(messageId string, userId string, adminId, fields []string, values []string) {

	m.DBAdapter.Update(fields, values).
		//Where(tableElements.UserId, userId, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.MessageId, messageId, m.DBAdapter.Wildcard.NotWildcard()).
		//Where(tableElements.AdminId, adminId, m.DBAdapter.Wildcard.NotWildcard()).
		//Where(tableElements.AnnonceId, annonceId, m.DBAdapter.Wildcard.NotWildcard())
		ExecUpdate()

}

func (m *ModelAnnnceMessage) UpdateById(messageId string, userId string, adminId string, fields, values []string) int {

	return m.DBAdapter.Update(fields, values).
		Where(tableElements.MessageId, messageId, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.AdminId, adminId, m.DBAdapter.Wildcard.NotWildcard()).
		//Where(tableElements.AnnonceId, annonceId, m.DBAdapter.Wildcard.NotWildcard()).
		ExecUpdate()
}

func (m *ModelAnnnceMessage) GetAllOrderByAdmin(adminId string, _modelUserApp ModelUserapp) map[int]map[string]string {
	m.DBAdapter.SetOrderByField(tableElements.MessageId, m.DBAdapter.OrderBy.Asc())
	return m.DBAdapter.SelectAll().
		RightJoin(_modelUserApp.DBAdapter.GetTableName(), _modelUserApp.GetTableElement_Id(), tableElements.UserId).
		Where(tableElements.AdminId, adminId, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()
}

func (m *ModelAnnnceMessage) GetMessageDetille(adminId string, userId string, _modelAnnonce ModelAnnonce, _modelUserApp ModelUserapp) map[int]map[string]string {
	m.DBAdapter.SetOrderByField(tableElements.MessageId, m.DBAdapter.OrderBy.Asc())
	return m.DBAdapter.SelectAll().
		RightJoin(_modelUserApp.DBAdapter.GetTableName(), _modelUserApp.GetTableElement_Id(), tableElements.UserId).
		RightJoin(_modelAnnonce.DBAdapter.GetTableName(), _modelAnnonce.GetTableElement_Id(), tableElements.AnnonceId).
		Where(tableElements.AdminId, adminId, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.UserId, userId, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()
}

func (m *ModelAnnnceMessage) GetNumberNotSeeMessage(message map[string]string) int {

	return len(m.DBAdapter.SelectAll().
		Where(tableElements.AdminId, message["Message_admin_id"], m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.UserId, message["Message_user_id"], m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.CreatedBy, message["Message_user_id"], m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.Status, "0", m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect())

}
