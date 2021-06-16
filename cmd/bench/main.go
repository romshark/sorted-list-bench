package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/romshark/sorted-list-bench/list"
	"github.com/romshark/sorted-list-bench/rnd"
)

func main() {
	flagImpl := flag.String("impl", "slice", "implementation")
	flagMinVal := flag.Int("min-val", 1, "minimum value")
	flagMaxVal := flag.Int("max-val", 1_000_000, "maximum value")
	flagSize := flag.Int("size", 1_000, "list size")
	flagScan := flag.Bool(
		"scan",
		true,
		"scans all entries after they're pushed",
	)
	flagDelete := flag.Bool(
		"delete",
		true,
		"deletes all entries after they're pushed",
	)
	flag.Parse()

	impl, err := getImpl(*flagImpl)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("conf.implementation = ", *flagImpl)
	fmt.Println("conf.min-value =", prettyUI64(uint64(*flagMinVal)))
	fmt.Println("conf.max-value =", prettyUI64(uint64(*flagMaxVal)))
	fmt.Println("conf.list-size =", prettyUI64(uint64(*flagSize)))
	fmt.Println("conf.scan =", *flagScan)
	fmt.Println("conf.delete =", *flagDelete)

	// Prepare
	log.Print("generating random input")
	start := time.Now()
	ints := rnd.MakeInts(*flagSize, *flagMinVal, *flagMaxVal)
	log.Printf("random input generated (%s)", time.Since(start))

	l := impl.Make()
	log.Printf("starting benchmark: %T", l)

	// Run benchmark
	start = time.Now()
	for _, in := range ints {
		l.Push(in)
	}
	if *flagScan {
		l.Scan(nil, func(y interface{}) bool {
			ScanVal = y.(int)
			return true
		})
	}
	if *flagDelete {
		for _, in := range ints {
			l.Delete(in)
		}
	}
	log.Printf("finished (%s)", time.Since(start))

	// Print runtime statistics
	var stat runtime.MemStats
	runtime.ReadMemStats(&stat)
	log.Print("total-alloc:", prettyUI64(stat.TotalAlloc))
	log.Print("heap-objects-freed:", prettyUI64(stat.Frees))
	log.Print("gc-cycles-completed:", prettyUI64(uint64(stat.NumGC)))
	log.Print("stw-ns-total:", time.Duration(stat.PauseTotalNs))
}

var ScanVal int

func getImpl(name string) (list.Implementation, error) {
	for _, i := range list.Implementations(sortFunc) {
		if i.Name == name {
			return i, nil
		}
	}
	return list.Implementation{}, fmt.Errorf(
		"no implementation for %s", name,
	)
}

func sortFunc(i, j interface{}) bool {
	return i.(int) < j.(int)
}

func prettyUI64(i uint64) string {
	return message.NewPrinter(language.English).Sprintf("%d", i)
}
