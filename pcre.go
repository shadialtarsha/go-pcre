// Package pcre is a library that provides pcre2 regular expressions
// in pure Go, allowing for features such as cross-compiling.
//
// The lib directory contains source code automatically translated from
// pcre2's C source code for each supported architecture and/or OS.
// This package wraps the automatically-translated source to provide a
// safe interface as close to Go's regexp library as possible.
package pcre

import (
	"runtime"
	"sync"

	"github.com/shadialtarsha/go-pcre/lib"

	"modernc.org/libc"
)

// Version returns the version of pcre2 embedded in this library.
func Version() string { return lib.DPACKAGE_VERSION }

// Regexp represents a pcre2 regular expression
type Regexp struct {
	mtx  *sync.Mutex
	expr string
	re   uintptr
	tls  *libc.TLS
}

// Compile runs CompileOpts with no options.
//
// Close() should be called on the returned expression
// once it is no longer needed.
func Compile(pattern string) (*Regexp, error) {
	return CompileOpts(pattern, 0)
}

// CompileOpts compiles the provided pattern using the given options.
//
// Close() should be called on the returned expression
// once it is no longer needed.
func CompileOpts(pattern string, options CompileOption) (*Regexp, error) {
	tls := libc.NewTLS()

	// Get C string of pattern
	cPattern, err := libc.CString(pattern)
	if err != nil {
		return nil, err
	}
	// Free the string when done
	defer libc.Xfree(tls, cPattern)

	// Allocate new error
	cErr := allocError(tls)
	// Free error when done
	defer libc.Xfree(tls, cErr)

	// Get error offsets
	errPtr := addErrCodeOffset(cErr)
	errOffsetPtr := addErrOffsetOffset(cErr)

	// Convert pattern length to size_t type
	cPatLen := lib.Tsize_t(len(pattern))

	// Compile expression
	r := lib.Xpcre2_compile_8(tls, cPattern, cPatLen, uint32(options), errPtr, errOffsetPtr, 0)
	if r == 0 {
		return nil, ptrToError(tls, cErr)
	}

	// Create regexp instance
	regex := Regexp{
		expr: pattern,
		mtx:  &sync.Mutex{},
		re:   r,
		tls:  tls,
	}

	// Make sure resources are freed if GC collects the
	// regular expression.
	runtime.SetFinalizer(&regex, func(r *Regexp) error {
		return r.Close()
	})

	return &regex, nil
}

// MustCompile compiles the given pattern and panics
// if there was an error
//
// Close() should be called on the returned expression
// once it is no longer needed.
func MustCompile(pattern string) *Regexp {
	rgx, err := Compile(pattern)
	if err != nil {
		panic(err)
	}
	return rgx
}

// MustCompileOpts compiles the given pattern with the given
// options and panics if there was an error.
//
// Close() should be called on the returned expression
// once it is no longer needed.
func MustCompileOpts(pattern string, options CompileOption) *Regexp {
	rgx, err := CompileOpts(pattern, options)
	if err != nil {
		panic(err)
	}
	return rgx
}

// Close frees resources used by the regular expression.
func (r *Regexp) Close() error {
	if r == nil {
		return nil
	}

	// Close thread-local storage
	defer r.tls.Close()

	// Free the compiled code
	lib.Xpcre2_code_free_8(r.tls, r.re)
	// Set regular expression to null
	r.re = 0

	return nil
}
