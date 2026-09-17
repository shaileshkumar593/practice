package main

import("context";"encoding/json";"fmt";"log";"groq-ai-models/internal")

func main(){
 c:=internal.LoadConfig();if c.APIKey==""{log.Fatal("GROQ_API_KEY is missing")}
 client:=internal.NewClient(c);ctx:=context.Background()

 msg,err:=client.Chat(ctx,internal.Request{Model:c.Model,Messages:[]internal.Message{
  {Role:"system",Content:"You are a precise technical assistant."},
  {Role:"user",Content:"Explain RAG in five points."}}});must(err)
 fmt.Println("LLM:\n",msg.Content)

 msg,err=client.Chat(ctx,internal.Request{Model:c.ReasoningModel,Messages:[]internal.Message{
  {Role:"user",Content:"Design safe payment retry handling after a gateway timeout. Cover idempotency, webhooks and reconciliation."}}});must(err)
 fmt.Println("\nREASONING:\n",msg.Content)

 tool:=internal.Tool{Type:"function",Function:internal.FunctionDef{
  Name:"get_order_status",Description:"Get order status.",
  Parameters:map[string]interface{}{"type":"object","properties":map[string]interface{}{"order_id":map[string]interface{}{"type":"string"}},"required":[]string{"order_id"}}}}
 msgs:=[]internal.Message{{Role:"user",Content:"What is order ORD-1001 status?"}}
 msg,err=client.Chat(ctx,internal.Request{Model:c.Model,Messages:msgs,Tools:[]internal.Tool{tool},ToolChoice:"auto"});must(err)
 msgs=append(msgs,msg)
 for _,call:=range msg.ToolCalls{
  if call.Function.Name=="get_order_status"{
   var a struct{OrderID string `json:"order_id"`};must(json.Unmarshal([]byte(call.Function.Arguments),&a))
   b,_:=json.Marshal(map[string]string{"order_id":a.OrderID,"status":"SHIPPED"})
   msgs=append(msgs,internal.Message{Role:"tool",ToolCallID:call.ID,Content:string(b)})
  }
 }
 msg,err=client.Chat(ctx,internal.Request{Model:c.Model,Messages:msgs,Tools:[]internal.Tool{tool}});must(err)
 fmt.Println("\nTOOL CALLING:\n",msg.Content)

 fmt.Println(`
EMBEDDING: call a dedicated embedding provider/local model -> vector DB
RERANKER: retrieve candidates -> cross-encoder -> relevance scores -> top K
MULTIMODAL: send supported image/audio/text input to a multimodal model
RAG: query -> retrieval -> reranking -> context -> Groq LLM
AGENT: model -> controlled tool -> result -> model loop
`)
}
func must(e error){if e!=nil{log.Fatal(e)}}
