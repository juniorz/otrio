package otrio_test

import (
	"fmt"
	"strings"

	"github.com/juniorz/otrio"
)

func Example_NewScanner() {
	const buff = "hi there?OTR?v3??OTR:AAMDJ+MVmSfjFZcAAAAAAQAAAAIAAADA1g5IjD1ZGLDVQEyCgCyn9hbrL3KAbGDdzE2ZkMyTKl7XfkSxh8YJnudstiB74i4BzT0W2haClg6dMary/jo9sMudwmUdlnKpIGEKXWdvJKT+hQ26h9nzMgEditLB8vjPEWAJ6gBXvZrY6ZQrx3gb4v0UaSMOMiR5sB7Eaulb2Yc6RmRnnlxgUUC2alosg4WIeFN951PLjScajVba6dqlDi+q1H5tPvI5SWMN7PCBWIJ41+WvF+5IAZzQZYgNaVLbAAAAAAAAAAEAAAAHwNiIi5Ms+4PsY/L2ipkTtquknfx6HodLvk3RAAAAAA==."

	s := otrio.NewScanner(strings.NewReader(buff))

	i := 0
	for s.Scan() {
		i++
	}

	fmt.Println("->", i)

	// Output: -> 3
}
