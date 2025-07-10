package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"snap-n-go/internal/executors"
	"snap-n-go/internal/logger"
	"snap-n-go/internal/types"
)

func initConfig() {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".cobra" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigType("yaml")
	viper.SetConfigName(".cobra")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

var (
	cp               types.ConnectionParams
	dbConnectionFile string
	l                *logger.Logger
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "SnapNGo",
	Short: "A CLI tool for processing DB backUps",
	Long:  `SnapNGo is a CLI utility that helps DBMS backup and restore operations`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize logger if not already initialized
		var err error
		if l == nil {
			l, err = logger.NewLogger("", "root")
			if err != nil {
				fmt.Println("Failed to initialize logger:", err)
				os.Exit(1)
			}
		}

		// Validate flag values
		if dbConnectionFile != "" {
			// concurrent executors
			fmt.Println("Using multiple DBs file:", dbConnectionFile)
			// TODO: Implement concurrent executors
			return
		}

		// Log the flag values for debugging
		fmt.Println("Command:", cp.Command)
		fmt.Println("Engine:", cp.Engine)
		fmt.Println("Host:", cp.Host)
		fmt.Println("Port:", cp.Port)
		fmt.Println("Username:", cp.Username)
		fmt.Println("Password:", cp.Password)
		fmt.Println("DbName:", cp.DbName)

		// Execute the command
		executors.Single(cp, l)
		return
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	// Define flags and configuration settings
	fmt.Println("Initializing root command")
	rootCmd.Flags().StringVarP(&cp.Command, "command", "c", "", "Command to execute (required)")
	rootCmd.Flags().StringVarP(&cp.Engine, "engine", "e", "", "DB mangment system (required)")
	rootCmd.Flags().StringVarP(&cp.Host, "dbhost", "x", "", "DB host (required)")
	rootCmd.Flags().StringVarP(&cp.Port, "port", "p", "", "DB port (required)")
	rootCmd.Flags().StringVarP(&cp.Username, "username", "u", "", "DB username (required)")
	rootCmd.Flags().StringVarP(&cp.Password, "password", "w", "", "DB password (required)")
	rootCmd.Flags().StringVarP(&cp.DbName, "dbName", "n", "", "DB name (required)")
	rootCmd.Flags().StringVarP(&dbConnectionFile, "multipleDBsFile", "", "", "path file to a list of DBs config and their commands")

	// Mark required flags
	rootCmd.MarkFlagRequired("command")
	rootCmd.MarkFlagRequired("engine")
	rootCmd.MarkFlagRequired("dbhost")
	rootCmd.MarkFlagRequired("port")
	//rootCmd.MarkFlagRequired("username")
	// rootCmd.MarkFlagRequired("password")
	// rootCmd.MarkFlagRequired("dbName")

	fmt.Println("Finishing init root command" + cp.Command)

	//info comes from config file OR cmd arguments.

	rootCmd.MarkFlagsMutuallyExclusive("multipleDBsFile", "command")
	rootCmd.MarkFlagsMutuallyExclusive("multipleDBsFile", "engine")
	rootCmd.MarkFlagsMutuallyExclusive("multipleDBsFile", "dbhost")
	rootCmd.MarkFlagsMutuallyExclusive("multipleDBsFile", "port")
	rootCmd.MarkFlagsMutuallyExclusive("multipleDBsFile", "username")
	rootCmd.MarkFlagsMutuallyExclusive("multipleDBsFile", "password")
}

// Execute executes the root command
func Execute() error {
	return rootCmd.Execute()
}
