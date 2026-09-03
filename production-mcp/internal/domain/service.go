package domain

type Service struct {
    ID           string   `json:"id"`
    Name         string   `json:"name"`
    Version      string   `json:"version"`
    Status       string   `json:"status"`
    Owner        string   `json:"owner"`
    Dependencies []string `json:"dependencies"`
}
