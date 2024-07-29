#include "headers/unique.stringToInt.hpp"
#include "headers/unique.get.hpp"
#include "../../__common/details/headers/int.hpp"


bool tl2::unique::Get::read(::basictl::tl_istream & s) {
	if (!::tl2::details::UniqueGetRead(s, *this)) { return false; }
	return true;
}

bool tl2::unique::Get::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::UniqueGetWrite(s, *this)) { return false; }
	return true;
}

bool tl2::unique::Get::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::UniqueGetReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::unique::Get::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::UniqueGetWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::UniqueGetReset(::tl2::unique::Get& item) {
	item.key.clear();
}

bool tl2::details::UniqueGetRead(::basictl::tl_istream & s, ::tl2::unique::Get& item) {
	if (!s.string_read(item.key)) { return false; }
	return true;
}

bool tl2::details::UniqueGetWrite(::basictl::tl_ostream & s, const ::tl2::unique::Get& item) {
	if (!s.string_write(item.key)) { return false;}
	return true;
}

bool tl2::details::UniqueGetReadBoxed(::basictl::tl_istream & s, ::tl2::unique::Get& item) {
	if (!s.nat_read_exact_tag(0xce89bbf2)) { return false; }
	return tl2::details::UniqueGetRead(s, item);
}

bool tl2::details::UniqueGetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::unique::Get& item) {
	if (!s.nat_write(0xce89bbf2)) { return false; }
	return tl2::details::UniqueGetWrite(s, item);
}

bool tl2::details::UniqueGetReadResult(::basictl::tl_istream & s, tl2::unique::Get& item, std::optional<int32_t>& result) {
	if (!::tl2::details::IntMaybeReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::UniqueGetWriteResult(::basictl::tl_ostream & s, tl2::unique::Get& item, std::optional<int32_t>& result) {
	if (!::tl2::details::IntMaybeWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::unique::Get::read_result(::basictl::tl_istream & s, std::optional<int32_t> & result) {
	return tl2::details::UniqueGetReadResult(s, *this, result);
}
bool tl2::unique::Get::write_result(::basictl::tl_ostream & s, std::optional<int32_t> & result) {
	return tl2::details::UniqueGetWriteResult(s, *this, result);
}

bool tl2::unique::StringToInt::read(::basictl::tl_istream & s) {
	if (!::tl2::details::UniqueStringToIntRead(s, *this)) { return false; }
	return true;
}

bool tl2::unique::StringToInt::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::UniqueStringToIntWrite(s, *this)) { return false; }
	return true;
}

bool tl2::unique::StringToInt::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::UniqueStringToIntReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::unique::StringToInt::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::UniqueStringToIntWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::UniqueStringToIntReset(::tl2::unique::StringToInt& item) {
	item.key.clear();
}

bool tl2::details::UniqueStringToIntRead(::basictl::tl_istream & s, ::tl2::unique::StringToInt& item) {
	if (!s.string_read(item.key)) { return false; }
	return true;
}

bool tl2::details::UniqueStringToIntWrite(::basictl::tl_ostream & s, const ::tl2::unique::StringToInt& item) {
	if (!s.string_write(item.key)) { return false;}
	return true;
}

bool tl2::details::UniqueStringToIntReadBoxed(::basictl::tl_istream & s, ::tl2::unique::StringToInt& item) {
	if (!s.nat_read_exact_tag(0x0f766c35)) { return false; }
	return tl2::details::UniqueStringToIntRead(s, item);
}

bool tl2::details::UniqueStringToIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::unique::StringToInt& item) {
	if (!s.nat_write(0x0f766c35)) { return false; }
	return tl2::details::UniqueStringToIntWrite(s, item);
}

bool tl2::details::UniqueStringToIntReadResult(::basictl::tl_istream & s, tl2::unique::StringToInt& item, int32_t& result) {
	if (!s.nat_read_exact_tag(0xa8509bda)) { return false;}
	if (!s.int_read(result)) { return false; }
	return true;
}
bool tl2::details::UniqueStringToIntWriteResult(::basictl::tl_ostream & s, tl2::unique::StringToInt& item, int32_t& result) {
	if (!s.nat_write(0xa8509bda)) { return false; }
	if (!s.int_write(result)) { return false;}
	return true;
}

bool tl2::unique::StringToInt::read_result(::basictl::tl_istream & s, int32_t & result) {
	return tl2::details::UniqueStringToIntReadResult(s, *this, result);
}
bool tl2::unique::StringToInt::write_result(::basictl::tl_ostream & s, int32_t & result) {
	return tl2::details::UniqueStringToIntWriteResult(s, *this, result);
}
