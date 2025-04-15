#include "headers/pkg2.t2.h"
#include "headers/pkg2.t1.h"
#include "headers/pkg2.foo.h"


bool tl2::pkg2::Foo::write_json(std::ostream& s)const {
	if (!::tl2::details::Pkg2FooWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::pkg2::Foo::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Pkg2FooRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::pkg2::Foo::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Pkg2FooWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::pkg2::Foo::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::pkg2::Foo::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::pkg2::Foo::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Pkg2FooReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::pkg2::Foo::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Pkg2FooWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::pkg2::Foo::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::pkg2::Foo::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Pkg2FooReset(::tl2::pkg2::Foo& item) noexcept {
	item.x = 0;
}

bool tl2::details::Pkg2FooWriteJSON(std::ostream& s, const ::tl2::pkg2::Foo& item) noexcept {
	s << "{";
	if (item.x != 0) {
		s << "\"x\":";
		s << item.x;
	}
	s << "}";
	return true;
}

bool tl2::details::Pkg2FooRead(::basictl::tl_istream & s, ::tl2::pkg2::Foo& item) noexcept {
	if (!s.int_read(item.x)) { return false; }
	return true;
}

bool tl2::details::Pkg2FooWrite(::basictl::tl_ostream & s, const ::tl2::pkg2::Foo& item) noexcept {
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

bool tl2::pkg2::T1::write_json(std::ostream& s)const {
	if (!::tl2::details::Pkg2T1WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::pkg2::T1::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Pkg2T1Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::pkg2::T1::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Pkg2T1Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::pkg2::T1::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::pkg2::T1::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::pkg2::T1::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Pkg2T1ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::pkg2::T1::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Pkg2T1WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::pkg2::T1::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::pkg2::T1::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Pkg2T1Reset(::tl2::pkg2::T1& item) noexcept {
	::tl2::details::Pkg2FooReset(item.x);
}

bool tl2::details::Pkg2T1WriteJSON(std::ostream& s, const ::tl2::pkg2::T1& item) noexcept {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::Pkg2FooWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::Pkg2T1Read(::basictl::tl_istream & s, ::tl2::pkg2::T1& item) noexcept {
	if (!::tl2::details::Pkg2FooReadBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::Pkg2T1Write(::basictl::tl_ostream & s, const ::tl2::pkg2::T1& item) noexcept {
	if (!::tl2::details::Pkg2FooWriteBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
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

void tl2::details::Pkg2T2Reset(::tl2::pkg2::T2& item) noexcept {
	::tl2::details::Pkg2FooReset(item);
}

bool tl2::details::Pkg2T2WriteJSON(std::ostream& s, const ::tl2::pkg2::T2& item) noexcept {
	if (!::tl2::details::Pkg2FooWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::Pkg2T2Read(::basictl::tl_istream & s, ::tl2::pkg2::T2& item) noexcept {
	if (!::tl2::details::Pkg2FooRead(s, item)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::Pkg2T2Write(::basictl::tl_ostream & s, const ::tl2::pkg2::T2& item) noexcept {
	if (!::tl2::details::Pkg2FooWrite(s, item)) { return s.set_error_unknown_scenario(); }
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
