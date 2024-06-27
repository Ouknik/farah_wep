package helpers

import (
	"app.roya_immobile.com/lib"
	"app.roya_immobile.com/persistence"
	"fmt"
	"net/http"
)

type HelperPrepareSearch struct {
	request                               *http.Request
	_dbAdapter                            lib.DbAdapter
	_handleSession                        *lib.HandleSession
	_handleRedirectAndPanic               *lib.HandleRedirectAndPanic
	dbTable                               string
	dbTableFieldPrefix                    string
	searchableDbFields                    []string
	searchableDbFieldsToIgnoreForWildcard []string
	additionalSelectAsFields              []string
	searchParamsAsUrlString               string
	rightJoins                            []_struct_joinParams
	leftJoins                             []_struct_joinParams
	baseWhereClause                       []whereClauseParams
	baseWhereOrClause                     []whereClauseParams
	baseWhereOrMathClause                 []whereMathClauseParams
	searchParamsForView                   map[string]string
	itemsPerPage                          int
	rowsInTotal                           int
	rowsRetrieved                         int
	pageNumberPresent                     int
	pageNumberTotal                       int
	queryLimitFrom                        int
	queryLimitUntil                       int
	orderBy                               string
	orderByCookieKey                      string
	orderByTypeCookieKey                  string
	orderType                             persistence.OrderBy
	groupBy                               []string
	controllerDetails                     persistence.ControllerDetails
	_envconfVars                          *lib.GetEnvConfigVars
}

type whereClauseParams struct {
	dbFieldToSearch string
	searchKey       string
	searchValue     string
	wildcard        persistence.Wildcard
}

type whereMathClauseParams struct {
	dbField         string
	value           string
	mathComparision persistence.MathComparision
}

type _struct_joinParams struct {
	tableToJoin string
	fieldOn     string
	fieldTo     string
}

func (h *HelperPrepareSearch) SetEnvironmentalParams(controllerDetails persistence.ControllerDetails, _handleRedirectAndPanic *lib.HandleRedirectAndPanic, _handleSessison *lib.HandleSession, _dbAdapter lib.DbAdapter, dbTable, dbTablePrefix string, _envconfVars *lib.GetEnvConfigVars) *lib.DbAdapter {
	h.controllerDetails = controllerDetails
	h._handleRedirectAndPanic = _handleRedirectAndPanic
	h._handleSession = _handleSessison
	h._dbAdapter = _dbAdapter
	h._dbAdapter.SetTableAndPrefix(dbTable, dbTablePrefix)
	h.dbTable = dbTable
	h.dbTableFieldPrefix = dbTablePrefix
	h._envconfVars = _envconfVars
	h.orderByCookieKey = fmt.Sprintf("_orderBy_%s%s", h.controllerDetails.Controller, h.controllerDetails.Action)
	h.orderByTypeCookieKey = fmt.Sprintf("_orderByType_%s%s", h.controllerDetails.Controller, h.controllerDetails.Action)
	h.setItemsPerPage()
	h.searchParamsForView = map[string]string{}
	return &h._dbAdapter
}

func (h *HelperPrepareSearch) SetHTTPRequest(request *http.Request) {
	h.request = request
}

func (h *HelperPrepareSearch) SetSearchableDbFields(searchableDBFields []string) {
	h.searchableDbFields = searchableDBFields
}

func (h *HelperPrepareSearch) SetSearchableDbFieldsToIgnoreForWildcard(searchableDBFieldsToIgnore []string) {
	h.searchableDbFieldsToIgnoreForWildcard = searchableDBFieldsToIgnore
}

func (h *HelperPrepareSearch) setJoinStruct(tableToJoin, fieldOn, fieldTo string) _struct_joinParams {
	return _struct_joinParams{
		tableToJoin: tableToJoin,
		fieldOn:     fieldOn,
		fieldTo:     fieldTo,
	}
}

func (h *HelperPrepareSearch) SetRightJoin(tableToJoin, fieldOn, fieldTo string) {
	h.rightJoins = append(h.rightJoins, h.setJoinStruct(tableToJoin, fieldOn, fieldTo))
}

func (h *HelperPrepareSearch) SetLeftJoin(tableToJoin, fieldOn, fieldTo string) {
	h.leftJoins = append(h.leftJoins, h.setJoinStruct(tableToJoin, fieldOn, fieldTo))
}

func (h *HelperPrepareSearch) SetBaseWhereClause(dbField, dbValue string, wildcard persistence.Wildcard) {
	h.baseWhereClause = append(h.baseWhereClause, whereClauseParams{dbFieldToSearch: dbField, searchValue: dbValue, wildcard: wildcard})
}

func (h *HelperPrepareSearch) SetBaseWhereOrClause(dbField, dbValue string, wildcard persistence.Wildcard) {
	h.baseWhereOrClause = append(h.baseWhereOrClause, whereClauseParams{dbFieldToSearch: dbField, searchValue: dbValue, wildcard: wildcard})
}

func (h *HelperPrepareSearch) SetBaseWhereOrMathClause(dbField, dbValue string, mathComparision persistence.MathComparision) {
	h.baseWhereOrMathClause = append(h.baseWhereOrMathClause, whereMathClauseParams{dbField: dbField, value: dbValue, mathComparision: mathComparision})
}

func (h *HelperPrepareSearch) SetOrderBy(orderByField string, orderByType persistence.OrderBy) {
	h.orderBy = orderByField
	h.orderType = orderByType
}

func (h *HelperPrepareSearch) SetGroupByField(groupBy string) {
	h.groupBy = append(h.groupBy, groupBy)
}

func (h *HelperPrepareSearch) SetAdditionalSelectAsField(selectAsField string) {
	h.additionalSelectAsFields = append(h.additionalSelectAsFields, selectAsField)
}

func (h *HelperPrepareSearch) GetPaginationDetails() persistence.PaginationDetails {

	var paginationDetails = persistence.PaginationDetails{
		RowsInTotal:  h.rowsInTotal,
		ItemsPerPage: h.itemsPerPage,
		ItemFrom:     (h.queryLimitFrom + 1),
		// dont return h.queryLimitUntil as it might defer
		// hence return (h.queryLimitFrom + h.rowsRetrieved)
		ItemUntil:               h.queryLimitFrom + h.rowsRetrieved,
		PagesTotal:              h.pageNumberTotal,
		PagePresent:             h.pageNumberPresent,
		SearchParamsAsUrlString: h.searchParamsAsUrlString,
		SearchParams:            h.searchParamsForView,
	}

	h.makePageNext(&paginationDetails)
	h.makePagePrevious(&paginationDetails)

	return paginationDetails

}

func (h *HelperPrepareSearch) makePageNext(paginationDetails *persistence.PaginationDetails) *persistence.PaginationDetails {
	if h.pageNumberPresent < h.pageNumberTotal {
		paginationDetails.PageNext = (h.pageNumberPresent + 1)
	} else {
		paginationDetails.PageNext = 0
	}
	return paginationDetails
}

func (h *HelperPrepareSearch) makePagePrevious(paginationDetails *persistence.PaginationDetails) *persistence.PaginationDetails {
	if h.pageNumberPresent > 1 {
		paginationDetails.PagePrevious = (h.pageNumberPresent - 1)
	} else {
		paginationDetails.PagePrevious = 0
	}
	return paginationDetails
}

func (h *HelperPrepareSearch) MakeSearchQuery() map[int]map[string]string {

	// override default order by and look
	// if the user has his sorting prefernces
	h.handleOrderBy()

	// first look for the search type param
	// if it is set, get the type of search
	// ....
	// is it a wild card search wc = wildcard
	// or a full search fs = full search
	// or is it empty or not set
	if h.request.FormValue("_typ") == "wc" {
		return h.handleWildcardSearch()
	} else if h.request.FormValue("_typ") == "fs" {
		return h.handleFullSearch()
	} else if h.request.FormValue("_typ") == "" {
		return h.handleList()
	}

	return nil

}

func (h *HelperPrepareSearch) handleOrderBy() {
	//fmt.Println(h.request.FormValue("_orderby"))

	// first if look if there was a forign field without prefix set in searchable fields
	// second if look if there was a native field, with prefix, set in searchable fields
	if h.request.FormValue("_orderby") != "" && lib.InArray(h.searchableDbFields, h.request.FormValue("_orderby")) {
		h.setOrderbyFromURL(true)
		return
	} else if h.request.FormValue("_orderby") != "" && h.concatUrlKeyAndDbPrefix(h.request.FormValue("_orderby")) != "" {
		h.setOrderbyFromURL(false)
		return
	}

	if h._handleSession.GetSessionCookie(h.orderByCookieKey) != "" {
		h.setOrderByFromCookie()
		return
	}
}

func (h *HelperPrepareSearch) setOrderbyFromURL(skipPrefix bool) {

	orderbyType := h._dbAdapter.OrderBy.Asc()
	orderField := ""
	if h.request.FormValue("_orderbytyp") == "desc" {
		orderbyType = h._dbAdapter.OrderBy.Desc()
	}

	// skip the db prefix the field is a forign field (ex. COUNT zsusrrsnt_comments AS Comments)
	if skipPrefix {
		orderField = h.request.FormValue("_orderby")
	} else {
		orderField = h.concatUrlKeyAndDbPrefix(h.request.FormValue("_orderby"))
	}

	h.SetOrderBy(
		orderField,
		orderbyType)

	//	fmt.Println(h.concatUrlKeyAndDbPrefix(h.request.FormValue("_orderby")), string(orderbyType))

	h.setOrderByToCookie(orderField, string(orderbyType))

}

func (h *HelperPrepareSearch) setOrderByToCookie(orderBy, orderByType string) {
	h._handleSession.SetSessionCookie(h.orderByCookieKey, orderBy)
	h._handleSession.SetSessionCookie(h.orderByTypeCookieKey, orderByType)
}

func (h *HelperPrepareSearch) setOrderByFromCookie() {
	h.orderBy = h._handleSession.GetSessionCookie(h.orderByCookieKey)
	h.orderType = persistence.OrderBy(h._handleSession.GetSessionCookie(h.orderByTypeCookieKey))
}

func (h *HelperPrepareSearch) handleList() map[int]map[string]string {

	//	h.rowsInTotal = h.execSelectCount(h.buildBaseQuery(h._dbAdapter.SelectCount(nil), false))
	h.rowsInTotal = h.execSelectCount(h.buildBaseQuery(h._dbAdapter.SelectAll(), false))

	//	fmt.Println(h._dbAdapter.GetLastExecutedQuery())

	if !h.setPagination() {
		return nil
	}

	return h.selectExecAndSetRowsRetrieved(h.buildBaseQuery(h._dbAdapter.SelectAll(), true))
}

func (h *HelperPrepareSearch) handleWildcardSearch() map[int]map[string]string {

	if "" == h.request.FormValue("search") {
		return nil
	}

	var whereClauses = h.prepareSearchFieldsFromDbFields(h.request.FormValue("search"))
	//	fmt.Println(whereClauses)
	if whereClauses == nil {
		return nil
	}

	// set search params as url string
	h.setSearchParamToUrlString("_typ=wc&search", h.request.FormValue("search"))
	h.setSearchParamForView("search", h.request.FormValue("search"))

	//h.rowsInTotal = h.execSelectCount(h.makeWhereOrClause(h.buildBaseQuery(h._dbAdapter.SelectCount(nil), false), whereClauses, h._dbAdapter.Wildcard.Wildcard(), false))
	h.rowsInTotal = h.execSelectCount(h.makeWhereOrClause(h.buildBaseQuery(h._dbAdapter.SelectAll(), false), whereClauses, h._dbAdapter.Wildcard.Wildcard(), false))
	//	fmt.Println(h._dbAdapter.GetLastExecutedQuery())
	if !h.setPagination() {
		return nil
	}

	return h.selectExecAndSetRowsRetrieved(h.makeWhereOrClause(h.buildBaseQuery(h._dbAdapter.SelectAll(), true), whereClauses, h._dbAdapter.Wildcard.Wildcard(), false))
}

func (h *HelperPrepareSearch) handleFullSearch() map[int]map[string]string {

	var whereClauses = h.prepareSearchFieldsFromUrlQuery()
	if whereClauses == nil {
		return nil
	}

	// set search params as url string
	h.setSearchParamToUrlString("_typ", "fs")

	rowsInNotWildcard, rowsInWildcard := 0, 0

	//	rowsInNotWildcard = h.execSelectCount(h.makeWhereClause(h.buildBaseQuery(h._dbAdapter.SelectCount(nil), false), whereClauses, h._dbAdapter.Wildcard.NotWildcard(), false))
	//	rowsInWildcard = h.execSelectCount(h.makeWhereClause(h.buildBaseQuery(h._dbAdapter.SelectCount(nil), false), whereClauses, h._dbAdapter.Wildcard.Wildcard(), true))

	rowsInNotWildcard = h.execSelectCount(h.makeWhereClause(h.buildBaseQuery(h._dbAdapter.SelectAll(), false), whereClauses, h._dbAdapter.Wildcard.NotWildcard(), false))
	rowsInWildcard = h.execSelectCount(h.makeWhereClause(h.buildBaseQuery(h._dbAdapter.SelectAll(), false), whereClauses, h._dbAdapter.Wildcard.Wildcard(), true))
	h.rowsInTotal = rowsInNotWildcard + rowsInWildcard

	rowsRetrievedFromFullSearch := 0

	if !h.setPagination() {
		return nil
	}

	seekWildcardResults := true

	res := map[int]map[string]string{}

	// There are more item than this page or same number of items in the present page
	if h.queryLimitFrom <= rowsInNotWildcard {

		res = h.selectExecAndSetRowsRetrieved(h.makeWhereClause(h.buildBaseQuery(h._dbAdapter.SelectAll(), true), whereClauses, h._dbAdapter.Wildcard.NotWildcard(), false))

		rowsRetrievedFromFullSearch = h.rowsRetrieved

		if (h.queryLimitFrom + h.queryLimitUntil) <= rowsInNotWildcard {
			seekWildcardResults = false
		}

	}

	if seekWildcardResults {
		h.calculateFullWildcardQueryLimit(rowsInNotWildcard, rowsRetrievedFromFullSearch)
		var wildcardRes = h.selectExecAndSetRowsRetrieved(h.makeWhereClause(h.buildBaseQuery(h._dbAdapter.SelectAll(), true), whereClauses, h._dbAdapter.Wildcard.Wildcard(), true))

		// for the view: increment the querylimitfrom with previous query count
		h.queryLimitFrom += rowsInNotWildcard - rowsRetrievedFromFullSearch

		var iteratorLimit = rowsRetrievedFromFullSearch + h.rowsRetrieved
		j := 0
		for i := rowsRetrievedFromFullSearch; i < iteratorLimit; i++ {
			res[i] = wildcardRes[j]
			j++
		}

	}

	return res
}

func (h *HelperPrepareSearch) execSelectCount(query *lib.DbAdapter) int {
	//	return lib.StringToInt(query.ExecSelect()[0]["count"])
	return len(query.ExecSelect())
}

func (h *HelperPrepareSearch) prepareSearchFieldsFromUrlQuery() []whereClauseParams {

	var urlQuery = h.request.URL.Query()
	fieldKey := ""

	var whereClauses []whereClauseParams

	for rawKey, val := range urlQuery {
		fieldKey = h.concatUrlKeyAndDbPrefix(rawKey)
		if "" != fieldKey && "" != val[0] {
			//fmt.Println(fieldKey, val[0])
			whereClauses = append(whereClauses, whereClauseParams{dbFieldToSearch: fieldKey, searchKey: rawKey, searchValue: val[0]})

			// set search params as url string
			h.setSearchParamToUrlString(fieldKey, val[0])
		}
	}

	//	fmt.Println(whereClauses)

	return whereClauses

}

func (h *HelperPrepareSearch) prepareSearchFieldsFromDbFields(wildcardValue string) []whereClauseParams {

	var whereClauses []whereClauseParams

	//	fmt.Println(h.searchableDbFieldsToIgnoreForWildcard)

	// create wildcard where OR clause
	for _, searchableDbField := range h.searchableDbFields {
		// skip the id or is_admin or is_blocked or is_trashed field by a wildcard search
		// skip fields which were not ment to be for wildcard search
		if lib.ContainsString("_id", searchableDbField) || lib.ContainsString("_is_", searchableDbField) || lib.InArray(h.searchableDbFieldsToIgnoreForWildcard, searchableDbField) {
			//fmt.Println(searchableDbField)
			continue
		}
		whereClauses = append(whereClauses, whereClauseParams{dbFieldToSearch: searchableDbField, searchValue: wildcardValue})
	}

	return whereClauses

}

func (h *HelperPrepareSearch) prepareRightJoins(query *lib.DbAdapter) {
	for _, rightJoin := range h.rightJoins {
		query.RightJoin(rightJoin.tableToJoin, rightJoin.fieldOn, rightJoin.fieldTo)
	}
}

func (h *HelperPrepareSearch) prepareLeftJoins(query *lib.DbAdapter) {
	for _, leftJoin := range h.leftJoins {
		query.LeftJoin(leftJoin.tableToJoin, leftJoin.fieldOn, leftJoin.fieldTo)
	}
}

func (h *HelperPrepareSearch) concatUrlKeyAndDbPrefix(rawkey string) string {
	//fmt.Println(h.searchableDbFields)
	if lib.InArray(h.searchableDbFields, fmt.Sprintf("%s%s", h.dbTableFieldPrefix, rawkey)) {
		return fmt.Sprintf("%s%s", h.dbTableFieldPrefix, rawkey)
	}
	return ""
}

func (h *HelperPrepareSearch) setPagination() bool {

	h.pageNumberPresent = lib.StringToInt(h.request.FormValue("_p"))

	if 0 == h.pageNumberPresent {
		h.pageNumberPresent = 1
	}

	// Check if the search reveald no results
	if h.rowsInTotal == 0 && h.pageNumberPresent == 1 {
		return true
	}

	h.calculateBaseQueryLimit()

	h.pageNumberTotal = lib.RoundToGreater(float64(h.rowsInTotal) / float64(h.itemsPerPage))

	if !h.isPageInLimit() {
		h._handleRedirectAndPanic.RedirectToControllerAction(h.controllerDetails.Controller, h.controllerDetails.Action, h._handleRedirectAndPanic.HttpStatusSeeOther)
		return false
	}

	return true

}

func (h *HelperPrepareSearch) isPageInLimit() bool {
	//	return h.rowsInTotal > h.queryLimitFrom
	return h.pageNumberPresent <= h.pageNumberTotal
}

func (h *HelperPrepareSearch) calculateBaseQueryLimit() {
	h.queryLimitFrom = ((h.pageNumberPresent - 1) * h.itemsPerPage)
	h.queryLimitUntil = h.itemsPerPage
}

func (h *HelperPrepareSearch) calculateFullWildcardQueryLimit(rowsInNotWildcard int, rowsRetrievedFromFullSearch int) {
	h.queryLimitFrom = h.queryLimitFrom - (rowsInNotWildcard - rowsRetrievedFromFullSearch)
	h.queryLimitUntil = h.queryLimitUntil - rowsRetrievedFromFullSearch
	if h.queryLimitUntil <= 0 {
		//		fmt.Println(h.queryLimitUntil)
		h.queryLimitUntil = h.itemsPerPage
	}
}

func (h *HelperPrepareSearch) setItemsPerPage() {
	//h.itemsPerPage = 1
	h.itemsPerPage = lib.StringToInt(h._envconfVars.GetConfVar("listItemsPerPage"))
}

func (h *HelperPrepareSearch) buildBaseQuery(query *lib.DbAdapter, setLimit bool) *lib.DbAdapter {

	h.makeBaseWhereClause(query)
	h.makeBaseWhereOrClause(query)
	h.makeBaseWhereOrMathClause(query)

	h.prepareRightJoins(query)
	h.prepareLeftJoins(query)

	if setLimit {
		query.SetLimit(h.queryLimitFrom, h.queryLimitUntil)
		query.SetOrderByField(h.orderBy, h.orderType)
	}

	h.prepareGroupBy(query)

	return query
}

func (h *HelperPrepareSearch) prepareAdditionalSelectAsFields(query *lib.DbAdapter) {
	for _, additionalSelectAsField := range h.additionalSelectAsFields {
		query.SetAdditionalSelectAsField(additionalSelectAsField)
	}
}

func (h *HelperPrepareSearch) prepareGroupBy(query *lib.DbAdapter) {
	for _, field := range h.groupBy {
		query.SetDistinctGroupByField(field)
	}
}

func (h *HelperPrepareSearch) makeWhereClause(query *lib.DbAdapter, whereClauseParams []whereClauseParams, wildcard persistence.Wildcard, appendNotLikeClause bool) *lib.DbAdapter {

	for _, whereClause := range whereClauseParams {

		query.Where(whereClause.dbFieldToSearch, whereClause.searchValue, wildcard)
		if appendNotLikeClause {
			query.WhereMathComparision(whereClause.dbFieldToSearch, whereClause.searchValue, h._dbAdapter.MathComparision.NotLike())
		}

		// set to search params for views
		h.setSearchParamForView(whereClause.searchKey, whereClause.searchValue)

	}

	return query

}

func (h *HelperPrepareSearch) makeWhereOrClause(query *lib.DbAdapter, whereClauseParams []whereClauseParams, wildcard persistence.Wildcard, appendNotLikeClause bool) *lib.DbAdapter {

	for _, whereClause := range whereClauseParams {
		query.WhereOr(whereClause.dbFieldToSearch, whereClause.searchValue, wildcard)
		if appendNotLikeClause {
			query.WhereOrMathComparision(whereClause.dbFieldToSearch, whereClause.searchValue, h._dbAdapter.MathComparision.NotLike())
		}
	}

	return query
}

func (h *HelperPrepareSearch) makeBaseWhereClause(query *lib.DbAdapter) *lib.DbAdapter {
	for _, whereClauseParam := range h.baseWhereClause {
		query.Where(whereClauseParam.dbFieldToSearch, whereClauseParam.searchValue, whereClauseParam.wildcard)
	}
	return query
}

func (h *HelperPrepareSearch) makeBaseWhereOrClause(query *lib.DbAdapter) *lib.DbAdapter {
	for _, whereOrClauseParam := range h.baseWhereOrClause {
		query.WhereOr(whereOrClauseParam.dbFieldToSearch, whereOrClauseParam.searchValue, whereOrClauseParam.wildcard)
	}
	return query
}

func (h *HelperPrepareSearch) makeBaseWhereOrMathClause(query *lib.DbAdapter) *lib.DbAdapter {
	for _, whereOrMathClauseParam := range h.baseWhereOrMathClause {
		query.WhereOrMathComparision(whereOrMathClauseParam.dbField, whereOrMathClauseParam.value, whereOrMathClauseParam.mathComparision)
	}
	return query
}

func (h *HelperPrepareSearch) selectExecAndSetRowsRetrieved(query *lib.DbAdapter) map[int]map[string]string {

	// prepare additional select as fields
	// ex.: COUNT(zsartcmt_comments) AS article_comments
	h.prepareAdditionalSelectAsFields(query)

	var res = query.ExecSelect()
	h.rowsRetrieved = len(res)
	//	fmt.Println(h._dbAdapter.GetLastExecutedQuery())
	//	fmt.Println(res)
	return res
}

func (h *HelperPrepareSearch) setSearchParamToUrlString(param, value string) {
	h.searchParamsAsUrlString = fmt.Sprintf("%s%s=%s&", h.searchParamsAsUrlString, param, value)
}

func (h *HelperPrepareSearch) setSearchParamForView(param, value string) {
	h.searchParamsForView[param] = value
}
