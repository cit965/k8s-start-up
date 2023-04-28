// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	list := PodList{Pods: []Pod{
		{
			"pod1",
			"tencent.io/nginx:1",
			"host1",
			"nginx",
		},
		{
			"pod2",
			"tencent.io/nginx:2",
			"host2",
			"nginx",
		},
	}}

	print := HumanReadablePrinter{}
	print.Print(list, os.Stdout)
}

type PodList struct {
	Pods []Pod
}
type Pod struct {
	name   string
	image  string
	host   string
	labels string
}

type ResourcePrinter interface {
	Print(PodList, io.Writer) error
}

type HumanReadablePrinter struct{}

func (h *HumanReadablePrinter) Print(podlist PodList, out io.Writer) error {
	h.printHeader(podColumns, out)
	for _, v := range podlist.Pods {
		fmt.Printf("%s\t%s\t%s\t%s\n", v.name, v.image, v.host, v.labels)
	}
	return nil
}

var podColumns = []string{"Name", "Image(s)", "Host", "Labels"}

func (h *HumanReadablePrinter) printHeader(columnNames []string, w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s\n", strings.Join(columnNames, "\t")); err != nil {
		return err
	}
	var lines []string
	for _, _ = range columnNames {
		lines = append(lines, "----------")
	}
	_, err := fmt.Fprintf(w, "%s\n", strings.Join(lines, "\t"))
	return err
}
