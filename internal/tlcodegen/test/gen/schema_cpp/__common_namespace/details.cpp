#include "headers/withFloat.hpp"
#include "headers/vector.hpp"
#include "headers/tuple.hpp"
#include "../tasks/headers/tasks.queueTypeInfo.hpp"
#include "headers/string.hpp"
#include "headers/statOne.hpp"
#include "../service6/headers/service6.findWithBoundsResult.hpp"
#include "../service2/headers/service2.counterSet.hpp"
#include "../service1/headers/service1.Value.hpp"
#include "headers/rpcInvokeReqExtra.hpp"
#include "headers/true.hpp"
#include "headers/nonOptNat.hpp"
#include "headers/myTwoDicts.hpp"
#include "headers/myMcValueVector.hpp"
#include "headers/myMcValueTuple.hpp"
#include "headers/myMcValue.hpp"
#include "headers/myMaybe2.hpp"
#include "headers/myMaybe1.hpp"
#include "headers/myTuple.hpp"
#include "headers/myBoxedVectorSlice.hpp"
#include "headers/myAnonMcValue.hpp"
#include "headers/map.hpp"
#include "headers/long.hpp"
#include "headers/issue3498.hpp"
#include "../service6/headers/service6.findResultRow.hpp"
#include "../service6/headers/service6.error.hpp"
#include "headers/int.hpp"
#include "headers/getStats.hpp"
#include "../tasks/headers/tasks.queueTypeStats.hpp"
#include "headers/getNonOptNat.hpp"
#include "headers/getMyValue.hpp"
#include "headers/MyValue.hpp"
#include "headers/myString.hpp"
#include "headers/myInt.hpp"
#include "headers/getMyDouble.hpp"
#include "headers/myDouble.hpp"
#include "headers/getMyDictOfInt.hpp"
#include "headers/myDictOfInt.hpp"
#include "headers/getMaybeIface.hpp"
#include "headers/getFloat.hpp"
#include "headers/getDouble.hpp"
#include "headers/get_arrays.hpp"
#include "headers/float.hpp"
#include "headers/fieldConflict4.hpp"
#include "headers/fieldConflict3.hpp"
#include "headers/fieldConflict2.hpp"
#include "headers/fieldConflict1.hpp"
#include "headers/Either.hpp"
#include "headers/right.hpp"
#include "headers/left.hpp"
#include "headers/double.hpp"
#include "headers/dictionary.hpp"
#include "headers/dictionaryField.hpp"
#include "headers/integer.hpp"
#include "headers/boxedVector64.hpp"
#include "headers/boxedVector32BoxedElem.hpp"
#include "headers/boxedVector32.hpp"
#include "headers/boxedTupleSlice3.hpp"
#include "headers/boxedTupleSlice2.hpp"
#include "headers/myBoxedTupleSlice.hpp"
#include "headers/boxedTupleSlice1.hpp"
#include "headers/boxedTuple.hpp"
#include "headers/boxedString.hpp"
#include "headers/boxedInt.hpp"
#include "headers/boxedArray.hpp"
#include "headers/myBoxedArray.hpp"
#include "headers/Bool.hpp"
#include "headers/benchObject.hpp"


bool tl2::BenchObject::write_json(std::ostream& s)const {
	if (!::tl2::details::BenchObjectWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BenchObject::read(::basictl::tl_istream & s) {
	if (!::tl2::details::BenchObjectRead(s, *this)) { return false; }
	return true;
}

bool tl2::BenchObject::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BenchObjectWrite(s, *this)) { return false; }
	return true;
}

bool tl2::BenchObject::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::BenchObjectReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::BenchObject::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BenchObjectWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::BenchObjectReset(::tl2::BenchObject& item) {
	item.xs.clear();
	item.ys.clear();
}

bool tl2::details::BenchObjectWriteJSON(std::ostream& s, const ::tl2::BenchObject& item) {
	s << "{";
	s << "\"xs\":";
	if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.xs)) { return false; }
	s << ",";
	s << "\"ys\":";
	if (!::tl2::details::BuiltinVectorIntegerWriteJSON(s, item.ys)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BenchObjectRead(::basictl::tl_istream & s, ::tl2::BenchObject& item) {
	if (!::tl2::details::BuiltinVectorIntRead(s, item.xs)) { return false; }
	if (!::tl2::details::BuiltinVectorIntegerRead(s, item.ys)) { return false; }
	return true;
}

bool tl2::details::BenchObjectWrite(::basictl::tl_ostream & s, const ::tl2::BenchObject& item) {
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

bool tl2::BoxedArray::read(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedArrayRead(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedArray::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedArrayWrite(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedArray::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedArrayReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedArray::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedArrayWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::BoxedArrayReset(::tl2::BoxedArray& item) {
	::tl2::details::MyBoxedArrayReset(item.x);
}

bool tl2::details::BoxedArrayWriteJSON(std::ostream& s, const ::tl2::BoxedArray& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::MyBoxedArrayWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BoxedArrayRead(::basictl::tl_istream & s, ::tl2::BoxedArray& item) {
	if (!::tl2::details::MyBoxedArrayReadBoxed(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedArrayWrite(::basictl::tl_ostream & s, const ::tl2::BoxedArray& item) {
	if (!::tl2::details::MyBoxedArrayWriteBoxed(s, item.x)) { return false; }
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
	if (!::tl2::details::MyBoxedArrayReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::BoxedArrayWriteResult(::basictl::tl_ostream & s, tl2::BoxedArray& item, ::tl2::MyBoxedArray& result) {
	if (!::tl2::details::MyBoxedArrayWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::BoxedArray::read_result(::basictl::tl_istream & s, ::tl2::MyBoxedArray & result) {
	return tl2::details::BoxedArrayReadResult(s, *this, result);
}
bool tl2::BoxedArray::write_result(::basictl::tl_ostream & s, ::tl2::MyBoxedArray & result) {
	return tl2::details::BoxedArrayWriteResult(s, *this, result);
}

bool tl2::BoxedInt::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedIntWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedInt::read(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedIntRead(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedInt::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedIntWrite(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedInt::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedIntReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedInt::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedIntWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::BoxedIntReset(::tl2::BoxedInt& item) {
	item.x = 0;
}

bool tl2::details::BoxedIntWriteJSON(std::ostream& s, const ::tl2::BoxedInt& item) {
	s << "{";
	s << "\"x\":";
	s << item.x;
	s << "}";
	return true;
}

bool tl2::details::BoxedIntRead(::basictl::tl_istream & s, ::tl2::BoxedInt& item) {
	if (!s.nat_read_exact_tag(0xa8509bda)) { return false;}
	if (!s.int_read(item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedIntWrite(::basictl::tl_ostream & s, const ::tl2::BoxedInt& item) {
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
	if (!s.nat_read_exact_tag(0xa8509bda)) { return false;}
	if (!s.int_read(result)) { return false; }
	return true;
}
bool tl2::details::BoxedIntWriteResult(::basictl::tl_ostream & s, tl2::BoxedInt& item, int32_t& result) {
	if (!s.nat_write(0xa8509bda)) { return false; }
	if (!s.int_write(result)) { return false;}
	return true;
}

bool tl2::BoxedInt::read_result(::basictl::tl_istream & s, int32_t & result) {
	return tl2::details::BoxedIntReadResult(s, *this, result);
}
bool tl2::BoxedInt::write_result(::basictl::tl_ostream & s, int32_t & result) {
	return tl2::details::BoxedIntWriteResult(s, *this, result);
}

bool tl2::BoxedString::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedStringWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedString::read(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedStringRead(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedString::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedStringWrite(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedString::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedStringReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedString::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedStringWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::BoxedStringReset(::tl2::BoxedString& item) {
	item.x.clear();
}

bool tl2::details::BoxedStringWriteJSON(std::ostream& s, const ::tl2::BoxedString& item) {
	s << "{";
	s << "\"x\":";
	s << "\"" << item.x << "\"";
	s << "}";
	return true;
}

bool tl2::details::BoxedStringRead(::basictl::tl_istream & s, ::tl2::BoxedString& item) {
	if (!s.nat_read_exact_tag(0xb5286e24)) { return false;}
	if (!s.string_read(item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedStringWrite(::basictl::tl_ostream & s, const ::tl2::BoxedString& item) {
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
	if (!s.nat_read_exact_tag(0xb5286e24)) { return false;}
	if (!s.string_read(result)) { return false; }
	return true;
}
bool tl2::details::BoxedStringWriteResult(::basictl::tl_ostream & s, tl2::BoxedString& item, std::string& result) {
	if (!s.nat_write(0xb5286e24)) { return false; }
	if (!s.string_write(result)) { return false;}
	return true;
}

bool tl2::BoxedString::read_result(::basictl::tl_istream & s, std::string & result) {
	return tl2::details::BoxedStringReadResult(s, *this, result);
}
bool tl2::BoxedString::write_result(::basictl::tl_ostream & s, std::string & result) {
	return tl2::details::BoxedStringWriteResult(s, *this, result);
}

bool tl2::BoxedTuple::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedTupleWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTuple::read(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedTupleRead(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTuple::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedTupleWrite(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTuple::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedTupleReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTuple::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedTupleWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::BoxedTupleReset(::tl2::BoxedTuple& item) {
	::tl2::details::BuiltinTuple3IntReset(item.x);
}

bool tl2::details::BoxedTupleWriteJSON(std::ostream& s, const ::tl2::BoxedTuple& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::BuiltinTuple3IntWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BoxedTupleRead(::basictl::tl_istream & s, ::tl2::BoxedTuple& item) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTuple3IntRead(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedTupleWrite(::basictl::tl_ostream & s, const ::tl2::BoxedTuple& item) {
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
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTuple3IntRead(s, result)) { return false; }
	return true;
}
bool tl2::details::BoxedTupleWriteResult(::basictl::tl_ostream & s, tl2::BoxedTuple& item, std::array<int32_t, 3>& result) {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTuple3IntWrite(s, result)) { return false; }
	return true;
}

bool tl2::BoxedTuple::read_result(::basictl::tl_istream & s, std::array<int32_t, 3> & result) {
	return tl2::details::BoxedTupleReadResult(s, *this, result);
}
bool tl2::BoxedTuple::write_result(::basictl::tl_ostream & s, std::array<int32_t, 3> & result) {
	return tl2::details::BoxedTupleWriteResult(s, *this, result);
}

bool tl2::BoxedTupleSlice1::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedTupleSlice1WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice1::read(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedTupleSlice1Read(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice1::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedTupleSlice1Write(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice1::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedTupleSlice1ReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice1::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedTupleSlice1WriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::BoxedTupleSlice1Reset(::tl2::BoxedTupleSlice1& item) {
	item.n = 0;
	item.x.clear();
}

bool tl2::details::BoxedTupleSlice1WriteJSON(std::ostream& s, const ::tl2::BoxedTupleSlice1& item) {
	s << "{";
	s << "\"n\":";
	s << item.n;
	s << ",";
	s << "\"x\":";
	if (!::tl2::details::BuiltinTupleIntBoxedWriteJSON(s, item.x, item.n)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BoxedTupleSlice1Read(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice1& item) {
	if (!s.nat_read(item.n)) { return false; }
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTupleIntBoxedRead(s, item.x, item.n)) { return false; }
	return true;
}

bool tl2::details::BoxedTupleSlice1Write(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice1& item) {
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
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTupleIntBoxedRead(s, result, item.n)) { return false; }
	return true;
}
bool tl2::details::BoxedTupleSlice1WriteResult(::basictl::tl_ostream & s, tl2::BoxedTupleSlice1& item, std::vector<int32_t>& result) {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntBoxedWrite(s, result, item.n)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice1::read_result(::basictl::tl_istream & s, std::vector<int32_t> & result) {
	return tl2::details::BoxedTupleSlice1ReadResult(s, *this, result);
}
bool tl2::BoxedTupleSlice1::write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result) {
	return tl2::details::BoxedTupleSlice1WriteResult(s, *this, result);
}

bool tl2::BoxedTupleSlice2::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedTupleSlice2WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice2::read(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedTupleSlice2Read(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice2::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedTupleSlice2Write(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice2::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedTupleSlice2ReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice2::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedTupleSlice2WriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::BoxedTupleSlice2Reset(::tl2::BoxedTupleSlice2& item) {
	::tl2::details::MyBoxedTupleSliceReset(item.x);
}

bool tl2::details::BoxedTupleSlice2WriteJSON(std::ostream& s, const ::tl2::BoxedTupleSlice2& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::MyBoxedTupleSliceWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BoxedTupleSlice2Read(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice2& item) {
	if (!::tl2::details::MyBoxedTupleSliceReadBoxed(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedTupleSlice2Write(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice2& item) {
	if (!::tl2::details::MyBoxedTupleSliceWriteBoxed(s, item.x)) { return false; }
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
	if (!::tl2::details::MyBoxedTupleSliceReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::BoxedTupleSlice2WriteResult(::basictl::tl_ostream & s, tl2::BoxedTupleSlice2& item, ::tl2::MyBoxedTupleSlice& result) {
	if (!::tl2::details::MyBoxedTupleSliceWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice2::read_result(::basictl::tl_istream & s, ::tl2::MyBoxedTupleSlice & result) {
	return tl2::details::BoxedTupleSlice2ReadResult(s, *this, result);
}
bool tl2::BoxedTupleSlice2::write_result(::basictl::tl_ostream & s, ::tl2::MyBoxedTupleSlice & result) {
	return tl2::details::BoxedTupleSlice2WriteResult(s, *this, result);
}

bool tl2::BoxedTupleSlice3::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedTupleSlice3WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice3::read(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedTupleSlice3Read(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice3::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedTupleSlice3Write(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice3::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedTupleSlice3ReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice3::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedTupleSlice3WriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::BoxedTupleSlice3Reset(::tl2::BoxedTupleSlice3& item) {
	item.n = 0;
	item.x.clear();
}

bool tl2::details::BoxedTupleSlice3WriteJSON(std::ostream& s, const ::tl2::BoxedTupleSlice3& item) {
	s << "{";
	s << "\"n\":";
	s << item.n;
	s << ",";
	s << "\"x\":";
	if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.x, item.n)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BoxedTupleSlice3Read(::basictl::tl_istream & s, ::tl2::BoxedTupleSlice3& item) {
	if (!s.nat_read(item.n)) { return false; }
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTupleIntRead(s, item.x, item.n)) { return false; }
	return true;
}

bool tl2::details::BoxedTupleSlice3Write(::basictl::tl_ostream & s, const ::tl2::BoxedTupleSlice3& item) {
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
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTupleIntRead(s, result, item.n)) { return false; }
	return true;
}
bool tl2::details::BoxedTupleSlice3WriteResult(::basictl::tl_ostream & s, tl2::BoxedTupleSlice3& item, std::vector<int32_t>& result) {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntWrite(s, result, item.n)) { return false; }
	return true;
}

bool tl2::BoxedTupleSlice3::read_result(::basictl::tl_istream & s, std::vector<int32_t> & result) {
	return tl2::details::BoxedTupleSlice3ReadResult(s, *this, result);
}
bool tl2::BoxedTupleSlice3::write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result) {
	return tl2::details::BoxedTupleSlice3WriteResult(s, *this, result);
}

bool tl2::BoxedVector32::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedVector32WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector32::read(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedVector32Read(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector32::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedVector32Write(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector32::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedVector32ReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector32::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedVector32WriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::BoxedVector32Reset(::tl2::BoxedVector32& item) {
	item.x.clear();
}

bool tl2::details::BoxedVector32WriteJSON(std::ostream& s, const ::tl2::BoxedVector32& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BoxedVector32Read(::basictl::tl_istream & s, ::tl2::BoxedVector32& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false;}
	if (!::tl2::details::BuiltinVectorIntRead(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedVector32Write(::basictl::tl_ostream & s, const ::tl2::BoxedVector32& item) {
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
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false;}
	if (!::tl2::details::BuiltinVectorIntRead(s, result)) { return false; }
	return true;
}
bool tl2::details::BoxedVector32WriteResult(::basictl::tl_ostream & s, tl2::BoxedVector32& item, std::vector<int32_t>& result) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntWrite(s, result)) { return false; }
	return true;
}

bool tl2::BoxedVector32::read_result(::basictl::tl_istream & s, std::vector<int32_t> & result) {
	return tl2::details::BoxedVector32ReadResult(s, *this, result);
}
bool tl2::BoxedVector32::write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result) {
	return tl2::details::BoxedVector32WriteResult(s, *this, result);
}

bool tl2::BoxedVector32BoxedElem::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedVector32BoxedElemWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector32BoxedElem::read(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedVector32BoxedElemRead(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector32BoxedElem::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedVector32BoxedElemWrite(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector32BoxedElem::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedVector32BoxedElemReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector32BoxedElem::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedVector32BoxedElemWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::BoxedVector32BoxedElemReset(::tl2::BoxedVector32BoxedElem& item) {
	item.x.clear();
}

bool tl2::details::BoxedVector32BoxedElemWriteJSON(std::ostream& s, const ::tl2::BoxedVector32BoxedElem& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::BuiltinVectorIntBoxedWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BoxedVector32BoxedElemRead(::basictl::tl_istream & s, ::tl2::BoxedVector32BoxedElem& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false;}
	if (!::tl2::details::BuiltinVectorIntBoxedRead(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedVector32BoxedElemWrite(::basictl::tl_ostream & s, const ::tl2::BoxedVector32BoxedElem& item) {
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
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false;}
	if (!::tl2::details::BuiltinVectorIntBoxedRead(s, result)) { return false; }
	return true;
}
bool tl2::details::BoxedVector32BoxedElemWriteResult(::basictl::tl_ostream & s, tl2::BoxedVector32BoxedElem& item, std::vector<int32_t>& result) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorIntBoxedWrite(s, result)) { return false; }
	return true;
}

bool tl2::BoxedVector32BoxedElem::read_result(::basictl::tl_istream & s, std::vector<int32_t> & result) {
	return tl2::details::BoxedVector32BoxedElemReadResult(s, *this, result);
}
bool tl2::BoxedVector32BoxedElem::write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result) {
	return tl2::details::BoxedVector32BoxedElemWriteResult(s, *this, result);
}

bool tl2::BoxedVector64::write_json(std::ostream& s)const {
	if (!::tl2::details::BoxedVector64WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector64::read(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedVector64Read(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector64::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedVector64Write(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector64::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::BoxedVector64ReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::BoxedVector64::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::BoxedVector64WriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::BoxedVector64Reset(::tl2::BoxedVector64& item) {
	item.x.clear();
}

bool tl2::details::BoxedVector64WriteJSON(std::ostream& s, const ::tl2::BoxedVector64& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::BuiltinVectorLongBoxedWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::BoxedVector64Read(::basictl::tl_istream & s, ::tl2::BoxedVector64& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false;}
	if (!::tl2::details::BuiltinVectorLongBoxedRead(s, item.x)) { return false; }
	return true;
}

bool tl2::details::BoxedVector64Write(::basictl::tl_ostream & s, const ::tl2::BoxedVector64& item) {
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
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false;}
	if (!::tl2::details::BuiltinVectorLongBoxedRead(s, result)) { return false; }
	return true;
}
bool tl2::details::BoxedVector64WriteResult(::basictl::tl_ostream & s, tl2::BoxedVector64& item, std::vector<int64_t>& result) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorLongBoxedWrite(s, result)) { return false; }
	return true;
}

bool tl2::BoxedVector64::read_result(::basictl::tl_istream & s, std::vector<int64_t> & result) {
	return tl2::details::BoxedVector64ReadResult(s, *this, result);
}
bool tl2::BoxedVector64::write_result(::basictl::tl_ostream & s, std::vector<int64_t> & result) {
	return tl2::details::BoxedVector64WriteResult(s, *this, result);
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
		if (!s.nat_read_exact_tag(0xa8509bda)) { return false;}
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
		if (!s.nat_read_exact_tag(0xa8509bda)) { return false;}
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
	if (item.size() != nat_n) {
		// TODO add exception
		return false;
	}
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
	if (item.size() != nat_n)
		return s.set_error_sequence_length();
	for(const auto & el : item) {
		if (!s.int_write(el)) { return false;}
	}
	return true;
}

void tl2::details::BuiltinTupleIntBoxedReset(std::vector<int32_t>& item) {
	item.resize(0);
}

bool tl2::details::BuiltinTupleIntBoxedWriteJSON(std::ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) {
	if (item.size() != nat_n) {
		// TODO add exception
		return false;
	}
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
		if (!s.nat_read_exact_tag(0xa8509bda)) { return false;}
	if (!s.int_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTupleIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) {
	if (item.size() != nat_n)
		return s.set_error_sequence_length();
	for(const auto & el : item) {
		if (!s.nat_write(0xa8509bda)) { return false; }
	if (!s.int_write(el)) { return false;}
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

void tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultReset(std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultWriteJSON(std::ostream & s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::EitherIntVectorService6FindWithBoundsResultWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::EitherIntVectorService6FindWithBoundsResultReadBoxed(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::EitherIntVectorService6FindWithBoundsResultWriteBoxed(s, el)) { return false; }
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
		if (!::tl2::details::EitherService6ErrorVectorService6FindResultRowReadBoxed(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::EitherService6ErrorVectorService6FindResultRowWriteBoxed(s, el)) { return false; }
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
		if (!s.nat_read_exact_tag(0xa8509bda)) { return false;}
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
		if (!::tl2::details::IntegerRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorIntegerWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Integer>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::IntegerWrite(s, el)) { return false; }
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
		if (!s.nat_read_exact_tag(0x22076cba)) { return false;}
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
		if (!::tl2::details::MapStringStringRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorMapStringStringWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Map<std::string, std::string>>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::MapStringStringWrite(s, el)) { return false; }
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

void tl2::details::DictionaryFieldIntReset(::tl2::DictionaryField<int32_t>& item) {
	item.key.clear();
	item.value = 0;
}

bool tl2::details::DictionaryFieldIntWriteJSON(std::ostream& s, const ::tl2::DictionaryField<int32_t>& item) {
	s << "{";
	s << "\"key\":";
	s << "\"" << item.key << "\"";
	s << ",";
	s << "\"value\":";
	s << item.value;
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

void tl2::details::DoubleReset(double& item) {
	item = 0;
}

bool tl2::details::DoubleWriteJSON(std::ostream& s, const double& item) {
	s << item;
	return true;
}

bool tl2::details::DoubleRead(::basictl::tl_istream & s, double& item) {
	if (!s.double_read(item)) { return false; }
	return true;
}

bool tl2::details::DoubleWrite(::basictl::tl_ostream & s, const double& item) {
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

static const std::string_view EitherIntVectorService6FindWithBoundsResult_tbl_tl_name[]{"left", "right"};
static const uint32_t EitherIntVectorService6FindWithBoundsResult_tbl_tl_tag[]{0x0a29cd5d, 0xdf3ecb3b};

void tl2::details::EitherIntVectorService6FindWithBoundsResultReset(::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::EitherIntVectorService6FindWithBoundsResultWriteJSON(std::ostream & s, const ::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	s << "{";
	s << "\"type\":";
	s << EitherIntVectorService6FindWithBoundsResult_tbl_tl_tag[item.value.index()];
	switch (item.value.index()) {
	case 0:
		s << ",\"value\":";
		if (!::tl2::details::LeftIntVectorService6FindWithBoundsResultWriteJSON(s, std::get<0>(item.value))) { return false; }
		break;
	case 1:
		s << ",\"value\":";
		if (!::tl2::details::RightIntVectorService6FindWithBoundsResultWriteJSON(s, std::get<1>(item.value))) { return false; }
		break;
	}
	s << "}";
	return true;
}
bool tl2::details::EitherIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	uint32_t nat;
	s.nat_read(nat);
	switch (nat) {
	case 0x0a29cd5d:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		if (!::tl2::details::LeftIntVectorService6FindWithBoundsResultRead(s, std::get<0>(item.value))) { return false; }
		break;
	case 0xdf3ecb3b:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		if (!::tl2::details::RightIntVectorService6FindWithBoundsResultRead(s, std::get<1>(item.value))) { return false; }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::EitherIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	s.nat_write(EitherIntVectorService6FindWithBoundsResult_tbl_tl_tag[item.value.index()]);
	switch (item.value.index()) {
	case 0:
		if (!::tl2::details::LeftIntVectorService6FindWithBoundsResultWrite(s, std::get<0>(item.value))) { return false; }
		break;
	case 1:
		if (!::tl2::details::RightIntVectorService6FindWithBoundsResultWrite(s, std::get<1>(item.value))) { return false; }
		break;
	}
	return true;
}

static const std::string_view EitherService6ErrorVectorService6FindResultRow_tbl_tl_name[]{"left", "right"};
static const uint32_t EitherService6ErrorVectorService6FindResultRow_tbl_tl_tag[]{0x0a29cd5d, 0xdf3ecb3b};

void tl2::details::EitherService6ErrorVectorService6FindResultRowReset(::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::EitherService6ErrorVectorService6FindResultRowWriteJSON(std::ostream & s, const ::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	s << "{";
	s << "\"type\":";
	s << EitherService6ErrorVectorService6FindResultRow_tbl_tl_tag[item.value.index()];
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
bool tl2::details::EitherService6ErrorVectorService6FindResultRowReadBoxed(::basictl::tl_istream & s, ::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	uint32_t nat;
	s.nat_read(nat);
	switch (nat) {
	case 0x0a29cd5d:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		if (!::tl2::details::LeftService6ErrorVectorService6FindResultRowRead(s, std::get<0>(item.value))) { return false; }
		break;
	case 0xdf3ecb3b:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		if (!::tl2::details::RightService6ErrorVectorService6FindResultRowRead(s, std::get<1>(item.value))) { return false; }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::EitherService6ErrorVectorService6FindResultRowWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	s.nat_write(EitherService6ErrorVectorService6FindResultRow_tbl_tl_tag[item.value.index()]);
	switch (item.value.index()) {
	case 0:
		if (!::tl2::details::LeftService6ErrorVectorService6FindResultRowWrite(s, std::get<0>(item.value))) { return false; }
		break;
	case 1:
		if (!::tl2::details::RightService6ErrorVectorService6FindResultRowWrite(s, std::get<1>(item.value))) { return false; }
		break;
	}
	return true;
}

bool tl2::FieldConflict1::write_json(std::ostream& s)const {
	if (!::tl2::details::FieldConflict1WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict1::read(::basictl::tl_istream & s) {
	if (!::tl2::details::FieldConflict1Read(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict1::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::FieldConflict1Write(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict1::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::FieldConflict1ReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict1::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::FieldConflict1WriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::FieldConflict1Reset(::tl2::FieldConflict1& item) {
	item.x = 0;
	item.set_x = 0;
}

bool tl2::details::FieldConflict1WriteJSON(std::ostream& s, const ::tl2::FieldConflict1& item) {
	s << "{";
	s << "\"x\":";
	s << item.x;
	s << ",";
	s << "\"set_x\":";
	s << item.set_x;
	s << "}";
	return true;
}

bool tl2::details::FieldConflict1Read(::basictl::tl_istream & s, ::tl2::FieldConflict1& item) {
	if (!s.int_read(item.x)) { return false; }
	if (!s.int_read(item.set_x)) { return false; }
	return true;
}

bool tl2::details::FieldConflict1Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict1& item) {
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

bool tl2::FieldConflict2::read(::basictl::tl_istream & s) {
	if (!::tl2::details::FieldConflict2Read(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict2::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::FieldConflict2Write(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict2::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::FieldConflict2ReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict2::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::FieldConflict2WriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::FieldConflict2Reset(::tl2::FieldConflict2& item) {
	item.x = 0;
	item.setX = 0;
}

bool tl2::details::FieldConflict2WriteJSON(std::ostream& s, const ::tl2::FieldConflict2& item) {
	s << "{";
	s << "\"x\":";
	s << item.x;
	s << ",";
	s << "\"setX\":";
	s << item.setX;
	s << "}";
	return true;
}

bool tl2::details::FieldConflict2Read(::basictl::tl_istream & s, ::tl2::FieldConflict2& item) {
	if (!s.int_read(item.x)) { return false; }
	if (!s.int_read(item.setX)) { return false; }
	return true;
}

bool tl2::details::FieldConflict2Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict2& item) {
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

bool tl2::FieldConflict3::read(::basictl::tl_istream & s) {
	if (!::tl2::details::FieldConflict3Read(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict3::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::FieldConflict3Write(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict3::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::FieldConflict3ReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict3::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::FieldConflict3WriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::FieldConflict3Reset(::tl2::FieldConflict3& item) {
	item.x = 0;
	item.SetX = 0;
}

bool tl2::details::FieldConflict3WriteJSON(std::ostream& s, const ::tl2::FieldConflict3& item) {
	s << "{";
	s << "\"x\":";
	s << item.x;
	s << ",";
	s << "\"SetX\":";
	s << item.SetX;
	s << "}";
	return true;
}

bool tl2::details::FieldConflict3Read(::basictl::tl_istream & s, ::tl2::FieldConflict3& item) {
	if (!s.int_read(item.x)) { return false; }
	if (!s.int_read(item.SetX)) { return false; }
	return true;
}

bool tl2::details::FieldConflict3Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict3& item) {
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

bool tl2::FieldConflict4::read(::basictl::tl_istream & s) {
	if (!::tl2::details::FieldConflict4Read(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict4::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::FieldConflict4Write(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict4::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::FieldConflict4ReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::FieldConflict4::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::FieldConflict4WriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::FieldConflict4Reset(::tl2::FieldConflict4& item) {
	item.X = 0;
	item.SetX = 0;
}

bool tl2::details::FieldConflict4WriteJSON(std::ostream& s, const ::tl2::FieldConflict4& item) {
	s << "{";
	s << "\"X\":";
	s << item.X;
	s << ",";
	s << "\"SetX\":";
	s << item.SetX;
	s << "}";
	return true;
}

bool tl2::details::FieldConflict4Read(::basictl::tl_istream & s, ::tl2::FieldConflict4& item) {
	if (!s.int_read(item.X)) { return false; }
	if (!s.int_read(item.SetX)) { return false; }
	return true;
}

bool tl2::details::FieldConflict4Write(::basictl::tl_ostream & s, const ::tl2::FieldConflict4& item) {
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

void tl2::details::FloatReset(float& item) {
	item = 0;
}

bool tl2::details::FloatWriteJSON(std::ostream& s, const float& item) {
	s << item;
	return true;
}

bool tl2::details::FloatRead(::basictl::tl_istream & s, float& item) {
	if (!s.float_read(item)) { return false; }
	return true;
}

bool tl2::details::FloatWrite(::basictl::tl_ostream & s, const float& item) {
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

bool tl2::Get_arrays::read(::basictl::tl_istream & s) {
	if (!::tl2::details::GetArraysRead(s, *this)) { return false; }
	return true;
}

bool tl2::Get_arrays::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetArraysWrite(s, *this)) { return false; }
	return true;
}

bool tl2::Get_arrays::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::GetArraysReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::Get_arrays::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetArraysWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::GetArraysReset(::tl2::Get_arrays& item) {
	item.n = 0;
	item.a.clear();
	::tl2::details::BuiltinTuple5IntReset(item.b);
}

bool tl2::details::GetArraysWriteJSON(std::ostream& s, const ::tl2::Get_arrays& item) {
	s << "{";
	s << "\"n\":";
	s << item.n;
	s << ",";
	s << "\"a\":";
	if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.a, item.n)) { return false; }
	s << ",";
	s << "\"b\":";
	if (!::tl2::details::BuiltinTuple5IntWriteJSON(s, item.b)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::GetArraysRead(::basictl::tl_istream & s, ::tl2::Get_arrays& item) {
	if (!s.nat_read(item.n)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.a, item.n)) { return false; }
	if (!::tl2::details::BuiltinTuple5IntRead(s, item.b)) { return false; }
	return true;
}

bool tl2::details::GetArraysWrite(::basictl::tl_ostream & s, const ::tl2::Get_arrays& item) {
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
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTuple5IntRead(s, result)) { return false; }
	return true;
}
bool tl2::details::GetArraysWriteResult(::basictl::tl_ostream & s, tl2::Get_arrays& item, std::array<int32_t, 5>& result) {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTuple5IntWrite(s, result)) { return false; }
	return true;
}

bool tl2::Get_arrays::read_result(::basictl::tl_istream & s, std::array<int32_t, 5> & result) {
	return tl2::details::GetArraysReadResult(s, *this, result);
}
bool tl2::Get_arrays::write_result(::basictl::tl_ostream & s, std::array<int32_t, 5> & result) {
	return tl2::details::GetArraysWriteResult(s, *this, result);
}

bool tl2::GetDouble::write_json(std::ostream& s)const {
	if (!::tl2::details::GetDoubleWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetDouble::read(::basictl::tl_istream & s) {
	if (!::tl2::details::GetDoubleRead(s, *this)) { return false; }
	return true;
}

bool tl2::GetDouble::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetDoubleWrite(s, *this)) { return false; }
	return true;
}

bool tl2::GetDouble::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::GetDoubleReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::GetDouble::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetDoubleWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::GetDoubleReset(::tl2::GetDouble& item) {
	item.x = 0;
}

bool tl2::details::GetDoubleWriteJSON(std::ostream& s, const ::tl2::GetDouble& item) {
	s << "{";
	s << "\"x\":";
	s << item.x;
	s << "}";
	return true;
}

bool tl2::details::GetDoubleRead(::basictl::tl_istream & s, ::tl2::GetDouble& item) {
	if (!s.nat_read_exact_tag(0x2210c154)) { return false;}
	if (!s.double_read(item.x)) { return false; }
	return true;
}

bool tl2::details::GetDoubleWrite(::basictl::tl_ostream & s, const ::tl2::GetDouble& item) {
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
	if (!s.nat_read_exact_tag(0x2210c154)) { return false;}
	if (!s.double_read(result)) { return false; }
	return true;
}
bool tl2::details::GetDoubleWriteResult(::basictl::tl_ostream & s, tl2::GetDouble& item, double& result) {
	if (!s.nat_write(0x2210c154)) { return false; }
	if (!s.double_write(result)) { return false;}
	return true;
}

bool tl2::GetDouble::read_result(::basictl::tl_istream & s, double & result) {
	return tl2::details::GetDoubleReadResult(s, *this, result);
}
bool tl2::GetDouble::write_result(::basictl::tl_ostream & s, double & result) {
	return tl2::details::GetDoubleWriteResult(s, *this, result);
}

bool tl2::GetFloat::write_json(std::ostream& s)const {
	if (!::tl2::details::GetFloatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetFloat::read(::basictl::tl_istream & s) {
	if (!::tl2::details::GetFloatRead(s, *this)) { return false; }
	return true;
}

bool tl2::GetFloat::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetFloatWrite(s, *this)) { return false; }
	return true;
}

bool tl2::GetFloat::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::GetFloatReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::GetFloat::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetFloatWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::GetFloatReset(::tl2::GetFloat& item) {
	item.x = 0;
}

bool tl2::details::GetFloatWriteJSON(std::ostream& s, const ::tl2::GetFloat& item) {
	s << "{";
	s << "\"x\":";
	s << item.x;
	s << "}";
	return true;
}

bool tl2::details::GetFloatRead(::basictl::tl_istream & s, ::tl2::GetFloat& item) {
	if (!s.float_read(item.x)) { return false; }
	return true;
}

bool tl2::details::GetFloatWrite(::basictl::tl_ostream & s, const ::tl2::GetFloat& item) {
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
	if (!s.nat_read_exact_tag(0x824dab22)) { return false;}
	if (!s.float_read(result)) { return false; }
	return true;
}
bool tl2::details::GetFloatWriteResult(::basictl::tl_ostream & s, tl2::GetFloat& item, float& result) {
	if (!s.nat_write(0x824dab22)) { return false; }
	if (!s.float_write(result)) { return false;}
	return true;
}

bool tl2::GetFloat::read_result(::basictl::tl_istream & s, float & result) {
	return tl2::details::GetFloatReadResult(s, *this, result);
}
bool tl2::GetFloat::write_result(::basictl::tl_ostream & s, float & result) {
	return tl2::details::GetFloatWriteResult(s, *this, result);
}

bool tl2::GetMaybeIface::write_json(std::ostream& s)const {
	if (!::tl2::details::GetMaybeIfaceWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetMaybeIface::read(::basictl::tl_istream & s) {
	if (!::tl2::details::GetMaybeIfaceRead(s, *this)) { return false; }
	return true;
}

bool tl2::GetMaybeIface::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetMaybeIfaceWrite(s, *this)) { return false; }
	return true;
}

bool tl2::GetMaybeIface::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::GetMaybeIfaceReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::GetMaybeIface::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetMaybeIfaceWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::GetMaybeIfaceReset(::tl2::GetMaybeIface& item) {
	::tl2::details::Service1ValueReset(item.x);
}

bool tl2::details::GetMaybeIfaceWriteJSON(std::ostream& s, const ::tl2::GetMaybeIface& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::Service1ValueWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::GetMaybeIfaceRead(::basictl::tl_istream & s, ::tl2::GetMaybeIface& item) {
	if (!::tl2::details::Service1ValueReadBoxed(s, item.x)) { return false; }
	return true;
}

bool tl2::details::GetMaybeIfaceWrite(::basictl::tl_ostream & s, const ::tl2::GetMaybeIface& item) {
	if (!::tl2::details::Service1ValueWriteBoxed(s, item.x)) { return false; }
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

bool tl2::GetMaybeIface::read_result(::basictl::tl_istream & s, std::optional<::tl2::service1::Value> & result) {
	return tl2::details::GetMaybeIfaceReadResult(s, *this, result);
}
bool tl2::GetMaybeIface::write_result(::basictl::tl_ostream & s, std::optional<::tl2::service1::Value> & result) {
	return tl2::details::GetMaybeIfaceWriteResult(s, *this, result);
}

bool tl2::GetMyDictOfInt::write_json(std::ostream& s)const {
	if (!::tl2::details::GetMyDictOfIntWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyDictOfInt::read(::basictl::tl_istream & s) {
	if (!::tl2::details::GetMyDictOfIntRead(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyDictOfInt::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetMyDictOfIntWrite(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyDictOfInt::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::GetMyDictOfIntReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyDictOfInt::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetMyDictOfIntWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::GetMyDictOfIntReset(::tl2::GetMyDictOfInt& item) {
	::tl2::details::MyDictOfIntReset(item.x);
}

bool tl2::details::GetMyDictOfIntWriteJSON(std::ostream& s, const ::tl2::GetMyDictOfInt& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::MyDictOfIntWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::GetMyDictOfIntRead(::basictl::tl_istream & s, ::tl2::GetMyDictOfInt& item) {
	if (!::tl2::details::MyDictOfIntReadBoxed(s, item.x)) { return false; }
	return true;
}

bool tl2::details::GetMyDictOfIntWrite(::basictl::tl_ostream & s, const ::tl2::GetMyDictOfInt& item) {
	if (!::tl2::details::MyDictOfIntWriteBoxed(s, item.x)) { return false; }
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
	if (!::tl2::details::MyDictOfIntReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::GetMyDictOfIntWriteResult(::basictl::tl_ostream & s, tl2::GetMyDictOfInt& item, ::tl2::MyDictOfInt& result) {
	if (!::tl2::details::MyDictOfIntWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::GetMyDictOfInt::read_result(::basictl::tl_istream & s, ::tl2::MyDictOfInt & result) {
	return tl2::details::GetMyDictOfIntReadResult(s, *this, result);
}
bool tl2::GetMyDictOfInt::write_result(::basictl::tl_ostream & s, ::tl2::MyDictOfInt & result) {
	return tl2::details::GetMyDictOfIntWriteResult(s, *this, result);
}

bool tl2::GetMyDouble::write_json(std::ostream& s)const {
	if (!::tl2::details::GetMyDoubleWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyDouble::read(::basictl::tl_istream & s) {
	if (!::tl2::details::GetMyDoubleRead(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyDouble::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetMyDoubleWrite(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyDouble::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::GetMyDoubleReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyDouble::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetMyDoubleWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::GetMyDoubleReset(::tl2::GetMyDouble& item) {
	::tl2::details::MyDoubleReset(item.x);
}

bool tl2::details::GetMyDoubleWriteJSON(std::ostream& s, const ::tl2::GetMyDouble& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::MyDoubleWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::GetMyDoubleRead(::basictl::tl_istream & s, ::tl2::GetMyDouble& item) {
	if (!::tl2::details::MyDoubleRead(s, item.x)) { return false; }
	return true;
}

bool tl2::details::GetMyDoubleWrite(::basictl::tl_ostream & s, const ::tl2::GetMyDouble& item) {
	if (!::tl2::details::MyDoubleWrite(s, item.x)) { return false; }
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
	if (!::tl2::details::MyDoubleReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::GetMyDoubleWriteResult(::basictl::tl_ostream & s, tl2::GetMyDouble& item, ::tl2::MyDouble& result) {
	if (!::tl2::details::MyDoubleWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::GetMyDouble::read_result(::basictl::tl_istream & s, ::tl2::MyDouble & result) {
	return tl2::details::GetMyDoubleReadResult(s, *this, result);
}
bool tl2::GetMyDouble::write_result(::basictl::tl_ostream & s, ::tl2::MyDouble & result) {
	return tl2::details::GetMyDoubleWriteResult(s, *this, result);
}

bool tl2::GetMyValue::write_json(std::ostream& s)const {
	if (!::tl2::details::GetMyValueWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyValue::read(::basictl::tl_istream & s) {
	if (!::tl2::details::GetMyValueRead(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyValue::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetMyValueWrite(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyValue::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::GetMyValueReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::GetMyValue::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetMyValueWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::GetMyValueReset(::tl2::GetMyValue& item) {
	::tl2::details::MyValueReset(item.x);
}

bool tl2::details::GetMyValueWriteJSON(std::ostream& s, const ::tl2::GetMyValue& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::MyValueWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::GetMyValueRead(::basictl::tl_istream & s, ::tl2::GetMyValue& item) {
	if (!::tl2::details::MyValueReadBoxed(s, item.x)) { return false; }
	return true;
}

bool tl2::details::GetMyValueWrite(::basictl::tl_ostream & s, const ::tl2::GetMyValue& item) {
	if (!::tl2::details::MyValueWriteBoxed(s, item.x)) { return false; }
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
	if (!::tl2::details::MyValueReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::GetMyValueWriteResult(::basictl::tl_ostream & s, tl2::GetMyValue& item, ::tl2::MyValue& result) {
	if (!::tl2::details::MyValueWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::GetMyValue::read_result(::basictl::tl_istream & s, ::tl2::MyValue & result) {
	return tl2::details::GetMyValueReadResult(s, *this, result);
}
bool tl2::GetMyValue::write_result(::basictl::tl_ostream & s, ::tl2::MyValue & result) {
	return tl2::details::GetMyValueWriteResult(s, *this, result);
}

bool tl2::GetNonOptNat::write_json(std::ostream& s)const {
	if (!::tl2::details::GetNonOptNatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetNonOptNat::read(::basictl::tl_istream & s) {
	if (!::tl2::details::GetNonOptNatRead(s, *this)) { return false; }
	return true;
}

bool tl2::GetNonOptNat::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetNonOptNatWrite(s, *this)) { return false; }
	return true;
}

bool tl2::GetNonOptNat::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::GetNonOptNatReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::GetNonOptNat::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetNonOptNatWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::GetNonOptNatReset(::tl2::GetNonOptNat& item) {
	item.n = 0;
	item.xs.clear();
}

bool tl2::details::GetNonOptNatWriteJSON(std::ostream& s, const ::tl2::GetNonOptNat& item) {
	s << "{";
	s << "\"n\":";
	s << item.n;
	s << ",";
	s << "\"xs\":";
	if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.xs, item.n)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::GetNonOptNatRead(::basictl::tl_istream & s, ::tl2::GetNonOptNat& item) {
	if (!s.nat_read(item.n)) { return false; }
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTupleIntRead(s, item.xs, item.n)) { return false; }
	return true;
}

bool tl2::details::GetNonOptNatWrite(::basictl::tl_ostream & s, const ::tl2::GetNonOptNat& item) {
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
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTupleIntRead(s, result, item.n)) { return false; }
	return true;
}
bool tl2::details::GetNonOptNatWriteResult(::basictl::tl_ostream & s, tl2::GetNonOptNat& item, std::vector<int32_t>& result) {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleIntWrite(s, result, item.n)) { return false; }
	return true;
}

bool tl2::GetNonOptNat::read_result(::basictl::tl_istream & s, std::vector<int32_t> & result) {
	return tl2::details::GetNonOptNatReadResult(s, *this, result);
}
bool tl2::GetNonOptNat::write_result(::basictl::tl_ostream & s, std::vector<int32_t> & result) {
	return tl2::details::GetNonOptNatWriteResult(s, *this, result);
}

bool tl2::GetStats::write_json(std::ostream& s)const {
	if (!::tl2::details::GetStatsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::GetStats::read(::basictl::tl_istream & s) {
	if (!::tl2::details::GetStatsRead(s, *this)) { return false; }
	return true;
}

bool tl2::GetStats::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetStatsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::GetStats::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::GetStatsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::GetStats::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::GetStatsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::GetStatsReset(::tl2::GetStats& item) {
	::tl2::details::TasksQueueTypeStatsReset(item.x);
}

bool tl2::details::GetStatsWriteJSON(std::ostream& s, const ::tl2::GetStats& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::TasksQueueTypeStatsWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::GetStatsRead(::basictl::tl_istream & s, ::tl2::GetStats& item) {
	if (!::tl2::details::TasksQueueTypeStatsRead(s, item.x)) { return false; }
	return true;
}

bool tl2::details::GetStatsWrite(::basictl::tl_ostream & s, const ::tl2::GetStats& item) {
	if (!::tl2::details::TasksQueueTypeStatsWrite(s, item.x)) { return false; }
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
	if (!::tl2::details::TasksQueueTypeStatsReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::GetStatsWriteResult(::basictl::tl_ostream & s, tl2::GetStats& item, ::tl2::tasks::QueueTypeStats& result) {
	if (!::tl2::details::TasksQueueTypeStatsWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::GetStats::read_result(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeStats & result) {
	return tl2::details::GetStatsReadResult(s, *this, result);
}
bool tl2::GetStats::write_result(::basictl::tl_ostream & s, ::tl2::tasks::QueueTypeStats & result) {
	return tl2::details::GetStatsWriteResult(s, *this, result);
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

bool tl2::details::IntMaybeWriteJSON(std::ostream & s, const std::optional<int32_t>& item) {
	s << "{";
	if (item) {
		s << "\"ok\":true,";
		s << "\"value\":";
		s << *item;
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

bool tl2::Integer::read(::basictl::tl_istream & s) {
	if (!::tl2::details::IntegerRead(s, *this)) { return false; }
	return true;
}

bool tl2::Integer::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::IntegerWrite(s, *this)) { return false; }
	return true;
}

bool tl2::Integer::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::IntegerReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::Integer::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::IntegerWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::IntegerReset(::tl2::Integer& item) {
	item.value = 0;
}

bool tl2::details::IntegerWriteJSON(std::ostream& s, const ::tl2::Integer& item) {
	s << "{";
	s << "\"value\":";
	s << item.value;
	s << "}";
	return true;
}

bool tl2::details::IntegerRead(::basictl::tl_istream & s, ::tl2::Integer& item) {
	if (!s.int_read(item.value)) { return false; }
	return true;
}

bool tl2::details::IntegerWrite(::basictl::tl_ostream & s, const ::tl2::Integer& item) {
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

bool tl2::Issue3498::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Issue3498Read(s, *this)) { return false; }
	return true;
}

bool tl2::Issue3498::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Issue3498Write(s, *this)) { return false; }
	return true;
}

bool tl2::Issue3498::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Issue3498ReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::Issue3498::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Issue3498WriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Issue3498Reset(::tl2::Issue3498& item) {
	item.x.clear();
}

bool tl2::details::Issue3498WriteJSON(std::ostream& s, const ::tl2::Issue3498& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::Issue3498Read(::basictl::tl_istream & s, ::tl2::Issue3498& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false;}
	if (!::tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowRead(s, item.x)) { return false; }
	return true;
}

bool tl2::details::Issue3498Write(::basictl::tl_ostream & s, const ::tl2::Issue3498& item) {
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

void tl2::details::LeftIntVectorService6FindWithBoundsResultReset(::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	item.value = 0;
}

bool tl2::details::LeftIntVectorService6FindWithBoundsResultWriteJSON(std::ostream& s, const ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	s << "{";
	s << "\"value\":";
	s << item.value;
	s << "}";
	return true;
}

bool tl2::details::LeftIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	if (!s.int_read(item.value)) { return false; }
	return true;
}

bool tl2::details::LeftIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	if (!s.int_write(item.value)) { return false;}
	return true;
}

bool tl2::details::LeftIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	if (!s.nat_read_exact_tag(0x0a29cd5d)) { return false; }
	return tl2::details::LeftIntVectorService6FindWithBoundsResultRead(s, item);
}

bool tl2::details::LeftIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Left<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	if (!s.nat_write(0x0a29cd5d)) { return false; }
	return tl2::details::LeftIntVectorService6FindWithBoundsResultWrite(s, item);
}

void tl2::details::LeftService6ErrorVectorService6FindResultRowReset(::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	::tl2::details::Service6ErrorReset(item.value);
}

bool tl2::details::LeftService6ErrorVectorService6FindResultRowWriteJSON(std::ostream& s, const ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::Service6ErrorWriteJSON(s, item.value)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::LeftService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	if (!::tl2::details::Service6ErrorRead(s, item.value)) { return false; }
	return true;
}

bool tl2::details::LeftService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const ::tl2::Left<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	if (!::tl2::details::Service6ErrorWrite(s, item.value)) { return false; }
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

void tl2::details::MapStringStringReset(::tl2::Map<std::string, std::string>& item) {
	item.key.clear();
	item.value.clear();
}

bool tl2::details::MapStringStringWriteJSON(std::ostream& s, const ::tl2::Map<std::string, std::string>& item) {
	s << "{";
	s << "\"key\":";
	s << "\"" << item.key << "\"";
	s << ",";
	s << "\"value\":";
	s << "\"" << item.value << "\"";
	s << "}";
	return true;
}

bool tl2::details::MapStringStringRead(::basictl::tl_istream & s, ::tl2::Map<std::string, std::string>& item) {
	if (!s.string_read(item.key)) { return false; }
	if (!s.string_read(item.value)) { return false; }
	return true;
}

bool tl2::details::MapStringStringWrite(::basictl::tl_ostream & s, const ::tl2::Map<std::string, std::string>& item) {
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

void tl2::details::MyAnonMcValueReset(::tl2::MyAnonMcValue& item) {
	::tl2::details::Service1ValueReset(item);
}

bool tl2::details::MyAnonMcValueWriteJSON(std::ostream& s, const ::tl2::MyAnonMcValue& item) {
	if (!::tl2::details::Service1ValueWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::MyAnonMcValueRead(::basictl::tl_istream & s, ::tl2::MyAnonMcValue& item) {
	if (!::tl2::details::Service1ValueReadBoxed(s, item)) { return false; }
	return true;
}

bool tl2::details::MyAnonMcValueWrite(::basictl::tl_ostream & s, const ::tl2::MyAnonMcValue& item) {
	if (!::tl2::details::Service1ValueWriteBoxed(s, item)) { return false; }
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

bool tl2::MyBoxedArray::read(::basictl::tl_istream & s) {
	if (!::tl2::details::MyBoxedArrayRead(s, *this)) { return false; }
	return true;
}

bool tl2::MyBoxedArray::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyBoxedArrayWrite(s, *this)) { return false; }
	return true;
}

bool tl2::MyBoxedArray::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::MyBoxedArrayReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::MyBoxedArray::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyBoxedArrayWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::MyBoxedArrayReset(::tl2::MyBoxedArray& item) {
	::tl2::details::BuiltinTuple2IntBoxedReset(item.data);
}

bool tl2::details::MyBoxedArrayWriteJSON(std::ostream& s, const ::tl2::MyBoxedArray& item) {
	s << "{";
	s << "\"data\":";
	if (!::tl2::details::BuiltinTuple2IntBoxedWriteJSON(s, item.data)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::MyBoxedArrayRead(::basictl::tl_istream & s, ::tl2::MyBoxedArray& item) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTuple2IntBoxedRead(s, item.data)) { return false; }
	return true;
}

bool tl2::details::MyBoxedArrayWrite(::basictl::tl_ostream & s, const ::tl2::MyBoxedArray& item) {
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

bool tl2::MyBoxedTupleSlice::read(::basictl::tl_istream & s) {
	if (!::tl2::details::MyBoxedTupleSliceRead(s, *this)) { return false; }
	return true;
}

bool tl2::MyBoxedTupleSlice::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyBoxedTupleSliceWrite(s, *this)) { return false; }
	return true;
}

bool tl2::MyBoxedTupleSlice::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::MyBoxedTupleSliceReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::MyBoxedTupleSlice::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyBoxedTupleSliceWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::MyBoxedTupleSliceReset(::tl2::MyBoxedTupleSlice& item) {
	item.n = 0;
	item.data.clear();
}

bool tl2::details::MyBoxedTupleSliceWriteJSON(std::ostream& s, const ::tl2::MyBoxedTupleSlice& item) {
	s << "{";
	s << "\"n\":";
	s << item.n;
	s << ",";
	s << "\"data\":";
	if (!::tl2::details::BuiltinTupleIntBoxedWriteJSON(s, item.data, item.n)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::MyBoxedTupleSliceRead(::basictl::tl_istream & s, ::tl2::MyBoxedTupleSlice& item) {
	if (!s.nat_read(item.n)) { return false; }
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTupleIntBoxedRead(s, item.data, item.n)) { return false; }
	return true;
}

bool tl2::details::MyBoxedTupleSliceWrite(::basictl::tl_ostream & s, const ::tl2::MyBoxedTupleSlice& item) {
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

bool tl2::MyBoxedVectorSlice::read(::basictl::tl_istream & s) {
	if (!::tl2::details::MyBoxedVectorSliceRead(s, *this)) { return false; }
	return true;
}

bool tl2::MyBoxedVectorSlice::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyBoxedVectorSliceWrite(s, *this)) { return false; }
	return true;
}

bool tl2::MyBoxedVectorSlice::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::MyBoxedVectorSliceReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::MyBoxedVectorSlice::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyBoxedVectorSliceWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::MyBoxedVectorSliceReset(::tl2::MyBoxedVectorSlice& item) {
	item.data.clear();
}

bool tl2::details::MyBoxedVectorSliceWriteJSON(std::ostream& s, const ::tl2::MyBoxedVectorSlice& item) {
	s << "{";
	s << "\"data\":";
	if (!::tl2::details::BuiltinVectorIntBoxedWriteJSON(s, item.data)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::MyBoxedVectorSliceRead(::basictl::tl_istream & s, ::tl2::MyBoxedVectorSlice& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false;}
	if (!::tl2::details::BuiltinVectorIntBoxedRead(s, item.data)) { return false; }
	return true;
}

bool tl2::details::MyBoxedVectorSliceWrite(::basictl::tl_ostream & s, const ::tl2::MyBoxedVectorSlice& item) {
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

void tl2::details::MyDictOfIntReset(::tl2::MyDictOfInt& item) {
	::tl2::details::DictionaryIntReset(item);
}

bool tl2::details::MyDictOfIntWriteJSON(std::ostream& s, const ::tl2::MyDictOfInt& item) {
	if (!::tl2::details::DictionaryIntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::MyDictOfIntRead(::basictl::tl_istream & s, ::tl2::MyDictOfInt& item) {
	if (!::tl2::details::DictionaryIntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::MyDictOfIntWrite(::basictl::tl_ostream & s, const ::tl2::MyDictOfInt& item) {
	if (!::tl2::details::DictionaryIntWrite(s, item)) { return false; }
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

void tl2::details::MyDoubleReset(::tl2::MyDouble& item) {
	item = 0;
}

bool tl2::details::MyDoubleWriteJSON(std::ostream& s, const ::tl2::MyDouble& item) {
	s << item;
	return true;
}

bool tl2::details::MyDoubleRead(::basictl::tl_istream & s, ::tl2::MyDouble& item) {
	if (!s.nat_read_exact_tag(0x2210c154)) { return false;}
	if (!s.double_read(item)) { return false; }
	return true;
}

bool tl2::details::MyDoubleWrite(::basictl::tl_ostream & s, const ::tl2::MyDouble& item) {
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

bool tl2::MyInt::read(::basictl::tl_istream & s) {
	if (!::tl2::details::MyIntRead(s, *this)) { return false; }
	return true;
}

bool tl2::MyInt::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyIntWrite(s, *this)) { return false; }
	return true;
}

bool tl2::MyInt::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::MyIntReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::MyInt::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyIntWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::MyIntReset(::tl2::MyInt& item) {
	item.val1 = 0;
}

bool tl2::details::MyIntWriteJSON(std::ostream& s, const ::tl2::MyInt& item) {
	s << "{";
	s << "\"val1\":";
	s << item.val1;
	s << "}";
	return true;
}

bool tl2::details::MyIntRead(::basictl::tl_istream & s, ::tl2::MyInt& item) {
	if (!s.nat_read_exact_tag(0xa8509bda)) { return false;}
	if (!s.int_read(item.val1)) { return false; }
	return true;
}

bool tl2::details::MyIntWrite(::basictl::tl_ostream & s, const ::tl2::MyInt& item) {
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

void tl2::details::MyMaybe1Reset(::tl2::MyMaybe1& item) {
	item.reset();
}

bool tl2::details::MyMaybe1WriteJSON(std::ostream& s, const ::tl2::MyMaybe1& item) {
	if (!::tl2::details::MyTuple10MaybeWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::MyMaybe1Read(::basictl::tl_istream & s, ::tl2::MyMaybe1& item) {
	if (!::tl2::details::MyTuple10MaybeReadBoxed(s, item)) { return false; }
	return true;
}

bool tl2::details::MyMaybe1Write(::basictl::tl_ostream & s, const ::tl2::MyMaybe1& item) {
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

void tl2::details::MyMaybe2Reset(::tl2::MyMaybe2& item) {
	item.reset();
}

bool tl2::details::MyMaybe2WriteJSON(std::ostream& s, const ::tl2::MyMaybe2& item) {
	if (!::tl2::details::MyTuple10MaybeWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::MyMaybe2Read(::basictl::tl_istream & s, ::tl2::MyMaybe2& item) {
	if (!::tl2::details::MyTuple10MaybeReadBoxed(s, item)) { return false; }
	return true;
}

bool tl2::details::MyMaybe2Write(::basictl::tl_ostream & s, const ::tl2::MyMaybe2& item) {
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

bool tl2::MyMcValue::read(::basictl::tl_istream & s) {
	if (!::tl2::details::MyMcValueRead(s, *this)) { return false; }
	return true;
}

bool tl2::MyMcValue::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyMcValueWrite(s, *this)) { return false; }
	return true;
}

bool tl2::MyMcValue::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::MyMcValueReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::MyMcValue::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyMcValueWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::MyMcValueReset(::tl2::MyMcValue& item) {
	::tl2::details::Service1ValueReset(item.x);
}

bool tl2::details::MyMcValueWriteJSON(std::ostream& s, const ::tl2::MyMcValue& item) {
	s << "{";
	s << "\"x\":";
	if (!::tl2::details::Service1ValueWriteJSON(s, item.x)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::MyMcValueRead(::basictl::tl_istream & s, ::tl2::MyMcValue& item) {
	if (!::tl2::details::Service1ValueReadBoxed(s, item.x)) { return false; }
	return true;
}

bool tl2::details::MyMcValueWrite(::basictl::tl_ostream & s, const ::tl2::MyMcValue& item) {
	if (!::tl2::details::Service1ValueWriteBoxed(s, item.x)) { return false; }
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

bool tl2::MyMcValueTuple::read(::basictl::tl_istream & s) {
	if (!::tl2::details::MyMcValueTupleRead(s, *this)) { return false; }
	return true;
}

bool tl2::MyMcValueTuple::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyMcValueTupleWrite(s, *this)) { return false; }
	return true;
}

bool tl2::MyMcValueTuple::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::MyMcValueTupleReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::MyMcValueTuple::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyMcValueTupleWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::MyMcValueTupleReset(::tl2::MyMcValueTuple& item) {
	::tl2::details::BuiltinTuple3Service1ValueReset(item.xs);
}

bool tl2::details::MyMcValueTupleWriteJSON(std::ostream& s, const ::tl2::MyMcValueTuple& item) {
	s << "{";
	s << "\"xs\":";
	if (!::tl2::details::BuiltinTuple3Service1ValueWriteJSON(s, item.xs)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::MyMcValueTupleRead(::basictl::tl_istream & s, ::tl2::MyMcValueTuple& item) {
	if (!::tl2::details::BuiltinTuple3Service1ValueRead(s, item.xs)) { return false; }
	return true;
}

bool tl2::details::MyMcValueTupleWrite(::basictl::tl_ostream & s, const ::tl2::MyMcValueTuple& item) {
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

bool tl2::MyMcValueVector::read(::basictl::tl_istream & s) {
	if (!::tl2::details::MyMcValueVectorRead(s, *this)) { return false; }
	return true;
}

bool tl2::MyMcValueVector::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyMcValueVectorWrite(s, *this)) { return false; }
	return true;
}

bool tl2::MyMcValueVector::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::MyMcValueVectorReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::MyMcValueVector::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyMcValueVectorWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::MyMcValueVectorReset(::tl2::MyMcValueVector& item) {
	item.xs.clear();
}

bool tl2::details::MyMcValueVectorWriteJSON(std::ostream& s, const ::tl2::MyMcValueVector& item) {
	s << "{";
	s << "\"xs\":";
	if (!::tl2::details::BuiltinVectorService1ValueWriteJSON(s, item.xs)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::MyMcValueVectorRead(::basictl::tl_istream & s, ::tl2::MyMcValueVector& item) {
	if (!::tl2::details::BuiltinVectorService1ValueRead(s, item.xs)) { return false; }
	return true;
}

bool tl2::details::MyMcValueVectorWrite(::basictl::tl_ostream & s, const ::tl2::MyMcValueVector& item) {
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

bool tl2::MyString::read(::basictl::tl_istream & s) {
	if (!::tl2::details::MyStringRead(s, *this)) { return false; }
	return true;
}

bool tl2::MyString::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyStringWrite(s, *this)) { return false; }
	return true;
}

bool tl2::MyString::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::MyStringReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::MyString::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyStringWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::MyStringReset(::tl2::MyString& item) {
	item.val2.clear();
}

bool tl2::details::MyStringWriteJSON(std::ostream& s, const ::tl2::MyString& item) {
	s << "{";
	s << "\"val2\":";
	s << "\"" << item.val2 << "\"";
	s << "}";
	return true;
}

bool tl2::details::MyStringRead(::basictl::tl_istream & s, ::tl2::MyString& item) {
	if (!s.nat_read_exact_tag(0xb5286e24)) { return false;}
	if (!s.string_read(item.val2)) { return false; }
	return true;
}

bool tl2::details::MyStringWrite(::basictl::tl_ostream & s, const ::tl2::MyString& item) {
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

void tl2::details::MyTuple10Reset(::tl2::MyTuplen<10>& item) {
	::tl2::details::BuiltinTuple10IntBoxedReset(item);
}

bool tl2::details::MyTuple10WriteJSON(std::ostream& s, const ::tl2::MyTuplen<10>& item) {
	if (!::tl2::details::BuiltinTuple10IntBoxedWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::MyTuple10Read(::basictl::tl_istream & s, ::tl2::MyTuplen<10>& item) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTuple10IntBoxedRead(s, item)) { return false; }
	return true;
}

bool tl2::details::MyTuple10Write(::basictl::tl_ostream & s, const ::tl2::MyTuplen<10>& item) {
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
		s << "\"ok\":true,";
		s << "\"value\":";
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
		if (!::tl2::details::MyTuple10Read(s, *item)) { return false; }
		return true;
	}
	item.reset();
	return true;
}

bool tl2::details::MyTuple10MaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::MyTuplen<10>>& item) {
	if (!s.nat_write(item ? 0x3f9c8ef8 : 0x27930a7b)) { return false; }
	if (item) {
		if (!::tl2::details::MyTuple10Write(s, *item)) { return false; }
	}
	return true;
}

bool tl2::MyTwoDicts::write_json(std::ostream& s)const {
	if (!::tl2::details::MyTwoDictsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::MyTwoDicts::read(::basictl::tl_istream & s) {
	if (!::tl2::details::MyTwoDictsRead(s, *this)) { return false; }
	return true;
}

bool tl2::MyTwoDicts::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyTwoDictsWrite(s, *this)) { return false; }
	return true;
}

bool tl2::MyTwoDicts::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::MyTwoDictsReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::MyTwoDicts::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyTwoDictsWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::MyTwoDictsReset(::tl2::MyTwoDicts& item) {
	::tl2::details::DictionaryIntReset(item.a);
	::tl2::details::DictionaryIntReset(item.b);
}

bool tl2::details::MyTwoDictsWriteJSON(std::ostream& s, const ::tl2::MyTwoDicts& item) {
	s << "{";
	s << "\"a\":";
	if (!::tl2::details::DictionaryIntWriteJSON(s, item.a)) { return false; }
	s << ",";
	s << "\"b\":";
	if (!::tl2::details::DictionaryIntWriteJSON(s, item.b)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::MyTwoDictsRead(::basictl::tl_istream & s, ::tl2::MyTwoDicts& item) {
	if (!::tl2::details::DictionaryIntRead(s, item.a)) { return false; }
	if (!::tl2::details::DictionaryIntRead(s, item.b)) { return false; }
	return true;
}

bool tl2::details::MyTwoDictsWrite(::basictl::tl_ostream & s, const ::tl2::MyTwoDicts& item) {
	if (!::tl2::details::DictionaryIntWrite(s, item.a)) { return false; }
	if (!::tl2::details::DictionaryIntWrite(s, item.b)) { return false; }
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
bool tl2::MyValue::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::MyValueReadBoxed(s, *this)) { return false; }
	return true;
}
bool tl2::MyValue::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::MyValueWriteBoxed(s, *this)) { return false; }
	return true;
}
std::string_view tl2::MyValue::tl_name() const {
	return MyValue_tbl_tl_name[value.index()];
}
uint32_t tl2::MyValue::tl_tag() const {
	return MyValue_tbl_tl_tag[value.index()];
}


void tl2::details::MyValueReset(::tl2::MyValue& item) {
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::MyValueWriteJSON(std::ostream & s, const ::tl2::MyValue& item) {
	s << "{";
	s << "\"type\":";
	s << MyValue_tbl_tl_tag[item.value.index()];
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
bool tl2::details::MyValueReadBoxed(::basictl::tl_istream & s, ::tl2::MyValue& item) {
	uint32_t nat;
	s.nat_read(nat);
	switch (nat) {
	case 0xc12375b7:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		if (!::tl2::details::MyIntRead(s, std::get<0>(item.value))) { return false; }
		break;
	case 0xc8bfa969:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		if (!::tl2::details::MyStringRead(s, std::get<1>(item.value))) { return false; }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::MyValueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::MyValue& item) {
	s.nat_write(MyValue_tbl_tl_tag[item.value.index()]);
	switch (item.value.index()) {
	case 0:
		if (!::tl2::details::MyIntWrite(s, std::get<0>(item.value))) { return false; }
		break;
	case 1:
		if (!::tl2::details::MyStringWrite(s, std::get<1>(item.value))) { return false; }
		break;
	}
	return true;
}

bool tl2::NonOptNat::write_json(std::ostream& s)const {
	if (!::tl2::details::NonOptNatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::NonOptNat::read(::basictl::tl_istream & s) {
	if (!::tl2::details::NonOptNatRead(s, *this)) { return false; }
	return true;
}

bool tl2::NonOptNat::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::NonOptNatWrite(s, *this)) { return false; }
	return true;
}

bool tl2::NonOptNat::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::NonOptNatReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::NonOptNat::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::NonOptNatWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::NonOptNatReset(::tl2::NonOptNat& item) {
	item.n = 0;
	item.xs.clear();
}

bool tl2::details::NonOptNatWriteJSON(std::ostream& s, const ::tl2::NonOptNat& item) {
	s << "{";
	s << "\"n\":";
	s << item.n;
	s << ",";
	s << "\"xs\":";
	if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.xs, item.n)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::NonOptNatRead(::basictl::tl_istream & s, ::tl2::NonOptNat& item) {
	if (!s.nat_read(item.n)) { return false; }
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTupleIntRead(s, item.xs, item.n)) { return false; }
	return true;
}

bool tl2::details::NonOptNatWrite(::basictl::tl_ostream & s, const ::tl2::NonOptNat& item) {
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

void tl2::details::RightIntVectorService6FindWithBoundsResultReset(::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	item.value.clear();
}

bool tl2::details::RightIntVectorService6FindWithBoundsResultWriteJSON(std::ostream& s, const ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::BuiltinVectorService6FindWithBoundsResultWriteJSON(s, item.value)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::RightIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	if (!::tl2::details::BuiltinVectorService6FindWithBoundsResultRead(s, item.value)) { return false; }
	return true;
}

bool tl2::details::RightIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	if (!::tl2::details::BuiltinVectorService6FindWithBoundsResultWrite(s, item.value)) { return false; }
	return true;
}

bool tl2::details::RightIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	if (!s.nat_read_exact_tag(0xdf3ecb3b)) { return false; }
	return tl2::details::RightIntVectorService6FindWithBoundsResultRead(s, item);
}

bool tl2::details::RightIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Right<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>& item) {
	if (!s.nat_write(0xdf3ecb3b)) { return false; }
	return tl2::details::RightIntVectorService6FindWithBoundsResultWrite(s, item);
}

void tl2::details::RightService6ErrorVectorService6FindResultRowReset(::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	item.value.clear();
}

bool tl2::details::RightService6ErrorVectorService6FindResultRowWriteJSON(std::ostream& s, const ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::BuiltinVectorService6FindResultRowWriteJSON(s, item.value)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::RightService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
	if (!::tl2::details::BuiltinVectorService6FindResultRowRead(s, item.value)) { return false; }
	return true;
}

bool tl2::details::RightService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const ::tl2::Right<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>& item) {
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

bool tl2::RpcInvokeReqExtra::read(::basictl::tl_istream & s) {
	if (!::tl2::details::RpcInvokeReqExtraRead(s, *this)) { return false; }
	return true;
}

bool tl2::RpcInvokeReqExtra::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::RpcInvokeReqExtraWrite(s, *this)) { return false; }
	return true;
}

bool tl2::RpcInvokeReqExtra::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::RpcInvokeReqExtraReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::RpcInvokeReqExtra::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::RpcInvokeReqExtraWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::RpcInvokeReqExtraReset(::tl2::RpcInvokeReqExtra& item) {
	item.fields_mask = 0;
	::tl2::details::TrueReset(item.query);
	::tl2::details::TrueReset(item.sort);
	::tl2::details::TrueReset(item.sort_reverse);
	item.wait_binlog_pos = 0;
	item.string_forward_keys.clear();
}

bool tl2::details::RpcInvokeReqExtraWriteJSON(std::ostream& s, const ::tl2::RpcInvokeReqExtra& item) {
	s << "{";
	s << "\"fields_mask\":";
	s << item.fields_mask;
	if ((item.fields_mask & (1<<0)) != 0) {
		s << ",";
		s << "\"query\":";
		if (!::tl2::details::TrueWriteJSON(s, item.query)) { return false; }
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		s << ",";
		s << "\"sort\":";
		if (!::tl2::details::TrueWriteJSON(s, item.sort)) { return false; }
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		s << ",";
		s << "\"sort_reverse\":";
		if (!::tl2::details::TrueWriteJSON(s, item.sort_reverse)) { return false; }
	}
	if ((item.fields_mask & (1<<16)) != 0) {
		s << ",";
		s << "\"wait_binlog_pos\":";
		s << item.wait_binlog_pos;
	}
	if ((item.fields_mask & (1<<18)) != 0) {
		s << ",";
		s << "\"string_forward_keys\":";
		if (!::tl2::details::BuiltinVectorStringWriteJSON(s, item.string_forward_keys)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::RpcInvokeReqExtraRead(::basictl::tl_istream & s, ::tl2::RpcInvokeReqExtra& item) {
	if (!s.nat_read(item.fields_mask)) { return false; }
	if ((item.fields_mask & (1<<0)) != 0) {
		if (!::tl2::details::TrueRead(s, item.query)) { return false; }
	} else {
			::tl2::details::TrueReset(item.query);
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (!::tl2::details::TrueRead(s, item.sort)) { return false; }
	} else {
			::tl2::details::TrueReset(item.sort);
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (!::tl2::details::TrueRead(s, item.sort_reverse)) { return false; }
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

bool tl2::details::RpcInvokeReqExtraWrite(::basictl::tl_ostream & s, const ::tl2::RpcInvokeReqExtra& item) {
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.query)) { return false; }
	}
	if ((item.fields_mask & (1<<1)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.sort)) { return false; }
	}
	if ((item.fields_mask & (1<<2)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.sort_reverse)) { return false; }
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

bool tl2::StatOne::read(::basictl::tl_istream & s) {
	if (!::tl2::details::StatOneRead(s, *this)) { return false; }
	return true;
}

bool tl2::StatOne::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::StatOneWrite(s, *this)) { return false; }
	return true;
}

bool tl2::StatOne::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::StatOneReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::StatOne::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::StatOneWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::StatOneReset(::tl2::StatOne& item) {
	item.key.clear();
	item.value.clear();
}

bool tl2::details::StatOneWriteJSON(std::ostream& s, const ::tl2::StatOne& item) {
	s << "{";
	s << "\"key\":";
	s << "\"" << item.key << "\"";
	s << ",";
	s << "\"value\":";
	s << "\"" << item.value << "\"";
	s << "}";
	return true;
}

bool tl2::details::StatOneRead(::basictl::tl_istream & s, ::tl2::StatOne& item) {
	if (!s.string_read(item.key)) { return false; }
	if (!s.string_read(item.value)) { return false; }
	return true;
}

bool tl2::details::StatOneWrite(::basictl::tl_ostream & s, const ::tl2::StatOne& item) {
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

bool tl2::details::TrueWriteJSON(std::ostream& s, const ::tl2::True& item) {
	s << "{";
	s << "}";
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

void tl2::details::TupleIntReset(std::vector<int32_t>& item) {
	item.clear();
}

bool tl2::details::TupleIntWriteJSON(std::ostream& s, const std::vector<int32_t>& item, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item, nat_n)) { return false; }
	return true;
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

void tl2::details::TupleInt3Reset(std::array<int32_t, 3>& item) {
	::tl2::details::BuiltinTuple3IntReset(item);
}

bool tl2::details::TupleInt3WriteJSON(std::ostream& s, const std::array<int32_t, 3>& item) {
	if (!::tl2::details::BuiltinTuple3IntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt3Read(::basictl::tl_istream & s, std::array<int32_t, 3>& item) {
	if (!::tl2::details::BuiltinTuple3IntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt3Write(::basictl::tl_ostream & s, const std::array<int32_t, 3>& item) {
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

void tl2::details::TupleInt5Reset(std::array<int32_t, 5>& item) {
	::tl2::details::BuiltinTuple5IntReset(item);
}

bool tl2::details::TupleInt5WriteJSON(std::ostream& s, const std::array<int32_t, 5>& item) {
	if (!::tl2::details::BuiltinTuple5IntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt5Read(::basictl::tl_istream & s, std::array<int32_t, 5>& item) {
	if (!::tl2::details::BuiltinTuple5IntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt5Write(::basictl::tl_ostream & s, const std::array<int32_t, 5>& item) {
	if (!::tl2::details::BuiltinTuple5IntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt5ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 5>& item) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleInt5Read(s, item);
}

bool tl2::details::TupleInt5WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 5>& item) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleInt5Write(s, item);
}

void tl2::details::TupleIntBoxedReset(std::vector<int32_t>& item) {
	item.clear();
}

bool tl2::details::TupleIntBoxedWriteJSON(std::ostream& s, const std::vector<int32_t>& item, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleIntBoxedWriteJSON(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleIntBoxedRead(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n) {
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

void tl2::details::TupleIntBoxed10Reset(std::array<int32_t, 10>& item) {
	::tl2::details::BuiltinTuple10IntBoxedReset(item);
}

bool tl2::details::TupleIntBoxed10WriteJSON(std::ostream& s, const std::array<int32_t, 10>& item) {
	if (!::tl2::details::BuiltinTuple10IntBoxedWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxed10Read(::basictl::tl_istream & s, std::array<int32_t, 10>& item) {
	if (!::tl2::details::BuiltinTuple10IntBoxedRead(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxed10Write(::basictl::tl_ostream & s, const std::array<int32_t, 10>& item) {
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

void tl2::details::TupleIntBoxed2Reset(std::array<int32_t, 2>& item) {
	::tl2::details::BuiltinTuple2IntBoxedReset(item);
}

bool tl2::details::TupleIntBoxed2WriteJSON(std::ostream& s, const std::array<int32_t, 2>& item) {
	if (!::tl2::details::BuiltinTuple2IntBoxedWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxed2Read(::basictl::tl_istream & s, std::array<int32_t, 2>& item) {
	if (!::tl2::details::BuiltinTuple2IntBoxedRead(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleIntBoxed2Write(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item) {
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

void tl2::details::TupleService1Value3Reset(std::array<::tl2::service1::Value, 3>& item) {
	::tl2::details::BuiltinTuple3Service1ValueReset(item);
}

bool tl2::details::TupleService1Value3WriteJSON(std::ostream& s, const std::array<::tl2::service1::Value, 3>& item) {
	if (!::tl2::details::BuiltinTuple3Service1ValueWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleService1Value3Read(::basictl::tl_istream & s, std::array<::tl2::service1::Value, 3>& item) {
	if (!::tl2::details::BuiltinTuple3Service1ValueRead(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleService1Value3Write(::basictl::tl_ostream & s, const std::array<::tl2::service1::Value, 3>& item) {
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

void tl2::details::TupleService2CounterSetReset(std::vector<::tl2::service2::CounterSet>& item) {
	item.clear();
}

bool tl2::details::TupleService2CounterSetWriteJSON(std::ostream& s, const std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleService2CounterSetWriteJSON(s, item, nat_n, nat_tintCountersNum, nat_tfloatCountersNum)) { return false; }
	return true;
}

bool tl2::details::TupleService2CounterSetRead(::basictl::tl_istream & s, std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleService2CounterSetRead(s, item, nat_n, nat_tintCountersNum, nat_tfloatCountersNum)) { return false; }
	return true;
}

bool tl2::details::TupleService2CounterSetWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleService2CounterSetWrite(s, item, nat_n, nat_tintCountersNum, nat_tfloatCountersNum)) { return false; }
	return true;
}

bool tl2::details::TupleService2CounterSetReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleService2CounterSetRead(s, item, nat_tintCountersNum, nat_tfloatCountersNum, nat_n);
}

bool tl2::details::TupleService2CounterSetWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleService2CounterSetWrite(s, item, nat_tintCountersNum, nat_tfloatCountersNum, nat_n);
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

void tl2::details::VectorEitherIntVectorService6FindWithBoundsResultReset(std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	item.clear();
}

bool tl2::details::VectorEitherIntVectorService6FindWithBoundsResultWriteJSON(std::ostream& s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	if (!::tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorEitherIntVectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	if (!::tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorEitherIntVectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	if (!::tl2::details::BuiltinVectorEitherIntVectorService6FindWithBoundsResultWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorEitherIntVectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorEitherIntVectorService6FindWithBoundsResultRead(s, item);
}

bool tl2::details::VectorEitherIntVectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::Either<int32_t, std::vector<::tl2::service6::FindWithBoundsResult>>>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorEitherIntVectorService6FindWithBoundsResultWrite(s, item);
}

void tl2::details::VectorEitherService6ErrorVectorService6FindResultRowReset(std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) {
	item.clear();
}

bool tl2::details::VectorEitherService6ErrorVectorService6FindResultRowWriteJSON(std::ostream& s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) {
	if (!::tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorEitherService6ErrorVectorService6FindResultRowRead(::basictl::tl_istream & s, std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) {
	if (!::tl2::details::BuiltinVectorEitherService6ErrorVectorService6FindResultRowRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorEitherService6ErrorVectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Either<::tl2::service6::Error, std::vector<::tl2::service6::FindResultRow>>>& item) {
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

void tl2::details::VectorIntReset(std::vector<int32_t>& item) {
	item.clear();
}

bool tl2::details::VectorIntWriteJSON(std::ostream& s, const std::vector<int32_t>& item) {
	if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item)) { return false; }
	return true;
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

void tl2::details::VectorIntBoxedReset(std::vector<int32_t>& item) {
	item.clear();
}

bool tl2::details::VectorIntBoxedWriteJSON(std::ostream& s, const std::vector<int32_t>& item) {
	if (!::tl2::details::BuiltinVectorIntBoxedWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item) {
	if (!::tl2::details::BuiltinVectorIntBoxedRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item) {
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

void tl2::details::VectorIntegerReset(std::vector<::tl2::Integer>& item) {
	item.clear();
}

bool tl2::details::VectorIntegerWriteJSON(std::ostream& s, const std::vector<::tl2::Integer>& item) {
	if (!::tl2::details::BuiltinVectorIntegerWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntegerRead(::basictl::tl_istream & s, std::vector<::tl2::Integer>& item) {
	if (!::tl2::details::BuiltinVectorIntegerRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorIntegerWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Integer>& item) {
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

void tl2::details::VectorLongBoxedReset(std::vector<int64_t>& item) {
	item.clear();
}

bool tl2::details::VectorLongBoxedWriteJSON(std::ostream& s, const std::vector<int64_t>& item) {
	if (!::tl2::details::BuiltinVectorLongBoxedWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorLongBoxedRead(::basictl::tl_istream & s, std::vector<int64_t>& item) {
	if (!::tl2::details::BuiltinVectorLongBoxedRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorLongBoxedWrite(::basictl::tl_ostream & s, const std::vector<int64_t>& item) {
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

void tl2::details::VectorMapStringStringReset(std::vector<::tl2::Map<std::string, std::string>>& item) {
	item.clear();
}

bool tl2::details::VectorMapStringStringWriteJSON(std::ostream& s, const std::vector<::tl2::Map<std::string, std::string>>& item) {
	if (!::tl2::details::BuiltinVectorMapStringStringWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorMapStringStringRead(::basictl::tl_istream & s, std::vector<::tl2::Map<std::string, std::string>>& item) {
	if (!::tl2::details::BuiltinVectorMapStringStringRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorMapStringStringWrite(::basictl::tl_ostream & s, const std::vector<::tl2::Map<std::string, std::string>>& item) {
	if (!::tl2::details::BuiltinVectorMapStringStringWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorMapStringStringReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::Map<std::string, std::string>>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorMapStringStringRead(s, item);
}

bool tl2::details::VectorMapStringStringWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::Map<std::string, std::string>>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorMapStringStringWrite(s, item);
}

void tl2::details::VectorService1ValueReset(std::vector<::tl2::service1::Value>& item) {
	item.clear();
}

bool tl2::details::VectorService1ValueWriteJSON(std::ostream& s, const std::vector<::tl2::service1::Value>& item) {
	if (!::tl2::details::BuiltinVectorService1ValueWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService1ValueRead(::basictl::tl_istream & s, std::vector<::tl2::service1::Value>& item) {
	if (!::tl2::details::BuiltinVectorService1ValueRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService1ValueWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service1::Value>& item) {
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

void tl2::details::VectorService6FindResultRowReset(std::vector<::tl2::service6::FindResultRow>& item) {
	item.clear();
}

bool tl2::details::VectorService6FindResultRowWriteJSON(std::ostream& s, const std::vector<::tl2::service6::FindResultRow>& item) {
	if (!::tl2::details::BuiltinVectorService6FindResultRowWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService6FindResultRowRead(::basictl::tl_istream & s, std::vector<::tl2::service6::FindResultRow>& item) {
	if (!::tl2::details::BuiltinVectorService6FindResultRowRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService6FindResultRowWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindResultRow>& item) {
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

void tl2::details::VectorService6FindWithBoundsResultReset(std::vector<::tl2::service6::FindWithBoundsResult>& item) {
	item.clear();
}

bool tl2::details::VectorService6FindWithBoundsResultWriteJSON(std::ostream& s, const std::vector<::tl2::service6::FindWithBoundsResult>& item) {
	if (!::tl2::details::BuiltinVectorService6FindWithBoundsResultWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService6FindWithBoundsResultRead(::basictl::tl_istream & s, std::vector<::tl2::service6::FindWithBoundsResult>& item) {
	if (!::tl2::details::BuiltinVectorService6FindWithBoundsResultRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService6FindWithBoundsResultWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindWithBoundsResult>& item) {
	if (!::tl2::details::BuiltinVectorService6FindWithBoundsResultWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorService6FindWithBoundsResultReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service6::FindWithBoundsResult>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorService6FindWithBoundsResultRead(s, item);
}

bool tl2::details::VectorService6FindWithBoundsResultWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service6::FindWithBoundsResult>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorService6FindWithBoundsResultWrite(s, item);
}

void tl2::details::VectorStringReset(std::vector<std::string>& item) {
	item.clear();
}

bool tl2::details::VectorStringWriteJSON(std::ostream& s, const std::vector<std::string>& item) {
	if (!::tl2::details::BuiltinVectorStringWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorStringRead(::basictl::tl_istream & s, std::vector<std::string>& item) {
	if (!::tl2::details::BuiltinVectorStringRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorStringWrite(::basictl::tl_ostream & s, const std::vector<std::string>& item) {
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

void tl2::details::VectorTasksQueueTypeInfoReset(std::vector<::tl2::tasks::QueueTypeInfo>& item) {
	item.clear();
}

bool tl2::details::VectorTasksQueueTypeInfoWriteJSON(std::ostream& s, const std::vector<::tl2::tasks::QueueTypeInfo>& item) {
	if (!::tl2::details::BuiltinVectorTasksQueueTypeInfoWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorTasksQueueTypeInfoRead(::basictl::tl_istream & s, std::vector<::tl2::tasks::QueueTypeInfo>& item) {
	if (!::tl2::details::BuiltinVectorTasksQueueTypeInfoRead(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorTasksQueueTypeInfoWrite(::basictl::tl_ostream & s, const std::vector<::tl2::tasks::QueueTypeInfo>& item) {
	if (!::tl2::details::BuiltinVectorTasksQueueTypeInfoWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::VectorTasksQueueTypeInfoReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::tasks::QueueTypeInfo>& item) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false; }
	return tl2::details::VectorTasksQueueTypeInfoRead(s, item);
}

bool tl2::details::VectorTasksQueueTypeInfoWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::tasks::QueueTypeInfo>& item) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	return tl2::details::VectorTasksQueueTypeInfoWrite(s, item);
}

bool tl2::WithFloat::write_json(std::ostream& s)const {
	if (!::tl2::details::WithFloatWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::WithFloat::read(::basictl::tl_istream & s) {
	if (!::tl2::details::WithFloatRead(s, *this)) { return false; }
	return true;
}

bool tl2::WithFloat::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::WithFloatWrite(s, *this)) { return false; }
	return true;
}

bool tl2::WithFloat::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::WithFloatReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::WithFloat::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::WithFloatWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::WithFloatReset(::tl2::WithFloat& item) {
	item.x = 0;
	item.y = 0;
	item.z = 0;
}

bool tl2::details::WithFloatWriteJSON(std::ostream& s, const ::tl2::WithFloat& item) {
	s << "{";
	s << "\"x\":";
	s << item.x;
	s << ",";
	s << "\"y\":";
	s << item.y;
	s << ",";
	s << "\"z\":";
	s << item.z;
	s << "}";
	return true;
}

bool tl2::details::WithFloatRead(::basictl::tl_istream & s, ::tl2::WithFloat& item) {
	if (!s.float_read(item.x)) { return false; }
	if (!s.nat_read_exact_tag(0x824dab22)) { return false;}
	if (!s.float_read(item.y)) { return false; }
	if (!s.float_read(item.z)) { return false; }
	return true;
}

bool tl2::details::WithFloatWrite(::basictl::tl_ostream & s, const ::tl2::WithFloat& item) {
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
