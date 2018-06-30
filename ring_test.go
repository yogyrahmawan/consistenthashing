package consistenthashing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type server struct {
	name string
	ip   string
}

func (s server) String() string {
	return fmt.Sprintf("name/key = %s, ip = %s", s.name, s.ip)
}

func TestRing(t *testing.T) {

	var servers = []server{
		{
			name: "server1",
			ip:   "10.13.13.0:8000",
		},
		{
			name: "server2",
			ip:   "10.13.13.0:8001",
		},
		{
			name: "server3",
			ip:   "10.13.13.0:8002",
		},
		{
			name: "server4",
			ip:   "10.13.13.0:8003",
		},
		{
			name: "server5",
			ip:   "10.13.13.0:8004",
		},
		{
			name: "server6",
			ip:   "10.13.13.0:8005",
		},
		{
			name: "server7",
			ip:   "10.13.13.0:8006",
		},
		{
			name: "server8",
			ip:   "10.13.13.0:8007",
		},
		{
			name: "server9",
			ip:   "10.13.13.0:8008",
		},
		{
			name: "server10",
			ip:   "10.13.13.0:8009",
		},
	}

	// init ring
	r := Init()
	for _, s := range servers {
		r.Insert(s.name, s)
	}

	// init key
	ks := []string{"abc", "a", "ter", "2232", "5454", "44343", "5454", "4354", "34343", "4re"}

	// test concurrent get
	resultChan := make(chan interface{}, len(ks))
	errChan := make(chan error)
	for i := 0; i < len(ks); i++ {
		go func(s string) {
			n, err := r.Get(s)
			if err != nil {
				errChan <- err
				return
			}
			resultChan <- n

		}(ks[i])
	}

	assert := assert.New(t)
	for i := 0; i < len(ks); i++ {
		select {
		case res := <-resultChan:
			assert.NotNil(res)
			_, ok := res.(server)

			assert.Equal(true, ok)
		case err := <-errChan:
			fmt.Println(err)
		}
	}

	// removing one node

	deletedServer := "server1"
	r.Remove(deletedServer)

	ns := r.GetAllElmt()

	deleted := true
	for _, v := range ns {
		val := v.(server)
		if val.name == deletedServer {
			deleted = false
			break
		}
	}

	assert.Equal(true, deleted)
}
