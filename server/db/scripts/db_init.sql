-- Create the 'teams' table
CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    strength INT NOT NULL,
    points INT DEFAULT 0,
    played INT DEFAULT 0,
    won INT DEFAULT 0,
    drawn INT DEFAULT 0,
    lost INT DEFAULT 0,
    goal_diff INT DEFAULT 0
);
-- Create the 'matches' table
CREATE TABLE matches (
    id SERIAL PRIMARY KEY,
    home_team_id INT REFERENCES teams(id),
    away_team_id INT REFERENCES teams(id),
    home_team_score INT DEFAULT 0,
    away_team_score INT DEFAULT 0,
    week INT NOT NULL,
    is_played BOOLEAN DEFAULT FALSE
);
-- Create the 'leagues' table
CREATE TABLE leagues (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);
-- Create the 'league_teams' table for many-to-many relationship
CREATE TABLE league_teams (
    league_id INT REFERENCES leagues(id),
    team_id INT REFERENCES teams(id),
    PRIMARY KEY (league_id, team_id)
);
-- Insert initial data into the 'teams' table
INSERT INTO teams (name, strength)
VALUES ('Chelsea', 85);
INSERT INTO teams (name, strength)
VALUES ('Arsenal', 80);
INSERT INTO teams (name, strength)
VALUES ('Manchester City', 90);
INSERT INTO teams (name, strength)
VALUES ('Liverpool', 82);
-- Insert initial data into the 'leagues' table
INSERT INTO leagues (name)
VALUES ('Premier League');
-- Insert initial data into the 'league_teams' table
INSERT INTO league_teams (league_id, team_id)
VALUES (1, 1);
INSERT INTO league_teams (league_id, team_id)
VALUES (1, 2);
INSERT INTO league_teams (league_id, team_id)
VALUES (1, 3);
INSERT INTO league_teams (league_id, team_id)
VALUES (1, 4);