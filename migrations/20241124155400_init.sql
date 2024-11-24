-- +goose Up
-- +goose StatementBegin

CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(255) NOT NULL,
    song_title VARCHAR(255) NOT NULL
);

CREATE TABLE verses (
    id SERIAL PRIMARY KEY,
    song_id INT REFERENCES songs(id),
    verse_number INT NOT NULL,
    text TEXT NOT NULL
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS verses;
DROP TABLE IF EXISTS songs;

-- +goose StatementEnd
