package main

import "unsafe"

type slice1 struct {
	cap int
	len int
	datas unsafe.Pointer{}
}

func make(length int) *slice1 {
	return &slice1 {
		len : length,
		cap : length,
		datas : malloc(unsafe.Pointer{},length),
	}
}

func this *slice1 Len() int {
	return this.len
}

func this *slice1 Cap() int {
	return this.cap
}

func Append(s *slice1,data int) *slice1 {
	newslice := (*s)

	if newslice.cap - newslice.len < 1 {
		tmp = malloc(unsafe.Pointer{},newslice.cap*2)
		copy(tmp,newslice.datas)
		newslice.datas = tmp
		newslice.cap = newslice.cap*2
	}

	newslice.datas[len] = unsafe.Pointer(data)
	newslice.len++

	return &newslice
}

//cap
//len
//make
//append

func main() {

}
