// Code generated by vktl/cmd/tlgen2; DO NOT EDIT.
#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "cd/types/cd.useCycle.h"

namespace tlgen { namespace details { 

void CdUseCycleReset(::tlgen::cd::UseCycle& item) noexcept;

bool CdUseCycleWriteJSON(std::ostream& s, const ::tlgen::cd::UseCycle& item) noexcept;
bool CdUseCycleRead(::tlgen::basictl::tl_istream & s, ::tlgen::cd::UseCycle& item) noexcept; 
bool CdUseCycleWrite(::tlgen::basictl::tl_ostream & s, const ::tlgen::cd::UseCycle& item) noexcept;
bool CdUseCycleReadBoxed(::tlgen::basictl::tl_istream & s, ::tlgen::cd::UseCycle& item);
bool CdUseCycleWriteBoxed(::tlgen::basictl::tl_ostream & s, const ::tlgen::cd::UseCycle& item);

}} // namespace tlgen::details

