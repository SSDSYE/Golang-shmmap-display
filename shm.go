package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"sync"
	"time"
	"unsafe"

	shm "github.com/tmthrgd/go-shm"
)

type cloudshmmap struct {
	tnext [20480][92]float32
}

var (
	addr = flag.String("addr", ":8088", "bind address")
)

var notify struct {
	mu sync.Mutex
	ch chan struct{}

	tnext [3000][92]float32 //20480

}

func main() {
	flag.Parse()
	log.Printf("golang shmmap display, listening on http://%s", *addr)

	// Initialize the notify channel
	notify.ch = make(chan struct{})

	// Open shmmap file name
	cloudshm, e := shm.Open("/cloudshm1234", os.O_RDWR, 0660)
	defer cloudshm.Close()
	if e != nil {
		fmt.Println(e)
		os.Exit(-1)
	}

	go func() {
		for {
			time.Sleep(900 * time.Millisecond)

			tnext, _ := ReadCloudTnext(cloudshm)

			notify.mu.Lock()
			for i := 0; i < 3000; i++ {
				for j := 0; j < 92; j++ {
					notify.tnext[i][j] = tnext[i][j]
				}
			}

			notify.mu.Unlock()

			close(notify.ch)
			notify.ch = make(chan struct{})

		}
	}()

	err := http.ListenAndServe(*addr, &handler{})
	if err != nil {
		log.Fatal(err)
	}
}

func ReadCloudTnext(file *os.File) ([20480][92]float32, int) {
	var tnext cloudshmmap

	buf := make([]byte, unsafe.Sizeof(tnext))
	buf, len1 := ReadBytes(file, 0, buf)

	var pct *cloudshmmap = *(**cloudshmmap)(unsafe.Pointer(&buf))

	return pct.tnext, len1
}

func ReadBytes(file *os.File, offset int64, buf []byte) ([]byte, int) {
	//var err error
	len1, _ := file.ReadAt(buf, offset)
	if len1 == 0 {
		fmt.Println("read bytes error")
		//err = errors.New("read 0 bytes")
	}

	return buf, len1
}
func Float32ToByte(float float32) []byte {
	bits := math.Float32bits(float)
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)

	return bytes
}

func ByteToFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)

	return math.Float32frombits(bits)
}
