// Code generated by cmd/cgo; DO NOT EDIT.

package mkcp

import "unsafe"

import _ "runtime/cgo"

import "syscall"

var _ syscall.Errno
func _Cgo_ptr(ptr unsafe.Pointer) unsafe.Pointer { return ptr }

//go:linkname _Cgo_always_false runtime.cgoAlwaysFalse
var _Cgo_always_false bool
//go:linkname _Cgo_use runtime.cgoUse
func _Cgo_use(interface{})
type _Ctype_IINT32 = _Ctype_ISTDINT32

type _Ctype_ISTDINT32 = _Ctype_int

type _Ctype_ISTDUINT32 = _Ctype_uint

type _Ctype_IUINT32 = _Ctype_ISTDUINT32

type _Ctype_char int8

type _Ctype_ikcpcb = _Ctype_struct_IKCPCB

type _Ctype_int int32

type _Ctype_intgo = _Ctype_ptrdiff_t

type _Ctype_long int64

type _Ctype_ptrdiff_t = _Ctype_long

type _Ctype_struct_IKCPCB struct {
	conv		_Ctype_IUINT32
	mtu		_Ctype_IUINT32
	mss		_Ctype_IUINT32
	state		_Ctype_IUINT32
	snd_una		_Ctype_IUINT32
	snd_nxt		_Ctype_IUINT32
	rcv_nxt		_Ctype_IUINT32
	ts_recent	_Ctype_IUINT32
	ts_lastack	_Ctype_IUINT32
	ssthresh	_Ctype_IUINT32
	rx_rttval	_Ctype_IINT32
	rx_srtt		_Ctype_IINT32
	rx_rto		_Ctype_IINT32
	rx_minrto	_Ctype_IINT32
	snd_wnd		_Ctype_IUINT32
	rcv_wnd		_Ctype_IUINT32
	rmt_wnd		_Ctype_IUINT32
	cwnd		_Ctype_IUINT32
	probe		_Ctype_IUINT32
	current		_Ctype_IUINT32
	interval	_Ctype_IUINT32
	ts_flush	_Ctype_IUINT32
	xmit		_Ctype_IUINT32
	nrcv_buf	_Ctype_IUINT32
	nsnd_buf	_Ctype_IUINT32
	nrcv_que	_Ctype_IUINT32
	nsnd_que	_Ctype_IUINT32
	nodelay		_Ctype_IUINT32
	updated		_Ctype_IUINT32
	ts_probe	_Ctype_IUINT32
	probe_wait	_Ctype_IUINT32
	dead_link	_Ctype_IUINT32
	incr		_Ctype_IUINT32
	snd_queue	_Ctype_struct_IQUEUEHEAD
	rcv_queue	_Ctype_struct_IQUEUEHEAD
	snd_buf		_Ctype_struct_IQUEUEHEAD
	rcv_buf		_Ctype_struct_IQUEUEHEAD
	acklist		*_Ctype_IUINT32
	ackcount	_Ctype_IUINT32
	ackblock	_Ctype_IUINT32
	user		unsafe.Pointer
	buffer		*_Ctype_char
	fastresend	_Ctype_int
	fastlimit	_Ctype_int
	nocwnd		_Ctype_int
	stream		_Ctype_int
	logmask		_Ctype_int
	output		*[0]byte
	writelog	*[0]byte
}

type _Ctype_struct_IQUEUEHEAD struct {
	next	*_Ctype_struct_IQUEUEHEAD
	prev	*_Ctype_struct_IQUEUEHEAD
}

type _Ctype_uint uint32

type _Ctype_uintptr_t = _Ctype_ulong

type _Ctype_ulong uint64

type _Ctype_void [0]byte

//go:linkname _cgo_runtime_cgocall runtime.cgocall
func _cgo_runtime_cgocall(unsafe.Pointer, uintptr) int32

//go:linkname _cgo_runtime_cgocallback runtime.cgocallback
func _cgo_runtime_cgocallback(unsafe.Pointer, unsafe.Pointer, uintptr, uintptr)

//go:linkname _cgoCheckPointer runtime.cgoCheckPointer
func _cgoCheckPointer(interface{}, ...interface{})

//go:linkname _cgoCheckResult runtime.cgoCheckResult
func _cgoCheckResult(interface{})


func _Cfunc_CBytes(b []byte) unsafe.Pointer {
	p := _cgo_cmalloc(uint64(len(b)))
	pp := (*[1<<30]byte)(p)
	copy(pp[:], b)
	return p
}
//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc_free
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc_free _cgo_0b42451c2bb2_Cfunc_free
var __cgofn__cgo_0b42451c2bb2_Cfunc_free byte
var _cgo_0b42451c2bb2_Cfunc_free = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc_free)

//go:cgo_unsafe_args
func _Cfunc_free(p0 unsafe.Pointer) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc_free, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc_ikcp_check
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_check _cgo_0b42451c2bb2_Cfunc_ikcp_check
var __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_check byte
var _cgo_0b42451c2bb2_Cfunc_ikcp_check = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_check)

//go:cgo_unsafe_args
func _Cfunc_ikcp_check(p0 *_Ctype_struct_IKCPCB, p1 _Ctype_IUINT32) (r1 _Ctype_IUINT32) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc_ikcp_check, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc_ikcp_flush
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_flush _cgo_0b42451c2bb2_Cfunc_ikcp_flush
var __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_flush byte
var _cgo_0b42451c2bb2_Cfunc_ikcp_flush = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_flush)

//go:cgo_unsafe_args
func _Cfunc_ikcp_flush(p0 *_Ctype_struct_IKCPCB) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc_ikcp_flush, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc_ikcp_getconv
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_getconv _cgo_0b42451c2bb2_Cfunc_ikcp_getconv
var __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_getconv byte
var _cgo_0b42451c2bb2_Cfunc_ikcp_getconv = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_getconv)

//go:cgo_unsafe_args
func _Cfunc_ikcp_getconv(p0 unsafe.Pointer) (r1 _Ctype_IUINT32) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc_ikcp_getconv, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc_ikcp_nodelay
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_nodelay _cgo_0b42451c2bb2_Cfunc_ikcp_nodelay
var __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_nodelay byte
var _cgo_0b42451c2bb2_Cfunc_ikcp_nodelay = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_nodelay)

//go:cgo_unsafe_args
func _Cfunc_ikcp_nodelay(p0 *_Ctype_struct_IKCPCB, p1 _Ctype_int, p2 _Ctype_int, p3 _Ctype_int, p4 _Ctype_int) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc_ikcp_nodelay, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
		_Cgo_use(p2)
		_Cgo_use(p3)
		_Cgo_use(p4)
	}
	return
}
//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc_ikcp_peeksize
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_peeksize _cgo_0b42451c2bb2_Cfunc_ikcp_peeksize
var __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_peeksize byte
var _cgo_0b42451c2bb2_Cfunc_ikcp_peeksize = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_peeksize)

//go:cgo_unsafe_args
func _Cfunc_ikcp_peeksize(p0 *_Ctype_struct_IKCPCB) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc_ikcp_peeksize, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc_ikcp_setmtu
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_setmtu _cgo_0b42451c2bb2_Cfunc_ikcp_setmtu
var __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_setmtu byte
var _cgo_0b42451c2bb2_Cfunc_ikcp_setmtu = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_setmtu)

//go:cgo_unsafe_args
func _Cfunc_ikcp_setmtu(p0 *_Ctype_struct_IKCPCB, p1 _Ctype_int) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc_ikcp_setmtu, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc_ikcp_update
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_update _cgo_0b42451c2bb2_Cfunc_ikcp_update
var __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_update byte
var _cgo_0b42451c2bb2_Cfunc_ikcp_update = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_update)

//go:cgo_unsafe_args
func _Cfunc_ikcp_update(p0 *_Ctype_struct_IKCPCB, p1 _Ctype_IUINT32) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc_ikcp_update, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc_ikcp_waitsnd
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_waitsnd _cgo_0b42451c2bb2_Cfunc_ikcp_waitsnd
var __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_waitsnd byte
var _cgo_0b42451c2bb2_Cfunc_ikcp_waitsnd = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_waitsnd)

//go:cgo_unsafe_args
func _Cfunc_ikcp_waitsnd(p0 *_Ctype_struct_IKCPCB) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc_ikcp_waitsnd, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc_ikcp_wndsize
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_wndsize _cgo_0b42451c2bb2_Cfunc_ikcp_wndsize
var __cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_wndsize byte
var _cgo_0b42451c2bb2_Cfunc_ikcp_wndsize = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc_ikcp_wndsize)

//go:cgo_unsafe_args
func _Cfunc_ikcp_wndsize(p0 *_Ctype_struct_IKCPCB, p1 _Ctype_int, p2 _Ctype_int) (r1 _Ctype_int) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc_ikcp_wndsize, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
		_Cgo_use(p2)
	}
	return
}
//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc_mkcp_create
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc_mkcp_create _cgo_0b42451c2bb2_Cfunc_mkcp_create
var __cgofn__cgo_0b42451c2bb2_Cfunc_mkcp_create byte
var _cgo_0b42451c2bb2_Cfunc_mkcp_create = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc_mkcp_create)

//go:cgo_unsafe_args
func _Cfunc_mkcp_create(p0 _Ctype_IUINT32, p1 _Ctype_uintptr_t) (r1 *_Ctype_struct_IKCPCB) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc_mkcp_create, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
		_Cgo_use(p1)
	}
	return
}
//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc_mkcp_release
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc_mkcp_release _cgo_0b42451c2bb2_Cfunc_mkcp_release
var __cgofn__cgo_0b42451c2bb2_Cfunc_mkcp_release byte
var _cgo_0b42451c2bb2_Cfunc_mkcp_release = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc_mkcp_release)

//go:cgo_unsafe_args
func _Cfunc_mkcp_release(p0 *_Ctype_struct_IKCPCB) (r1 _Ctype_void) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc_mkcp_release, uintptr(unsafe.Pointer(&p0)))
	if _Cgo_always_false {
		_Cgo_use(p0)
	}
	return
}
//go:cgo_export_dynamic go_output
//go:linkname _cgoexp_0b42451c2bb2_go_output _cgoexp_0b42451c2bb2_go_output
//go:cgo_export_static _cgoexp_0b42451c2bb2_go_output
//go:nosplit
//go:norace
func _cgoexp_0b42451c2bb2_go_output(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_0b42451c2bb2_go_output
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_0b42451c2bb2_go_output(p0 []byte, p1 uintptr) (r0 int) {
	return go_output(p0, p1)
}
//go:cgo_export_dynamic go_malloc
//go:linkname _cgoexp_0b42451c2bb2_go_malloc _cgoexp_0b42451c2bb2_go_malloc
//go:cgo_export_static _cgoexp_0b42451c2bb2_go_malloc
//go:nosplit
//go:norace
func _cgoexp_0b42451c2bb2_go_malloc(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_0b42451c2bb2_go_malloc
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_0b42451c2bb2_go_malloc(p0 uint) (r0 uintptr) {
	defer func() {
		_cgoCheckResult(r0)
	}()
	return go_malloc(p0)
}
//go:cgo_export_dynamic go_free
//go:linkname _cgoexp_0b42451c2bb2_go_free _cgoexp_0b42451c2bb2_go_free
//go:cgo_export_static _cgoexp_0b42451c2bb2_go_free
//go:nosplit
//go:norace
func _cgoexp_0b42451c2bb2_go_free(a unsafe.Pointer, n int32, ctxt uintptr) {
	fn := _cgoexpwrap_0b42451c2bb2_go_free
	_cgo_runtime_cgocallback(**(**unsafe.Pointer)(unsafe.Pointer(&fn)), a, uintptr(n), ctxt);
}

func _cgoexpwrap_0b42451c2bb2_go_free(p0 uintptr) {
	go_free(p0)
}

//go:cgo_import_static _cgo_0b42451c2bb2_Cfunc__Cmalloc
//go:linkname __cgofn__cgo_0b42451c2bb2_Cfunc__Cmalloc _cgo_0b42451c2bb2_Cfunc__Cmalloc
var __cgofn__cgo_0b42451c2bb2_Cfunc__Cmalloc byte
var _cgo_0b42451c2bb2_Cfunc__Cmalloc = unsafe.Pointer(&__cgofn__cgo_0b42451c2bb2_Cfunc__Cmalloc)

//go:linkname runtime_throw runtime.throw
func runtime_throw(string)

//go:cgo_unsafe_args
func _cgo_cmalloc(p0 uint64) (r1 unsafe.Pointer) {
	_cgo_runtime_cgocall(_cgo_0b42451c2bb2_Cfunc__Cmalloc, uintptr(unsafe.Pointer(&p0)))
	if r1 == nil {
		runtime_throw("runtime: C malloc failed")
	}
	return
}
