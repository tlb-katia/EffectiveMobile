package repository

import (
	"EffectiveMobile_Project/internal/entities"
	"EffectiveMobile_Project/pkg/storage/postgres"
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"strings"
)

type Repository struct {
	db      *sql.DB
	builder sq.StatementBuilderType
}

func NewRepository(storage *postgres.Storage) *Repository {
	return &Repository{
		db:      storage.Db,
		builder: storage.Builder,
	}
}

// todo: time display and verses

func (r *Repository) GetAllSongsFiltered(ctx context.Context, req *entities.AllSongsRequest) (*[]entities.AllSongsResponse, error) {
	query := r.builder.Select("group_name", "song_name", "release_date", "link").
		From("songs")

	if req.SongName != nil {
		query = query.Where(sq.Eq{"song_name": &req.SongName})
	}
	if req.GroupName != nil {
		query = query.Where(sq.Eq{"group_name": &req.GroupName})
	}
	if req.ReleaseDate != nil {
		query = query.Where(sq.Eq{"release_date": &req.ReleaseDate})
	}

	q, args, err := query.Limit(req.Limit).Offset(req.Offset).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var resps []entities.AllSongsResponse

	for rows.Next() {
		var resp entities.AllSongsResponse
		err := rows.Scan(&resp.GroupName, &resp.SongName, &resp.ReleaseDate, &resp.Link)
		if err != nil {
			return nil, err
		}
		resps = append(resps, resp)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &resps, nil
}

func (r *Repository) GetLyricsPaginated(ctx context.Context, req *entities.LyricsRequest) (*[]string, error) {
	subQuery, subArgs, err := r.builder.Select("id").
		From("songs").
		Where(sq.And{
			sq.Eq{"song_name": req.SongName},
			sq.Eq{"group_name": req.GroupName},
		}).ToSql()
	if err != nil {
		return nil, err
	}

	query := r.builder.Select("verse_text").
		From("verses").
		Where("song_id = ("+subQuery+")", subArgs...).
		OrderBy("verse_number").
		Limit(req.Limit).
		Offset(req.Offset)

	q, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}
	var resps []string

	for rows.Next() {
		var resp string
		err := rows.Scan(&resp)
		if err != nil {
			return nil, err
		}
		resps = append(resps, resp)
	}

	return &resps, nil
}

func (r *Repository) DeleteSong(ctx context.Context, songId int) error {
	query, args, err := r.builder.Delete("songs").
		Where(sq.Eq{"id": &songId}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) ChangeSongData(ctx context.Context, req *entities.ChangeSongReq) (*entities.AddSong, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query := r.builder.Update("songs")

	if req.SongName != nil {
		query = query.Set("song_name", *req.SongName)
	}
	if req.ReleaseDate != nil {
		query = query.Set("release_date", *req.ReleaseDate)
	}
	if req.GroupName != nil {
		query = query.Set("group_name", *req.GroupName)
	}

	if req.Link != nil {
		query = query.Set("link", *req.Link)
	}

	q, args, err := query.Where(sq.Eq{"id": req.Id}).ToSql()
	if err != nil {
		return nil, err
	}

	_, err = tx.ExecContext(ctx, q, args...)
	if err != nil {
		return nil, err
	}

	if req.VerseNumText != nil {
		for verseNum, verseText := range *req.VerseNumText {
			updateQuery := r.builder.Update("verses").
				Set("verse_text", verseText).
				Where(sq.Eq{
					"song_id":      req.Id,
					"verse_number": verseNum,
				})
			q, args, err = updateQuery.ToSql()
			if err != nil {
				return nil, err
			}
			result, err := tx.ExecContext(ctx, q, args...)
			if err != nil {
				return nil, err
			}

			rowsAffected, err := result.RowsAffected()
			if err != nil {
				return nil, err
			}
			if rowsAffected == 0 {
				insertQuery := r.builder.Insert("verses").
					Columns("song_id", "verse_number", "verse_text").
					Values(req.Id, verseNum, verseText)

				q, args, err = insertQuery.ToSql()
				if err != nil {
					fmt.Println("!!!!!", err.Error())
					return nil, err
				}

				_, err = tx.ExecContext(ctx, q, args...)
				if err != nil {
					fmt.Println("errrror", err.Error())
					return nil, err
				}

			}

		}

	}
	return r.getASong(ctx, req.Id)
}

func (r *Repository) AddSong(ctx context.Context, req *entities.AddSong) (*entities.AddSong, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query, args, err := r.builder.Insert("songs").
		Columns("group_name", "song_name", "release_date", "link").
		Values(req.GroupName, req.SongName, req.ReleaseDate, req.Link).
		Suffix("RETURNING id").
		ToSql()

	if err != nil {
		return nil, err
	}
	var songId uint64

	err = tx.QueryRow(query, args...).Scan(&songId)
	if err != nil {
		return nil, err
	}

	q := r.builder.Insert("verses").
		Columns("song_id", "verse_number", "verse_text")

	versesSlice := r.breakVersesIntoSlices(req.Text)

	for i := 1; i <= len(versesSlice); i++ {
		q = q.Values(songId, i, versesSlice[i])
	}

	queryVerses, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	_, err = tx.Exec(queryVerses, args...)
	if err != nil {
		return nil, err
	}

	return r.getASong(ctx, songId)
}

//todo: how do i display text?

func (r *Repository) getASong(ctc context.Context, songId uint64) (*entities.AddSong, error) {
	res := &entities.AddSong{}

	query, args, err := r.builder.
		Select("group_name", "song_name", "release_date", "link").
		From("songs").
		Where(sq.Eq{"id": songId}).
		ToSql()

	if err != nil {
		return nil, err
	}

	err = r.db.QueryRowContext(ctc, query, args...).Scan(&res.GroupName, &res.SongName, &res.ReleaseDate, &res.Link)
	if err != nil {
		return nil, err
	}

	queryVerses, args, err := r.builder.
		Select("verse_text").
		From("verses").
		Where(sq.Eq{"song_id": songId}).
		ToSql()

	if err != nil {
		return nil, err
	}

	rows, err := r.db.QueryContext(ctc, queryVerses, args...)
	if err != nil {
		return nil, err
	}

	var sb strings.Builder

	for rows.Next() {
		var verse string
		err := rows.Scan(&verse)
		if err != nil {
			return nil, err
		}
		sb.WriteString(verse)
	}
	res.Text = sb.String()

	return res, nil
}

func (r *Repository) breakVersesIntoSlices(text string) []string {
	result := strings.Split(text, "\n")
	return result
}
