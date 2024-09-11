CREATE TABLE IF NOT EXISTS user_teams (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    team_id INTEGER REFERENCES teams(id) ON DELETE CASCADE,
    user_role VARCHAR(100) DEFAULT 'None',
    PRIMARY KEY (user_id, team_id)
);

