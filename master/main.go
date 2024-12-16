package main

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"math/rand"
	"net"
	"os"
	"sync"

	pb "SortingMapReduce/proto"
)

type MasterServer struct {
	pb.UnimplementedMasterServiceServer
	workers []string // Indirizzi dei worker
}

// Run avvia il processo MapReduce
func (m *MasterServer) Run(data []int32) {
	nmin := data[0]
	nmax := data[0]
	for _, num := range data {
		if num < nmin {
			nmin = num
		}
		if num > nmax {
			nmax = num
		}
	}
	// Numero di worker
	numWorkers := len(m.workers)
	if numWorkers == 0 {
		log.Fatalf("Nessun worker disponibile")
	}

	// Suddividi i dati in chunk per i worker
	chunkSize := (len(data) + numWorkers - 1) / numWorkers
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	// Fase di Map
	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[start:end]
		workerAddr := m.workers[i]

		go func(workerAddr string, chunk []int32) {
			defer wg.Done()

			conn, client := createWorkerClient(workerAddr)
			defer conn.Close()

			req := &pb.ArrayInt32{
				Data: chunk,
			}

			_, err := client.MapExecute(context.Background(), req)
			if err != nil {
				log.Fatalf("Errore durante la fase di map: %v", err)
			}
		}(workerAddr, chunk)
	}
	wg.Wait()

	//calcola partizioni delle chiavi
	totalMin, totalMax := nmin, nmax
	rangeSize := (totalMax - totalMin + 1) / int32(numWorkers)
	var partition []*pb.Tuple
	for i := 0; i < numWorkers; i++ {
		key_min := totalMin + int32(i)*rangeSize
		key_max := key_min + rangeSize - 1
		if i == numWorkers-1 {
			key_max = totalMax // L'ultimo worker prende fino al massimo
		}
		tuple := &pb.Tuple{Worker: m.workers[i], Max: key_max, Min: key_min}
		partition = append(partition, tuple)
	}

	//Fase di shuffle
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		workerAddr := m.workers[i]
		go func(workerAddr string, partition []*pb.Tuple) {
			defer wg.Done()
			req := &pb.Partition{Partition: partition,
				Worker: workerAddr}

			conn, client := createWorkerClient(workerAddr)
			defer conn.Close()

			_, err := client.StartShuffle(context.Background(), req)
			if err != nil {
				log.Fatalf("Errore durante la fase di reduce: %v", err)
			}
		}(workerAddr, partition)
	}
	wg.Wait()

	// Fase di Reduce
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		workerAddr := m.workers[i]
		go func(workerAddr string) {
			defer wg.Done()

			conn, client := createWorkerClient(workerAddr)
			defer conn.Close()

			worker := &pb.Worker{
				Worker: m.workers[i],
			}

			_, err := client.Reduce(context.Background(), worker)
			if err != nil {
				log.Fatalf("Errore durante la fase di reduce: %v", err)
			}
		}(workerAddr)
	}
	wg.Wait()
}

// Funzione di supporto per creare il client gRPC del worker
func createWorkerClient(workerAddr string) (*grpc.ClientConn, pb.WorkerServiceClient) {
	conn, err := grpc.Dial(workerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Impossibile connettersi al worker %s: %v", workerAddr, err)
	}
	client := pb.NewWorkerServiceClient(conn)
	return conn, client
}

type Conf struct {
	Master  string   `json:"master"`
	Workers []string `json:"workers"`
}

// Main per avviare il Master
func main() {
	filename := os.Getenv("CONFIG_FILE")
	configData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Errore nell'apertura del file:", err)
		return
	}
	var conf Conf
	if err := json.Unmarshal(configData, &conf); err != nil {
		fmt.Println("Errore nella decodifica del JSON:", err)
		return
	}

	listener, err := net.Listen("tcp", conf.Master)
	if err != nil {
		log.Fatalf("Errore nel creare il listener: %v", err)
	}

	grpcServer := grpc.NewServer()
	master := &MasterServer{
		workers: conf.Workers,
	}

	pb.RegisterMasterServiceServer(grpcServer, master)
	log.Println("Master in ascolto su ", conf.Master)

	// Avvio del Master in una goroutine
	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Errore nel servire il Master: %v", err)
		}
	}()
	n := 1000
	var data []int32
	for range n {
		data = append(data, rand.Int31()%1000)
	}
	log.Printf("Master: Dati di input: %v", data)

	// Avvio del processo MapReduce
	master.Run(data)
}
