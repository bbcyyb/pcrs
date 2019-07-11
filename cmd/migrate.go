package cmd

import (
	"fmt"
	. "github.com/bbcyyb/pcrs/infra/logger"
	"github.com/bbcyyb/pcrs/infra/migrate"
	migrateV4 "github.com/golang-migrate/migrate/v4"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/sqlserver"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
)

const migrateVersion = "v1.0.0"

var (
	help        bool
	version     bool
	prefetch    uint
	lockTimeout uint
	path        string
	database    string
	source      string
	migrater    *migrateV4.Migrate
	migraterErr error
	startTime   time.Time
)

/**
 * sample:
 * look version: `pcrs migrate look --database sqlserver://PowerCalc:Power%401433@10.35.83.61:1433?database=PowerCalcFor46  --source file://migrate_sql`
 * up:           `pcrs migrate up 1 --database sqlserver://PowerCalc:Power%401433@10.35.83.61:1433?database=PowerCalcFor46  --source file://migrate_sql`
 * down:         `pcrs migrate down 1 --database sqlserver://PowerCalc:Power%401433@10.35.83.61:1433?database=PowerCalcFor46  --source file://migrate_sql`
 */
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	Args:  cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		Logger.Debugf("Inside migrateCmd PersistentPreRun with args: %v", strings.Join(args, " "))
		preRun()
	},
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Debugf("migrateCmd arguments : %v", strings.Join(args, " "))
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		postRun()
	},
}

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "migrate up",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		arg := ""
		if len(args) > 0 {
			arg = args[0]
		}
		migrate.ExecuteUp(migrater, arg)
	},
}

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "migrate down",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		migrate.ExecuteDown(migrater, args[0])
	},
}

var gotoCmd = &cobra.Command{
	Use:   "goto",
	Short: "migrate to specified version",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Debugf("goto command args: %v", strings.Join(args, " "))
		migrate.ExecuteGoto(migrater, args[0])
	},
}

var lookCmd = &cobra.Command{
	Use:   "look",
	Short: "have a look at the current migration database version",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		migrate.ExecuteLook(migrater)
	},
}

func preRun() {
	Logger.Debugf("help:%v, version:%v, prefetch:%v, lockTimeout:%v, path:%v, database:%v, source:%v\n",
		help, version, prefetch, lockTimeout, path, database, source)

	if version {
		Logger.Infof("migrate version is %v", migrateVersion)
		os.Exit(0)
	}

	// translate -path into -source if given
	if source == "" && path != "" {
		source = fmt.Sprintf("file://%v", path)
	}

	// initialize migrate
	migrater, migraterErr = migrateV4.New(source, database)
	if migraterErr == nil {
		//migrater.Log = Logger // TODO
		migrater.PrefetchMigrations = prefetch
		migrater.LockTimeout = time.Duration(int64(lockTimeout)) * time.Second

		// handle Ctrl+c
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGINT)
		go func() {
			for range signals {
				fmt.Println("Stopping after this running migration ...")
				migrater.GracefulStop <- true
				return
			}
		}()
	} else {
		Logger.Errorf("preRun: %v", migraterErr)
		os.Exit(-1)
	}
	startTime = time.Now()
}

func postRun() {
	Logger.Infof("Finished after %v", time.Since(startTime))
	defer func() {
		if migraterErr == nil {
			if _, err := migrater.Close(); err != nil {
				Logger.Errorf("MigrateError:%v", err)
				fmt.Println(err)
			}
		}
	}()
}

func init() {
	migrateCmd.Flags().BoolVarP(&help, "help", "h", false, "Print usage")
	migrateCmd.Flags().BoolVarP(&version, "version", "v", false, "Print version")
	migrateCmd.PersistentFlags().UintVarP(&prefetch, "prefetch", "f", 10, "Number of migrations to load in advance before executing (default 10)")
	migrateCmd.PersistentFlags().UintVarP(&lockTimeout, "lockTimeout", "l", 15, "Allow N seconds to acquire database lock (default 15)")
	migrateCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "Shorthand for -source file://path")
	migrateCmd.PersistentFlags().StringVarP(&database, "database", "d", "", "Run migrations against this database (driver://url)")
	migrateCmd.PersistentFlags().StringVarP(&source, "source", "s", "", "Location of the migrations")

	migrateCmd.MarkPersistentFlagRequired("source")
	migrateCmd.MarkPersistentFlagRequired("database")

	rootCmd.AddCommand(migrateCmd)
	migrateCmd.AddCommand(upCmd)
	migrateCmd.AddCommand(downCmd)
	migrateCmd.AddCommand(gotoCmd)
	migrateCmd.AddCommand(lookCmd)
}
