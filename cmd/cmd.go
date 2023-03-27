package cmd

import (
	"fmt"
	"github.com/raylax/rayx/dsl"
	"github.com/raylax/rayx/executor"
	"github.com/raylax/rayx/version"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var cmd = &cobra.Command{
	Use:     "rayx",
	Short:   "xray poc executor",
	Version: fmt.Sprintf("%s %s %s", version.BuildVersion, version.BuildHash, version.BuildDate),
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

var pocs string
var url string
var headers = newMapValue()
var cookies = newMapValue()
var threads int
var sleep int
var timeout int

func init() {
	flags := cmd.Flags()
	flags.StringVarP(&pocs, "poc", "p", "", "POC file/folder (eg. path/to/poc,rce.yaml)")
	flags.StringVarP(&url, "url", "u", "", "target url (eg. https://httpbin.org)")
	flags.IntVarP(&threads, "threads", "t", 1, "threads")
	flags.IntVarP(&sleep, "sleep", "s", 0, "sleep seconds")

	flags.IntVarP(&timeout, "timeout", "", 30, "timeout seconds")
	flags.Var(&headers, "header", "(eg header1=value1,header1=value2)")
	flags.Var(&cookies, "cookie", "(eg cookie1=value1,cookie2=value2)")

	cmd.MarkFlagRequired("poc")
	cmd.MarkFlagRequired("url")
}

func Execute() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getFiles(path string) []string {
	stat, err := os.Stat(path)
	if err != nil {
		return nil
	}
	if !stat.IsDir() {
		return []string{path}
	}
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil
	}

	var files []string
	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".yaml") || strings.HasSuffix(entry.Name(), ".yml") {
			files = append(files, filepath.Join(path, entry.Name()))
		}
	}

	return files
}

func execute() {

	var files []string
	for _, poc := range strings.Split(pocs, ",") {
		files = append(files, getFiles(strings.TrimSpace(poc))...)
	}
	if len(files) == 0 {
		println("no poc file")
		os.Exit(0)
	}

	fmt.Printf("POCs: %d threads:%d sleep:%d - %s\n", len(files), threads, sleep, url)

	ch := make(chan *workerTask)
	config := &executor.Config{
		Url:     url,
		Timeout: timeout,
		Headers: *headers.value,
		Cookies: *cookies.value,
	}

	var wg sync.WaitGroup
	wg.Add(len(files))

	for i := 0; i < threads; i++ {
		w := &worker{
			e:  executor.NewPocExecutor(),
			c:  config,
			wg: &wg,
			ch: ch,
		}
		go w.poll()
	}

	for _, file := range files {
		poc, err := dsl.ParsePoc(file)
		if err != nil {
			printResult(executor.StateError, file, err)
			continue
		}
		ch <- &workerTask{
			poc:  poc,
			file: file,
		}
	}

	wg.Wait()
	close(ch)

}

func printResult(state executor.State, file string, err error) {
	if err != nil {
		fmt.Printf("%s - [E] - %s\n", file, err.Error())
		return
	}
	switch state {
	case executor.StateHit:
		fmt.Printf("%s - [%s] √√√\n", file, state)
	case executor.StateMISS:
		fmt.Printf("%s - [%s]\n", file, state)
	}
}

type workerTask struct {
	file string
	poc  *dsl.Poc
}

type worker struct {
	e  *executor.PocExecutor
	c  *executor.Config
	wg *sync.WaitGroup
	ch <-chan *workerTask
}

func (w *worker) poll() {
	for {
		t := <-w.ch
		if t == nil {
			break
		}
		state, err := w.e.Execute(w.c, t.poc)
		printResult(state, t.file, err)
		w.wg.Done()
		if sleep > 0 {
			time.Sleep(time.Duration(sleep) * time.Second)
		}
	}
}
