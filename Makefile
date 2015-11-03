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
	gox -osarch="linux/arm" ./cmd/guess
	gox -osarch="linux/arm" ./cmd/days-until-birthday
	gox -osarch="linux/arm" ./cmd/days-until-christmas
	gox -osarch="linux/arm" ./cmd/math

lint:
	golint cmd/guess/guess.go
	golint cmd/days-until-christmas/days-until-christmas.go
	golint cmd/days-until-birthday/days-until-birthday.go
	golint cmd/math/math.go
	golint lib/lib.go
