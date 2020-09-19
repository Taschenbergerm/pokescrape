-- pokescraper.pokemon_references definition
CREATE TABLE IF NOT EXISTS pokemon_references (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  name TEXT NOT NULL,
  url TEXT DEFAULT NULL
); 
