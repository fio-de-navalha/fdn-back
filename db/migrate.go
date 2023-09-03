package migrations

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
	"time"

	"github.com/fio-de-navalha/fdn-back/internal/infra/database"
	"golang.org/x/exp/slices"
	"gorm.io/gorm"
)

type Migrations struct {
	ID        uint `gorm:"primaryKey"`
	File      string
	CreatedAt time.Time
}

func Migrate() error {
	db := database.DB
	if err := db.AutoMigrate(&Migrations{}); err != nil {
		return err
	}

	runMigrations(db)

	return nil
}

func runMigrations(db *gorm.DB) error {
	basePath := "./db/"

	migrationFiles := getRanMigrations(db)
	files := getMigrationFiles(basePath)

	for _, file := range files {
		if strings.Contains(file.Name(), ".sql") && !slices.Contains(migrationFiles, file.Name()) {
			executeMigrationFile(db, basePath, file)
		}
	}

	return nil
}

func getRanMigrations(db *gorm.DB) []string {
	var migrationFile []string
	db.Table("migrations").Select("file").Find(&migrationFile)
	return migrationFile
}

func getMigrationFiles(basePath string) []fs.DirEntry {
	files, err := os.ReadDir(basePath)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return files
}

func executeMigrationFile(db *gorm.DB, basePath string, file fs.DirEntry) {
	c, err := os.ReadFile(basePath + file.Name())
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println("Running migration for: ", file.Name())
	db.Exec(string(c))
	db.Exec("INSERT INTO migrations (file, created_at) VALUES (?, ?)", file.Name(), time.Now())
}
