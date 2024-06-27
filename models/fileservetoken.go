package models

import (
	"database/sql"
	"errors"

	tableElements "app.roya_immobile.com/models/fileservetoken"

	"app.roya_immobile.com/persistence"

	"app.roya_immobile.com/lib"
)

type ModelFileservetoken struct {
	DBAdapter lib.DbAdapter
}

func (m *ModelFileservetoken) Init(_handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter = lib.DbAdapter{}
	m.DBAdapter.Connect("zs_file_serve_token", "zsflsrvtkn_", _handleRedirectAndPanic)
}

// Extend the exisisting DBAdapter to use sub tables
// on the main DB Connected thread
// @ param db *sql.DB
// @ return void
func (m *ModelFileservetoken) InitWithoutConnection(db *sql.DB, _handleRedirectAndPanic *lib.HandleRedirectAndPanic) {
	m.DBAdapter.SetHandleRedirectAndPanic(_handleRedirectAndPanic)
	m.DBAdapter.SetSqlConnection(db)
	m.DBAdapter.SetTableAndPrefix("zs_file_serve_token", "zsflsrvtkn_")
}

// Returns db fiedls for frontend iteration
// []string
func (m *ModelFileservetoken) GetDbFieldsAsKeysForIteration() []string {
	return tableElements.KeysForIteration_zs_file_serve_token
}

// Returns db fiedls which are serachable
// []string
func (m *ModelFileservetoken) GetSearchableDbFields() []string {
	return tableElements.SearchableFields_zs_file_serve_token
}

// Globalise DB Fields

// Type		: tinyint(4)
// Null		: NO
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Destination = "zsflsrvtkn_destination"
func (m *ModelFileservetoken) GetTableElement_Destination() string {
	return tableElements.Destination
}

// Type		: varchar(20)
// Null		: YES
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// File_size = "zsflsrvtkn_file_size"
func (m *ModelFileservetoken) GetTableElement_File_size() string {
	return tableElements.File_size
}

// Type		: int(20)
// Null		: NO
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Created_by = "zsflsrvtkn_created_by"
func (m *ModelFileservetoken) GetTableElement_Created_by() string {
	return tableElements.Created_by
}

// Type		: timestamp
// Null		: NO
// Key		:
// Default	: current_timestamp()
// Extra	:
// Crdate = "zsflsrvtkn_crdate"
func (m *ModelFileservetoken) GetTableElement_Crdate() string {
	return tableElements.Crdate
}

// Type		: varchar(32)
// Null		: NO
// Key		: PRI
// Default	: %!s(<nil>)
// Extra	:
// Token = "zsflsrvtkn_token"
func (m *ModelFileservetoken) GetTableElement_Token() string {
	return tableElements.Token
}

// Type		: varchar(255)
// Null		: NO
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// Portal_key = "zsflsrvtkn_portal_key"
func (m *ModelFileservetoken) GetTableElement_Portal_key() string {
	return tableElements.Portal_key
}

// Type		: tinytext
// Null		: YES
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// File_name = "zsflsrvtkn_file_name"
func (m *ModelFileservetoken) GetTableElement_File_name() string {
	return tableElements.File_name
}

// Type		: varchar(20)
// Null		: YES
// Key		:
// Default	: %!s(<nil>)
// Extra	:
// File_type = "zsflsrvtkn_file_mime_type"
func (m *ModelFileservetoken) GetTableElement_File_mime_type() string {
	return tableElements.File_mime_type
}

func (m *ModelFileservetoken) InsertFileDetails(_persistenceCloudToken persistence.CloudToken) (int64, string) {
	var fields = []string{
		tableElements.Token,
		tableElements.Portal_key,
		tableElements.File_name,
		tableElements.File_mime_type,
		tableElements.File_size,
		tableElements.Destination,
		tableElements.Created_by,
	}
	var values = []string{
		_persistenceCloudToken.Token,
		_persistenceCloudToken.PortalKey,
		_persistenceCloudToken.FileName,
		_persistenceCloudToken.FileMimeType,
		_persistenceCloudToken.FileSize,
		_persistenceCloudToken.Destination,
		_persistenceCloudToken.CreatedBy,
	}

	return m.DBAdapter.Insert(fields, values)
}

func (m *ModelFileservetoken) FindByFileTokenForDelete(portalKey, fileName string) (string, error) {
	// m.DBAdapter.SetLimit(0, 1)
	var result = m.DBAdapter.SelectAll().
		Where(tableElements.Portal_key, portalKey, m.DBAdapter.Wildcard.NotWildcard()).
		Where(tableElements.File_name, fileName, m.DBAdapter.Wildcard.NotWildcard()).
		ExecSelect()[0]
	if len(result) != 0 {
		return lib.StringReplace(result[tableElements.Token], "__", "", -1), nil
	}
	return "", errors.New("file not found")
}
