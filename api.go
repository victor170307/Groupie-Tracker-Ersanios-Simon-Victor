package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

var client = &http.Client{
	Timeout: 10 * time.Second,
}

type LocationsResponse struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type DatesResponse struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

