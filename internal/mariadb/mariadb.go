package mariadb

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var recipesSchema = `CREATE TABLE IF NOT EXISTS Recipes(
  ID int UNSIGNED NOT NULL AUTO_INCREMENT,
  Name varchar(50) NOT NULL,
  URL varchar(2083) NOT NULL,
  PRIMARY KEY (ID)
);
`

var recipeIngredientsSchema = `CREATE TABLE IF NOT EXISTS RecipeIngredients(
  ID int UNSIGNED NOT NULL AUTO_INCREMENT,
  Ingredient int UNSIGNED,
  Amount float NOT NULL,
  Unit varchar(25) NOT NULL,
  PRIMARY KEY (ID),
  FOREIGN KEY (Ingredient) REFERENCES Ingredients(ID)
);
`

var ingredientMembershipSchema = `CREATE TABLE IF NOT EXISTS IngredientMemberships(
  ID int UNSIGNED NOT NULL AUTO_INCREMENT,
  Ingredient int UNSIGNED,
  Recipe int UNSIGNED,
  PRIMARY KEY (ID),
  FOREIGN KEY (Ingredient) REFERENCES RecipeIngredients(ID),
  FOREIGN KEY (Recipe) REFERENCES Recipes(ID)
);
`

var ingredientsSchema = `CREATE TABLE IF NOT EXISTS Ingredients(
  ID int UNSIGNED NOT NULL AUTO_INCREMENT,
  Name varchar(50) NOT NULL,
  Recipes text(65535),
  PRIMARY KEY (ID)
);
`

func Connect() {
	db, err := sqlx.Connect("mysql",  "rotr:rotrpwd@tcp(127.0.0.1:3306)/rotr")
	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	db.MustExec(ingredientsSchema)
	db.MustExec(recipeIngredientsSchema)
	db.MustExec(recipesSchema)
	db.MustExec(ingredientMembershipSchema)
}