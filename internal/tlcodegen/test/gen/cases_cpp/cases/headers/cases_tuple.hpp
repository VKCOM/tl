#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/tuple.hpp"
#include "../../__common_namespace/types/pair.hpp"
#include "../../__common_namespace/types/int.hpp"

namespace tl2 { namespace details { 

void BuiltinTupleTupleInt2Reset(std::vector<std::array<int32_t, 2>>& item);

bool BuiltinTupleTupleInt2WriteJSON(std::ostream & s, const std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n);
bool BuiltinTupleTupleInt2Read(::basictl::tl_istream & s, std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n);
bool BuiltinTupleTupleInt2Write(::basictl::tl_ostream & s, const std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTupleTuplePairTupleIntTupleInt2Reset(std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item);

bool BuiltinTupleTuplePairTupleIntTupleInt2WriteJSON(std::ostream & s, const std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_n, uint32_t nat_ttXn, uint32_t nat_ttYn);
bool BuiltinTupleTuplePairTupleIntTupleInt2Read(::basictl::tl_istream & s, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_n, uint32_t nat_ttXn, uint32_t nat_ttYn);
bool BuiltinTupleTuplePairTupleIntTupleInt2Write(::basictl::tl_ostream & s, const std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_n, uint32_t nat_ttXn, uint32_t nat_ttYn);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleIntReset(std::vector<int32_t>& item);

bool TupleIntWriteJSON(std::ostream& s, const std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntReadBoxed(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool TupleIntWriteBoxed(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleInt2Reset(std::array<int32_t, 2>& item);

bool TupleInt2WriteJSON(std::ostream& s, const std::array<int32_t, 2>& item);
bool TupleInt2Read(::basictl::tl_istream & s, std::array<int32_t, 2>& item);
bool TupleInt2Write(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item);
bool TupleInt2ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 2>& item);
bool TupleInt2WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleInt4Reset(std::array<int32_t, 4>& item);

bool TupleInt4WriteJSON(std::ostream& s, const std::array<int32_t, 4>& item);
bool TupleInt4Read(::basictl::tl_istream & s, std::array<int32_t, 4>& item);
bool TupleInt4Write(::basictl::tl_ostream & s, const std::array<int32_t, 4>& item);
bool TupleInt4ReadBoxed(::basictl::tl_istream & s, std::array<int32_t, 4>& item);
bool TupleInt4WriteBoxed(::basictl::tl_ostream & s, const std::array<int32_t, 4>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TuplePairTupleIntTupleInt2Reset(std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item);

bool TuplePairTupleIntTupleInt2WriteJSON(std::ostream& s, const std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn);
bool TuplePairTupleIntTupleInt2Read(::basictl::tl_istream & s, std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn);
bool TuplePairTupleIntTupleInt2Write(::basictl::tl_ostream & s, const std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn);
bool TuplePairTupleIntTupleInt2ReadBoxed(::basictl::tl_istream & s, std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn);
bool TuplePairTupleIntTupleInt2WriteBoxed(::basictl::tl_ostream & s, const std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleTupleInt2Reset(std::vector<std::array<int32_t, 2>>& item);

bool TupleTupleInt2WriteJSON(std::ostream& s, const std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n);
bool TupleTupleInt2Read(::basictl::tl_istream & s, std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n);
bool TupleTupleInt2Write(::basictl::tl_ostream & s, const std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n);
bool TupleTupleInt2ReadBoxed(::basictl::tl_istream & s, std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n);
bool TupleTupleInt2WriteBoxed(::basictl::tl_ostream & s, const std::vector<std::array<int32_t, 2>>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void TupleTuplePairTupleIntTupleInt2Reset(std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item);

bool TupleTuplePairTupleIntTupleInt2WriteJSON(std::ostream& s, const std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_ttXn, uint32_t nat_ttYn, uint32_t nat_n);
bool TupleTuplePairTupleIntTupleInt2Read(::basictl::tl_istream & s, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_ttXn, uint32_t nat_ttYn, uint32_t nat_n);
bool TupleTuplePairTupleIntTupleInt2Write(::basictl::tl_ostream & s, const std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_ttXn, uint32_t nat_ttYn, uint32_t nat_n);
bool TupleTuplePairTupleIntTupleInt2ReadBoxed(::basictl::tl_istream & s, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_ttXn, uint32_t nat_ttYn, uint32_t nat_n);
bool TupleTuplePairTupleIntTupleInt2WriteBoxed(::basictl::tl_ostream & s, const std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_ttXn, uint32_t nat_ttYn, uint32_t nat_n);

}} // namespace tl2::details

