package scripts

import (
	"fmt"
	"github.com/taschenbergerm/pokescraper/pkg"
	"github.com/taschenbergerm/pokescraper/config"
	"github.com/taschenbergerm/pokescraper/log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Mysql driver 
	_ "github.com/mattn/go-sqlite3" // SQLite driver 
)

// InitDb initalizes a Database 
func InitDb(args *pkg.InitDBArgs){
	
	var URI string
	var dialect string

	if args.MySQL{
		dialect =  "mysql"
		URI = pkg.InitMysqlURI(args.User, args.Password, args.Host, "sys", args.Port)
	}else  {
		dialect = "sqlite3"
		URI = pkg.InitSQLiteURI()
	}
	log.Debugf("Use %s as dialect", dialect)
	log.Debugf("Use %s as URI", URI)
	assetsFolder := config.Config().GetString("asset_folder")
	assets := fmt.Sprintf("%s%s/", assetsFolder,dialect)

	db, err := sql.Open(dialect,URI)
	pkg.HandleErrorSoftly(err)
	defer db.Close()
	err = db.Ping()
	pkg.HandleErrorSoftly(err)

	stmt := pkg.PrepareStatementFromFile(assets + "create.sql", db)
	_ = pkg.ExecStmtAndClose(stmt)
	stmt = pkg.PrepareStatementFromFile(assets + "create_poketable.sql", db)
	_ = pkg.ExecStmtAndClose(stmt)
	fmt.Println("Exit InitDB")
}

