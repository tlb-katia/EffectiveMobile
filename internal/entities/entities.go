package entities

type AllSongsRequest struct {
	GroupName   *string     `json:"omitempty,group_name"`
	SongName    *string     `json:"song_name,omitempty"`
	ReleaseDate *CustomTime `json:"release_date,omitempty"`
	Link        *string     `json:"link,omitempty"`
	Offset      uint64      `json:"offset,omitempty"`
	Limit       uint64      `json:"limit,omitempty"`
}

type AllSongsResponse struct {
	GroupName   string     `json:"omitempty,group_name"`
	SongName    string     `json:"song_name,omitempty"`
	ReleaseDate CustomTime `json:"release_date,omitempty"`
	Text        string     `json:"text,omitempty"`
	Link        string     `json:"link,omitempty"`
}

type LyricsRequest struct {
	SongName  string `json:"song_name" validate:"required"`
	GroupName string `json:"group_name" validate:"required"`
	Offset    uint64 `json:"offset,omitempty"`
	Limit     uint64 `json:"limit,omitempty"`
}

type AddSong struct {
	GroupName   string     `json:"group_name" validate:"required"`
	SongName    string     `json:"song_name" validate:"required"`
	ReleaseDate CustomTime `json:"release_date" validate:"required"`
	Text        string     `json:"text" validate:"required"`
	Link        string     `json:"link" validate:"required"`
}

type ChangeSongReq struct {
	Id           uint64
	GroupName    *string         `json:"group_name,omitempty"`
	SongName     *string         `json:"song_name,omitempty"`
	ReleaseDate  *CustomTime     `json:"release_date,omitempty"`
	VerseNumText *map[int]string `json:"text,omitempty"`
	Link         *string         `json:"link,omitempty"`
}
