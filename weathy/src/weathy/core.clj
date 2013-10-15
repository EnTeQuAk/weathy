(ns weathy.core
  (:gen-class)
  (:use [clojure.string :only [join]]
        [environ.core])
  (:require [cheshire.core :as json]
            [clj-http.client :as client]))


(defn forecast
  "Retrieve the forecast for a given latitude and longitude"
  [lat lon]
  (let [base-url "https://api.forecast.io/forecast"
        params {:units "si"}
        api-key (env :forecast-api-key)
        url (join "/" [base-url api-key (join "," (filter #(not-empty %) (map str [lat lon])))])
        response (client/get url {:query-params params :throw-exceptions true})]
    (cond (= 200 (:status response))
          (json/parse-string (:body response) true))))

(defn geocode
  "Retrieve latitude and longitude for an address"
  [address sensor]
  (let [base-url "http://maps.googleapis.com/maps/api/geocode/json"
        params {:address address :sensor sensor}
        response (client/get base-url {:query-params params :throw-exceptions true})]
    (cond (= 200 (:status response))
          (get-in (json/parse-string (:body response) true) [:results 0 :geometry :location]))))

(defn -main
  [& args]
  (let [location (geocode "Berlin" "false")
        info (apply forecast (map location [:lat :lng]))]
    (println (get-in info [:daily :summary]))))
