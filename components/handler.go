package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.5" )

func Gl19XXLNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mix8rKuCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDKpLoD1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PqOypT2BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hrguOr7NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 85HdhNbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nMsZr4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1z2odT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4aL3M2I7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDcf7vyfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZxQv6ACZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWhaoy7NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJAXHAQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QD333dcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rUVK2jAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ogMoL0bKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AhmxkVKnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNY5rbb0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iuOs1CsnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1ktheHgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXC0oaElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c6Tf3lMNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CqL9HQzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JyfoEwhSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rv36JoaAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CgSubY5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FBy5b9tCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8FJNeQ4SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JmyL1HngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQ9ZSEBBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TzFOo4EwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f7HulotIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func reLWRzCuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EHfzDoaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wR5DRRcpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hev61z81Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BVzESIjtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwaDflRoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DNwjbM7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mhfPgFqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HdBapXTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tq8jLcKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l6344JrDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3X7gz5LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EZyXxQJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4JWoEkKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TtdhSv8HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NAwyXpQNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DaIlYXjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1QOAbT2PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dncGrMM4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hxU4PSftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UoI9Vs2TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kHNYD933Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func soWaGfc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9f0sqJjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ARVPz6kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bbEc7RFAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func npAOFSVIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aeHWta8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kt1K5zyTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S1kjtAoQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rJ8qNzZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0R2uf6uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cushe77VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcXOEi2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7KIxYN9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UG5sPFKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2l2TwUIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AzsXcMVrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yJl8OKgTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nrFTzd1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FwQPkLJTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F8GbMOFQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B6di78fZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iRmiw0KtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPZDsDHCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bu7BIoTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gOLLYEvBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IsueFe8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zAdJDfcTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBE3nk8uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func flRmejs2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRngcWVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xSxtmGDvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hago9tLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4U5NpFcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4zm4CoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ny2inPnFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MzuE1oCqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8c3gGPSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NrIljln0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func syX5i0yTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y71ORUBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UWFNTXApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Puw4hpJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LHFA09lqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SKWDvdj8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 74Z3DcDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rj8D6IFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 604ea3LtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RfOGYI7ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3jXWeF76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pTNIckmDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kh0c9KAxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KDgY7OoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IalltaThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sHDNmGL9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func toMJuz5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2sOT2H5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OoApbJfCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3HAwlotsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cty3BFj6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9pg6BJdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qK4yPQpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkX2sxgtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qaFpjek4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QLWmJ5m5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YWZK3v5xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d8U6jIknWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hOoKpg7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q0H2TJi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26MFUuxkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwzqGsp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iFmrYsUvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SB7emeCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UzgffDYGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfaNc2RRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EGp10u9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4fYmZxigWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5EXChdEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RBGl94DAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQQ7CNxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmWzkso1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjIZ62rhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UZ8EBYVNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p7M7oAvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35xuJ5ntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V22R6QqXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HHmKLK5gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func natAqvXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func udAQWIdFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1bHb5kwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MgIhp5JVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JHbVwmrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8nQM8o1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J03YCaQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NNW5EPi4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h8I6tTuQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AMgdIx41Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e21CK8oGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JwBsVNSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFO9pHGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R5Nf7qYTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XOiwY7Z8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nzc6FlZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L82hopaeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wntkmKJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwsrojhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2iZy9SO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GpyCylvNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JiWCRsCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fgt5t95LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i6KGtG4uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vw8BBttHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4EvkK9rSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T3eYALHeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func apc7abJWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fYo0K031Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w8ifxyEDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4tMziVyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9VUdcs7sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wByZkEFmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c560KyOkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWkvyCKdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R0BZcrpeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X0QQbIxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nv9KAsl0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C38CqqdzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ml12ttyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Q7sOqPKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51uFfy0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YXXit9MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yPuV7aVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MikDJs6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 03nZzijBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cS9Hk5saWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1vobYNZCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBcPCp0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mwTwTG5SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YblHWTYBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7rL4U0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjYHKUwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NBlCqcD7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJvIMH3FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CNa5KwmzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBhrKKsfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sryR1u0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaXmVhDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uV7TZZVhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ATFgl8tGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LtM2iPBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybjVoB8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uIW92uftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qjBO0CpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xWN7KcXEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xl8XltxqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LiALNDq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gln8Cw1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtPka8FdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PtwxHcTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJB330ZrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C71CMlo5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vv5GACt2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZINKfP2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nTkjipe7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGxLxI7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2GwnDca5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YpMU8c1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NqgsijX8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Fofdri1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qPCevRPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func geZR2hyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P9GsfdpvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zxiG01iKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8avgl6BLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ESWkUeIRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FqFt2c1AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VsA4xwAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6q14zGTOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k4whRVN0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFFXMObJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QHJAJhvKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WMaypLkaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vpHfTD0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5wj6BbrCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVPnfubyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C9omMMdLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dqFniSOiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W21S1XkyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2lZcnD0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsqbTncbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTGplGlPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GjPTmUaFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7miLtoZJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cOQMEJScWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4pJPIIQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SD2PTsYhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func utHSHrKIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywumFEYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9nGomUJRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swpNFIflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5HcCZuIaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6PpWSDrtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JPS0IZxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yf1MF7kfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjUZ5cK6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AVuPEqacWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vGNkGEFaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C4oj3daEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JRUwbVuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pbWGudbtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MRLtWBd5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q8HV3O6xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LeLGaE3uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PK1PAq5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kBoZNHHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 98LZx8FvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yi0HIU1YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EzYygju5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyA42K47Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func thADeMbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0i17HNMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6RPzksHiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ndzulNYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 25AZXqiUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xDTyWdX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z8k2PryhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qYCY1cWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZkP5rcRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gQSWavK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9yjSADMIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxrPLbN1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xm3R8Q6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hj4I03huWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 829rQDQpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5EIuWrzgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xbqRcNNoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ednpf4gIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSKOoidJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jn9Yzi7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QADrChR6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MfElXYjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 107rtm8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DoMceyIyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NxwzFHFDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pmtp7ThmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 63xdoOzCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tJLzsDaZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IE6P5L7uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgVf2ZAtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yOyXKC3JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func spu3C4vxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0VgUAfudWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ou3tvoXeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jAseLa1OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G9M5iIQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XMkVsafxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lmCO0pT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdLMDu9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H02P5sTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YRmD4saCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fv9gLp38Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9bmleS3qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78HSFQOrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BbGxWIQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4gyyHcL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UTyudMaQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eZw2bYCiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zH9un8K9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func up5PYTpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4hn2tHflWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xeqqmbqLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mej2WwSGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5v6p0iUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ryaQEmOyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AdzLocZdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KNQj3K3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8ogvsTSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qoY7JJUqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOP0NkDpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VP6pDPfpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7V2rISZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6iSEk67dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wGdayPK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7diFtzB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TvzftBayWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFnxKTpZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXdphcuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODhwvrvuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJ9OL5hTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vmRIAOEXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 52kWtnhiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4H0N6vejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jA0ofPwqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NzZ4oGt1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3fS747rIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NVwaVeKXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TGMHQ6KuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ww2j07bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4SNakqT2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fwY5iNfDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GKHoXy6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q3BZOckYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func heM0xKgNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8OyDye1MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HUW3DeQ3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sg2cofkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tNlB6MTWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EsUWxFDWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9KhGINmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PEYo5rvkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YHDUUGMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vJkndDLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GZhNML1SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4H5AzDsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cw9fI90mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2hTkfqRnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AiM3BOPpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43CyX6DIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJYL151PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P47O3e7hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1j9IYZ0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k1qrkx9HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0IrpL9XKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mu27O91eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tXSnaNnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hao1L25TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IiZmquqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7NGtLeJAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVwf7PwhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lGFtPpNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wJxcXYNgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VxK7sGwlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dblnMfvAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSqVelyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aMQLEYWIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PoEAeMUNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NGfVAnAvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 805aGJKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ICUJydqbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MJxwsCVGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lcq7aWElWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RYz5lB4RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 227cGFK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gwKI0RNzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGAy005JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g22jN8FVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knj9lrEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OWlNNDJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ih0oZKCFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t1PEbTyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBJg7FxCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yr2qb6k4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rEPkGuUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WCC9lpE9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ydctNE2QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tLtPAEqdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vpl37F3GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func at85BHynWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func O5wQdOdUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KOz4AQZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTetUmMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5n8Ae2EhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7p5pJPr6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7K35EhDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func si6PKNQiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h23xUbmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H3MIsuJrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZTqE23MGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KkA6FwELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vDfxBKO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WgEg0DIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3q3CGfeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jwiKcwwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EVkupbeuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4XENEKeCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sNCDCtlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wsQ54ajWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KxkjkM8ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dtv0iHyEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PCnqVHsFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fpMBxdUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Uyg1HTNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pzTUu9n7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vxVNs1r6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5hcyf3a1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f8uy8GyyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rkIncC2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rNih9I7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2SW347d6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func reFsAlCNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mbSOmwL7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oX3A6EtVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fpHhxhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Aa9afYJEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0wNY93zRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RxJJ31zeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z0UTr68JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zCsrWCXYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vg7lYSWxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u3oopMDcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxPCeC6pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htYCJlFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LyjyUgZQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d6BFDhqCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1isWYPc8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Xu7FENZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r0OB8s0lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VdXwZc5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6dCaazeaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFiKEAyqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b4ETxJXzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6BVTGZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mt6ud6oIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mWVIhUJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5AQ0BcZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5GywISCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func un1xd39iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZzt3iXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func myyDCOpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aczAp5J0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bq8pbQjGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9hPcebNeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func etmeZQL0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LuF4Hno2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OBeDEurFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbwQcEtNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T1rtvP4DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kaDXBOhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jOmmDgDZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CLQDNY6gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wkpsDNoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h1nKJJxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mC30EZ8AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YVwnB7YjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0lf2XJlEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9SpD8v2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nSz4YpV4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5crkTBMzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5R2VBhZNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cNHNqEaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jvCDQ5gxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6HFi85FdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3uKPXD4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lqc4A2idWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bPMIH8JLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QNTEXkDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1pNx7YRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D54xtNh9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cygKapZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w4Ldx1mMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HNjc9TS2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qPyOootjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i2o5c7csWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dgySj0wUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N94nBahRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MldC7USoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZGtpCKcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IiBeWDpGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QZOOELnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qIALdcxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q9GlqhwkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GtCH8jqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SfmSu6JXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WW5gVSVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rXUelKdVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y72O0fvRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JfBz1vftWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGSaZrdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yFbAnU1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8XveLyJdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4qKzM6dNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGQd1ibEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cTEP64L0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4dzoeiSyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bvLf2JUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BIOu1dNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gg1p8jv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2FfTdj4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TO56xpBPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LMnw0ElIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6mIJDoTIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DZPNI2yMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qs8Y7QNBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mnr1yWg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLbsqlYkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8G4O0D7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A4znz3JUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3QcAqriRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgaJPJitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzlgeEJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 47BX1wvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LzZ8XgMiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JROzxgUmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mn4XhjoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Frvkaw45Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I3dELEopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhdgrFr1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFb22ioLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2MS9EPpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tf6SBZamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8fA8kNOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LUQGNsxNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uLfChYw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YD9zjTAwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jpF97zSsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZPBIQHPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J7UqWAHKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QK3lpCsZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4rWBjmmxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xzUFO8O6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i5GU8jmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EhDcrThjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnwRZPrsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eeV9AyUcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W1J15KpaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RBjGSrwaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9wRncJuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aJRrGFnUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWZBPL7dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rap6krGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZxPsTs9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ffBYazLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YoQE8JCTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPjjMuaOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFAUp4MoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5lBgojGgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N9ov4ddbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6doqRwLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVgNKvxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXXXwm2fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hXKHrEIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CPGfsE0KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gwukkcveWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E2Pb58SJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2U3eYa3MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QnmFDG2nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVEYP6DzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nmuVYEhjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BGxGpRo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jja6mEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HSfaBVFlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8I2IeUHQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUFPYAuLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 30ulm3MMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W3iPgZUGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyRpn1wyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ky8HlAqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InkUvYnEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hCMIGJcjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C51gi79ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func blAwZKsHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KqtJqxwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NpnxDji5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZANgwIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rjC25SSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8ZB85GsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WhKQhQvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b2Smb48LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eTbu9yv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywCdaA9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4jkldZ54Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XAxStBu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E0M2ZBoKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mkaBbiI2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DSXpFbcWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JthwAXFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pgGiRJCYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func anfIBovYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YtmLpWeiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RKEnQtKMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHWbuj7eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 86hb3zNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6B3kaMp0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x2ospwMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k21CJU2tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oMkvgtVkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lxbIXkT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func htjBeQMKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sh3Hv9nDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rnV5AFlnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e9g1VatoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lcBvvPLCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8I9P9F5ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DoCGRmsEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tc5THfiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jVCuTvmnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ea2APNEIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QiJ8rgDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eJhF6fSMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6xO1NfheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IiTtIsx6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BbmmTjyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dGNwcVgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZhI3ZehjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8uWQ1SuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CAYCf1IAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QOCOdy19Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fzUzEIRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nVD1oh12Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pyk34ZkmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x02kqgamWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bL2dh8LIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bO6YuIHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func efxBbCeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fUn4FwX9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBDAIwBuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nEdygeyaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HGkhOYrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VhD3bRVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cRzWxznWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hO4ODwMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3k1zgftVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWIDJgWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VSicYF3bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a3MZCHbHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yKMTq9eyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hL1cYivuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZzoCYFRXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iOUW2yysWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IuW1Xuj1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 51
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func toL86XQFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hWJkKd1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qUY9FquqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kismjUWKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func drKaaO07Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3OCzWIUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ax9MLfskWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWHpqwldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FIKSPsPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VvDyohmqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FB33CCz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i4lK0GTKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func brDG3xEBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CncLtM5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W9L9KdW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKVKLaKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Mc88u0PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7kNIq1PuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8mNOaR5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMzwOerqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m3xmn6szWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZMsw5UcGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mdjx0GO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r5qqLZxXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DMJJRxwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N94AfcgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DdaPmdnjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D58yc1DEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yNsbuEFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0SvHWrVuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oWO5TIvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J8V2AW9EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UcnMenHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LB8aJg4MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mSS64qOIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vWE37YgjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezGBMeqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s8c1xlhcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SxKiHESyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ciFUnPnLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tkg6eiGlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eH2QA5NzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nUsGTp0rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MTphy9R9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3KZnqTk7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wpyjL1uwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wiAzpcvmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2sknVyGiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnxC9PqVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nYYnS5iIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g38JlAbyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YSX2LnzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G1CekRmXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YZe2PAuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dd7ObIPwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rkn2nwEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2nOmsDEEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlRtjdnOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKHmYeTQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FWSwrIiCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tYlDh9BEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 80FrTHerWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOOBBEbzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eLK50AQWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5t3HGUXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aBopnG0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qHKBfyP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZjoZrpRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0KG610wUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywrWKqpwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gzqm0lL3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GHQUno2IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NSp3v35SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6HN9aNFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJPEEt2ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OysuikheWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JKJ30edqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFsndvJdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZLQUHI9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pwE7GUUpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Xy7bpe1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zViLU9t6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rt8E6h0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VsiCrZXxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KYUbHioaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kT05BPF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EeoJhoZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9HqVCRVPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OTI0YILaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gNdoUl06Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BhUk74HRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WSya4XlnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lBu0YZzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rvZXzUNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

