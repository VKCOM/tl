#include "headers/vector.hpp"
#include "headers/dictionary.hpp"
#include "headers/dictionaryField.hpp"
#include "headers/dictionaryAny.hpp"
#include "headers/dictionaryFieldAny.hpp"
#include "headers/true.hpp"


void tl2::details::BuiltinVectorDictionaryFieldAnyDoubleIntReset(std::vector<::tl2::DictionaryFieldAny<double, int32_t>>& item) {
	item.resize(0); // TODO - unwrap
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

bool tl2::True::read(::basictl::tl_istream & s) {
	if (!::tl2::details::TrueRead(s, *this)) { return false; }
	return true;
}

bool tl2::True::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::TrueWrite(s, *this)) { return false; }
	return true;
}

bool tl2::True::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::TrueReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::True::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::TrueWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::TrueReset(::tl2::True& item) {
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
