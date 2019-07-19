package migrate

import (
	"errors"
	"fmt"
	"github.com/bbcyyb/pcrs/pkg/log"
	migrateV4 "github.com/golang-migrate/migrate/v4"
	"strconv"
	"strings"
)

func ExecuteUp(migrater *migrateV4.Migrate, arg string) {
	limit := -1
	if arg != "" {
		n, err := strconv.ParseUint(arg, 10, 64)
		if err != nil {
			log.Error("error: can't read limit argument N")
			return
		}
		limit = int(n)
	}
	upVersion(migrater, limit)
}

func ExecuteDown(migrater *migrateV4.Migrate, downArg string) {
	num, needsConfirm, err := numDownMigrationsFromArgs(downArg)
	if err != nil {
		log.Error(err)
		return
	}
	if needsConfirm {
		log.Info("Are you sure you want to apply all down migrations? [y/N]")
		var response string
		fmt.Scanln(&response)
		response = strings.ToLower(strings.TrimSpace(response))

		if response == "y" {
			log.Info("Applying all down migrations")
		} else {
			log.Error("Not applying all down migrations")
		}
	}
	downVersion(migrater, num)
}

func ExecuteGoto(migrater *migrateV4.Migrate, arg string) {
	v, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		log.Error("error: can't read version argument V")
		return
	}

	gotoVersion(migrater, uint(v))
}

func ExecuteLook(migrater *migrateV4.Migrate) {
	versionLook(migrater)
}

// numDownMigrationsFromArgs returns an int for number of migrations to apply and a bool indicating if we need a confirm before applying
func numDownMigrationsFromArgs(arg string) (int, bool, error) {
	if strings.TrimSpace(arg) == "all" {
		return -1, true, nil
	}

	num, err := strconv.ParseUint(arg, 10, 64)
	if err != nil {
		return 0, false, errors.New("can't read limit argument N, only accept num or 'all'")
	}
	return int(num), false, nil
}

func upVersion(m *migrateV4.Migrate, limit int) {
	if limit >= 0 {
		if err := m.Steps(limit); err != nil {
			log.Error(err)
			return
		}
	} else {
		if err := m.Up(); err != nil {
			log.Error(err)
			return
		}
	}
	versionLook(m)
}

func downVersion(m *migrateV4.Migrate, limit int) {
	if limit >= 0 {
		if err := m.Steps(-limit); err != nil {
			log.Error(err)
			return
		}
	} else {
		if err := m.Down(); err != nil {
			log.Error(err)
			return
		}
	}
	versionLook(m)
}

func gotoVersion(m *migrateV4.Migrate, v uint) {
	if err := m.Migrate(v); err != nil {
		log.Error(err)
		return
	}
	versionLook(m)
}

func versionLook(m *migrateV4.Migrate) {
	v, dirty, err := m.Version()
	if err != nil {
		log.Error(err)
	}
	if dirty {
		log.Errorf("version %v is (dirty)\n", v)
	} else {
		log.Infof("current version is:%v", v)
	}
}

func SampleTest() {
	connectString := "sqlserver://PowerCalc:Power%401433@10.35.83.61:1433?database=PowerCalcFor46"
	m, err := migrateV4.New(
		"file://migrate_sql",
		connectString)
	defer m.Close()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	m.Steps(2)
	log.Info(connectString)
}
