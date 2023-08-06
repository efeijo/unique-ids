FILE_BIN=maelstrom-unique_ids
EXERCISE_FOLDER=echoEx
MAELSTROM_PATH=~/coding/maelstrom/maelstrom
BIN_PATH=~/coding/EXERCISE/bin/${FILE_BIN}
run: build
	${MAELSTROM_PATH} test -w echo --bin ${BIN_PATH} --node-count 1 --time-limit 10

build:
	go install .
	go build -o ${BIN_PATH} main.go

clean:
	rm -rf bin 
	rm -rf store