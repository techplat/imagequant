//+build !windows

package imagequant

///*
//#cgo CFLAGS: -O3 -fno-math-errno -fopenmp -funroll-loops -fomit-frame-pointer -Wall -Wno-attributes -std=c99 -DNDEBUG -DUSE_SSE=1 -msse -fexcess-precision=fast
//#cgo LDFLAGS: -lm -fopenmp
//*/
//import "C"

/*
#cgo CFLAGS: -O3 -fomit-frame-pointer -Wall -Wno-attributes -std=c99 -DNDEBUG -DUSE_SSE=1 -msse
#cgo LDFLAGS: -lm
*/
import "C"
