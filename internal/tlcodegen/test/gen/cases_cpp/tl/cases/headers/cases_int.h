#pragma once

#include "../../../basics/basictl.h"
#include "../../__common_namespace/types/int.h"

namespace tl2 { namespace details { 

void BuiltinTuple4IntReset(std::array<int32_t, 4>& item);

bool BuiltinTuple4IntWriteJSON(std::ostream & s, const std::array<int32_t, 4>& item);
bool BuiltinTuple4IntRead(::basictl::tl_istream & s, std::array<int32_t, 4>& item);
bool BuiltinTuple4IntWrite(::basictl::tl_ostream & s, const std::array<int32_t, 4>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTupleIntReset(std::vector<int32_t>& item);

bool BuiltinTupleIntWriteJSON(std::ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);
bool BuiltinTupleIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item, uint32_t nat_n);
bool BuiltinTupleIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item, uint32_t nat_n);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinTupleTupleIntReset(std::vector<std::vector<int32_t>>& item);

bool BuiltinTupleTupleIntWriteJSON(std::ostream & s, const std::vector<std::vector<int32_t>>& item, uint32_t nat_n, uint32_t nat_t);
bool BuiltinTupleTupleIntRead(::basictl::tl_istream & s, std::vector<std::vector<int32_t>>& item, uint32_t nat_n, uint32_t nat_t);
bool BuiltinTupleTupleIntWrite(::basictl::tl_ostream & s, const std::vector<std::vector<int32_t>>& item, uint32_t nat_n, uint32_t nat_t);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void BuiltinVectorIntReset(std::vector<int32_t>& item);

bool BuiltinVectorIntWriteJSON(std::ostream & s, const std::vector<int32_t>& item);
bool BuiltinVectorIntRead(::basictl::tl_istream & s, std::vector<int32_t>& item);
bool BuiltinVectorIntWrite(::basictl::tl_ostream & s, const std::vector<int32_t>& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

bool IntMaybeWriteJSON(std::ostream & s, const std::optional<int32_t>& item);

bool IntMaybeReadBoxed(::basictl::tl_istream & s, std::optional<int32_t>& item);
bool IntMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<int32_t>& item);


}} // namespace tl2::details

