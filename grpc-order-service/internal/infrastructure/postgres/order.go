package postgres

import (
    "context"; "fmt"
    "github.com/jackc/pgx/v5/pgxpool"
    "grpc-order-service/internal/domain"
)
type OrderRepository struct { db *pgxpool.Pool }
func NewOrderRepository(db *pgxpool.Pool)*OrderRepository{return &OrderRepository{db:db}}

func (r *OrderRepository) Create(ctx context.Context,o *domain.Order)error{
    tx,err:=r.db.Begin(ctx);if err!=nil{return err};defer tx.Rollback(ctx)
    _,err=tx.Exec(ctx,`INSERT INTO orders (id,customer_id,total_amount,currency,status,created_at,updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`,o.ID,o.CustomerID,o.TotalAmount,o.Currency,o.Status,o.CreatedAt,o.UpdatedAt);if err!=nil{return err}
    for _,i:=range o.Items{_,err=tx.Exec(ctx,`INSERT INTO order_items (order_id,product_id,quantity,price) VALUES ($1,$2,$3,$4)`,o.ID,i.ProductID,i.Quantity,i.Price);if err!=nil{return err}}
    return tx.Commit(ctx)
}
func (r *OrderRepository) Get(ctx context.Context,id string)(*domain.Order,error){
    var o domain.Order;var st string
    err:=r.db.QueryRow(ctx,`SELECT id,customer_id,total_amount,currency,status,created_at,updated_at FROM orders WHERE id=$1`,id).Scan(&o.ID,&o.CustomerID,&o.TotalAmount,&o.Currency,&st,&o.CreatedAt,&o.UpdatedAt)
    if err!=nil{return nil,fmt.Errorf("get order: %w",err)};o.Status=domain.OrderStatus(st);return &o,nil
}
func (r *OrderRepository) List(ctx context.Context,cid string,limit int)([]*domain.Order,error){
    rows,err:=r.db.Query(ctx,`SELECT id,customer_id,total_amount,currency,status,created_at,updated_at FROM orders WHERE customer_id=$1 ORDER BY created_at DESC LIMIT $2`,cid,limit);if err!=nil{return nil,err};defer rows.Close()
    var out []*domain.Order
    for rows.Next(){var o domain.Order;var st string;if err:=rows.Scan(&o.ID,&o.CustomerID,&o.TotalAmount,&o.Currency,&st,&o.CreatedAt,&o.UpdatedAt);err!=nil{return nil,err};o.Status=domain.OrderStatus(st);out=append(out,&o)}
    return out,rows.Err()
}
func (r *OrderRepository) UpdateStatus(ctx context.Context,id string,status domain.OrderStatus)error{
    tag,err:=r.db.Exec(ctx,`UPDATE orders SET status=$1,updated_at=NOW() WHERE id=$2`,status,id);if err!=nil{return err}
    if tag.RowsAffected()==0{return fmt.Errorf("order %q not found",id)};return nil
}
