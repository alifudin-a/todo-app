// Code generated by sqlc. DO NOT EDIT.

package db

import ()

type Task struct {
	ID          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Complete    string `json:"complete"`
}