package main

import (
	pb "SortingMapReduce/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"sort"
	"sync"
)

var orderChunk []int32
var shuffleMap map[string][]int32
var inputReduce map[string][]int32
var lock sync.Mutex

type WorkerServer struct {
	pb.UnimplementedWorkerServiceServer
}

func (w *WorkerServer) MapExecute(ctx context.Context, req *pb.ArrayInt32) (*pb.Bool, error) {
	orderChunk = make([]int32, 0)
	log.Printf("Fase Map ricevuti dati: %v", req.Data)
	data := req.Data
	sort.Slice(data, func(i, j int) bool { return data[i] < data[j] })
	orderChunk = data
	return nil, nil
}

func (w *WorkerServer) StartShuffle(ctx context.Context, partition *pb.Partition) (*pb.Bool, error) {
	inputReduce = make(map[string][]int32)
	shuffleMap = make(map[string][]int32)
	var indexMin = 0
	var i = 0

	for _, worker := range partition.Partition {
		indexMin = i
		for _, value := range orderChunk[indexMin:] {
			if value > worker.Max {
				break
			}
			i++
		}
		shuffleMap[worker.Worker] = orderChunk[indexMin:i]
	}
	log.Printf("Shuffle map: %v", shuffleMap)

	for worker, reduceInput := range shuffleMap {
		conn, err := grpc.Dial(worker, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Impossibile connettersi al worker %s: %v", worker, err)
		}
		client := pb.NewWorkerServiceClient(conn)

		partialInputReduce := &pb.PartialInputReduce{Worker: partition.Worker,
			Data: reduceInput}
		client.AddPartialInputReduce(context.Background(), partialInputReduce)
		if false {
			log.Fatalf("Errore durante l'invio dei dati per il reduce: %v", err)
		}
	}

	return nil, nil
}
func (w *WorkerServer) AddPartialInputReduce(ctx context.Context, partialInputReduce *pb.PartialInputReduce) (*pb.Bool, error) {
	lock.Lock()
	inputReduce[partialInputReduce.Worker] = partialInputReduce.Data
	lock.Unlock()

	return nil, nil

}
func (w *WorkerServer) Reduce(ctx context.Context, worker *pb.Worker) (*pb.Bool, error) {
	log.Printf("Input completo di reduce: %v", inputReduce)
	var orderData string
	result := mergeDivideAndConquer(inputReduce)
	orderData = fmt.Sprintf("%v", result)

	log.Printf("Dati ordinati: %v", orderData)
	filename := os.Getenv("RESULT_FILE")

	err := os.WriteFile(filename, []byte(orderData), 0644)
	if err != nil {
		log.Printf("Errore nella scrittura del file")
	}
	return nil, nil
}

func mergeSortedArrays(a, b []int32) []int32 {
	result := make([]int32, 0, len(a)+len(b))
	i, j := 0, 0

	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			result = append(result, a[i])
			i++
		} else {
			result = append(result, b[j])
			j++
		}
	}
	result = append(result, a[i:]...)
	result = append(result, b[j:]...)

	return result
}
func mergeDivideAndConquer(inputReduce map[string][]int32) []int32 {
	chunks := make([][]int32, 0, len(inputReduce))
	for _, chunk := range inputReduce {
		chunks = append(chunks, chunk)
	}
	for len(chunks) > 1 {
		var newChunks [][]int32

		// Unisci gli array a coppie
		for i := 0; i < len(chunks); i += 2 {
			if i+1 < len(chunks) {
				// Merge tra due array
				merged := mergeSortedArrays(chunks[i], chunks[i+1])
				newChunks = append(newChunks, merged)
			} else {
				// Aggiungi l'ultimo array se è dispari
				newChunks = append(newChunks, chunks[i])
			}
		}
		chunks = newChunks
	}
	return chunks[0]
}

// Main per avviare il Worker
func main() {
	portPtr := flag.Int("port", 8888, "Port")
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *portPtr))
	if err != nil {
		log.Fatalf("Errore nel creare il listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	worker := &WorkerServer{}

	pb.RegisterWorkerServiceServer(grpcServer, worker)
	log.Printf("Worker in ascolto su :%d", *portPtr)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Errore nel servire il Worker: %v", err)
	}
}
