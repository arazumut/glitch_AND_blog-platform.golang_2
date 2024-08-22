package main

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
    Posts    []Post
}

type Post struct {
    gorm.Model
    Title    string
    Content  string
    UserID   uint
    Comments []Comment
}

type Comment struct {
    gorm.Model
    Content string
    PostID  uint
}
