package parking

import (
    "container/heap"
    "errors"
)

// ParkingLot merepresentasikan parkiran
type ParkingLot struct {
    capacity int
    slots    map[int]*Car    // slot terisi -> mobil
    free     *MinHeap        // heap untuk cari slot terdekat
}

// NewParkingLot inisialisasi parkiran dengan kapasitas tertentu
func NewParkingLot(capacity int) *ParkingLot {
    h := &MinHeap{}
    for i := 1; i <= capacity; i++ {
        heap.Push(h, i) // isi heap dengan semua slot kosong
    }
    return &ParkingLot{
        capacity: capacity,
        slots:    make(map[int]*Car),
        free:     h,
    }
}

// Park mobil ke slot terdekat
func (p *ParkingLot) Park(car *Car) (int, error) {
    if p.free.Len() == 0 {
        return 0, errors.New("Sorry, parking lot is full")
    }
    slot := heap.Pop(p.free).(int)
    p.slots[slot] = car
    return slot, nil
}

// Leave mobil keluar + hitung biaya
func (p *ParkingLot) Leave(carNumber string, hours int) (int, int, error) {
    slot := -1
    for s, c := range p.slots {
        if c.Number == carNumber {
            slot = s
            delete(p.slots, s)
            heap.Push(p.free, s)
            break
        }
    }
    if slot == -1 {
        return 0, 0, errors.New("Registration number not found")
    }

    // Hitung biaya
    charge := 10
    if hours > 2 {
        charge += (hours - 2) * 10
    }
    return slot, charge, nil
}

// Status menampilkan slot yang terisi
func (p *ParkingLot) Status() map[int]string {
    result := make(map[int]string)
    for s, c := range p.slots {
        result[s] = c.Number
    }
    return result
}
func (p *ParkingLot) Capacity() int {
    return p.capacity
}
