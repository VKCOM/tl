#include "headers/string.hpp"
#include "headers/long.hpp"
#include "headers/int64.hpp"
#include "headers/int32.hpp"
#include "headers/int.hpp"


void tl2::details::IntReset(int32_t& item) {
	item = 0;
}

bool tl2::details::IntRead(::basictl::tl_istream & s, int32_t& item) {
	if (!s.int_read(item)) { return false; }
	return true;
}

bool tl2::details::IntWrite(::basictl::tl_ostream & s, const int32_t& item) {
	if (!s.int_write(item)) { return false;}
	return true;
}

bool tl2::details::IntReadBoxed(::basictl::tl_istream & s, int32_t& item) {
	if (!s.nat_read_exact_tag(0xa8509bda)) { return false; }
	return tl2::details::IntRead(s, item);
}

bool tl2::details::IntWriteBoxed(::basictl::tl_ostream & s, const int32_t& item) {
	if (!s.nat_write(0xa8509bda)) { return false; }
	return tl2::details::IntWrite(s, item);
}

void tl2::details::Int32Reset(::tl2::Int32& item) {
	item = 0;
}

bool tl2::details::Int32Read(::basictl::tl_istream & s, ::tl2::Int32& item) {
	if (!s.int_read(item)) { return false; }
	return true;
}

bool tl2::details::Int32Write(::basictl::tl_ostream & s, const ::tl2::Int32& item) {
	if (!s.int_write(item)) { return false;}
	return true;
}

bool tl2::details::Int32ReadBoxed(::basictl::tl_istream & s, ::tl2::Int32& item) {
	if (!s.nat_read_exact_tag(0x7934e71f)) { return false; }
	return tl2::details::Int32Read(s, item);
}

bool tl2::details::Int32WriteBoxed(::basictl::tl_ostream & s, const ::tl2::Int32& item) {
	if (!s.nat_write(0x7934e71f)) { return false; }
	return tl2::details::Int32Write(s, item);
}

void tl2::details::Int64Reset(::tl2::Int64& item) {
	item = 0;
}

bool tl2::details::Int64Read(::basictl::tl_istream & s, ::tl2::Int64& item) {
	if (!s.long_read(item)) { return false; }
	return true;
}

bool tl2::details::Int64Write(::basictl::tl_ostream & s, const ::tl2::Int64& item) {
	if (!s.long_write(item)) { return false;}
	return true;
}

bool tl2::details::Int64ReadBoxed(::basictl::tl_istream & s, ::tl2::Int64& item) {
	if (!s.nat_read_exact_tag(0xf5609de0)) { return false; }
	return tl2::details::Int64Read(s, item);
}

bool tl2::details::Int64WriteBoxed(::basictl::tl_ostream & s, const ::tl2::Int64& item) {
	if (!s.nat_write(0xf5609de0)) { return false; }
	return tl2::details::Int64Write(s, item);
}

void tl2::details::LongReset(int64_t& item) {
	item = 0;
}

bool tl2::details::LongRead(::basictl::tl_istream & s, int64_t& item) {
	if (!s.long_read(item)) { return false; }
	return true;
}

bool tl2::details::LongWrite(::basictl::tl_ostream & s, const int64_t& item) {
	if (!s.long_write(item)) { return false;}
	return true;
}

bool tl2::details::LongReadBoxed(::basictl::tl_istream & s, int64_t& item) {
	if (!s.nat_read_exact_tag(0x22076cba)) { return false; }
	return tl2::details::LongRead(s, item);
}

bool tl2::details::LongWriteBoxed(::basictl::tl_ostream & s, const int64_t& item) {
	if (!s.nat_write(0x22076cba)) { return false; }
	return tl2::details::LongWrite(s, item);
}

void tl2::details::StringReset(std::string& item) {
	item.clear();
}

bool tl2::details::StringRead(::basictl::tl_istream & s, std::string& item) {
	if (!s.string_read(item)) { return false; }
	return true;
}

bool tl2::details::StringWrite(::basictl::tl_ostream & s, const std::string& item) {
	if (!s.string_write(item)) { return false;}
	return true;
}

bool tl2::details::StringReadBoxed(::basictl::tl_istream & s, std::string& item) {
	if (!s.nat_read_exact_tag(0xb5286e24)) { return false; }
	return tl2::details::StringRead(s, item);
}

bool tl2::details::StringWriteBoxed(::basictl::tl_ostream & s, const std::string& item) {
	if (!s.nat_write(0xb5286e24)) { return false; }
	return tl2::details::StringWrite(s, item);
}
