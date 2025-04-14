#include "headers/vector.h"
#include "headers/int64.h"
#include "headers/int32.h"
#include "headers/dictionaryField.h"
#include "headers/dictionary.h"
#include "headers/dictionaryAny.h"
#include "headers/dictionaryFieldAny.h"
#include "headers/true.h"


void tl2::details::BuiltinVectorDictionaryFieldAnyDoubleIntReset(std::vector<::tl2::DictionaryFieldAny<double, int32_t>>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorDictionaryFieldAnyDoubleIntWriteJSON(std::ostream & s, const std::vector<::tl2::DictionaryFieldAny<double, int32_t>>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::DictionaryFieldAnyDoubleIntWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldAnyDoubleIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryFieldAny<double, int32_t>>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::DictionaryFieldAnyDoubleIntRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldAnyDoubleIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryFieldAny<double, int32_t>>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::DictionaryFieldAnyDoubleIntWrite(s, el)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinVectorDictionaryFieldAnyIntIntReset(std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorDictionaryFieldAnyIntIntWriteJSON(std::ostream & s, const std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::DictionaryFieldAnyIntIntWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldAnyIntIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::DictionaryFieldAnyIntIntRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldAnyIntIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::DictionaryFieldAnyIntIntWrite(s, el)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinVectorDictionaryFieldIntReset(std::map<std::string, int32_t>& item) {
	item.clear(); // TODO - unwrap
}

bool tl2::details::BuiltinVectorDictionaryFieldIntWriteJSON(std::ostream & s, const std::map<std::string, int32_t>& item) {
	s << "{";
	size_t index = 0;
	for(const auto & el : item) {
		s << "\"" << el.first << "\"";
		s << ":";
		s << el.second;
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "}";
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldIntRead(::basictl::tl_istream & s, std::map<std::string, int32_t>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	item.clear();
	for(uint32_t i = 0; i < len; i++) {
		std::string key;
		if (!s.string_read(key)) { return false; }
		if (!s.int_read(item[key])) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldIntWrite(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!s.string_write(el.first)) { return false;}
		if (!s.int_write(el.second)) { return false;}
	}
	return true;
}

void tl2::details::DictionaryAnyDoubleIntReset(::tl2::DictionaryAny<double, int32_t>& item) noexcept {
	item.clear();
}

bool tl2::details::DictionaryAnyDoubleIntWriteJSON(std::ostream& s, const ::tl2::DictionaryAny<double, int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyDoubleIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryAnyDoubleIntRead(::basictl::tl_istream & s, ::tl2::DictionaryAny<double, int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyDoubleIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryAnyDoubleIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryAny<double, int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyDoubleIntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryAnyDoubleIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryAny<double, int32_t>& item) {
	if (!s.nat_read_exact_tag(0x1f4c6190)) { return false; }
	return tl2::details::DictionaryAnyDoubleIntRead(s, item);
}

bool tl2::details::DictionaryAnyDoubleIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryAny<double, int32_t>& item) {
	if (!s.nat_write(0x1f4c6190)) { return false; }
	return tl2::details::DictionaryAnyDoubleIntWrite(s, item);
}

void tl2::details::DictionaryAnyIntIntReset(std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) noexcept {
	item.clear();
}

bool tl2::details::DictionaryAnyIntIntWriteJSON(std::ostream& s, const std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyIntIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryAnyIntIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyIntIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryAnyIntIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyIntIntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryAnyIntIntReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) {
	if (!s.nat_read_exact_tag(0x1f4c6190)) { return false; }
	return tl2::details::DictionaryAnyIntIntRead(s, item);
}

bool tl2::details::DictionaryAnyIntIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) {
	if (!s.nat_write(0x1f4c6190)) { return false; }
	return tl2::details::DictionaryAnyIntIntWrite(s, item);
}

void tl2::details::DictionaryFieldAnyDoubleIntReset(::tl2::DictionaryFieldAny<double, int32_t>& item) noexcept {
	item.key = 0;
	item.value = 0;
}

bool tl2::details::DictionaryFieldAnyDoubleIntWriteJSON(std::ostream& s, const ::tl2::DictionaryFieldAny<double, int32_t>& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key != 0) {
		add_comma = true;
		s << "\"key\":";
		s << item.key;
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

bool tl2::details::DictionaryFieldAnyDoubleIntRead(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<double, int32_t>& item) noexcept {
	if (!s.double_read(item.key)) { return false; }
	if (!s.int_read(item.value)) { return false; }
	return true;
}

bool tl2::details::DictionaryFieldAnyDoubleIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<double, int32_t>& item) noexcept {
	if (!s.double_write(item.key)) { return false;}
	if (!s.int_write(item.value)) { return false;}
	return true;
}

bool tl2::details::DictionaryFieldAnyDoubleIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<double, int32_t>& item) {
	if (!s.nat_read_exact_tag(0x2c43a65b)) { return false; }
	return tl2::details::DictionaryFieldAnyDoubleIntRead(s, item);
}

bool tl2::details::DictionaryFieldAnyDoubleIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<double, int32_t>& item) {
	if (!s.nat_write(0x2c43a65b)) { return false; }
	return tl2::details::DictionaryFieldAnyDoubleIntWrite(s, item);
}

void tl2::details::DictionaryFieldAnyIntIntReset(::tl2::DictionaryFieldAny<int32_t, int32_t>& item) noexcept {
	item.key = 0;
	item.value = 0;
}

bool tl2::details::DictionaryFieldAnyIntIntWriteJSON(std::ostream& s, const ::tl2::DictionaryFieldAny<int32_t, int32_t>& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key != 0) {
		add_comma = true;
		s << "\"key\":";
		s << item.key;
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

bool tl2::details::DictionaryFieldAnyIntIntRead(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<int32_t, int32_t>& item) noexcept {
	if (!s.int_read(item.key)) { return false; }
	if (!s.int_read(item.value)) { return false; }
	return true;
}

bool tl2::details::DictionaryFieldAnyIntIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<int32_t, int32_t>& item) noexcept {
	if (!s.int_write(item.key)) { return false;}
	if (!s.int_write(item.value)) { return false;}
	return true;
}

bool tl2::details::DictionaryFieldAnyIntIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<int32_t, int32_t>& item) {
	if (!s.nat_read_exact_tag(0x2c43a65b)) { return false; }
	return tl2::details::DictionaryFieldAnyIntIntRead(s, item);
}

bool tl2::details::DictionaryFieldAnyIntIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<int32_t, int32_t>& item) {
	if (!s.nat_write(0x2c43a65b)) { return false; }
	return tl2::details::DictionaryFieldAnyIntIntWrite(s, item);
}

void tl2::details::DictionaryFieldIntReset(::tl2::DictionaryField<int32_t>& item) noexcept {
	item.key.clear();
	item.value = 0;
}

bool tl2::details::DictionaryFieldIntWriteJSON(std::ostream& s, const ::tl2::DictionaryField<int32_t>& item) noexcept {
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

bool tl2::details::DictionaryFieldIntRead(::basictl::tl_istream & s, ::tl2::DictionaryField<int32_t>& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.int_read(item.value)) { return false; }
	return true;
}

bool tl2::details::DictionaryFieldIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<int32_t>& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.int_write(item.value)) { return false;}
	return true;
}

bool tl2::details::DictionaryFieldIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<int32_t>& item) {
	if (!s.nat_read_exact_tag(0x239c1b62)) { return false; }
	return tl2::details::DictionaryFieldIntRead(s, item);
}

bool tl2::details::DictionaryFieldIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<int32_t>& item) {
	if (!s.nat_write(0x239c1b62)) { return false; }
	return tl2::details::DictionaryFieldIntWrite(s, item);
}

void tl2::details::DictionaryFieldStringReset(::tl2::DictionaryField<std::string>& item) noexcept {
	item.key.clear();
	item.value.clear();
}

bool tl2::details::DictionaryFieldStringWriteJSON(std::ostream& s, const ::tl2::DictionaryField<std::string>& item) noexcept {
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

bool tl2::details::DictionaryFieldStringRead(::basictl::tl_istream & s, ::tl2::DictionaryField<std::string>& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.string_read(item.value)) { return false; }
	return true;
}

bool tl2::details::DictionaryFieldStringWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::string>& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.string_write(item.value)) { return false;}
	return true;
}

bool tl2::details::DictionaryFieldStringReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<std::string>& item) {
	if (!s.nat_read_exact_tag(0x239c1b62)) { return false; }
	return tl2::details::DictionaryFieldStringRead(s, item);
}

bool tl2::details::DictionaryFieldStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::string>& item) {
	if (!s.nat_write(0x239c1b62)) { return false; }
	return tl2::details::DictionaryFieldStringWrite(s, item);
}

void tl2::details::DictionaryIntReset(std::map<std::string, int32_t>& item) noexcept {
	item.clear();
}

bool tl2::details::DictionaryIntWriteJSON(std::ostream& s, const std::map<std::string, int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryIntRead(::basictl::tl_istream & s, std::map<std::string, int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryIntWrite(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryIntReadBoxed(::basictl::tl_istream & s, std::map<std::string, int32_t>& item) {
	if (!s.nat_read_exact_tag(0x1f4c618f)) { return false; }
	return tl2::details::DictionaryIntRead(s, item);
}

bool tl2::details::DictionaryIntWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item) {
	if (!s.nat_write(0x1f4c618f)) { return false; }
	return tl2::details::DictionaryIntWrite(s, item);
}

void tl2::details::Int32Reset(::tl2::Int32& item) noexcept {
	item = 0;
}

bool tl2::details::Int32WriteJSON(std::ostream& s, const ::tl2::Int32& item) noexcept {
	s << item;
	return true;
}

bool tl2::details::Int32Read(::basictl::tl_istream & s, ::tl2::Int32& item) noexcept {
	if (!s.int_read(item)) { return false; }
	return true;
}

bool tl2::details::Int32Write(::basictl::tl_ostream & s, const ::tl2::Int32& item) noexcept {
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

void tl2::details::Int64Reset(::tl2::Int64& item) noexcept {
	item = 0;
}

bool tl2::details::Int64WriteJSON(std::ostream& s, const ::tl2::Int64& item) noexcept {
	s << item;
	return true;
}

bool tl2::details::Int64Read(::basictl::tl_istream & s, ::tl2::Int64& item) noexcept {
	if (!s.long_read(item)) { return false; }
	return true;
}

bool tl2::details::Int64Write(::basictl::tl_ostream & s, const ::tl2::Int64& item) noexcept {
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

bool tl2::True::write_json(std::ostream& s)const {
	if (!::tl2::details::TrueWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::True::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TrueRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::True::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TrueWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::True::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::True::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::True::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TrueReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::True::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TrueWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::True::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::True::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TrueReset(::tl2::True& item) noexcept {
}

bool tl2::details::TrueWriteJSON(std::ostream& s, const ::tl2::True& item) noexcept {
	s << "true";
	return true;
}

bool tl2::details::TrueRead(::basictl::tl_istream & s, ::tl2::True& item) noexcept {
	return true;
}

bool tl2::details::TrueWrite(::basictl::tl_ostream & s, const ::tl2::True& item) noexcept {
	return true;
}

bool tl2::details::TrueReadBoxed(::basictl::tl_istream & s, ::tl2::True& item) {
	if (!s.nat_read_exact_tag(0x3fedd339)) { return false; }
	return tl2::details::TrueRead(s, item);
}

bool tl2::details::TrueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::True& item) {
	if (!s.nat_write(0x3fedd339)) { return false; }
	return tl2::details::TrueWrite(s, item);
}

void tl2::details::VectorDictionaryFieldIntReset(std::map<std::string, int32_t>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorDictionaryFieldIntWriteJSON(std::ostream& s, const std::map<std::string, int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorDictionaryFieldIntRead(::basictl::tl_istream & s, std::map<std::string, int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorDictionaryFieldIntWrite(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorDictionaryFieldIntReadBoxed(::basictl::tl_istream & s, std::map<std::string, int32_t>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorDictionaryFieldIntRead(s, item);
}

bool tl2::details::VectorDictionaryFieldIntWriteBoxed(::basictl::tl_ostream & s, const std::map<std::string, int32_t>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorDictionaryFieldIntWrite(s, item);
}
