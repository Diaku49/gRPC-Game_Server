-- create game table
CREATE TABLE IF NOT EXISTS games (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   player1_id UUID NOT NULL,
   player2_id UUID NOT NULL,
   winner_id UUID,

   status GAME_STATUS NOT NULL,
   rounds_num SMALLINT,

   created_at TIMESTAMPTZ DEFAULT NOW(),
   updated_at TIMESTAMPTZ DEFAULT NOW(),
   -- Constraints
   CONSTRAINT fk_player1 FOREIGN KEY (player1_id) REFERENCES users(id),
   CONSTRAINT fk_player2 FOREIGN KEY (player2_id) REFERENCES users(id),
   CONSTRAINT fk_winner FOREIGN KEY (winner_id) REFERENCES users(id),
   CONSTRAINT winner_is_player CHECK (winner_id IS NULL OR winner_id IS IN (player1_id, player2_id))
   CONSTRAINT different_players CHECK (player1_id != player2_id),
);

-- indexes
CREATE INDEX IF NOT EXISTS idx_games_player1 ON games(player1_id);
CREATE INDEX IF NOT EXISTS idx_games_player2 ON games(player2_id);
CREATE INDEX IF NOT EXISTS idx_games_status ON games(status);

-- enum
CREATE TYPE GAME_STATUS AS ENUM (
    'open'
    'finished'
    'in_progress'
    'cancelled'
);