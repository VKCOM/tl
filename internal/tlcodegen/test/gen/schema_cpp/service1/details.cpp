#include "service1/headers/service1_vector.h"
#include "service1/headers/service1.Value.h"
#include "service1/headers/service1.touch.h"
#include "service1/headers/service1.setOrIncr.h"
#include "service1/headers/service1.set.h"
#include "service1/headers/service1.replaceOrIncr.h"
#include "service1/headers/service1.replace.h"
#include "service1/headers/service1.keysStat.h"
#include "service1/headers/service1.incr.h"
#include "service1/headers/service1.getWildcardWithFlags.h"
#include "service1/headers/service1.getWildcardList.h"
#include "service1/headers/service1.getWildcardDict.h"
#include "service1/headers/service1.getWildcard.h"
#include "service1/headers/service1.getKeysStatPeriods.h"
#include "service1/headers/service1.getKeysStat.h"
#include "service1/headers/service1.getExpireTime.h"
#include "service1/headers/service1.get.h"
#include "service1/headers/service1.exists.h"
#include "service1/headers/service1.enableKeysStat.h"
#include "service1/headers/service1.enableExpiration.h"
#include "service1/headers/service1.disableKeysStat.h"
#include "service1/headers/service1.disableExpiration.h"
#include "service1/headers/service1.delete.h"
#include "service1/headers/service1.decr.h"
#include "service1/headers/service1.cas.h"
#include "service1/headers/service1.append.h"
#include "service1/headers/service1.addOrIncr.h"
#include "service1/headers/service1.addOrGet.h"
#include "service1/headers/service1.add.h"
#include "service1/headers/service1_map.h"
#include "__common_namespace/headers/int.h"
#include "service1/headers/service1.strvalueWithTime.h"
#include "service1/headers/service1.strvalue.h"
#include "service1/headers/service1.not_found.h"
#include "service1/headers/service1.longvalueWithTime.h"
#include "service1/headers/service1.longvalue.h"
#include "service1/headers/service1_dictionary.h"
#include "__common_namespace/headers/string.h"
#include "service1/headers/service1_dictionaryField.h"
#include "__common_namespace/headers/dictionaryField.h"
#include "__common_namespace/headers/Bool.h"


void tl2::details::BuiltinTuple3Service1ValueReset(std::array<::tl2::service1::Value, 3>& item) {
	for(auto && el : item) {
		::tl2::details::Service1ValueReset(el);
	}
}

bool tl2::details::BuiltinTuple3Service1ValueWriteJSON(std::ostream &s, const std::array<::tl2::service1::Value, 3>& item) {
	s << "[";
	size_t index = 0;
	for(auto && el : item) {
		if (!::tl2::details::Service1ValueWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTuple3Service1ValueRead(::basictl::tl_istream & s, std::array<::tl2::service1::Value, 3>& item) {
	for(auto && el : item) {
		if (!::tl2::details::Service1ValueReadBoxed(s, el)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

bool tl2::details::BuiltinTuple3Service1ValueWrite(::basictl::tl_ostream & s, const std::array<::tl2::service1::Value, 3>& item) {
	for(const auto & el : item) {
		if (!::tl2::details::Service1ValueWriteBoxed(s, el)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

void tl2::details::BuiltinVectorDictionaryFieldDictionaryIntReset(std::map<std::string, std::map<std::string, int32_t>>& item) {
	item.clear(); // TODO - unwrap
}

bool tl2::details::BuiltinVectorDictionaryFieldDictionaryIntWriteJSON(std::ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item) {
	s << "{";
	size_t index = 0;
	for(const auto & el : item) {
		s << "\"" << el.first << "\"";
		s << ":";
		if (!::tl2::details::BuiltinVectorDictionaryFieldIntWriteJSON(s, el.second)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "}";
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldDictionaryIntRead(::basictl::tl_istream & s, std::map<std::string, std::map<std::string, int32_t>>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	item.clear();
	for(uint32_t i = 0; i < len; i++) {
		std::string key;
		if (!s.string_read(key)) { return false; }
		if (!::tl2::details::BuiltinVectorDictionaryFieldIntRead(s, item[key])) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldDictionaryIntWrite(::basictl::tl_ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!s.string_write(el.first)) { return false;}
		if (!::tl2::details::BuiltinVectorDictionaryFieldIntWrite(s, el.second)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinVectorDictionaryFieldService1ValueReset(std::map<std::string, ::tl2::service1::Value>& item) {
	item.clear(); // TODO - unwrap
}

bool tl2::details::BuiltinVectorDictionaryFieldService1ValueWriteJSON(std::ostream & s, const std::map<std::string, ::tl2::service1::Value>& item) {
	s << "{";
	size_t index = 0;
	for(const auto & el : item) {
		s << "\"" << el.first << "\"";
		s << ":";
		if (!::tl2::details::Service1ValueWriteJSON(s, el.second)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "}";
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldService1ValueRead(::basictl::tl_istream & s, std::map<std::string, ::tl2::service1::Value>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	item.clear();
	for(uint32_t i = 0; i < len; i++) {
		std::string key;
		if (!s.string_read(key)) { return false; }
		if (!::tl2::details::Service1ValueReadBoxed(s, item[key])) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldService1ValueWrite(::basictl::tl_ostream & s, const std::map<std::string, ::tl2::service1::Value>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!s.string_write(el.first)) { return false;}
		if (!::tl2::details::Service1ValueWriteBoxed(s, el.second)) { return s.set_error_unknown_scenario(); }
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

void tl2::details::BuiltinVectorMapStringStringReset(std::vector<::tl2::Map<std::string, std::string>>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorMapStringStringWriteJSON(std::ostream & s, const std::vector<::tl2::Map<std::string, std::string>>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::MapStringStringWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorMapStringStringRead(::basictl::tl_istream & s, std::vector<::tl2::Map<std::string, std::string>>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::MapStringStringRead(s, el)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

bool tl2::details::BuiltinVectorMapStringStringWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Map<std::string, std::string>>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::MapStringStringWrite(s, el)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

void tl2::details::BuiltinVectorService1ValueReset(std::vector<::tl2::service1::Value>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorService1ValueWriteJSON(std::ostream & s, const std::vector<::tl2::service1::Value>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::Service1ValueWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorService1ValueRead(::basictl::tl_istream & s, std::vector<::tl2::service1::Value>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::Service1ValueReadBoxed(s, el)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

bool tl2::details::BuiltinVectorService1ValueWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service1::Value>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::Service1ValueWriteBoxed(s, el)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

void tl2::details::DictionaryDictionaryIntReset(std::map<std::string, std::map<std::string, int32_t>>& item) noexcept {
	item.clear();
}

bool tl2::details::DictionaryDictionaryIntWriteJSON(std::ostream& s, const std::map<std::string, std::map<std::string, int32_t>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldDictionaryIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryDictionaryIntRead(::basictl::tl_istream & s, std::map<std::string, std::map<std::string, int32_t>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldDictionaryIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryDictionaryIntWrite(::basictl::tl_ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldDictionaryIntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryDictionaryIntReadBoxed(::basictl::tl_istream & s, std::map<std::string, std::map<std::string, int32_t>>& item) {
	if (!s.nat_read_exact_tag(0x1f4c618f)) { return false; }
	return tl2::details::DictionaryDictionaryIntRead(s, item);
}

bool tl2::details::DictionaryDictionaryIntWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item) {
	if (!s.nat_write(0x1f4c618f)) { return false; }
	return tl2::details::DictionaryDictionaryIntWrite(s, item);
}

void tl2::details::MapStringStringReset(::tl2::Map<std::string, std::string>& item) noexcept {
	item.key.clear();
	item.value.clear();
}

bool tl2::details::MapStringStringWriteJSON(std::ostream& s, const ::tl2::Map<std::string, std::string>& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.value.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"value\":";
		s << "\"" << item.value << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::MapStringStringRead(::basictl::tl_istream & s, ::tl2::Map<std::string, std::string>& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.string_read(item.value)) { return false; }
	return true;
}

bool tl2::details::MapStringStringWrite(::basictl::tl_ostream & s, const ::tl2::Map<std::string, std::string>& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.string_write(item.value)) { return false;}
	return true;
}

bool tl2::details::MapStringStringReadBoxed(::basictl::tl_istream & s, ::tl2::Map<std::string, std::string>& item) {
	if (!s.nat_read_exact_tag(0x79c473a4)) { return false; }
	return tl2::details::MapStringStringRead(s, item);
}

bool tl2::details::MapStringStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Map<std::string, std::string>& item) {
	if (!s.nat_write(0x79c473a4)) { return false; }
	return tl2::details::MapStringStringWrite(s, item);
}

bool tl2::service1::Add::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1AddWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Add::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1AddRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Add::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1AddWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Add::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Add::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Add::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1AddReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Add::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1AddWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Add::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Add::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1AddReset(::tl2::service1::Add& item) noexcept {
	item.key.clear();
	item.flags = 0;
	item.delay = 0;
	item.value.clear();
}

bool tl2::details::Service1AddWriteJSON(std::ostream& s, const ::tl2::service1::Add& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.delay != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"delay\":";
		s << item.delay;
	}
	if (item.value.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"value\":";
		s << "\"" << item.value << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1AddRead(::basictl::tl_istream & s, ::tl2::service1::Add& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.delay)) { return false; }
	if (!s.string_read(item.value)) { return false; }
	return true;
}

bool tl2::details::Service1AddWrite(::basictl::tl_ostream & s, const ::tl2::service1::Add& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!s.int_write(item.delay)) { return false;}
	if (!s.string_write(item.value)) { return false;}
	return true;
}

bool tl2::details::Service1AddReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Add& item) {
	if (!s.nat_read_exact_tag(0x481df8be)) { return false; }
	return tl2::details::Service1AddRead(s, item);
}

bool tl2::details::Service1AddWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Add& item) {
	if (!s.nat_write(0x481df8be)) { return false; }
	return tl2::details::Service1AddWrite(s, item);
}

bool tl2::details::Service1AddReadResult(::basictl::tl_istream & s, tl2::service1::Add& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1AddWriteResult(::basictl::tl_ostream & s, tl2::service1::Add& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::Add::read_result(::basictl::tl_istream & s, bool & result) noexcept {
	bool success = tl2::details::Service1AddReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::Add::write_result(::basictl::tl_ostream & s, bool & result) noexcept {
	bool success = tl2::details::Service1AddWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::Add::read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::Add::write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::AddOrGet::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1AddOrGetWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::AddOrGet::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1AddOrGetRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::AddOrGet::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1AddOrGetWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::AddOrGet::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::AddOrGet::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::AddOrGet::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1AddOrGetReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::AddOrGet::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1AddOrGetWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::AddOrGet::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::AddOrGet::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1AddOrGetReset(::tl2::service1::AddOrGet& item) noexcept {
	item.key.clear();
	item.flags = 0;
	item.delay = 0;
	item.value.clear();
}

bool tl2::details::Service1AddOrGetWriteJSON(std::ostream& s, const ::tl2::service1::AddOrGet& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.delay != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"delay\":";
		s << item.delay;
	}
	if (item.value.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"value\":";
		s << "\"" << item.value << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1AddOrGetRead(::basictl::tl_istream & s, ::tl2::service1::AddOrGet& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.delay)) { return false; }
	if (!s.string_read(item.value)) { return false; }
	return true;
}

bool tl2::details::Service1AddOrGetWrite(::basictl::tl_ostream & s, const ::tl2::service1::AddOrGet& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!s.int_write(item.delay)) { return false;}
	if (!s.string_write(item.value)) { return false;}
	return true;
}

bool tl2::details::Service1AddOrGetReadBoxed(::basictl::tl_istream & s, ::tl2::service1::AddOrGet& item) {
	if (!s.nat_read_exact_tag(0x6a42faad)) { return false; }
	return tl2::details::Service1AddOrGetRead(s, item);
}

bool tl2::details::Service1AddOrGetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::AddOrGet& item) {
	if (!s.nat_write(0x6a42faad)) { return false; }
	return tl2::details::Service1AddOrGetWrite(s, item);
}

bool tl2::details::Service1AddOrGetReadResult(::basictl::tl_istream & s, tl2::service1::AddOrGet& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::Service1AddOrGetWriteResult(::basictl::tl_ostream & s, tl2::service1::AddOrGet& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::service1::AddOrGet::read_result(::basictl::tl_istream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1AddOrGetReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::AddOrGet::write_result(::basictl::tl_ostream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1AddOrGetWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::AddOrGet::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::service1::Value & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::AddOrGet::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::service1::Value & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::AddOrIncr::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1AddOrIncrWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::AddOrIncr::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1AddOrIncrRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::AddOrIncr::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1AddOrIncrWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::AddOrIncr::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::AddOrIncr::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::AddOrIncr::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1AddOrIncrReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::AddOrIncr::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1AddOrIncrWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::AddOrIncr::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::AddOrIncr::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1AddOrIncrReset(::tl2::service1::AddOrIncr& item) noexcept {
	item.key.clear();
	item.flags = 0;
	item.delay = 0;
	item.value = 0;
}

bool tl2::details::Service1AddOrIncrWriteJSON(std::ostream& s, const ::tl2::service1::AddOrIncr& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.delay != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"delay\":";
		s << item.delay;
	}
	if (item.value != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"value\":";
		s << item.value;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1AddOrIncrRead(::basictl::tl_istream & s, ::tl2::service1::AddOrIncr& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.delay)) { return false; }
	if (!s.long_read(item.value)) { return false; }
	return true;
}

bool tl2::details::Service1AddOrIncrWrite(::basictl::tl_ostream & s, const ::tl2::service1::AddOrIncr& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!s.int_write(item.delay)) { return false;}
	if (!s.long_write(item.value)) { return false;}
	return true;
}

bool tl2::details::Service1AddOrIncrReadBoxed(::basictl::tl_istream & s, ::tl2::service1::AddOrIncr& item) {
	if (!s.nat_read_exact_tag(0x90c4b402)) { return false; }
	return tl2::details::Service1AddOrIncrRead(s, item);
}

bool tl2::details::Service1AddOrIncrWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::AddOrIncr& item) {
	if (!s.nat_write(0x90c4b402)) { return false; }
	return tl2::details::Service1AddOrIncrWrite(s, item);
}

bool tl2::details::Service1AddOrIncrReadResult(::basictl::tl_istream & s, tl2::service1::AddOrIncr& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::Service1AddOrIncrWriteResult(::basictl::tl_ostream & s, tl2::service1::AddOrIncr& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::service1::AddOrIncr::read_result(::basictl::tl_istream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1AddOrIncrReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::AddOrIncr::write_result(::basictl::tl_ostream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1AddOrIncrWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::AddOrIncr::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::service1::Value & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::AddOrIncr::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::service1::Value & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::Append::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1AppendWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Append::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1AppendRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Append::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1AppendWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Append::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Append::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Append::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1AppendReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Append::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1AppendWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Append::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Append::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1AppendReset(::tl2::service1::Append& item) noexcept {
	item.key.clear();
	item.flags = 0;
	item.delay = 0;
	item.suffix.clear();
}

bool tl2::details::Service1AppendWriteJSON(std::ostream& s, const ::tl2::service1::Append& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.delay != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"delay\":";
		s << item.delay;
	}
	if (item.suffix.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"suffix\":";
		s << "\"" << item.suffix << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1AppendRead(::basictl::tl_istream & s, ::tl2::service1::Append& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.delay)) { return false; }
	if (!s.string_read(item.suffix)) { return false; }
	return true;
}

bool tl2::details::Service1AppendWrite(::basictl::tl_ostream & s, const ::tl2::service1::Append& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!s.int_write(item.delay)) { return false;}
	if (!s.string_write(item.suffix)) { return false;}
	return true;
}

bool tl2::details::Service1AppendReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Append& item) {
	if (!s.nat_read_exact_tag(0x04dec671)) { return false; }
	return tl2::details::Service1AppendRead(s, item);
}

bool tl2::details::Service1AppendWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Append& item) {
	if (!s.nat_write(0x04dec671)) { return false; }
	return tl2::details::Service1AppendWrite(s, item);
}

bool tl2::details::Service1AppendReadResult(::basictl::tl_istream & s, tl2::service1::Append& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1AppendWriteResult(::basictl::tl_ostream & s, tl2::service1::Append& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::Append::read_result(::basictl::tl_istream & s, bool & result) noexcept {
	bool success = tl2::details::Service1AppendReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::Append::write_result(::basictl::tl_ostream & s, bool & result) noexcept {
	bool success = tl2::details::Service1AppendWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::Append::read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::Append::write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::Cas::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1CasWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Cas::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1CasRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Cas::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1CasWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Cas::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Cas::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Cas::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1CasReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Cas::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1CasWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Cas::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Cas::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1CasReset(::tl2::service1::Cas& item) noexcept {
	item.key.clear();
	item.flags = 0;
	item.delay = 0;
	item.casToken.clear();
	item.newValue.clear();
}

bool tl2::details::Service1CasWriteJSON(std::ostream& s, const ::tl2::service1::Cas& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.delay != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"delay\":";
		s << item.delay;
	}
	if (item.casToken.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"casToken\":";
		s << "\"" << item.casToken << "\"";
	}
	if (item.newValue.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"newValue\":";
		s << "\"" << item.newValue << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1CasRead(::basictl::tl_istream & s, ::tl2::service1::Cas& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.delay)) { return false; }
	if (!s.string_read(item.casToken)) { return false; }
	if (!s.string_read(item.newValue)) { return false; }
	return true;
}

bool tl2::details::Service1CasWrite(::basictl::tl_ostream & s, const ::tl2::service1::Cas& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!s.int_write(item.delay)) { return false;}
	if (!s.string_write(item.casToken)) { return false;}
	if (!s.string_write(item.newValue)) { return false;}
	return true;
}

bool tl2::details::Service1CasReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Cas& item) {
	if (!s.nat_read_exact_tag(0x51851964)) { return false; }
	return tl2::details::Service1CasRead(s, item);
}

bool tl2::details::Service1CasWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Cas& item) {
	if (!s.nat_write(0x51851964)) { return false; }
	return tl2::details::Service1CasWrite(s, item);
}

bool tl2::details::Service1CasReadResult(::basictl::tl_istream & s, tl2::service1::Cas& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1CasWriteResult(::basictl::tl_ostream & s, tl2::service1::Cas& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::Cas::read_result(::basictl::tl_istream & s, bool & result) noexcept {
	bool success = tl2::details::Service1CasReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::Cas::write_result(::basictl::tl_ostream & s, bool & result) noexcept {
	bool success = tl2::details::Service1CasWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::Cas::read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::Cas::write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::Decr::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1DecrWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Decr::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1DecrRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Decr::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1DecrWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Decr::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Decr::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Decr::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1DecrReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Decr::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1DecrWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Decr::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Decr::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1DecrReset(::tl2::service1::Decr& item) noexcept {
	item.key.clear();
	item.value = 0;
}

bool tl2::details::Service1DecrWriteJSON(std::ostream& s, const ::tl2::service1::Decr& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.value != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"value\":";
		s << item.value;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1DecrRead(::basictl::tl_istream & s, ::tl2::service1::Decr& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.long_read(item.value)) { return false; }
	return true;
}

bool tl2::details::Service1DecrWrite(::basictl::tl_ostream & s, const ::tl2::service1::Decr& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.long_write(item.value)) { return false;}
	return true;
}

bool tl2::details::Service1DecrReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Decr& item) {
	if (!s.nat_read_exact_tag(0xeb179ce7)) { return false; }
	return tl2::details::Service1DecrRead(s, item);
}

bool tl2::details::Service1DecrWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Decr& item) {
	if (!s.nat_write(0xeb179ce7)) { return false; }
	return tl2::details::Service1DecrWrite(s, item);
}

bool tl2::details::Service1DecrReadResult(::basictl::tl_istream & s, tl2::service1::Decr& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::Service1DecrWriteResult(::basictl::tl_ostream & s, tl2::service1::Decr& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::service1::Decr::read_result(::basictl::tl_istream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1DecrReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::Decr::write_result(::basictl::tl_ostream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1DecrWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::Decr::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::service1::Value & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::Decr::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::service1::Value & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::Delete::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1DeleteWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Delete::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1DeleteRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Delete::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1DeleteWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Delete::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Delete::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Delete::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1DeleteReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Delete::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1DeleteWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Delete::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Delete::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1DeleteReset(::tl2::service1::Delete& item) noexcept {
	item.key.clear();
}

bool tl2::details::Service1DeleteWriteJSON(std::ostream& s, const ::tl2::service1::Delete& item) noexcept {
	s << "{";
	if (item.key.size() != 0) {
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1DeleteRead(::basictl::tl_istream & s, ::tl2::service1::Delete& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	return true;
}

bool tl2::details::Service1DeleteWrite(::basictl::tl_ostream & s, const ::tl2::service1::Delete& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	return true;
}

bool tl2::details::Service1DeleteReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Delete& item) {
	if (!s.nat_read_exact_tag(0x83277767)) { return false; }
	return tl2::details::Service1DeleteRead(s, item);
}

bool tl2::details::Service1DeleteWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Delete& item) {
	if (!s.nat_write(0x83277767)) { return false; }
	return tl2::details::Service1DeleteWrite(s, item);
}

bool tl2::details::Service1DeleteReadResult(::basictl::tl_istream & s, tl2::service1::Delete& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1DeleteWriteResult(::basictl::tl_ostream & s, tl2::service1::Delete& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::Delete::read_result(::basictl::tl_istream & s, bool & result) noexcept {
	bool success = tl2::details::Service1DeleteReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::Delete::write_result(::basictl::tl_ostream & s, bool & result) noexcept {
	bool success = tl2::details::Service1DeleteWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::Delete::read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::Delete::write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::DisableExpiration::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1DisableExpirationWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::DisableExpiration::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1DisableExpirationRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::DisableExpiration::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1DisableExpirationWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::DisableExpiration::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::DisableExpiration::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::DisableExpiration::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1DisableExpirationReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::DisableExpiration::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1DisableExpirationWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::DisableExpiration::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::DisableExpiration::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1DisableExpirationReset(::tl2::service1::DisableExpiration& item) noexcept {
	item.prefix.clear();
}

bool tl2::details::Service1DisableExpirationWriteJSON(std::ostream& s, const ::tl2::service1::DisableExpiration& item) noexcept {
	s << "{";
	if (item.prefix.size() != 0) {
		s << "\"prefix\":";
		s << "\"" << item.prefix << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1DisableExpirationRead(::basictl::tl_istream & s, ::tl2::service1::DisableExpiration& item) noexcept {
	if (!s.string_read(item.prefix)) { return false; }
	return true;
}

bool tl2::details::Service1DisableExpirationWrite(::basictl::tl_ostream & s, const ::tl2::service1::DisableExpiration& item) noexcept {
	if (!s.string_write(item.prefix)) { return false;}
	return true;
}

bool tl2::details::Service1DisableExpirationReadBoxed(::basictl::tl_istream & s, ::tl2::service1::DisableExpiration& item) {
	if (!s.nat_read_exact_tag(0xf1c39c2d)) { return false; }
	return tl2::details::Service1DisableExpirationRead(s, item);
}

bool tl2::details::Service1DisableExpirationWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::DisableExpiration& item) {
	if (!s.nat_write(0xf1c39c2d)) { return false; }
	return tl2::details::Service1DisableExpirationWrite(s, item);
}

bool tl2::details::Service1DisableExpirationReadResult(::basictl::tl_istream & s, tl2::service1::DisableExpiration& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1DisableExpirationWriteResult(::basictl::tl_ostream & s, tl2::service1::DisableExpiration& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::DisableExpiration::read_result(::basictl::tl_istream & s, bool & result) noexcept {
	bool success = tl2::details::Service1DisableExpirationReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::DisableExpiration::write_result(::basictl::tl_ostream & s, bool & result) noexcept {
	bool success = tl2::details::Service1DisableExpirationWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::DisableExpiration::read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::DisableExpiration::write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::DisableKeysStat::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1DisableKeysStatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::DisableKeysStat::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1DisableKeysStatRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::DisableKeysStat::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1DisableKeysStatWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::DisableKeysStat::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::DisableKeysStat::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::DisableKeysStat::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1DisableKeysStatReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::DisableKeysStat::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1DisableKeysStatWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::DisableKeysStat::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::DisableKeysStat::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1DisableKeysStatReset(::tl2::service1::DisableKeysStat& item) noexcept {
	item.period = 0;
}

bool tl2::details::Service1DisableKeysStatWriteJSON(std::ostream& s, const ::tl2::service1::DisableKeysStat& item) noexcept {
	s << "{";
	if (item.period != 0) {
		s << "\"period\":";
		s << item.period;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1DisableKeysStatRead(::basictl::tl_istream & s, ::tl2::service1::DisableKeysStat& item) noexcept {
	if (!s.int_read(item.period)) { return false; }
	return true;
}

bool tl2::details::Service1DisableKeysStatWrite(::basictl::tl_ostream & s, const ::tl2::service1::DisableKeysStat& item) noexcept {
	if (!s.int_write(item.period)) { return false;}
	return true;
}

bool tl2::details::Service1DisableKeysStatReadBoxed(::basictl::tl_istream & s, ::tl2::service1::DisableKeysStat& item) {
	if (!s.nat_read_exact_tag(0x79d6160f)) { return false; }
	return tl2::details::Service1DisableKeysStatRead(s, item);
}

bool tl2::details::Service1DisableKeysStatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::DisableKeysStat& item) {
	if (!s.nat_write(0x79d6160f)) { return false; }
	return tl2::details::Service1DisableKeysStatWrite(s, item);
}

bool tl2::details::Service1DisableKeysStatReadResult(::basictl::tl_istream & s, tl2::service1::DisableKeysStat& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1DisableKeysStatWriteResult(::basictl::tl_ostream & s, tl2::service1::DisableKeysStat& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::DisableKeysStat::read_result(::basictl::tl_istream & s, bool & result) noexcept {
	bool success = tl2::details::Service1DisableKeysStatReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::DisableKeysStat::write_result(::basictl::tl_ostream & s, bool & result) noexcept {
	bool success = tl2::details::Service1DisableKeysStatWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::DisableKeysStat::read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::DisableKeysStat::write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::EnableExpiration::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1EnableExpirationWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::EnableExpiration::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1EnableExpirationRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::EnableExpiration::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1EnableExpirationWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::EnableExpiration::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::EnableExpiration::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::EnableExpiration::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1EnableExpirationReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::EnableExpiration::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1EnableExpirationWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::EnableExpiration::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::EnableExpiration::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1EnableExpirationReset(::tl2::service1::EnableExpiration& item) noexcept {
	item.prefix.clear();
}

bool tl2::details::Service1EnableExpirationWriteJSON(std::ostream& s, const ::tl2::service1::EnableExpiration& item) noexcept {
	s << "{";
	if (item.prefix.size() != 0) {
		s << "\"prefix\":";
		s << "\"" << item.prefix << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1EnableExpirationRead(::basictl::tl_istream & s, ::tl2::service1::EnableExpiration& item) noexcept {
	if (!s.string_read(item.prefix)) { return false; }
	return true;
}

bool tl2::details::Service1EnableExpirationWrite(::basictl::tl_ostream & s, const ::tl2::service1::EnableExpiration& item) noexcept {
	if (!s.string_write(item.prefix)) { return false;}
	return true;
}

bool tl2::details::Service1EnableExpirationReadBoxed(::basictl::tl_istream & s, ::tl2::service1::EnableExpiration& item) {
	if (!s.nat_read_exact_tag(0x2b51ad67)) { return false; }
	return tl2::details::Service1EnableExpirationRead(s, item);
}

bool tl2::details::Service1EnableExpirationWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::EnableExpiration& item) {
	if (!s.nat_write(0x2b51ad67)) { return false; }
	return tl2::details::Service1EnableExpirationWrite(s, item);
}

bool tl2::details::Service1EnableExpirationReadResult(::basictl::tl_istream & s, tl2::service1::EnableExpiration& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1EnableExpirationWriteResult(::basictl::tl_ostream & s, tl2::service1::EnableExpiration& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::EnableExpiration::read_result(::basictl::tl_istream & s, bool & result) noexcept {
	bool success = tl2::details::Service1EnableExpirationReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::EnableExpiration::write_result(::basictl::tl_ostream & s, bool & result) noexcept {
	bool success = tl2::details::Service1EnableExpirationWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::EnableExpiration::read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::EnableExpiration::write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::EnableKeysStat::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1EnableKeysStatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::EnableKeysStat::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1EnableKeysStatRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::EnableKeysStat::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1EnableKeysStatWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::EnableKeysStat::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::EnableKeysStat::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::EnableKeysStat::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1EnableKeysStatReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::EnableKeysStat::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1EnableKeysStatWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::EnableKeysStat::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::EnableKeysStat::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1EnableKeysStatReset(::tl2::service1::EnableKeysStat& item) noexcept {
	item.period = 0;
}

bool tl2::details::Service1EnableKeysStatWriteJSON(std::ostream& s, const ::tl2::service1::EnableKeysStat& item) noexcept {
	s << "{";
	if (item.period != 0) {
		s << "\"period\":";
		s << item.period;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1EnableKeysStatRead(::basictl::tl_istream & s, ::tl2::service1::EnableKeysStat& item) noexcept {
	if (!s.int_read(item.period)) { return false; }
	return true;
}

bool tl2::details::Service1EnableKeysStatWrite(::basictl::tl_ostream & s, const ::tl2::service1::EnableKeysStat& item) noexcept {
	if (!s.int_write(item.period)) { return false;}
	return true;
}

bool tl2::details::Service1EnableKeysStatReadBoxed(::basictl::tl_istream & s, ::tl2::service1::EnableKeysStat& item) {
	if (!s.nat_read_exact_tag(0x29a7090e)) { return false; }
	return tl2::details::Service1EnableKeysStatRead(s, item);
}

bool tl2::details::Service1EnableKeysStatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::EnableKeysStat& item) {
	if (!s.nat_write(0x29a7090e)) { return false; }
	return tl2::details::Service1EnableKeysStatWrite(s, item);
}

bool tl2::details::Service1EnableKeysStatReadResult(::basictl::tl_istream & s, tl2::service1::EnableKeysStat& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1EnableKeysStatWriteResult(::basictl::tl_ostream & s, tl2::service1::EnableKeysStat& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::EnableKeysStat::read_result(::basictl::tl_istream & s, bool & result) noexcept {
	bool success = tl2::details::Service1EnableKeysStatReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::EnableKeysStat::write_result(::basictl::tl_ostream & s, bool & result) noexcept {
	bool success = tl2::details::Service1EnableKeysStatWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::EnableKeysStat::read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::EnableKeysStat::write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::Exists::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1ExistsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Exists::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1ExistsRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Exists::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1ExistsWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Exists::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Exists::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Exists::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1ExistsReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Exists::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1ExistsWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Exists::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Exists::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1ExistsReset(::tl2::service1::Exists& item) noexcept {
	item.key.clear();
}

bool tl2::details::Service1ExistsWriteJSON(std::ostream& s, const ::tl2::service1::Exists& item) noexcept {
	s << "{";
	if (item.key.size() != 0) {
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1ExistsRead(::basictl::tl_istream & s, ::tl2::service1::Exists& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	return true;
}

bool tl2::details::Service1ExistsWrite(::basictl::tl_ostream & s, const ::tl2::service1::Exists& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	return true;
}

bool tl2::details::Service1ExistsReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Exists& item) {
	if (!s.nat_read_exact_tag(0xe0284c9e)) { return false; }
	return tl2::details::Service1ExistsRead(s, item);
}

bool tl2::details::Service1ExistsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Exists& item) {
	if (!s.nat_write(0xe0284c9e)) { return false; }
	return tl2::details::Service1ExistsWrite(s, item);
}

bool tl2::details::Service1ExistsReadResult(::basictl::tl_istream & s, tl2::service1::Exists& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1ExistsWriteResult(::basictl::tl_ostream & s, tl2::service1::Exists& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::Exists::read_result(::basictl::tl_istream & s, bool & result) noexcept {
	bool success = tl2::details::Service1ExistsReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::Exists::write_result(::basictl::tl_ostream & s, bool & result) noexcept {
	bool success = tl2::details::Service1ExistsWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::Exists::read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::Exists::write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::Get::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1GetWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Get::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Get::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Get::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Get::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Get::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Get::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Get::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Get::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1GetReset(::tl2::service1::Get& item) noexcept {
	item.key.clear();
}

bool tl2::details::Service1GetWriteJSON(std::ostream& s, const ::tl2::service1::Get& item) noexcept {
	s << "{";
	if (item.key.size() != 0) {
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1GetRead(::basictl::tl_istream & s, ::tl2::service1::Get& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	return true;
}

bool tl2::details::Service1GetWrite(::basictl::tl_ostream & s, const ::tl2::service1::Get& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	return true;
}

bool tl2::details::Service1GetReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Get& item) {
	if (!s.nat_read_exact_tag(0x29099b19)) { return false; }
	return tl2::details::Service1GetRead(s, item);
}

bool tl2::details::Service1GetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Get& item) {
	if (!s.nat_write(0x29099b19)) { return false; }
	return tl2::details::Service1GetWrite(s, item);
}

bool tl2::details::Service1GetReadResult(::basictl::tl_istream & s, tl2::service1::Get& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::Service1GetWriteResult(::basictl::tl_ostream & s, tl2::service1::Get& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::service1::Get::read_result(::basictl::tl_istream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1GetReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::Get::write_result(::basictl::tl_ostream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1GetWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::Get::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::service1::Value & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::Get::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::service1::Value & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::GetExpireTime::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1GetExpireTimeWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::GetExpireTime::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetExpireTimeRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetExpireTime::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetExpireTimeWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetExpireTime::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::GetExpireTime::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::GetExpireTime::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetExpireTimeReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetExpireTime::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetExpireTimeWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetExpireTime::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::GetExpireTime::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1GetExpireTimeReset(::tl2::service1::GetExpireTime& item) noexcept {
	item.key.clear();
}

bool tl2::details::Service1GetExpireTimeWriteJSON(std::ostream& s, const ::tl2::service1::GetExpireTime& item) noexcept {
	s << "{";
	if (item.key.size() != 0) {
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1GetExpireTimeRead(::basictl::tl_istream & s, ::tl2::service1::GetExpireTime& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	return true;
}

bool tl2::details::Service1GetExpireTimeWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetExpireTime& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	return true;
}

bool tl2::details::Service1GetExpireTimeReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetExpireTime& item) {
	if (!s.nat_read_exact_tag(0x5a731070)) { return false; }
	return tl2::details::Service1GetExpireTimeRead(s, item);
}

bool tl2::details::Service1GetExpireTimeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetExpireTime& item) {
	if (!s.nat_write(0x5a731070)) { return false; }
	return tl2::details::Service1GetExpireTimeWrite(s, item);
}

bool tl2::details::Service1GetExpireTimeReadResult(::basictl::tl_istream & s, tl2::service1::GetExpireTime& item, std::optional<int32_t>& result) {
	if (!::tl2::details::IntMaybeReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1GetExpireTimeWriteResult(::basictl::tl_ostream & s, tl2::service1::GetExpireTime& item, std::optional<int32_t>& result) {
	if (!::tl2::details::IntMaybeWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::GetExpireTime::read_result(::basictl::tl_istream & s, std::optional<int32_t> & result) noexcept {
	bool success = tl2::details::Service1GetExpireTimeReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::GetExpireTime::write_result(::basictl::tl_ostream & s, std::optional<int32_t> & result) noexcept {
	bool success = tl2::details::Service1GetExpireTimeWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::GetExpireTime::read_result_or_throw(::basictl::tl_throwable_istream & s, std::optional<int32_t> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::GetExpireTime::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::optional<int32_t> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::GetKeysStat::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1GetKeysStatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::GetKeysStat::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetKeysStatRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetKeysStat::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetKeysStatWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetKeysStat::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::GetKeysStat::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::GetKeysStat::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetKeysStatReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetKeysStat::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetKeysStatWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetKeysStat::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::GetKeysStat::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1GetKeysStatReset(::tl2::service1::GetKeysStat& item) noexcept {
	item.period = 0;
}

bool tl2::details::Service1GetKeysStatWriteJSON(std::ostream& s, const ::tl2::service1::GetKeysStat& item) noexcept {
	s << "{";
	if (item.period != 0) {
		s << "\"period\":";
		s << item.period;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1GetKeysStatRead(::basictl::tl_istream & s, ::tl2::service1::GetKeysStat& item) noexcept {
	if (!s.int_read(item.period)) { return false; }
	return true;
}

bool tl2::details::Service1GetKeysStatWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetKeysStat& item) noexcept {
	if (!s.int_write(item.period)) { return false;}
	return true;
}

bool tl2::details::Service1GetKeysStatReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetKeysStat& item) {
	if (!s.nat_read_exact_tag(0x06cecd58)) { return false; }
	return tl2::details::Service1GetKeysStatRead(s, item);
}

bool tl2::details::Service1GetKeysStatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetKeysStat& item) {
	if (!s.nat_write(0x06cecd58)) { return false; }
	return tl2::details::Service1GetKeysStatWrite(s, item);
}

bool tl2::details::Service1GetKeysStatReadResult(::basictl::tl_istream & s, tl2::service1::GetKeysStat& item, std::optional<::tl2::service1::KeysStat>& result) {
	if (!::tl2::details::Service1KeysStatMaybeReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1GetKeysStatWriteResult(::basictl::tl_ostream & s, tl2::service1::GetKeysStat& item, std::optional<::tl2::service1::KeysStat>& result) {
	if (!::tl2::details::Service1KeysStatMaybeWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::GetKeysStat::read_result(::basictl::tl_istream & s, std::optional<::tl2::service1::KeysStat> & result) noexcept {
	bool success = tl2::details::Service1GetKeysStatReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::GetKeysStat::write_result(::basictl::tl_ostream & s, std::optional<::tl2::service1::KeysStat> & result) noexcept {
	bool success = tl2::details::Service1GetKeysStatWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::GetKeysStat::read_result_or_throw(::basictl::tl_throwable_istream & s, std::optional<::tl2::service1::KeysStat> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::GetKeysStat::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::optional<::tl2::service1::KeysStat> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::GetKeysStatPeriods::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1GetKeysStatPeriodsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::GetKeysStatPeriods::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetKeysStatPeriodsRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetKeysStatPeriods::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetKeysStatPeriodsWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetKeysStatPeriods::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::GetKeysStatPeriods::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::GetKeysStatPeriods::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetKeysStatPeriodsReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetKeysStatPeriods::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetKeysStatPeriodsWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetKeysStatPeriods::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::GetKeysStatPeriods::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1GetKeysStatPeriodsReset(::tl2::service1::GetKeysStatPeriods& item) noexcept {
}

bool tl2::details::Service1GetKeysStatPeriodsWriteJSON(std::ostream& s, const ::tl2::service1::GetKeysStatPeriods& item) noexcept {
	s << "true";
	return true;
}

bool tl2::details::Service1GetKeysStatPeriodsRead(::basictl::tl_istream & s, ::tl2::service1::GetKeysStatPeriods& item) noexcept {
	return true;
}

bool tl2::details::Service1GetKeysStatPeriodsWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetKeysStatPeriods& item) noexcept {
	return true;
}

bool tl2::details::Service1GetKeysStatPeriodsReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetKeysStatPeriods& item) {
	if (!s.nat_read_exact_tag(0x8cdf39e3)) { return false; }
	return tl2::details::Service1GetKeysStatPeriodsRead(s, item);
}

bool tl2::details::Service1GetKeysStatPeriodsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetKeysStatPeriods& item) {
	if (!s.nat_write(0x8cdf39e3)) { return false; }
	return tl2::details::Service1GetKeysStatPeriodsWrite(s, item);
}

bool tl2::details::Service1GetKeysStatPeriodsReadResult(::basictl::tl_istream & s, tl2::service1::GetKeysStatPeriods& item, std::vector<int32_t>& result) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1GetKeysStatPeriodsWriteResult(::basictl::tl_ostream & s, tl2::service1::GetKeysStatPeriods& item, std::vector<int32_t>& result) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntWrite(s, result)) { return false; }
	return true;
}

bool tl2::service1::GetKeysStatPeriods::read_result(::basictl::tl_istream & s, std::vector<int32_t> & result) noexcept {
	bool success = tl2::details::Service1GetKeysStatPeriodsReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::GetKeysStatPeriods::write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result) noexcept {
	bool success = tl2::details::Service1GetKeysStatPeriodsWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::GetKeysStatPeriods::read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<int32_t> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::GetKeysStatPeriods::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<int32_t> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::GetWildcard::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1GetWildcardWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::GetWildcard::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetWildcardRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetWildcard::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetWildcardWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetWildcard::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::GetWildcard::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::GetWildcard::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetWildcardReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetWildcard::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetWildcardWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetWildcard::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::GetWildcard::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1GetWildcardReset(::tl2::service1::GetWildcard& item) noexcept {
	item.prefix.clear();
}

bool tl2::details::Service1GetWildcardWriteJSON(std::ostream& s, const ::tl2::service1::GetWildcard& item) noexcept {
	s << "{";
	if (item.prefix.size() != 0) {
		s << "\"prefix\":";
		s << "\"" << item.prefix << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1GetWildcardRead(::basictl::tl_istream & s, ::tl2::service1::GetWildcard& item) noexcept {
	if (!s.string_read(item.prefix)) { return false; }
	return true;
}

bool tl2::details::Service1GetWildcardWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcard& item) noexcept {
	if (!s.string_write(item.prefix)) { return false;}
	return true;
}

bool tl2::details::Service1GetWildcardReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetWildcard& item) {
	if (!s.nat_read_exact_tag(0x2f2abf13)) { return false; }
	return tl2::details::Service1GetWildcardRead(s, item);
}

bool tl2::details::Service1GetWildcardWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcard& item) {
	if (!s.nat_write(0x2f2abf13)) { return false; }
	return tl2::details::Service1GetWildcardWrite(s, item);
}

bool tl2::details::Service1GetWildcardReadResult(::basictl::tl_istream & s, tl2::service1::GetWildcard& item, std::vector<::tl2::Map<std::string, std::string>>& result) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorMapStringStringRead(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1GetWildcardWriteResult(::basictl::tl_ostream & s, tl2::service1::GetWildcard& item, std::vector<::tl2::Map<std::string, std::string>>& result) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorMapStringStringWrite(s, result)) { return false; }
	return true;
}

bool tl2::service1::GetWildcard::read_result(::basictl::tl_istream & s, std::vector<::tl2::Map<std::string, std::string>> & result) noexcept {
	bool success = tl2::details::Service1GetWildcardReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::GetWildcard::write_result(::basictl::tl_ostream & s, std::vector<::tl2::Map<std::string, std::string>> & result) noexcept {
	bool success = tl2::details::Service1GetWildcardWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::GetWildcard::read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<::tl2::Map<std::string, std::string>> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::GetWildcard::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<::tl2::Map<std::string, std::string>> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::GetWildcardDict::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1GetWildcardDictWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::GetWildcardDict::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetWildcardDictRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetWildcardDict::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetWildcardDictWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetWildcardDict::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::GetWildcardDict::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::GetWildcardDict::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetWildcardDictReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetWildcardDict::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetWildcardDictWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetWildcardDict::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::GetWildcardDict::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1GetWildcardDictReset(::tl2::service1::GetWildcardDict& item) noexcept {
	item.prefix.clear();
}

bool tl2::details::Service1GetWildcardDictWriteJSON(std::ostream& s, const ::tl2::service1::GetWildcardDict& item) noexcept {
	s << "{";
	if (item.prefix.size() != 0) {
		s << "\"prefix\":";
		s << "\"" << item.prefix << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1GetWildcardDictRead(::basictl::tl_istream & s, ::tl2::service1::GetWildcardDict& item) noexcept {
	if (!s.string_read(item.prefix)) { return false; }
	return true;
}

bool tl2::details::Service1GetWildcardDictWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardDict& item) noexcept {
	if (!s.string_write(item.prefix)) { return false;}
	return true;
}

bool tl2::details::Service1GetWildcardDictReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetWildcardDict& item) {
	if (!s.nat_read_exact_tag(0x72bbc81b)) { return false; }
	return tl2::details::Service1GetWildcardDictRead(s, item);
}

bool tl2::details::Service1GetWildcardDictWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardDict& item) {
	if (!s.nat_write(0x72bbc81b)) { return false; }
	return tl2::details::Service1GetWildcardDictWrite(s, item);
}

bool tl2::details::Service1GetWildcardDictReadResult(::basictl::tl_istream & s, tl2::service1::GetWildcardDict& item, std::map<std::string, std::string>& result) {
	if (!s.nat_read_exact_tag(0x1f4c618f)) { return false; }
	if (!::tl2::details::BuiltinVectorDictionaryFieldStringRead(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1GetWildcardDictWriteResult(::basictl::tl_ostream & s, tl2::service1::GetWildcardDict& item, std::map<std::string, std::string>& result) {
	if (!s.nat_write(0x1f4c618f)) { return false; }
	if (!::tl2::details::BuiltinVectorDictionaryFieldStringWrite(s, result)) { return false; }
	return true;
}

bool tl2::service1::GetWildcardDict::read_result(::basictl::tl_istream & s, std::map<std::string, std::string> & result) noexcept {
	bool success = tl2::details::Service1GetWildcardDictReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::GetWildcardDict::write_result(::basictl::tl_ostream & s, std::map<std::string, std::string> & result) noexcept {
	bool success = tl2::details::Service1GetWildcardDictWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::GetWildcardDict::read_result_or_throw(::basictl::tl_throwable_istream & s, std::map<std::string, std::string> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::GetWildcardDict::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::map<std::string, std::string> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::GetWildcardList::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1GetWildcardListWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::GetWildcardList::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetWildcardListRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetWildcardList::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetWildcardListWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetWildcardList::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::GetWildcardList::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::GetWildcardList::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetWildcardListReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetWildcardList::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetWildcardListWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetWildcardList::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::GetWildcardList::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1GetWildcardListReset(::tl2::service1::GetWildcardList& item) noexcept {
	item.prefix.clear();
}

bool tl2::details::Service1GetWildcardListWriteJSON(std::ostream& s, const ::tl2::service1::GetWildcardList& item) noexcept {
	s << "{";
	if (item.prefix.size() != 0) {
		s << "\"prefix\":";
		s << "\"" << item.prefix << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1GetWildcardListRead(::basictl::tl_istream & s, ::tl2::service1::GetWildcardList& item) noexcept {
	if (!s.string_read(item.prefix)) { return false; }
	return true;
}

bool tl2::details::Service1GetWildcardListWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardList& item) noexcept {
	if (!s.string_write(item.prefix)) { return false;}
	return true;
}

bool tl2::details::Service1GetWildcardListReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetWildcardList& item) {
	if (!s.nat_read_exact_tag(0x56b6ead4)) { return false; }
	return tl2::details::Service1GetWildcardListRead(s, item);
}

bool tl2::details::Service1GetWildcardListWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardList& item) {
	if (!s.nat_write(0x56b6ead4)) { return false; }
	return tl2::details::Service1GetWildcardListWrite(s, item);
}

bool tl2::details::Service1GetWildcardListReadResult(::basictl::tl_istream & s, tl2::service1::GetWildcardList& item, std::vector<std::string>& result) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorStringRead(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1GetWildcardListWriteResult(::basictl::tl_ostream & s, tl2::service1::GetWildcardList& item, std::vector<std::string>& result) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorStringWrite(s, result)) { return false; }
	return true;
}

bool tl2::service1::GetWildcardList::read_result(::basictl::tl_istream & s, std::vector<std::string> & result) noexcept {
	bool success = tl2::details::Service1GetWildcardListReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::GetWildcardList::write_result(::basictl::tl_ostream & s, std::vector<std::string> & result) noexcept {
	bool success = tl2::details::Service1GetWildcardListWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::GetWildcardList::read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<std::string> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::GetWildcardList::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<std::string> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::GetWildcardWithFlags::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1GetWildcardWithFlagsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::GetWildcardWithFlags::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetWildcardWithFlagsRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetWildcardWithFlags::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetWildcardWithFlagsWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetWildcardWithFlags::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::GetWildcardWithFlags::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::GetWildcardWithFlags::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1GetWildcardWithFlagsReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::GetWildcardWithFlags::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1GetWildcardWithFlagsWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::GetWildcardWithFlags::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::GetWildcardWithFlags::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1GetWildcardWithFlagsReset(::tl2::service1::GetWildcardWithFlags& item) noexcept {
	item.prefix.clear();
}

bool tl2::details::Service1GetWildcardWithFlagsWriteJSON(std::ostream& s, const ::tl2::service1::GetWildcardWithFlags& item) noexcept {
	s << "{";
	if (item.prefix.size() != 0) {
		s << "\"prefix\":";
		s << "\"" << item.prefix << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1GetWildcardWithFlagsRead(::basictl::tl_istream & s, ::tl2::service1::GetWildcardWithFlags& item) noexcept {
	if (!s.string_read(item.prefix)) { return false; }
	return true;
}

bool tl2::details::Service1GetWildcardWithFlagsWrite(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardWithFlags& item) noexcept {
	if (!s.string_write(item.prefix)) { return false;}
	return true;
}

bool tl2::details::Service1GetWildcardWithFlagsReadBoxed(::basictl::tl_istream & s, ::tl2::service1::GetWildcardWithFlags& item) {
	if (!s.nat_read_exact_tag(0x5f6a1f78)) { return false; }
	return tl2::details::Service1GetWildcardWithFlagsRead(s, item);
}

bool tl2::details::Service1GetWildcardWithFlagsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::GetWildcardWithFlags& item) {
	if (!s.nat_write(0x5f6a1f78)) { return false; }
	return tl2::details::Service1GetWildcardWithFlagsWrite(s, item);
}

bool tl2::details::Service1GetWildcardWithFlagsReadResult(::basictl::tl_istream & s, tl2::service1::GetWildcardWithFlags& item, std::map<std::string, ::tl2::service1::Value>& result) {
	if (!s.nat_read_exact_tag(0x1f4c618f)) { return false; }
	if (!::tl2::details::BuiltinVectorDictionaryFieldService1ValueRead(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1GetWildcardWithFlagsWriteResult(::basictl::tl_ostream & s, tl2::service1::GetWildcardWithFlags& item, std::map<std::string, ::tl2::service1::Value>& result) {
	if (!s.nat_write(0x1f4c618f)) { return false; }
	if (!::tl2::details::BuiltinVectorDictionaryFieldService1ValueWrite(s, result)) { return false; }
	return true;
}

bool tl2::service1::GetWildcardWithFlags::read_result(::basictl::tl_istream & s, std::map<std::string, ::tl2::service1::Value> & result) noexcept {
	bool success = tl2::details::Service1GetWildcardWithFlagsReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::GetWildcardWithFlags::write_result(::basictl::tl_ostream & s, std::map<std::string, ::tl2::service1::Value> & result) noexcept {
	bool success = tl2::details::Service1GetWildcardWithFlagsWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::GetWildcardWithFlags::read_result_or_throw(::basictl::tl_throwable_istream & s, std::map<std::string, ::tl2::service1::Value> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::GetWildcardWithFlags::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::map<std::string, ::tl2::service1::Value> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::Incr::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1IncrWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Incr::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1IncrRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Incr::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1IncrWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Incr::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Incr::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Incr::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1IncrReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Incr::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1IncrWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Incr::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Incr::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1IncrReset(::tl2::service1::Incr& item) noexcept {
	item.key.clear();
	item.value = 0;
}

bool tl2::details::Service1IncrWriteJSON(std::ostream& s, const ::tl2::service1::Incr& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.value != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"value\":";
		s << item.value;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1IncrRead(::basictl::tl_istream & s, ::tl2::service1::Incr& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.long_read(item.value)) { return false; }
	return true;
}

bool tl2::details::Service1IncrWrite(::basictl::tl_ostream & s, const ::tl2::service1::Incr& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.long_write(item.value)) { return false;}
	return true;
}

bool tl2::details::Service1IncrReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Incr& item) {
	if (!s.nat_read_exact_tag(0x0f96b56e)) { return false; }
	return tl2::details::Service1IncrRead(s, item);
}

bool tl2::details::Service1IncrWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Incr& item) {
	if (!s.nat_write(0x0f96b56e)) { return false; }
	return tl2::details::Service1IncrWrite(s, item);
}

bool tl2::details::Service1IncrReadResult(::basictl::tl_istream & s, tl2::service1::Incr& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::Service1IncrWriteResult(::basictl::tl_ostream & s, tl2::service1::Incr& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::service1::Incr::read_result(::basictl::tl_istream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1IncrReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::Incr::write_result(::basictl::tl_ostream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1IncrWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::Incr::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::service1::Value & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::Incr::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::service1::Value & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::KeysStat::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1KeysStatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::KeysStat::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1KeysStatRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::KeysStat::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1KeysStatWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::KeysStat::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::KeysStat::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::KeysStat::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1KeysStatReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::KeysStat::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1KeysStatWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::KeysStat::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::KeysStat::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1KeysStatReset(::tl2::service1::KeysStat& item) noexcept {
	item.start_time = 0;
	item.keys_tops.clear();
}

bool tl2::details::Service1KeysStatWriteJSON(std::ostream& s, const ::tl2::service1::KeysStat& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.start_time != 0) {
		add_comma = true;
		s << "\"start_time\":";
		s << item.start_time;
	}
	if (item.keys_tops.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"keys_tops\":";
		if (!::tl2::details::BuiltinVectorDictionaryFieldDictionaryIntWriteJSON(s, item.keys_tops)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service1KeysStatRead(::basictl::tl_istream & s, ::tl2::service1::KeysStat& item) noexcept {
	if (!s.int_read(item.start_time)) { return false; }
	if (!::tl2::details::BuiltinVectorDictionaryFieldDictionaryIntRead(s, item.keys_tops)) { return false; }
	return true;
}

bool tl2::details::Service1KeysStatWrite(::basictl::tl_ostream & s, const ::tl2::service1::KeysStat& item) noexcept {
	if (!s.int_write(item.start_time)) { return false;}
	if (!::tl2::details::BuiltinVectorDictionaryFieldDictionaryIntWrite(s, item.keys_tops)) { return false; }
	return true;
}

bool tl2::details::Service1KeysStatReadBoxed(::basictl::tl_istream & s, ::tl2::service1::KeysStat& item) {
	if (!s.nat_read_exact_tag(0xf0f6bc68)) { return false; }
	return tl2::details::Service1KeysStatRead(s, item);
}

bool tl2::details::Service1KeysStatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::KeysStat& item) {
	if (!s.nat_write(0xf0f6bc68)) { return false; }
	return tl2::details::Service1KeysStatWrite(s, item);
}

bool tl2::details::Service1KeysStatMaybeWriteJSON(std::ostream & s, const std::optional<::tl2::service1::KeysStat>& item) {
	s << "{";
	if (item) {
		s << "\"ok\":true";
		s << ",\"value\":";
		if (!::tl2::details::Service1KeysStatWriteJSON(s, *item)) { return false; }
	}
	s << "}";
	return true;
}
bool tl2::details::Service1KeysStatMaybeReadBoxed(::basictl::tl_istream & s, std::optional<::tl2::service1::KeysStat>& item) {
	bool has_item = false;
	if (!s.bool_read(has_item, 0x27930a7b, 0x3f9c8ef8)) { return false; }
	if (has_item) {
		if (!item) {
			item.emplace();
		}
		if (!::tl2::details::Service1KeysStatRead(s, *item)) { return s.set_error_unknown_scenario(); }
		return true;
	}
	item.reset();
	return true;
}

bool tl2::details::Service1KeysStatMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::service1::KeysStat>& item) {
	if (!s.nat_write(item ? 0x3f9c8ef8 : 0x27930a7b)) { return false; }
	if (item) {
		if (!::tl2::details::Service1KeysStatWrite(s, *item)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

bool tl2::service1::Longvalue::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1LongvalueWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Longvalue::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1LongvalueRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Longvalue::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1LongvalueWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Longvalue::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Longvalue::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Longvalue::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1LongvalueReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Longvalue::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1LongvalueWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Longvalue::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Longvalue::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1LongvalueReset(::tl2::service1::Longvalue& item) noexcept {
	item.value = 0;
	item.flags = 0;
}

bool tl2::details::Service1LongvalueWriteJSON(std::ostream& s, const ::tl2::service1::Longvalue& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.value != 0) {
		add_comma = true;
		s << "\"value\":";
		s << item.value;
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1LongvalueRead(::basictl::tl_istream & s, ::tl2::service1::Longvalue& item) noexcept {
	if (!s.long_read(item.value)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	return true;
}

bool tl2::details::Service1LongvalueWrite(::basictl::tl_ostream & s, const ::tl2::service1::Longvalue& item) noexcept {
	if (!s.long_write(item.value)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	return true;
}

bool tl2::details::Service1LongvalueReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Longvalue& item) {
	if (!s.nat_read_exact_tag(0x082e0945)) { return false; }
	return tl2::details::Service1LongvalueRead(s, item);
}

bool tl2::details::Service1LongvalueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Longvalue& item) {
	if (!s.nat_write(0x082e0945)) { return false; }
	return tl2::details::Service1LongvalueWrite(s, item);
}

bool tl2::service1::LongvalueWithTime::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1LongvalueWithTimeWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::LongvalueWithTime::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1LongvalueWithTimeRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::LongvalueWithTime::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1LongvalueWithTimeWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::LongvalueWithTime::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::LongvalueWithTime::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::LongvalueWithTime::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1LongvalueWithTimeReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::LongvalueWithTime::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1LongvalueWithTimeWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::LongvalueWithTime::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::LongvalueWithTime::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1LongvalueWithTimeReset(::tl2::service1::LongvalueWithTime& item) noexcept {
	item.value = 0;
	item.flags = 0;
	item.modificationTime = 0;
}

bool tl2::details::Service1LongvalueWithTimeWriteJSON(std::ostream& s, const ::tl2::service1::LongvalueWithTime& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.value != 0) {
		add_comma = true;
		s << "\"value\":";
		s << item.value;
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.modificationTime != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"modificationTime\":";
		s << item.modificationTime;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1LongvalueWithTimeRead(::basictl::tl_istream & s, ::tl2::service1::LongvalueWithTime& item) noexcept {
	if (!s.long_read(item.value)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.modificationTime)) { return false; }
	return true;
}

bool tl2::details::Service1LongvalueWithTimeWrite(::basictl::tl_ostream & s, const ::tl2::service1::LongvalueWithTime& item) noexcept {
	if (!s.long_write(item.value)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!s.int_write(item.modificationTime)) { return false;}
	return true;
}

bool tl2::details::Service1LongvalueWithTimeReadBoxed(::basictl::tl_istream & s, ::tl2::service1::LongvalueWithTime& item) {
	if (!s.nat_read_exact_tag(0xa04606ec)) { return false; }
	return tl2::details::Service1LongvalueWithTimeRead(s, item);
}

bool tl2::details::Service1LongvalueWithTimeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::LongvalueWithTime& item) {
	if (!s.nat_write(0xa04606ec)) { return false; }
	return tl2::details::Service1LongvalueWithTimeWrite(s, item);
}

bool tl2::service1::Not_found::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1NotFoundWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Not_found::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1NotFoundRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Not_found::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1NotFoundWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Not_found::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Not_found::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Not_found::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1NotFoundReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Not_found::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1NotFoundWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Not_found::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Not_found::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1NotFoundReset(::tl2::service1::Not_found& item) noexcept {
}

bool tl2::details::Service1NotFoundWriteJSON(std::ostream& s, const ::tl2::service1::Not_found& item) noexcept {
	s << "true";
	return true;
}

bool tl2::details::Service1NotFoundRead(::basictl::tl_istream & s, ::tl2::service1::Not_found& item) noexcept {
	return true;
}

bool tl2::details::Service1NotFoundWrite(::basictl::tl_ostream & s, const ::tl2::service1::Not_found& item) noexcept {
	return true;
}

bool tl2::details::Service1NotFoundReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Not_found& item) {
	if (!s.nat_read_exact_tag(0x1d670b96)) { return false; }
	return tl2::details::Service1NotFoundRead(s, item);
}

bool tl2::details::Service1NotFoundWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Not_found& item) {
	if (!s.nat_write(0x1d670b96)) { return false; }
	return tl2::details::Service1NotFoundWrite(s, item);
}

bool tl2::service1::Replace::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1ReplaceWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Replace::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1ReplaceRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Replace::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1ReplaceWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Replace::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Replace::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Replace::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1ReplaceReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Replace::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1ReplaceWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Replace::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Replace::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1ReplaceReset(::tl2::service1::Replace& item) noexcept {
	item.key.clear();
	item.flags = 0;
	item.delay = 0;
	item.value.clear();
}

bool tl2::details::Service1ReplaceWriteJSON(std::ostream& s, const ::tl2::service1::Replace& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.delay != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"delay\":";
		s << item.delay;
	}
	if (item.value.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"value\":";
		s << "\"" << item.value << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1ReplaceRead(::basictl::tl_istream & s, ::tl2::service1::Replace& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.delay)) { return false; }
	if (!s.string_read(item.value)) { return false; }
	return true;
}

bool tl2::details::Service1ReplaceWrite(::basictl::tl_ostream & s, const ::tl2::service1::Replace& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!s.int_write(item.delay)) { return false;}
	if (!s.string_write(item.value)) { return false;}
	return true;
}

bool tl2::details::Service1ReplaceReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Replace& item) {
	if (!s.nat_read_exact_tag(0x7f2c447d)) { return false; }
	return tl2::details::Service1ReplaceRead(s, item);
}

bool tl2::details::Service1ReplaceWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Replace& item) {
	if (!s.nat_write(0x7f2c447d)) { return false; }
	return tl2::details::Service1ReplaceWrite(s, item);
}

bool tl2::details::Service1ReplaceReadResult(::basictl::tl_istream & s, tl2::service1::Replace& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1ReplaceWriteResult(::basictl::tl_ostream & s, tl2::service1::Replace& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::Replace::read_result(::basictl::tl_istream & s, bool & result) noexcept {
	bool success = tl2::details::Service1ReplaceReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::Replace::write_result(::basictl::tl_ostream & s, bool & result) noexcept {
	bool success = tl2::details::Service1ReplaceWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::Replace::read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::Replace::write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::ReplaceOrIncr::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1ReplaceOrIncrWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::ReplaceOrIncr::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1ReplaceOrIncrRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::ReplaceOrIncr::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1ReplaceOrIncrWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::ReplaceOrIncr::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::ReplaceOrIncr::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::ReplaceOrIncr::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1ReplaceOrIncrReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::ReplaceOrIncr::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1ReplaceOrIncrWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::ReplaceOrIncr::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::ReplaceOrIncr::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1ReplaceOrIncrReset(::tl2::service1::ReplaceOrIncr& item) noexcept {
	item.key.clear();
	item.flags = 0;
	item.delay = 0;
	item.value = 0;
}

bool tl2::details::Service1ReplaceOrIncrWriteJSON(std::ostream& s, const ::tl2::service1::ReplaceOrIncr& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.delay != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"delay\":";
		s << item.delay;
	}
	if (item.value != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"value\":";
		s << item.value;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1ReplaceOrIncrRead(::basictl::tl_istream & s, ::tl2::service1::ReplaceOrIncr& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.delay)) { return false; }
	if (!s.long_read(item.value)) { return false; }
	return true;
}

bool tl2::details::Service1ReplaceOrIncrWrite(::basictl::tl_ostream & s, const ::tl2::service1::ReplaceOrIncr& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!s.int_write(item.delay)) { return false;}
	if (!s.long_write(item.value)) { return false;}
	return true;
}

bool tl2::details::Service1ReplaceOrIncrReadBoxed(::basictl::tl_istream & s, ::tl2::service1::ReplaceOrIncr& item) {
	if (!s.nat_read_exact_tag(0x9d1bdcfd)) { return false; }
	return tl2::details::Service1ReplaceOrIncrRead(s, item);
}

bool tl2::details::Service1ReplaceOrIncrWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::ReplaceOrIncr& item) {
	if (!s.nat_write(0x9d1bdcfd)) { return false; }
	return tl2::details::Service1ReplaceOrIncrWrite(s, item);
}

bool tl2::details::Service1ReplaceOrIncrReadResult(::basictl::tl_istream & s, tl2::service1::ReplaceOrIncr& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::Service1ReplaceOrIncrWriteResult(::basictl::tl_ostream & s, tl2::service1::ReplaceOrIncr& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::service1::ReplaceOrIncr::read_result(::basictl::tl_istream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1ReplaceOrIncrReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::ReplaceOrIncr::write_result(::basictl::tl_ostream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1ReplaceOrIncrWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::ReplaceOrIncr::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::service1::Value & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::ReplaceOrIncr::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::service1::Value & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::Set::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1SetWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Set::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1SetRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Set::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1SetWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Set::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Set::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Set::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1SetReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Set::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1SetWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Set::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Set::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1SetReset(::tl2::service1::Set& item) noexcept {
	item.key.clear();
	item.flags = 0;
	item.delay = 0;
	item.value.clear();
}

bool tl2::details::Service1SetWriteJSON(std::ostream& s, const ::tl2::service1::Set& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.delay != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"delay\":";
		s << item.delay;
	}
	if (item.value.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"value\":";
		s << "\"" << item.value << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::Service1SetRead(::basictl::tl_istream & s, ::tl2::service1::Set& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.delay)) { return false; }
	if (!s.string_read(item.value)) { return false; }
	return true;
}

bool tl2::details::Service1SetWrite(::basictl::tl_ostream & s, const ::tl2::service1::Set& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!s.int_write(item.delay)) { return false;}
	if (!s.string_write(item.value)) { return false;}
	return true;
}

bool tl2::details::Service1SetReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Set& item) {
	if (!s.nat_read_exact_tag(0x05ae5f66)) { return false; }
	return tl2::details::Service1SetRead(s, item);
}

bool tl2::details::Service1SetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Set& item) {
	if (!s.nat_write(0x05ae5f66)) { return false; }
	return tl2::details::Service1SetWrite(s, item);
}

bool tl2::details::Service1SetReadResult(::basictl::tl_istream & s, tl2::service1::Set& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1SetWriteResult(::basictl::tl_ostream & s, tl2::service1::Set& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::Set::read_result(::basictl::tl_istream & s, bool & result) noexcept {
	bool success = tl2::details::Service1SetReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::Set::write_result(::basictl::tl_ostream & s, bool & result) noexcept {
	bool success = tl2::details::Service1SetWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::Set::read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::Set::write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::SetOrIncr::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1SetOrIncrWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::SetOrIncr::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1SetOrIncrRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::SetOrIncr::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1SetOrIncrWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::SetOrIncr::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::SetOrIncr::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::SetOrIncr::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1SetOrIncrReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::SetOrIncr::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1SetOrIncrWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::SetOrIncr::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::SetOrIncr::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1SetOrIncrReset(::tl2::service1::SetOrIncr& item) noexcept {
	item.key.clear();
	item.flags = 0;
	item.delay = 0;
	item.value = 0;
}

bool tl2::details::Service1SetOrIncrWriteJSON(std::ostream& s, const ::tl2::service1::SetOrIncr& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.delay != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"delay\":";
		s << item.delay;
	}
	if (item.value != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"value\":";
		s << item.value;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1SetOrIncrRead(::basictl::tl_istream & s, ::tl2::service1::SetOrIncr& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.delay)) { return false; }
	if (!s.long_read(item.value)) { return false; }
	return true;
}

bool tl2::details::Service1SetOrIncrWrite(::basictl::tl_ostream & s, const ::tl2::service1::SetOrIncr& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!s.int_write(item.delay)) { return false;}
	if (!s.long_write(item.value)) { return false;}
	return true;
}

bool tl2::details::Service1SetOrIncrReadBoxed(::basictl::tl_istream & s, ::tl2::service1::SetOrIncr& item) {
	if (!s.nat_read_exact_tag(0x772e390d)) { return false; }
	return tl2::details::Service1SetOrIncrRead(s, item);
}

bool tl2::details::Service1SetOrIncrWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::SetOrIncr& item) {
	if (!s.nat_write(0x772e390d)) { return false; }
	return tl2::details::Service1SetOrIncrWrite(s, item);
}

bool tl2::details::Service1SetOrIncrReadResult(::basictl::tl_istream & s, tl2::service1::SetOrIncr& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::Service1SetOrIncrWriteResult(::basictl::tl_ostream & s, tl2::service1::SetOrIncr& item, ::tl2::service1::Value& result) {
	if (!::tl2::details::Service1ValueWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::service1::SetOrIncr::read_result(::basictl::tl_istream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1SetOrIncrReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::SetOrIncr::write_result(::basictl::tl_ostream & s, ::tl2::service1::Value & result) noexcept {
	bool success = tl2::details::Service1SetOrIncrWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::SetOrIncr::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::service1::Value & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::SetOrIncr::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::service1::Value & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::service1::Strvalue::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1StrvalueWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Strvalue::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1StrvalueRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Strvalue::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1StrvalueWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Strvalue::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Strvalue::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Strvalue::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1StrvalueReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Strvalue::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1StrvalueWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Strvalue::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Strvalue::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1StrvalueReset(::tl2::service1::Strvalue& item) noexcept {
	item.value.clear();
	item.flags = 0;
}

bool tl2::details::Service1StrvalueWriteJSON(std::ostream& s, const ::tl2::service1::Strvalue& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.value.size() != 0) {
		add_comma = true;
		s << "\"value\":";
		s << "\"" << item.value << "\"";
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1StrvalueRead(::basictl::tl_istream & s, ::tl2::service1::Strvalue& item) noexcept {
	if (!s.string_read(item.value)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	return true;
}

bool tl2::details::Service1StrvalueWrite(::basictl::tl_ostream & s, const ::tl2::service1::Strvalue& item) noexcept {
	if (!s.string_write(item.value)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	return true;
}

bool tl2::details::Service1StrvalueReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Strvalue& item) {
	if (!s.nat_read_exact_tag(0x5faa0c52)) { return false; }
	return tl2::details::Service1StrvalueRead(s, item);
}

bool tl2::details::Service1StrvalueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Strvalue& item) {
	if (!s.nat_write(0x5faa0c52)) { return false; }
	return tl2::details::Service1StrvalueWrite(s, item);
}

bool tl2::service1::StrvalueWithTime::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1StrvalueWithTimeWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::StrvalueWithTime::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1StrvalueWithTimeRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::StrvalueWithTime::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1StrvalueWithTimeWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::StrvalueWithTime::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::StrvalueWithTime::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::StrvalueWithTime::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1StrvalueWithTimeReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::StrvalueWithTime::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1StrvalueWithTimeWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::StrvalueWithTime::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::StrvalueWithTime::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1StrvalueWithTimeReset(::tl2::service1::StrvalueWithTime& item) noexcept {
	item.value.clear();
	item.flags = 0;
	item.modificationTime = 0;
}

bool tl2::details::Service1StrvalueWithTimeWriteJSON(std::ostream& s, const ::tl2::service1::StrvalueWithTime& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.value.size() != 0) {
		add_comma = true;
		s << "\"value\":";
		s << "\"" << item.value << "\"";
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.modificationTime != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"modificationTime\":";
		s << item.modificationTime;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1StrvalueWithTimeRead(::basictl::tl_istream & s, ::tl2::service1::StrvalueWithTime& item) noexcept {
	if (!s.string_read(item.value)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!s.int_read(item.modificationTime)) { return false; }
	return true;
}

bool tl2::details::Service1StrvalueWithTimeWrite(::basictl::tl_ostream & s, const ::tl2::service1::StrvalueWithTime& item) noexcept {
	if (!s.string_write(item.value)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!s.int_write(item.modificationTime)) { return false;}
	return true;
}

bool tl2::details::Service1StrvalueWithTimeReadBoxed(::basictl::tl_istream & s, ::tl2::service1::StrvalueWithTime& item) {
	if (!s.nat_read_exact_tag(0x98b1a484)) { return false; }
	return tl2::details::Service1StrvalueWithTimeRead(s, item);
}

bool tl2::details::Service1StrvalueWithTimeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::StrvalueWithTime& item) {
	if (!s.nat_write(0x98b1a484)) { return false; }
	return tl2::details::Service1StrvalueWithTimeWrite(s, item);
}

bool tl2::service1::Touch::write_json(std::ostream& s)const {
	if (!::tl2::details::Service1TouchWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service1::Touch::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1TouchRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Touch::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1TouchWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Touch::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::service1::Touch::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::service1::Touch::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1TouchReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::service1::Touch::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1TouchWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::service1::Touch::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Touch::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Service1TouchReset(::tl2::service1::Touch& item) noexcept {
	item.key.clear();
	item.delay = 0;
}

bool tl2::details::Service1TouchWriteJSON(std::ostream& s, const ::tl2::service1::Touch& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (item.delay != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"delay\":";
		s << item.delay;
	}
	s << "}";
	return true;
}

bool tl2::details::Service1TouchRead(::basictl::tl_istream & s, ::tl2::service1::Touch& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.int_read(item.delay)) { return false; }
	return true;
}

bool tl2::details::Service1TouchWrite(::basictl::tl_ostream & s, const ::tl2::service1::Touch& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.int_write(item.delay)) { return false;}
	return true;
}

bool tl2::details::Service1TouchReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Touch& item) {
	if (!s.nat_read_exact_tag(0xb737aa03)) { return false; }
	return tl2::details::Service1TouchRead(s, item);
}

bool tl2::details::Service1TouchWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Touch& item) {
	if (!s.nat_write(0xb737aa03)) { return false; }
	return tl2::details::Service1TouchWrite(s, item);
}

bool tl2::details::Service1TouchReadResult(::basictl::tl_istream & s, tl2::service1::Touch& item, bool& result) {
	if (!::tl2::details::BoolReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service1TouchWriteResult(::basictl::tl_ostream & s, tl2::service1::Touch& item, bool& result) {
	if (!::tl2::details::BoolWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service1::Touch::read_result(::basictl::tl_istream & s, bool & result) noexcept {
	bool success = tl2::details::Service1TouchReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::service1::Touch::write_result(::basictl::tl_ostream & s, bool & result) noexcept {
	bool success = tl2::details::Service1TouchWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::service1::Touch::read_result_or_throw(::basictl::tl_throwable_istream & s, bool & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::service1::Touch::write_result_or_throw(::basictl::tl_throwable_ostream & s, bool & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

static const std::string_view Service1Value_tbl_tl_name[]{"service1.not_found", "service1.strvalue", "service1.longvalue", "service1.strvalueWithTime", "service1.longvalueWithTime"};
static const uint32_t Service1Value_tbl_tl_tag[]{0x1d670b96, 0x5faa0c52, 0x082e0945, 0x98b1a484, 0xa04606ec};

bool tl2::service1::Value::write_json(std::ostream & s)const {
	if (!::tl2::details::Service1ValueWriteJSON(s, *this)) { return false; }
	return true;
}
bool tl2::service1::Value::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Service1ValueReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::service1::Value::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Service1ValueWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	return true;
}

void tl2::service1::Value::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::service1::Value::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

std::string_view tl2::service1::Value::tl_name() const {
	return Service1Value_tbl_tl_name[value.index()];
}
uint32_t tl2::service1::Value::tl_tag() const {
	return Service1Value_tbl_tl_tag[value.index()];
}


void tl2::details::Service1ValueReset(::tl2::service1::Value& item) noexcept{
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::Service1ValueWriteJSON(std::ostream & s, const ::tl2::service1::Value& item) noexcept {
	s << "{";
	s << "\"type\":";
	s << "\"" << Service1Value_tbl_tl_name[item.value.index()] << "\"";
	switch (item.value.index()) {
	case 1:
		s << ",\"value\":";
		if (!::tl2::details::Service1StrvalueWriteJSON(s, std::get<1>(item.value))) { return false; }
		break;
	case 2:
		s << ",\"value\":";
		if (!::tl2::details::Service1LongvalueWriteJSON(s, std::get<2>(item.value))) { return false; }
		break;
	case 3:
		s << ",\"value\":";
		if (!::tl2::details::Service1StrvalueWithTimeWriteJSON(s, std::get<3>(item.value))) { return false; }
		break;
	case 4:
		s << ",\"value\":";
		if (!::tl2::details::Service1LongvalueWithTimeWriteJSON(s, std::get<4>(item.value))) { return false; }
		break;
	}
	s << "}";
	return true;
}
bool tl2::details::Service1ValueReadBoxed(::basictl::tl_istream & s, ::tl2::service1::Value& item) noexcept {
	uint32_t nat;
	if (!s.nat_read(nat)) { return false; }
	switch (nat) {
	case 0x1d670b96:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		break;
	case 0x5faa0c52:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		if (!::tl2::details::Service1StrvalueRead(s, std::get<1>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	case 0x082e0945:
		if (item.value.index() != 2) { item.value.emplace<2>(); }
		if (!::tl2::details::Service1LongvalueRead(s, std::get<2>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	case 0x98b1a484:
		if (item.value.index() != 3) { item.value.emplace<3>(); }
		if (!::tl2::details::Service1StrvalueWithTimeRead(s, std::get<3>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	case 0xa04606ec:
		if (item.value.index() != 4) { item.value.emplace<4>(); }
		if (!::tl2::details::Service1LongvalueWithTimeRead(s, std::get<4>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::Service1ValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service1::Value& item) noexcept{
	if (!s.nat_write(Service1Value_tbl_tl_tag[item.value.index()])) { return false; }
	switch (item.value.index()) {
	case 1:
		if (!::tl2::details::Service1StrvalueWrite(s, std::get<1>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	case 2:
		if (!::tl2::details::Service1LongvalueWrite(s, std::get<2>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	case 3:
		if (!::tl2::details::Service1StrvalueWithTimeWrite(s, std::get<3>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	case 4:
		if (!::tl2::details::Service1LongvalueWithTimeWrite(s, std::get<4>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	}
	return true;
}

bool tl2::details::Service1ValueBoxedMaybeWriteJSON(std::ostream & s, const std::optional<::tl2::service1::Value>& item) {
	s << "{";
	if (item) {
		s << "\"ok\":true";
		s << ",\"value\":";
		if (!::tl2::details::Service1ValueWriteJSON(s, *item)) { return false; }
	}
	s << "}";
	return true;
}
bool tl2::details::Service1ValueBoxedMaybeReadBoxed(::basictl::tl_istream & s, std::optional<::tl2::service1::Value>& item) {
	bool has_item = false;
	if (!s.bool_read(has_item, 0x27930a7b, 0x3f9c8ef8)) { return false; }
	if (has_item) {
		if (!item) {
			item.emplace();
		}
		if (!::tl2::details::Service1ValueReadBoxed(s, *item)) { return s.set_error_unknown_scenario(); }
		return true;
	}
	item.reset();
	return true;
}

bool tl2::details::Service1ValueBoxedMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::service1::Value>& item) {
	if (!s.nat_write(item ? 0x3f9c8ef8 : 0x27930a7b)) { return false; }
	if (item) {
		if (!::tl2::details::Service1ValueWriteBoxed(s, *item)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

void tl2::details::VectorDictionaryFieldDictionaryIntReset(std::map<std::string, std::map<std::string, int32_t>>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorDictionaryFieldDictionaryIntWriteJSON(std::ostream& s, const std::map<std::string, std::map<std::string, int32_t>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldDictionaryIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorDictionaryFieldDictionaryIntRead(::basictl::tl_istream & s, std::map<std::string, std::map<std::string, int32_t>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldDictionaryIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorDictionaryFieldDictionaryIntWrite(::basictl::tl_ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldDictionaryIntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorDictionaryFieldDictionaryIntReadBoxed(::basictl::tl_istream & s, std::map<std::string, std::map<std::string, int32_t>>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorDictionaryFieldDictionaryIntRead(s, item);
}

bool tl2::details::VectorDictionaryFieldDictionaryIntWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, std::map<std::string, int32_t>>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorDictionaryFieldDictionaryIntWrite(s, item);
}
