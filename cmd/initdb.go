package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/taschenbergerm/pokescraper/config"
	"github.com/taschenbergerm/pokescraper/log"
	"github.com/taschenbergerm/pokescraper/pkg"
	"github.com/taschenbergerm/pokescraper/scripts"
)


func init() {	
	
	Config := config.Config()
	var CmdArgs pkg.InitDBArgs
	
	pflag.StringVarP(&CmdArgs.User, "user", "u", "root", "user for the database")
    pflag.StringVarP(&CmdArgs.Password, "password", "p", "my-secret-pw", "URI of the database")
    pflag.StringVarP(&CmdArgs.Host, "host", "H", "localhost", "URL of the database host")
    pflag.IntVarP(&CmdArgs.Port, "port", "P", 3306, "port on the database host")
	pflag.BoolVar(&CmdArgs.SQLite, "sqlite", true, "use SQLite Database (default)")
	pflag.BoolVar(&CmdArgs.MySQL, "mysql", false, "use MySQL Database instead")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	if CmdArgs.MySQL {
		CmdArgs.SQLite = false
	}

	var initdbCmd = &cobra.Command{
		Use: "initdb",
		Short: "Initialize the DB ",
		Long: "Initialize the database ", 
		Args: cobra.ExactArgs(0),
		Run: func (cmd *cobra.Command, args []string){
			fmt.Println("Pseudo Initdb: "+ strings.Join(args, " "))
			scripts.InitDb(&CmdArgs)
			viper.Set("configured", true)
			fmt.Println("Write Config File: "+ string(Config.GetString("config_path")))
			err := os.MkdirAll(Config.GetString("config_folder"), 0777)
			if err != nil{				
				fmt.Println(fmt.Errorf("An error occoured %s", err))
			}
			err = viper.WriteConfigAs(config.Config().GetString("config_path"))
			if err != nil{				
				fmt.Println(fmt.Errorf("An error occoured %s", err))

			}else{
				log.Info("PokeScraper is now configured - start scraping ")
			}
			pkg.Main()
			log.Info("Database is filled")
		},
	}	
	rootCmd.AddCommand(initdbCmd)


}
  


