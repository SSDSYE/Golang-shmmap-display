# Golang Shmmap Display

A demo of real-time displaying shared memory content with P5.js, Go, and server-side events.

## Overview

This Golang-based program is used to read data from a shared memory and show it visually on a browser on time. As for the case from which this program is devloped, a temperature profile that describes the 2D temperature distribution in a steel slab in a continuous casting process is stored in the shared memory, as a large matrix with the dimension of 3001x91. The program is designed to read the temperature data from the shared memory and show it as a real-time updated graph to demonstrate the temperature distribution on the steel slab with different temperature represented by different colors. The real-time updated graph can be considered as a simple temperature point cloud. 

## Demo

```shell
git clone https://github.com/SSDSYE/golang-shmmap-display.git
cd golang-shmmap-display
go run *.go
```

Then browse to http://localhost:8088

💥

![Screenshot](screenshot.png)
