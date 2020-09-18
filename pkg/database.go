package pkg

import (
	"database/sql"
	"io/ioutil"
	"os"

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

	pokeRefStmt := prepareStatementFromFile("assets/insert_pokemon_ref.sql", db)
	defer pokeRefStmt.Close()

	pokeStmt := prepareStatementFromFile("assets/insert_pokemon.sql", db) 
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

func prepareStatementFromFile(path string,db *sql.DB) *sql.Stmt{
	dir, err := os.Getwd()
	HandleErrorSoftly(err)
	log.Debugf("Looking relativly from %v",dir )
	file ,err := os.Open(path)
	HandleErrorSoftly(err)
	content, err := ioutil.ReadAll(file)
	HandleErrorSoftly(err)
	stmt, err := db.Prepare(string(content))
	HandleErrorSoftly(err)
	return stmt
}

func printAffected(rowProxy sql.Result){
	affected, err := rowProxy.RowsAffected()
			HandleErrorSoftly(err)
			log.Debugf("Insert - Affected %v rows",affected)
}