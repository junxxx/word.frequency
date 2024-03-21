a word counter

args
-f file

usage:
./wordcount -f "path/to/file"

Actually this can be done with shell 
```
cat file | tr ' ' '\n' | tr -d ',' | tr -d '.' | sort | uniq -c | sort -r
```