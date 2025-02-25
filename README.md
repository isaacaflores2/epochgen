# epochgen
Simple CLI to work with epoch times

# Build and Run
```shell
go build
./epochgen
```

# Copy to bin dir
```shell
cp ./epochgen ~/bin/epochgen 
```

## Sample output
```
All epoch times are milliseconds ⏱ 
 - current epoch: 1660055225761 
 - epoch 48 hours from now: 1660228025761 
 - epoch 14 days from now: 1661264825761 
```

# Parse time
```shell
./epochgen -t 2022-08-17T23:00:00.000Z
````

## Parsed time output
```
Successfully Parsed time: 2022-08-17T23:00:00.000Z  ⏱ 
 - epoch in milliseconds: 1660777200000 
```

# Add duration to epoch time
```shell
./epochgen -e 1660055225761 -d 48h
```

## New epoch time output
```
Successfully added 48h0m0s to epoch time: 1660055225761 ⏱ 
 - New time: 1660228025761 
```

