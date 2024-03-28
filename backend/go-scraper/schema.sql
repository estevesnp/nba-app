DROP DATABASE IF EXISTS nbaappdb;

CREATE DATABASE nbaappdb;

DROP TABLE IF EXISTS players;

CREATE TABLE players (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    position VARCHAR(3) NOT NULL CHECK (position IN(
        'G', 'F', 'C', 'G/F', 'F/C'
    )),
    team VARCHAR(3) NOT NULL CHECK (team IN (
        'ATL', 'BOS', 'BKN', 'CHA', 'CHI', 'CLE',
        'DAL', 'DEN', 'DET', 'GSW', 'HOU', 'IND',
        'LAC', 'LAL', 'MEM', 'MIA', 'MIL', 'MIN',
        'NOP', 'NYK', 'OKC', 'ORL', 'PHI', 'PHX',
        'POR', 'SAC', 'SAS', 'TOR', 'UTA', 'WAS'
    ))
);