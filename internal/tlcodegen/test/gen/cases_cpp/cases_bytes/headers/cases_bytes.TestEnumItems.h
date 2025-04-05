#pragma once

#include "../../basictl/io_streams.h"
#include "../types/cases_bytes.TestEnumItems.h"

namespace tl2 { namespace details { 

void CasesBytesTestEnum1Reset(::tl2::cases_bytes::TestEnum1& item);

bool CasesBytesTestEnum1WriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestEnum1& item);
bool CasesBytesTestEnum1Read(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum1& item);
bool CasesBytesTestEnum1Write(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum1& item);
bool CasesBytesTestEnum1ReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum1& item);
bool CasesBytesTestEnum1WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum1& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void CasesBytesTestEnum2Reset(::tl2::cases_bytes::TestEnum2& item);

bool CasesBytesTestEnum2WriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestEnum2& item);
bool CasesBytesTestEnum2Read(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum2& item);
bool CasesBytesTestEnum2Write(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum2& item);
bool CasesBytesTestEnum2ReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum2& item);
bool CasesBytesTestEnum2WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum2& item);

}} // namespace tl2::details

namespace tl2 { namespace details { 

void CasesBytesTestEnum3Reset(::tl2::cases_bytes::TestEnum3& item);

bool CasesBytesTestEnum3WriteJSON(std::ostream& s, const ::tl2::cases_bytes::TestEnum3& item);
bool CasesBytesTestEnum3Read(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum3& item);
bool CasesBytesTestEnum3Write(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum3& item);
bool CasesBytesTestEnum3ReadBoxed(::basictl::tl_istream & s, ::tl2::cases_bytes::TestEnum3& item);
bool CasesBytesTestEnum3WriteBoxed(::basictl::tl_ostream & s, const ::tl2::cases_bytes::TestEnum3& item);

}} // namespace tl2::details

