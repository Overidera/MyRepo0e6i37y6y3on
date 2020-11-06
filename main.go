
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