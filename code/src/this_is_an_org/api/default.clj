(ns this-is-an-org.api.default
  (:require [this-is-an-org.core :refer [call-api check-required-params with-collection-format]])
  (:import (java.io File)))

(defn josh-get-with-http-info
  "hi
  hi"
  []
  (call-api "/josh" :get
            {:path-params   {}
             :header-params {}
             :query-params  {}
             :form-params   {}
             :content-types []
             :accepts       []
             :auth-names    []}))

(defn josh-get
  "hi
  hi"
  []
  (:data (josh-get-with-http-info)))
