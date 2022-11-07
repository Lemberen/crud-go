    CREATE TABLE IF NOT EXISTS Books (
        [id]        INTEGER primary key,
        [title]     TEXT    not null,
        [author]    TEXT    not null,
        [publisher] TEXT,
        [language]  TEXT    not null,
        [pages]     INTEGER not null,
        [edition]   INTEGER,
        [year]      INTEGER,
        [isbn]      TEXT    UNIQUE
    );

