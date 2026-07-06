package service

import "orders/internal/repository"

func GetAll() map[string]interface{} {
return map[string]interface{}{
"service": "orders",
"data": repository.FindAll(),
}
}
