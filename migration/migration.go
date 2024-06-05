package migration

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"spurt-cms/config"
	"spurt-cms/migration/mysql"
	"spurt-cms/migration/postgres"
	"strings"
)

func TableMigration() {

	if os.Getenv("DATABASE_TYPE") == "postgres" {

		postgres.MigrationTables() //auto migrate table

	} else if os.Getenv("DATABASE_TYPE") == "mysql" {

		mysql.MigrationTables() //auto migrate table
	}

}

func InsertDefaultValues() {

	// Open the SQL file
	file, err := os.Open("cms.sql")

	if err != nil {

		log.Println(err)
	}

	defer file.Close()

	// Start a transaction
	db := config.SetupDB()

	// Create a new scanner to read the SQL file line by line
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		query := scanner.Text()

		// Skip empty lines and comments
		if len(strings.TrimSpace(query)) == 0 || strings.HasPrefix(query, "--") {

			continue
		}

		// Execute the SQL query
		err := db.Exec(query)

		if err != nil {

			fmt.Println("Error executing query: \nQuery:", err, query)
		}
	}

	var permissionmaxID int

	var rolemaxID int

	var usermaxID int

	var moduleId int

	var ChannelMaxid int

	var CategoriesMaxId int

	var LangMaxId int

	var ChannelCatMaxId int

	var GroupMaxId int

	if os.Getenv("DATABASE_TYPE") == "postgres" {

		if err := db.Raw("SELECT max(id) FROM tbl_modules").Row().Scan(&moduleId); err != nil {
			// Handle error
			log.Println(err)
		}

		if err := db.Raw("SELECT max(id) FROM tbl_module_permissions").Row().Scan(&permissionmaxID); err != nil {
			// Handle error
			log.Println(err)
		}

		if err := db.Raw("SELECT max(id) FROM tbl_roles").Row().Scan(&rolemaxID); err != nil {
			// Handle error
			log.Println(err)
		}

		if err := db.Raw("SELECT max(id) FROM tbl_users").Row().Scan(&usermaxID); err != nil {
			// Handle error
			log.Println(err)
		}

		if err := db.Raw("SELECT max(id) FROM tbl_channels").Row().Scan(&ChannelMaxid); err != nil {
			// Handle error
			log.Println(err)
		}

		if err := db.Raw("SELECT max(id) FROM tbl_categories").Row().Scan(&CategoriesMaxId); err != nil {
			// Handle error
			log.Println(err)
		}

		if err := db.Raw("SELECT max(id) FROM tbl_languages").Row().Scan(&LangMaxId); err != nil {
			// Handle error
			log.Println(err)
		}

		if err := db.Raw("SELECT max(id) FROM tbl_channel_categories").Row().Scan(&ChannelCatMaxId); err != nil {
			// Handle error
			log.Println(err)
		}

		if err := db.Raw("SELECT max(id) FROM tbl_member_groups").Row().Scan(&GroupMaxId); err != nil {
			// Handle error
			log.Println(err)
		}

		db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_modules_id_seq RESTART WITH %d", moduleId+1))

		db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_module_permissions_id_seq RESTART WITH %d", permissionmaxID+1))

		db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_roles_id_seq RESTART WITH %d", rolemaxID+1))

		db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_users_id_seq RESTART WITH %d", usermaxID+1))

		db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_channels_id_seq RESTART WITH %d", ChannelMaxid+1))

		db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_categories_id_seq RESTART WITH %d", CategoriesMaxId+1))

		db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_languages_id_seq RESTART WITH %d", LangMaxId+1))

		db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_channel_categories_id_seq RESTART WITH %d", ChannelCatMaxId+1))

		db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_member_groups_id_seq RESTART WITH %d", GroupMaxId+1))

	}

	fmt.Println("SQL file loaded successfully.")
}
