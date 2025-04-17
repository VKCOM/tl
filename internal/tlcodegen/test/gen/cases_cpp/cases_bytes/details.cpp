#include "cases_bytes/headers/cases_bytes_vector.h"
#include "cases_bytes/headers/cases_bytes_tuple.h"
#include "cases_bytes/headers/cases_bytes.testVector.h"
#include "cases_bytes/headers/cases_bytes.testTuple.h"
#include "cases_bytes/headers/cases_bytes.testEnumContainer.h"
#include "cases/headers/cases.TestEnum.h"
#include "cases_bytes/headers/cases_bytes.TestEnum.h"
#include "cases_bytes/headers/cases_bytes.TestEnumItems.h"
#include "cases_bytes/headers/cases_bytes.testDictStringString.h"
#include "cases_bytes/headers/cases_bytes_dictionary.h"
#include "cases_bytes/headers/cases_bytes.testDictString.h"
#include "cases_bytes/headers/cases_bytes.testDictInt.h"
#include "__common_namespace/headers/dictionaryFieldAny.h"
#include "cases_bytes/headers/cases_bytes.testDictAny.h"
#include "__common_namespace/headers/dictionaryAny.h"
#include "cases_bytes/headers/cases_bytes.testArray.h"
#include "cases_bytes/headers/cases_bytes_string.h"
#include "cases_bytes/headers/cases_bytes_dictionaryField.h"
#include "__common_namespace/headers/dictionaryField.h"


void tl2::details::BuiltinTuple4StringReset(std::array<std::string, 4>& item) {
	for(auto && el : item) {
		el.clear();
	}
}

bool tl2::details::BuiltinTuple4StringWriteJSON(std::ostream &s, const std::array<std::string, 4>& item) {
	s << "[";
	size_t index = 0;
	for(auto && el : item) {
		s << "\"" << el << "\"";
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTuple4StringRead(::basictl::tl_istream & s, std::array<std::string, 4>& item) {
	for(auto && el : item) {
		if (!s.string_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTuple4StringWrite(::basictl::tl_ostream & s, const std::array<std::string, 4>& item) {
	for(const auto & el : item) {
		if (!s.string_write(el)) { return false;}
	}
	return true;
}

void tl2::details::BuiltinTupleStringReset(std::vector<std::string>& item) {
	item.resize(0);
}

bool tl2::details::BuiltinTupleStringWriteJSON(std::ostream & s, const std::vector<std::string>& item, uint32_t nat_n) {
	if (item.size() != nat_n) { return false; }
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		s << "\"" << el << "\"";
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTupleStringRead(::basictl::tl_istream & s, std::vector<std::string>& item, uint32_t nat_n) {
	// TODO - check length sanity
	item.resize(nat_n);
	for(auto && el : item) {
		if (!s.string_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTupleStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item, uint32_t nat_n) {
	if (item.size() != nat_n) { return s.set_error_sequence_length(); }
	for(const auto & el : item) {
		if (!s.string_write(el)) { return false;}
	}
	return true;
}

void tl2::details::BuiltinVectorDictionaryFieldStringReset(std::map<std::string, std::string>& item) {
	item.clear(); // TODO - unwrap
}

bool tl2::details::BuiltinVectorDictionaryFieldStringWriteJSON(std::ostream & s, const std::map<std::string, std::string>& item) {
	s << "{";
	size_t index = 0;
	for(const auto & el : item) {
		s << "\"" << el.first << "\"";
		s << ":";
		s << "\"" << el.second << "\"";
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "}";
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldStringRead(::basictl::tl_istream & s, std::map<std::string, std::string>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	item.clear();
	for(uint32_t i = 0; i < len; i++) {
		std::string key;
		if (!s.string_read(key)) { return false; }
		if (!s.string_read(item[key])) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldStringWrite(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!s.string_write(el.first)) { return false;}
		if (!s.string_write(el.second)) { return false;}
	}
	return true;
}

void tl2::details::BuiltinVectorStringReset(std::vector<std::string>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorStringWriteJSON(std::ostream & s, const std::vector<std::string>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		s << "\"" << el << "\"";
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorStringRead(::basictl::tl_istream & s, std::vector<std::string>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!s.string_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!s.string_write(el)) { return false;}
	}
	return true;
}

bool tl2::cases_bytes::TestArray::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesBytesTestArrayWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases_bytes::TestArray::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestArrayRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestArray::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestArrayWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestArray::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestArray::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::cases_bytes::TestArray::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestArrayReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestArray::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestArrayWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestArray::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestArray::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::CasesBytesTestArrayReset(::tl2::cases_bytes::TestArray& item) noexcept {
	item.n = 0;
	item.arr.clear();
}

bool tl2::details::CasesBytesTestArrayWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestArray& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.n != 0) {
		add_comma = true;
		s << "\"n\":";
		s << item.n;
	}
	if ((item.arr.size() != 0) || (item.n != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"arr\":";
		if (!::tl2::details::BuiltinTupleStringWriteJSON(s, item.arr, item.n)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesBytesTestArrayRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestArray& item) noexcept {
	if (!s.nat_read(item.n)) { return false; }
	if (!::tl2::details::BuiltinTupleStringRead(s, item.arr, item.n)) { return false; }
	return true;
}

bool tl2::details::CasesBytesTestArrayWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestArray& item) noexcept {
	if (!s.nat_write(item.n)) { return false;}
	if (!::tl2::details::BuiltinTupleStringWrite(s, item.arr, item.n)) { return false; }
	return true;
}

bool tl2::details::CasesBytesTestArrayReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestArray& item) {
	if (!s.nat_read_exact_tag(0x3762fb81)) { return false; }
	return tl2::details::CasesBytesTestArrayRead(s, item);
}

bool tl2::details::CasesBytesTestArrayWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestArray& item) {
	if (!s.nat_write(0x3762fb81)) { return false; }
	return tl2::details::CasesBytesTestArrayWrite(s, item);
}

bool tl2::cases_bytes::TestDictAny::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesBytesTestDictAnyWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases_bytes::TestDictAny::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestDictAnyRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestDictAny::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestDictAnyWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestDictAny::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestDictAny::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::cases_bytes::TestDictAny::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestDictAnyReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestDictAny::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestDictAnyWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestDictAny::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestDictAny::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::CasesBytesTestDictAnyReset(::tl2::cases_bytes::TestDictAny& item) noexcept {
	::tl2::details::DictionaryAnyDoubleIntReset(item.dict);
}

bool tl2::details::CasesBytesTestDictAnyWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestDictAny& item) noexcept {
	s << "{";
	if (item.dict.size() != 0) {
		s << "\"dict\":";
		if (!::tl2::details::DictionaryAnyDoubleIntWriteJSON(s, item.dict)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesBytesTestDictAnyRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictAny& item) noexcept {
	if (!::tl2::details::DictionaryAnyDoubleIntRead(s, item.dict)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::CasesBytesTestDictAnyWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictAny& item) noexcept {
	if (!::tl2::details::DictionaryAnyDoubleIntWrite(s, item.dict)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::CasesBytesTestDictAnyReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictAny& item) {
	if (!s.nat_read_exact_tag(0x5a5fce57)) { return false; }
	return tl2::details::CasesBytesTestDictAnyRead(s, item);
}

bool tl2::details::CasesBytesTestDictAnyWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictAny& item) {
	if (!s.nat_write(0x5a5fce57)) { return false; }
	return tl2::details::CasesBytesTestDictAnyWrite(s, item);
}

bool tl2::cases_bytes::TestDictInt::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesBytesTestDictIntWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases_bytes::TestDictInt::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestDictIntRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestDictInt::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestDictIntWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestDictInt::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestDictInt::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::cases_bytes::TestDictInt::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestDictIntReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestDictInt::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestDictIntWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestDictInt::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestDictInt::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::CasesBytesTestDictIntReset(::tl2::cases_bytes::TestDictInt& item) noexcept {
	item.dict.clear();
}

bool tl2::details::CasesBytesTestDictIntWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestDictInt& item) noexcept {
	s << "{";
	if (item.dict.size() != 0) {
		s << "\"dict\":";
		if (!::tl2::details::BuiltinVectorDictionaryFieldAnyIntIntWriteJSON(s, item.dict)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesBytesTestDictIntRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictInt& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyIntIntRead(s, item.dict)) { return false; }
	return true;
}

bool tl2::details::CasesBytesTestDictIntWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictInt& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyIntIntWrite(s, item.dict)) { return false; }
	return true;
}

bool tl2::details::CasesBytesTestDictIntReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictInt& item) {
	if (!s.nat_read_exact_tag(0x453ace07)) { return false; }
	return tl2::details::CasesBytesTestDictIntRead(s, item);
}

bool tl2::details::CasesBytesTestDictIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictInt& item) {
	if (!s.nat_write(0x453ace07)) { return false; }
	return tl2::details::CasesBytesTestDictIntWrite(s, item);
}

bool tl2::cases_bytes::TestDictString::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesBytesTestDictStringWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases_bytes::TestDictString::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestDictStringRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestDictString::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestDictStringWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestDictString::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestDictString::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::cases_bytes::TestDictString::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestDictStringReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestDictString::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestDictStringWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestDictString::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestDictString::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::CasesBytesTestDictStringReset(::tl2::cases_bytes::TestDictString& item) noexcept {
	item.dict.clear();
}

bool tl2::details::CasesBytesTestDictStringWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestDictString& item) noexcept {
	s << "{";
	if (item.dict.size() != 0) {
		s << "\"dict\":";
		if (!::tl2::details::BuiltinVectorDictionaryFieldIntWriteJSON(s, item.dict)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesBytesTestDictStringRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictString& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntRead(s, item.dict)) { return false; }
	return true;
}

bool tl2::details::CasesBytesTestDictStringWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictString& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWrite(s, item.dict)) { return false; }
	return true;
}

bool tl2::details::CasesBytesTestDictStringReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictString& item) {
	if (!s.nat_read_exact_tag(0x6c04d6ce)) { return false; }
	return tl2::details::CasesBytesTestDictStringRead(s, item);
}

bool tl2::details::CasesBytesTestDictStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictString& item) {
	if (!s.nat_write(0x6c04d6ce)) { return false; }
	return tl2::details::CasesBytesTestDictStringWrite(s, item);
}

bool tl2::cases_bytes::TestDictStringString::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesBytesTestDictStringStringWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases_bytes::TestDictStringString::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestDictStringStringRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestDictStringString::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestDictStringStringWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestDictStringString::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestDictStringString::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::cases_bytes::TestDictStringString::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestDictStringStringReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestDictStringString::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestDictStringStringWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestDictStringString::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestDictStringString::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::CasesBytesTestDictStringStringReset(::tl2::cases_bytes::TestDictStringString& item) noexcept {
	item.dict.clear();
}

bool tl2::details::CasesBytesTestDictStringStringWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestDictStringString& item) noexcept {
	s << "{";
	if (item.dict.size() != 0) {
		s << "\"dict\":";
		if (!::tl2::details::BuiltinVectorDictionaryFieldStringWriteJSON(s, item.dict)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesBytesTestDictStringStringRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictStringString& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldStringRead(s, item.dict)) { return false; }
	return true;
}

bool tl2::details::CasesBytesTestDictStringStringWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictStringString& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldStringWrite(s, item.dict)) { return false; }
	return true;
}

bool tl2::details::CasesBytesTestDictStringStringReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestDictStringString& item) {
	if (!s.nat_read_exact_tag(0xad69c772)) { return false; }
	return tl2::details::CasesBytesTestDictStringStringRead(s, item);
}

bool tl2::details::CasesBytesTestDictStringStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestDictStringString& item) {
	if (!s.nat_write(0xad69c772)) { return false; }
	return tl2::details::CasesBytesTestDictStringStringWrite(s, item);
}

static const std::string_view CasesBytesTestEnum_tbl_tl_name[]{"cases_bytes.testEnum1", "cases_bytes.testEnum2", "cases_bytes.testEnum3"};
static const uint32_t CasesBytesTestEnum_tbl_tl_tag[]{0x58aad3f5, 0x00b47add, 0x81911ffa};

bool tl2::cases_bytes::TestEnum::write_json(std::ostream & s)const {
	if (!::tl2::details::CasesBytesTestEnumWriteJSON(s, *this)) { return false; }
	return true;
}
bool tl2::cases_bytes::TestEnum::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestEnumReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::cases_bytes::TestEnum::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestEnumWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	return true;
}

void tl2::cases_bytes::TestEnum::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestEnum::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

std::string_view tl2::cases_bytes::TestEnum::tl_name() const {
	return CasesBytesTestEnum_tbl_tl_name[value.index()];
}
uint32_t tl2::cases_bytes::TestEnum::tl_tag() const {
	return CasesBytesTestEnum_tbl_tl_tag[value.index()];
}


void tl2::details::CasesBytesTestEnumReset(::tl2::cases_bytes::TestEnum& item) noexcept{
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::CasesBytesTestEnumWriteJSON(std::ostream & s, const ::tl2::cases_bytes::TestEnum& item) noexcept {
	s << "\"" << CasesBytesTestEnum_tbl_tl_name[item.value.index()] << "\"";
	return true;
}
bool tl2::details::CasesBytesTestEnumReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum& item) noexcept {
	uint32_t nat;
	if (!s.nat_read(nat)) { return false; }
	switch (nat) {
	case 0x58aad3f5:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		break;
	case 0x00b47add:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		break;
	case 0x81911ffa:
		if (item.value.index() != 2) { item.value.emplace<2>(); }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::CasesBytesTestEnumWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum& item) noexcept{
	if (!s.nat_write(CasesBytesTestEnum_tbl_tl_tag[item.value.index()])) { return false; }
	switch (item.value.index()) {
	}
	return true;
}

bool tl2::cases_bytes::TestEnum1::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesBytesTestEnum1WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases_bytes::TestEnum1::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestEnum1Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestEnum1::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestEnum1Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestEnum1::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestEnum1::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::cases_bytes::TestEnum1::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestEnum1ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestEnum1::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestEnum1WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestEnum1::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestEnum1::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::CasesBytesTestEnum1Reset(::tl2::cases_bytes::TestEnum1& item) noexcept {
}

bool tl2::details::CasesBytesTestEnum1WriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestEnum1& item) noexcept {
	s << "true";
	return true;
}

bool tl2::details::CasesBytesTestEnum1Read(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum1& item) noexcept {
	return true;
}

bool tl2::details::CasesBytesTestEnum1Write(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum1& item) noexcept {
	return true;
}

bool tl2::details::CasesBytesTestEnum1ReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum1& item) {
	if (!s.nat_read_exact_tag(0x58aad3f5)) { return false; }
	return tl2::details::CasesBytesTestEnum1Read(s, item);
}

bool tl2::details::CasesBytesTestEnum1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum1& item) {
	if (!s.nat_write(0x58aad3f5)) { return false; }
	return tl2::details::CasesBytesTestEnum1Write(s, item);
}

bool tl2::cases_bytes::TestEnum2::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesBytesTestEnum2WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases_bytes::TestEnum2::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestEnum2Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestEnum2::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestEnum2Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestEnum2::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestEnum2::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::cases_bytes::TestEnum2::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestEnum2ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestEnum2::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestEnum2WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestEnum2::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestEnum2::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::CasesBytesTestEnum2Reset(::tl2::cases_bytes::TestEnum2& item) noexcept {
}

bool tl2::details::CasesBytesTestEnum2WriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestEnum2& item) noexcept {
	s << "true";
	return true;
}

bool tl2::details::CasesBytesTestEnum2Read(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum2& item) noexcept {
	return true;
}

bool tl2::details::CasesBytesTestEnum2Write(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum2& item) noexcept {
	return true;
}

bool tl2::details::CasesBytesTestEnum2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum2& item) {
	if (!s.nat_read_exact_tag(0x00b47add)) { return false; }
	return tl2::details::CasesBytesTestEnum2Read(s, item);
}

bool tl2::details::CasesBytesTestEnum2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum2& item) {
	if (!s.nat_write(0x00b47add)) { return false; }
	return tl2::details::CasesBytesTestEnum2Write(s, item);
}

bool tl2::cases_bytes::TestEnum3::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesBytesTestEnum3WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases_bytes::TestEnum3::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestEnum3Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestEnum3::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestEnum3Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestEnum3::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestEnum3::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::cases_bytes::TestEnum3::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestEnum3ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestEnum3::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestEnum3WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestEnum3::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestEnum3::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::CasesBytesTestEnum3Reset(::tl2::cases_bytes::TestEnum3& item) noexcept {
}

bool tl2::details::CasesBytesTestEnum3WriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestEnum3& item) noexcept {
	s << "true";
	return true;
}

bool tl2::details::CasesBytesTestEnum3Read(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum3& item) noexcept {
	return true;
}

bool tl2::details::CasesBytesTestEnum3Write(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum3& item) noexcept {
	return true;
}

bool tl2::details::CasesBytesTestEnum3ReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum3& item) {
	if (!s.nat_read_exact_tag(0x81911ffa)) { return false; }
	return tl2::details::CasesBytesTestEnum3Read(s, item);
}

bool tl2::details::CasesBytesTestEnum3WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum3& item) {
	if (!s.nat_write(0x81911ffa)) { return false; }
	return tl2::details::CasesBytesTestEnum3Write(s, item);
}

bool tl2::cases_bytes::TestEnumContainer::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesBytesTestEnumContainerWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases_bytes::TestEnumContainer::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestEnumContainerRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestEnumContainer::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestEnumContainerWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestEnumContainer::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestEnumContainer::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::cases_bytes::TestEnumContainer::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestEnumContainerReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestEnumContainer::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestEnumContainerWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestEnumContainer::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestEnumContainer::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::CasesBytesTestEnumContainerReset(::tl2::cases_bytes::TestEnumContainer& item) noexcept {
	::tl2::details::CasesTestEnumReset(item.value);
}

bool tl2::details::CasesBytesTestEnumContainerWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestEnumContainer& item) noexcept {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::CasesTestEnumWriteJSON(s, item.value)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesBytesTestEnumContainerRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnumContainer& item) noexcept {
	if (!::tl2::details::CasesTestEnumReadBoxed(s, item.value)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::CasesBytesTestEnumContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnumContainer& item) noexcept {
	if (!::tl2::details::CasesTestEnumWriteBoxed(s, item.value)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::CasesBytesTestEnumContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnumContainer& item) {
	if (!s.nat_read_exact_tag(0x32b92037)) { return false; }
	return tl2::details::CasesBytesTestEnumContainerRead(s, item);
}

bool tl2::details::CasesBytesTestEnumContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnumContainer& item) {
	if (!s.nat_write(0x32b92037)) { return false; }
	return tl2::details::CasesBytesTestEnumContainerWrite(s, item);
}

bool tl2::cases_bytes::TestTuple::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesBytesTestTupleWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases_bytes::TestTuple::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestTupleRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestTuple::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestTupleWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestTuple::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestTuple::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::cases_bytes::TestTuple::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestTupleReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestTuple::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestTupleWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestTuple::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestTuple::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::CasesBytesTestTupleReset(::tl2::cases_bytes::TestTuple& item) noexcept {
	::tl2::details::BuiltinTuple4StringReset(item.tpl);
}

bool tl2::details::CasesBytesTestTupleWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestTuple& item) noexcept {
	s << "{";
	s << "\"tpl\":";
	if (!::tl2::details::BuiltinTuple4StringWriteJSON(s, item.tpl)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesBytesTestTupleRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestTuple& item) noexcept {
	if (!::tl2::details::BuiltinTuple4StringRead(s, item.tpl)) { return false; }
	return true;
}

bool tl2::details::CasesBytesTestTupleWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestTuple& item) noexcept {
	if (!::tl2::details::BuiltinTuple4StringWrite(s, item.tpl)) { return false; }
	return true;
}

bool tl2::details::CasesBytesTestTupleReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestTuple& item) {
	if (!s.nat_read_exact_tag(0x2dd3bacf)) { return false; }
	return tl2::details::CasesBytesTestTupleRead(s, item);
}

bool tl2::details::CasesBytesTestTupleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestTuple& item) {
	if (!s.nat_write(0x2dd3bacf)) { return false; }
	return tl2::details::CasesBytesTestTupleWrite(s, item);
}

bool tl2::cases_bytes::TestVector::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesBytesTestVectorWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases_bytes::TestVector::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestVectorRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestVector::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestVectorWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestVector::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestVector::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::cases_bytes::TestVector::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::CasesBytesTestVectorReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::cases_bytes::TestVector::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::CasesBytesTestVectorWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::cases_bytes::TestVector::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::cases_bytes::TestVector::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::CasesBytesTestVectorReset(::tl2::cases_bytes::TestVector& item) noexcept {
	item.arr.clear();
}

bool tl2::details::CasesBytesTestVectorWriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestVector& item) noexcept {
	s << "{";
	if (item.arr.size() != 0) {
		s << "\"arr\":";
		if (!::tl2::details::BuiltinVectorStringWriteJSON(s, item.arr)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesBytesTestVectorRead(::basictl::tl_istream & s, ::tl2::cases_bytes::TestVector& item) noexcept {
	if (!::tl2::details::BuiltinVectorStringRead(s, item.arr)) { return false; }
	return true;
}

bool tl2::details::CasesBytesTestVectorWrite(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestVector& item) noexcept {
	if (!::tl2::details::BuiltinVectorStringWrite(s, item.arr)) { return false; }
	return true;
}

bool tl2::details::CasesBytesTestVectorReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestVector& item) {
	if (!s.nat_read_exact_tag(0x3647c8ae)) { return false; }
	return tl2::details::CasesBytesTestVectorRead(s, item);
}

bool tl2::details::CasesBytesTestVectorWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestVector& item) {
	if (!s.nat_write(0x3647c8ae)) { return false; }
	return tl2::details::CasesBytesTestVectorWrite(s, item);
}

void tl2::details::DictionaryStringReset(std::map<std::string, std::string>& item) noexcept {
	item.clear();
}

bool tl2::details::DictionaryStringWriteJSON(std::ostream& s, const std::map<std::string, std::string>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldStringWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryStringRead(::basictl::tl_istream & s, std::map<std::string, std::string>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldStringRead(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryStringWrite(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldStringWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryStringReadBoxed(::basictl::tl_istream & s, std::map<std::string, std::string>& item) {
	if (!s.nat_read_exact_tag(0x1f4c618f)) { return false; }
	return tl2::details::DictionaryStringRead(s, item);
}

bool tl2::details::DictionaryStringWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item) {
	if (!s.nat_write(0x1f4c618f)) { return false; }
	return tl2::details::DictionaryStringWrite(s, item);
}

void tl2::details::TupleString4Reset(std::array<std::string, 4>& item) noexcept {
	::tl2::details::BuiltinTuple4StringReset(item);
}

bool tl2::details::TupleString4WriteJSON(std::ostream& s, const std::array<std::string, 4>& item) noexcept {
	if (!::tl2::details::BuiltinTuple4StringWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleString4Read(::basictl::tl_istream & s, std::array<std::string, 4>& item) noexcept {
	if (!::tl2::details::BuiltinTuple4StringRead(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleString4Write(::basictl::tl_ostream & s, const std::array<std::string, 4>& item) noexcept {
	if (!::tl2::details::BuiltinTuple4StringWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleString4ReadBoxed(::basictl::tl_istream & s, std::array<std::string, 4>& item) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleString4Read(s, item);
}

bool tl2::details::TupleString4WriteBoxed(::basictl::tl_ostream & s, const std::array<std::string, 4>& item) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleString4Write(s, item);
}

void tl2::details::VectorDictionaryFieldStringReset(std::map<std::string, std::string>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorDictionaryFieldStringWriteJSON(std::ostream& s, const std::map<std::string, std::string>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldStringWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorDictionaryFieldStringRead(::basictl::tl_istream & s, std::map<std::string, std::string>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldStringRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorDictionaryFieldStringWrite(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldStringWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorDictionaryFieldStringReadBoxed(::basictl::tl_istream & s, std::map<std::string, std::string>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorDictionaryFieldStringRead(s, item);
}

bool tl2::details::VectorDictionaryFieldStringWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, std::string>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorDictionaryFieldStringWrite(s, item);
}

void tl2::details::VectorStringReset(std::vector<std::string>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorStringWriteJSON(std::ostream& s, const std::vector<std::string>& item) noexcept {
	if (!::tl2::details::BuiltinVectorStringWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorStringRead(::basictl::tl_istream & s, std::vector<std::string>& item) noexcept {
	if (!::tl2::details::BuiltinVectorStringRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item) noexcept {
	if (!::tl2::details::BuiltinVectorStringWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorStringReadBoxed(::basictl::tl_istream & s, std::vector<std::string>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorStringRead(s, item);
}

bool tl2::details::VectorStringWriteBoxed(::basictl::tl_ostream & s, const std::vector<std::string>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorStringWrite(s, item);
}
