-- table
DROP TABLE IF EXISTS games CASCADE;
-- indexes
DROP INDEX IF EXISTS idx_games_player1 ON game(player1_id);
DROP INDEX IF EXISTS idx_games_player2 ON game(player2_id);
DROP INDEX IF EXISTS idx_games_status ON game(status);