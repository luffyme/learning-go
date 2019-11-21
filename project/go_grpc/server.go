package main

import (
	"context"
	"log"
	"net"

	"google.golang.net/grpc"

	pb "go_grpc/proto"
)

type SearchService struct{}


