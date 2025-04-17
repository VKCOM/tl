#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.inplace2.h"
#include "__common_namespace/types/pair.h"
#include "__common_namespace/types/int.h"

namespace tl2 { namespace details { 

void CasesInplace2IntReset(::tl2::cases::Inplace2<int32_t>& item) noexcept;

bool CasesInplace2IntWriteJSON(std::ostream& s, const ::tl2::cases::Inplace2<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) noexcept;
bool CasesInplace2IntRead(::basictl::tl_istream & s, ::tl2::cases::Inplace2<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) noexcept; 
bool CasesInplace2IntWrite(::basictl::tl_ostream & s, const ::tl2::cases::Inplace2<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) noexcept;
bool CasesInplace2IntReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Inplace2<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3);
bool CasesInplace2IntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Inplace2<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void CasesInplace2PairTupleIntTupleIntReset(::tl2::cases::Inplace2<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item) noexcept;

bool CasesInplace2PairTupleIntTupleIntWriteJSON(std::ostream& s, const ::tl2::cases::Inplace2<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) noexcept;
bool CasesInplace2PairTupleIntTupleIntRead(::basictl::tl_istream & s, ::tl2::cases::Inplace2<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) noexcept; 
bool CasesInplace2PairTupleIntTupleIntWrite(::basictl::tl_ostream & s, const ::tl2::cases::Inplace2<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) noexcept;
bool CasesInplace2PairTupleIntTupleIntReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Inplace2<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn);
bool CasesInplace2PairTupleIntTupleIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Inplace2<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn);

}} // namespace tl2::details

