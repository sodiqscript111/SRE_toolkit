.PHONY: run load profile-cli profile-web clean

run:
	go run app/main.go

load:
	k6 run load.js

profile-cli:
	go tool pprof http://localhost:8085/debug/pprof/profile?seconds=20

profile-web:
	go tool pprof -http=:8081 http://localhost:8085/debug/pprof/profile?seconds=20

clean:
	rm -rf *.pprof profile.pb.gz
