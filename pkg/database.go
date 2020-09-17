package pkg

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/taschenbergerm/pokescraper/log"
)

// InsertPokemonLink inserts a Pokemon-references from a channel into the database
func InsertPokemonLink(pokeChannel chan PokemonEntry, quit chan bool) {
	log.Debug("Inserted Active")
	db, err := sql.Open("mysql",
		"root:my-secret-pw@tcp(127.0.0.1:3306)/pokescraper")
	HandleErrorStrictly(err)
	defer db.Close()
	log.Debug("DB Opened")

	err = db.Ping()
	log.Debug("DB Connected")

	stmt, err := db.Prepare("INSERT INTO pokemon_references (name, url) VALUES(?,?)")
	HandleErrorStrictly(err)
	defer stmt.Close()
	log.Debug("Statement Prepared")

	for {
		select {
		case pokemon := <-pokeChannel:
			log.Debug("Enter Insert select")
			p := pokemon.PokemonSpecies
			log.Debugf("Recieved Pokemon %s", pokemon.PokemonSpecies.Name)
			proxy, err := stmt.Exec(p.Name, p.URL)
			HandleErrorSoftly(err)
			log.Debugf("Insert %s  - Affected %v rows", p.Name, proxy.RowsAffected)
		case <-quit:
			log.Info("Inserted Shut down")
			return
		}
	}

}
