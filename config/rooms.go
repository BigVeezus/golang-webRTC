package config

import (
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/exp/rand"
)

// participant struct
type Participant struct {
	Host bool
	Conn *websocket.Conn
}

// mapping room connection to participants
type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[string][]Participant
}

// function initialize room
func (r *RoomMap) Init() {
	r.Map = make(map[string][]Participant)
}

// function to get room by Id
func (r *RoomMap) Get(roomID string) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[roomID]
}

// function to create room
func (r *RoomMap) CreateRoom() string {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	rand.Seed(uint64(time.Now().UnixNano()))
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	roomID := string(b)

	r.Map[roomID] = []Participant{}

	return roomID
}

// function to insert host into room
func (r *RoomMap) InsertIntoRoom(roomID string, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{host, conn}

	log.Println("Inserting into Room with RoomID: ", roomID)
	r.Map[roomID] = append(r.Map[roomID], p)
}

// function to delete room
func (r *RoomMap) DeleteRoom(roomID string) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, roomID)

}
