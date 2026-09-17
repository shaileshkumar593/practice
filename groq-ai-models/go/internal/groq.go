package internal

import("bytes";"context";"encoding/json";"fmt";"net/http";"time")

type Client struct{BaseURL,APIKey string; HTTP *http.Client}
type Message struct{Role,Content string; ToolCalls []ToolCall `json:"tool_calls,omitempty"`; ToolCallID string `json:"tool_call_id,omitempty"`}
type Tool struct{Type string `json:"type"`; Function FunctionDef `json:"function"`}
type FunctionDef struct{Name,Description string; Parameters map[string]interface{} `json:"parameters"`}
type ToolCall struct{ID,Type string; Function struct{Name,Arguments string} `json:"function"`}
type Request struct{Model string `json:"model"`; Messages []Message `json:"messages"`; Temperature float64 `json:"temperature,omitempty"`; Tools []Tool `json:"tools,omitempty"`; ToolChoice string `json:"tool_choice,omitempty"`}
type Response struct{Choices []struct{Message Message `json:"message"`} `json:"choices"`}

func NewClient(c Config)*Client{return &Client{c.BaseURL,c.APIKey,&http.Client{Timeout:60*time.Second}}}
func(c *Client)Chat(ctx context.Context,r Request)(Message,error){
 b,e:=json.Marshal(r);if e!=nil{return Message{},e}
 q,e:=http.NewRequestWithContext(ctx,"POST",c.BaseURL+"/chat/completions",bytes.NewReader(b));if e!=nil{return Message{},e}
 q.Header.Set("Authorization","Bearer "+c.APIKey);q.Header.Set("Content-Type","application/json")
 resp,e:=c.HTTP.Do(q);if e!=nil{return Message{},e};defer resp.Body.Close()
 if resp.StatusCode<200||resp.StatusCode>=300{return Message{},fmt.Errorf("groq HTTP %d",resp.StatusCode)}
 var out Response;if e=json.NewDecoder(resp.Body).Decode(&out);e!=nil{return Message{},e}
 if len(out.Choices)==0{return Message{},fmt.Errorf("no choices")};return out.Choices[0].Message,nil
}
