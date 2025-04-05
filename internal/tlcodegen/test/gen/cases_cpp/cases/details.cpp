#include "headers/cases_vector.h"
#include "headers/cases_tuple.h"
#include "headers/cases.testVector.h"
#include "headers/cases.testUnionContainer.h"
#include "headers/cases.TestUnion.h"
#include "headers/cases.testUnion2.h"
#include "headers/cases.testUnion1.h"
#include "headers/cases.testTuple.h"
#include "headers/cases.testRecursiveFieldMask.h"
#include "headers/cases.testOutFieldMaskContainer.h"
#include "headers/cases.testOutFieldMask.h"
#include "headers/cases.testMaybe.h"
#include "headers/cases_int.h"
#include "headers/cases.testLocalFieldmask.h"
#include "headers/cases.testInplaceStructArgs2.h"
#include "headers/cases_pair.h"
#include "headers/cases.testInplaceStructArgs.h"
#include "headers/cases.testEnumContainer.h"
#include "headers/cases.testDictString.h"
#include "headers/cases.testDictInt.h"
#include "headers/cases.testDictAny.h"
#include "headers/cases.testBeforeReadBitValidation.h"
#include "headers/cases.testArray.h"
#include "headers/cases.testAllPossibleFieldConfigsContainer.h"
#include "headers/cases.testAllPossibleFieldConfigs.h"
#include "headers/cases.replace7plusplus.h"
#include "headers/cases.replace7plus.h"
#include "headers/cases.replace7.h"
#include "headers/cases.myCycle1.h"
#include "headers/cases.myCycle2.h"
#include "headers/cases.myCycle3.h"
#include "headers/cases.inplace1.h"
#include "headers/cases.inplace2.h"
#include "headers/cases.inplace3.h"
#include "headers/cases.TestEnum.h"
#include "headers/cases.TestEnumItems.h"
#include "../__common_namespace/headers/dictionary.h"
#include "../__common_namespace/headers/dictionaryFieldAny.h"
#include "../__common_namespace/headers/dictionaryAny.h"
#include "../__common_namespace/headers/true.h"


void tl2::details::BuiltinTuple2IntReset(std::array<int32_t, 2>& item) {
	for(auto && el : item) {
		el = 0;
	}
}

bool tl2::details::BuiltinTuple2IntWriteJSON(std::ostream &s, const std::array<int32_t, 2>& item) {
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

bool tl2::details::BuiltinTuple2IntRead(::basictl::tl_istream & s, std::array<int32_t, 2>& item) {
	for(auto && el : item) {
		if (!s.int_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTuple2IntWrite(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item) {
	for(const auto & el : item) {
		if (!s.int_write(el)) { return false;}
	}
	return true;
}

void tl2::details::BuiltinTuple2PairTupleIntTupleIntReset(std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item) {
	for(auto && el : item) {
		::tl2::details::PairTupleIntTupleIntReset(el);
	}
}

bool tl2::details::BuiltinTuple2PairTupleIntTupleIntWriteJSON(std::ostream &s, const std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn) {
	s << "[";
	size_t index = 0;
	for(auto && el : item) {
		if (!::tl2::details::PairTupleIntTupleIntWriteJSON(s, el, nat_tXn, nat_tYn)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTuple2PairTupleIntTupleIntRead(::basictl::tl_istream & s, std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn) {
	for(auto && el : item) {
		if (!::tl2::details::PairTupleIntTupleIntRead(s, el, nat_tXn, nat_tYn)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTuple2PairTupleIntTupleIntWrite(::basictl::tl_ostream & s, const std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn) {
	for(const auto & el : item) {
		if (!::tl2::details::PairTupleIntTupleIntWrite(s, el, nat_tXn, nat_tYn)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinTuple4IntReset(std::array<int32_t, 4>& item) {
	for(auto && el : item) {
		el = 0;
	}
}

bool tl2::details::BuiltinTuple4IntWriteJSON(std::ostream &s, const std::array<int32_t, 4>& item) {
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

bool tl2::details::BuiltinTuple4IntRead(::basictl::tl_istream & s, std::array<int32_t, 4>& item) {
	for(auto && el : item) {
		if (!s.int_read(el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTuple4IntWrite(::basictl::tl_ostream & s, const std::array<int32_t, 4>& item) {
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

void tl2::details::BuiltinTupleTupleIntReset(std::vector<std::vector<int32_t>>& item) {
	item.resize(0);
}

bool tl2::details::BuiltinTupleTupleIntWriteJSON(std::ostream & s, const std::vector<std::vector<int32_t>>& item, uint32_t nat_n, uint32_t nat_t) {
	if (item.size() != nat_n) {
		// TODO add exception
		return false;
	}
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, el, nat_t)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTupleTupleIntRead(::basictl::tl_istream & s, std::vector<std::vector<int32_t>>& item, uint32_t nat_n, uint32_t nat_t) {
	// TODO - check length sanity
	item.resize(nat_n);
	for(auto && el : item) {
		if (!::tl2::details::BuiltinTupleIntRead(s, el, nat_t)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTupleTupleIntWrite(::basictl::tl_ostream & s, const std::vector<std::vector<int32_t>>& item, uint32_t nat_n, uint32_t nat_t) {
	if (item.size() != nat_n)
		return s.set_error_sequence_length();
	for(const auto & el : item) {
		if (!::tl2::details::BuiltinTupleIntWrite(s, el, nat_t)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinTupleTupleInt2Reset(std::vector<std::array<int32_t, 2>>& item) {
	item.resize(0);
}

bool tl2::details::BuiltinTupleTupleInt2WriteJSON(std::ostream & s, const std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n) {
	if (item.size() != nat_n) {
		// TODO add exception
		return false;
	}
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::BuiltinTuple2IntWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTupleTupleInt2Read(::basictl::tl_istream & s, std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n) {
	// TODO - check length sanity
	item.resize(nat_n);
	for(auto && el : item) {
		if (!::tl2::details::BuiltinTuple2IntRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTupleTupleInt2Write(::basictl::tl_ostream & s, const std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n) {
	if (item.size() != nat_n)
		return s.set_error_sequence_length();
	for(const auto & el : item) {
		if (!::tl2::details::BuiltinTuple2IntWrite(s, el)) { return false; }
	}
	return true;
}

void tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2Reset(std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item) {
	item.resize(0);
}

bool tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2WriteJSON(std::ostream & s, const std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_n, uint32_t nat_ttXn, uint32_t nat_ttYn) {
	if (item.size() != nat_n) {
		// TODO add exception
		return false;
	}
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::BuiltinTuple2PairTupleIntTupleIntWriteJSON(s, el, nat_ttXn, nat_ttYn)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2Read(::basictl::tl_istream & s, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_n, uint32_t nat_ttXn, uint32_t nat_ttYn) {
	// TODO - check length sanity
	item.resize(nat_n);
	for(auto && el : item) {
		if (!::tl2::details::BuiltinTuple2PairTupleIntTupleIntRead(s, el, nat_ttXn, nat_ttYn)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2Write(::basictl::tl_ostream & s, const std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_n, uint32_t nat_ttXn, uint32_t nat_ttYn) {
	if (item.size() != nat_n)
		return s.set_error_sequence_length();
	for(const auto & el : item) {
		if (!::tl2::details::BuiltinTuple2PairTupleIntTupleIntWrite(s, el, nat_ttXn, nat_ttYn)) { return false; }
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

void tl2::details::CasesInplace1IntReset(::tl2::cases::Inplace1<int32_t>& item) {
	::tl2::details::CasesInplace2IntReset(item.value);
}

bool tl2::details::CasesInplace1IntWriteJSON(std::ostream& s, const ::tl2::cases::Inplace1<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::CasesInplace2IntWriteJSON(s, item.value, nat_a2, nat_a3, nat_a1)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesInplace1IntRead(::basictl::tl_istream & s, ::tl2::cases::Inplace1<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	if (!::tl2::details::CasesInplace2IntRead(s, item.value, nat_a2, nat_a3, nat_a1)) { return false; }
	return true;
}

bool tl2::details::CasesInplace1IntWrite(::basictl::tl_ostream & s, const ::tl2::cases::Inplace1<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	if (!::tl2::details::CasesInplace2IntWrite(s, item.value, nat_a2, nat_a3, nat_a1)) { return false; }
	return true;
}

bool tl2::details::CasesInplace1IntReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Inplace1<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	if (!s.nat_read_exact_tag(0x5533e8e9)) { return false; }
	return tl2::details::CasesInplace1IntRead(s, item, nat_a1, nat_a2, nat_a3);
}

bool tl2::details::CasesInplace1IntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Inplace1<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	if (!s.nat_write(0x5533e8e9)) { return false; }
	return tl2::details::CasesInplace1IntWrite(s, item, nat_a1, nat_a2, nat_a3);
}

void tl2::details::CasesInplace1PairTupleIntTupleIntReset(::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item) {
	::tl2::details::CasesInplace2PairTupleIntTupleIntReset(item.value);
}

bool tl2::details::CasesInplace1PairTupleIntTupleIntWriteJSON(std::ostream& s, const ::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::CasesInplace2PairTupleIntTupleIntWriteJSON(s, item.value, nat_a2, nat_a3, nat_a1, nat_XXn, nat_XYn)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesInplace1PairTupleIntTupleIntRead(::basictl::tl_istream & s, ::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) {
	if (!::tl2::details::CasesInplace2PairTupleIntTupleIntRead(s, item.value, nat_a2, nat_a3, nat_a1, nat_XXn, nat_XYn)) { return false; }
	return true;
}

bool tl2::details::CasesInplace1PairTupleIntTupleIntWrite(::basictl::tl_ostream & s, const ::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) {
	if (!::tl2::details::CasesInplace2PairTupleIntTupleIntWrite(s, item.value, nat_a2, nat_a3, nat_a1, nat_XXn, nat_XYn)) { return false; }
	return true;
}

bool tl2::details::CasesInplace1PairTupleIntTupleIntReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) {
	if (!s.nat_read_exact_tag(0x5533e8e9)) { return false; }
	return tl2::details::CasesInplace1PairTupleIntTupleIntRead(s, item, nat_a1, nat_a2, nat_a3, nat_XXn, nat_XYn);
}

bool tl2::details::CasesInplace1PairTupleIntTupleIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) {
	if (!s.nat_write(0x5533e8e9)) { return false; }
	return tl2::details::CasesInplace1PairTupleIntTupleIntWrite(s, item, nat_a1, nat_a2, nat_a3, nat_XXn, nat_XYn);
}

void tl2::details::CasesInplace2IntReset(::tl2::cases::Inplace2<int32_t>& item) {
	::tl2::details::CasesInplace3TupleInt2Reset(item.value);
}

bool tl2::details::CasesInplace2IntWriteJSON(std::ostream& s, const ::tl2::cases::Inplace2<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::CasesInplace3TupleInt2WriteJSON(s, item.value, nat_a2, nat_a3, nat_a1)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesInplace2IntRead(::basictl::tl_istream & s, ::tl2::cases::Inplace2<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	if (!::tl2::details::CasesInplace3TupleInt2Read(s, item.value, nat_a2, nat_a3, nat_a1)) { return false; }
	return true;
}

bool tl2::details::CasesInplace2IntWrite(::basictl::tl_ostream & s, const ::tl2::cases::Inplace2<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	if (!::tl2::details::CasesInplace3TupleInt2Write(s, item.value, nat_a2, nat_a3, nat_a1)) { return false; }
	return true;
}

bool tl2::details::CasesInplace2IntReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Inplace2<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	if (!s.nat_read_exact_tag(0x869fcff5)) { return false; }
	return tl2::details::CasesInplace2IntRead(s, item, nat_a1, nat_a2, nat_a3);
}

bool tl2::details::CasesInplace2IntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Inplace2<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	if (!s.nat_write(0x869fcff5)) { return false; }
	return tl2::details::CasesInplace2IntWrite(s, item, nat_a1, nat_a2, nat_a3);
}

void tl2::details::CasesInplace2PairTupleIntTupleIntReset(::tl2::cases::Inplace2<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item) {
	::tl2::details::CasesInplace3TuplePairTupleIntTupleInt2Reset(item.value);
}

bool tl2::details::CasesInplace2PairTupleIntTupleIntWriteJSON(std::ostream& s, const ::tl2::cases::Inplace2<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::CasesInplace3TuplePairTupleIntTupleInt2WriteJSON(s, item.value, nat_a2, nat_a3, nat_a1, nat_XXn, nat_XYn)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesInplace2PairTupleIntTupleIntRead(::basictl::tl_istream & s, ::tl2::cases::Inplace2<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) {
	if (!::tl2::details::CasesInplace3TuplePairTupleIntTupleInt2Read(s, item.value, nat_a2, nat_a3, nat_a1, nat_XXn, nat_XYn)) { return false; }
	return true;
}

bool tl2::details::CasesInplace2PairTupleIntTupleIntWrite(::basictl::tl_ostream & s, const ::tl2::cases::Inplace2<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) {
	if (!::tl2::details::CasesInplace3TuplePairTupleIntTupleInt2Write(s, item.value, nat_a2, nat_a3, nat_a1, nat_XXn, nat_XYn)) { return false; }
	return true;
}

bool tl2::details::CasesInplace2PairTupleIntTupleIntReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Inplace2<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) {
	if (!s.nat_read_exact_tag(0x869fcff5)) { return false; }
	return tl2::details::CasesInplace2PairTupleIntTupleIntRead(s, item, nat_a1, nat_a2, nat_a3, nat_XXn, nat_XYn);
}

bool tl2::details::CasesInplace2PairTupleIntTupleIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Inplace2<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) {
	if (!s.nat_write(0x869fcff5)) { return false; }
	return tl2::details::CasesInplace2PairTupleIntTupleIntWrite(s, item, nat_a1, nat_a2, nat_a3, nat_XXn, nat_XYn);
}

void tl2::details::CasesInplace3TupleInt2Reset(::tl2::cases::Inplace3<std::array<int32_t, 2>>& item) {
	::tl2::details::PairTupleTupleInt2TupleTupleInt2Reset(item.value);
}

bool tl2::details::CasesInplace3TupleInt2WriteJSON(std::ostream& s, const ::tl2::cases::Inplace3<std::array<int32_t, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::PairTupleTupleInt2TupleTupleInt2WriteJSON(s, item.value, nat_a2, nat_a3)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesInplace3TupleInt2Read(::basictl::tl_istream & s, ::tl2::cases::Inplace3<std::array<int32_t, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	if (!::tl2::details::PairTupleTupleInt2TupleTupleInt2Read(s, item.value, nat_a2, nat_a3)) { return false; }
	return true;
}

bool tl2::details::CasesInplace3TupleInt2Write(::basictl::tl_ostream & s, const ::tl2::cases::Inplace3<std::array<int32_t, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	if (!::tl2::details::PairTupleTupleInt2TupleTupleInt2Write(s, item.value, nat_a2, nat_a3)) { return false; }
	return true;
}

bool tl2::details::CasesInplace3TupleInt2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Inplace3<std::array<int32_t, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	if (!s.nat_read_exact_tag(0x4ffb95cb)) { return false; }
	return tl2::details::CasesInplace3TupleInt2Read(s, item, nat_a1, nat_a2, nat_a3);
}

bool tl2::details::CasesInplace3TupleInt2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Inplace3<std::array<int32_t, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) {
	if (!s.nat_write(0x4ffb95cb)) { return false; }
	return tl2::details::CasesInplace3TupleInt2Write(s, item, nat_a1, nat_a2, nat_a3);
}

void tl2::details::CasesInplace3TuplePairTupleIntTupleInt2Reset(::tl2::cases::Inplace3<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item) {
	::tl2::details::PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2Reset(item.value);
}

bool tl2::details::CasesInplace3TuplePairTupleIntTupleInt2WriteJSON(std::ostream& s, const ::tl2::cases::Inplace3<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XtXn, uint32_t nat_XtYn) {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2WriteJSON(s, item.value, nat_XtXn, nat_XtYn, nat_a2, nat_XtXn, nat_XtYn, nat_a3)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesInplace3TuplePairTupleIntTupleInt2Read(::basictl::tl_istream & s, ::tl2::cases::Inplace3<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XtXn, uint32_t nat_XtYn) {
	if (!::tl2::details::PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2Read(s, item.value, nat_XtXn, nat_XtYn, nat_a2, nat_XtXn, nat_XtYn, nat_a3)) { return false; }
	return true;
}

bool tl2::details::CasesInplace3TuplePairTupleIntTupleInt2Write(::basictl::tl_ostream & s, const ::tl2::cases::Inplace3<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XtXn, uint32_t nat_XtYn) {
	if (!::tl2::details::PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2Write(s, item.value, nat_XtXn, nat_XtYn, nat_a2, nat_XtXn, nat_XtYn, nat_a3)) { return false; }
	return true;
}

bool tl2::details::CasesInplace3TuplePairTupleIntTupleInt2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Inplace3<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XtXn, uint32_t nat_XtYn) {
	if (!s.nat_read_exact_tag(0x4ffb95cb)) { return false; }
	return tl2::details::CasesInplace3TuplePairTupleIntTupleInt2Read(s, item, nat_a1, nat_a2, nat_a3, nat_XtXn, nat_XtYn);
}

bool tl2::details::CasesInplace3TuplePairTupleIntTupleInt2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Inplace3<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XtXn, uint32_t nat_XtYn) {
	if (!s.nat_write(0x4ffb95cb)) { return false; }
	return tl2::details::CasesInplace3TuplePairTupleIntTupleInt2Write(s, item, nat_a1, nat_a2, nat_a3, nat_XtXn, nat_XtYn);
}

bool tl2::cases::MyCycle1::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesMyCycle1WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::MyCycle1::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesMyCycle1Read(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::MyCycle1::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesMyCycle1Write(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::MyCycle1::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesMyCycle1ReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::MyCycle1::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesMyCycle1WriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesMyCycle1Reset(::tl2::cases::MyCycle1& item) {
	item.fields_mask = 0;
	::tl2::details::CasesMyCycle2Reset(item.a);
}

bool tl2::details::CasesMyCycle1WriteJSON(std::ostream& s, const ::tl2::cases::MyCycle1& item) {
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
		s << "\"a\":";
		if (!::tl2::details::CasesMyCycle2WriteJSON(s, item.a)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesMyCycle1Read(::basictl::tl_istream & s, ::tl2::cases::MyCycle1& item) {
	if (!s.nat_read(item.fields_mask)) { return false; }
	if ((item.fields_mask & (1<<0)) != 0) {
		if (!::tl2::details::CasesMyCycle2Read(s, item.a)) { return false; }
	} else {
			::tl2::details::CasesMyCycle2Reset(item.a);
	}
	return true;
}

bool tl2::details::CasesMyCycle1Write(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle1& item) {
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!::tl2::details::CasesMyCycle2Write(s, item.a)) { return false; }
	}
	return true;
}

bool tl2::details::CasesMyCycle1ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::MyCycle1& item) {
	if (!s.nat_read_exact_tag(0xd3ca919d)) { return false; }
	return tl2::details::CasesMyCycle1Read(s, item);
}

bool tl2::details::CasesMyCycle1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle1& item) {
	if (!s.nat_write(0xd3ca919d)) { return false; }
	return tl2::details::CasesMyCycle1Write(s, item);
}

bool tl2::cases::MyCycle2::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesMyCycle2WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::MyCycle2::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesMyCycle2Read(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::MyCycle2::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesMyCycle2Write(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::MyCycle2::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesMyCycle2ReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::MyCycle2::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesMyCycle2WriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesMyCycle2Reset(::tl2::cases::MyCycle2& item) {
	item.fields_mask = 0;
	::tl2::details::CasesMyCycle3Reset(item.a);
}

bool tl2::details::CasesMyCycle2WriteJSON(std::ostream& s, const ::tl2::cases::MyCycle2& item) {
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
		s << "\"a\":";
		if (!::tl2::details::CasesMyCycle3WriteJSON(s, item.a)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesMyCycle2Read(::basictl::tl_istream & s, ::tl2::cases::MyCycle2& item) {
	if (!s.nat_read(item.fields_mask)) { return false; }
	if ((item.fields_mask & (1<<0)) != 0) {
		if (!::tl2::details::CasesMyCycle3Read(s, item.a)) { return false; }
	} else {
			::tl2::details::CasesMyCycle3Reset(item.a);
	}
	return true;
}

bool tl2::details::CasesMyCycle2Write(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle2& item) {
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!::tl2::details::CasesMyCycle3Write(s, item.a)) { return false; }
	}
	return true;
}

bool tl2::details::CasesMyCycle2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::MyCycle2& item) {
	if (!s.nat_read_exact_tag(0x5444c9a2)) { return false; }
	return tl2::details::CasesMyCycle2Read(s, item);
}

bool tl2::details::CasesMyCycle2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle2& item) {
	if (!s.nat_write(0x5444c9a2)) { return false; }
	return tl2::details::CasesMyCycle2Write(s, item);
}

bool tl2::cases::MyCycle3::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesMyCycle3WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::MyCycle3::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesMyCycle3Read(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::MyCycle3::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesMyCycle3Write(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::MyCycle3::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesMyCycle3ReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::MyCycle3::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesMyCycle3WriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesMyCycle3Reset(::tl2::cases::MyCycle3& item) {
	item.fields_mask = 0;

	if (item.a) {
			::tl2::details::CasesMyCycle1Reset((*item.a));
	}
}

bool tl2::details::CasesMyCycle3WriteJSON(std::ostream& s, const ::tl2::cases::MyCycle3& item) {
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
		s << "\"a\":";
		if (!::tl2::details::CasesMyCycle1WriteJSON(s, *item.a)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesMyCycle3Read(::basictl::tl_istream & s, ::tl2::cases::MyCycle3& item) {
	if (!s.nat_read(item.fields_mask)) { return false; }
	if ((item.fields_mask & (1<<0)) != 0) {
		if (!item.a) { item.a = std::make_shared<::tl2::cases::MyCycle1>(); }
		if (!::tl2::details::CasesMyCycle1Read(s, *item.a)) { return false; }
	} else {
		if (item.a) {
			::tl2::details::CasesMyCycle1Reset(*item.a);
		}
	}
	return true;
}

bool tl2::details::CasesMyCycle3Write(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle3& item) {
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!::tl2::details::CasesMyCycle1Write(s, *item.a)) { return false; }
	}
	return true;
}

bool tl2::details::CasesMyCycle3ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::MyCycle3& item) {
	if (!s.nat_read_exact_tag(0x7624f86b)) { return false; }
	return tl2::details::CasesMyCycle3Read(s, item);
}

bool tl2::details::CasesMyCycle3WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::MyCycle3& item) {
	if (!s.nat_write(0x7624f86b)) { return false; }
	return tl2::details::CasesMyCycle3Write(s, item);
}

bool tl2::cases::Replace7::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesReplace7WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::Replace7::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesReplace7Read(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::Replace7::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesReplace7Write(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::Replace7::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesReplace7ReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::Replace7::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesReplace7WriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesReplace7Reset(::tl2::cases::Replace7& item) {
	item.n = 0;
	item.m = 0;
	item.a.clear();
}

bool tl2::details::CasesReplace7WriteJSON(std::ostream& s, const ::tl2::cases::Replace7& item) {
	auto add_comma = false;
	s << "{";
	if (item.n != 0) {
		add_comma = true;
		s << "\"n\":";
		s << item.n;
	}
	if (item.m != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"m\":";
		s << item.m;
	}
	if ((item.a.size() != 0) || (item.n != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a\":";
		if (!::tl2::details::BuiltinTupleTupleIntWriteJSON(s, item.a, item.n, item.m)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesReplace7Read(::basictl::tl_istream & s, ::tl2::cases::Replace7& item) {
	if (!s.nat_read(item.n)) { return false; }
	if (!s.nat_read(item.m)) { return false; }
	if (!::tl2::details::BuiltinTupleTupleIntRead(s, item.a, item.n, item.m)) { return false; }
	return true;
}

bool tl2::details::CasesReplace7Write(::basictl::tl_ostream & s, const ::tl2::cases::Replace7& item) {
	if (!s.nat_write(item.n)) { return false;}
	if (!s.nat_write(item.m)) { return false;}
	if (!::tl2::details::BuiltinTupleTupleIntWrite(s, item.a, item.n, item.m)) { return false; }
	return true;
}

bool tl2::details::CasesReplace7ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Replace7& item) {
	if (!s.nat_read_exact_tag(0x6ccce4be)) { return false; }
	return tl2::details::CasesReplace7Read(s, item);
}

bool tl2::details::CasesReplace7WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Replace7& item) {
	if (!s.nat_write(0x6ccce4be)) { return false; }
	return tl2::details::CasesReplace7Write(s, item);
}

bool tl2::cases::Replace7plus::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesReplace7plusWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::Replace7plus::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesReplace7plusRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::Replace7plus::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesReplace7plusWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::Replace7plus::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesReplace7plusReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::Replace7plus::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesReplace7plusWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesReplace7plusReset(::tl2::cases::Replace7plus& item) {
	item.n = 0;
	item.m = 0;
	item.a.clear();
}

bool tl2::details::CasesReplace7plusWriteJSON(std::ostream& s, const ::tl2::cases::Replace7plus& item) {
	auto add_comma = false;
	s << "{";
	if (item.n != 0) {
		add_comma = true;
		s << "\"n\":";
		s << item.n;
	}
	if (item.m != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"m\":";
		s << item.m;
	}
	if ((item.n & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a\":";
		if (!::tl2::details::BuiltinTupleTupleIntWriteJSON(s, item.a, item.n, item.m)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesReplace7plusRead(::basictl::tl_istream & s, ::tl2::cases::Replace7plus& item) {
	if (!s.nat_read(item.n)) { return false; }
	if (!s.nat_read(item.m)) { return false; }
	if ((item.n & (1<<0)) != 0) {
		if (!::tl2::details::BuiltinTupleTupleIntRead(s, item.a, item.n, item.m)) { return false; }
	} else {
			item.a.clear();
	}
	return true;
}

bool tl2::details::CasesReplace7plusWrite(::basictl::tl_ostream & s, const ::tl2::cases::Replace7plus& item) {
	if (!s.nat_write(item.n)) { return false;}
	if (!s.nat_write(item.m)) { return false;}
	if ((item.n & (1<<0)) != 0) {
			if (!::tl2::details::BuiltinTupleTupleIntWrite(s, item.a, item.n, item.m)) { return false; }
	}
	return true;
}

bool tl2::details::CasesReplace7plusReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Replace7plus& item) {
	if (!s.nat_read_exact_tag(0x197858f5)) { return false; }
	return tl2::details::CasesReplace7plusRead(s, item);
}

bool tl2::details::CasesReplace7plusWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Replace7plus& item) {
	if (!s.nat_write(0x197858f5)) { return false; }
	return tl2::details::CasesReplace7plusWrite(s, item);
}

bool tl2::cases::Replace7plusplus::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesReplace7plusplusWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::Replace7plusplus::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesReplace7plusplusRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::Replace7plusplus::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesReplace7plusplusWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::Replace7plusplus::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesReplace7plusplusReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::Replace7plusplus::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesReplace7plusplusWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesReplace7plusplusReset(::tl2::cases::Replace7plusplus& item) {
	item.N = 0;
	item.M = 0;
	item.A.clear();
}

bool tl2::details::CasesReplace7plusplusWriteJSON(std::ostream& s, const ::tl2::cases::Replace7plusplus& item) {
	auto add_comma = false;
	s << "{";
	if (item.N != 0) {
		add_comma = true;
		s << "\"N\":";
		s << item.N;
	}
	if (item.M != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"M\":";
		s << item.M;
	}
	if ((item.N & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"A\":";
		if (!::tl2::details::BuiltinTupleTupleIntWriteJSON(s, item.A, item.N, item.M)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesReplace7plusplusRead(::basictl::tl_istream & s, ::tl2::cases::Replace7plusplus& item) {
	if (!s.nat_read(item.N)) { return false; }
	if (!s.nat_read(item.M)) { return false; }
	if ((item.N & (1<<0)) != 0) {
		if (!::tl2::details::BuiltinTupleTupleIntRead(s, item.A, item.N, item.M)) { return false; }
	} else {
			item.A.clear();
	}
	return true;
}

bool tl2::details::CasesReplace7plusplusWrite(::basictl::tl_ostream & s, const ::tl2::cases::Replace7plusplus& item) {
	if (!s.nat_write(item.N)) { return false;}
	if (!s.nat_write(item.M)) { return false;}
	if ((item.N & (1<<0)) != 0) {
			if (!::tl2::details::BuiltinTupleTupleIntWrite(s, item.A, item.N, item.M)) { return false; }
	}
	return true;
}

bool tl2::details::CasesReplace7plusplusReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Replace7plusplus& item) {
	if (!s.nat_read_exact_tag(0xabc39b68)) { return false; }
	return tl2::details::CasesReplace7plusplusRead(s, item);
}

bool tl2::details::CasesReplace7plusplusWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Replace7plusplus& item) {
	if (!s.nat_write(0xabc39b68)) { return false; }
	return tl2::details::CasesReplace7plusplusWrite(s, item);
}

bool tl2::cases::TestAllPossibleFieldConfigs::write_json(std::ostream& s, uint32_t nat_outer)const {
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsWriteJSON(s, *this, nat_outer)) { return false; }
	return true;
}

bool tl2::cases::TestAllPossibleFieldConfigs::read(::basictl::tl_istream & s, uint32_t nat_outer) {
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsRead(s, *this, nat_outer)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestAllPossibleFieldConfigs::write(::basictl::tl_ostream & s, uint32_t nat_outer)const {
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsWrite(s, *this, nat_outer)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestAllPossibleFieldConfigs::read_boxed(::basictl::tl_istream & s, uint32_t nat_outer) {
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsReadBoxed(s, *this, nat_outer)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestAllPossibleFieldConfigs::write_boxed(::basictl::tl_ostream & s, uint32_t nat_outer)const {
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsWriteBoxed(s, *this, nat_outer)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestAllPossibleFieldConfigsReset(::tl2::cases::TestAllPossibleFieldConfigs& item) {
	item.local = 0;
	item.f00 = 0;
	::tl2::details::TrueReset(item.f01);
	item.f02.clear();
	item.f03.clear();
	item.f10 = 0;
	::tl2::details::TrueReset(item.f11);
	item.f12.clear();
	item.f13.clear();
	item.f20 = 0;
	::tl2::details::TrueReset(item.f21);
	item.f22.clear();
	item.f23.clear();
}

bool tl2::details::CasesTestAllPossibleFieldConfigsWriteJSON(std::ostream& s, const ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer) {
	auto add_comma = false;
	s << "{";
	if (item.local != 0) {
		add_comma = true;
		s << "\"local\":";
		s << item.local;
	}
	if (item.f00 != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f00\":";
		s << item.f00;
	}
	if ((item.f02.size() != 0) || (item.local != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f02\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.f02, item.local)) { return false; }
	}
	if ((item.f03.size() != 0) || (nat_outer != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f03\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.f03, nat_outer)) { return false; }
	}
	if ((item.local & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f10\":";
		s << item.f10;
	}
	if ((item.local & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f11\":";
		if (!::tl2::details::TrueWriteJSON(s, item.f11)) { return false; }
	}
	if ((item.local & (1<<2)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f12\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.f12, item.local)) { return false; }
	}
	if ((item.local & (1<<3)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f13\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.f13, nat_outer)) { return false; }
	}
	if ((nat_outer & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f20\":";
		s << item.f20;
	}
	if ((nat_outer & (1<<2)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f22\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.f22, item.local)) { return false; }
	}
	if ((nat_outer & (1<<3)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f23\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.f23, nat_outer)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestAllPossibleFieldConfigsRead(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer) {
	if (!s.nat_read(item.local)) { return false; }
	if (!s.int_read(item.f00)) { return false; }
	if (!::tl2::details::TrueRead(s, item.f01)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.f02, item.local)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.f03, nat_outer)) { return false; }
	if ((item.local & (1<<0)) != 0) {
		if (!s.int_read(item.f10)) { return false; }
	} else {
			item.f10 = 0;
	}
	if ((item.local & (1<<1)) != 0) {
		if (!::tl2::details::TrueRead(s, item.f11)) { return false; }
	} else {
			::tl2::details::TrueReset(item.f11);
	}
	if ((item.local & (1<<2)) != 0) {
		if (!::tl2::details::BuiltinTupleIntRead(s, item.f12, item.local)) { return false; }
	} else {
			item.f12.clear();
	}
	if ((item.local & (1<<3)) != 0) {
		if (!::tl2::details::BuiltinTupleIntRead(s, item.f13, nat_outer)) { return false; }
	} else {
			item.f13.clear();
	}
	if ((nat_outer & (1<<0)) != 0) {
		if (!s.int_read(item.f20)) { return false; }
	} else {
			item.f20 = 0;
	}
	if ((nat_outer & (1<<1)) != 0) {
		if (!::tl2::details::TrueRead(s, item.f21)) { return false; }
	} else {
			::tl2::details::TrueReset(item.f21);
	}
	if ((nat_outer & (1<<2)) != 0) {
		if (!::tl2::details::BuiltinTupleIntRead(s, item.f22, item.local)) { return false; }
	} else {
			item.f22.clear();
	}
	if ((nat_outer & (1<<3)) != 0) {
		if (!::tl2::details::BuiltinTupleIntRead(s, item.f23, nat_outer)) { return false; }
	} else {
			item.f23.clear();
	}
	return true;
}

bool tl2::details::CasesTestAllPossibleFieldConfigsWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer) {
	if (!s.nat_write(item.local)) { return false;}
	if (!s.int_write(item.f00)) { return false;}
	if (!::tl2::details::TrueWrite(s, item.f01)) { return false; }
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.f02, item.local)) { return false; }
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.f03, nat_outer)) { return false; }
	if ((item.local & (1<<0)) != 0) {
			if (!s.int_write(item.f10)) { return false;}
	}
	if ((item.local & (1<<1)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.f11)) { return false; }
	}
	if ((item.local & (1<<2)) != 0) {
			if (!::tl2::details::BuiltinTupleIntWrite(s, item.f12, item.local)) { return false; }
	}
	if ((item.local & (1<<3)) != 0) {
			if (!::tl2::details::BuiltinTupleIntWrite(s, item.f13, nat_outer)) { return false; }
	}
	if ((nat_outer & (1<<0)) != 0) {
			if (!s.int_write(item.f20)) { return false;}
	}
	if ((nat_outer & (1<<1)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.f21)) { return false; }
	}
	if ((nat_outer & (1<<2)) != 0) {
			if (!::tl2::details::BuiltinTupleIntWrite(s, item.f22, item.local)) { return false; }
	}
	if ((nat_outer & (1<<3)) != 0) {
			if (!::tl2::details::BuiltinTupleIntWrite(s, item.f23, nat_outer)) { return false; }
	}
	return true;
}

bool tl2::details::CasesTestAllPossibleFieldConfigsReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer) {
	if (!s.nat_read_exact_tag(0xfb6836d3)) { return false; }
	return tl2::details::CasesTestAllPossibleFieldConfigsRead(s, item, nat_outer);
}

bool tl2::details::CasesTestAllPossibleFieldConfigsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigs& item, uint32_t nat_outer) {
	if (!s.nat_write(0xfb6836d3)) { return false; }
	return tl2::details::CasesTestAllPossibleFieldConfigsWrite(s, item, nat_outer);
}

bool tl2::cases::TestAllPossibleFieldConfigsContainer::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsContainerWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestAllPossibleFieldConfigsContainer::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsContainerRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestAllPossibleFieldConfigsContainer::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsContainerWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestAllPossibleFieldConfigsContainer::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsContainerReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestAllPossibleFieldConfigsContainer::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsContainerWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestAllPossibleFieldConfigsContainerReset(::tl2::cases::TestAllPossibleFieldConfigsContainer& item) {
	item.outer = 0;
	::tl2::details::CasesTestAllPossibleFieldConfigsReset(item.value);
}

bool tl2::details::CasesTestAllPossibleFieldConfigsContainerWriteJSON(std::ostream& s, const ::tl2::cases::TestAllPossibleFieldConfigsContainer& item) {
	auto add_comma = false;
	s << "{";
	if (item.outer != 0) {
		add_comma = true;
		s << "\"outer\":";
		s << item.outer;
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"value\":";
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsWriteJSON(s, item.value, item.outer)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesTestAllPossibleFieldConfigsContainerRead(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigsContainer& item) {
	if (!s.nat_read(item.outer)) { return false; }
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsRead(s, item.value, item.outer)) { return false; }
	return true;
}

bool tl2::details::CasesTestAllPossibleFieldConfigsContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigsContainer& item) {
	if (!s.nat_write(item.outer)) { return false;}
	if (!::tl2::details::CasesTestAllPossibleFieldConfigsWrite(s, item.value, item.outer)) { return false; }
	return true;
}

bool tl2::details::CasesTestAllPossibleFieldConfigsContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestAllPossibleFieldConfigsContainer& item) {
	if (!s.nat_read_exact_tag(0xe3fae936)) { return false; }
	return tl2::details::CasesTestAllPossibleFieldConfigsContainerRead(s, item);
}

bool tl2::details::CasesTestAllPossibleFieldConfigsContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestAllPossibleFieldConfigsContainer& item) {
	if (!s.nat_write(0xe3fae936)) { return false; }
	return tl2::details::CasesTestAllPossibleFieldConfigsContainerWrite(s, item);
}

bool tl2::cases::TestArray::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestArrayWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestArray::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestArrayRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestArray::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestArrayWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestArray::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestArrayReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestArray::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestArrayWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestArrayReset(::tl2::cases::TestArray& item) {
	item.n = 0;
	item.arr.clear();
}

bool tl2::details::CasesTestArrayWriteJSON(std::ostream& s, const ::tl2::cases::TestArray& item) {
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
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.arr, item.n)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestArrayRead(::basictl::tl_istream & s, ::tl2::cases::TestArray& item) {
	if (!s.nat_read(item.n)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.arr, item.n)) { return false; }
	return true;
}

bool tl2::details::CasesTestArrayWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestArray& item) {
	if (!s.nat_write(item.n)) { return false;}
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.arr, item.n)) { return false; }
	return true;
}

bool tl2::details::CasesTestArrayReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestArray& item) {
	if (!s.nat_read_exact_tag(0xa888030d)) { return false; }
	return tl2::details::CasesTestArrayRead(s, item);
}

bool tl2::details::CasesTestArrayWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestArray& item) {
	if (!s.nat_write(0xa888030d)) { return false; }
	return tl2::details::CasesTestArrayWrite(s, item);
}

bool tl2::cases::TestBeforeReadBitValidation::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestBeforeReadBitValidationWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestBeforeReadBitValidation::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestBeforeReadBitValidationRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestBeforeReadBitValidation::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestBeforeReadBitValidationWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestBeforeReadBitValidation::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestBeforeReadBitValidationReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestBeforeReadBitValidation::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestBeforeReadBitValidationWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestBeforeReadBitValidationReset(::tl2::cases::TestBeforeReadBitValidation& item) {
	item.n = 0;
	item.a.clear();
	item.b.clear();
}

bool tl2::details::CasesTestBeforeReadBitValidationWriteJSON(std::ostream& s, const ::tl2::cases::TestBeforeReadBitValidation& item) {
	auto add_comma = false;
	s << "{";
	if (item.n != 0) {
		add_comma = true;
		s << "\"n\":";
		s << item.n;
	}
	if ((item.n & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.a, item.n)) { return false; }
	}
	if ((item.n & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"b\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.b, item.n)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestBeforeReadBitValidationRead(::basictl::tl_istream & s, ::tl2::cases::TestBeforeReadBitValidation& item) {
	if (!s.nat_read(item.n)) { return false; }
	if ((item.n & (1<<0)) != 0) {
		if (!::tl2::details::BuiltinTupleIntRead(s, item.a, item.n)) { return false; }
	} else {
			item.a.clear();
	}
	if ((item.n & (1<<1)) != 0) {
		if (!::tl2::details::BuiltinTupleIntRead(s, item.b, item.n)) { return false; }
	} else {
			item.b.clear();
	}
	return true;
}

bool tl2::details::CasesTestBeforeReadBitValidationWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestBeforeReadBitValidation& item) {
	if (!s.nat_write(item.n)) { return false;}
	if ((item.n & (1<<0)) != 0) {
			if (!::tl2::details::BuiltinTupleIntWrite(s, item.a, item.n)) { return false; }
	}
	if ((item.n & (1<<1)) != 0) {
			if (!::tl2::details::BuiltinTupleIntWrite(s, item.b, item.n)) { return false; }
	}
	return true;
}

bool tl2::details::CasesTestBeforeReadBitValidationReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestBeforeReadBitValidation& item) {
	if (!s.nat_read_exact_tag(0x9b2396db)) { return false; }
	return tl2::details::CasesTestBeforeReadBitValidationRead(s, item);
}

bool tl2::details::CasesTestBeforeReadBitValidationWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestBeforeReadBitValidation& item) {
	if (!s.nat_write(0x9b2396db)) { return false; }
	return tl2::details::CasesTestBeforeReadBitValidationWrite(s, item);
}

bool tl2::cases::TestDictAny::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestDictAnyWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestDictAny::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestDictAnyRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestDictAny::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestDictAnyWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestDictAny::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestDictAnyReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestDictAny::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestDictAnyWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestDictAnyReset(::tl2::cases::TestDictAny& item) {
	::tl2::details::DictionaryAnyDoubleIntReset(item.dict);
}

bool tl2::details::CasesTestDictAnyWriteJSON(std::ostream& s, const ::tl2::cases::TestDictAny& item) {
	s << "{";
	if (item.dict.size() != 0) {
		s << "\"dict\":";
		if (!::tl2::details::DictionaryAnyDoubleIntWriteJSON(s, item.dict)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestDictAnyRead(::basictl::tl_istream & s, ::tl2::cases::TestDictAny& item) {
	if (!::tl2::details::DictionaryAnyDoubleIntRead(s, item.dict)) { return false; }
	return true;
}

bool tl2::details::CasesTestDictAnyWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestDictAny& item) {
	if (!::tl2::details::DictionaryAnyDoubleIntWrite(s, item.dict)) { return false; }
	return true;
}

bool tl2::details::CasesTestDictAnyReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestDictAny& item) {
	if (!s.nat_read_exact_tag(0xe29b8ae6)) { return false; }
	return tl2::details::CasesTestDictAnyRead(s, item);
}

bool tl2::details::CasesTestDictAnyWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestDictAny& item) {
	if (!s.nat_write(0xe29b8ae6)) { return false; }
	return tl2::details::CasesTestDictAnyWrite(s, item);
}

bool tl2::cases::TestDictInt::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestDictIntWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestDictInt::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestDictIntRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestDictInt::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestDictIntWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestDictInt::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestDictIntReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestDictInt::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestDictIntWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestDictIntReset(::tl2::cases::TestDictInt& item) {
	item.dict.clear();
}

bool tl2::details::CasesTestDictIntWriteJSON(std::ostream& s, const ::tl2::cases::TestDictInt& item) {
	s << "{";
	if (item.dict.size() != 0) {
		s << "\"dict\":";
		if (!::tl2::details::BuiltinVectorDictionaryFieldAnyIntIntWriteJSON(s, item.dict)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestDictIntRead(::basictl::tl_istream & s, ::tl2::cases::TestDictInt& item) {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyIntIntRead(s, item.dict)) { return false; }
	return true;
}

bool tl2::details::CasesTestDictIntWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestDictInt& item) {
	if (!::tl2::details::BuiltinVectorDictionaryFieldAnyIntIntWrite(s, item.dict)) { return false; }
	return true;
}

bool tl2::details::CasesTestDictIntReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestDictInt& item) {
	if (!s.nat_read_exact_tag(0xd3877643)) { return false; }
	return tl2::details::CasesTestDictIntRead(s, item);
}

bool tl2::details::CasesTestDictIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestDictInt& item) {
	if (!s.nat_write(0xd3877643)) { return false; }
	return tl2::details::CasesTestDictIntWrite(s, item);
}

bool tl2::cases::TestDictString::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestDictStringWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestDictString::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestDictStringRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestDictString::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestDictStringWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestDictString::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestDictStringReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestDictString::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestDictStringWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestDictStringReset(::tl2::cases::TestDictString& item) {
	::tl2::details::DictionaryIntReset(item.dict);
}

bool tl2::details::CasesTestDictStringWriteJSON(std::ostream& s, const ::tl2::cases::TestDictString& item) {
	s << "{";
	if (item.dict.size() != 0) {
		s << "\"dict\":";
		if (!::tl2::details::DictionaryIntWriteJSON(s, item.dict)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestDictStringRead(::basictl::tl_istream & s, ::tl2::cases::TestDictString& item) {
	if (!::tl2::details::DictionaryIntRead(s, item.dict)) { return false; }
	return true;
}

bool tl2::details::CasesTestDictStringWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestDictString& item) {
	if (!::tl2::details::DictionaryIntWrite(s, item.dict)) { return false; }
	return true;
}

bool tl2::details::CasesTestDictStringReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestDictString& item) {
	if (!s.nat_read_exact_tag(0xc463c79b)) { return false; }
	return tl2::details::CasesTestDictStringRead(s, item);
}

bool tl2::details::CasesTestDictStringWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestDictString& item) {
	if (!s.nat_write(0xc463c79b)) { return false; }
	return tl2::details::CasesTestDictStringWrite(s, item);
}

static const std::string_view CasesTestEnum_tbl_tl_name[]{"cases.testEnum1", "cases.testEnum2", "cases.testEnum3"};
static const uint32_t CasesTestEnum_tbl_tl_tag[]{0x6c6c55ac, 0x86ea88ce, 0x69b83e2f};

bool tl2::cases::TestEnum::write_json(std::ostream & s)const {
	if (!::tl2::details::CasesTestEnumWriteJSON(s, *this)) { return false; }
	return true;
}
bool tl2::cases::TestEnum::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestEnumReadBoxed(s, *this)) { return false; }
	return true;
}
bool tl2::cases::TestEnum::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestEnumWriteBoxed(s, *this)) { return false; }
	return true;
}
std::string_view tl2::cases::TestEnum::tl_name() const {
	return CasesTestEnum_tbl_tl_name[value.index()];
}
uint32_t tl2::cases::TestEnum::tl_tag() const {
	return CasesTestEnum_tbl_tl_tag[value.index()];
}


void tl2::details::CasesTestEnumReset(::tl2::cases::TestEnum& item) {
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::CasesTestEnumWriteJSON(std::ostream & s, const ::tl2::cases::TestEnum& item) {
	s << "\"" << CasesTestEnum_tbl_tl_name[item.value.index()] << "\"";
	return true;
}
bool tl2::details::CasesTestEnumReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestEnum& item) {
	uint32_t nat;
	s.nat_read(nat);
	switch (nat) {
	case 0x6c6c55ac:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		break;
	case 0x86ea88ce:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		break;
	case 0x69b83e2f:
		if (item.value.index() != 2) { item.value.emplace<2>(); }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::CasesTestEnumWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestEnum& item) {
	s.nat_write(CasesTestEnum_tbl_tl_tag[item.value.index()]);
	switch (item.value.index()) {
	}
	return true;
}

bool tl2::cases::TestEnum1::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestEnum1WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestEnum1::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestEnum1Read(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestEnum1::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestEnum1Write(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestEnum1::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestEnum1ReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestEnum1::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestEnum1WriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestEnum1Reset(::tl2::cases::TestEnum1& item) {
}

bool tl2::details::CasesTestEnum1WriteJSON(std::ostream& s, const ::tl2::cases::TestEnum1& item) {
	s << "true";
	return true;
}

bool tl2::details::CasesTestEnum1Read(::basictl::tl_istream & s, ::tl2::cases::TestEnum1& item) {
	return true;
}

bool tl2::details::CasesTestEnum1Write(::basictl::tl_ostream & s, const ::tl2::cases::TestEnum1& item) {
	return true;
}

bool tl2::details::CasesTestEnum1ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestEnum1& item) {
	if (!s.nat_read_exact_tag(0x6c6c55ac)) { return false; }
	return tl2::details::CasesTestEnum1Read(s, item);
}

bool tl2::details::CasesTestEnum1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestEnum1& item) {
	if (!s.nat_write(0x6c6c55ac)) { return false; }
	return tl2::details::CasesTestEnum1Write(s, item);
}

bool tl2::cases::TestEnum2::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestEnum2WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestEnum2::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestEnum2Read(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestEnum2::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestEnum2Write(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestEnum2::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestEnum2ReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestEnum2::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestEnum2WriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestEnum2Reset(::tl2::cases::TestEnum2& item) {
}

bool tl2::details::CasesTestEnum2WriteJSON(std::ostream& s, const ::tl2::cases::TestEnum2& item) {
	s << "true";
	return true;
}

bool tl2::details::CasesTestEnum2Read(::basictl::tl_istream & s, ::tl2::cases::TestEnum2& item) {
	return true;
}

bool tl2::details::CasesTestEnum2Write(::basictl::tl_ostream & s, const ::tl2::cases::TestEnum2& item) {
	return true;
}

bool tl2::details::CasesTestEnum2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestEnum2& item) {
	if (!s.nat_read_exact_tag(0x86ea88ce)) { return false; }
	return tl2::details::CasesTestEnum2Read(s, item);
}

bool tl2::details::CasesTestEnum2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestEnum2& item) {
	if (!s.nat_write(0x86ea88ce)) { return false; }
	return tl2::details::CasesTestEnum2Write(s, item);
}

bool tl2::cases::TestEnum3::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestEnum3WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestEnum3::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestEnum3Read(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestEnum3::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestEnum3Write(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestEnum3::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestEnum3ReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestEnum3::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestEnum3WriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestEnum3Reset(::tl2::cases::TestEnum3& item) {
}

bool tl2::details::CasesTestEnum3WriteJSON(std::ostream& s, const ::tl2::cases::TestEnum3& item) {
	s << "true";
	return true;
}

bool tl2::details::CasesTestEnum3Read(::basictl::tl_istream & s, ::tl2::cases::TestEnum3& item) {
	return true;
}

bool tl2::details::CasesTestEnum3Write(::basictl::tl_ostream & s, const ::tl2::cases::TestEnum3& item) {
	return true;
}

bool tl2::details::CasesTestEnum3ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestEnum3& item) {
	if (!s.nat_read_exact_tag(0x69b83e2f)) { return false; }
	return tl2::details::CasesTestEnum3Read(s, item);
}

bool tl2::details::CasesTestEnum3WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestEnum3& item) {
	if (!s.nat_write(0x69b83e2f)) { return false; }
	return tl2::details::CasesTestEnum3Write(s, item);
}

bool tl2::cases::TestEnumContainer::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestEnumContainerWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestEnumContainer::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestEnumContainerRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestEnumContainer::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestEnumContainerWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestEnumContainer::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestEnumContainerReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestEnumContainer::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestEnumContainerWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestEnumContainerReset(::tl2::cases::TestEnumContainer& item) {
	::tl2::details::CasesTestEnumReset(item.value);
}

bool tl2::details::CasesTestEnumContainerWriteJSON(std::ostream& s, const ::tl2::cases::TestEnumContainer& item) {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::CasesTestEnumWriteJSON(s, item.value)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesTestEnumContainerRead(::basictl::tl_istream & s, ::tl2::cases::TestEnumContainer& item) {
	if (!::tl2::details::CasesTestEnumReadBoxed(s, item.value)) { return false; }
	return true;
}

bool tl2::details::CasesTestEnumContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestEnumContainer& item) {
	if (!::tl2::details::CasesTestEnumWriteBoxed(s, item.value)) { return false; }
	return true;
}

bool tl2::details::CasesTestEnumContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestEnumContainer& item) {
	if (!s.nat_read_exact_tag(0xcb684231)) { return false; }
	return tl2::details::CasesTestEnumContainerRead(s, item);
}

bool tl2::details::CasesTestEnumContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestEnumContainer& item) {
	if (!s.nat_write(0xcb684231)) { return false; }
	return tl2::details::CasesTestEnumContainerWrite(s, item);
}

bool tl2::cases::TestInplaceStructArgs::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestInplaceStructArgsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestInplaceStructArgs::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestInplaceStructArgsRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestInplaceStructArgs::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestInplaceStructArgsWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestInplaceStructArgs::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestInplaceStructArgsReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestInplaceStructArgs::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestInplaceStructArgsWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestInplaceStructArgsReset(::tl2::cases::TestInplaceStructArgs& item) {
	item.a1 = 0;
	item.a2 = 0;
	item.a3 = 0;
	::tl2::details::CasesInplace1IntReset(item.arg);
}

bool tl2::details::CasesTestInplaceStructArgsWriteJSON(std::ostream& s, const ::tl2::cases::TestInplaceStructArgs& item) {
	auto add_comma = false;
	s << "{";
	if (item.a1 != 0) {
		add_comma = true;
		s << "\"a1\":";
		s << item.a1;
	}
	if (item.a2 != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a2\":";
		s << item.a2;
	}
	if (item.a3 != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a3\":";
		s << item.a3;
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"arg\":";
	if (!::tl2::details::CasesInplace1IntWriteJSON(s, item.arg, item.a1, item.a2, item.a3)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesTestInplaceStructArgsRead(::basictl::tl_istream & s, ::tl2::cases::TestInplaceStructArgs& item) {
	if (!s.nat_read(item.a1)) { return false; }
	if (!s.nat_read(item.a2)) { return false; }
	if (!s.nat_read(item.a3)) { return false; }
	if (!::tl2::details::CasesInplace1IntRead(s, item.arg, item.a1, item.a2, item.a3)) { return false; }
	return true;
}

bool tl2::details::CasesTestInplaceStructArgsWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestInplaceStructArgs& item) {
	if (!s.nat_write(item.a1)) { return false;}
	if (!s.nat_write(item.a2)) { return false;}
	if (!s.nat_write(item.a3)) { return false;}
	if (!::tl2::details::CasesInplace1IntWrite(s, item.arg, item.a1, item.a2, item.a3)) { return false; }
	return true;
}

bool tl2::details::CasesTestInplaceStructArgsReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestInplaceStructArgs& item) {
	if (!s.nat_read_exact_tag(0xa9e4441e)) { return false; }
	return tl2::details::CasesTestInplaceStructArgsRead(s, item);
}

bool tl2::details::CasesTestInplaceStructArgsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestInplaceStructArgs& item) {
	if (!s.nat_write(0xa9e4441e)) { return false; }
	return tl2::details::CasesTestInplaceStructArgsWrite(s, item);
}

bool tl2::cases::TestInplaceStructArgs2::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestInplaceStructArgs2WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestInplaceStructArgs2::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestInplaceStructArgs2Read(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestInplaceStructArgs2::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestInplaceStructArgs2Write(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestInplaceStructArgs2::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestInplaceStructArgs2ReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestInplaceStructArgs2::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestInplaceStructArgs2WriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestInplaceStructArgs2Reset(::tl2::cases::TestInplaceStructArgs2& item) {
	item.a1 = 0;
	item.a2 = 0;
	item.a3 = 0;
	::tl2::details::CasesInplace1PairTupleIntTupleIntReset(item.arg);
}

bool tl2::details::CasesTestInplaceStructArgs2WriteJSON(std::ostream& s, const ::tl2::cases::TestInplaceStructArgs2& item) {
	auto add_comma = false;
	s << "{";
	if (item.a1 != 0) {
		add_comma = true;
		s << "\"a1\":";
		s << item.a1;
	}
	if (item.a2 != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a2\":";
		s << item.a2;
	}
	if (item.a3 != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a3\":";
		s << item.a3;
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"arg\":";
	if (!::tl2::details::CasesInplace1PairTupleIntTupleIntWriteJSON(s, item.arg, item.a1, item.a2, item.a3, item.a3, item.a2)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesTestInplaceStructArgs2Read(::basictl::tl_istream & s, ::tl2::cases::TestInplaceStructArgs2& item) {
	if (!s.nat_read(item.a1)) { return false; }
	if (!s.nat_read(item.a2)) { return false; }
	if (!s.nat_read(item.a3)) { return false; }
	if (!::tl2::details::CasesInplace1PairTupleIntTupleIntRead(s, item.arg, item.a1, item.a2, item.a3, item.a3, item.a2)) { return false; }
	return true;
}

bool tl2::details::CasesTestInplaceStructArgs2Write(::basictl::tl_ostream & s, const ::tl2::cases::TestInplaceStructArgs2& item) {
	if (!s.nat_write(item.a1)) { return false;}
	if (!s.nat_write(item.a2)) { return false;}
	if (!s.nat_write(item.a3)) { return false;}
	if (!::tl2::details::CasesInplace1PairTupleIntTupleIntWrite(s, item.arg, item.a1, item.a2, item.a3, item.a3, item.a2)) { return false; }
	return true;
}

bool tl2::details::CasesTestInplaceStructArgs2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestInplaceStructArgs2& item) {
	if (!s.nat_read_exact_tag(0xaa9f2480)) { return false; }
	return tl2::details::CasesTestInplaceStructArgs2Read(s, item);
}

bool tl2::details::CasesTestInplaceStructArgs2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestInplaceStructArgs2& item) {
	if (!s.nat_write(0xaa9f2480)) { return false; }
	return tl2::details::CasesTestInplaceStructArgs2Write(s, item);
}

bool tl2::cases::TestLocalFieldmask::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestLocalFieldmaskWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestLocalFieldmask::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestLocalFieldmaskRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestLocalFieldmask::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestLocalFieldmaskWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestLocalFieldmask::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestLocalFieldmaskReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestLocalFieldmask::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestLocalFieldmaskWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestLocalFieldmaskReset(::tl2::cases::TestLocalFieldmask& item) {
	item.f1 = 0;
	item.f2 = 0;
	::tl2::details::TrueReset(item.f3);
	::tl2::details::TrueReset(item.f4);
}

bool tl2::details::CasesTestLocalFieldmaskWriteJSON(std::ostream& s, const ::tl2::cases::TestLocalFieldmask& item) {
	auto add_comma = false;
	s << "{";
	if (item.f1 != 0) {
		add_comma = true;
		s << "\"f1\":";
		s << item.f1;
	}
	if ((item.f1 & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f2\":";
		s << item.f2;
	}
	if ((item.f2 & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f3\":";
		if (!::tl2::details::TrueWriteJSON(s, item.f3)) { return false; }
	}
	if ((item.f2 & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f4\":";
		if (!::tl2::details::TrueWriteJSON(s, item.f4)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestLocalFieldmaskRead(::basictl::tl_istream & s, ::tl2::cases::TestLocalFieldmask& item) {
	if (!s.nat_read(item.f1)) { return false; }
	if ((item.f1 & (1<<0)) != 0) {
		if (!s.nat_read(item.f2)) { return false; }
	} else {
			item.f2 = 0;
	}
	if ((item.f2 & (1<<1)) != 0) {
		if (!::tl2::details::TrueRead(s, item.f3)) { return false; }
	} else {
			::tl2::details::TrueReset(item.f3);
	}
	if ((item.f2 & (1<<1)) != 0) {
		if (!::tl2::details::TrueRead(s, item.f4)) { return false; }
	} else {
			::tl2::details::TrueReset(item.f4);
	}
	return true;
}

bool tl2::details::CasesTestLocalFieldmaskWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestLocalFieldmask& item) {
	if (!s.nat_write(item.f1)) { return false;}
	if ((item.f1 & (1<<0)) != 0) {
			if (!s.nat_write(item.f2)) { return false;}
	}
	if ((item.f2 & (1<<1)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.f3)) { return false; }
	}
	if ((item.f2 & (1<<1)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.f4)) { return false; }
	}
	return true;
}

bool tl2::details::CasesTestLocalFieldmaskReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestLocalFieldmask& item) {
	if (!s.nat_read_exact_tag(0xf68fd3f9)) { return false; }
	return tl2::details::CasesTestLocalFieldmaskRead(s, item);
}

bool tl2::details::CasesTestLocalFieldmaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestLocalFieldmask& item) {
	if (!s.nat_write(0xf68fd3f9)) { return false; }
	return tl2::details::CasesTestLocalFieldmaskWrite(s, item);
}

bool tl2::cases::TestMaybe::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestMaybeWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestMaybe::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestMaybeRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestMaybe::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestMaybeWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestMaybe::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestMaybeReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestMaybe::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestMaybeWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestMaybeReset(::tl2::cases::TestMaybe& item) {
	item.value.reset();
}

bool tl2::details::CasesTestMaybeWriteJSON(std::ostream& s, const ::tl2::cases::TestMaybe& item) {
	s << "{";
	if (item.value.has_value()) {
		s << "\"value\":";
		if (!::tl2::details::IntMaybeWriteJSON(s, item.value)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestMaybeRead(::basictl::tl_istream & s, ::tl2::cases::TestMaybe& item) {
	if (!::tl2::details::IntMaybeReadBoxed(s, item.value)) { return false; }
	return true;
}

bool tl2::details::CasesTestMaybeWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestMaybe& item) {
	if (!::tl2::details::IntMaybeWriteBoxed(s, item.value)) { return false; }
	return true;
}

bool tl2::details::CasesTestMaybeReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestMaybe& item) {
	if (!s.nat_read_exact_tag(0xd6602613)) { return false; }
	return tl2::details::CasesTestMaybeRead(s, item);
}

bool tl2::details::CasesTestMaybeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestMaybe& item) {
	if (!s.nat_write(0xd6602613)) { return false; }
	return tl2::details::CasesTestMaybeWrite(s, item);
}

bool tl2::cases::TestOutFieldMask::write_json(std::ostream& s, uint32_t nat_f)const {
	if (!::tl2::details::CasesTestOutFieldMaskWriteJSON(s, *this, nat_f)) { return false; }
	return true;
}

bool tl2::cases::TestOutFieldMask::read(::basictl::tl_istream & s, uint32_t nat_f) {
	if (!::tl2::details::CasesTestOutFieldMaskRead(s, *this, nat_f)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestOutFieldMask::write(::basictl::tl_ostream & s, uint32_t nat_f)const {
	if (!::tl2::details::CasesTestOutFieldMaskWrite(s, *this, nat_f)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestOutFieldMask::read_boxed(::basictl::tl_istream & s, uint32_t nat_f) {
	if (!::tl2::details::CasesTestOutFieldMaskReadBoxed(s, *this, nat_f)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestOutFieldMask::write_boxed(::basictl::tl_ostream & s, uint32_t nat_f)const {
	if (!::tl2::details::CasesTestOutFieldMaskWriteBoxed(s, *this, nat_f)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestOutFieldMaskReset(::tl2::cases::TestOutFieldMask& item) {
	item.f1 = 0;
	::tl2::details::TrueReset(item.f2);
	item.f3.clear();
}

bool tl2::details::CasesTestOutFieldMaskWriteJSON(std::ostream& s, const ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f) {
	auto add_comma = false;
	s << "{";
	if ((nat_f & (1<<0)) != 0) {
		add_comma = true;
		s << "\"f1\":";
		s << item.f1;
	}
	if ((item.f3.size() != 0) || (nat_f != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f3\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.f3, nat_f)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestOutFieldMaskRead(::basictl::tl_istream & s, ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f) {
	if ((nat_f & (1<<0)) != 0) {
		if (!s.nat_read(item.f1)) { return false; }
	} else {
			item.f1 = 0;
	}
	if ((nat_f & (1<<3)) != 0) {
		if (!::tl2::details::TrueRead(s, item.f2)) { return false; }
	} else {
			::tl2::details::TrueReset(item.f2);
	}
	if (!::tl2::details::BuiltinTupleIntRead(s, item.f3, nat_f)) { return false; }
	return true;
}

bool tl2::details::CasesTestOutFieldMaskWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f) {
	if ((nat_f & (1<<0)) != 0) {
			if (!s.nat_write(item.f1)) { return false;}
	}
	if ((nat_f & (1<<3)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.f2)) { return false; }
	}
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.f3, nat_f)) { return false; }
	return true;
}

bool tl2::details::CasesTestOutFieldMaskReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f) {
	if (!s.nat_read_exact_tag(0xbd6b4b3c)) { return false; }
	return tl2::details::CasesTestOutFieldMaskRead(s, item, nat_f);
}

bool tl2::details::CasesTestOutFieldMaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestOutFieldMask& item, uint32_t nat_f) {
	if (!s.nat_write(0xbd6b4b3c)) { return false; }
	return tl2::details::CasesTestOutFieldMaskWrite(s, item, nat_f);
}

bool tl2::cases::TestOutFieldMaskContainer::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestOutFieldMaskContainerWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestOutFieldMaskContainer::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestOutFieldMaskContainerRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestOutFieldMaskContainer::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestOutFieldMaskContainerWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestOutFieldMaskContainer::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestOutFieldMaskContainerReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestOutFieldMaskContainer::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestOutFieldMaskContainerWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestOutFieldMaskContainerReset(::tl2::cases::TestOutFieldMaskContainer& item) {
	item.f = 0;
	::tl2::details::CasesTestOutFieldMaskReset(item.inner);
}

bool tl2::details::CasesTestOutFieldMaskContainerWriteJSON(std::ostream& s, const ::tl2::cases::TestOutFieldMaskContainer& item) {
	auto add_comma = false;
	s << "{";
	if (item.f != 0) {
		add_comma = true;
		s << "\"f\":";
		s << item.f;
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"inner\":";
	if (!::tl2::details::CasesTestOutFieldMaskWriteJSON(s, item.inner, item.f)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesTestOutFieldMaskContainerRead(::basictl::tl_istream & s, ::tl2::cases::TestOutFieldMaskContainer& item) {
	if (!s.nat_read(item.f)) { return false; }
	if (!::tl2::details::CasesTestOutFieldMaskRead(s, item.inner, item.f)) { return false; }
	return true;
}

bool tl2::details::CasesTestOutFieldMaskContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestOutFieldMaskContainer& item) {
	if (!s.nat_write(item.f)) { return false;}
	if (!::tl2::details::CasesTestOutFieldMaskWrite(s, item.inner, item.f)) { return false; }
	return true;
}

bool tl2::details::CasesTestOutFieldMaskContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestOutFieldMaskContainer& item) {
	if (!s.nat_read_exact_tag(0x1850ffe4)) { return false; }
	return tl2::details::CasesTestOutFieldMaskContainerRead(s, item);
}

bool tl2::details::CasesTestOutFieldMaskContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestOutFieldMaskContainer& item) {
	if (!s.nat_write(0x1850ffe4)) { return false; }
	return tl2::details::CasesTestOutFieldMaskContainerWrite(s, item);
}

bool tl2::cases::TestRecursiveFieldMask::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestRecursiveFieldmaskWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestRecursiveFieldMask::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestRecursiveFieldmaskRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestRecursiveFieldMask::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestRecursiveFieldmaskWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestRecursiveFieldMask::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestRecursiveFieldmaskReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestRecursiveFieldMask::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestRecursiveFieldmaskWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestRecursiveFieldmaskReset(::tl2::cases::TestRecursiveFieldMask& item) {
	item.f0 = 0;
	item.f1 = 0;
	item.f2 = 0;
	::tl2::details::TrueReset(item.t1);
	::tl2::details::TrueReset(item.t2);
	::tl2::details::TrueReset(item.t3);
}

bool tl2::details::CasesTestRecursiveFieldmaskWriteJSON(std::ostream& s, const ::tl2::cases::TestRecursiveFieldMask& item) {
	auto add_comma = false;
	s << "{";
	if (item.f0 != 0) {
		add_comma = true;
		s << "\"f0\":";
		s << item.f0;
	}
	if ((item.f0 & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f1\":";
		s << item.f1;
	}
	if ((item.f1 & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"f2\":";
		s << item.f2;
	}
	if ((item.f0 & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"t1\":";
		if (!::tl2::details::TrueWriteJSON(s, item.t1)) { return false; }
	}
	if ((item.f1 & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"t2\":";
		if (!::tl2::details::TrueWriteJSON(s, item.t2)) { return false; }
	}
	if ((item.f2 & (1<<2)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"t3\":";
		if (!::tl2::details::TrueWriteJSON(s, item.t3)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestRecursiveFieldmaskRead(::basictl::tl_istream & s, ::tl2::cases::TestRecursiveFieldMask& item) {
	if (!s.nat_read(item.f0)) { return false; }
	if ((item.f0 & (1<<0)) != 0) {
		if (!s.nat_read(item.f1)) { return false; }
	} else {
			item.f1 = 0;
	}
	if ((item.f1 & (1<<1)) != 0) {
		if (!s.nat_read(item.f2)) { return false; }
	} else {
			item.f2 = 0;
	}
	if ((item.f0 & (1<<0)) != 0) {
		if (!::tl2::details::TrueRead(s, item.t1)) { return false; }
	} else {
			::tl2::details::TrueReset(item.t1);
	}
	if ((item.f1 & (1<<1)) != 0) {
		if (!::tl2::details::TrueRead(s, item.t2)) { return false; }
	} else {
			::tl2::details::TrueReset(item.t2);
	}
	if ((item.f2 & (1<<2)) != 0) {
		if (!::tl2::details::TrueRead(s, item.t3)) { return false; }
	} else {
			::tl2::details::TrueReset(item.t3);
	}
	return true;
}

bool tl2::details::CasesTestRecursiveFieldmaskWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestRecursiveFieldMask& item) {
	if (!s.nat_write(item.f0)) { return false;}
	if ((item.f0 & (1<<0)) != 0) {
			if (!s.nat_write(item.f1)) { return false;}
	}
	if ((item.f1 & (1<<1)) != 0) {
			if (!s.nat_write(item.f2)) { return false;}
	}
	if ((item.f0 & (1<<0)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.t1)) { return false; }
	}
	if ((item.f1 & (1<<1)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.t2)) { return false; }
	}
	if ((item.f2 & (1<<2)) != 0) {
			if (!::tl2::details::TrueWrite(s, item.t3)) { return false; }
	}
	return true;
}

bool tl2::details::CasesTestRecursiveFieldmaskReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestRecursiveFieldMask& item) {
	if (!s.nat_read_exact_tag(0xc58cf85e)) { return false; }
	return tl2::details::CasesTestRecursiveFieldmaskRead(s, item);
}

bool tl2::details::CasesTestRecursiveFieldmaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestRecursiveFieldMask& item) {
	if (!s.nat_write(0xc58cf85e)) { return false; }
	return tl2::details::CasesTestRecursiveFieldmaskWrite(s, item);
}

bool tl2::cases::TestTuple::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestTupleWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestTuple::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestTupleRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestTuple::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestTupleWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestTuple::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestTupleReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestTuple::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestTupleWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestTupleReset(::tl2::cases::TestTuple& item) {
	::tl2::details::BuiltinTuple4IntReset(item.tpl);
}

bool tl2::details::CasesTestTupleWriteJSON(std::ostream& s, const ::tl2::cases::TestTuple& item) {
	s << "{";
	s << "\"tpl\":";
	if (!::tl2::details::BuiltinTuple4IntWriteJSON(s, item.tpl)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesTestTupleRead(::basictl::tl_istream & s, ::tl2::cases::TestTuple& item) {
	if (!::tl2::details::BuiltinTuple4IntRead(s, item.tpl)) { return false; }
	return true;
}

bool tl2::details::CasesTestTupleWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestTuple& item) {
	if (!::tl2::details::BuiltinTuple4IntWrite(s, item.tpl)) { return false; }
	return true;
}

bool tl2::details::CasesTestTupleReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestTuple& item) {
	if (!s.nat_read_exact_tag(0x4b9caf8f)) { return false; }
	return tl2::details::CasesTestTupleRead(s, item);
}

bool tl2::details::CasesTestTupleWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestTuple& item) {
	if (!s.nat_write(0x4b9caf8f)) { return false; }
	return tl2::details::CasesTestTupleWrite(s, item);
}

static const std::string_view CasesTestUnion_tbl_tl_name[]{"cases.testUnion1", "cases.testUnion2"};
static const uint32_t CasesTestUnion_tbl_tl_tag[]{0x4b4f09b1, 0x464f96c4};

bool tl2::cases::TestUnion::write_json(std::ostream & s)const {
	if (!::tl2::details::CasesTestUnionWriteJSON(s, *this)) { return false; }
	return true;
}
bool tl2::cases::TestUnion::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestUnionReadBoxed(s, *this)) { return false; }
	return true;
}
bool tl2::cases::TestUnion::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestUnionWriteBoxed(s, *this)) { return false; }
	return true;
}
std::string_view tl2::cases::TestUnion::tl_name() const {
	return CasesTestUnion_tbl_tl_name[value.index()];
}
uint32_t tl2::cases::TestUnion::tl_tag() const {
	return CasesTestUnion_tbl_tl_tag[value.index()];
}


void tl2::details::CasesTestUnionReset(::tl2::cases::TestUnion& item) {
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::CasesTestUnionWriteJSON(std::ostream & s, const ::tl2::cases::TestUnion& item) {
	s << "{";
	s << "\"type\":";
	s << "\"" << CasesTestUnion_tbl_tl_name[item.value.index()] << "\"";
	switch (item.value.index()) {
	case 0:
		s << ",\"value\":";
		if (!::tl2::details::CasesTestUnion1WriteJSON(s, std::get<0>(item.value))) { return false; }
		break;
	case 1:
		s << ",\"value\":";
		if (!::tl2::details::CasesTestUnion2WriteJSON(s, std::get<1>(item.value))) { return false; }
		break;
	}
	s << "}";
	return true;
}
bool tl2::details::CasesTestUnionReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestUnion& item) {
	uint32_t nat;
	s.nat_read(nat);
	switch (nat) {
	case 0x4b4f09b1:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		if (!::tl2::details::CasesTestUnion1Read(s, std::get<0>(item.value))) { return false; }
		break;
	case 0x464f96c4:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		if (!::tl2::details::CasesTestUnion2Read(s, std::get<1>(item.value))) { return false; }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::CasesTestUnionWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion& item) {
	s.nat_write(CasesTestUnion_tbl_tl_tag[item.value.index()]);
	switch (item.value.index()) {
	case 0:
		if (!::tl2::details::CasesTestUnion1Write(s, std::get<0>(item.value))) { return false; }
		break;
	case 1:
		if (!::tl2::details::CasesTestUnion2Write(s, std::get<1>(item.value))) { return false; }
		break;
	}
	return true;
}

bool tl2::cases::TestUnion1::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestUnion1WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestUnion1::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestUnion1Read(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestUnion1::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestUnion1Write(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestUnion1::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestUnion1ReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestUnion1::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestUnion1WriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestUnion1Reset(::tl2::cases::TestUnion1& item) {
	item.value = 0;
}

bool tl2::details::CasesTestUnion1WriteJSON(std::ostream& s, const ::tl2::cases::TestUnion1& item) {
	s << "{";
	if (item.value != 0) {
		s << "\"value\":";
		s << item.value;
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestUnion1Read(::basictl::tl_istream & s, ::tl2::cases::TestUnion1& item) {
	if (!s.int_read(item.value)) { return false; }
	return true;
}

bool tl2::details::CasesTestUnion1Write(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion1& item) {
	if (!s.int_write(item.value)) { return false;}
	return true;
}

bool tl2::details::CasesTestUnion1ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestUnion1& item) {
	if (!s.nat_read_exact_tag(0x4b4f09b1)) { return false; }
	return tl2::details::CasesTestUnion1Read(s, item);
}

bool tl2::details::CasesTestUnion1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion1& item) {
	if (!s.nat_write(0x4b4f09b1)) { return false; }
	return tl2::details::CasesTestUnion1Write(s, item);
}

bool tl2::cases::TestUnion2::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestUnion2WriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestUnion2::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestUnion2Read(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestUnion2::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestUnion2Write(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestUnion2::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestUnion2ReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestUnion2::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestUnion2WriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestUnion2Reset(::tl2::cases::TestUnion2& item) {
	item.value.clear();
}

bool tl2::details::CasesTestUnion2WriteJSON(std::ostream& s, const ::tl2::cases::TestUnion2& item) {
	s << "{";
	if (item.value.size() != 0) {
		s << "\"value\":";
		s << "\"" << item.value << "\"";
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestUnion2Read(::basictl::tl_istream & s, ::tl2::cases::TestUnion2& item) {
	if (!s.string_read(item.value)) { return false; }
	return true;
}

bool tl2::details::CasesTestUnion2Write(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion2& item) {
	if (!s.string_write(item.value)) { return false;}
	return true;
}

bool tl2::details::CasesTestUnion2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestUnion2& item) {
	if (!s.nat_read_exact_tag(0x464f96c4)) { return false; }
	return tl2::details::CasesTestUnion2Read(s, item);
}

bool tl2::details::CasesTestUnion2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestUnion2& item) {
	if (!s.nat_write(0x464f96c4)) { return false; }
	return tl2::details::CasesTestUnion2Write(s, item);
}

bool tl2::cases::TestUnionContainer::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestUnionContainerWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestUnionContainer::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestUnionContainerRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestUnionContainer::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestUnionContainerWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestUnionContainer::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestUnionContainerReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestUnionContainer::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestUnionContainerWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestUnionContainerReset(::tl2::cases::TestUnionContainer& item) {
	::tl2::details::CasesTestUnionReset(item.value);
}

bool tl2::details::CasesTestUnionContainerWriteJSON(std::ostream& s, const ::tl2::cases::TestUnionContainer& item) {
	s << "{";
	s << "\"value\":";
	if (!::tl2::details::CasesTestUnionWriteJSON(s, item.value)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::CasesTestUnionContainerRead(::basictl::tl_istream & s, ::tl2::cases::TestUnionContainer& item) {
	if (!::tl2::details::CasesTestUnionReadBoxed(s, item.value)) { return false; }
	return true;
}

bool tl2::details::CasesTestUnionContainerWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestUnionContainer& item) {
	if (!::tl2::details::CasesTestUnionWriteBoxed(s, item.value)) { return false; }
	return true;
}

bool tl2::details::CasesTestUnionContainerReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestUnionContainer& item) {
	if (!s.nat_read_exact_tag(0x4497a381)) { return false; }
	return tl2::details::CasesTestUnionContainerRead(s, item);
}

bool tl2::details::CasesTestUnionContainerWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestUnionContainer& item) {
	if (!s.nat_write(0x4497a381)) { return false; }
	return tl2::details::CasesTestUnionContainerWrite(s, item);
}

bool tl2::cases::TestVector::write_json(std::ostream& s)const {
	if (!::tl2::details::CasesTestVectorWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::cases::TestVector::read(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestVectorRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestVector::write(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestVectorWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestVector::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::CasesTestVectorReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::cases::TestVector::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::CasesTestVectorWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::details::CasesTestVectorReset(::tl2::cases::TestVector& item) {
	item.arr.clear();
}

bool tl2::details::CasesTestVectorWriteJSON(std::ostream& s, const ::tl2::cases::TestVector& item) {
	s << "{";
	if (item.arr.size() != 0) {
		s << "\"arr\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.arr)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::CasesTestVectorRead(::basictl::tl_istream & s, ::tl2::cases::TestVector& item) {
	if (!::tl2::details::BuiltinVectorIntRead(s, item.arr)) { return false; }
	return true;
}

bool tl2::details::CasesTestVectorWrite(::basictl::tl_ostream & s, const ::tl2::cases::TestVector& item) {
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.arr)) { return false; }
	return true;
}

bool tl2::details::CasesTestVectorReadBoxed(::basictl::tl_istream & s, ::tl2::cases::TestVector& item) {
	if (!s.nat_read_exact_tag(0x4975695c)) { return false; }
	return tl2::details::CasesTestVectorRead(s, item);
}

bool tl2::details::CasesTestVectorWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::TestVector& item) {
	if (!s.nat_write(0x4975695c)) { return false; }
	return tl2::details::CasesTestVectorWrite(s, item);
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

void tl2::details::PairTupleIntTupleIntReset(::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>& item) {
	item.x.clear();
	item.y.clear();
}

bool tl2::details::PairTupleIntTupleIntWriteJSON(std::ostream& s, const ::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>& item, uint32_t nat_X, uint32_t nat_Y) {
	auto add_comma = false;
	s << "{";
	if ((item.x.size() != 0) || (nat_X != 0)) {
		add_comma = true;
		s << "\"x\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.x, nat_X)) { return false; }
	}
	if ((item.y.size() != 0) || (nat_Y != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"y\":";
		if (!::tl2::details::BuiltinTupleIntWriteJSON(s, item.y, nat_Y)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::PairTupleIntTupleIntRead(::basictl::tl_istream & s, ::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>& item, uint32_t nat_X, uint32_t nat_Y) {
	if (!::tl2::details::BuiltinTupleIntRead(s, item.x, nat_X)) { return false; }
	if (!::tl2::details::BuiltinTupleIntRead(s, item.y, nat_Y)) { return false; }
	return true;
}

bool tl2::details::PairTupleIntTupleIntWrite(::basictl::tl_ostream & s, const ::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>& item, uint32_t nat_X, uint32_t nat_Y) {
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.x, nat_X)) { return false; }
	if (!::tl2::details::BuiltinTupleIntWrite(s, item.y, nat_Y)) { return false; }
	return true;
}

bool tl2::details::PairTupleIntTupleIntReadBoxed(::basictl::tl_istream & s, ::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>& item, uint32_t nat_X, uint32_t nat_Y) {
	if (!s.nat_read_exact_tag(0xf01604df)) { return false; }
	return tl2::details::PairTupleIntTupleIntRead(s, item, nat_X, nat_Y);
}

bool tl2::details::PairTupleIntTupleIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>& item, uint32_t nat_X, uint32_t nat_Y) {
	if (!s.nat_write(0xf01604df)) { return false; }
	return tl2::details::PairTupleIntTupleIntWrite(s, item, nat_X, nat_Y);
}

void tl2::details::PairTupleTupleInt2TupleTupleInt2Reset(::tl2::Pair<std::vector<std::array<int32_t, 2>>, std::vector<std::array<int32_t, 2>>>& item) {
	item.x.clear();
	item.y.clear();
}

bool tl2::details::PairTupleTupleInt2TupleTupleInt2WriteJSON(std::ostream& s, const ::tl2::Pair<std::vector<std::array<int32_t, 2>>, std::vector<std::array<int32_t, 2>>>& item, uint32_t nat_X, uint32_t nat_Y) {
	auto add_comma = false;
	s << "{";
	if ((item.x.size() != 0) || (nat_X != 0)) {
		add_comma = true;
		s << "\"x\":";
		if (!::tl2::details::BuiltinTupleTupleInt2WriteJSON(s, item.x, nat_X)) { return false; }
	}
	if ((item.y.size() != 0) || (nat_Y != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"y\":";
		if (!::tl2::details::BuiltinTupleTupleInt2WriteJSON(s, item.y, nat_Y)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::PairTupleTupleInt2TupleTupleInt2Read(::basictl::tl_istream & s, ::tl2::Pair<std::vector<std::array<int32_t, 2>>, std::vector<std::array<int32_t, 2>>>& item, uint32_t nat_X, uint32_t nat_Y) {
	if (!::tl2::details::BuiltinTupleTupleInt2Read(s, item.x, nat_X)) { return false; }
	if (!::tl2::details::BuiltinTupleTupleInt2Read(s, item.y, nat_Y)) { return false; }
	return true;
}

bool tl2::details::PairTupleTupleInt2TupleTupleInt2Write(::basictl::tl_ostream & s, const ::tl2::Pair<std::vector<std::array<int32_t, 2>>, std::vector<std::array<int32_t, 2>>>& item, uint32_t nat_X, uint32_t nat_Y) {
	if (!::tl2::details::BuiltinTupleTupleInt2Write(s, item.x, nat_X)) { return false; }
	if (!::tl2::details::BuiltinTupleTupleInt2Write(s, item.y, nat_Y)) { return false; }
	return true;
}

bool tl2::details::PairTupleTupleInt2TupleTupleInt2ReadBoxed(::basictl::tl_istream & s, ::tl2::Pair<std::vector<std::array<int32_t, 2>>, std::vector<std::array<int32_t, 2>>>& item, uint32_t nat_X, uint32_t nat_Y) {
	if (!s.nat_read_exact_tag(0xf01604df)) { return false; }
	return tl2::details::PairTupleTupleInt2TupleTupleInt2Read(s, item, nat_X, nat_Y);
}

bool tl2::details::PairTupleTupleInt2TupleTupleInt2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::Pair<std::vector<std::array<int32_t, 2>>, std::vector<std::array<int32_t, 2>>>& item, uint32_t nat_X, uint32_t nat_Y) {
	if (!s.nat_write(0xf01604df)) { return false; }
	return tl2::details::PairTupleTupleInt2TupleTupleInt2Write(s, item, nat_X, nat_Y);
}

void tl2::details::PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2Reset(::tl2::Pair<std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>>& item) {
	item.x.clear();
	item.y.clear();
}

bool tl2::details::PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2WriteJSON(std::ostream& s, const ::tl2::Pair<std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>>& item, uint32_t nat_XttXn, uint32_t nat_XttYn, uint32_t nat_Xn, uint32_t nat_YttXn, uint32_t nat_YttYn, uint32_t nat_Yn) {
	auto add_comma = false;
	s << "{";
	if ((item.x.size() != 0) || (nat_XttXn != 0)) {
		add_comma = true;
		s << "\"x\":";
		if (!::tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2WriteJSON(s, item.x, nat_Xn, nat_XttXn, nat_XttYn)) { return false; }
	}
	if ((item.y.size() != 0) || (nat_YttXn != 0)) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"y\":";
		if (!::tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2WriteJSON(s, item.y, nat_Yn, nat_YttXn, nat_YttYn)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2Read(::basictl::tl_istream & s, ::tl2::Pair<std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>>& item, uint32_t nat_XttXn, uint32_t nat_XttYn, uint32_t nat_Xn, uint32_t nat_YttXn, uint32_t nat_YttYn, uint32_t nat_Yn) {
	if (!::tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2Read(s, item.x, nat_Xn, nat_XttXn, nat_XttYn)) { return false; }
	if (!::tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2Read(s, item.y, nat_Yn, nat_YttXn, nat_YttYn)) { return false; }
	return true;
}

bool tl2::details::PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2Write(::basictl::tl_ostream & s, const ::tl2::Pair<std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>>& item, uint32_t nat_XttXn, uint32_t nat_XttYn, uint32_t nat_Xn, uint32_t nat_YttXn, uint32_t nat_YttYn, uint32_t nat_Yn) {
	if (!::tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2Write(s, item.x, nat_Xn, nat_XttXn, nat_XttYn)) { return false; }
	if (!::tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2Write(s, item.y, nat_Yn, nat_YttXn, nat_YttYn)) { return false; }
	return true;
}

bool tl2::details::PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2ReadBoxed(::basictl::tl_istream & s, ::tl2::Pair<std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>>& item, uint32_t nat_XttXn, uint32_t nat_XttYn, uint32_t nat_Xn, uint32_t nat_YttXn, uint32_t nat_YttYn, uint32_t nat_Yn) {
	if (!s.nat_read_exact_tag(0xf01604df)) { return false; }
	return tl2::details::PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2Read(s, item, nat_XttXn, nat_XttYn, nat_Xn, nat_YttXn, nat_YttYn, nat_Yn);
}

bool tl2::details::PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::Pair<std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>>& item, uint32_t nat_XttXn, uint32_t nat_XttYn, uint32_t nat_Xn, uint32_t nat_YttXn, uint32_t nat_YttYn, uint32_t nat_Yn) {
	if (!s.nat_write(0xf01604df)) { return false; }
	return tl2::details::PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2Write(s, item, nat_XttXn, nat_XttYn, nat_Xn, nat_YttXn, nat_YttYn, nat_Yn);
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

void tl2::details::TupleInt2Reset(std::array<int32_t, 2>& item) {
	::tl2::details::BuiltinTuple2IntReset(item);
}

bool tl2::details::TupleInt2WriteJSON(std::ostream& s, const std::array<int32_t, 2>& item) {
	if (!::tl2::details::BuiltinTuple2IntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt2Read(::basictl::tl_istream & s, std::array<int32_t, 2>& item) {
	if (!::tl2::details::BuiltinTuple2IntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt2Write(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item) {
	if (!::tl2::details::BuiltinTuple2IntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt2ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 2>& item) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleInt2Read(s, item);
}

bool tl2::details::TupleInt2WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleInt2Write(s, item);
}

void tl2::details::TupleInt4Reset(std::array<int32_t, 4>& item) {
	::tl2::details::BuiltinTuple4IntReset(item);
}

bool tl2::details::TupleInt4WriteJSON(std::ostream& s, const std::array<int32_t, 4>& item) {
	if (!::tl2::details::BuiltinTuple4IntWriteJSON(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt4Read(::basictl::tl_istream & s, std::array<int32_t, 4>& item) {
	if (!::tl2::details::BuiltinTuple4IntRead(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt4Write(::basictl::tl_ostream & s, const std::array<int32_t, 4>& item) {
	if (!::tl2::details::BuiltinTuple4IntWrite(s, item)) { return false; }
	return true;
}

bool tl2::details::TupleInt4ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 4>& item) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleInt4Read(s, item);
}

bool tl2::details::TupleInt4WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 4>& item) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleInt4Write(s, item);
}

void tl2::details::TuplePairTupleIntTupleInt2Reset(std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item) {
	::tl2::details::BuiltinTuple2PairTupleIntTupleIntReset(item);
}

bool tl2::details::TuplePairTupleIntTupleInt2WriteJSON(std::ostream& s, const std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn) {
	if (!::tl2::details::BuiltinTuple2PairTupleIntTupleIntWriteJSON(s, item, nat_tXn, nat_tYn)) { return false; }
	return true;
}

bool tl2::details::TuplePairTupleIntTupleInt2Read(::basictl::tl_istream & s, std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn) {
	if (!::tl2::details::BuiltinTuple2PairTupleIntTupleIntRead(s, item, nat_tXn, nat_tYn)) { return false; }
	return true;
}

bool tl2::details::TuplePairTupleIntTupleInt2Write(::basictl::tl_ostream & s, const std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn) {
	if (!::tl2::details::BuiltinTuple2PairTupleIntTupleIntWrite(s, item, nat_tXn, nat_tYn)) { return false; }
	return true;
}

bool tl2::details::TuplePairTupleIntTupleInt2ReadBoxed(::basictl::tl_istream & s, std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TuplePairTupleIntTupleInt2Read(s, item, nat_tXn, nat_tYn);
}

bool tl2::details::TuplePairTupleIntTupleInt2WriteBoxed(::basictl::tl_ostream & s, const std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TuplePairTupleIntTupleInt2Write(s, item, nat_tXn, nat_tYn);
}

void tl2::details::TupleTupleInt2Reset(std::vector<std::array<int32_t, 2>>& item) {
	item.clear();
}

bool tl2::details::TupleTupleInt2WriteJSON(std::ostream& s, const std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleTupleInt2WriteJSON(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleTupleInt2Read(::basictl::tl_istream & s, std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleTupleInt2Read(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleTupleInt2Write(::basictl::tl_ostream & s, const std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleTupleInt2Write(s, item, nat_n)) { return false; }
	return true;
}

bool tl2::details::TupleTupleInt2ReadBoxed(::basictl::tl_istream & s, std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleTupleInt2Read(s, item, nat_n);
}

bool tl2::details::TupleTupleInt2WriteBoxed(::basictl::tl_ostream & s, const std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleTupleInt2Write(s, item, nat_n);
}

void tl2::details::TupleTuplePairTupleIntTupleInt2Reset(std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item) {
	item.clear();
}

bool tl2::details::TupleTuplePairTupleIntTupleInt2WriteJSON(std::ostream& s, const std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_ttXn, uint32_t nat_ttYn, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2WriteJSON(s, item, nat_n, nat_ttXn, nat_ttYn)) { return false; }
	return true;
}

bool tl2::details::TupleTuplePairTupleIntTupleInt2Read(::basictl::tl_istream & s, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_ttXn, uint32_t nat_ttYn, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2Read(s, item, nat_n, nat_ttXn, nat_ttYn)) { return false; }
	return true;
}

bool tl2::details::TupleTuplePairTupleIntTupleInt2Write(::basictl::tl_ostream & s, const std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_ttXn, uint32_t nat_ttYn, uint32_t nat_n) {
	if (!::tl2::details::BuiltinTupleTuplePairTupleIntTupleInt2Write(s, item, nat_n, nat_ttXn, nat_ttYn)) { return false; }
	return true;
}

bool tl2::details::TupleTuplePairTupleIntTupleInt2ReadBoxed(::basictl::tl_istream & s, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_ttXn, uint32_t nat_ttYn, uint32_t nat_n) {
	if (!s.nat_read_exact_tag(0x9770768a)) { return false; }
	return tl2::details::TupleTuplePairTupleIntTupleInt2Read(s, item, nat_ttXn, nat_ttYn, nat_n);
}

bool tl2::details::TupleTuplePairTupleIntTupleInt2WriteBoxed(::basictl::tl_ostream & s, const std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_ttXn, uint32_t nat_ttYn, uint32_t nat_n) {
	if (!s.nat_write(0x9770768a)) { return false; }
	return tl2::details::TupleTuplePairTupleIntTupleInt2Write(s, item, nat_ttXn, nat_ttYn, nat_n);
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
