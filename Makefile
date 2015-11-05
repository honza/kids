all: guess days-until-christmas days-until-birthday math

guess: cmd/guess/guess.go
	go build ./cmd/guess

days-until-christmas: cmd/days-until-christmas/days-until-christmas.go
	go build ./cmd/days-until-christmas

days-until-birthday: cmd/days-until-birthday/days-until-birthday.go
	go build ./cmd/days-until-birthday

math: cmd/math/math.go
	go build ./cmd/math

clean:
	rm guess days-until-christmas days-until-birthday math

raspberry:
	GOOS=linux GOARCH=arm GOARM=6 ~/Github/go/bin/go build ./cmd/guess
	GOOS=linux GOARCH=arm GOARM=6 ~/Github/go/bin/go build ./cmd/days-until-birthday
	GOOS=linux GOARCH=arm GOARM=6 ~/Github/go/bin/go build ./cmd/days-until-christmas
	GOOS=linux GOARCH=arm GOARM=6 ~/Github/go/bin/go build ./cmd/math

lint:
	golint cmd/guess/guess.go
	golint cmd/days-until-christmas/days-until-christmas.go
	golint cmd/days-until-birthday/days-until-birthday.go
	golint cmd/math/math.go
	golint lib/lib.go
