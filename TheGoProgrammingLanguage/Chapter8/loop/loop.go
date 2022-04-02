package main

import (
	_ "fmt"
)

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfiles string
		err error
	}

	ch := make(chan item, len(filenames))

	for _, f := range filenames {
		go func(f string) {
			var it item

			it.thumbfiles, it.err = thumbfiles.ImageFile(f)

			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch

		if it.err != nil {
			return nil, it.err
		}

		thumbfiles = append(thumbfiles, it.thumbfiles)
		// fmt.Println("1")
	}

	return thumbfiles, nil
}
