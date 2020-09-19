SELECT 
    id,
    Name,
    BaseHappiness,
    CaptureRate ,
    Color  ,
    EvolvesFrom  ,
    GenderRate  ,
    Generation  ,
    GrowthRate  ,
    HasGenderDifference,
    HatchCounter 
FROM 
    pokemon 
WHERE 
    Name = ?;