# README


#### AWS test

```{shell}
./chord.exe -a 172.31.35.162 -p 8100 --ts 3000 --tff 1000 --tcp 3000 -r 4
```

```{shell}
./chord.exe -a 172.31.36.232 -p 8102 --ja 172.31.35.162 --jp 8100 --ts 3000 --tff 1000 --tcp 3000 -r 4
```

```{shell}
./chord.exe -a 172.31.33.223 -p 8104 --ja 172.31.36.232 --jp 8102 --ts 3000 --tff 1000 --tcp 3000 -r 4
```