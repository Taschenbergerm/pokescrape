package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/taschenbergerm/pokescraper/config"
	"github.com/taschenbergerm/pokescraper/log"
	"github.com/taschenbergerm/pokescraper/pkg"
)

func init() {
	var lookupCmd = &cobra.Command{
		Use: "lookup [id or name]",
		Short: "Look up a pokemon by Id or name",
		Long: `Lookup will search the database for the specified name or id (if numeric only)
		to find the repspective entry in the pokemon table`,
		Args: cobra.ExactValidArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Pseudo Lookup for: " + strings.Join(args, " "))
			fmt.Printf("The password is %s \n", config.Config().GetString("password"))
			if config.Config().GetBool("configured"){
				log.Infoln("The programm is fully configured")
			}else{
				log.Infoln("The app needs more configuration, please run")
				log.Infoln("\t $ pokescraper initdb [Options]")
				os.Exit(0)
			}
			pokemon := pkg.LookUpPokemonByName(args[0])
			repr, err := json.Marshal(pokemon)
			pkg.HandleErrorSoftly(err)
			println(string(repr))
		  },
	}

	
	rootCmd.AddCommand(lookupCmd)
  }

