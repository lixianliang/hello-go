package main

func lineParser(ctx context.Context, base int, in <-chan string) (<-chan int64, <-chan error, error) {
	out := make(chan int64)
	errc := make(chan error)
	go func() {
		defer close(out)
		defer close(errc)

		for line := range in {
			n, err := strconv.ParseInt(line, base, 64)
			if err != nil {
				errc <- err
				return
			}

			select {
			case out <- n: // 结果
			case <-ctx.Done(): // 关闭处理
				return
			}
		}
	}()

	return out, errc, nil
}
