// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "service4/types/service4.modifiedNewsEntry.h"

namespace tlgen { namespace details { 

void Service4ModifiedNewsEntryReset(::tlgen::service4::ModifiedNewsEntry& item) noexcept;

bool Service4ModifiedNewsEntryWriteJSON(std::ostream& s, const ::tlgen::service4::ModifiedNewsEntry& item) noexcept;
bool Service4ModifiedNewsEntryRead(::tlgen::basictl::tl_istream & s, ::tlgen::service4::ModifiedNewsEntry& item) noexcept; 
bool Service4ModifiedNewsEntryWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::service4::ModifiedNewsEntry& item) noexcept;
bool Service4ModifiedNewsEntryReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::service4::ModifiedNewsEntry& item);
bool Service4ModifiedNewsEntryWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::service4::ModifiedNewsEntry& item);

}} // namespace tlgen::details

