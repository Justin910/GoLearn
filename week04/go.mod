module week04

go 1.17

require (
	github.com/golang/protobuf v1.5.2
	google.golang.org/grpc v1.46.2
)

require (
	github.com/google/go-cmp v0.5.7 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/text v0.3.5 // indirect
	google.golang.org/genproto v0.0.0-20220519153652-3a47de7e79bd // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

replace (
	week04/app/ => ./app/
)