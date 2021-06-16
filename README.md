<a href="https://github.com/romshark/sorted-list-bench/actions?query=workflow%3ACI">
    <img src="https://github.com/romshark/sorted-list-bench/workflows/CI/badge.svg" alt="GitHub Actions: CI">
</a>
<a href="https://coveralls.io/github/romshark/sorted-list-bench">
    <img src="https://coveralls.io/repos/github/romshark/sorted-list-bench/badge.svg" alt="Coverage Status" />
</a>
<a href="https://goreportcard.com/report/github.com/sorted-list-bench/eventlog">
    <img src="https://goreportcard.com/badge/github.com/romshark/sorted-list-bench" alt="GoReportCard">
</a>

# sorted-list-bench

Example:
```
cd cmd/bench
go build && ./bench -size 10_000 -impl <implementation>
```

## Results
### Linked List
```
conf.implementation =  linklist
conf.min-value = 1
conf.max-value = 1,000,000
conf.list-size = 100,000
conf.scan = true
conf.delete = true
2021/06/16 21:54:41 generating random input
2021/06/16 21:54:41 random input generated (3.370403ms)
2021/06/16 21:54:41 starting benchmark: *linklist.List
2021/06/16 21:55:09 finished (28.142398182s)
2021/06/16 21:55:09 total-alloc:15,469,064
2021/06/16 21:55:09 heap-objects-freed:101,142
2021/06/16 21:55:09 gc-cycles-completed:3
2021/06/16 21:55:09 stw-ns-total:89.996µs
```

### Slice
```
conf.implementation =  slice
conf.min-value = 1
conf.max-value = 1,000,000
conf.list-size = 100,000
conf.scan = true
conf.delete = true
2021/06/16 21:55:30 generating random input
2021/06/16 21:55:30 random input generated (2.061907ms)
2021/06/16 21:55:30 starting benchmark: *slice.List
2021/06/16 21:55:58 finished (27.697782571s)
2021/06/16 21:55:58 total-alloc:11,815,128
2021/06/16 21:55:58 heap-objects-freed:100,097
2021/06/16 21:55:58 gc-cycles-completed:3
2021/06/16 21:55:58 stw-ns-total:170.067µs
```
