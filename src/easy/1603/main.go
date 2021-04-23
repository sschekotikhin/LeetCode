package main

type ParkingSystem struct {
	kinds map[int]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
    return ParkingSystem{kinds: map[int]int{1: big, 2: medium, 3: small}}
}


func (this *ParkingSystem) AddCar(carType int) bool {
    this.kinds[carType]--

	return this.kinds[carType] >= 0
}
