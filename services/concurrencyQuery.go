package services

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/JoseLuisTorrentera/academy-go-q42021/models"
)

type GoroutinePool struct {
	queue   chan []string
	results chan models.Spell
	errors  chan error
	wg      sync.WaitGroup
}

func (gp *GoroutinePool) Close() {
	close(gp.queue)
	gp.wg.Wait()
	close(gp.results)
	close(gp.errors)

}

func (gp *GoroutinePool) AddWorkers(numWorkers int) {
	gp.wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func(workerID int) {

			count := 0
			for job := range gp.queue {
				id, err := strconv.Atoi(job[0])
				if err != nil {
					gp.errors <- err
				}

				level, err := strconv.Atoi(job[3])
				if err != nil {
					gp.errors <- err
				}

				spell := models.Spell{
					ID:      id,
					Name:    job[1],
					Classes: job[2],
					Level:   level,
					School:  job[4],
				}
				count++
				gp.results <- spell

			}
			fmt.Println(fmt.Sprintf("Worker %d executed %d jobs", workerID, count))
			gp.wg.Done()
		}(i)
	}
}

type QueryService struct {
	file  string
	gpool *GoroutinePool
}

func NewQueryService(file string) QueryService {
	return QueryService{
		file: file,
	}
}

func (qs QueryService) GetSpellsByQuery(itemType int, numItems int, numItemsWorker int) ([]*models.Spell, error) {

	csvFile, err := os.Open(qs.file)
	defer csvFile.Close()
	if err != nil {
		return nil, err
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return nil, err
	}
	qs.gpool = &GoroutinePool{
		queue:   make(chan []string, numItems),
		results: make(chan models.Spell, numItems),
		errors:  make(chan error, 1),
	}

	numWorkers := numItems / numItemsWorker
	qs.gpool.AddWorkers(numWorkers)

	items := 0
	for _, line := range csvLines {
		if items == numItems {
			break
		}
		id, err := strconv.Atoi(line[0])
		if err != nil {
			return nil, err
		}
		if id%2 == itemType {
			qs.gpool.queue <- line
			items++
		}
	}

	qs.gpool.Close()

	spells := []*models.Spell{}

	err = <-qs.gpool.errors
	if err != nil {
		return nil, err
	}

	for i := 0; i < numItems; i++ {
		spell := <-qs.gpool.results
		spells = append(spells, &spell)
	}

	return spells, nil
}
