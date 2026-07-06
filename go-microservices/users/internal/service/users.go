package service

import "users/internal/repository"

func GetAll() map[string]interface{} {
return map[string]interface{}{
"service": "users",
"data": repository.FindAll(),
}
}
