package storage

import (
	"time"
	"ydb-client/internal/config"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

func seriesData(id uint64, released time.Time, title, info, comment string) types.Value {
	// Comment type value.
	var commentv types.Value
	if comment == "" {
		commentv = types.NullValue(types.TypeUTF8)
	} else {
		commentv = types.OptionalValue(types.TextValue(comment))
	}

	return types.StructValue(
		types.StructFieldValue("series_id", types.Uint64Value(id)),
		types.StructFieldValue("release_date", types.DateValueFromTime(released)),
		types.StructFieldValue("title", types.TextValue(title)),
		types.StructFieldValue("series_info", types.TextValue(info)),
		types.StructFieldValue("comment", commentv),
	)
}

func seasonData(seriesID, seasonID uint64, title string, first, last time.Time) types.Value {
	return types.StructValue(
		types.StructFieldValue("series_id", types.Uint64Value(seriesID)),
		types.StructFieldValue("season_id", types.Uint64Value(seasonID)),
		types.StructFieldValue("title", types.TextValue(title)),
		types.StructFieldValue("first_aired", types.DateValueFromTime(first)),
		types.StructFieldValue("last_aired", types.DateValueFromTime(last)),
	)
}

func episodeData(seriesID, seasonID, episodeID uint64, title string, date time.Time) types.Value {
	return types.StructValue(
		types.StructFieldValue("series_id", types.Uint64Value(seriesID)),
		types.StructFieldValue("season_id", types.Uint64Value(seasonID)),
		types.StructFieldValue("episode_id", types.Uint64Value(episodeID)),
		types.StructFieldValue("title", types.TextValue(title)),
		types.StructFieldValue("air_date", types.DateValueFromTime(date)),
	)
}

func getSeriesData(cnf config.Config) types.Value {
	a := make([]types.Value, 0)
	var serial uint64
	for serial = 1; serial < cnf.Series; serial++ {
		// TODO: issue before 1970 linux date.
		f, _ := time.Parse(time.DateOnly, "1970-01-01")
		t, _ := time.Parse(time.DateOnly, "2024-01-01")
		a = append(a, seriesData(serial, gofakeit.DateRange(f, t), gofakeit.Adverb()+" "+gofakeit.Color(), gofakeit.Noun()+" "+gofakeit.Animal(), "comment"))
	}

	return types.ListValue(a...)
}

func getSeasonsData(cnf config.Config) types.Value {
	s := make([]types.Value, 0)
	var serial uint64
	for serial = 1; serial < cnf.Series; serial++ {
		var season uint64
		for season = 1; season < cnf.Seasons; season++ {
			f, _ := time.Parse(time.DateOnly, "1970-01-01")
			t, _ := time.Parse(time.DateOnly, "2024-01-01")
			s = append(s, seasonData(serial, season, gofakeit.Drink(), gofakeit.DateRange(f, t), gofakeit.DateRange(f, t)))
		}
	}

	return types.ListValue(s...)
}

func getEpisodesData(cnf config.Config) types.Value {
	s := make([]types.Value, 0)
	var serial uint64
	for serial = 1; serial < cnf.Series; serial++ {
		var season uint64
		for season = 1; season < cnf.Seasons; season++ {
			var episodes uint64
			for episodes = 1; episodes < cnf.Episodes; episodes++ {
				f, _ := time.Parse(time.DateOnly, "1970-01-01")
				t, _ := time.Parse(time.DateOnly, "2024-01-01")
				s = append(s, episodeData(serial, season, episodes, gofakeit.Snack(), gofakeit.DateRange(f, t)))
			}

		}

	}

	return types.ListValue(s...)
}

// episodeData(1, 1, 2, "Calamity Jen", days("2006-02-03")),
