-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS songs (
     song_id SERIAL PRIMARY KEY,
     group_name VARCHAR(255),
     song_title VARCHAR(255),
     release_date VARCHAR(10),
     text TEXT,
     link VARCHAR(255),
     created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
     updated_at TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS songs_title_id_idx ON songs (song_title);
CREATE INDEX IF NOT EXISTS songs_group_id_idx ON songs (group_name);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS songs;
-- +goose StatementEnd