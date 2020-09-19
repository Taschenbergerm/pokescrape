
CREATE TABLE IF NOT EXISTS pokemon (
id INTEGER PRIMARY KEY,
Name TEXT NOT NULL,
BaseHappiness INTEGER ,
CaptureRate INTEGER,
Color TEXT ,
EvolvesFrom TEXT ,
GenderRate INTEGER ,
Generation TEXT ,
GrowthRate TEXT ,
HasGenderDifference INTEGER(4) ,
HatchCounter INTEGER
);