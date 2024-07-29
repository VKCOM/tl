#include "headers/pkg2.t2.hpp"
#include "headers/pkg2.t1.hpp"
#include "headers/pkg2.foo.hpp"


bool tl2::pkg2::Foo::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Pkg2FooRead(s, *this)) { return false; }
	return true;
}

bool tl2::pkg2::Foo::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Pkg2FooWrite(s, *this)) { return false; }
	return true;
}

bool tl2::pkg2::Foo::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Pkg2FooReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::pkg2::Foo::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Pkg2FooWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Pkg2FooReset(::tl2::pkg2::Foo& item) {
	item.x = 0;
}

bool tl2::details::Pkg2FooRead(::basictl::tl_istream & s, ::tl2::pkg2::Foo& item) {
	if (!s.int_read(item.x)) { return false; }
	return true;
}

bool tl2::details::Pkg2FooWrite(::basictl::tl_ostream & s, const ::tl2::pkg2::Foo& item) {
	if (!s.int_write(item.x)) { return false;}
	return true;
}

bool tl2::details::Pkg2FooReadBoxed(::basictl::tl_istream & s, ::tl2::pkg2::Foo& item) {
	if (!s.nat_read_exact_tag(0xe144703d)) { return false; }
	return tl2::details::Pkg2FooRead(s, item);
}

bool tl2::details::Pkg2FooWriteBoxed(::basictl::tl_ostream & s, const ::tl2::pkg2::Foo& item) {
	if (!s.nat_write(0xe144703d)) { return false; }
	return tl2::details::Pkg2FooWrite(s, item);
}

bool tl2::pkg2::T1::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Pkg2T1Read(s, *this)) { return false; }
	return true;
}

bool tl2::pkg2::T1::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Pkg2T1Write(s, *this)) { return false; }
	return true;
}

bool tl2::pkg2::T1::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Pkg2T1ReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::pkg2::T1::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Pkg2T1WriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Pkg2T1Reset(::tl2::pkg2::T1& item) {
	::tl2::details::Pkg2FooReset(item.x);
}

bool tl2::details::Pkg2T1Read(::basictl::tl_istream & s, ::tl2::pkg2::T1& item) {
	if (!::tl2::details::Pkg2FooReadBoxed(s, item.x)) { return false; }
	return true;
}

bool tl2::details::Pkg2T1Write(::basictl::tl_ostream & s, const ::tl2::pkg2::T1& item) {
	if (!::tl2::details::Pkg2FooWriteBoxed(s, item.x)) { return false; }
	return true;
}

bool tl2::details::Pkg2T1ReadBoxed(::basictl::tl_istream & s, ::tl2::pkg2::T1& item) {
	if (!s.nat_read_exact_tag(0x638206ec)) { return false; }
	return tl2::details::Pkg2T1Read(s, item);
}

bool tl2::details::Pkg2T1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::pkg2::T1& item) {
	if (!s.nat_write(0x638206ec)) { return false; }
	return tl2::details::Pkg2T1Write(s, item);
}

void tl2::details::Pkg2T2Reset(::tl2::pkg2::T2& item) {
	::tl2::details::Pkg2FooReset(item);
}

bool tl2::details::Pkg2T2Read(::basictl::tl_istream & s, ::tl2::pkg2::T2& item) {
	if (!::tl2::details::Pkg2FooRead(s, item)) { return false; }
	return true;
}

bool tl2::details::Pkg2T2Write(::basictl::tl_ostream & s, const ::tl2::pkg2::T2& item) {
	if (!::tl2::details::Pkg2FooWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::Pkg2T2ReadBoxed(::basictl::tl_istream & s, ::tl2::pkg2::T2& item) {
	if (!s.nat_read_exact_tag(0xd6e5af9c)) { return false; }
	return tl2::details::Pkg2T2Read(s, item);
}

bool tl2::details::Pkg2T2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::pkg2::T2& item) {
	if (!s.nat_write(0xd6e5af9c)) { return false; }
	return tl2::details::Pkg2T2Write(s, item);
}
