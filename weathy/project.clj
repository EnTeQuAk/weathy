(defproject weathy "0.1.0-SNAPSHOT"
  :description "FIXME: write description"
  :url "http://example.com/FIXME"
  :license {:name "New and Simplified BSD licenses"
            :url "http://www.opensource.org/licenses/bsd-license.php"}
  :dependencies [[org.clojure/clojure "1.5.1"]
                 [org.clojure/tools.namespace "0.2.4"]
                 [clj-http "0.7.7"]
                 [cheshire "5.2.0"]
                 [environ "0.4.0"]]
  :main ^:skip-aot weathy.core
  :min-lein-version "2.0.0"
  :target-path "target/%s"
  :profiles {:uberjar {:aot :all}})
