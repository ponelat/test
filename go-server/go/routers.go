package 

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v2/user/analytics/",
		Index,
	},

	Route{
		"Optimise",
		"POST",
		"/v2/user/analytics/{storeId}/optimisations/{actionName}",
		Optimise,
	},

	Route{
		"OptimiseByCategory",
		"POST",
		"/v2/user/analytics/{storeId}/optimisations/bycategory/{categoryId}/{actionName}",
		OptimiseByCategory,
	},

	Route{
		"OptimiseByChannel",
		"POST",
		"/v2/user/analytics/{storeId}/optimisations/bychannel/{channelId}/{actionName}",
		OptimiseByChannel,
	},

	Route{
		"OptimiseByProduct",
		"POST",
		"/v2/user/analytics/{storeId}/optimisations/byproduct/{productId}/{actionName}",
		OptimiseByProduct,
	},

	Route{
		"DeleteReportFilter",
		"DELETE",
		"/v2/user/analytics/{storeId}/reports/filters/{reportFilterId}",
		DeleteReportFilter,
	},

	Route{
		"ExportStoreReportByCategory",
		"POST",
		"/v2/user/analytics/{storeId}/reports/bycategory/export",
		ExportStoreReportByCategory,
	},

	Route{
		"ExportStoreReportByChannel",
		"POST",
		"/v2/user/analytics/{storeId}/reports/bychannel/export",
		ExportStoreReportByChannel,
	},

	Route{
		"ExportStoreReportByProduct",
		"POST",
		"/v2/user/analytics/{storeId}/reports/byproduct/export",
		ExportStoreReportByProduct,
	},

	Route{
		"GetReportFilter",
		"GET",
		"/v2/user/analytics/{storeId}/reports/filters/{reportFilterId}",
		GetReportFilter,
	},

	Route{
		"GetReportFilters",
		"GET",
		"/v2/user/analytics/{storeId}/reports/filters",
		GetReportFilters,
	},

	Route{
		"GetStoreReportByCategory",
		"POST",
		"/v2/user/analytics/{storeId}/reports/bycategory",
		GetStoreReportByCategory,
	},

	Route{
		"GetStoreReportByChannel",
		"POST",
		"/v2/user/analytics/{storeId}/reports/bychannel",
		GetStoreReportByChannel,
	},

	Route{
		"GetStoreReportByDay",
		"POST",
		"/v2/user/analytics/{storeId}/reports/byday",
		GetStoreReportByDay,
	},

	Route{
		"GetStoreReportByDayExport",
		"POST",
		"/v2/user/analytics/{storeId}/reports/byday/export",
		GetStoreReportByDayExport,
	},

	Route{
		"GetStoreReportByProduct",
		"POST",
		"/v2/user/analytics/{storeId}/reports/byproduct",
		GetStoreReportByProduct,
	},

	Route{
		"SaveReportFilter",
		"PUT",
		"/v2/user/analytics/{storeId}/reports/filters/{reportFilterId}",
		SaveReportFilter,
	},

	Route{
		"CreateRule",
		"POST",
		"/v2/user/analytics/{storeId}/rules",
		CreateRule,
	},

	Route{
		"DeleteRule",
		"DELETE",
		"/v2/user/analytics/{storeId}/rules/{ruleId}",
		DeleteRule,
	},

	Route{
		"DisableRule",
		"POST",
		"/v2/user/analytics/{storeId}/rules/{ruleId}/disable",
		DisableRule,
	},

	Route{
		"EnableRule",
		"POST",
		"/v2/user/analytics/{storeId}/rules/{ruleId}/enable",
		EnableRule,
	},

	Route{
		"GetRule",
		"GET",
		"/v2/user/analytics/{storeId}/rules/{ruleId}",
		GetRule,
	},

	Route{
		"GetRules",
		"GET",
		"/v2/user/analytics/{storeId}/rules",
		GetRules,
	},

	Route{
		"GetRulesExecutions",
		"GET",
		"/v2/user/analytics/{storeId}/rules/executions",
		GetRulesExecutions,
	},

	Route{
		"MoveDownRule",
		"POST",
		"/v2/user/analytics/{storeId}/rules/{ruleId}/movedown",
		MoveDownRule,
	},

	Route{
		"MoveUpRule",
		"POST",
		"/v2/user/analytics/{storeId}/rules/{ruleId}/moveup",
		MoveUpRule,
	},

	Route{
		"RunRule",
		"POST",
		"/v2/user/analytics/{storeId}/rules/{ruleId}/run",
		RunRule,
	},

	Route{
		"RunRules",
		"POST",
		"/v2/user/analytics/{storeId}/rules/run",
		RunRules,
	},

	Route{
		"UpdateRule",
		"PATCH",
		"/v2/user/analytics/{storeId}/rules/{ruleId}",
		UpdateRule,
	},

	Route{
		"GetStoreTrackedClicks",
		"GET",
		"/v2/user/analytics/{storeId}/tracking/clicks",
		GetStoreTrackedClicks,
	},

	Route{
		"GetStoreTrackedExternalOrders",
		"GET",
		"/v2/user/analytics/{storeId}/tracking/externalorders",
		GetStoreTrackedExternalOrders,
	},

	Route{
		"GetStoreTrackedOrders",
		"GET",
		"/v2/user/analytics/{storeId}/tracking/orders",
		GetStoreTrackedOrders,
	},

	Route{
		"GetStoreTrackingStatus",
		"GET",
		"/v2/user/analytics/{storeId}/tracking/status",
		GetStoreTrackingStatus,
	},

	Route{
		"GetTrackingStatus",
		"GET",
		"/v2/user/analytics/tracking/status",
		GetTrackingStatus,
	},

}