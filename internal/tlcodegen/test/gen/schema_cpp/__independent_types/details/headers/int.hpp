#pragma once

#include "../../../a_tlgen_helpers_code.hpp"
#include "../../../__common/types/int.hpp"

namespace tl2 { namespace details { 

void BuiltinTuple10IntBoxedReset(std::array<int32_t, 10>& item);
bool BuiltinTuple10IntBoxedRead(::basictl::tl_istream & s, std::array<int32_t, 10>& item);
bool BuiltinTuple10IntBoxedWrite(::basictl::tl_ostream & s, const std::array<int32_t, 10>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTuple2IntBoxedReset(std::array<int32_t, 2>& item);
bool BuiltinTuple2IntBoxedRead(::basictl::tl_istream & s, std::array<int32_t, 2>& item);
bool BuiltinTuple2IntBoxedWrite(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTuple3IntReset(std::array<int32_t, 3>& item);
bool BuiltinTuple3IntRead(::basictl::tl_istream & s, std::array<int32_t, 3>& item);
bool BuiltinTuple3IntWrite(::basictl::tl_ostream & s, const std::array<int32_t, 3>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTuple5IntReset(std::array<int32_t, 5>& item);
bool BuiltinTuple5IntRead(::basictl::tl_istream & s, std::array<int32_t, 5>& item);
bool BuiltinTuple5IntWrite(::basictl::tl_ostream & s, const std::array<int32_t, 5>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTupleIntBoxedReset(std::vector<int32_t>& item);
bool BuiltinTupleIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool BuiltinTupleIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorIntBoxedReset(std::vector<int32_t>& item);
bool BuiltinVectorIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item);
bool BuiltinVectorIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void IntReset(int32_t& item);
bool IntRead(::basictl::tl_istream & s, int32_t& item);
bool IntWrite(::basictl::tl_ostream & s, const int32_t& item);
bool IntReadBoxed(::basictl::tl_istream & s, int32_t& item);
bool IntWriteBoxed(::basictl::tl_ostream & s, const int32_t& item);

}} // namespace tl2::details

