package models

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/ouqiang/gocron/internal/modules/setting"
)

func TestSQLiteMigrationInstall(t *testing.T) {
	dir, err := ioutil.TempDir("", "gocron-sqlite-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	s := &setting.Setting{}
	s.Db.Engine = "sqlite"
	s.Db.Database = filepath.Join(dir, "gocron.db")

	engine, err := CreateTmpDb(s)
	if err != nil {
		t.Fatal(err)
	}
	defer engine.Close()

	if engine.DriverName() != "sqlite3" {
		t.Fatalf("expected sqlite3 driver, got %s", engine.DriverName())
	}

	oldDb := Db
	oldTablePrefix := TablePrefix
	Db = engine
	TablePrefix = ""
	defer func() {
		Db = oldDb
		TablePrefix = oldTablePrefix
	}()

	if err = new(Migration).Install(s.Db.Database); err != nil {
		t.Fatal(err)
	}

	exist, err := Db.IsTableExist(new(Task))
	if err != nil {
		t.Fatal(err)
	}
	if !exist {
		t.Fatal("expected task table to exist")
	}
}
