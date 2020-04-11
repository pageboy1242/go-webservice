package models

import (
    "fmt"
    "testing"
)

func TestGetUsers(t *testing.T) {
    u := GetUsers()
    
    fmt.Printf("TestGetUsers")
    
    if len(u) != 0 {
        t.Errorf("Expected empty array")
    }
}
