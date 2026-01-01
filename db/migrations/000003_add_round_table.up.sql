CREATE TABLE IF NOT EXISTS rounds (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    game_id UUID NOT NULL,
     win_id UUID NOT NULL,

    sequence SMALLINT NOT NULL,
    status ROUND_STATUS NOT NULL DEFAULT "started",

    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),

    CONSTRAINT fk_game_id FOREIGN KEY (game_id) REFERENCES game(id),
    CONSTRAINT fk_win_id FOREIGN KEY (win_id) REFERENCES users(id),
);

-- indexes
CREATE INDEX IF NOT EXISTS idx_win_id ON rounds(win_id);
CREATE INDEX IF NOT EXISTS idx_game_id ON rounds(game_id);


-- enum
CREATE TYPE ROUND_STATUS AS ENUM (
    "started",
    "finished",
    "cancelled",
);