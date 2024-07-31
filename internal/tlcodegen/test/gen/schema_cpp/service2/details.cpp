#include "headers/service2_tuple.hpp"
#include "headers/service2.setObjectTtl.hpp"
#include "headers/service2.set.hpp"
#include "headers/service2.addOrIncrMany.hpp"
#include "headers/service2.deltaSet.hpp"
#include "headers/service2.objectId.hpp"
#include "headers/service2.counterSet.hpp"
#include "../__common_namespace/headers/true.hpp"
#include "../__common_namespace/headers/int.hpp"
#include "headers/service2_double.hpp"


void tl2::details::BuiltinTupleDoubleReset(std::vector<double>& item) {
	item.resize(0);
}

bool tl2::details::BuiltinTupleDoubleWriteJSON(std::ostream & s, const std::vector<double>& item, uint32_t nat_n) {
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

bool tl2::details::BuiltinTupleDoubleRead(::basictl::tl_istream & s, std::vector<double>& item, uint32_t nat_n) {
	// TODO - check length sanity
	item.resize(nat_n);
	for(auto && el : item) {
		if (!s.double_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTupleDoubleWrite(::basictl::tl_ostream & s, const std::vector<double>& item, uint32_t nat_n) {
	if (item.size() != nat_n)
		return s.set_error_sequence_length();
	for(const auto & el : item) {
		if (!s.double_write(el)) { return false;}
	}
	return true;
}

void tl2::details::BuiltinTupleService2CounterSetReset(std::vector<::tl2::service2::CounterSet>& item) {
	item.resize(0);
}

bool tl2::details::BuiltinTupleService2CounterSetWriteJSON(std::ostream & s, const std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_n, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum) {
	if (item.size() != nat_n) {
		// TODO add exception
		return false;
	}
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::Service2CounterSetWriteJSON(s, el, nat_tintCountersNum, nat_tfloatCountersNum)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTupleService2CounterSetRead(::basictl::tl_istream & s, std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_n, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum) {
	// TODO - check length sanity
	item.resize(nat_n);
	for(auto && el : item) {
		if (!::tl2::details::Service2CounterSetRead(s, el, nat_tintCountersNum, nat_tfloatCountersNum)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTupleService2CounterSetWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service2::CounterSet>& item, uint32_t nat_n, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum) {
	if (item.size() != nat_n)
		return s.set_error_sequence_length();
	for(const auto & el : item) {
		if (!::tl2::details::Service2CounterSetWrite(s, el, nat_tintCountersNum, nat_tfloatCountersNum)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinTupleService2DeltaSetReset(std::vector<::tl2::service2::DeltaSet>& item) {
	item.resize(0);
}

bool tl2::details::BuiltinTupleService2DeltaSetWriteJSON(std::ostream & s, const std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_n, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum) {
	if (item.size() != nat_n) {
		// TODO add exception
		return false;
	}
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::Service2DeltaSetWriteJSON(s, el, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTupleService2DeltaSetRead(::basictl::tl_istream & s, std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_n, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum) {
	// TODO - check length sanity
	item.resize(nat_n);
	for(auto && el : item) {
		if (!::tl2::details::Service2DeltaSetRead(s, el, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTupleService2DeltaSetWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_n, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum) {
	if (item.size() != nat_n)
		return s.set_error_sequence_length();
	for(const auto & el : item) {
		if (!::tl2::details::Service2DeltaSetWrite(s, el, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum)) { return false; }
	}
	return true;
}

bool tl2::service2::AddOrIncrMany::write_json(std::ostream& s)const {
	if (!::tl2::details::Service2AddOrIncrManyWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service2::AddOrIncrMany::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service2AddOrIncrManyRead(s, *this)) { return false; }
	return true;
}

bool tl2::service2::AddOrIncrMany::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service2AddOrIncrManyWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service2::AddOrIncrMany::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service2AddOrIncrManyReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service2::AddOrIncrMany::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service2AddOrIncrManyWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service2AddOrIncrManyReset(::tl2::service2::AddOrIncrMany& item) {
	item.objectIdLength = 0;
	item.intCountersNum = 0;
	item.floatCountersNum = 0;
	item.objectsNum = 0;
	item.intCounters.clear();
	item.floatCounters.clear();
	item.deltas.clear();
}

bool tl2::details::Service2AddOrIncrManyWriteJSON(std::ostream& s, const ::tl2::service2::AddOrIncrMany& item) {
	s << "{";
	if (item.objectIdLength != 0) {
		s << "\"objectIdLength\":";
		s << item.objectIdLength;
	}
	if (item.intCountersNum != 0) {
		s << ",";
		s << "\"intCountersNum\":";
		s << item.intCountersNum;
	}
	if (item.floatCountersNum != 0) {
		s << ",";
		s << "\"floatCountersNum\":";
		s << item.floatCountersNum;
	}
	if (item.objectsNum != 0) {
		s << ",";
		s << "\"objectsNum\":";
		s << item.objectsNum;
	}
	if ((item.intCounters.size() != 0) || (item.intCountersNum != 0)) {
		s << ",";
		s << "\"intCounters\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.intCounters, item.intCountersNum)) { return false; }
	}
	if ((item.floatCounters.size() != 0) || (item.floatCountersNum != 0)) {
		s << ",";
		s << "\"floatCounters\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.floatCounters, item.floatCountersNum)) { return false; }
	}
	if ((item.deltas.size() != 0) || (item.objectIdLength != 0)) {
		s << ",";
		s << "\"deltas\":";
		if (!::tl2::details::BuiltinTupleService2DeltaSetWriteJSON(s, item.deltas, item.objectsNum, item.objectIdLength, item.intCountersNum, item.floatCountersNum)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service2AddOrIncrManyRead(::basictl::tl_istream & s, ::tl2::service2::AddOrIncrMany& item) {
	if (!s.nat_read(item.objectIdLength)) { return false; }
	if (!s.nat_read(item.intCountersNum)) { return false; }
	if (!s.nat_read(item.floatCountersNum)) { return false; }
	if (!s.nat_read(item.objectsNum)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.intCounters, item.intCountersNum)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.floatCounters, item.floatCountersNum)) { return false; }
	if (!::tl2::details::BuiltinTupleService2DeltaSetRead(s, item.deltas, item.objectsNum, item.objectIdLength, item.intCountersNum, item.floatCountersNum)) { return false; }
	return true;
}

bool tl2::details::Service2AddOrIncrManyWrite(::basictl::tl_ostream & s, const ::tl2::service2::AddOrIncrMany& item) {
	if (!s.nat_write(item.objectIdLength)) { return false;}
	if (!s.nat_write(item.intCountersNum)) { return false;}
	if (!s.nat_write(item.floatCountersNum)) { return false;}
	if (!s.nat_write(item.objectsNum)) { return false;}
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.intCounters, item.intCountersNum)) { return false; }
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.floatCounters, item.floatCountersNum)) { return false; }
	if (!::tl2::details::BuiltinTupleService2DeltaSetWrite(s, item.deltas, item.objectsNum, item.objectIdLength, item.intCountersNum, item.floatCountersNum)) { return false; }
	return true;
}

bool tl2::details::Service2AddOrIncrManyReadBoxed(::basictl::tl_istream & s, ::tl2::service2::AddOrIncrMany& item) {
	if (!s.nat_read_exact_tag(0x5aa52489)) { return false; }
	return tl2::details::Service2AddOrIncrManyRead(s, item);
}

bool tl2::details::Service2AddOrIncrManyWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::AddOrIncrMany& item) {
	if (!s.nat_write(0x5aa52489)) { return false; }
	return tl2::details::Service2AddOrIncrManyWrite(s, item);
}

bool tl2::details::Service2AddOrIncrManyReadResult(::basictl::tl_istream & s, tl2::service2::AddOrIncrMany& item, std::vector<::tl2::service2::CounterSet>& result) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false;}
	if (!::tl2::details::BuiltinTupleService2CounterSetRead(s, result, item.objectsNum, item.intCountersNum, item.floatCountersNum)) { return false; }
	return true;
}
bool tl2::details::Service2AddOrIncrManyWriteResult(::basictl::tl_ostream & s, tl2::service2::AddOrIncrMany& item, std::vector<::tl2::service2::CounterSet>& result) {
	if (!s.nat_write(0x9770768a)) { return false; }
	if (!::tl2::details::BuiltinTupleService2CounterSetWrite(s, result, item.objectsNum, item.intCountersNum, item.floatCountersNum)) { return false; }
	return true;
}

bool tl2::service2::AddOrIncrMany::read_result(::basictl::tl_istream & s, std::vector<::tl2::service2::CounterSet> & result) {
	return tl2::details::Service2AddOrIncrManyReadResult(s, *this, result);
}
bool tl2::service2::AddOrIncrMany::write_result(::basictl::tl_ostream & s, std::vector<::tl2::service2::CounterSet> & result) {
	return tl2::details::Service2AddOrIncrManyWriteResult(s, *this, result);
}

bool tl2::service2::CounterSet::write_json(std::ostream& s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const {
	if (!::tl2::details::Service2CounterSetWriteJSON(s, *this, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	return true;
}

bool tl2::service2::CounterSet::read(::basictl::tl_istream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	if (!::tl2::details::Service2CounterSetRead(s, *this, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	return true;
}

bool tl2::service2::CounterSet::write(::basictl::tl_ostream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const {
	if (!::tl2::details::Service2CounterSetWrite(s, *this, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	return true;
}

bool tl2::service2::CounterSet::read_boxed(::basictl::tl_istream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	if (!::tl2::details::Service2CounterSetReadBoxed(s, *this, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	return true;
}

bool tl2::service2::CounterSet::write_boxed(::basictl::tl_ostream & s, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const {
	if (!::tl2::details::Service2CounterSetWriteBoxed(s, *this, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	return true;
}

void tl2::details::Service2CounterSetReset(::tl2::service2::CounterSet& item) {
	item.intCounters.clear();
	item.floatCounters.clear();
}

bool tl2::details::Service2CounterSetWriteJSON(std::ostream& s, const ::tl2::service2::CounterSet& item, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	s << "{";
	if ((item.intCounters.size() != 0) || (nat_intCountersNum != 0)) {
		s << "\"intCounters\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.intCounters, nat_intCountersNum)) { return false; }
	}
	if ((item.floatCounters.size() != 0) || (nat_floatCountersNum != 0)) {
		s << ",";
		s << "\"floatCounters\":";
		if (!::tl2::details::BuiltinTupleDoubleWriteJSON(s, item.floatCounters, nat_floatCountersNum)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service2CounterSetRead(::basictl::tl_istream & s, ::tl2::service2::CounterSet& item, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	if (!::tl2::details::BuiltinTupleIntRead(s, item.intCounters, nat_intCountersNum)) { return false; }
	if (!::tl2::details::BuiltinTupleDoubleRead(s, item.floatCounters, nat_floatCountersNum)) { return false; }
	return true;
}

bool tl2::details::Service2CounterSetWrite(::basictl::tl_ostream & s, const ::tl2::service2::CounterSet& item, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.intCounters, nat_intCountersNum)) { return false; }
	if (!::tl2::details::BuiltinTupleDoubleWrite(s, item.floatCounters, nat_floatCountersNum)) { return false; }
	return true;
}

bool tl2::details::Service2CounterSetReadBoxed(::basictl::tl_istream & s, ::tl2::service2::CounterSet& item, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	if (!s.nat_read_exact_tag(0xf5403fd9)) { return false; }
	return tl2::details::Service2CounterSetRead(s, item, nat_intCountersNum, nat_floatCountersNum);
}

bool tl2::details::Service2CounterSetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::CounterSet& item, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	if (!s.nat_write(0xf5403fd9)) { return false; }
	return tl2::details::Service2CounterSetWrite(s, item, nat_intCountersNum, nat_floatCountersNum);
}

bool tl2::service2::DeltaSet::write_json(std::ostream& s, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const {
	if (!::tl2::details::Service2DeltaSetWriteJSON(s, *this, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	return true;
}

bool tl2::service2::DeltaSet::read(::basictl::tl_istream & s, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	if (!::tl2::details::Service2DeltaSetRead(s, *this, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	return true;
}

bool tl2::service2::DeltaSet::write(::basictl::tl_ostream & s, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const {
	if (!::tl2::details::Service2DeltaSetWrite(s, *this, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	return true;
}

bool tl2::service2::DeltaSet::read_boxed(::basictl::tl_istream & s, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	if (!::tl2::details::Service2DeltaSetReadBoxed(s, *this, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	return true;
}

bool tl2::service2::DeltaSet::write_boxed(::basictl::tl_ostream & s, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum)const {
	if (!::tl2::details::Service2DeltaSetWriteBoxed(s, *this, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	return true;
}

void tl2::details::Service2DeltaSetReset(::tl2::service2::DeltaSet& item) {
	::tl2::details::Service2ObjectIdReset(item.id);
	::tl2::details::Service2CounterSetReset(item.counters);
}

bool tl2::details::Service2DeltaSetWriteJSON(std::ostream& s, const ::tl2::service2::DeltaSet& item, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	s << "{";
	s << "\"id\":";
	if (!::tl2::details::Service2ObjectIdWriteJSON(s, item.id, nat_objectIdLength)) { return false; }
	s << ",";
	s << "\"counters\":";
	if (!::tl2::details::Service2CounterSetWriteJSON(s, item.counters, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::Service2DeltaSetRead(::basictl::tl_istream & s, ::tl2::service2::DeltaSet& item, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	if (!::tl2::details::Service2ObjectIdRead(s, item.id, nat_objectIdLength)) { return false; }
	if (!::tl2::details::Service2CounterSetRead(s, item.counters, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	return true;
}

bool tl2::details::Service2DeltaSetWrite(::basictl::tl_ostream & s, const ::tl2::service2::DeltaSet& item, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	if (!::tl2::details::Service2ObjectIdWrite(s, item.id, nat_objectIdLength)) { return false; }
	if (!::tl2::details::Service2CounterSetWrite(s, item.counters, nat_intCountersNum, nat_floatCountersNum)) { return false; }
	return true;
}

bool tl2::details::Service2DeltaSetReadBoxed(::basictl::tl_istream & s, ::tl2::service2::DeltaSet& item, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	if (!s.nat_read_exact_tag(0xbf49abc2)) { return false; }
	return tl2::details::Service2DeltaSetRead(s, item, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum);
}

bool tl2::details::Service2DeltaSetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::DeltaSet& item, uint32_t nat_objectIdLength, uint32_t nat_intCountersNum, uint32_t nat_floatCountersNum) {
	if (!s.nat_write(0xbf49abc2)) { return false; }
	return tl2::details::Service2DeltaSetWrite(s, item, nat_objectIdLength, nat_intCountersNum, nat_floatCountersNum);
}

bool tl2::service2::ObjectId::write_json(std::ostream& s, uint32_t nat_objectIdLength)const {
	if (!::tl2::details::Service2ObjectIdWriteJSON(s, *this, nat_objectIdLength)) { return false; }
	return true;
}

bool tl2::service2::ObjectId::read(::basictl::tl_istream & s, uint32_t nat_objectIdLength) {
	if (!::tl2::details::Service2ObjectIdRead(s, *this, nat_objectIdLength)) { return false; }
	return true;
}

bool tl2::service2::ObjectId::write(::basictl::tl_ostream & s, uint32_t nat_objectIdLength)const {
	if (!::tl2::details::Service2ObjectIdWrite(s, *this, nat_objectIdLength)) { return false; }
	return true;
}

bool tl2::service2::ObjectId::read_boxed(::basictl::tl_istream & s, uint32_t nat_objectIdLength) {
	if (!::tl2::details::Service2ObjectIdReadBoxed(s, *this, nat_objectIdLength)) { return false; }
	return true;
}

bool tl2::service2::ObjectId::write_boxed(::basictl::tl_ostream & s, uint32_t nat_objectIdLength)const {
	if (!::tl2::details::Service2ObjectIdWriteBoxed(s, *this, nat_objectIdLength)) { return false; }
	return true;
}

void tl2::details::Service2ObjectIdReset(::tl2::service2::ObjectId& item) {
	item.id.clear();
}

bool tl2::details::Service2ObjectIdWriteJSON(std::ostream& s, const ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength) {
	s << "{";
	if ((item.id.size() != 0) || (nat_objectIdLength != 0)) {
		s << "\"id\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.id, nat_objectIdLength)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::Service2ObjectIdRead(::basictl::tl_istream & s, ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength) {
	if (!::tl2::details::BuiltinTupleIntRead(s, item.id, nat_objectIdLength)) { return false; }
	return true;
}

bool tl2::details::Service2ObjectIdWrite(::basictl::tl_ostream & s, const ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength) {
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.id, nat_objectIdLength)) { return false; }
	return true;
}

bool tl2::details::Service2ObjectIdReadBoxed(::basictl::tl_istream & s, ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength) {
	if (!s.nat_read_exact_tag(0xaa0af282)) { return false; }
	return tl2::details::Service2ObjectIdRead(s, item, nat_objectIdLength);
}

bool tl2::details::Service2ObjectIdWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::ObjectId& item, uint32_t nat_objectIdLength) {
	if (!s.nat_write(0xaa0af282)) { return false; }
	return tl2::details::Service2ObjectIdWrite(s, item, nat_objectIdLength);
}

bool tl2::service2::Set::write_json(std::ostream& s)const {
	if (!::tl2::details::Service2SetWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service2::Set::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service2SetRead(s, *this)) { return false; }
	return true;
}

bool tl2::service2::Set::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service2SetWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service2::Set::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service2SetReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service2::Set::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service2SetWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service2SetReset(::tl2::service2::Set& item) {
	item.objectIdLength = 0;
	item.intCountersNum = 0;
	item.floatCountersNum = 0;
	item.intCounters.clear();
	item.floatCounters.clear();
	::tl2::details::Service2DeltaSetReset(item.newValues);
}

bool tl2::details::Service2SetWriteJSON(std::ostream& s, const ::tl2::service2::Set& item) {
	s << "{";
	if (item.objectIdLength != 0) {
		s << "\"objectIdLength\":";
		s << item.objectIdLength;
	}
	if (item.intCountersNum != 0) {
		s << ",";
		s << "\"intCountersNum\":";
		s << item.intCountersNum;
	}
	if (item.floatCountersNum != 0) {
		s << ",";
		s << "\"floatCountersNum\":";
		s << item.floatCountersNum;
	}
	if ((item.intCounters.size() != 0) || (item.intCountersNum != 0)) {
		s << ",";
		s << "\"intCounters\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.intCounters, item.intCountersNum)) { return false; }
	}
	if ((item.floatCounters.size() != 0) || (item.floatCountersNum != 0)) {
		s << ",";
		s << "\"floatCounters\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.floatCounters, item.floatCountersNum)) { return false; }
	}
	s << ",";
	s << "\"newValues\":";
	if (!::tl2::details::Service2DeltaSetWriteJSON(s, item.newValues, item.objectIdLength, item.intCountersNum, item.floatCountersNum)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::Service2SetRead(::basictl::tl_istream & s, ::tl2::service2::Set& item) {
	if (!s.nat_read(item.objectIdLength)) { return false; }
	if (!s.nat_read(item.intCountersNum)) { return false; }
	if (!s.nat_read(item.floatCountersNum)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.intCounters, item.intCountersNum)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.floatCounters, item.floatCountersNum)) { return false; }
	if (!::tl2::details::Service2DeltaSetRead(s, item.newValues, item.objectIdLength, item.intCountersNum, item.floatCountersNum)) { return false; }
	return true;
}

bool tl2::details::Service2SetWrite(::basictl::tl_ostream & s, const ::tl2::service2::Set& item) {
	if (!s.nat_write(item.objectIdLength)) { return false;}
	if (!s.nat_write(item.intCountersNum)) { return false;}
	if (!s.nat_write(item.floatCountersNum)) { return false;}
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.intCounters, item.intCountersNum)) { return false; }
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.floatCounters, item.floatCountersNum)) { return false; }
	if (!::tl2::details::Service2DeltaSetWrite(s, item.newValues, item.objectIdLength, item.intCountersNum, item.floatCountersNum)) { return false; }
	return true;
}

bool tl2::details::Service2SetReadBoxed(::basictl::tl_istream & s, ::tl2::service2::Set& item) {
	if (!s.nat_read_exact_tag(0x0d31f63d)) { return false; }
	return tl2::details::Service2SetRead(s, item);
}

bool tl2::details::Service2SetWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::Set& item) {
	if (!s.nat_write(0x0d31f63d)) { return false; }
	return tl2::details::Service2SetWrite(s, item);
}

bool tl2::details::Service2SetReadResult(::basictl::tl_istream & s, tl2::service2::Set& item, ::tl2::True& result) {
	if (!::tl2::details::TrueReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service2SetWriteResult(::basictl::tl_ostream & s, tl2::service2::Set& item, ::tl2::True& result) {
	if (!::tl2::details::TrueWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service2::Set::read_result(::basictl::tl_istream & s, ::tl2::True & result) {
	return tl2::details::Service2SetReadResult(s, *this, result);
}
bool tl2::service2::Set::write_result(::basictl::tl_ostream & s, ::tl2::True & result) {
	return tl2::details::Service2SetWriteResult(s, *this, result);
}

bool tl2::service2::SetObjectTtl::write_json(std::ostream& s)const {
	if (!::tl2::details::Service2SetObjectTtlWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::service2::SetObjectTtl::read(::basictl::tl_istream & s) {
	if (!::tl2::details::Service2SetObjectTtlRead(s, *this)) { return false; }
	return true;
}

bool tl2::service2::SetObjectTtl::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service2SetObjectTtlWrite(s, *this)) { return false; }
	return true;
}

bool tl2::service2::SetObjectTtl::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::Service2SetObjectTtlReadBoxed(s, *this)) { return false; }
	return true;
}

bool tl2::service2::SetObjectTtl::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::Service2SetObjectTtlWriteBoxed(s, *this)) { return false; }
	return true;
}

void tl2::details::Service2SetObjectTtlReset(::tl2::service2::SetObjectTtl& item) {
	item.objectIdLength = 0;
	::tl2::details::Service2ObjectIdReset(item.objectId);
	item.ttl = 0;
}

bool tl2::details::Service2SetObjectTtlWriteJSON(std::ostream& s, const ::tl2::service2::SetObjectTtl& item) {
	s << "{";
	if (item.objectIdLength != 0) {
		s << "\"objectIdLength\":";
		s << item.objectIdLength;
	}
	s << ",";
	s << "\"objectId\":";
	if (!::tl2::details::Service2ObjectIdWriteJSON(s, item.objectId, item.objectIdLength)) { return false; }
	if (item.ttl != 0) {
		s << ",";
		s << "\"ttl\":";
		s << item.ttl;
	}
	s << "}";
	return true;
}

bool tl2::details::Service2SetObjectTtlRead(::basictl::tl_istream & s, ::tl2::service2::SetObjectTtl& item) {
	if (!s.nat_read(item.objectIdLength)) { return false; }
	if (!::tl2::details::Service2ObjectIdRead(s, item.objectId, item.objectIdLength)) { return false; }
	if (!s.int_read(item.ttl)) { return false; }
	return true;
}

bool tl2::details::Service2SetObjectTtlWrite(::basictl::tl_ostream & s, const ::tl2::service2::SetObjectTtl& item) {
	if (!s.nat_write(item.objectIdLength)) { return false;}
	if (!::tl2::details::Service2ObjectIdWrite(s, item.objectId, item.objectIdLength)) { return false; }
	if (!s.int_write(item.ttl)) { return false;}
	return true;
}

bool tl2::details::Service2SetObjectTtlReadBoxed(::basictl::tl_istream & s, ::tl2::service2::SetObjectTtl& item) {
	if (!s.nat_read_exact_tag(0x6f98f025)) { return false; }
	return tl2::details::Service2SetObjectTtlRead(s, item);
}

bool tl2::details::Service2SetObjectTtlWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service2::SetObjectTtl& item) {
	if (!s.nat_write(0x6f98f025)) { return false; }
	return tl2::details::Service2SetObjectTtlWrite(s, item);
}

bool tl2::details::Service2SetObjectTtlReadResult(::basictl::tl_istream & s, tl2::service2::SetObjectTtl& item, ::tl2::True& result) {
	if (!::tl2::details::TrueReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::Service2SetObjectTtlWriteResult(::basictl::tl_ostream & s, tl2::service2::SetObjectTtl& item, ::tl2::True& result) {
	if (!::tl2::details::TrueWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::service2::SetObjectTtl::read_result(::basictl::tl_istream & s, ::tl2::True & result) {
	return tl2::details::Service2SetObjectTtlReadResult(s, *this, result);
}
bool tl2::service2::SetObjectTtl::write_result(::basictl::tl_ostream & s, ::tl2::True & result) {
	return tl2::details::Service2SetObjectTtlWriteResult(s, *this, result);
}

void tl2::details::TupleDoubleReset(std::vector<double>& item) {
	item.clear();
}

bool tl2::details::TupleDoubleWriteJSON(std::ostream& s, const std::vector<double>& item, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleDoubleWriteJSON(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleDoubleRead(::basictl::tl_istream & s, std::vector<double>& item, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleDoubleRead(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleDoubleWrite(::basictl::tl_ostream & s, const std::vector<double>& item, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleDoubleWrite(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleDoubleReadBoxed(::basictl::tl_istream & s, std::vector<double>& item, uint32_t nat_n) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleDoubleRead(s, item, nat_n);
}

bool tl2::details::TupleDoubleWriteBoxed(::basictl::tl_ostream & s, const std::vector<double>& item, uint32_t nat_n) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleDoubleWrite(s, item, nat_n);
}

void tl2::details::TupleService2DeltaSetReset(std::vector<::tl2::service2::DeltaSet>& item) {
	item.clear();
}

bool tl2::details::TupleService2DeltaSetWriteJSON(std::ostream& s, const std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleService2DeltaSetWriteJSON(s, item, nat_n, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum)) { return false; }
	return true;
}

bool tl2::details::TupleService2DeltaSetRead(::basictl::tl_istream & s, std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleService2DeltaSetRead(s, item, nat_n, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum)) { return false; }
	return true;
}

bool tl2::details::TupleService2DeltaSetWrite(::basictl::tl_ostream & s, const std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleService2DeltaSetWrite(s, item, nat_n, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum)) { return false; }
	return true;
}

bool tl2::details::TupleService2DeltaSetReadBoxed(::basictl::tl_istream & s, std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleService2DeltaSetRead(s, item, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum, nat_n);
}

bool tl2::details::TupleService2DeltaSetWriteBoxed(::basictl::tl_ostream & s, const std::vector<::tl2::service2::DeltaSet>& item, uint32_t nat_tobjectIdLength, uint32_t nat_tintCountersNum, uint32_t nat_tfloatCountersNum, uint32_t nat_n) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleService2DeltaSetWrite(s, item, nat_tobjectIdLength, nat_tintCountersNum, nat_tfloatCountersNum, nat_n);
}
