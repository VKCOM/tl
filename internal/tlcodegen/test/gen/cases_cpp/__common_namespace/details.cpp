#include "headers/vector.h"
#include "headers/string.h"
#include "headers/long.h"
#include "headers/int64.h"
#include "headers/int32.h"
#include "headers/int.h"
#include "headers/dictionary.h"
#include "headers/dictionaryField.h"
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

void tl2::details::BuiltinVectorDictionaryFieldIntReset(std::vector<::tl2::DictionaryField<int32_t>>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorDictionaryFieldIntWriteJSON(std::ostream & s, const std::vector<::tl2::DictionaryField<int32_t>>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::DictionaryFieldIntWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<int32_t>>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::DictionaryFieldIntRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorDictionaryFieldIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<int32_t>>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::DictionaryFieldIntWrite(s, el)) { return false; }
	}
	return true;
}

void tl2::details::DictionaryAnyDoubleIntReset(::tl2::DictionaryAny<double, int32_t>& item) {
	item.clear();
}

bool tl2::details::DictionaryAnyDoubleIntWriteJSON(std::ostream& s, const ::tl2::DictionaryAny<double, int32_t>& item) {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyDoubleIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryAnyDoubleIntRead(::basictl::tl_istream & s, ::tl2::DictionaryAny<double, int32_t>& item) {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyDoubleIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryAnyDoubleIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryAny<double, int32_t>& item) {
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

void tl2::details::DictionaryAnyIntIntReset(std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) {
	item.clear();
}

bool tl2::details::DictionaryAnyIntIntWriteJSON(std::ostream& s, const std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyIntIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryAnyIntIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyIntIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryAnyIntIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryFieldAny<int32_t, int32_t>>& item) {
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

void tl2::details::DictionaryFieldAnyDoubleIntReset(::tl2::DictionaryFieldAny<double, int32_t>& item) {
	item.key = 0;
	item.value = 0;
}

bool tl2::details::DictionaryFieldAnyDoubleIntWriteJSON(std::ostream& s, const ::tl2::DictionaryFieldAny<double, int32_t>& item) {
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

bool tl2::details::DictionaryFieldAnyDoubleIntRead(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<double, int32_t>& item) {
	if (!s.double_read(item.key)) { return false; }
	if (!s.int_read(item.value)) { return false; }
	return true;
}

bool tl2::details::DictionaryFieldAnyDoubleIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<double, int32_t>& item) {
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

void tl2::details::DictionaryFieldAnyIntIntReset(::tl2::DictionaryFieldAny<int32_t, int32_t>& item) {
	item.key = 0;
	item.value = 0;
}

bool tl2::details::DictionaryFieldAnyIntIntWriteJSON(std::ostream& s, const ::tl2::DictionaryFieldAny<int32_t, int32_t>& item) {
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

bool tl2::details::DictionaryFieldAnyIntIntRead(::basictl::tl_istream & s, ::tl2::DictionaryFieldAny<int32_t, int32_t>& item) {
	if (!s.int_read(item.key)) { return false; }
	if (!s.int_read(item.value)) { return false; }
	return true;
}

bool tl2::details::DictionaryFieldAnyIntIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryFieldAny<int32_t, int32_t>& item) {
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

void tl2::details::DictionaryFieldIntReset(::tl2::DictionaryField<int32_t>& item) {
	item.key.clear();
	item.value = 0;
}

bool tl2::details::DictionaryFieldIntWriteJSON(std::ostream& s, const ::tl2::DictionaryField<int32_t>& item) {
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

bool tl2::details::DictionaryFieldIntRead(::basictl::tl_istream & s, ::tl2::DictionaryField<int32_t>& item) {
	if (!s.string_read(item.key)) { return false; }
	if (!s.int_read(item.value)) { return false; }
	return true;
}

bool tl2::details::DictionaryFieldIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<int32_t>& item) {
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

void tl2::details::DictionaryIntReset(::tl2::Dictionary<int32_t>& item) {
	item.clear();
}

bool tl2::details::DictionaryIntWriteJSON(std::ostream& s, const ::tl2::Dictionary<int32_t>& item) {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryIntRead(::basictl::tl_istream & s, ::tl2::Dictionary<int32_t>& item) {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryIntWrite(::basictl::tl_ostream & s, const ::tl2::Dictionary<int32_t>& item) {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::DictionaryIntReadBoxed(::basictl::tl_istream & s, ::tl2::Dictionary<int32_t>& item) {
	if (!s.nat_read_exact_tag(0x1f4c618f)) { return false; }
	return tl2::details::DictionaryIntRead(s, item);
}

bool tl2::details::DictionaryIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Dictionary<int32_t>& item) {
	if (!s.nat_write(0x1f4c618f)) { return false; }
	return tl2::details::DictionaryIntWrite(s, item);
}

void tl2::details::IntReset(int32_t& item) {
	item = 0;
}

bool tl2::details::IntWriteJSON(std::ostream& s, const int32_t& item) {
	s << item;
	return true;
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

bool tl2::details::Int32WriteJSON(std::ostream& s, const ::tl2::Int32& item) {
	s << item;
	return true;
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

bool tl2::details::Int64WriteJSON(std::ostream& s, const ::tl2::Int64& item) {
	s << item;
	return true;
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

bool tl2::details::LongWriteJSON(std::ostream& s, const int64_t& item) {
	s << item;
	return true;
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

bool tl2::details::StringWriteJSON(std::ostream& s, const std::string& item) {
	s << "\"" << item << "\"";
	return true;
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

void tl2::details::TrueReset(::tl2::True& item) {
}

bool tl2::details::TrueWriteJSON(std::ostream& s, const ::tl2::True& item) {
	s << "true";
	return true;
}

bool tl2::details::TrueRead(::basictl::tl_istream & s, ::tl2::True& item) {
	return true;
}

bool tl2::details::TrueWrite(::basictl::tl_ostream & s, const ::tl2::True& item) {
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

void tl2::details::VectorDictionaryFieldIntReset(std::vector<::tl2::DictionaryField<int32_t>>& item) {
	item.clear();
}

bool tl2::details::VectorDictionaryFieldIntWriteJSON(std::ostream& s, const std::vector<::tl2::DictionaryField<int32_t>>& item) {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorDictionaryFieldIntRead(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<int32_t>>& item) {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorDictionaryFieldIntWrite(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<int32_t>>& item) {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorDictionaryFieldIntReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::DictionaryField<int32_t>>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorDictionaryFieldIntRead(s, item);
}

bool tl2::details::VectorDictionaryFieldIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::DictionaryField<int32_t>>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorDictionaryFieldIntWrite(s, item);
}
