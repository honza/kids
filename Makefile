all: guess days-until-christmas days-until-birthday

guess: cmd/guess/guess.go
	go build ./cmd/guess

days-until-christmas:
	go build ./cmd/days-until-christmas

days-until-birthday:
	go build ./cmd/days-until-birthday

clean:
	rm guess days-until-christmas days-until-birthday

raspberry:
	gox -osarch="linux/arm" ./cmd/guess
	gox -osarch="linux/arm" ./cmd/days-until-birthday
	gox -osarch="linux/arm" ./cmd/days-until-christmas
