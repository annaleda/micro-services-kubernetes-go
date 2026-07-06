package service

import "cart/internal/repository"

func GetAll() map[string]interface{} {
return map[string]interface{}{
"service": "cart",
"data": repository.FindAll(),
}
}
