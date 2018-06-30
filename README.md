ConsistentHasing 

#### Source 
* akamai consistenthasing [https://www.akamai.com/es/es/multimedia/documents/technical-publication/consistent-hashing-and-random-trees-distributed-caching-protocols-for-relieving-hot-spots-on-the-world-wide-web-technical-publication.pdf]

#### Short Introduction 
* Hash Function : fnv1a 
* Data structure : bst 

#### usage example
```
package main
  
import (
        "fmt"
        "log"

        ch "github.com/yogyrahmawan/consistenthashing"
)

func main() {
        r := ch.Init()
        r.Insert("a", "10.88.10.200:8000")
        r.Insert("b", "10.88.10.201:8001")
        fmt.Println(r.GetAllElmt())

        // get assigned node
        assigned, err := r.Get("cde")
        if err != nil {
                log.Fatalf("error getting assigned node, err = %v \n", err)
        }
        fmt.Println(assigned)

        // removing node
        r.Remove("a")

        fmt.Println(r.GetAllElmt())
}
```