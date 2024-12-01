# Advent of Code

Learning Go properly with this, because Google doesn't teach me shit.


## 2024

### Day 1 ✅

It was a super straightforward solution, for part 1, we would just need to sort the two arrays then iterate over the values, taking the differences from each list. I got choked up with converting strings to ints, because for some reason the strconv.Atoi wasn't importing properly the first few times trying to read the input

Part 2 was also straightforward, but I did find key checking in Go weird, since it doesn't look like anything I am used to.

```golang
if _, ok := seenInList2[v]; ok {
  ...
}
```

The underscore is something I'm used to to, but apparently, it's even more common in Go, which is interesting.

Overall, it's been a long time since I've actually used Go for these types of coding challenges, but overall, it took me 18 minutes 13 seconds to complete part 1, and 7 minutes 23 seconds to complete part 2. It took some time to look up syntax again, and how scanners in Go work, since this is the first time taking file input in Go, but overall, it was pretty fun.
