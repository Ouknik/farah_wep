package models

import (
	"database/sql"

	"app.roya_immobile.com/lib"
	tableElements "app.roya_immobile.com/models/invoice"
)

type Modelinvoice struct {
	DBAdapter lib.DbAdapter
}

func (m *Modelinvoice) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m.DBAdapter.Connect("ry_invoices", "ryinvoc_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *Modelinvoice) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("ry_invoices", "ryinvoc_")
}

// Returns db fiedls for frontend iteration
// []string
func (m *Modelinvoice) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_ry_invoice
}

// Returns db fiedls which are serachable
// []string
func (m *Modelinvoice) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_ry_invoice
}

func (m *Modelinvoice) GetTableElement_Crdate() string {
	return tableElements.InvoiceCrdate
}

func (m *Modelinvoice) GetTableElement_Id() string {
	return tableElements.Id
}

func (m *Modelinvoice) GetTableElement_Prix() string {
	return tableElements.InvoicePrix
}

func (m *Modelinvoice) GetTableElement_ClienId() string {
	return tableElements.ClienId
}

func (m *Modelinvoice) GetTableElement_AnnonceId() string {
	return tableElements.AnnonceId
}

func (m *Modelinvoice) GetTableElement_InvoiceDate() string {
	return tableElements.InvoiceDate
}

func (m *Modelinvoice) GetTableElement_InvoiceMonth() string {
	return tableElements.InvoiceMonth
}

func (m *Modelinvoice) GetTableElement_InvoiceStatus() string {
	return tableElements.InvoiceStatus
}

func (m *Modelinvoice) AddIvoice(fields []string, values []string) (int64, string) {
	return m.DBAdapter.Insert(fields, values)
}

func (m *Modelinvoice) GetAllIvoices(month string) map[int]map[string]string {

	return m.DBAdapter.SelectAll().
		Where(tableElements.InvoiceMonth, month, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()
}

func (m *Modelinvoice) FindInvoceByClient(clienId string, month string) map[string]string {

	return m.DBAdapter.SelectAll().
		Where(tableElements.InvoiceMonth, month, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.ClienId, clienId, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()[0]
}

func (m *Modelinvoice) InsertInvoice(fildes, values []string) {

	m.DBAdapter.Insert(fildes, values)

}

func (m *Modelinvoice) GetIvoicesByClienAndMonth(_modelAnnonce ModelAnnonce, _modelUserapp ModelUserapp, _modelClient ModelClients, clientId string, month string) map[int]map[string]string {

	return m.DBAdapter.SelectAll().
		Where(tableElements.ClienId, clientId, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.InvoiceMonth, month, m.DBAdapter.Wildcard.NotWildcard()).
		RightJoin(_modelClient.DBAdapter.GetTableName(), _modelClient.GetTableElement_Id(), tableElements.ClienId).
		RightJoin(_modelAnnonce.DBAdapter.GetTableName(), _modelAnnonce.GetTableElement_Id(), tableElements.AnnonceId).
		RightJoin(_modelUserapp.DBAdapter.GetTableName(), _modelUserapp.GetTableElement_Id(), _modelClient.GetTableElement_UserId()).
		ExecSelect()
}

func (m *Modelinvoice) GetAllIvoicesBuClien(_modelAnnonce ModelAnnonce, _modelUserapp ModelUserapp, _modelClient ModelClients, clientId string) map[int]map[string]string {

	return m.DBAdapter.SelectAll().
		Where(tableElements.ClienId, clientId, m.DBAdapter.Wildcard.NotWildcard()).
		RightJoin(_modelClient.DBAdapter.GetTableName(), _modelClient.GetTableElement_Id(), tableElements.ClienId).
		RightJoin(_modelAnnonce.DBAdapter.GetTableName(), _modelAnnonce.GetTableElement_Id(), tableElements.AnnonceId).
		RightJoin(_modelUserapp.DBAdapter.GetTableName(), _modelUserapp.GetTableElement_Id(), _modelClient.GetTableElement_UserId()).
		ExecSelect()
}

func (m *Modelinvoice) GetAllIvoicesBuMonth(_modelAnnonce ModelAnnonce, _modelUserapp ModelUserapp, _modelClient ModelClients, month string) map[int]map[string]string {

	return m.DBAdapter.SelectAll().
		Where(tableElements.InvoiceMonth, month, m.DBAdapter.Wildcard.NotWildcard()).
		RightJoin(_modelClient.DBAdapter.GetTableName(), _modelClient.GetTableElement_Id(), tableElements.ClienId).
		RightJoin(_modelAnnonce.DBAdapter.GetTableName(), _modelAnnonce.GetTableElement_Id(), tableElements.AnnonceId).
		RightJoin(_modelUserapp.DBAdapter.GetTableName(), _modelUserapp.GetTableElement_Id(), _modelClient.GetTableElement_UserId()).
		ExecSelect()
}
