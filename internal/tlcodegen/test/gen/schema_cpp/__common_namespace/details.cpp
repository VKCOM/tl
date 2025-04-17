#include "__common_namespace/headers/withFloat.h"
#include "__common_namespace/headers/vector.h"
#include "__common_namespace/headers/tuple.h"
#include "__common_namespace/headers/string.h"
#include "__common_namespace/headers/statOne.h"
#include "service1/headers/service1.Value.h"
#include "__common_namespace/headers/rpcInvokeReqExtra.h"
#include "__common_namespace/headers/true.h"
#include "__common_namespace/headers/nonOptNat.h"
#include "__common_namespace/headers/myTwoDicts.h"
#include "__common_namespace/headers/myMcValueVector.h"
#include "__common_namespace/headers/myMcValueTuple.h"
#include "__common_namespace/headers/myMcValue.h"
#include "__common_namespace/headers/myMaybe2.h"
#include "__common_namespace/headers/myMaybe1.h"
#include "__common_namespace/headers/myTuple.h"
#include "__common_namespace/headers/myBoxedVectorSlice.h"
#include "__common_namespace/headers/myAnonMcValue.h"
#include "__common_namespace/headers/long.h"
#include "__common_namespace/headers/issue3498.h"
#include "service6/headers/service6.findResultRow.h"
#include "service6/headers/service6.error.h"
#include "__common_namespace/headers/int.h"
#include "__common_namespace/headers/getStats.h"
#include "tasks/headers/tasks.queueTypeStats.h"
#include "__common_namespace/headers/getNonOptNat.h"
#include "__common_namespace/headers/getMyValue.h"
#include "__common_namespace/headers/MyValue.h"
#include "__common_namespace/headers/myString.h"
#include "__common_namespace/headers/myInt.h"
#include "__common_namespace/headers/getMyDouble.h"
#include "__common_namespace/headers/myDouble.h"
#include "__common_namespace/headers/getMyDictOfInt.h"
#include "__common_namespace/headers/myDictOfInt.h"
#include "__common_namespace/headers/getMaybeIface.h"
#include "__common_namespace/headers/getFloat.h"
#include "__common_namespace/headers/getDouble.h"
#include "__common_namespace/headers/get_arrays.h"
#include "__common_namespace/headers/float.h"
#include "__common_namespace/headers/fieldConflict4.h"
#include "__common_namespace/headers/fieldConflict3.h"
#include "__common_namespace/headers/fieldConflict2.h"
#include "__common_namespace/headers/fieldConflict1.h"
#include "__common_namespace/headers/Either.h"
#include "__common_namespace/headers/right.h"
#include "__common_namespace/headers/left.h"
#include "__common_namespace/headers/double.h"
#include "__common_namespace/headers/dictionary.h"
#include "__common_namespace/headers/dictionaryField.h"
#include "__common_namespace/headers/boxedVector64.h"
#include "__common_namespace/headers/boxedVector32BoxedElem.h"
#include "__common_namespace/headers/boxedVector32.h"
#include "__common_namespace/headers/boxedTupleSlice3.h"
#include "__common_namespace/headers/boxedTupleSlice2.h"
#include "__common_namespace/headers/myBoxedTupleSlice.h"
#include "__common_namespace/headers/boxedTupleSlice1.h"
#include "__common_namespace/headers/boxedTuple.h"
#include "__common_namespace/headers/boxedString.h"
#include "__common_namespace/headers/boxedInt.h"
#include "__common_namespace/headers/boxedArray.h"
#include "__common_namespace/headers/myBoxedArray.h"
#include "__common_namespace/headers/Bool.h"
#include "__common_namespace/headers/benchObject.h"
#include "__common_namespace/headers/integer.h"


bool tl2::BenchObject::write_json(std::ostream& s)const {
	if (!::tl2::details::BenchObjectWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BenchObject::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchObjectRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BenchObject::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchObjectWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BenchObject::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::BenchObject::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::BenchObject::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BenchObjectReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BenchObject::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BenchObjectWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BenchObject::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::BenchObject::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BenchObjectReset(::tl2::BenchObject& item) noexcept {
	item.xs.clear();
	item.ys.clear();
}

bool tl2::details::BenchObjectWriteJSON(std::ostream& s, const ::tl2::BenchObject& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.xs.size() != 0) {
		add_comma = true;
		s << "\"xs\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.xs)) { return false; }
	}
	if (item.ys.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"ys\":";
		if (!::tl2::details::BuiltinVectorIntegerWriteJSON(s, item.ys)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::BenchObjectRead(::basictl::tl_istream & s, ::tl2::BenchObject& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntRead(s, item.xs)) { return false; }
	if (!::tl2::details::BuiltinVectorIntegerRead(s, item.ys)) { return false; }
	return true;
}

bool tl2::details::BenchObjectWrite(::basictl::tl_ostream & s, const ::tl2::BenchObject& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.xs)) { return false; }
	if (!::tl2::details::BuiltinVectorIntegerWrite(s, item.ys)) { return false; }
	return true;
}

bool tl2::details::BenchObjectReadBoxed(::basictl::tl_istream & s, ::tl2::BenchObject& item) {
	if (!s.nat_read_exact_tag(0xb697e865)) { return false; }
	return tl2::details::BenchObjectRead(s, item);
}

bool tl2::details::BenchObjectWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BenchObject& item) {
	if (!s.nat_write(0xb697e865)) { return false; }
	return tl2::details::BenchObjectWrite(s, item);
}

bool tl2::details::BoolWriteJSON(std::ostream & s, bool item) {
	if (item) {
		s << "true";
	} else {
		s << "false";
	}
	return true;
}

bool tl2::details::BoolReadBoxed(::basictl::tl_istream & s, bool& item) {
	return s.bool_read(item, 0xbc799737, 0x997275b5);
}

bool tl2::details::BoolWriteBoxed(::basictl::tl_ostream & s, bool item) {
	return s.nat_write(item ? 0x997275b5 : 0xbc799737);
}

bool tl2::BoxedArray::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedArrayWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedArray::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedArrayRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedArray::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedArrayWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedArray::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::BoxedArray::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::BoxedArray::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedArrayReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedArray::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedArrayWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedArray::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::BoxedArray::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BoxedArrayReset(::tl2::BoxedArray& item) noexcept {
	::tl2::details::MyBoxedArrayReset(item.x);
}

bool tl2::details::BoxedArrayWriteJSON(std::ostream& s, const ::tl2::BoxedArray& item) noexcept {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::MyBoxedArrayWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BoxedArrayRead(::basictl::tl_istream & s, ::tl2::BoxedArray& item) noexcept {
	if (!::tl2::details::MyBoxedArrayReadBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::BoxedArrayWrite(::basictl::tl_ostream & s, const ::tl2::BoxedArray& item) noexcept {
	if (!::tl2::details::MyBoxedArrayWriteBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::BoxedArrayReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedArray& item) {
	if (!s.nat_read_exact_tag(0x95dcc8b7)) { return false; }
	return tl2::details::BoxedArrayRead(s, item);
}

bool tl2::details::BoxedArrayWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedArray& item) {
	if (!s.nat_write(0x95dcc8b7)) { return false; }
	return tl2::details::BoxedArrayWrite(s, item);
}

bool tl2::details::BoxedArrayReadResult(::basictl::tl_istream & s, tl2::BoxedArray& item, ::tl2::MyBoxedArray& result) {
	if (!::tl2::details::MyBoxedArrayReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::BoxedArrayWriteResult(::basictl::tl_ostream & s, tl2::BoxedArray& item, ::tl2::MyBoxedArray& result) {
	if (!::tl2::details::MyBoxedArrayWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::BoxedArray::read_result(::basictl::tl_istream & s, ::tl2::MyBoxedArray & result) noexcept {
	bool success = tl2::details::BoxedArrayReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::BoxedArray::write_result(::basictl::tl_ostream & s, ::tl2::MyBoxedArray & result) noexcept {
	bool success = tl2::details::BoxedArrayWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::BoxedArray::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::MyBoxedArray & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::BoxedArray::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::MyBoxedArray & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::BoxedInt::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedIntWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedInt::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedIntRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedInt::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedIntWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedInt::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::BoxedInt::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::BoxedInt::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedIntReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedInt::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedIntWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedInt::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::BoxedInt::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BoxedIntReset(::tl2::BoxedInt& item) noexcept {
	item.x = 0;
}

bool tl2::details::BoxedIntWriteJSON(std::ostream& s, const ::tl2::BoxedInt& item) noexcept {
	s << "{";
	if (item.x != 0) {
		s << "\"x\":";
		s << item.x;
	}
	s << "}";
	return true;
}

bool tl2::details::BoxedIntRead(::basictl::tl_istream & s, ::tl2::BoxedInt& item) noexcept {
	if (!s.nat_read_exact_tag(0xa8509bda)) { return false; }
	if (!s.int_read(item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedIntWrite(::basictl::tl_ostream & s, const ::tl2::BoxedInt& item) noexcept {
	if (!s.nat_write(0xa8509bda)) { return false; }
	if (!s.int_write(item.x)) { return false;}
	return true;
}

bool tl2::details::BoxedIntReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedInt& item) {
	if (!s.nat_read_exact_tag(0x5688ebaf)) { return false; }
	return tl2::details::BoxedIntRead(s, item);
}

bool tl2::details::BoxedIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedInt& item) {
	if (!s.nat_write(0x5688ebaf)) { return false; }
	return tl2::details::BoxedIntWrite(s, item);
}

bool tl2::details::BoxedIntReadResult(::basictl::tl_istream & s, tl2::BoxedInt& item, int32_t& result) {
	if (!s.nat_read_exact_tag(0xa8509bda)) { return false; }
	if (!s.int_read(result)) { return false; }
	return true;
}
bool tl2::details::BoxedIntWriteResult(::basictl::tl_ostream & s, tl2::BoxedInt& item, int32_t& result) {
	if (!s.nat_write(0xa8509bda)) { return false; }
	if (!s.int_write(result)) { return false;}
	return true;
}

bool tl2::BoxedInt::read_result(::basictl::tl_istream & s, int32_t & result) noexcept {
	bool success = tl2::details::BoxedIntReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::BoxedInt::write_result(::basictl::tl_ostream & s, int32_t & result) noexcept {
	bool success = tl2::details::BoxedIntWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::BoxedInt::read_result_or_throw(::basictl::tl_throwable_istream & s, int32_t & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::BoxedInt::write_result_or_throw(::basictl::tl_throwable_ostream & s, int32_t & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::BoxedString::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedStringWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedString::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedStringRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedString::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedStringWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedString::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::BoxedString::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::BoxedString::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedStringReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedString::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedStringWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedString::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::BoxedString::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BoxedStringReset(::tl2::BoxedString& item) noexcept {
	item.x.clear();
}

bool tl2::details::BoxedStringWriteJSON(std::ostream& s, const ::tl2::BoxedString& item) noexcept {
	s << "{";
	if (item.x.size() != 0) {
		s << "\"x\":";
		s << "\"" << item.x << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::BoxedStringRead(::basictl::tl_istream & s, ::tl2::BoxedString& item) noexcept {
	if (!s.nat_read_exact_tag(0xb5286e24)) { return false; }
	if (!s.string_read(item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedStringWrite(::basictl::tl_ostream & s, const ::tl2::BoxedString& item) noexcept {
	if (!s.nat_write(0xb5286e24)) { return false; }
	if (!s.string_write(item.x)) { return false;}
	return true;
}

bool tl2::details::BoxedStringReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedString& item) {
	if (!s.nat_read_exact_tag(0x548994db)) { return false; }
	return tl2::details::BoxedStringRead(s, item);
}

bool tl2::details::BoxedStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedString& item) {
	if (!s.nat_write(0x548994db)) { return false; }
	return tl2::details::BoxedStringWrite(s, item);
}

bool tl2::details::BoxedStringReadResult(::basictl::tl_istream & s, tl2::BoxedString& item, std::string& result) {
	if (!s.nat_read_exact_tag(0xb5286e24)) { return false; }
	if (!s.string_read(result)) { return false; }
	return true;
}
bool tl2::details::BoxedStringWriteResult(::basictl::tl_ostream & s, tl2::BoxedString& item, std::string& result) {
	if (!s.nat_write(0xb5286e24)) { return false; }
	if (!s.string_write(result)) { return false;}
	return true;
}

bool tl2::BoxedString::read_result(::basictl::tl_istream & s, std::string & result) noexcept {
	bool success = tl2::details::BoxedStringReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::BoxedString::write_result(::basictl::tl_ostream & s, std::string & result) noexcept {
	bool success = tl2::details::BoxedStringWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::BoxedString::read_result_or_throw(::basictl::tl_throwable_istream & s, std::string & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::BoxedString::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::string & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::BoxedTuple::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedTupleWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTuple::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedTupleRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedTuple::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedTupleWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedTuple::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::BoxedTuple::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::BoxedTuple::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedTupleReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedTuple::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedTupleWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedTuple::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::BoxedTuple::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BoxedTupleReset(::tl2::BoxedTuple& item) noexcept {
	::tl2::details::BuiltinTuple3IntReset(item.x);
}

bool tl2::details::BoxedTupleWriteJSON(std::ostream& s, const ::tl2::BoxedTuple& item) noexcept {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::BuiltinTuple3IntWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BoxedTupleRead(::basictl::tl_istream & s, ::tl2::BoxedTuple& item) noexcept {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTuple3IntRead(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedTupleWrite(::basictl::tl_ostream & s, const ::tl2::BoxedTuple& item) noexcept {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTuple3IntWrite(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedTupleReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedTuple& item) {
	if (!s.nat_read_exact_tag(0x30c9d533)) { return false; }
	return tl2::details::BoxedTupleRead(s, item);
}

bool tl2::details::BoxedTupleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedTuple& item) {
	if (!s.nat_write(0x30c9d533)) { return false; }
	return tl2::details::BoxedTupleWrite(s, item);
}

bool tl2::details::BoxedTupleReadResult(::basictl::tl_istream & s, tl2::BoxedTuple& item, std::array<int32_t, 3>& result) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTuple3IntRead(s, result)) { return false; }
	return true;
}
bool tl2::details::BoxedTupleWriteResult(::basictl::tl_ostream & s, tl2::BoxedTuple& item, std::array<int32_t, 3>& result) {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTuple3IntWrite(s, result)) { return false; }
	return true;
}

bool tl2::BoxedTuple::read_result(::basictl::tl_istream & s, std::array<int32_t, 3> & result) noexcept {
	bool success = tl2::details::BoxedTupleReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::BoxedTuple::write_result(::basictl::tl_ostream & s, std::array<int32_t, 3> & result) noexcept {
	bool success = tl2::details::BoxedTupleWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::BoxedTuple::read_result_or_throw(::basictl::tl_throwable_istream & s, std::array<int32_t, 3> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::BoxedTuple::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::array<int32_t, 3> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::BoxedTupleSlice1::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedTupleSlice1WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice1::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedTupleSlice1Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedTupleSlice1::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedTupleSlice1Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedTupleSlice1::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::BoxedTupleSlice1::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::BoxedTupleSlice1::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedTupleSlice1ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedTupleSlice1::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedTupleSlice1WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedTupleSlice1::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::BoxedTupleSlice1::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BoxedTupleSlice1Reset(::tl2::BoxedTupleSlice1& item) noexcept {
	item.n = 0;
	item.x.clear();
}

bool tl2::details::BoxedTupleSlice1WriteJSON(std::ostream& s, const ::tl2::BoxedTupleSlice1& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.n != 0) {
		add_comma = true;
		s << "\"n\":";
		s << item.n;
	}
	if ((item.x.size() != 0) || (item.n != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"x\":";
		if (!::tl2::details::BuiltinTupleIntBoxedWriteJSON(s, item.x, item.n)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::BoxedTupleSlice1Read(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice1& item) noexcept {
	if (!s.nat_read(item.n)) { return false; }
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntBoxedRead(s, item.x, item.n)) { return false; }
	return true;
}

bool tl2::details::BoxedTupleSlice1Write(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice1& item) noexcept {
	if (!s.nat_write(item.n)) { return false;}
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntBoxedWrite(s, item.x, item.n)) { return false; }
	return true;
}

bool tl2::details::BoxedTupleSlice1ReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice1& item) {
	if (!s.nat_read_exact_tag(0x25230d40)) { return false; }
	return tl2::details::BoxedTupleSlice1Read(s, item);
}

bool tl2::details::BoxedTupleSlice1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice1& item) {
	if (!s.nat_write(0x25230d40)) { return false; }
	return tl2::details::BoxedTupleSlice1Write(s, item);
}

bool tl2::details::BoxedTupleSlice1ReadResult(::basictl::tl_istream & s, tl2::BoxedTupleSlice1& item, std::vector<int32_t>& result) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntBoxedRead(s, result, item.n)) { return false; }
	return true;
}
bool tl2::details::BoxedTupleSlice1WriteResult(::basictl::tl_ostream & s, tl2::BoxedTupleSlice1& item, std::vector<int32_t>& result) {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntBoxedWrite(s, result, item.n)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice1::read_result(::basictl::tl_istream & s, std::vector<int32_t> & result) noexcept {
	bool success = tl2::details::BoxedTupleSlice1ReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::BoxedTupleSlice1::write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result) noexcept {
	bool success = tl2::details::BoxedTupleSlice1WriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::BoxedTupleSlice1::read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<int32_t> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::BoxedTupleSlice1::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<int32_t> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::BoxedTupleSlice2::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedTupleSlice2WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice2::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedTupleSlice2Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedTupleSlice2::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedTupleSlice2Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedTupleSlice2::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::BoxedTupleSlice2::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::BoxedTupleSlice2::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedTupleSlice2ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedTupleSlice2::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedTupleSlice2WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedTupleSlice2::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::BoxedTupleSlice2::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BoxedTupleSlice2Reset(::tl2::BoxedTupleSlice2& item) noexcept {
	::tl2::details::MyBoxedTupleSliceReset(item.x);
}

bool tl2::details::BoxedTupleSlice2WriteJSON(std::ostream& s, const ::tl2::BoxedTupleSlice2& item) noexcept {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::MyBoxedTupleSliceWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BoxedTupleSlice2Read(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice2& item) noexcept {
	if (!::tl2::details::MyBoxedTupleSliceReadBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::BoxedTupleSlice2Write(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice2& item) noexcept {
	if (!::tl2::details::MyBoxedTupleSliceWriteBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::BoxedTupleSlice2ReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice2& item) {
	if (!s.nat_read_exact_tag(0x1cdf4705)) { return false; }
	return tl2::details::BoxedTupleSlice2Read(s, item);
}

bool tl2::details::BoxedTupleSlice2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice2& item) {
	if (!s.nat_write(0x1cdf4705)) { return false; }
	return tl2::details::BoxedTupleSlice2Write(s, item);
}

bool tl2::details::BoxedTupleSlice2ReadResult(::basictl::tl_istream & s, tl2::BoxedTupleSlice2& item, ::tl2::MyBoxedTupleSlice& result) {
	if (!::tl2::details::MyBoxedTupleSliceReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::BoxedTupleSlice2WriteResult(::basictl::tl_ostream & s, tl2::BoxedTupleSlice2& item, ::tl2::MyBoxedTupleSlice& result) {
	if (!::tl2::details::MyBoxedTupleSliceWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::BoxedTupleSlice2::read_result(::basictl::tl_istream & s, ::tl2::MyBoxedTupleSlice & result) noexcept {
	bool success = tl2::details::BoxedTupleSlice2ReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::BoxedTupleSlice2::write_result(::basictl::tl_ostream & s, ::tl2::MyBoxedTupleSlice & result) noexcept {
	bool success = tl2::details::BoxedTupleSlice2WriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::BoxedTupleSlice2::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::MyBoxedTupleSlice & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::BoxedTupleSlice2::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::MyBoxedTupleSlice & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::BoxedTupleSlice3::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedTupleSlice3WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice3::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedTupleSlice3Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedTupleSlice3::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedTupleSlice3Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedTupleSlice3::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::BoxedTupleSlice3::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::BoxedTupleSlice3::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedTupleSlice3ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedTupleSlice3::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedTupleSlice3WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedTupleSlice3::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::BoxedTupleSlice3::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BoxedTupleSlice3Reset(::tl2::BoxedTupleSlice3& item) noexcept {
	item.n = 0;
	item.x.clear();
}

bool tl2::details::BoxedTupleSlice3WriteJSON(std::ostream& s, const ::tl2::BoxedTupleSlice3& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.n != 0) {
		add_comma = true;
		s << "\"n\":";
		s << item.n;
	}
	if ((item.x.size() != 0) || (item.n != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"x\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.x, item.n)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::BoxedTupleSlice3Read(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice3& item) noexcept {
	if (!s.nat_read(item.n)) { return false; }
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.x, item.n)) { return false; }
	return true;
}

bool tl2::details::BoxedTupleSlice3Write(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice3& item) noexcept {
	if (!s.nat_write(item.n)) { return false;}
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.x, item.n)) { return false; }
	return true;
}

bool tl2::details::BoxedTupleSlice3ReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice3& item) {
	if (!s.nat_read_exact_tag(0xa19b8106)) { return false; }
	return tl2::details::BoxedTupleSlice3Read(s, item);
}

bool tl2::details::BoxedTupleSlice3WriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice3& item) {
	if (!s.nat_write(0xa19b8106)) { return false; }
	return tl2::details::BoxedTupleSlice3Write(s, item);
}

bool tl2::details::BoxedTupleSlice3ReadResult(::basictl::tl_istream & s, tl2::BoxedTupleSlice3& item, std::vector<int32_t>& result) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, result, item.n)) { return false; }
	return true;
}
bool tl2::details::BoxedTupleSlice3WriteResult(::basictl::tl_ostream & s, tl2::BoxedTupleSlice3& item, std::vector<int32_t>& result) {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntWrite(s, result, item.n)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice3::read_result(::basictl::tl_istream & s, std::vector<int32_t> & result) noexcept {
	bool success = tl2::details::BoxedTupleSlice3ReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::BoxedTupleSlice3::write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result) noexcept {
	bool success = tl2::details::BoxedTupleSlice3WriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::BoxedTupleSlice3::read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<int32_t> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::BoxedTupleSlice3::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<int32_t> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::BoxedVector32::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedVector32WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector32::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedVector32Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedVector32::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedVector32Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedVector32::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::BoxedVector32::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::BoxedVector32::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedVector32ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedVector32::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedVector32WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedVector32::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::BoxedVector32::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BoxedVector32Reset(::tl2::BoxedVector32& item) noexcept {
	item.x.clear();
}

bool tl2::details::BoxedVector32WriteJSON(std::ostream& s, const ::tl2::BoxedVector32& item) noexcept {
	s << "{";
	if (item.x.size() != 0) {
		s << "\"x\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.x)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::BoxedVector32Read(::basictl::tl_istream & s, ::tl2::BoxedVector32& item) noexcept {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedVector32Write(::basictl::tl_ostream & s, const ::tl2::BoxedVector32& item) noexcept {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedVector32ReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedVector32& item) {
	if (!s.nat_read_exact_tag(0xbbadef07)) { return false; }
	return tl2::details::BoxedVector32Read(s, item);
}

bool tl2::details::BoxedVector32WriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedVector32& item) {
	if (!s.nat_write(0xbbadef07)) { return false; }
	return tl2::details::BoxedVector32Write(s, item);
}

bool tl2::details::BoxedVector32ReadResult(::basictl::tl_istream & s, tl2::BoxedVector32& item, std::vector<int32_t>& result) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, result)) { return false; }
	return true;
}
bool tl2::details::BoxedVector32WriteResult(::basictl::tl_ostream & s, tl2::BoxedVector32& item, std::vector<int32_t>& result) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntWrite(s, result)) { return false; }
	return true;
}

bool tl2::BoxedVector32::read_result(::basictl::tl_istream & s, std::vector<int32_t> & result) noexcept {
	bool success = tl2::details::BoxedVector32ReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::BoxedVector32::write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result) noexcept {
	bool success = tl2::details::BoxedVector32WriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::BoxedVector32::read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<int32_t> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::BoxedVector32::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<int32_t> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::BoxedVector32BoxedElem::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedVector32BoxedElemWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector32BoxedElem::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedVector32BoxedElemRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedVector32BoxedElem::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedVector32BoxedElemWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedVector32BoxedElem::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::BoxedVector32BoxedElem::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::BoxedVector32BoxedElem::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedVector32BoxedElemReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedVector32BoxedElem::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedVector32BoxedElemWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedVector32BoxedElem::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::BoxedVector32BoxedElem::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BoxedVector32BoxedElemReset(::tl2::BoxedVector32BoxedElem& item) noexcept {
	item.x.clear();
}

bool tl2::details::BoxedVector32BoxedElemWriteJSON(std::ostream& s, const ::tl2::BoxedVector32BoxedElem& item) noexcept {
	s << "{";
	if (item.x.size() != 0) {
		s << "\"x\":";
		if (!::tl2::details::BuiltinVectorIntBoxedWriteJSON(s, item.x)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::BoxedVector32BoxedElemRead(::basictl::tl_istream & s, ::tl2::BoxedVector32BoxedElem& item) noexcept {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntBoxedRead(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedVector32BoxedElemWrite(::basictl::tl_ostream & s, const ::tl2::BoxedVector32BoxedElem& item) noexcept {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntBoxedWrite(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedVector32BoxedElemReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedVector32BoxedElem& item) {
	if (!s.nat_read_exact_tag(0x591cecd4)) { return false; }
	return tl2::details::BoxedVector32BoxedElemRead(s, item);
}

bool tl2::details::BoxedVector32BoxedElemWriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedVector32BoxedElem& item) {
	if (!s.nat_write(0x591cecd4)) { return false; }
	return tl2::details::BoxedVector32BoxedElemWrite(s, item);
}

bool tl2::details::BoxedVector32BoxedElemReadResult(::basictl::tl_istream & s, tl2::BoxedVector32BoxedElem& item, std::vector<int32_t>& result) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntBoxedRead(s, result)) { return false; }
	return true;
}
bool tl2::details::BoxedVector32BoxedElemWriteResult(::basictl::tl_ostream & s, tl2::BoxedVector32BoxedElem& item, std::vector<int32_t>& result) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntBoxedWrite(s, result)) { return false; }
	return true;
}

bool tl2::BoxedVector32BoxedElem::read_result(::basictl::tl_istream & s, std::vector<int32_t> & result) noexcept {
	bool success = tl2::details::BoxedVector32BoxedElemReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::BoxedVector32BoxedElem::write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result) noexcept {
	bool success = tl2::details::BoxedVector32BoxedElemWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::BoxedVector32BoxedElem::read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<int32_t> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::BoxedVector32BoxedElem::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<int32_t> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::BoxedVector64::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedVector64WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector64::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedVector64Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedVector64::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedVector64Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedVector64::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::BoxedVector64::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::BoxedVector64::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::BoxedVector64ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::BoxedVector64::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::BoxedVector64WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::BoxedVector64::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::BoxedVector64::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::BoxedVector64Reset(::tl2::BoxedVector64& item) noexcept {
	item.x.clear();
}

bool tl2::details::BoxedVector64WriteJSON(std::ostream& s, const ::tl2::BoxedVector64& item) noexcept {
	s << "{";
	if (item.x.size() != 0) {
		s << "\"x\":";
		if (!::tl2::details::BuiltinVectorLongBoxedWriteJSON(s, item.x)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::BoxedVector64Read(::basictl::tl_istream & s, ::tl2::BoxedVector64& item) noexcept {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorLongBoxedRead(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedVector64Write(::basictl::tl_ostream & s, const ::tl2::BoxedVector64& item) noexcept {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorLongBoxedWrite(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedVector64ReadBoxed(::basictl::tl_istream & s, ::tl2::BoxedVector64& item) {
	if (!s.nat_read_exact_tag(0x83659ba8)) { return false; }
	return tl2::details::BoxedVector64Read(s, item);
}

bool tl2::details::BoxedVector64WriteBoxed(::basictl::tl_ostream & s, const ::tl2::BoxedVector64& item) {
	if (!s.nat_write(0x83659ba8)) { return false; }
	return tl2::details::BoxedVector64Write(s, item);
}

bool tl2::details::BoxedVector64ReadResult(::basictl::tl_istream & s, tl2::BoxedVector64& item, std::vector<int64_t>& result) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorLongBoxedRead(s, result)) { return false; }
	return true;
}
bool tl2::details::BoxedVector64WriteResult(::basictl::tl_ostream & s, tl2::BoxedVector64& item, std::vector<int64_t>& result) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorLongBoxedWrite(s, result)) { return false; }
	return true;
}

bool tl2::BoxedVector64::read_result(::basictl::tl_istream & s, std::vector<int64_t> & result) noexcept {
	bool success = tl2::details::BoxedVector64ReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::BoxedVector64::write_result(::basictl::tl_ostream & s, std::vector<int64_t> & result) noexcept {
	bool success = tl2::details::BoxedVector64WriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::BoxedVector64::read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<int64_t> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::BoxedVector64::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<int64_t> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

void tl2::details::BuiltinTuple10IntBoxedReset(std::array<int32_t, 10>& item) {
	for(auto && el : item) {
		el = 0;
	}
}

bool tl2::details::BuiltinTuple10IntBoxedWriteJSON(std::ostream &s, const std::array<int32_t, 10>& item) {
	s << "[";
	size_t index = 0;
	for(auto && el : item) {
		s << el;
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTuple10IntBoxedRead(::basictl::tl_istream & s, std::array<int32_t, 10>& item) {
	for(auto && el : item) {
		if (!s.nat_read_exact_tag(0xa8509bda)) { return false; }
	if (!s.int_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTuple10IntBoxedWrite(::basictl::tl_ostream & s, const std::array<int32_t, 10>& item) {
	for(const auto & el : item) {
		if (!s.nat_write(0xa8509bda)) { return false; }
	if (!s.int_write(el)) { return false;}
	}
	return true;
}

void tl2::details::BuiltinTuple2IntBoxedReset(std::array<int32_t, 2>& item) {
	for(auto && el : item) {
		el = 0;
	}
}

bool tl2::details::BuiltinTuple2IntBoxedWriteJSON(std::ostream &s, const std::array<int32_t, 2>& item) {
	s << "[";
	size_t index = 0;
	for(auto && el : item) {
		s << el;
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTuple2IntBoxedRead(::basictl::tl_istream & s, std::array<int32_t, 2>& item) {
	for(auto && el : item) {
		if (!s.nat_read_exact_tag(0xa8509bda)) { return false; }
	if (!s.int_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTuple2IntBoxedWrite(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item) {
	for(const auto & el : item) {
		if (!s.nat_write(0xa8509bda)) { return false; }
	if (!s.int_write(el)) { return false;}
	}
	return true;
}

void tl2::details::BuiltinTuple3IntReset(std::array<int32_t, 3>& item) {
	for(auto && el : item) {
		el = 0;
	}
}

bool tl2::details::BuiltinTuple3IntWriteJSON(std::ostream &s, const std::array<int32_t, 3>& item) {
	s << "[";
	size_t index = 0;
	for(auto && el : item) {
		s << el;
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTuple3IntRead(::basictl::tl_istream & s, std::array<int32_t, 3>& item) {
	for(auto && el : item) {
		if (!s.int_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTuple3IntWrite(::basictl::tl_ostream & s, const std::array<int32_t, 3>& item) {
	for(const auto & el : item) {
		if (!s.int_write(el)) { return false;}
	}
	return true;
}

void tl2::details::BuiltinTuple5IntReset(std::array<int32_t, 5>& item) {
	for(auto && el : item) {
		el = 0;
	}
}

bool tl2::details::BuiltinTuple5IntWriteJSON(std::ostream &s, const std::array<int32_t, 5>& item) {
	s << "[";
	size_t index = 0;
	for(auto && el : item) {
		s << el;
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTuple5IntRead(::basictl::tl_istream & s, std::array<int32_t, 5>& item) {
	for(auto && el : item) {
		if (!s.int_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTuple5IntWrite(::basictl::tl_ostream & s, const std::array<int32_t, 5>& item) {
	for(const auto & el : item) {
		if (!s.int_write(el)) { return false;}
	}
	return true;
}

void tl2::details::BuiltinTupleIntReset(std::vector<int32_t>& item) {
	item.resize(0);
}

bool tl2::details::BuiltinTupleIntWriteJSON(std::ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) {
	if (item.size() != nat_n) { return false; }
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		s << el;
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
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
	if (item.size() != nat_n) { return s.set_error_sequence_length(); }
	for(const auto & el : item) {
		if (!s.int_write(el)) { return false;}
	}
	return true;
}

void tl2::details::BuiltinTupleIntBoxedReset(std::vector<int32_t>& item) {
	item.resize(0);
}

bool tl2::details::BuiltinTupleIntBoxedWriteJSON(std::ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) {
	if (item.size() != nat_n) { return false; }
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		s << el;
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTupleIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n) {
	// TODO - check length sanity
	item.resize(nat_n);
	for(auto && el : item) {
		if (!s.nat_read_exact_tag(0xa8509bda)) { return false; }
	if (!s.int_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTupleIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) {
	if (item.size() != nat_n) { return s.set_error_sequence_length(); }
	for(const auto & el : item) {
		if (!s.nat_write(0xa8509bda)) { return false; }
	if (!s.int_write(el)) { return false;}
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

void tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowReset(std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowWriteJSON(std::ostream & s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::EitherService6ErrorVectorService6FindResultRowWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::EitherService6ErrorVectorService6FindResultRowReadBoxed(s, el)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

bool tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::EitherService6ErrorVectorService6FindResultRowWriteBoxed(s, el)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

void tl2::details::BuiltinVectorIntReset(std::vector<int32_t>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorIntWriteJSON(std::ostream & s, const std::vector<int32_t>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		s << el;
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
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

void tl2::details::BuiltinVectorIntBoxedReset(std::vector<int32_t>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorIntBoxedWriteJSON(std::ostream & s, const std::vector<int32_t>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		s << el;
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!s.nat_read_exact_tag(0xa8509bda)) { return false; }
	if (!s.int_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!s.nat_write(0xa8509bda)) { return false; }
	if (!s.int_write(el)) { return false;}
	}
	return true;
}

void tl2::details::BuiltinVectorIntegerReset(std::vector<::tl2::Integer>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorIntegerWriteJSON(std::ostream & s, const std::vector<::tl2::Integer>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::IntegerWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorIntegerRead(::basictl::tl_istream & s, std::vector<::tl2::Integer>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::IntegerRead(s, el)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

bool tl2::details::BuiltinVectorIntegerWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Integer>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::IntegerWrite(s, el)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

void tl2::details::BuiltinVectorLongBoxedReset(std::vector<int64_t>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorLongBoxedWriteJSON(std::ostream & s, const std::vector<int64_t>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		s << el;
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorLongBoxedRead(::basictl::tl_istream & s, std::vector<int64_t>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!s.nat_read_exact_tag(0x22076cba)) { return false; }
	if (!s.long_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorLongBoxedWrite(::basictl::tl_ostream & s, const std::vector<int64_t>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!s.nat_write(0x22076cba)) { return false; }
	if (!s.long_write(el)) { return false;}
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

void tl2::details::DictionaryFieldDictionaryIntReset(::tl2::DictionaryField<std::map<std::string, int32_t>>& item) noexcept {
	item.key.clear();
	item.value.clear();
}

bool tl2::details::DictionaryFieldDictionaryIntWriteJSON(std::ostream& s, const ::tl2::DictionaryField<std::map<std::string, int32_t>>& item) noexcept {
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
		if (!::tl2::details::BuiltinVectorDictionaryFieldIntWriteJSON(s, item.value)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::DictionaryFieldDictionaryIntRead(::basictl::tl_istream & s, ::tl2::DictionaryField<std::map<std::string, int32_t>>& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntRead(s, item.value)) { return false; }
	return true;
}

bool tl2::details::DictionaryFieldDictionaryIntWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::map<std::string, int32_t>>& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWrite(s, item.value)) { return false; }
	return true;
}

bool tl2::details::DictionaryFieldDictionaryIntReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<std::map<std::string, int32_t>>& item) {
	if (!s.nat_read_exact_tag(0x239c1b62)) { return false; }
	return tl2::details::DictionaryFieldDictionaryIntRead(s, item);
}

bool tl2::details::DictionaryFieldDictionaryIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<std::map<std::string, int32_t>>& item) {
	if (!s.nat_write(0x239c1b62)) { return false; }
	return tl2::details::DictionaryFieldDictionaryIntWrite(s, item);
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

void tl2::details::DictionaryFieldService1ValueReset(::tl2::DictionaryField<::tl2::service1::Value>& item) noexcept {
	item.key.clear();
	::tl2::details::Service1ValueReset(item.value);
}

bool tl2::details::DictionaryFieldService1ValueWriteJSON(std::ostream& s, const ::tl2::DictionaryField<::tl2::service1::Value>& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.key.size() != 0) {
		add_comma = true;
		s << "\"key\":";
		s << "\"" << item.key << "\"";
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"value\":";
	if (!::tl2::details::Service1ValueWriteJSON(s, item.value)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::DictionaryFieldService1ValueRead(::basictl::tl_istream & s, ::tl2::DictionaryField<::tl2::service1::Value>& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!::tl2::details::Service1ValueReadBoxed(s, item.value)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::DictionaryFieldService1ValueWrite(::basictl::tl_ostream & s, const ::tl2::DictionaryField<::tl2::service1::Value>& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!::tl2::details::Service1ValueWriteBoxed(s, item.value)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::DictionaryFieldService1ValueReadBoxed(::basictl::tl_istream & s, ::tl2::DictionaryField<::tl2::service1::Value>& item) {
	if (!s.nat_read_exact_tag(0x239c1b62)) { return false; }
	return tl2::details::DictionaryFieldService1ValueRead(s, item);
}

bool tl2::details::DictionaryFieldService1ValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::DictionaryField<::tl2::service1::Value>& item) {
	if (!s.nat_write(0x239c1b62)) { return false; }
	return tl2::details::DictionaryFieldService1ValueWrite(s, item);
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

void tl2::details::DoubleReset(double& item) noexcept {
	item = 0;
}

bool tl2::details::DoubleWriteJSON(std::ostream& s, const double& item) noexcept {
	s << item;
	return true;
}

bool tl2::details::DoubleRead(::basictl::tl_istream & s, double& item) noexcept {
	if (!s.double_read(item)) { return false; }
	return true;
}

bool tl2::details::DoubleWrite(::basictl::tl_ostream & s, const double& item) noexcept {
	if (!s.double_write(item)) { return false;}
	return true;
}

bool tl2::details::DoubleReadBoxed(::basictl::tl_istream & s, double& item) {
	if (!s.nat_read_exact_tag(0x2210c154)) { return false; }
	return tl2::details::DoubleRead(s, item);
}

bool tl2::details::DoubleWriteBoxed(::basictl::tl_ostream & s, const double& item) {
	if (!s.nat_write(0x2210c154)) { return false; }
	return tl2::details::DoubleWrite(s, item);
}

static const std::string_view EitherService6ErrorVectorService6FindResultRow_tbl_tl_name[]{"left", "right"};
static const uint32_t EitherService6ErrorVectorService6FindResultRow_tbl_tl_tag[]{0x0a29cd5d, 0xdf3ecb3b};

void tl2::details::EitherService6ErrorVectorService6FindResultRowReset(::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept{
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::EitherService6ErrorVectorService6FindResultRowWriteJSON(std::ostream & s, const ::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept {
	s << "{";
	s << "\"type\":";
	s << "\"" << EitherService6ErrorVectorService6FindResultRow_tbl_tl_name[item.value.index()] << "\"";
	switch (item.value.index()) {
	case 0:
		s << ",\"value\":";
		if (!::tl2::details::LeftService6ErrorVectorService6FindResultRowWriteJSON(s, std::get<0>(item.value))) { return false; }
		break;
	case 1:
		s << ",\"value\":";
		if (!::tl2::details::RightService6ErrorVectorService6FindResultRowWriteJSON(s, std::get<1>(item.value))) { return false; }
		break;
	}
	s << "}";
	return true;
}
bool tl2::details::EitherService6ErrorVectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, ::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept {
	uint32_t nat;
	if (!s.nat_read(nat)) { return false; }
	switch (nat) {
	case 0x0a29cd5d:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		if (!::tl2::details::LeftService6ErrorVectorService6FindResultRowRead(s, std::get<0>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	case 0xdf3ecb3b:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		if (!::tl2::details::RightService6ErrorVectorService6FindResultRowRead(s, std::get<1>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::EitherService6ErrorVectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept{
	if (!s.nat_write(EitherService6ErrorVectorService6FindResultRow_tbl_tl_tag[item.value.index()])) { return false; }
	switch (item.value.index()) {
	case 0:
		if (!::tl2::details::LeftService6ErrorVectorService6FindResultRowWrite(s, std::get<0>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	case 1:
		if (!::tl2::details::RightService6ErrorVectorService6FindResultRowWrite(s, std::get<1>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	}
	return true;
}

bool tl2::FieldConflict1::write_json(std::ostream& s)const {
	if (!::tl2::details::FieldConflict1WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict1::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::FieldConflict1Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::FieldConflict1::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::FieldConflict1Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::FieldConflict1::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::FieldConflict1::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::FieldConflict1::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::FieldConflict1ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::FieldConflict1::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::FieldConflict1WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::FieldConflict1::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::FieldConflict1::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::FieldConflict1Reset(::tl2::FieldConflict1& item) noexcept {
	item.x = 0;
	item.set_x = 0;
}

bool tl2::details::FieldConflict1WriteJSON(std::ostream& s, const ::tl2::FieldConflict1& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.x != 0) {
		add_comma = true;
		s << "\"x\":";
		s << item.x;
	}
	if (item.set_x != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"set_x\":";
		s << item.set_x;
	}
	s << "}";
	return true;
}

bool tl2::details::FieldConflict1Read(::basictl::tl_istream & s, ::tl2::FieldConflict1& item) noexcept {
	if (!s.int_read(item.x)) { return false; }
	if (!s.int_read(item.set_x)) { return false; }
	return true;
}

bool tl2::details::FieldConflict1Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict1& item) noexcept {
	if (!s.int_write(item.x)) { return false;}
	if (!s.int_write(item.set_x)) { return false;}
	return true;
}

bool tl2::details::FieldConflict1ReadBoxed(::basictl::tl_istream & s, ::tl2::FieldConflict1& item) {
	if (!s.nat_read_exact_tag(0xf314bd09)) { return false; }
	return tl2::details::FieldConflict1Read(s, item);
}

bool tl2::details::FieldConflict1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::FieldConflict1& item) {
	if (!s.nat_write(0xf314bd09)) { return false; }
	return tl2::details::FieldConflict1Write(s, item);
}

bool tl2::FieldConflict2::write_json(std::ostream& s)const {
	if (!::tl2::details::FieldConflict2WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict2::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::FieldConflict2Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::FieldConflict2::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::FieldConflict2Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::FieldConflict2::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::FieldConflict2::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::FieldConflict2::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::FieldConflict2ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::FieldConflict2::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::FieldConflict2WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::FieldConflict2::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::FieldConflict2::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::FieldConflict2Reset(::tl2::FieldConflict2& item) noexcept {
	item.x = 0;
	item.setX = 0;
}

bool tl2::details::FieldConflict2WriteJSON(std::ostream& s, const ::tl2::FieldConflict2& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.x != 0) {
		add_comma = true;
		s << "\"x\":";
		s << item.x;
	}
	if (item.setX != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"setX\":";
		s << item.setX;
	}
	s << "}";
	return true;
}

bool tl2::details::FieldConflict2Read(::basictl::tl_istream & s, ::tl2::FieldConflict2& item) noexcept {
	if (!s.int_read(item.x)) { return false; }
	if (!s.int_read(item.setX)) { return false; }
	return true;
}

bool tl2::details::FieldConflict2Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict2& item) noexcept {
	if (!s.int_write(item.x)) { return false;}
	if (!s.int_write(item.setX)) { return false;}
	return true;
}

bool tl2::details::FieldConflict2ReadBoxed(::basictl::tl_istream & s, ::tl2::FieldConflict2& item) {
	if (!s.nat_read_exact_tag(0x1bba76b8)) { return false; }
	return tl2::details::FieldConflict2Read(s, item);
}

bool tl2::details::FieldConflict2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::FieldConflict2& item) {
	if (!s.nat_write(0x1bba76b8)) { return false; }
	return tl2::details::FieldConflict2Write(s, item);
}

bool tl2::FieldConflict3::write_json(std::ostream& s)const {
	if (!::tl2::details::FieldConflict3WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict3::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::FieldConflict3Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::FieldConflict3::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::FieldConflict3Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::FieldConflict3::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::FieldConflict3::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::FieldConflict3::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::FieldConflict3ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::FieldConflict3::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::FieldConflict3WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::FieldConflict3::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::FieldConflict3::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::FieldConflict3Reset(::tl2::FieldConflict3& item) noexcept {
	item.x = 0;
	item.SetX = 0;
}

bool tl2::details::FieldConflict3WriteJSON(std::ostream& s, const ::tl2::FieldConflict3& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.x != 0) {
		add_comma = true;
		s << "\"x\":";
		s << item.x;
	}
	if (item.SetX != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"SetX\":";
		s << item.SetX;
	}
	s << "}";
	return true;
}

bool tl2::details::FieldConflict3Read(::basictl::tl_istream & s, ::tl2::FieldConflict3& item) noexcept {
	if (!s.int_read(item.x)) { return false; }
	if (!s.int_read(item.SetX)) { return false; }
	return true;
}

bool tl2::details::FieldConflict3Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict3& item) noexcept {
	if (!s.int_write(item.x)) { return false;}
	if (!s.int_write(item.SetX)) { return false;}
	return true;
}

bool tl2::details::FieldConflict3ReadBoxed(::basictl::tl_istream & s, ::tl2::FieldConflict3& item) {
	if (!s.nat_read_exact_tag(0x2cf6e157)) { return false; }
	return tl2::details::FieldConflict3Read(s, item);
}

bool tl2::details::FieldConflict3WriteBoxed(::basictl::tl_ostream & s, const ::tl2::FieldConflict3& item) {
	if (!s.nat_write(0x2cf6e157)) { return false; }
	return tl2::details::FieldConflict3Write(s, item);
}

bool tl2::FieldConflict4::write_json(std::ostream& s)const {
	if (!::tl2::details::FieldConflict4WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict4::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::FieldConflict4Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::FieldConflict4::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::FieldConflict4Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::FieldConflict4::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::FieldConflict4::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::FieldConflict4::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::FieldConflict4ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::FieldConflict4::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::FieldConflict4WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::FieldConflict4::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::FieldConflict4::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::FieldConflict4Reset(::tl2::FieldConflict4& item) noexcept {
	item.X = 0;
	item.SetX = 0;
}

bool tl2::details::FieldConflict4WriteJSON(std::ostream& s, const ::tl2::FieldConflict4& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.X != 0) {
		add_comma = true;
		s << "\"X\":";
		s << item.X;
	}
	if (item.SetX != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"SetX\":";
		s << item.SetX;
	}
	s << "}";
	return true;
}

bool tl2::details::FieldConflict4Read(::basictl::tl_istream & s, ::tl2::FieldConflict4& item) noexcept {
	if (!s.int_read(item.X)) { return false; }
	if (!s.int_read(item.SetX)) { return false; }
	return true;
}

bool tl2::details::FieldConflict4Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict4& item) noexcept {
	if (!s.int_write(item.X)) { return false;}
	if (!s.int_write(item.SetX)) { return false;}
	return true;
}

bool tl2::details::FieldConflict4ReadBoxed(::basictl::tl_istream & s, ::tl2::FieldConflict4& item) {
	if (!s.nat_read_exact_tag(0xd93c186a)) { return false; }
	return tl2::details::FieldConflict4Read(s, item);
}

bool tl2::details::FieldConflict4WriteBoxed(::basictl::tl_ostream & s, const ::tl2::FieldConflict4& item) {
	if (!s.nat_write(0xd93c186a)) { return false; }
	return tl2::details::FieldConflict4Write(s, item);
}

void tl2::details::FloatReset(float& item) noexcept {
	item = 0;
}

bool tl2::details::FloatWriteJSON(std::ostream& s, const float& item) noexcept {
	s << item;
	return true;
}

bool tl2::details::FloatRead(::basictl::tl_istream & s, float& item) noexcept {
	if (!s.float_read(item)) { return false; }
	return true;
}

bool tl2::details::FloatWrite(::basictl::tl_ostream & s, const float& item) noexcept {
	if (!s.float_write(item)) { return false;}
	return true;
}

bool tl2::details::FloatReadBoxed(::basictl::tl_istream & s, float& item) {
	if (!s.nat_read_exact_tag(0x824dab22)) { return false; }
	return tl2::details::FloatRead(s, item);
}

bool tl2::details::FloatWriteBoxed(::basictl::tl_ostream & s, const float& item) {
	if (!s.nat_write(0x824dab22)) { return false; }
	return tl2::details::FloatWrite(s, item);
}

bool tl2::Get_arrays::write_json(std::ostream& s)const {
	if (!::tl2::details::GetArraysWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::Get_arrays::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetArraysRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::Get_arrays::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetArraysWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::Get_arrays::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::Get_arrays::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::Get_arrays::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetArraysReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::Get_arrays::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetArraysWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::Get_arrays::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::Get_arrays::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::GetArraysReset(::tl2::Get_arrays& item) noexcept {
	item.n = 0;
	item.a.clear();
	::tl2::details::BuiltinTuple5IntReset(item.b);
}

bool tl2::details::GetArraysWriteJSON(std::ostream& s, const ::tl2::Get_arrays& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.n != 0) {
		add_comma = true;
		s << "\"n\":";
		s << item.n;
	}
	if ((item.a.size() != 0) || (item.n != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.a, item.n)) { return false; }
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"b\":";
	if (!::tl2::details::BuiltinTuple5IntWriteJSON(s, item.b)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::GetArraysRead(::basictl::tl_istream & s, ::tl2::Get_arrays& item) noexcept {
	if (!s.nat_read(item.n)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.a, item.n)) { return false; }
	if (!::tl2::details::BuiltinTuple5IntRead(s, item.b)) { return false; }
	return true;
}

bool tl2::details::GetArraysWrite(::basictl::tl_ostream & s, const ::tl2::Get_arrays& item) noexcept {
	if (!s.nat_write(item.n)) { return false;}
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.a, item.n)) { return false; }
	if (!::tl2::details::BuiltinTuple5IntWrite(s, item.b)) { return false; }
	return true;
}

bool tl2::details::GetArraysReadBoxed(::basictl::tl_istream & s, ::tl2::Get_arrays& item) {
	if (!s.nat_read_exact_tag(0x90658cdb)) { return false; }
	return tl2::details::GetArraysRead(s, item);
}

bool tl2::details::GetArraysWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Get_arrays& item) {
	if (!s.nat_write(0x90658cdb)) { return false; }
	return tl2::details::GetArraysWrite(s, item);
}

bool tl2::details::GetArraysReadResult(::basictl::tl_istream & s, tl2::Get_arrays& item, std::array<int32_t, 5>& result) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTuple5IntRead(s, result)) { return false; }
	return true;
}
bool tl2::details::GetArraysWriteResult(::basictl::tl_ostream & s, tl2::Get_arrays& item, std::array<int32_t, 5>& result) {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTuple5IntWrite(s, result)) { return false; }
	return true;
}

bool tl2::Get_arrays::read_result(::basictl::tl_istream & s, std::array<int32_t, 5> & result) noexcept {
	bool success = tl2::details::GetArraysReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::Get_arrays::write_result(::basictl::tl_ostream & s, std::array<int32_t, 5> & result) noexcept {
	bool success = tl2::details::GetArraysWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::Get_arrays::read_result_or_throw(::basictl::tl_throwable_istream & s, std::array<int32_t, 5> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::Get_arrays::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::array<int32_t, 5> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::GetDouble::write_json(std::ostream& s)const {
	if (!::tl2::details::GetDoubleWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetDouble::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetDoubleRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetDouble::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetDoubleWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetDouble::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::GetDouble::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::GetDouble::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetDoubleReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetDouble::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetDoubleWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetDouble::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::GetDouble::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::GetDoubleReset(::tl2::GetDouble& item) noexcept {
	item.x = 0;
}

bool tl2::details::GetDoubleWriteJSON(std::ostream& s, const ::tl2::GetDouble& item) noexcept {
	s << "{";
	if (item.x != 0) {
		s << "\"x\":";
		s << item.x;
	}
	s << "}";
	return true;
}

bool tl2::details::GetDoubleRead(::basictl::tl_istream & s, ::tl2::GetDouble& item) noexcept {
	if (!s.nat_read_exact_tag(0x2210c154)) { return false; }
	if (!s.double_read(item.x)) { return false; }
	return true;
}

bool tl2::details::GetDoubleWrite(::basictl::tl_ostream & s, const ::tl2::GetDouble& item) noexcept {
	if (!s.nat_write(0x2210c154)) { return false; }
	if (!s.double_write(item.x)) { return false;}
	return true;
}

bool tl2::details::GetDoubleReadBoxed(::basictl::tl_istream & s, ::tl2::GetDouble& item) {
	if (!s.nat_read_exact_tag(0x39711d7b)) { return false; }
	return tl2::details::GetDoubleRead(s, item);
}

bool tl2::details::GetDoubleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetDouble& item) {
	if (!s.nat_write(0x39711d7b)) { return false; }
	return tl2::details::GetDoubleWrite(s, item);
}

bool tl2::details::GetDoubleReadResult(::basictl::tl_istream & s, tl2::GetDouble& item, double& result) {
	if (!s.nat_read_exact_tag(0x2210c154)) { return false; }
	if (!s.double_read(result)) { return false; }
	return true;
}
bool tl2::details::GetDoubleWriteResult(::basictl::tl_ostream & s, tl2::GetDouble& item, double& result) {
	if (!s.nat_write(0x2210c154)) { return false; }
	if (!s.double_write(result)) { return false;}
	return true;
}

bool tl2::GetDouble::read_result(::basictl::tl_istream & s, double & result) noexcept {
	bool success = tl2::details::GetDoubleReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::GetDouble::write_result(::basictl::tl_ostream & s, double & result) noexcept {
	bool success = tl2::details::GetDoubleWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::GetDouble::read_result_or_throw(::basictl::tl_throwable_istream & s, double & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::GetDouble::write_result_or_throw(::basictl::tl_throwable_ostream & s, double & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::GetFloat::write_json(std::ostream& s)const {
	if (!::tl2::details::GetFloatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetFloat::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetFloatRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetFloat::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetFloatWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetFloat::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::GetFloat::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::GetFloat::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetFloatReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetFloat::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetFloatWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetFloat::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::GetFloat::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::GetFloatReset(::tl2::GetFloat& item) noexcept {
	item.x = 0;
}

bool tl2::details::GetFloatWriteJSON(std::ostream& s, const ::tl2::GetFloat& item) noexcept {
	s << "{";
	if (item.x != 0) {
		s << "\"x\":";
		s << item.x;
	}
	s << "}";
	return true;
}

bool tl2::details::GetFloatRead(::basictl::tl_istream & s, ::tl2::GetFloat& item) noexcept {
	if (!s.float_read(item.x)) { return false; }
	return true;
}

bool tl2::details::GetFloatWrite(::basictl::tl_ostream & s, const ::tl2::GetFloat& item) noexcept {
	if (!s.float_write(item.x)) { return false;}
	return true;
}

bool tl2::details::GetFloatReadBoxed(::basictl::tl_istream & s, ::tl2::GetFloat& item) {
	if (!s.nat_read_exact_tag(0x25a7bc68)) { return false; }
	return tl2::details::GetFloatRead(s, item);
}

bool tl2::details::GetFloatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetFloat& item) {
	if (!s.nat_write(0x25a7bc68)) { return false; }
	return tl2::details::GetFloatWrite(s, item);
}

bool tl2::details::GetFloatReadResult(::basictl::tl_istream & s, tl2::GetFloat& item, float& result) {
	if (!s.nat_read_exact_tag(0x824dab22)) { return false; }
	if (!s.float_read(result)) { return false; }
	return true;
}
bool tl2::details::GetFloatWriteResult(::basictl::tl_ostream & s, tl2::GetFloat& item, float& result) {
	if (!s.nat_write(0x824dab22)) { return false; }
	if (!s.float_write(result)) { return false;}
	return true;
}

bool tl2::GetFloat::read_result(::basictl::tl_istream & s, float & result) noexcept {
	bool success = tl2::details::GetFloatReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::GetFloat::write_result(::basictl::tl_ostream & s, float & result) noexcept {
	bool success = tl2::details::GetFloatWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::GetFloat::read_result_or_throw(::basictl::tl_throwable_istream & s, float & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::GetFloat::write_result_or_throw(::basictl::tl_throwable_ostream & s, float & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::GetMaybeIface::write_json(std::ostream& s)const {
	if (!::tl2::details::GetMaybeIfaceWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetMaybeIface::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetMaybeIfaceRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetMaybeIface::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetMaybeIfaceWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetMaybeIface::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::GetMaybeIface::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::GetMaybeIface::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetMaybeIfaceReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetMaybeIface::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetMaybeIfaceWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetMaybeIface::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::GetMaybeIface::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::GetMaybeIfaceReset(::tl2::GetMaybeIface& item) noexcept {
	::tl2::details::Service1ValueReset(item.x);
}

bool tl2::details::GetMaybeIfaceWriteJSON(std::ostream& s, const ::tl2::GetMaybeIface& item) noexcept {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::Service1ValueWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::GetMaybeIfaceRead(::basictl::tl_istream & s, ::tl2::GetMaybeIface& item) noexcept {
	if (!::tl2::details::Service1ValueReadBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::GetMaybeIfaceWrite(::basictl::tl_ostream & s, const ::tl2::GetMaybeIface& item) noexcept {
	if (!::tl2::details::Service1ValueWriteBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::GetMaybeIfaceReadBoxed(::basictl::tl_istream & s, ::tl2::GetMaybeIface& item) {
	if (!s.nat_read_exact_tag(0x6b055ae4)) { return false; }
	return tl2::details::GetMaybeIfaceRead(s, item);
}

bool tl2::details::GetMaybeIfaceWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetMaybeIface& item) {
	if (!s.nat_write(0x6b055ae4)) { return false; }
	return tl2::details::GetMaybeIfaceWrite(s, item);
}

bool tl2::details::GetMaybeIfaceReadResult(::basictl::tl_istream & s, tl2::GetMaybeIface& item, std::optional<::tl2::service1::Value>& result) {
	if (!::tl2::details::Service1ValueBoxedMaybeReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::GetMaybeIfaceWriteResult(::basictl::tl_ostream & s, tl2::GetMaybeIface& item, std::optional<::tl2::service1::Value>& result) {
	if (!::tl2::details::Service1ValueBoxedMaybeWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::GetMaybeIface::read_result(::basictl::tl_istream & s, std::optional<::tl2::service1::Value> & result) noexcept {
	bool success = tl2::details::GetMaybeIfaceReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::GetMaybeIface::write_result(::basictl::tl_ostream & s, std::optional<::tl2::service1::Value> & result) noexcept {
	bool success = tl2::details::GetMaybeIfaceWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::GetMaybeIface::read_result_or_throw(::basictl::tl_throwable_istream & s, std::optional<::tl2::service1::Value> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::GetMaybeIface::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::optional<::tl2::service1::Value> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::GetMyDictOfInt::write_json(std::ostream& s)const {
	if (!::tl2::details::GetMyDictOfIntWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyDictOfInt::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetMyDictOfIntRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetMyDictOfInt::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetMyDictOfIntWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetMyDictOfInt::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::GetMyDictOfInt::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::GetMyDictOfInt::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetMyDictOfIntReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetMyDictOfInt::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetMyDictOfIntWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetMyDictOfInt::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::GetMyDictOfInt::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::GetMyDictOfIntReset(::tl2::GetMyDictOfInt& item) noexcept {
	::tl2::details::MyDictOfIntReset(item.x);
}

bool tl2::details::GetMyDictOfIntWriteJSON(std::ostream& s, const ::tl2::GetMyDictOfInt& item) noexcept {
	s << "{";
	if (item.x.size() != 0) {
		s << "\"x\":";
		if (!::tl2::details::MyDictOfIntWriteJSON(s, item.x)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::GetMyDictOfIntRead(::basictl::tl_istream & s, ::tl2::GetMyDictOfInt& item) noexcept {
	if (!::tl2::details::MyDictOfIntReadBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::GetMyDictOfIntWrite(::basictl::tl_ostream & s, const ::tl2::GetMyDictOfInt& item) noexcept {
	if (!::tl2::details::MyDictOfIntWriteBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::GetMyDictOfIntReadBoxed(::basictl::tl_istream & s, ::tl2::GetMyDictOfInt& item) {
	if (!s.nat_read_exact_tag(0x166f962c)) { return false; }
	return tl2::details::GetMyDictOfIntRead(s, item);
}

bool tl2::details::GetMyDictOfIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetMyDictOfInt& item) {
	if (!s.nat_write(0x166f962c)) { return false; }
	return tl2::details::GetMyDictOfIntWrite(s, item);
}

bool tl2::details::GetMyDictOfIntReadResult(::basictl::tl_istream & s, tl2::GetMyDictOfInt& item, ::tl2::MyDictOfInt& result) {
	if (!::tl2::details::MyDictOfIntReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::GetMyDictOfIntWriteResult(::basictl::tl_ostream & s, tl2::GetMyDictOfInt& item, ::tl2::MyDictOfInt& result) {
	if (!::tl2::details::MyDictOfIntWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::GetMyDictOfInt::read_result(::basictl::tl_istream & s, ::tl2::MyDictOfInt & result) noexcept {
	bool success = tl2::details::GetMyDictOfIntReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::GetMyDictOfInt::write_result(::basictl::tl_ostream & s, ::tl2::MyDictOfInt & result) noexcept {
	bool success = tl2::details::GetMyDictOfIntWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::GetMyDictOfInt::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::MyDictOfInt & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::GetMyDictOfInt::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::MyDictOfInt & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::GetMyDouble::write_json(std::ostream& s)const {
	if (!::tl2::details::GetMyDoubleWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyDouble::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetMyDoubleRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetMyDouble::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetMyDoubleWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetMyDouble::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::GetMyDouble::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::GetMyDouble::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetMyDoubleReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetMyDouble::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetMyDoubleWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetMyDouble::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::GetMyDouble::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::GetMyDoubleReset(::tl2::GetMyDouble& item) noexcept {
	::tl2::details::MyDoubleReset(item.x);
}

bool tl2::details::GetMyDoubleWriteJSON(std::ostream& s, const ::tl2::GetMyDouble& item) noexcept {
	s << "{";
	if (item.x != 0) {
		s << "\"x\":";
		if (!::tl2::details::MyDoubleWriteJSON(s, item.x)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::GetMyDoubleRead(::basictl::tl_istream & s, ::tl2::GetMyDouble& item) noexcept {
	if (!::tl2::details::MyDoubleRead(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::GetMyDoubleWrite(::basictl::tl_ostream & s, const ::tl2::GetMyDouble& item) noexcept {
	if (!::tl2::details::MyDoubleWrite(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::GetMyDoubleReadBoxed(::basictl::tl_istream & s, ::tl2::GetMyDouble& item) {
	if (!s.nat_read_exact_tag(0xb660ad10)) { return false; }
	return tl2::details::GetMyDoubleRead(s, item);
}

bool tl2::details::GetMyDoubleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetMyDouble& item) {
	if (!s.nat_write(0xb660ad10)) { return false; }
	return tl2::details::GetMyDoubleWrite(s, item);
}

bool tl2::details::GetMyDoubleReadResult(::basictl::tl_istream & s, tl2::GetMyDouble& item, ::tl2::MyDouble& result) {
	if (!::tl2::details::MyDoubleReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::GetMyDoubleWriteResult(::basictl::tl_ostream & s, tl2::GetMyDouble& item, ::tl2::MyDouble& result) {
	if (!::tl2::details::MyDoubleWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::GetMyDouble::read_result(::basictl::tl_istream & s, ::tl2::MyDouble & result) noexcept {
	bool success = tl2::details::GetMyDoubleReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::GetMyDouble::write_result(::basictl::tl_ostream & s, ::tl2::MyDouble & result) noexcept {
	bool success = tl2::details::GetMyDoubleWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::GetMyDouble::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::MyDouble & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::GetMyDouble::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::MyDouble & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::GetMyValue::write_json(std::ostream& s)const {
	if (!::tl2::details::GetMyValueWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyValue::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetMyValueRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetMyValue::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetMyValueWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetMyValue::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::GetMyValue::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::GetMyValue::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetMyValueReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetMyValue::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetMyValueWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetMyValue::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::GetMyValue::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::GetMyValueReset(::tl2::GetMyValue& item) noexcept {
	::tl2::details::MyValueReset(item.x);
}

bool tl2::details::GetMyValueWriteJSON(std::ostream& s, const ::tl2::GetMyValue& item) noexcept {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::MyValueWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::GetMyValueRead(::basictl::tl_istream & s, ::tl2::GetMyValue& item) noexcept {
	if (!::tl2::details::MyValueReadBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::GetMyValueWrite(::basictl::tl_ostream & s, const ::tl2::GetMyValue& item) noexcept {
	if (!::tl2::details::MyValueWriteBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::GetMyValueReadBoxed(::basictl::tl_istream & s, ::tl2::GetMyValue& item) {
	if (!s.nat_read_exact_tag(0xb3df27fe)) { return false; }
	return tl2::details::GetMyValueRead(s, item);
}

bool tl2::details::GetMyValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetMyValue& item) {
	if (!s.nat_write(0xb3df27fe)) { return false; }
	return tl2::details::GetMyValueWrite(s, item);
}

bool tl2::details::GetMyValueReadResult(::basictl::tl_istream & s, tl2::GetMyValue& item, ::tl2::MyValue& result) {
	if (!::tl2::details::MyValueReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::GetMyValueWriteResult(::basictl::tl_ostream & s, tl2::GetMyValue& item, ::tl2::MyValue& result) {
	if (!::tl2::details::MyValueWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::GetMyValue::read_result(::basictl::tl_istream & s, ::tl2::MyValue & result) noexcept {
	bool success = tl2::details::GetMyValueReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::GetMyValue::write_result(::basictl::tl_ostream & s, ::tl2::MyValue & result) noexcept {
	bool success = tl2::details::GetMyValueWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::GetMyValue::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::MyValue & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::GetMyValue::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::MyValue & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::GetNonOptNat::write_json(std::ostream& s)const {
	if (!::tl2::details::GetNonOptNatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetNonOptNat::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetNonOptNatRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetNonOptNat::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetNonOptNatWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetNonOptNat::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::GetNonOptNat::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::GetNonOptNat::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetNonOptNatReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetNonOptNat::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetNonOptNatWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetNonOptNat::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::GetNonOptNat::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::GetNonOptNatReset(::tl2::GetNonOptNat& item) noexcept {
	item.n = 0;
	item.xs.clear();
}

bool tl2::details::GetNonOptNatWriteJSON(std::ostream& s, const ::tl2::GetNonOptNat& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.n != 0) {
		add_comma = true;
		s << "\"n\":";
		s << item.n;
	}
	if ((item.xs.size() != 0) || (item.n != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"xs\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.xs, item.n)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::GetNonOptNatRead(::basictl::tl_istream & s, ::tl2::GetNonOptNat& item) noexcept {
	if (!s.nat_read(item.n)) { return false; }
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.xs, item.n)) { return false; }
	return true;
}

bool tl2::details::GetNonOptNatWrite(::basictl::tl_ostream & s, const ::tl2::GetNonOptNat& item) noexcept {
	if (!s.nat_write(item.n)) { return false;}
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.xs, item.n)) { return false; }
	return true;
}

bool tl2::details::GetNonOptNatReadBoxed(::basictl::tl_istream & s, ::tl2::GetNonOptNat& item) {
	if (!s.nat_read_exact_tag(0x67665961)) { return false; }
	return tl2::details::GetNonOptNatRead(s, item);
}

bool tl2::details::GetNonOptNatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetNonOptNat& item) {
	if (!s.nat_write(0x67665961)) { return false; }
	return tl2::details::GetNonOptNatWrite(s, item);
}

bool tl2::details::GetNonOptNatReadResult(::basictl::tl_istream & s, tl2::GetNonOptNat& item, std::vector<int32_t>& result) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, result, item.n)) { return false; }
	return true;
}
bool tl2::details::GetNonOptNatWriteResult(::basictl::tl_ostream & s, tl2::GetNonOptNat& item, std::vector<int32_t>& result) {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntWrite(s, result, item.n)) { return false; }
	return true;
}

bool tl2::GetNonOptNat::read_result(::basictl::tl_istream & s, std::vector<int32_t> & result) noexcept {
	bool success = tl2::details::GetNonOptNatReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::GetNonOptNat::write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result) noexcept {
	bool success = tl2::details::GetNonOptNatWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::GetNonOptNat::read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<int32_t> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::GetNonOptNat::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<int32_t> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::GetStats::write_json(std::ostream& s)const {
	if (!::tl2::details::GetStatsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetStats::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetStatsRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetStats::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetStatsWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetStats::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::GetStats::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::GetStats::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::GetStatsReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::GetStats::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::GetStatsWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::GetStats::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::GetStats::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::GetStatsReset(::tl2::GetStats& item) noexcept {
	::tl2::details::TasksQueueTypeStatsReset(item.x);
}

bool tl2::details::GetStatsWriteJSON(std::ostream& s, const ::tl2::GetStats& item) noexcept {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::TasksQueueTypeStatsWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::GetStatsRead(::basictl::tl_istream & s, ::tl2::GetStats& item) noexcept {
	if (!::tl2::details::TasksQueueTypeStatsRead(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::GetStatsWrite(::basictl::tl_ostream & s, const ::tl2::GetStats& item) noexcept {
	if (!::tl2::details::TasksQueueTypeStatsWrite(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::GetStatsReadBoxed(::basictl::tl_istream & s, ::tl2::GetStats& item) {
	if (!s.nat_read_exact_tag(0xbaa6da35)) { return false; }
	return tl2::details::GetStatsRead(s, item);
}

bool tl2::details::GetStatsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::GetStats& item) {
	if (!s.nat_write(0xbaa6da35)) { return false; }
	return tl2::details::GetStatsWrite(s, item);
}

bool tl2::details::GetStatsReadResult(::basictl::tl_istream & s, tl2::GetStats& item, ::tl2::tasks::QueueTypeStats& result) {
	if (!::tl2::details::TasksQueueTypeStatsReadBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::details::GetStatsWriteResult(::basictl::tl_ostream & s, tl2::GetStats& item, ::tl2::tasks::QueueTypeStats& result) {
	if (!::tl2::details::TasksQueueTypeStatsWriteBoxed(s, result)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::GetStats::read_result(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeStats & result) noexcept {
	bool success = tl2::details::GetStatsReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::GetStats::write_result(::basictl::tl_ostream & s, ::tl2::tasks::QueueTypeStats & result) noexcept {
	bool success = tl2::details::GetStatsWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::GetStats::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::tasks::QueueTypeStats & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::GetStats::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::tasks::QueueTypeStats & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

void tl2::details::IntReset(int32_t& item) noexcept {
	item = 0;
}

bool tl2::details::IntWriteJSON(std::ostream& s, const int32_t& item) noexcept {
	s << item;
	return true;
}

bool tl2::details::IntRead(::basictl::tl_istream & s, int32_t& item) noexcept {
	if (!s.int_read(item)) { return false; }
	return true;
}

bool tl2::details::IntWrite(::basictl::tl_ostream & s, const int32_t& item) noexcept {
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

bool tl2::details::IntMaybeWriteJSON(std::ostream & s, const std::optional<int32_t>& item) {
	s << "{";
	if (item) {
		s << "\"ok\":true";
		if((*item) != 0) {
			s << ",\"value\":";
			s << *item;
		}
	}
	s << "}";
	return true;
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

bool tl2::Integer::write_json(std::ostream& s)const {
	if (!::tl2::details::IntegerWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::Integer::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::IntegerRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::Integer::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::IntegerWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::Integer::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::Integer::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::Integer::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::IntegerReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::Integer::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::IntegerWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::Integer::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::Integer::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::IntegerReset(::tl2::Integer& item) noexcept {
	item.value = 0;
}

bool tl2::details::IntegerWriteJSON(std::ostream& s, const ::tl2::Integer& item) noexcept {
	s << "{";
	if (item.value != 0) {
		s << "\"value\":";
		s << item.value;
	}
	s << "}";
	return true;
}

bool tl2::details::IntegerRead(::basictl::tl_istream & s, ::tl2::Integer& item) noexcept {
	if (!s.int_read(item.value)) { return false; }
	return true;
}

bool tl2::details::IntegerWrite(::basictl::tl_ostream & s, const ::tl2::Integer& item) noexcept {
	if (!s.int_write(item.value)) { return false;}
	return true;
}

bool tl2::details::IntegerReadBoxed(::basictl::tl_istream & s, ::tl2::Integer& item) {
	if (!s.nat_read_exact_tag(0x7e194796)) { return false; }
	return tl2::details::IntegerRead(s, item);
}

bool tl2::details::IntegerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Integer& item) {
	if (!s.nat_write(0x7e194796)) { return false; }
	return tl2::details::IntegerWrite(s, item);
}

bool tl2::Issue3498::write_json(std::ostream& s)const {
	if (!::tl2::details::Issue3498WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::Issue3498::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Issue3498Read(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::Issue3498::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Issue3498Write(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::Issue3498::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::Issue3498::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::Issue3498::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::Issue3498ReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::Issue3498::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::Issue3498WriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::Issue3498::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::Issue3498::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::Issue3498Reset(::tl2::Issue3498& item) noexcept {
	item.x.clear();
}

bool tl2::details::Issue3498WriteJSON(std::ostream& s, const ::tl2::Issue3498& item) noexcept {
	s << "{";
	if (item.x.size() != 0) {
		s << "\"x\":";
		if (!::tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowWriteJSON(s, item.x)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Issue3498Read(::basictl::tl_istream & s, ::tl2::Issue3498& item) noexcept {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowRead(s, item.x)) { return false; }
	return true;
}

bool tl2::details::Issue3498Write(::basictl::tl_ostream & s, const ::tl2::Issue3498& item) noexcept {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowWrite(s, item.x)) { return false; }
	return true;
}

bool tl2::details::Issue3498ReadBoxed(::basictl::tl_istream & s, ::tl2::Issue3498& item) {
	if (!s.nat_read_exact_tag(0xf54b7b0a)) { return false; }
	return tl2::details::Issue3498Read(s, item);
}

bool tl2::details::Issue3498WriteBoxed(::basictl::tl_ostream & s, const ::tl2::Issue3498& item) {
	if (!s.nat_write(0xf54b7b0a)) { return false; }
	return tl2::details::Issue3498Write(s, item);
}

void tl2::details::LeftService6ErrorVectorService6FindResultRowReset(::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept {
	::tl2::details::Service6ErrorReset(item.value);
}

bool tl2::details::LeftService6ErrorVectorService6FindResultRowWriteJSON(std::ostream& s, const ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::Service6ErrorWriteJSON(s, item.value)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::LeftService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept {
	if (!::tl2::details::Service6ErrorRead(s, item.value)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::LeftService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept {
	if (!::tl2::details::Service6ErrorWrite(s, item.value)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::LeftService6ErrorVectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	if (!s.nat_read_exact_tag(0x0a29cd5d)) { return false; }
	return tl2::details::LeftService6ErrorVectorService6FindResultRowRead(s, item);
}

bool tl2::details::LeftService6ErrorVectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	if (!s.nat_write(0x0a29cd5d)) { return false; }
	return tl2::details::LeftService6ErrorVectorService6FindResultRowWrite(s, item);
}

void tl2::details::LongReset(int64_t& item) noexcept {
	item = 0;
}

bool tl2::details::LongWriteJSON(std::ostream& s, const int64_t& item) noexcept {
	s << item;
	return true;
}

bool tl2::details::LongRead(::basictl::tl_istream & s, int64_t& item) noexcept {
	if (!s.long_read(item)) { return false; }
	return true;
}

bool tl2::details::LongWrite(::basictl::tl_ostream & s, const int64_t& item) noexcept {
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

void tl2::details::MyAnonMcValueReset(::tl2::MyAnonMcValue& item) noexcept {
	::tl2::details::Service1ValueReset(item);
}

bool tl2::details::MyAnonMcValueWriteJSON(std::ostream& s, const ::tl2::MyAnonMcValue& item) noexcept {
	if (!::tl2::details::Service1ValueWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::MyAnonMcValueRead(::basictl::tl_istream & s, ::tl2::MyAnonMcValue& item) noexcept {
	if (!::tl2::details::Service1ValueReadBoxed(s, item)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::MyAnonMcValueWrite(::basictl::tl_ostream & s, const ::tl2::MyAnonMcValue& item) noexcept {
	if (!::tl2::details::Service1ValueWriteBoxed(s, item)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::MyAnonMcValueReadBoxed(::basictl::tl_istream & s, ::tl2::MyAnonMcValue& item) {
	if (!s.nat_read_exact_tag(0x569310db)) { return false; }
	return tl2::details::MyAnonMcValueRead(s, item);
}

bool tl2::details::MyAnonMcValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyAnonMcValue& item) {
	if (!s.nat_write(0x569310db)) { return false; }
	return tl2::details::MyAnonMcValueWrite(s, item);
}

bool tl2::MyBoxedArray::write_json(std::ostream& s)const {
	if (!::tl2::details::MyBoxedArrayWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::MyBoxedArray::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyBoxedArrayRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyBoxedArray::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyBoxedArrayWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyBoxedArray::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::MyBoxedArray::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::MyBoxedArray::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyBoxedArrayReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyBoxedArray::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyBoxedArrayWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyBoxedArray::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::MyBoxedArray::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::MyBoxedArrayReset(::tl2::MyBoxedArray& item) noexcept {
	::tl2::details::BuiltinTuple2IntBoxedReset(item.data);
}

bool tl2::details::MyBoxedArrayWriteJSON(std::ostream& s, const ::tl2::MyBoxedArray& item) noexcept {
	s << "{";
	s << "\"data\":";
	if (!::tl2::details::BuiltinTuple2IntBoxedWriteJSON(s, item.data)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::MyBoxedArrayRead(::basictl::tl_istream & s, ::tl2::MyBoxedArray& item) noexcept {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTuple2IntBoxedRead(s, item.data)) { return false; }
	return true;
}

bool tl2::details::MyBoxedArrayWrite(::basictl::tl_ostream & s, const ::tl2::MyBoxedArray& item) noexcept {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTuple2IntBoxedWrite(s, item.data)) { return false; }
	return true;
}

bool tl2::details::MyBoxedArrayReadBoxed(::basictl::tl_istream & s, ::tl2::MyBoxedArray& item) {
	if (!s.nat_read_exact_tag(0x288f64f0)) { return false; }
	return tl2::details::MyBoxedArrayRead(s, item);
}

bool tl2::details::MyBoxedArrayWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyBoxedArray& item) {
	if (!s.nat_write(0x288f64f0)) { return false; }
	return tl2::details::MyBoxedArrayWrite(s, item);
}

bool tl2::MyBoxedTupleSlice::write_json(std::ostream& s)const {
	if (!::tl2::details::MyBoxedTupleSliceWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::MyBoxedTupleSlice::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyBoxedTupleSliceRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyBoxedTupleSlice::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyBoxedTupleSliceWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyBoxedTupleSlice::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::MyBoxedTupleSlice::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::MyBoxedTupleSlice::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyBoxedTupleSliceReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyBoxedTupleSlice::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyBoxedTupleSliceWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyBoxedTupleSlice::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::MyBoxedTupleSlice::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::MyBoxedTupleSliceReset(::tl2::MyBoxedTupleSlice& item) noexcept {
	item.n = 0;
	item.data.clear();
}

bool tl2::details::MyBoxedTupleSliceWriteJSON(std::ostream& s, const ::tl2::MyBoxedTupleSlice& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.n != 0) {
		add_comma = true;
		s << "\"n\":";
		s << item.n;
	}
	if ((item.data.size() != 0) || (item.n != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"data\":";
		if (!::tl2::details::BuiltinTupleIntBoxedWriteJSON(s, item.data, item.n)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::MyBoxedTupleSliceRead(::basictl::tl_istream & s, ::tl2::MyBoxedTupleSlice& item) noexcept {
	if (!s.nat_read(item.n)) { return false; }
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntBoxedRead(s, item.data, item.n)) { return false; }
	return true;
}

bool tl2::details::MyBoxedTupleSliceWrite(::basictl::tl_ostream & s, const ::tl2::MyBoxedTupleSlice& item) noexcept {
	if (!s.nat_write(item.n)) { return false;}
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntBoxedWrite(s, item.data, item.n)) { return false; }
	return true;
}

bool tl2::details::MyBoxedTupleSliceReadBoxed(::basictl::tl_istream & s, ::tl2::MyBoxedTupleSlice& item) {
	if (!s.nat_read_exact_tag(0x25d1a1be)) { return false; }
	return tl2::details::MyBoxedTupleSliceRead(s, item);
}

bool tl2::details::MyBoxedTupleSliceWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyBoxedTupleSlice& item) {
	if (!s.nat_write(0x25d1a1be)) { return false; }
	return tl2::details::MyBoxedTupleSliceWrite(s, item);
}

bool tl2::MyBoxedVectorSlice::write_json(std::ostream& s)const {
	if (!::tl2::details::MyBoxedVectorSliceWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::MyBoxedVectorSlice::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyBoxedVectorSliceRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyBoxedVectorSlice::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyBoxedVectorSliceWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyBoxedVectorSlice::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::MyBoxedVectorSlice::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::MyBoxedVectorSlice::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyBoxedVectorSliceReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyBoxedVectorSlice::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyBoxedVectorSliceWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyBoxedVectorSlice::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::MyBoxedVectorSlice::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::MyBoxedVectorSliceReset(::tl2::MyBoxedVectorSlice& item) noexcept {
	item.data.clear();
}

bool tl2::details::MyBoxedVectorSliceWriteJSON(std::ostream& s, const ::tl2::MyBoxedVectorSlice& item) noexcept {
	s << "{";
	if (item.data.size() != 0) {
		s << "\"data\":";
		if (!::tl2::details::BuiltinVectorIntBoxedWriteJSON(s, item.data)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::MyBoxedVectorSliceRead(::basictl::tl_istream & s, ::tl2::MyBoxedVectorSlice& item) noexcept {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntBoxedRead(s, item.data)) { return false; }
	return true;
}

bool tl2::details::MyBoxedVectorSliceWrite(::basictl::tl_ostream & s, const ::tl2::MyBoxedVectorSlice& item) noexcept {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntBoxedWrite(s, item.data)) { return false; }
	return true;
}

bool tl2::details::MyBoxedVectorSliceReadBoxed(::basictl::tl_istream & s, ::tl2::MyBoxedVectorSlice& item) {
	if (!s.nat_read_exact_tag(0x57d164bb)) { return false; }
	return tl2::details::MyBoxedVectorSliceRead(s, item);
}

bool tl2::details::MyBoxedVectorSliceWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyBoxedVectorSlice& item) {
	if (!s.nat_write(0x57d164bb)) { return false; }
	return tl2::details::MyBoxedVectorSliceWrite(s, item);
}

void tl2::details::MyDictOfIntReset(::tl2::MyDictOfInt& item) noexcept {
	item.clear();
}

bool tl2::details::MyDictOfIntWriteJSON(std::ostream& s, const ::tl2::MyDictOfInt& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::MyDictOfIntRead(::basictl::tl_istream & s, ::tl2::MyDictOfInt& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::MyDictOfIntWrite(::basictl::tl_ostream & s, const ::tl2::MyDictOfInt& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::MyDictOfIntReadBoxed(::basictl::tl_istream & s, ::tl2::MyDictOfInt& item) {
	if (!s.nat_read_exact_tag(0xb8019a3d)) { return false; }
	return tl2::details::MyDictOfIntRead(s, item);
}

bool tl2::details::MyDictOfIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyDictOfInt& item) {
	if (!s.nat_write(0xb8019a3d)) { return false; }
	return tl2::details::MyDictOfIntWrite(s, item);
}

void tl2::details::MyDoubleReset(::tl2::MyDouble& item) noexcept {
	item = 0;
}

bool tl2::details::MyDoubleWriteJSON(std::ostream& s, const ::tl2::MyDouble& item) noexcept {
	s << item;
	return true;
}

bool tl2::details::MyDoubleRead(::basictl::tl_istream & s, ::tl2::MyDouble& item) noexcept {
	if (!s.nat_read_exact_tag(0x2210c154)) { return false; }
	if (!s.double_read(item)) { return false; }
	return true;
}

bool tl2::details::MyDoubleWrite(::basictl::tl_ostream & s, const ::tl2::MyDouble& item) noexcept {
	if (!s.nat_write(0x2210c154)) { return false; }
	if (!s.double_write(item)) { return false;}
	return true;
}

bool tl2::details::MyDoubleReadBoxed(::basictl::tl_istream & s, ::tl2::MyDouble& item) {
	if (!s.nat_read_exact_tag(0x90a6c726)) { return false; }
	return tl2::details::MyDoubleRead(s, item);
}

bool tl2::details::MyDoubleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyDouble& item) {
	if (!s.nat_write(0x90a6c726)) { return false; }
	return tl2::details::MyDoubleWrite(s, item);
}

bool tl2::MyInt::write_json(std::ostream& s)const {
	if (!::tl2::details::MyIntWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::MyInt::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyIntRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyInt::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyIntWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyInt::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::MyInt::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::MyInt::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyIntReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyInt::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyIntWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyInt::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::MyInt::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::MyIntReset(::tl2::MyInt& item) noexcept {
	item.val1 = 0;
}

bool tl2::details::MyIntWriteJSON(std::ostream& s, const ::tl2::MyInt& item) noexcept {
	s << "{";
	if (item.val1 != 0) {
		s << "\"val1\":";
		s << item.val1;
	}
	s << "}";
	return true;
}

bool tl2::details::MyIntRead(::basictl::tl_istream & s, ::tl2::MyInt& item) noexcept {
	if (!s.nat_read_exact_tag(0xa8509bda)) { return false; }
	if (!s.int_read(item.val1)) { return false; }
	return true;
}

bool tl2::details::MyIntWrite(::basictl::tl_ostream & s, const ::tl2::MyInt& item) noexcept {
	if (!s.nat_write(0xa8509bda)) { return false; }
	if (!s.int_write(item.val1)) { return false;}
	return true;
}

bool tl2::details::MyIntReadBoxed(::basictl::tl_istream & s, ::tl2::MyInt& item) {
	if (!s.nat_read_exact_tag(0xc12375b7)) { return false; }
	return tl2::details::MyIntRead(s, item);
}

bool tl2::details::MyIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyInt& item) {
	if (!s.nat_write(0xc12375b7)) { return false; }
	return tl2::details::MyIntWrite(s, item);
}

void tl2::details::MyMaybe1Reset(::tl2::MyMaybe1& item) noexcept {
	item.reset();
}

bool tl2::details::MyMaybe1WriteJSON(std::ostream& s, const ::tl2::MyMaybe1& item) noexcept {
	if (!::tl2::details::MyTuple10MaybeWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::MyMaybe1Read(::basictl::tl_istream & s, ::tl2::MyMaybe1& item) noexcept {
	if (!::tl2::details::MyTuple10MaybeReadBoxed(s, item)) { return false; }
	return true;
}

bool tl2::details::MyMaybe1Write(::basictl::tl_ostream & s, const ::tl2::MyMaybe1& item) noexcept {
	if (!::tl2::details::MyTuple10MaybeWriteBoxed(s, item)) { return false; }
	return true;
}

bool tl2::details::MyMaybe1ReadBoxed(::basictl::tl_istream & s, ::tl2::MyMaybe1& item) {
	if (!s.nat_read_exact_tag(0x32c541fe)) { return false; }
	return tl2::details::MyMaybe1Read(s, item);
}

bool tl2::details::MyMaybe1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMaybe1& item) {
	if (!s.nat_write(0x32c541fe)) { return false; }
	return tl2::details::MyMaybe1Write(s, item);
}

void tl2::details::MyMaybe2Reset(::tl2::MyMaybe2& item) noexcept {
	item.reset();
}

bool tl2::details::MyMaybe2WriteJSON(std::ostream& s, const ::tl2::MyMaybe2& item) noexcept {
	if (!::tl2::details::MyTuple10MaybeWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::MyMaybe2Read(::basictl::tl_istream & s, ::tl2::MyMaybe2& item) noexcept {
	if (!::tl2::details::MyTuple10MaybeReadBoxed(s, item)) { return false; }
	return true;
}

bool tl2::details::MyMaybe2Write(::basictl::tl_ostream & s, const ::tl2::MyMaybe2& item) noexcept {
	if (!::tl2::details::MyTuple10MaybeWriteBoxed(s, item)) { return false; }
	return true;
}

bool tl2::details::MyMaybe2ReadBoxed(::basictl::tl_istream & s, ::tl2::MyMaybe2& item) {
	if (!s.nat_read_exact_tag(0xef6d355c)) { return false; }
	return tl2::details::MyMaybe2Read(s, item);
}

bool tl2::details::MyMaybe2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMaybe2& item) {
	if (!s.nat_write(0xef6d355c)) { return false; }
	return tl2::details::MyMaybe2Write(s, item);
}

bool tl2::MyMcValue::write_json(std::ostream& s)const {
	if (!::tl2::details::MyMcValueWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::MyMcValue::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyMcValueRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyMcValue::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyMcValueWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyMcValue::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::MyMcValue::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::MyMcValue::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyMcValueReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyMcValue::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyMcValueWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyMcValue::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::MyMcValue::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::MyMcValueReset(::tl2::MyMcValue& item) noexcept {
	::tl2::details::Service1ValueReset(item.x);
}

bool tl2::details::MyMcValueWriteJSON(std::ostream& s, const ::tl2::MyMcValue& item) noexcept {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::Service1ValueWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::MyMcValueRead(::basictl::tl_istream & s, ::tl2::MyMcValue& item) noexcept {
	if (!::tl2::details::Service1ValueReadBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::MyMcValueWrite(::basictl::tl_ostream & s, const ::tl2::MyMcValue& item) noexcept {
	if (!::tl2::details::Service1ValueWriteBoxed(s, item.x)) { return s.set_error_unknown_scenario(); }
	return true;
}

bool tl2::details::MyMcValueReadBoxed(::basictl::tl_istream & s, ::tl2::MyMcValue& item) {
	if (!s.nat_read_exact_tag(0xe2ffd978)) { return false; }
	return tl2::details::MyMcValueRead(s, item);
}

bool tl2::details::MyMcValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMcValue& item) {
	if (!s.nat_write(0xe2ffd978)) { return false; }
	return tl2::details::MyMcValueWrite(s, item);
}

bool tl2::MyMcValueTuple::write_json(std::ostream& s)const {
	if (!::tl2::details::MyMcValueTupleWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::MyMcValueTuple::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyMcValueTupleRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyMcValueTuple::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyMcValueTupleWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyMcValueTuple::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::MyMcValueTuple::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::MyMcValueTuple::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyMcValueTupleReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyMcValueTuple::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyMcValueTupleWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyMcValueTuple::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::MyMcValueTuple::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::MyMcValueTupleReset(::tl2::MyMcValueTuple& item) noexcept {
	::tl2::details::BuiltinTuple3Service1ValueReset(item.xs);
}

bool tl2::details::MyMcValueTupleWriteJSON(std::ostream& s, const ::tl2::MyMcValueTuple& item) noexcept {
	s << "{";
	s << "\"xs\":";
	if (!::tl2::details::BuiltinTuple3Service1ValueWriteJSON(s, item.xs)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::MyMcValueTupleRead(::basictl::tl_istream & s, ::tl2::MyMcValueTuple& item) noexcept {
	if (!::tl2::details::BuiltinTuple3Service1ValueRead(s, item.xs)) { return false; }
	return true;
}

bool tl2::details::MyMcValueTupleWrite(::basictl::tl_ostream & s, const ::tl2::MyMcValueTuple& item) noexcept {
	if (!::tl2::details::BuiltinTuple3Service1ValueWrite(s, item.xs)) { return false; }
	return true;
}

bool tl2::details::MyMcValueTupleReadBoxed(::basictl::tl_istream & s, ::tl2::MyMcValueTuple& item) {
	if (!s.nat_read_exact_tag(0x1287d116)) { return false; }
	return tl2::details::MyMcValueTupleRead(s, item);
}

bool tl2::details::MyMcValueTupleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMcValueTuple& item) {
	if (!s.nat_write(0x1287d116)) { return false; }
	return tl2::details::MyMcValueTupleWrite(s, item);
}

bool tl2::MyMcValueVector::write_json(std::ostream& s)const {
	if (!::tl2::details::MyMcValueVectorWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::MyMcValueVector::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyMcValueVectorRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyMcValueVector::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyMcValueVectorWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyMcValueVector::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::MyMcValueVector::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::MyMcValueVector::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyMcValueVectorReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyMcValueVector::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyMcValueVectorWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyMcValueVector::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::MyMcValueVector::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::MyMcValueVectorReset(::tl2::MyMcValueVector& item) noexcept {
	item.xs.clear();
}

bool tl2::details::MyMcValueVectorWriteJSON(std::ostream& s, const ::tl2::MyMcValueVector& item) noexcept {
	s << "{";
	if (item.xs.size() != 0) {
		s << "\"xs\":";
		if (!::tl2::details::BuiltinVectorService1ValueWriteJSON(s, item.xs)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::MyMcValueVectorRead(::basictl::tl_istream & s, ::tl2::MyMcValueVector& item) noexcept {
	if (!::tl2::details::BuiltinVectorService1ValueRead(s, item.xs)) { return false; }
	return true;
}

bool tl2::details::MyMcValueVectorWrite(::basictl::tl_ostream & s, const ::tl2::MyMcValueVector& item) noexcept {
	if (!::tl2::details::BuiltinVectorService1ValueWrite(s, item.xs)) { return false; }
	return true;
}

bool tl2::details::MyMcValueVectorReadBoxed(::basictl::tl_istream & s, ::tl2::MyMcValueVector& item) {
	if (!s.nat_read_exact_tag(0x761d6d58)) { return false; }
	return tl2::details::MyMcValueVectorRead(s, item);
}

bool tl2::details::MyMcValueVectorWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyMcValueVector& item) {
	if (!s.nat_write(0x761d6d58)) { return false; }
	return tl2::details::MyMcValueVectorWrite(s, item);
}

bool tl2::MyString::write_json(std::ostream& s)const {
	if (!::tl2::details::MyStringWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::MyString::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyStringRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyString::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyStringWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyString::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::MyString::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::MyString::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyStringReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyString::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyStringWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyString::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::MyString::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::MyStringReset(::tl2::MyString& item) noexcept {
	item.val2.clear();
}

bool tl2::details::MyStringWriteJSON(std::ostream& s, const ::tl2::MyString& item) noexcept {
	s << "{";
	if (item.val2.size() != 0) {
		s << "\"val2\":";
		s << "\"" << item.val2 << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::MyStringRead(::basictl::tl_istream & s, ::tl2::MyString& item) noexcept {
	if (!s.nat_read_exact_tag(0xb5286e24)) { return false; }
	if (!s.string_read(item.val2)) { return false; }
	return true;
}

bool tl2::details::MyStringWrite(::basictl::tl_ostream & s, const ::tl2::MyString& item) noexcept {
	if (!s.nat_write(0xb5286e24)) { return false; }
	if (!s.string_write(item.val2)) { return false;}
	return true;
}

bool tl2::details::MyStringReadBoxed(::basictl::tl_istream & s, ::tl2::MyString& item) {
	if (!s.nat_read_exact_tag(0xc8bfa969)) { return false; }
	return tl2::details::MyStringRead(s, item);
}

bool tl2::details::MyStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyString& item) {
	if (!s.nat_write(0xc8bfa969)) { return false; }
	return tl2::details::MyStringWrite(s, item);
}

void tl2::details::MyTuple10Reset(::tl2::MyTuplen<10>& item) noexcept {
	::tl2::details::BuiltinTuple10IntBoxedReset(item);
}

bool tl2::details::MyTuple10WriteJSON(std::ostream& s, const ::tl2::MyTuplen<10>& item) noexcept {
	if (!::tl2::details::BuiltinTuple10IntBoxedWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::MyTuple10Read(::basictl::tl_istream & s, ::tl2::MyTuplen<10>& item) noexcept {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTuple10IntBoxedRead(s, item)) { return false; }
	return true;
}

bool tl2::details::MyTuple10Write(::basictl::tl_ostream & s, const ::tl2::MyTuplen<10>& item) noexcept {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTuple10IntBoxedWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::MyTuple10ReadBoxed(::basictl::tl_istream & s, ::tl2::MyTuplen<10>& item) {
	if (!s.nat_read_exact_tag(0x62c51172)) { return false; }
	return tl2::details::MyTuple10Read(s, item);
}

bool tl2::details::MyTuple10WriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyTuplen<10>& item) {
	if (!s.nat_write(0x62c51172)) { return false; }
	return tl2::details::MyTuple10Write(s, item);
}

bool tl2::details::MyTuple10MaybeWriteJSON(std::ostream & s, const std::optional<::tl2::MyTuplen<10>>& item) {
	s << "{";
	if (item) {
		s << "\"ok\":true";
		s << ",\"value\":";
		if (!::tl2::details::MyTuple10WriteJSON(s, *item)) { return false; }
	}
	s << "}";
	return true;
}
bool tl2::details::MyTuple10MaybeReadBoxed(::basictl::tl_istream & s, std::optional<::tl2::MyTuplen<10>>& item) {
	bool has_item = false;
	if (!s.bool_read(has_item, 0x27930a7b, 0x3f9c8ef8)) { return false; }
	if (has_item) {
		if (!item) {
			item.emplace();
		}
		if (!::tl2::details::MyTuple10Read(s, *item)) { return s.set_error_unknown_scenario(); }
		return true;
	}
	item.reset();
	return true;
}

bool tl2::details::MyTuple10MaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::MyTuplen<10>>& item) {
	if (!s.nat_write(item ? 0x3f9c8ef8 : 0x27930a7b)) { return false; }
	if (item) {
		if (!::tl2::details::MyTuple10Write(s, *item)) { return s.set_error_unknown_scenario(); }
	}
	return true;
}

bool tl2::MyTwoDicts::write_json(std::ostream& s)const {
	if (!::tl2::details::MyTwoDictsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::MyTwoDicts::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyTwoDictsRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyTwoDicts::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyTwoDictsWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyTwoDicts::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::MyTwoDicts::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::MyTwoDicts::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyTwoDictsReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::MyTwoDicts::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyTwoDictsWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::MyTwoDicts::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::MyTwoDicts::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::MyTwoDictsReset(::tl2::MyTwoDicts& item) noexcept {
	item.a.clear();
	item.b.clear();
}

bool tl2::details::MyTwoDictsWriteJSON(std::ostream& s, const ::tl2::MyTwoDicts& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.a.size() != 0) {
		add_comma = true;
		s << "\"a\":";
		if (!::tl2::details::BuiltinVectorDictionaryFieldIntWriteJSON(s, item.a)) { return false; }
	}
	if (item.b.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"b\":";
		if (!::tl2::details::BuiltinVectorDictionaryFieldIntWriteJSON(s, item.b)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::MyTwoDictsRead(::basictl::tl_istream & s, ::tl2::MyTwoDicts& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntRead(s, item.a)) { return false; }
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntRead(s, item.b)) { return false; }
	return true;
}

bool tl2::details::MyTwoDictsWrite(::basictl::tl_ostream & s, const ::tl2::MyTwoDicts& item) noexcept {
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWrite(s, item.a)) { return false; }
	if (!::tl2::details::BuiltinVectorDictionaryFieldIntWrite(s, item.b)) { return false; }
	return true;
}

bool tl2::details::MyTwoDictsReadBoxed(::basictl::tl_istream & s, ::tl2::MyTwoDicts& item) {
	if (!s.nat_read_exact_tag(0xa859581d)) { return false; }
	return tl2::details::MyTwoDictsRead(s, item);
}

bool tl2::details::MyTwoDictsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyTwoDicts& item) {
	if (!s.nat_write(0xa859581d)) { return false; }
	return tl2::details::MyTwoDictsWrite(s, item);
}

static const std::string_view MyValue_tbl_tl_name[]{"myInt", "myString"};
static const uint32_t MyValue_tbl_tl_tag[]{0xc12375b7, 0xc8bfa969};

bool tl2::MyValue::write_json(std::ostream & s)const {
	if (!::tl2::details::MyValueWriteJSON(s, *this)) { return false; }
	return true;
}
bool tl2::MyValue::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::MyValueReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	return true;
}
bool tl2::MyValue::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::MyValueWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	return true;
}

void tl2::MyValue::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::MyValue::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

std::string_view tl2::MyValue::tl_name() const {
	return MyValue_tbl_tl_name[value.index()];
}
uint32_t tl2::MyValue::tl_tag() const {
	return MyValue_tbl_tl_tag[value.index()];
}


void tl2::details::MyValueReset(::tl2::MyValue& item) noexcept{
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::MyValueWriteJSON(std::ostream & s, const ::tl2::MyValue& item) noexcept {
	s << "{";
	s << "\"type\":";
	s << "\"" << MyValue_tbl_tl_name[item.value.index()] << "\"";
	switch (item.value.index()) {
	case 0:
		s << ",\"value\":";
		if (!::tl2::details::MyIntWriteJSON(s, std::get<0>(item.value))) { return false; }
		break;
	case 1:
		s << ",\"value\":";
		if (!::tl2::details::MyStringWriteJSON(s, std::get<1>(item.value))) { return false; }
		break;
	}
	s << "}";
	return true;
}
bool tl2::details::MyValueReadBoxed(::basictl::tl_istream & s, ::tl2::MyValue& item) noexcept {
	uint32_t nat;
	if (!s.nat_read(nat)) { return false; }
	switch (nat) {
	case 0xc12375b7:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		if (!::tl2::details::MyIntRead(s, std::get<0>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	case 0xc8bfa969:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		if (!::tl2::details::MyStringRead(s, std::get<1>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::MyValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyValue& item) noexcept{
	if (!s.nat_write(MyValue_tbl_tl_tag[item.value.index()])) { return false; }
	switch (item.value.index()) {
	case 0:
		if (!::tl2::details::MyIntWrite(s, std::get<0>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	case 1:
		if (!::tl2::details::MyStringWrite(s, std::get<1>(item.value))) { return s.set_error_unknown_scenario(); }
		break;
	}
	return true;
}

bool tl2::NonOptNat::write_json(std::ostream& s)const {
	if (!::tl2::details::NonOptNatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::NonOptNat::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::NonOptNatRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::NonOptNat::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::NonOptNatWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::NonOptNat::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::NonOptNat::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::NonOptNat::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::NonOptNatReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::NonOptNat::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::NonOptNatWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::NonOptNat::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::NonOptNat::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::NonOptNatReset(::tl2::NonOptNat& item) noexcept {
	item.n = 0;
	item.xs.clear();
}

bool tl2::details::NonOptNatWriteJSON(std::ostream& s, const ::tl2::NonOptNat& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.n != 0) {
		add_comma = true;
		s << "\"n\":";
		s << item.n;
	}
	if ((item.xs.size() != 0) || (item.n != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"xs\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.xs, item.n)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::NonOptNatRead(::basictl::tl_istream & s, ::tl2::NonOptNat& item) noexcept {
	if (!s.nat_read(item.n)) { return false; }
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.xs, item.n)) { return false; }
	return true;
}

bool tl2::details::NonOptNatWrite(::basictl::tl_ostream & s, const ::tl2::NonOptNat& item) noexcept {
	if (!s.nat_write(item.n)) { return false;}
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.xs, item.n)) { return false; }
	return true;
}

bool tl2::details::NonOptNatReadBoxed(::basictl::tl_istream & s, ::tl2::NonOptNat& item) {
	if (!s.nat_read_exact_tag(0x45366605)) { return false; }
	return tl2::details::NonOptNatRead(s, item);
}

bool tl2::details::NonOptNatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::NonOptNat& item) {
	if (!s.nat_write(0x45366605)) { return false; }
	return tl2::details::NonOptNatWrite(s, item);
}

void tl2::details::RightService6ErrorVectorService6FindResultRowReset(::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept {
	item.value.clear();
}

bool tl2::details::RightService6ErrorVectorService6FindResultRowWriteJSON(std::ostream& s, const ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept {
	s << "{";
	if (item.value.size() != 0) {
		s << "\"value\":";
		if (!::tl2::details::BuiltinVectorService6FindResultRowWriteJSON(s, item.value)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::RightService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService6FindResultRowRead(s, item.value)) { return false; }
	return true;
}

bool tl2::details::RightService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService6FindResultRowWrite(s, item.value)) { return false; }
	return true;
}

bool tl2::details::RightService6ErrorVectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	if (!s.nat_read_exact_tag(0xdf3ecb3b)) { return false; }
	return tl2::details::RightService6ErrorVectorService6FindResultRowRead(s, item);
}

bool tl2::details::RightService6ErrorVectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	if (!s.nat_write(0xdf3ecb3b)) { return false; }
	return tl2::details::RightService6ErrorVectorService6FindResultRowWrite(s, item);
}

bool tl2::RpcInvokeReqExtra::write_json(std::ostream& s)const {
	if (!::tl2::details::RpcInvokeReqExtraWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::RpcInvokeReqExtra::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::RpcInvokeReqExtraRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::RpcInvokeReqExtra::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::RpcInvokeReqExtraWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::RpcInvokeReqExtra::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::RpcInvokeReqExtra::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::RpcInvokeReqExtra::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::RpcInvokeReqExtraReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::RpcInvokeReqExtra::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::RpcInvokeReqExtraWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::RpcInvokeReqExtra::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::RpcInvokeReqExtra::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::RpcInvokeReqExtraReset(::tl2::RpcInvokeReqExtra& item) noexcept {
	item.fields_mask = 0;
	::tl2::details::TrueReset(item.query);
	::tl2::details::TrueReset(item.sort);
	::tl2::details::TrueReset(item.sort_reverse);
	item.wait_binlog_pos = 0;
	item.string_forward_keys.clear();
}

bool tl2::details::RpcInvokeReqExtraWriteJSON(std::ostream& s, const ::tl2::RpcInvokeReqExtra& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.fields_mask != 0) {
		add_comma = true;
		s << "\"fields_mask\":";
		s << item.fields_mask;
	}
	if ((item.fields_mask & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"query\":";
		if (!::tl2::details::TrueWriteJSON(s, item.query)) { return false; }
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"sort\":";
		if (!::tl2::details::TrueWriteJSON(s, item.sort)) { return false; }
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"sort_reverse\":";
		if (!::tl2::details::TrueWriteJSON(s, item.sort_reverse)) { return false; }
	}
	if ((item.fields_mask & (1<<16)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"wait_binlog_pos\":";
		s << item.wait_binlog_pos;
	}
	if ((item.fields_mask & (1<<18)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"string_forward_keys\":";
		if (!::tl2::details::BuiltinVectorStringWriteJSON(s, item.string_forward_keys)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::RpcInvokeReqExtraRead(::basictl::tl_istream & s, ::tl2::RpcInvokeReqExtra& item) noexcept {
	if (!s.nat_read(item.fields_mask)) { return false; }
	if ((item.fields_mask & (1<<0)) != 0) {
		if (!::tl2::details::TrueRead(s, item.query)) { return s.set_error_unknown_scenario(); }
	} else {
			::tl2::details::TrueReset(item.query);
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (!::tl2::details::TrueRead(s, item.sort)) { return s.set_error_unknown_scenario(); }
	} else {
			::tl2::details::TrueReset(item.sort);
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (!::tl2::details::TrueRead(s, item.sort_reverse)) { return s.set_error_unknown_scenario(); }
	} else {
			::tl2::details::TrueReset(item.sort_reverse);
	}
	if ((item.fields_mask & (1<<16)) != 0) {
		if (!s.long_read(item.wait_binlog_pos)) { return false; }
	} else {
			item.wait_binlog_pos = 0;
	}
	if ((item.fields_mask & (1<<18)) != 0) {
		if (!::tl2::details::BuiltinVectorStringRead(s, item.string_forward_keys)) { return false; }
	} else {
			item.string_forward_keys.clear();
	}
	return true;
}

bool tl2::details::RpcInvokeReqExtraWrite(::basictl::tl_ostream & s, const ::tl2::RpcInvokeReqExtra& item) noexcept {
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.query)) { return s.set_error_unknown_scenario(); }
	}
	if ((item.fields_mask & (1<<1)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.sort)) { return s.set_error_unknown_scenario(); }
	}
	if ((item.fields_mask & (1<<2)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.sort_reverse)) { return s.set_error_unknown_scenario(); }
	}
	if ((item.fields_mask & (1<<16)) != 0) {
			if (!s.long_write(item.wait_binlog_pos)) { return false;}
	}
	if ((item.fields_mask & (1<<18)) != 0) {
			if (!::tl2::details::BuiltinVectorStringWrite(s, item.string_forward_keys)) { return false; }
	}
	return true;
}

bool tl2::details::RpcInvokeReqExtraReadBoxed(::basictl::tl_istream & s, ::tl2::RpcInvokeReqExtra& item) {
	if (!s.nat_read_exact_tag(0xf3ef81a9)) { return false; }
	return tl2::details::RpcInvokeReqExtraRead(s, item);
}

bool tl2::details::RpcInvokeReqExtraWriteBoxed(::basictl::tl_ostream & s, const ::tl2::RpcInvokeReqExtra& item) {
	if (!s.nat_write(0xf3ef81a9)) { return false; }
	return tl2::details::RpcInvokeReqExtraWrite(s, item);
}

bool tl2::StatOne::write_json(std::ostream& s)const {
	if (!::tl2::details::StatOneWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::StatOne::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::StatOneRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::StatOne::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::StatOneWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::StatOne::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::StatOne::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::StatOne::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::StatOneReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::StatOne::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::StatOneWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::StatOne::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::StatOne::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::StatOneReset(::tl2::StatOne& item) noexcept {
	item.key.clear();
	item.value.clear();
}

bool tl2::details::StatOneWriteJSON(std::ostream& s, const ::tl2::StatOne& item) noexcept {
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

bool tl2::details::StatOneRead(::basictl::tl_istream & s, ::tl2::StatOne& item) noexcept {
	if (!s.string_read(item.key)) { return false; }
	if (!s.string_read(item.value)) { return false; }
	return true;
}

bool tl2::details::StatOneWrite(::basictl::tl_ostream & s, const ::tl2::StatOne& item) noexcept {
	if (!s.string_write(item.key)) { return false;}
	if (!s.string_write(item.value)) { return false;}
	return true;
}

bool tl2::details::StatOneReadBoxed(::basictl::tl_istream & s, ::tl2::StatOne& item) {
	if (!s.nat_read_exact_tag(0x74b0604b)) { return false; }
	return tl2::details::StatOneRead(s, item);
}

bool tl2::details::StatOneWriteBoxed(::basictl::tl_ostream & s, const ::tl2::StatOne& item) {
	if (!s.nat_write(0x74b0604b)) { return false; }
	return tl2::details::StatOneWrite(s, item);
}

void tl2::details::StringReset(std::string& item) noexcept {
	item.clear();
}

bool tl2::details::StringWriteJSON(std::ostream& s, const std::string& item) noexcept {
	s << "\"" << item << "\"";
	return true;
}

bool tl2::details::StringRead(::basictl::tl_istream & s, std::string& item) noexcept {
	if (!s.string_read(item)) { return false; }
	return true;
}

bool tl2::details::StringWrite(::basictl::tl_ostream & s, const std::string& item) noexcept {
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
	if (!::tl2::details::TrueRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::True::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TrueWrite(s, *this)) { return s.set_error_unknown_scenario(); }
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
	if (!::tl2::details::TrueReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::True::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TrueWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
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

void tl2::details::TupleIntReset(std::vector<int32_t>& item) noexcept {
	item.clear();
}

bool tl2::details::TupleIntWriteJSON(std::ostream& s, const std::vector<int32_t>& item, uint32_t nat_n) noexcept {
	if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n) noexcept {
	if (!::tl2::details::BuiltinTupleIntRead(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) noexcept {
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

void tl2::details::TupleInt3Reset(std::array<int32_t, 3>& item) noexcept {
	::tl2::details::BuiltinTuple3IntReset(item);
}

bool tl2::details::TupleInt3WriteJSON(std::ostream& s, const std::array<int32_t, 3>& item) noexcept {
	if (!::tl2::details::BuiltinTuple3IntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt3Read(::basictl::tl_istream & s, std::array<int32_t, 3>& item) noexcept {
	if (!::tl2::details::BuiltinTuple3IntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt3Write(::basictl::tl_ostream & s, const std::array<int32_t, 3>& item) noexcept {
	if (!::tl2::details::BuiltinTuple3IntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt3ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 3>& item) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleInt3Read(s, item);
}

bool tl2::details::TupleInt3WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 3>& item) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleInt3Write(s, item);
}

void tl2::details::TupleIntBoxedReset(std::vector<int32_t>& item) noexcept {
	item.clear();
}

bool tl2::details::TupleIntBoxedWriteJSON(std::ostream& s, const std::vector<int32_t>& item, uint32_t nat_n) noexcept {
	if (!::tl2::details::BuiltinTupleIntBoxedWriteJSON(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n) noexcept {
	if (!::tl2::details::BuiltinTupleIntBoxedRead(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) noexcept {
	if (!::tl2::details::BuiltinTupleIntBoxedWrite(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxedReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleIntBoxedRead(s, item, nat_n);
}

bool tl2::details::TupleIntBoxedWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleIntBoxedWrite(s, item, nat_n);
}

void tl2::details::TupleIntBoxed10Reset(std::array<int32_t, 10>& item) noexcept {
	::tl2::details::BuiltinTuple10IntBoxedReset(item);
}

bool tl2::details::TupleIntBoxed10WriteJSON(std::ostream& s, const std::array<int32_t, 10>& item) noexcept {
	if (!::tl2::details::BuiltinTuple10IntBoxedWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxed10Read(::basictl::tl_istream & s, std::array<int32_t, 10>& item) noexcept {
	if (!::tl2::details::BuiltinTuple10IntBoxedRead(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxed10Write(::basictl::tl_ostream & s, const std::array<int32_t, 10>& item) noexcept {
	if (!::tl2::details::BuiltinTuple10IntBoxedWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxed10ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 10>& item) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleIntBoxed10Read(s, item);
}

bool tl2::details::TupleIntBoxed10WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 10>& item) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleIntBoxed10Write(s, item);
}

void tl2::details::TupleIntBoxed2Reset(std::array<int32_t, 2>& item) noexcept {
	::tl2::details::BuiltinTuple2IntBoxedReset(item);
}

bool tl2::details::TupleIntBoxed2WriteJSON(std::ostream& s, const std::array<int32_t, 2>& item) noexcept {
	if (!::tl2::details::BuiltinTuple2IntBoxedWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxed2Read(::basictl::tl_istream & s, std::array<int32_t, 2>& item) noexcept {
	if (!::tl2::details::BuiltinTuple2IntBoxedRead(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxed2Write(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item) noexcept {
	if (!::tl2::details::BuiltinTuple2IntBoxedWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxed2ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 2>& item) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleIntBoxed2Read(s, item);
}

bool tl2::details::TupleIntBoxed2WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleIntBoxed2Write(s, item);
}

void tl2::details::TupleService1Value3Reset(std::array<::tl2::service1::Value, 3>& item) noexcept {
	::tl2::details::BuiltinTuple3Service1ValueReset(item);
}

bool tl2::details::TupleService1Value3WriteJSON(std::ostream& s, const std::array<::tl2::service1::Value, 3>& item) noexcept {
	if (!::tl2::details::BuiltinTuple3Service1ValueWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleService1Value3Read(::basictl::tl_istream & s, std::array<::tl2::service1::Value, 3>& item) noexcept {
	if (!::tl2::details::BuiltinTuple3Service1ValueRead(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleService1Value3Write(::basictl::tl_ostream & s, const std::array<::tl2::service1::Value, 3>& item) noexcept {
	if (!::tl2::details::BuiltinTuple3Service1ValueWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleService1Value3ReadBoxed(::basictl::tl_istream & s, std::array<::tl2::service1::Value, 3>& item) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleService1Value3Read(s, item);
}

bool tl2::details::TupleService1Value3WriteBoxed(::basictl::tl_ostream & s, const std::array<::tl2::service1::Value, 3>& item) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleService1Value3Write(s, item);
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

void tl2::details::VectorEitherService6ErrorVectorService6FindResultRowReset(std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorEitherService6ErrorVectorService6FindResultRowWriteJSON(std::ostream& s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorEitherService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorEitherService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) noexcept {
	if (!::tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorEitherService6ErrorVectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorEitherService6ErrorVectorService6FindResultRowRead(s, item);
}

bool tl2::details::VectorEitherService6ErrorVectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorEitherService6ErrorVectorService6FindResultRowWrite(s, item);
}

void tl2::details::VectorIntReset(std::vector<int32_t>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorIntWriteJSON(std::ostream& s, const std::vector<int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item) noexcept {
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

void tl2::details::VectorIntBoxedReset(std::vector<int32_t>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorIntBoxedWriteJSON(std::ostream& s, const std::vector<int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntBoxedWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntBoxedRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntBoxedWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntBoxedReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorIntBoxedRead(s, item);
}

bool tl2::details::VectorIntBoxedWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorIntBoxedWrite(s, item);
}

void tl2::details::VectorIntegerReset(std::vector<::tl2::Integer>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorIntegerWriteJSON(std::ostream& s, const std::vector<::tl2::Integer>& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntegerWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntegerRead(::basictl::tl_istream & s, std::vector<::tl2::Integer>& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntegerRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntegerWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Integer>& item) noexcept {
	if (!::tl2::details::BuiltinVectorIntegerWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntegerReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::Integer>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorIntegerRead(s, item);
}

bool tl2::details::VectorIntegerWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::Integer>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorIntegerWrite(s, item);
}

void tl2::details::VectorLongBoxedReset(std::vector<int64_t>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorLongBoxedWriteJSON(std::ostream& s, const std::vector<int64_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorLongBoxedWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorLongBoxedRead(::basictl::tl_istream & s, std::vector<int64_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorLongBoxedRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorLongBoxedWrite(::basictl::tl_ostream & s, const std::vector<int64_t>& item) noexcept {
	if (!::tl2::details::BuiltinVectorLongBoxedWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorLongBoxedReadBoxed(::basictl::tl_istream & s, std::vector<int64_t>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorLongBoxedRead(s, item);
}

bool tl2::details::VectorLongBoxedWriteBoxed(::basictl::tl_ostream & s, const std::vector<int64_t>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorLongBoxedWrite(s, item);
}

void tl2::details::VectorService1ValueReset(std::vector<::tl2::service1::Value>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorService1ValueWriteJSON(std::ostream& s, const std::vector<::tl2::service1::Value>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService1ValueWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService1ValueRead(::basictl::tl_istream & s, std::vector<::tl2::service1::Value>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService1ValueRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService1ValueWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service1::Value>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService1ValueWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService1ValueReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service1::Value>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorService1ValueRead(s, item);
}

bool tl2::details::VectorService1ValueWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service1::Value>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorService1ValueWrite(s, item);
}

void tl2::details::VectorService6FindResultRowReset(std::vector<::tl2::service6::FindResultRow>& item) noexcept {
	item.clear();
}

bool tl2::details::VectorService6FindResultRowWriteJSON(std::ostream& s, const std::vector<::tl2::service6::FindResultRow>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService6FindResultRowWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService6FindResultRowRead(::basictl::tl_istream & s, std::vector<::tl2::service6::FindResultRow>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService6FindResultRowRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindResultRow>& item) noexcept {
	if (!::tl2::details::BuiltinVectorService6FindResultRowWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service6::FindResultRow>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorService6FindResultRowRead(s, item);
}

bool tl2::details::VectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindResultRow>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorService6FindResultRowWrite(s, item);
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

bool tl2::WithFloat::write_json(std::ostream& s)const {
	if (!::tl2::details::WithFloatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::WithFloat::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::WithFloatRead(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::WithFloat::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::WithFloatWrite(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::WithFloat::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::WithFloat::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::WithFloat::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::WithFloatReadBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

bool tl2::WithFloat::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::WithFloatWriteBoxed(s, *this)) { return s.set_error_unknown_scenario(); }
	s.last_release();
	return true;
}

void tl2::WithFloat::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::WithFloat::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::WithFloatReset(::tl2::WithFloat& item) noexcept {
	item.x = 0;
	item.y = 0;
	item.z = 0;
}

bool tl2::details::WithFloatWriteJSON(std::ostream& s, const ::tl2::WithFloat& item) noexcept {
	auto add_comma = false;
	s << "{";
	if (item.x != 0) {
		add_comma = true;
		s << "\"x\":";
		s << item.x;
	}
	if (item.y != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"y\":";
		s << item.y;
	}
	if (item.z != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"z\":";
		s << item.z;
	}
	s << "}";
	return true;
}

bool tl2::details::WithFloatRead(::basictl::tl_istream & s, ::tl2::WithFloat& item) noexcept {
	if (!s.float_read(item.x)) { return false; }
	if (!s.nat_read_exact_tag(0x824dab22)) { return false; }
	if (!s.float_read(item.y)) { return false; }
	if (!s.float_read(item.z)) { return false; }
	return true;
}

bool tl2::details::WithFloatWrite(::basictl::tl_ostream & s, const ::tl2::WithFloat& item) noexcept {
	if (!s.float_write(item.x)) { return false;}
	if (!s.nat_write(0x824dab22)) { return false; }
	if (!s.float_write(item.y)) { return false;}
	if (!s.float_write(item.z)) { return false;}
	return true;
}

bool tl2::details::WithFloatReadBoxed(::basictl::tl_istream & s, ::tl2::WithFloat& item) {
	if (!s.nat_read_exact_tag(0x071b8685)) { return false; }
	return tl2::details::WithFloatRead(s, item);
}

bool tl2::details::WithFloatWriteBoxed(::basictl::tl_ostream & s, const ::tl2::WithFloat& item) {
	if (!s.nat_write(0x071b8685)) { return false; }
	return tl2::details::WithFloatWrite(s, item);
}
