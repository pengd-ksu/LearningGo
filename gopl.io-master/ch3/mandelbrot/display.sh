#!/usr/bin/env bash

go run main.go > out/mandelbrot.png
open -a Preview out/mandelbrot.png