package handlers

import (
    "fmt"
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

        w.Header().Set("Content-Type", "text/html; charset=utf-8")

        // Simple HTML table
        fmt.Fprintln(w, `<table class="weather-data">`)
        fmt.Fprintln(w, `<thead>
            <tr>
                <th>Time</th>
                <th>Temp (°C)</th>
                <th>Humidity (%)</th>
                <th>Pressure (hPa)</th>
                <th>Gas</th>
                <th>Light</th>
                <th>PM2.5</th>
                <th>Wind Speed</th>
                <th>Wind Dir</th>
            </tr>
        </thead>
        <tbody>`)

        for rows.Next() {
            var (
                id            int
                timestamp     string
                temp          float64
                humidity      float64
                pressure      float64
                gas           float64
                light         int
                pm2_5          int
                windSpeed     float64
                windDirection float64
            )

            if err := rows.Scan(&id, &timestamp, &temp, &humidity, &pressure, &gas, &light, &pm2_5, &windSpeed, &windDirection); err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            fmt.Fprintf(w,
                `<tr>
                    <td>%s</td>
                    <td>%.1f</td>
                    <td>%.1f</td>
                    <td>%.1f</td>
                    <td>%.1f</td>
                    <td>%d</td>
                    <td>%d</td>
                    <td>%.1f</td>
                    <td>%.1f</td>
                </tr>`,
                timestamp, temp, humidity, pressure, gas, light, pm2_5, windSpeed, windDirection,
            )
        }

        fmt.Fprintln(w, `</tbody></table>`)
    }
}
