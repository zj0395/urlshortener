# UrlShorten
shorten a url to 6-size char combine from [0-9a-zA-Z].  

## Quick Start
### Basic
1. Change every template file and delete `.template` from its name
2. Just `go run main.go`

### Custom
#### Generate a random algorithm for yourself
1. You can generate a random sequence by `go run meta/functions.go`, repleace `utils/shorten/defines.go` to get your unique sequence.

## Randomness test
When number incr by one, the generate code will change a lot. Your can see the result from `bash unit_test.sh`.  
```
        	changeCharNum	cnt
        	1	0
        	2	0
        	3	0
        	4	0
        	5	977621
        	6	22379
```
