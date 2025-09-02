package handlers

import (
    "encoding/json"
	"net/http"
	"database/sql"
        _ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// InitDB opens the SQLite database once and keeps the handle
func InitDB(path string) error {
    var err error
    db, err = sql.Open("sqlite3", path)
    if err != nil {
        return err
    }
    return db.Ping()
}

func GetWeatherData() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT city, temperature FROM weather")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var results []map[string]interface{}
        for rows.Next() {
            var city string
            var temp float64
            if err := rows.Scan(&city, &temp); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            results = append(results, map[string]interface{}{
                "city":        city,
                "temperature": temp,
            })
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(results)
    }
}
