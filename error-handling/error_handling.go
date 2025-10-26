package erratum

import (
	"errors"
	"fmt"
)

func Use(opener ResourceOpener, input string) (err error) {
	r, err := opener()
	if err != nil {
		var te TransientError
		if errors.As(err, &te) {
			return Use(opener, input)
		}
		return err
	}

	// In Go, deferred functions run in LIFO order — like a stack.
	// That means the `defer` declared last executes first when the
	// surrounding function returns (or panics).
	defer func() {
		if cerr := r.Close(); cerr != nil {
			if err != nil {
				err = fmt.Errorf("close error: %w (previous error: %w)", cerr, err)
			} else {
				err = cerr
			}
		}
	}()

	defer func() {
		// When the goroutine is not panicking, `recover` returns `nil`.
		if rec := recover(); rec != nil {
			// rec: any
			e := toError(rec)
			var fe FrobError
			if errors.As(e, &fe) {
				r.Defrob(fe.defrobTag)
			}
			err = e
		}
	}()

	r.Frob(input)
	return err
}

func toError(rec any) error {
	if err, ok := rec.(error); ok {
		return err
	}
	return fmt.Errorf("%v", rec)
}
