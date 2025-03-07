package config

type Config struct {
    DBHost     string
    DBUser     string
    DBPassword string
    DBName     string
    RedisURL   string
    ServerPort string
}

func Load() (*Config, error) {
    return &Config{
        DBHost:     "localhost",
        DBUser:     "postgres",
        DBPassword: "postgres",
        DBName:     "coderank",
        RedisURL:   "localhost:6379",
        ServerPort: ":8080",
    }, nil
} 