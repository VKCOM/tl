#include "headers/vector.hpp"
#include "headers/tuple.hpp"
#include "headers/true.hpp"
#include "headers/int.hpp"
#include "headers/dictionary.hpp"
#include "headers/dictionaryField.hpp"
#include "headers/Bool.hpp"


bool tl2::details::BoolReadBoxed(::basictl::tl_istream & s, bool& item) {
	return s.bool_read(item, 0xbc799737, 0x997275b5);
}

bool tl2::details::BoolWriteBoxed(::basictl::tl_ostream & s, bool item) {
	return s.nat_write(item ? 0x997275b5 : 0xbc799737);
}

void tl2::details::BuiltinTupleIntReset(std::vector<int32_t>& item) {
	item.resize(0);
}

bool tl2::details::BuiltinTupleIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n) {
	// TODO - check length sanity
	item.resize(nat_n);
	for(auto && el : item) {
		if (!s.int_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTupleIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) {
	if (item.size() != nat_n)
		return s.set_error_sequence_length();
	for(const auto & el : item) {
		if (!s.int_write(el)) { return false;}
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

void tl2::details::BuiltinVectorIntReset(std::vector<int32_t>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!s.int_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!s.int_write(el)) { return false;}
	}
	return true;
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

bool tl2::details::IntMaybeReadBoxed(::basictl::tl_istream & s, std::optional<int32_t>& item) {
	bool has_item = false;
	if (!s.bool_read(has_item, 0x27930a7b, 0x3f9c8ef8)) { return false; }
	if (has_item) {
		if (!item) {
			item.emplace();
		}
		if (!s.int_read(*item)) { return false; }
		return true;
	}
	item.reset();
	return true;
}

bool tl2::details::IntMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<int32_t>& item) {
	if (!s.nat_write(item ? 0x3f9c8ef8 : 0x27930a7b)) { return false; }
	if (item) {
		if (!s.int_write(*item)) { return false;}
	}
	return true;
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

void tl2::details::TupleIntReset(std::vector<int32_t>& item) {
	item.clear();
}

bool tl2::details::TupleIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleIntRead(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleIntWrite(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleIntReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleIntRead(s, item, nat_n);
}

bool tl2::details::TupleIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleIntWrite(s, item, nat_n);
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

void tl2::details::VectorIntReset(std::vector<int32_t>& item) {
	item.clear();
}

bool tl2::details::VectorIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item) {
	if (!::tl2::details::BuiltinVectorIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item) {
	if (!::tl2::details::BuiltinVectorIntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorIntRead(s, item);
}

bool tl2::details::VectorIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorIntWrite(s, item);
}
