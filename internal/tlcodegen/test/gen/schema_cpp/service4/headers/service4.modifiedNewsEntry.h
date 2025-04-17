#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service4/types/service4.modifiedNewsEntry.h"

namespace tl2 { namespace details { 

void Service4ModifiedNewsEntryReset(::tl2::service4::ModifiedNewsEntry& item) noexcept;

bool Service4ModifiedNewsEntryWriteJSON(std::ostream& s, const ::tl2::service4::ModifiedNewsEntry& item) noexcept;
bool Service4ModifiedNewsEntryRead(::basictl::tl_istream & s, ::tl2::service4::ModifiedNewsEntry& item) noexcept; 
bool Service4ModifiedNewsEntryWrite(::basictl::tl_ostream & s, const ::tl2::service4::ModifiedNewsEntry& item) noexcept;
bool Service4ModifiedNewsEntryReadBoxed(::basictl::tl_istream & s, ::tl2::service4::ModifiedNewsEntry& item);
bool Service4ModifiedNewsEntryWriteBoxed(::basictl::tl_ostream & s, const ::tl2::service4::ModifiedNewsEntry& item);

}} // namespace tl2::details

