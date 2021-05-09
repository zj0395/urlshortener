# UrlShorten
shorten a url to 6-size char combine from [0-9a-zA-Z].  

## Quick Start
### Basic
1. Change every `config/*.tpl` file and delete `.tpl` from file name
2. Just `go run main.go`

### Custom
#### Generate a random algorithm for yourself
1. You can generate a unique random sequence, see `utils/shorten/defines.go`

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
