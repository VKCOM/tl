#pragma once

#include "../../basictl/io_streams.h"
#include "../../basictl/io_throwable_streams.h"
#include "../types/cases.inplace1.h"
#include "../../__common_namespace/types/pair.h"
#include "../../__common_namespace/types/int.h"

namespace tl2 { namespace details { 

void CasesInplace1IntReset(::tl2::cases::Inplace1<int32_t>& item) noexcept;

bool CasesInplace1IntWriteJSON(std::ostream& s, const ::tl2::cases::Inplace1<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) noexcept;
bool CasesInplace1IntRead(::basictl::tl_istream & s, ::tl2::cases::Inplace1<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) noexcept; 
bool CasesInplace1IntWrite(::basictl::tl_ostream & s, const ::tl2::cases::Inplace1<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) noexcept;
bool CasesInplace1IntReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Inplace1<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3);
bool CasesInplace1IntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Inplace1<int32_t>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void CasesInplace1PairTupleIntTupleIntReset(::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item) noexcept;

bool CasesInplace1PairTupleIntTupleIntWriteJSON(std::ostream& s, const ::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) noexcept;
bool CasesInplace1PairTupleIntTupleIntRead(::basictl::tl_istream & s, ::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) noexcept; 
bool CasesInplace1PairTupleIntTupleIntWrite(::basictl::tl_ostream & s, const ::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn) noexcept;
bool CasesInplace1PairTupleIntTupleIntReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn);
bool CasesInplace1PairTupleIntTupleIntWriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Inplace1<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XXn, uint32_t nat_XYn);

}} // namespace tl2::details

