package vars

import "sync"

// Used to help keep track of jobs in a WaitGroup
type RunningJobs struct {
	JobCount int
	Mw       sync.RWMutex
}

func (job *RunningJobs) GetJobs() int {
	job.Mw.RLock()
	defer job.Mw.RUnlock()
	return job.JobCount
}
func (job *RunningJobs) AddJob() {
	job.Mw.Lock()
	defer job.Mw.Unlock()
	job.JobCount += 1
}
func (job *RunningJobs) SubJob() {
	job.Mw.Lock()
	defer job.Mw.Unlock()
	job.JobCount -= 1
}
