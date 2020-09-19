package pkg

import (
	"fmt"
	"database/sql"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3" // SQLite Driver 
	"github.com/taschenbergerm/pokescraper/log"
	"github.com/taschenbergerm/pokescraper/config"
)
//LookUpPokemonByName queries a pokemon by its name 
func LookUpPokemonByName(name string) Pokemon{
	var pokemon Pokemon
	Config := config.Config()
	log.Debug("Inserted Active")
	db, err := sql.Open(assambleURI(Config))
	HandleErrorStrictly(err)
	defer db.Close()
	log.Debug("DB Opened")
	stmt := PrepareStatementFromFile("assets/lookup.sql", db)
	err = stmt.QueryRow(name).Scan(&pokemon.ID,
		&pokemon.Name,
		&pokemon.BaseHappiness,
		&pokemon.CaptureRate,
		&pokemon.Color.Name,
		&pokemon.EvolvesFrom.Name,
		&pokemon.GenderRate,
		&pokemon.Generation.Name,
		&pokemon.GrowthRate.Name,
		&pokemon.HasGenderDifference,
		&pokemon.HatchCounter)
	HandleErrorSoftly(err)
	log.Debugf("Found entry for %s (id=%v) \n", pokemon.Name, pokemon.ID)
	return pokemon
}

// InsertPokemonLink inserts a Pokemon-references from a channel into the database
func InsertPokemonLink(pokeRefChannel chan PokemonEntry,pokeChannel chan Pokemon, quit chan bool) {
	Config := config.Config()
	log.Debug("Inserted Active")
	dialect, uri := assambleURI(Config)
	dialectSQLFolder := fmt.Sprintf("assets/%s/", dialect)
	db, err := sql.Open(dialect, uri)
	HandleErrorStrictly(err)
	defer db.Close()
	log.Debug("DB Opened")

	err = db.Ping()
	log.Debug("DB Connected")

	pokeRefStmt := PrepareStatementFromFile(dialectSQLFolder+"insert_pokemon_ref.sql", db)
	defer pokeRefStmt.Close()

	pokeStmt := PrepareStatementFromFile(dialectSQLFolder+"insert_pokemon.sql", db) 
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

// PrepareStatementFromFile reads a sql file and prepares a stmt from it 
func PrepareStatementFromFile(path string,db *sql.DB) *sql.Stmt{
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


func assambleURI(c config.Provider) (string, string) {
	var dialect string 
	var URI string 

	if c.GetBool("mysql"){
		dialect =  "mysql"
		URI = InitMysqlURI(
			c.GetString("user"),
			c.GetString("password"),
			c.GetString("host"),
			c.GetString("pokescraper"),
			c.GetInt("port"))
	}else  {
		dialect = "sqlite3"
		URI = InitSQLiteURI()
	}
	return dialect, URI
}

//ExecStmtAndClose executes a statement and caputes panics 
func ExecStmtAndClose(stmt *sql.Stmt) (err error) {
	defer func() {
        if r := recover(); r != nil {
			err = fmt.Errorf("Recovered in Exec due to:\n %s", r)
        }
	}()
	_, err = stmt.Exec()
	defer stmt.Close()
	return err
}

//InitMysqlURI assembles the URI for a mysql database
func InitMysqlURI(user, password, host, database string,  port int ) string{
	URI := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s",
						user,
						password, 
						host,
						port,
						database)
	return URI 
}

//InitSQLiteURI assembles the uri for a sqlite db (which is a file path)
func InitSQLiteURI() string{
	URI := config.Config().GetString("sqlite_db")
	return URI 
}