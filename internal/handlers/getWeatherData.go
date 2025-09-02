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
        rows, err := db.Query("SELECT id, timestamp, temperature, humidity, pressure, gas_resistance, light, pm2_5, wind_speed, wind_direction FROM sensor_data")
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var results []map[string]interface{}
        for rows.Next() {
            var timestamp string
	    var id int
	    var temp float64
	    var humidity float64
	    var pressure float64
	    var gas float64
	    var light int
	    var pm2_5 int
	    var wind_speed float64
	    var wind_direction float64

            if err := rows.Scan(&id, &timestamp, &temp, &humidity, &pressure, &gas, &light, &pm2_5, &wind_speed, &wind_direction); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            results = append(results, map[string]interface{}{
                "id":        id,
		"timestamp": timestamp,
                "temperature": temp,
		"humidity": humidity,
		"pressure": pressure,
		"gas_resistance": gas,
		"light": light,
		"pm2_5": pm2_5,
		"wind_speed": wind_speed,
		"wind_direction": wind_direction,

            })
        }

        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(results)
    }
}
