
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/darcys22/godbledger/proto"
	"google.golang.org/grpc"
	"math"
	"math/rand"
)

const (
	address       = "localhost:50051"
	iterationDays = 365
	sdBPS         = 500
	decimalsBPS   = 10000
	tradProb      = 10
	startPrice    = 100.00
)

// Account holds the name and balance
type Trade struct {
	amount int
	price  float64