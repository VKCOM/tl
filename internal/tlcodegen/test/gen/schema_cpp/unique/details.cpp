#include "headers/unique.stringToInt.h"
#include "headers/unique.get.h"
#include "../__common_namespace/headers/int.h"


bool tl2::unique::Get::write_json(std::ostream& s)const {
	if (!::tl2::details::UniqueGetWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::unique::Get::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::UniqueGetRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::unique::Get::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::UniqueGetWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::unique::Get::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::unique::Get::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::unique::Get::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::UniqueGetReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::unique::Get::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::UniqueGetWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::unique::Get::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::unique::Get::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::UniqueGetReset(::tl2::unique::Get& item) noexcept {
	item.key.clear();
}

bool tl2::details::UniqueGetWriteJSON(std::ostream& s, const ::tl2::unique::Get& item) noexcept {
	s << "{";
	if (item.key.size() != 0) {
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::UniqueGetRead(::basictl::tl_istream & s, ::tl2::unique::Get& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	return true;
}

bool tl2::details::UniqueGetWrite(::basictl::tl_ostream & s, const ::tl2::unique::Get& item) noexcept {
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

bool tl2::unique::Get::read_result(::basictl::tl_istream & s, std::optional<int32_t> & result) noexcept {
	bool success = tl2::details::UniqueGetReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::unique::Get::write_result(::basictl::tl_ostream & s, std::optional<int32_t> & result) noexcept {
	bool success = tl2::details::UniqueGetWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::unique::Get::read_result_or_throw(::basictl::tl_throwable_istream & s, std::optional<int32_t> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::unique::Get::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::optional<int32_t> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::unique::StringToInt::write_json(std::ostream& s)const {
	if (!::tl2::details::UniqueStringToIntWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::unique::StringToInt::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::UniqueStringToIntRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::unique::StringToInt::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::UniqueStringToIntWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::unique::StringToInt::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::unique::StringToInt::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::unique::StringToInt::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::UniqueStringToIntReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::unique::StringToInt::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::UniqueStringToIntWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::unique::StringToInt::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::unique::StringToInt::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::UniqueStringToIntReset(::tl2::unique::StringToInt& item) noexcept {
	item.key.clear();
}

bool tl2::details::UniqueStringToIntWriteJSON(std::ostream& s, const ::tl2::unique::StringToInt& item) noexcept {
	s << "{";
	if (item.key.size() != 0) {
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::UniqueStringToIntRead(::basictl::tl_istream & s, ::tl2::unique::StringToInt& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	return true;
}

bool tl2::details::UniqueStringToIntWrite(::basictl::tl_ostream & s, const ::tl2::unique::StringToInt& item) noexcept {
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
	if (!s.nat_read_exact_tag(0xa8509bda)) { return false; }
	if (!s.int_read(result)) { return false; }
	return true;
}
bool tl2::details::UniqueStringToIntWriteResult(::basictl::tl_ostream & s, tl2::unique::StringToInt& item, int32_t& result) {
	if (!s.nat_write(0xa8509bda)) { return false; }
	if (!s.int_write(result)) { return false;}
	return true;
}

bool tl2::unique::StringToInt::read_result(::basictl::tl_istream & s, int32_t & result) noexcept {
	bool success = tl2::details::UniqueStringToIntReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::unique::StringToInt::write_result(::basictl::tl_ostream & s, int32_t & result) noexcept {
	bool success = tl2::details::UniqueStringToIntWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::unique::StringToInt::read_result_or_throw(::basictl::tl_throwable_istream & s, int32_t & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::unique::StringToInt::write_result_or_throw(::basictl::tl_throwable_ostream & s, int32_t & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}
