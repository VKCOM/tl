#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../../__common_namespace/types/pair.h"
#include "../../__common_namespace/types/tuple.h"
#include "../../__common_namespace/types/int.h"

namespace tl2 { namespace details { 

void BuiltinTuple2PairTupleIntTupleIntReset(std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item);

bool BuiltinTuple2PairTupleIntTupleIntWriteJSON(std::ostream & s, const std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn);
bool BuiltinTuple2PairTupleIntTupleIntRead(::basictl::tl_istream & s, std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn);
bool BuiltinTuple2PairTupleIntTupleIntWrite(::basictl::tl_ostream & s, const std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>& item, uint32_t nat_tXn, uint32_t nat_tYn);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void PairTupleIntTupleIntReset(::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>& item);

bool PairTupleIntTupleIntWriteJSON(std::ostream& s, const ::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>& item, uint32_t nat_X, uint32_t nat_Y);
bool PairTupleIntTupleIntRead(::basictl::tl_istream & s, ::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>& item, uint32_t nat_X, uint32_t nat_Y);
bool PairTupleIntTupleIntWrite(::basictl::tl_ostream & s, const ::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>& item, uint32_t nat_X, uint32_t nat_Y);
bool PairTupleIntTupleIntReadBoxed(::basictl::tl_istream & s, ::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>& item, uint32_t nat_X, uint32_t nat_Y);
bool PairTupleIntTupleIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>& item, uint32_t nat_X, uint32_t nat_Y);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void PairTupleTupleInt2TupleTupleInt2Reset(::tl2::Pair<std::vector<std::array<int32_t, 2>>, std::vector<std::array<int32_t, 2>>>& item);

bool PairTupleTupleInt2TupleTupleInt2WriteJSON(std::ostream& s, const ::tl2::Pair<std::vector<std::array<int32_t, 2>>, std::vector<std::array<int32_t, 2>>>& item, uint32_t nat_X, uint32_t nat_Y);
bool PairTupleTupleInt2TupleTupleInt2Read(::basictl::tl_istream & s, ::tl2::Pair<std::vector<std::array<int32_t, 2>>, std::vector<std::array<int32_t, 2>>>& item, uint32_t nat_X, uint32_t nat_Y);
bool PairTupleTupleInt2TupleTupleInt2Write(::basictl::tl_ostream & s, const ::tl2::Pair<std::vector<std::array<int32_t, 2>>, std::vector<std::array<int32_t, 2>>>& item, uint32_t nat_X, uint32_t nat_Y);
bool PairTupleTupleInt2TupleTupleInt2ReadBoxed(::basictl::tl_istream & s, ::tl2::Pair<std::vector<std::array<int32_t, 2>>, std::vector<std::array<int32_t, 2>>>& item, uint32_t nat_X, uint32_t nat_Y);
bool PairTupleTupleInt2TupleTupleInt2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::Pair<std::vector<std::array<int32_t, 2>>, std::vector<std::array<int32_t, 2>>>& item, uint32_t nat_X, uint32_t nat_Y);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2Reset(::tl2::Pair<std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>>& item);

bool PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2WriteJSON(std::ostream& s, const ::tl2::Pair<std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>>& item, uint32_t nat_XttXn, uint32_t nat_XttYn, uint32_t nat_Xn, uint32_t nat_YttXn, uint32_t nat_YttYn, uint32_t nat_Yn);
bool PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2Read(::basictl::tl_istream & s, ::tl2::Pair<std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>>& item, uint32_t nat_XttXn, uint32_t nat_XttYn, uint32_t nat_Xn, uint32_t nat_YttXn, uint32_t nat_YttYn, uint32_t nat_Yn);
bool PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2Write(::basictl::tl_ostream & s, const ::tl2::Pair<std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>>& item, uint32_t nat_XttXn, uint32_t nat_XttYn, uint32_t nat_Xn, uint32_t nat_YttXn, uint32_t nat_YttYn, uint32_t nat_Yn);
bool PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2ReadBoxed(::basictl::tl_istream & s, ::tl2::Pair<std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>>& item, uint32_t nat_XttXn, uint32_t nat_XttYn, uint32_t nat_Xn, uint32_t nat_YttXn, uint32_t nat_YttYn, uint32_t nat_Yn);
bool PairTupleTuplePairTupleIntTupleInt2TupleTuplePairTupleIntTupleInt2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::Pair<std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>, std::vector<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>>& item, uint32_t nat_XttXn, uint32_t nat_XttYn, uint32_t nat_Xn, uint32_t nat_YttXn, uint32_t nat_YttYn, uint32_t nat_Yn);

}} // namespace tl2::details

