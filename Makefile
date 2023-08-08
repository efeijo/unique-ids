FILE_BIN=maelstrom-unique_ids
EXERCISE_FOLDER=unique-ids
MAELSTROM_PATH=~/coding/maelstrom/maelstrom
BIN_PATH=~/coding/EXERCISE/bin/${FILE_BIN}
run: build
	${MAELSTROM_PATH} test -w unique-ids --bin ${BIN_PATH} --time-limit 30 --rate 1000 --node-count 3 --availability total --nemesis partition

build:
	go install .
	go build -o ${BIN_PATH} main.go

clean:
	rm -rf bin 
	rm -rf store

debug:
	${MAELSTROM_PATH} serve