package main

import {
	"fmt"
	"math/rand"
	"time"
)

func main(){
	rand.Seed(time.now) 

	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(100))
		