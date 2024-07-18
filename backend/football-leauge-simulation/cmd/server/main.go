package main

import (
	"log"
	"net/http"

	"football-league-simulation/internal/config"
	"football-league-simulation/internal/handler"
	"football-league-simulation/internal/repository"
	"football-league-simulation/internal/service"
	"football-league-simulation/internal/util"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := util.NewDB(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories
	teamRepo := repository.NewTeamRepository(db)
	matchRepo := repository.NewMatchRepository(db)
	leagueRepo := repository.NewLeagueRepository(db)

	// Initialize services
	teamService := service.NewTeamService(teamRepo)
	matchService := service.NewMatchService(matchRepo)
	leagueService := service.NewLeagueService(leagueRepo)
	simulationService := service.NewSimulationService(teamService, matchService, leagueService)

	// Initialize handlers
	teamHandler := handler.NewTeamHandler(teamService)
	matchHandler := handler.NewMatchHandler(matchService)
	simulationHandler := handler.NewSimulationHandler(simulationService)

	// Initialize router
	router := mux.NewRouter()

	// CORS middleware
	router.Use(enableCORS)

	// Team routes
	router.HandleFunc("/teams", teamHandler.CreateTeam).Methods("POST")
	router.HandleFunc("/teams", teamHandler.GetTeams).Methods("GET")
	router.HandleFunc("/teams/{id}", teamHandler.GetTeamByID).Methods("GET")

	// Match routes
	router.HandleFunc("/matches", matchHandler.CreateMatch).Methods("POST")
	router.HandleFunc("/matches", matchHandler.GetMatches).Methods("GET")
	router.HandleFunc("/matches/{id}", matchHandler.GetMatchByID).Methods("GET")

	// Simulation routes
	router.HandleFunc("/simulate/week", simulationHandler.SimulateWeek).Methods("POST")
	router.HandleFunc("/simulate/season", simulationHandler.SimulateSeason).Methods("POST")

	// Start server
	log.Printf("Server is running on port %s", cfg.Server.Port)
	if err := http.ListenAndServe(":"+cfg.Server.Port, router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// enableCORS is a middleware function to enable CORS headers
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
