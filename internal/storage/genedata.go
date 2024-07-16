package storage

import (
	"bytes"
	"context"
	"text/template"
	"ydb-client/internal/config"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

type templateConfig struct {
	TablePathPrefix string
}

var fill = template.Must(template.New("fill database").Parse(`
PRAGMA TablePathPrefix("{{ .TablePathPrefix }}");

DECLARE $seriesData AS List<Struct<
	series_id: Uint64,
	title: Text,
	series_info: Text,
	release_date: Date,
	comment: Optional<Text>>>;

DECLARE $seasonsData AS List<Struct<
	series_id: Uint64,
	season_id: Uint64,
	title: Text,
	first_aired: Date,
	last_aired: Date>>;

DECLARE $episodesData AS List<Struct<
	series_id: Uint64,
	season_id: Uint64,
	episode_id: Uint64,
	title: Text,
	air_date: Date>>;

REPLACE INTO series
SELECT
	series_id,
	title,
	series_info,
	CAST(release_date AS Uint64) AS release_date,
	comment
FROM AS_TABLE($seriesData);

REPLACE INTO seasons
SELECT
	series_id,
	season_id,
	title,
	CAST(first_aired AS Uint64) AS first_aired,
	CAST(last_aired AS Uint64) AS last_aired
FROM AS_TABLE($seasonsData);

REPLACE INTO episodes
SELECT
	series_id,
	season_id,
	episode_id,
	title,
	CAST(air_date AS Uint64) AS air_date
FROM AS_TABLE($episodesData);
`))

func FacialillTablesWithData(ctx context.Context, cfg config.Config, c table.Client, prefix string) error {
	// Tx control.
	writeTx := table.TxControl(
		table.BeginTx(
			table.WithSerializableReadWrite(),
		),
		table.CommitTx(),
	)

	r := render(fill, templateConfig{
		TablePathPrefix: prefix,
	})

	return c.Do(ctx,
		func(ctx context.Context, s table.Session) (err error) {
			
			_, _, err = s.Execute(ctx, writeTx, r, table.NewQueryParameters(
				table.ValueParam("$seriesData", getSeriesData(cfg)),
				table.ValueParam("$seasonsData", getSeasonsData(cfg)),
				table.ValueParam("$episodesData", getEpisodesData(cfg)),
			))

			return err
		},
	)
}

func render(t *template.Template, data interface{}) string {
	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		panic(err)
	}

	return buf.String()
}
