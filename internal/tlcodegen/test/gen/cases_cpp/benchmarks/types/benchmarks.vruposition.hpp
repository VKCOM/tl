#pragma once

#include "../../a_tlgen_helpers_code.hpp"
#include "../../__common_namespace/types/true.hpp"
#include "benchmarks.vruhash.hpp"


namespace tl2 { namespace benchmarks { 
struct Vruposition {
	uint32_t fields_mask = 0;
	::tl2::True commit_bit{};
	::tl2::True meta_block{};
	::tl2::True split_payload{};
	::tl2::True rotation_block{};
	::tl2::True canonical_hash{};
	int64_t payload_offset = 0;
	int64_t block_time_nano = 0;
	::tl2::benchmarks::Vruhash hash{};
	int64_t file_offset = 0;
	int64_t seq_number = 0;

	std::string_view tl_name() const { return "benchmarks.vruposition"; }
	uint32_t tl_tag() const { return 0x32792c04; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;
};

}} // namespace tl2::benchmarks

