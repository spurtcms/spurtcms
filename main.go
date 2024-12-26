package main

import (
	"os"
	"spurt-cms/controllers"
	"spurt-cms/graphql"
	"spurt-cms/migration"
	templateview "spurt-cms/page-view"
	"spurt-cms/routes"
	storagecontroller "spurt-cms/storage-controller"
	"sync"
)

func RunAdminPanel(wg *sync.WaitGroup) {

	defer wg.Done()

	r := routes.SetupRoutes() // setup routes

	err := r.Run(":" + os.Getenv("PORT"))

	if err != nil {

		panic(err)
	}
}

func main() {

	migration.TableMigration() //migration all spurtcms related tables

	migration.InsertDefaultValues() //insert default values....

	storagecontroller.LocalStorageCreation() //create all project related files and folders

	controllers.PackageInitialize() //initialize all spurtcms packages

	var wg sync.WaitGroup

	wg.Add(3)

	go RunAdminPanel(&wg)

	go templateview.RunTemplateView(&wg)

	go graphql.RunGraphqlServer(&wg)

	wg.Wait()

}
