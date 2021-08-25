package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/cmd/rest-api/lib"
	records "github.com/lunarnuts/go-course/tree/course-project/course-project/Backend/src/db/models"
)

func GetCurrentWeather(p *pgxpool.Pool, w http.ResponseWriter, r *http.Request) {
	cityName := lib.CityNameFromVars(r)
	timeRequested := time.Now()
	temperature := 0.0
	rec := records.Record{
		CityName:      cityName,
		TimeRequested: timeRequested.String(),
		Temperature:   float64(temperature),
	}
	id, err := records.Insert(p, rec)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	weather, err := lib.GetCurrentWeatherFromAPI(ctx, cityName)
	if err != nil {
		lib.ReturnInternalError(w, err)
	}
	rec.Temperature = weather.Temperature
	err = records.Update(p, id, rec)
	if err != nil {
		lib.ReturnInternalError(w, err)
		return
	}
	rec, err = records.Select(p, id)
	if err != nil {
		w.WriteHeader(500)
		return
	}
	lib.ReturnJSON(w, rec)
}