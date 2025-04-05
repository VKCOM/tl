#pragma once

#include "../../basictl/io_streams.h"
#include "../types/int.h"

namespace tl2 { namespace details { 

void BuiltinTuple10IntBoxedReset(std::array<int32_t, 10>& item);

bool BuiltinTuple10IntBoxedWriteJSON(std::ostream & s, const std::array<int32_t, 10>& item);
bool BuiltinTuple10IntBoxedRead(::basictl::tl_istream & s, std::array<int32_t, 10>& item);
bool BuiltinTuple10IntBoxedWrite(::basictl::tl_ostream & s, const std::array<int32_t, 10>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTuple2IntBoxedReset(std::array<int32_t, 2>& item);

bool BuiltinTuple2IntBoxedWriteJSON(std::ostream & s, const std::array<int32_t, 2>& item);
bool BuiltinTuple2IntBoxedRead(::basictl::tl_istream & s, std::array<int32_t, 2>& item);
bool BuiltinTuple2IntBoxedWrite(::basictl::tl_ostream & s, const std::array<int32_t, 2>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTuple3IntReset(std::array<int32_t, 3>& item);

bool BuiltinTuple3IntWriteJSON(std::ostream & s, const std::array<int32_t, 3>& item);
bool BuiltinTuple3IntRead(::basictl::tl_istream & s, std::array<int32_t, 3>& item);
bool BuiltinTuple3IntWrite(::basictl::tl_ostream & s, const std::array<int32_t, 3>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTuple5IntReset(std::array<int32_t, 5>& item);

bool BuiltinTuple5IntWriteJSON(std::ostream & s, const std::array<int32_t, 5>& item);
bool BuiltinTuple5IntRead(::basictl::tl_istream & s, std::array<int32_t, 5>& item);
bool BuiltinTuple5IntWrite(::basictl::tl_ostream & s, const std::array<int32_t, 5>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTupleIntReset(std::vector<int32_t>& item);

bool BuiltinTupleIntWriteJSON(std::ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);
bool BuiltinTupleIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool BuiltinTupleIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTupleIntBoxedReset(std::vector<int32_t>& item);

bool BuiltinTupleIntBoxedWriteJSON(std::ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);
bool BuiltinTupleIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool BuiltinTupleIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorIntReset(std::vector<int32_t>& item);

bool BuiltinVectorIntWriteJSON(std::ostream & s, const std::vector<int32_t>& item);
bool BuiltinVectorIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item);
bool BuiltinVectorIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorIntBoxedReset(std::vector<int32_t>& item);

bool BuiltinVectorIntBoxedWriteJSON(std::ostream & s, const std::vector<int32_t>& item);
bool BuiltinVectorIntBoxedRead(::basictl::tl_istream & s, std::vector<int32_t>& item);
bool BuiltinVectorIntBoxedWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void IntReset(int32_t& item);

bool IntWriteJSON(std::ostream& s, const int32_t& item);
bool IntRead(::basictl::tl_istream & s, int32_t& item);
bool IntWrite(::basictl::tl_ostream & s, const int32_t& item);
bool IntReadBoxed(::basictl::tl_istream & s, int32_t& item);
bool IntWriteBoxed(::basictl::tl_ostream & s, const int32_t& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool IntMaybeWriteJSON(std::ostream & s, const std::optional<int32_t>& item);

bool IntMaybeReadBoxed(::basictl::tl_istream & s, std::optional<int32_t>& item);
bool IntMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<int32_t>& item);


}} // namespace tl2::details

