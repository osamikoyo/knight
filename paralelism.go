package main

import "sync"

func (m *Manifest) Paralelism(tasks ...string) {
	var wg sync.WaitGroup

	for _, t := range tasks {
		wg.Add(1)

		go func(w *sync.WaitGroup) {
			m.Run(t)

			w.Done()
		}(&wg)
	}

	wg.Wait()
}
