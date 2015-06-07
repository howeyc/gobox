#wc
a faster wc (mostly)

- Same speed as native wc for line/byte counts.
- **3x** faster word counts
- **30x** faster multibyte runes counts

Relies on performance gained from multithreaded processing on buffer chunks.

```shell
./wc [-l] [-m] [-w] [-b] file [...fileN]
  -c=false: count bytes
  -l=false: count lines
  -m=false: count multibyte runes
  -w=false: count words

# count lines
time ./wc -l 100m.lines;
 100000000 100m.lines

real	0m1.155s
user	0m3.875s
sys	    0m0.389s

time wc -l 100m.lines 
 100000000 100m.lines

real	0m1.138s
user	0m0.920s
sys	    0m0.209s

# count words
time ./wc -w 100m.lines;
 108339058 100m.lines

real	0m2.119s
user	0m7.289s
sys	    0m0.438s

time wc -w 100m.lines 
 108339058 100m.lines

real	0m5.888s
user	0m5.650s
sys	    0m0.237s

# count multibyte characters
time ./wc -m 100m.lines;
 1093761047 100m.lines

real	0m0.677s
user	0m1.981s
sys	    0m0.348s

time wc -m 100m.lines 
 1093761047 100m.lines

real	0m18.468s
user	0m18.186s
sys	    0m0.268s

# count bytes
time ./wc -c 100m.lines;
 1093761047 100m.lines

real	0m0.420s
user	0m0.176s
sys	    0m0.298s

time wc -c 100m.lines 
 1093761047 100m.lines

real	0m0.003s
user	0m0.001s
sys	    0m0.002s

/usr/sbin/system_profiler -detailLevel full SPHardwareDataType
Hardware:

    Hardware Overview:

      Model Name: MacBook Air
      Model Identifier: MacBookAir6,2
      Processor Name: Intel Core i5
      Processor Speed: 1.3 GHz
      Number of Processors: 1
      Total Number of Cores: 2
      L2 Cache (per Core): 256 KB
      L3 Cache: 3 MB
      Memory: 4 GB
      ...
```

##Build and Use

```shell
git clone git@github.com:jasonmoo/wc.git
cd wc/cmd
go get
go build wc.go
```
