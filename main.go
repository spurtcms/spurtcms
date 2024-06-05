package main

import (
	"os"
	"spurt-cms/controllers"
	"spurt-cms/migration"
	"spurt-cms/routes"
	storagecontroller "spurt-cms/storage-controller"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	storagecontroller.LocalStorageCreation() //create all project related files and folders

	migration.TableMigration() //migration all spurtcms related tables

	migration.InsertDefaultValues() //insert default values....

	controllers.PackageInitialize() //initialize all spurtcms packages

	r := routes.SetupRoutes() // setup routes

	err := r.Run(":" + os.Getenv("PORT"))

	if err != nil {

		panic(err)
	}

}
