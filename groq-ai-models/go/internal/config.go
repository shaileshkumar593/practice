package internal
import "os"
import "github.com/joho/godotenv"

type Config struct{ APIKey,BaseURL,Model,ReasoningModel string }
func LoadConfig() Config {
 _=godotenv.Load()
 return Config{os.Getenv("GROQ_API_KEY"), getenv("GROQ_BASE_URL","https://api.groq.com/openai/v1"),
 getenv("GROQ_MODEL","openai/gpt-oss-120b"),getenv("GROQ_REASONING_MODEL","openai/gpt-oss-120b")}
}
func getenv(k,f string)string{if v:=os.Getenv(k);v!=""{return v};return f}
