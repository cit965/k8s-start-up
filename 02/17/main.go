package main

import (
	"fmt"

	"github.com/kubernetes/apimachinery/pkg/labels"
)

func main() {
	labelSet := labels.Set{"app": "web", "tier": "frontend"}

	query := &labels.hasTerm{label: "app", value: "web"}

	if query.Matches(labelSet) {
		fmt.Println("Matched!")
	} else {
		fmt.Println("Not matched!")
	}
}
