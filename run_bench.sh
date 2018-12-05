#!/bin/bash

curr_wkdir=$PWD

# for curr_dir in $(ls ); do
for curr_dir in */; do

	echo 
	echo curr_dir $curr_dir

	cd ${curr_wkdir}/$curr_dir

	go test -bench=. -benchmem
done




