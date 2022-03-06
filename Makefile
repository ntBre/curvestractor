test:
	go run . input.png > output.dat
	gnuplot -p -e "plot 'output.dat'"

install:
	go install .
