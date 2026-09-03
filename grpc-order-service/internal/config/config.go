package config
import "os"
type Config struct{Port,DatabaseURL,RedisURL,Environment string}
func Load()Config{p:=os.Getenv("PORT");if p==""{p="50051"};return Config{Port:p,DatabaseURL:os.Getenv("DATABASE_URL"),RedisURL:os.Getenv("REDIS_URL"),Environment:os.Getenv("ENVIRONMENT")}}
