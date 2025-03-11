package repository

const (
	createSong = `INSERT INTO songs (group_name, song_title, release_date, text, link, created_at) 
					VALUES ($1, $2, NULLIF($3, ''), NULLIF($4, ''), NULLIF($5, ''), now()) 
					RETURNING
						song_id,
						group_name,
						song_title,
						COALESCE(release_date, '') AS release_date,
						COALESCE(text, '') AS text,
						COALESCE(link, '') AS link,
						created_at,
						updated_at`

	updateSong = `UPDATE songs 
					SET group_name = COALESCE(NULLIF($1, ''), group_name),
						song_title = COALESCE(NULLIF($2, ''), song_title), 
					    release_date = COALESCE(NULLIF($3, ''), release_date), 
					    text = COALESCE(NULLIF($4, ''), text), 
					    link = COALESCE(NULLIF($5, ''), link), 
					    updated_at = now() 
					WHERE song_id = $6
					RETURNING
						song_id,
						group_name,
						song_title,
						COALESCE(release_date, '') AS release_date,
						COALESCE(text, '') AS text,
						COALESCE(link, '') AS link,
						created_at,
						updated_at`

	deleteSong = `DELETE FROM songs WHERE song_id = $1`

	findByTitleAndGroupCount = `SELECT COUNT(*)
					FROM songs
					WHERE 
					    (song_title ILIKE '%' || $1 || '%') AND
						(group_name ILIKE '%' || $2 || '%')`

	findByTitleAndGroup = `SELECT song_id, group_name, song_title, COALESCE(release_date, '') AS release_date, COALESCE(text, '') AS text, COALESCE(link, '') AS link, created_at, updated_at
					FROM songs
					WHERE 
					    (song_title ILIKE '%' || $1 || '%') AND
						(group_name ILIKE '%' || $2 || '%')
					ORDER BY song_title, created_at, updated_at
					OFFSET $3 LIMIT $4`

	findTextByID = `SELECT text
					FROM songs
					WHERE song_id = $1`
)
