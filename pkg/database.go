package pkg

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/taschenbergerm/pokescraper/log"
)

// InsertPokemonLink inserts a Pokemon-references from a channel into the database
func InsertPokemonLink(pokeRefChannel chan PokemonEntry,pokeChannel chan Pokemon, quit chan bool) {
	log.Debug("Inserted Active")
	db, err := sql.Open("mysql",
		"root:my-secret-pw@tcp(127.0.0.1:3306)/pokescraper")
	HandleErrorStrictly(err)
	defer db.Close()
	log.Debug("DB Opened")

	err = db.Ping()
	log.Debug("DB Connected")

	pokeRefStmt, err := db.Prepare("REPLACE INTO pokemon_references (name, url) VALUES(?,?)")
	HandleErrorStrictly(err)
	defer pokeRefStmt.Close()

	pokeStmt, err := db.Prepare("REPLACE INTO pokemon"+
	"(`id`,`Name`,`BaseHappiness`,`CaptureRate`,`Color`,`EvolvesFrom`,"+
	"`GenderRate`,`Generation`,`GrowthRate`,`HasGenderDifference`,"+
	"`HatchCounter`) VALUES(?,?,?,?,?,?,?,?,?,?,?)")
	HandleErrorStrictly(err)
	defer pokeStmt.Close()
	log.Debug("Statement Prepared")

	for {
		select {
		case pokemonRef := <-pokeRefChannel:
			log.Debug("Enter Insert pokeRef select")
			p := pokemonRef.PokemonSpecies
			log.Debugf("Recieved Pokemon %s", pokemonRef.PokemonSpecies.Name)
			proxy, err := pokeRefStmt.Exec(p.Name, p.URL)
			HandleErrorSoftly(err)
			printAffected(proxy)
		case pokemon := <- pokeChannel:
			log.Debug("Enter Insert pokemon select")
			proxy, err := pokeStmt.Exec(pokemon.ID,
			pokemon.Name,
			pokemon.BaseHappiness,
			pokemon.CaptureRate,
			pokemon.Color.Name,
			pokemon.EvolvesFrom.Name,
			pokemon.GenderRate,
			pokemon.Generation.Name,
			pokemon.GrowthRate.Name,
			pokemon.HasGenderDifference,
			pokemon.HatchCounter)
			HandleErrorSoftly(err)
			printAffected(proxy)
		case <-quit:
			log.Info("Inserted Shut down")
			return
		}
	}
}


func printAffected(rowProxy sql.Result){
	affected, err := rowProxy.RowsAffected()
			HandleErrorSoftly(err)
			log.Debugf("Insert - Affected %v rows",affected)
}