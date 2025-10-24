package migration

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"spurt-cms/config"
	"spurt-cms/migration/mysql"
	"spurt-cms/migration/postgres"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
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

		currentTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().UTC().Format("2006-01-02 15:04:05"))

		parts := strings.Fields(fmt.Sprint(currentTime))

		timeStamp := fmt.Sprintf("%s %s", parts[0], parts[1])

		uuid := (uuid.New()).String()

		arr := strings.Split(uuid, "-")

		Uuid := arr[len(arr)-1]

		query = strings.ReplaceAll(query, "current_time", timeStamp)

		baseurl := os.Getenv("BASE_URL")
		baseurl = strings.TrimPrefix(baseurl, "https://")
		baseurl = strings.TrimPrefix(baseurl, "http://")
		baseurl = strings.TrimSuffix(baseurl, "/")

		query = strings.ReplaceAll(query, "base_url", baseurl)

		if strings.Contains(query, "tbl_forms") {

			query = strings.ReplaceAll(query, "uu_id", Uuid)
		}

		if strings.Contains(query, "tbl_block_tags") {

			var defBlockIds []int

			if err = db.Table("tbl_blocks").Select("id").Where("is_deleted=0 and tenant_id is NULL and block_description=?", "spurtcms").Scan(&defBlockIds).Error; err != nil {
				log.Println(err)
			}

			for _, val := range defBlockIds {
				modifiedQuery := strings.ReplaceAll(query, "blo_id", strconv.Itoa(val))

				var exists bool
				db.Table("tbl_block_tags").Select("count(*) > 0").Where("block_id = ?", val).Scan(&exists)

				if !exists {

					err := db.Exec(modifiedQuery).Error
					if err != nil {
						fmt.Println("Error executing query: \nQuery:", modifiedQuery, "\nError:", err)
					}
				} else {
				}
			}
			continue
		}

		err := db.Exec(query).Error

		if err != nil {

			fmt.Println("Error executing query: \nQuery:", err, query)
		}
	}

	if os.Getenv("DATABASE_TYPE") == "postgres" {

		var (
			module_count, mod_perm_count, role_count int64

			user_count, channel_count, category_count int64

			language_count, chancat_count, memgroup_count int64

			personalize_count, block_count, blockmstrTag_count int64

			blocktag_count, blockcolln_count, gensetting_count int64

			template_count, tempmodule_count, emailConf_count, emailTemp_count, form_count, website_count, gotemp_count, Subscription_count int64
		)

		if err := db.Table("tbl_modules").Count(&module_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_module_permissions").Count(&mod_perm_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_roles").Count(&role_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_users").Count(&user_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_channels").Count(&channel_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_categories").Count(&category_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_channel_categories").Count(&chancat_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_member_groups").Count(&memgroup_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_user_personalizes").Count(&personalize_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_blocks").Count(&block_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_block_mstr_tags").Count(&blockmstrTag_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_block_tags").Count(&blocktag_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_block_collections").Count(&blockcolln_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_general_settings").Count(&gensetting_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_templates").Count(&template_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_template_modules").Count(&tempmodule_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_email_configurations").Count(&emailConf_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_email_templates").Count(&emailTemp_count).Error; err != nil {

			log.Println(err)
		}
		if err := db.Table("tbl_forms").Count(&form_count).Error; err != nil {

			log.Println(err)
		}

		if err := db.Table("tbl_mstr_membershiplevels").Count(&Subscription_count).Error; err != nil {

			log.Println(err)
		}
		if err := db.Table("tbl_websites").Count(&website_count).Error; err != nil {

			log.Println(err)
		}
		if err := db.Table("tbl_go_templates").Count(&gotemp_count).Error; err != nil {

			log.Println(err)
		}
		if module_count > 0 {

			var moduleId int

			if err := db.Raw("SELECT max(id) FROM tbl_modules").Row().Scan(&moduleId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_modules_id_seq RESTART WITH %d", moduleId+1))

		}

		if mod_perm_count > 0 {

			var permissionmaxID int

			if err := db.Raw("SELECT max(id) FROM tbl_module_permissions").Row().Scan(&permissionmaxID); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_module_permissions_id_seq RESTART WITH %d", permissionmaxID+1))

		}

		if role_count > 0 {

			var rolemaxID int

			if err := db.Raw("SELECT max(id) FROM tbl_roles").Row().Scan(&rolemaxID); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_roles_id_seq RESTART WITH %d", rolemaxID+1))

		}

		if user_count > 0 {

			var usermaxID int

			if err := db.Raw("SELECT max(id) FROM tbl_users").Row().Scan(&usermaxID); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_users_id_seq RESTART WITH %d", usermaxID+1))

		}

		if channel_count > 0 {

			var ChannelMaxid int

			if err := db.Raw("SELECT max(id) FROM tbl_channels").Row().Scan(&ChannelMaxid); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_channels_id_seq RESTART WITH %d", ChannelMaxid+1))

		}

		if category_count > 0 {

			var CategoriesMaxId int

			if err := db.Raw("SELECT max(id) FROM tbl_categories").Row().Scan(&CategoriesMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_categories_id_seq RESTART WITH %d", CategoriesMaxId+1))

		}

		if language_count > 0 {

			var LangMaxId int

			if err := db.Raw("SELECT max(id) FROM tbl_languages").Row().Scan(&LangMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_languages_id_seq RESTART WITH %d", LangMaxId+1))

		}

		if chancat_count > 0 {

			var ChannelCatMaxId int

			if err := db.Raw("SELECT max(id) FROM tbl_channel_categories").Row().Scan(&ChannelCatMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_channel_categories_id_seq RESTART WITH %d", ChannelCatMaxId+1))

		}

		if memgroup_count > 0 {

			var GroupMaxId int

			if err := db.Raw("SELECT max(id) FROM tbl_member_groups").Row().Scan(&GroupMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_member_groups_id_seq RESTART WITH %d", GroupMaxId+1))

		}

		if personalize_count > 0 {

			var UserPersonalizeId int

			if err := db.Raw("SELECT max(id) FROM tbl_user_personalizes").Row().Scan(&UserPersonalizeId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_user_personalizes_id_seq RESTART WITH %d", UserPersonalizeId+1))

		}

		if block_count > 0 {

			var BlockMaxId int

			if err := db.Raw("SELECT max(id) FROM tbl_blocks").Row().Scan(&BlockMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_blocks_id_seq RESTART WITH %d", BlockMaxId+1))

		}

		if blockmstrTag_count > 0 {

			var MstrTagMaxId int

			if err := db.Raw("SELECT max(id) FROM tbl_block_mstr_tags").Row().Scan(&MstrTagMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_block_mstr_tags_id_seq RESTART WITH %d", MstrTagMaxId+1))

		}

		if blocktag_count > 0 {

			var TagMaxId int

			if err := db.Raw("SELECT max(id) FROM tbl_block_tags").Row().Scan(&TagMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_block_tags_id_seq RESTART WITH %d", TagMaxId+1))

		}

		if blockcolln_count > 0 {

			var CollectionMaxId int

			if err := db.Raw("SELECT max(id) FROM tbl_block_collections").Row().Scan(&CollectionMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_block_collections_id_seq RESTART WITH %d", CollectionMaxId+1))

		}

		if gensetting_count > 0 {

			var GenSettingMaxId int

			if err := db.Raw("SELECT COALESCE((select max(id) from tbl_general_settings),0)").Scan(&GenSettingMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_general_settings_id_seq RESTART WITH %v", GenSettingMaxId+1))

		}

		if template_count > 0 {

			var TemplateMaxId int

			if err := db.Raw("SELECT COALESCE((select max(id) from tbl_templates),0)").Scan(&TemplateMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_templates_id_seq RESTART WITH %v", TemplateMaxId+1))

		}

		if tempmodule_count > 0 {

			var templateModuleMaxId int

			if err := db.Raw("SELECT COALESCE((select max(id) from tbl_template_modules),0)").Scan(&templateModuleMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_template_modules_id_seq RESTART WITH %v", templateModuleMaxId+1))

		}

		if emailConf_count > 0 {

			var emailConfMaxId int

			if err := db.Raw("SELECT COALESCE((select max(id) from tbl_email_configurations),0)").Scan(&emailConfMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_email_configurations_id_seq RESTART WITH %v", emailConfMaxId+1))
		}

		if emailTemp_count > 0 {

			var emailTempMaxId int

			if err := db.Raw("SELECT COALESCE((select max(id) from tbl_email_templates),0)").Scan(&emailTempMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_email_templates_id_seq RESTART WITH %v", emailTempMaxId+1))
		}

		if Subscription_count > 0 {

			var subscriptionMaxId int

			if err := db.Raw("SELECT COALESCE((select max(id) from tbl_mstr_membershiplevels),0)").Scan(&subscriptionMaxId); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_mstr_membershiplevels_id_seq RESTART WITH %v", subscriptionMaxId+1))
		}
		if website_count > 0 {

			var WebsiteMaxid int

			if err := db.Raw("SELECT max(id) FROM tbl_websites").Row().Scan(&WebsiteMaxid); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_websites_id_seq RESTART WITH %d", WebsiteMaxid+1))

		}

		if gotemp_count > 0 {

			var gotempMaxid int

			if err := db.Raw("SELECT max(id) FROM tbl_go_templates").Row().Scan(&gotempMaxid); err != nil {
				// Handle error
				log.Println(err)
			}

			db.Exec(fmt.Sprintf("ALTER SEQUENCE tbl_go_templates_id_seq RESTART WITH %d", gotempMaxid+1))

		}

	}

	fmt.Println("SQL file loaded successfully.")
}
