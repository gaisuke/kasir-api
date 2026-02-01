package main

import (
	"encoding/json"
	"fmt"
	"kasir-api/database"
	"kasir-api/handlers"
	"kasir-api/repositories"
	"kasir-api/services"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

const (
	apiRoutes           = "/api"
	apiProductsRoutes   = apiRoutes + "/products"
	apiCategoriesRoutes = apiRoutes + "/categories"
)

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database: ", err)
	}
	defer db.Close()

	productsRepository := repositories.NewProductRepository(db)
	productsService := services.NewProductService(productsRepository)
	productHandler := handlers.NewProductHandler(productsService)

	categoryRepository := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepository)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	mux := http.NewServeMux()

	mux.HandleFunc(apiCategoriesRoutes, categoryHandler.HandleCategories)
	mux.HandleFunc(apiCategoriesRoutes+"/", categoryHandler.HandleCategoriesById)

	mux.HandleFunc(apiProductsRoutes, productHandler.HandleProducts)
	mux.HandleFunc(apiProductsRoutes+"/", productHandler.HandleProductsById)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	addr := "0.0.0.0:" + config.Port
	fmt.Println("Server running on => ", addr)

	err = http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
