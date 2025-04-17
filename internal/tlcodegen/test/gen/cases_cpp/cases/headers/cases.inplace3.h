#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cases/types/cases.inplace3.h"
#include "__common_namespace/types/pair.h"
#include "__common_namespace/types/int.h"

namespace tl2 { namespace details { 

void CasesInplace3TupleInt2Reset(::tl2::cases::Inplace3<std::array<int32_t, 2>>& item) noexcept;

bool CasesInplace3TupleInt2WriteJSON(std::ostream& s, const ::tl2::cases::Inplace3<std::array<int32_t, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) noexcept;
bool CasesInplace3TupleInt2Read(::basictl::tl_istream & s, ::tl2::cases::Inplace3<std::array<int32_t, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) noexcept; 
bool CasesInplace3TupleInt2Write(::basictl::tl_ostream & s, const ::tl2::cases::Inplace3<std::array<int32_t, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3) noexcept;
bool CasesInplace3TupleInt2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Inplace3<std::array<int32_t, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3);
bool CasesInplace3TupleInt2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Inplace3<std::array<int32_t, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void CasesInplace3TuplePairTupleIntTupleInt2Reset(::tl2::cases::Inplace3<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item) noexcept;

bool CasesInplace3TuplePairTupleIntTupleInt2WriteJSON(std::ostream& s, const ::tl2::cases::Inplace3<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XtXn, uint32_t nat_XtYn) noexcept;
bool CasesInplace3TuplePairTupleIntTupleInt2Read(::basictl::tl_istream & s, ::tl2::cases::Inplace3<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XtXn, uint32_t nat_XtYn) noexcept; 
bool CasesInplace3TuplePairTupleIntTupleInt2Write(::basictl::tl_ostream & s, const ::tl2::cases::Inplace3<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XtXn, uint32_t nat_XtYn) noexcept;
bool CasesInplace3TuplePairTupleIntTupleInt2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases::Inplace3<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XtXn, uint32_t nat_XtYn);
bool CasesInplace3TuplePairTupleIntTupleInt2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases::Inplace3<std::array<::tl2::Pair<std::vector<int32_t>, std::vector<int32_t>>, 2>>& item, uint32_t nat_a1, uint32_t nat_a2, uint32_t nat_a3, uint32_t nat_XtXn, uint32_t nat_XtYn);

}} // namespace tl2::details

