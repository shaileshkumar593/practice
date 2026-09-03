package config

import "os"

type Config struct {
    Port        string
    Neo4jURI    string
    Neo4jUser   string
    Neo4jPass   string
    Environment string
}

func Load() Config {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    return Config{
        Port:        port,
        Neo4jURI:    os.Getenv("NEO4J_URI"),
        Neo4jUser:   os.Getenv("NEO4J_USER"),
        Neo4jPass:   os.Getenv("NEO4J_PASSWORD"),
        Environment: os.Getenv("ENVIRONMENT"),
    }
}
