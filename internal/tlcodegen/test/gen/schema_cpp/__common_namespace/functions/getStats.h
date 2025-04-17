#pragma once

#include "basictl/io_streams.h"
#include "basictl/io_throwable_streams.h"
#include "tasks/types/tasks.queueTypeStats.h"


namespace tl2 { 
struct GetStats {
	::tl2::tasks::QueueTypeStats x{};

	std::string_view tl_name() const { return "getStats"; }
	uint32_t tl_tag() const { return 0xbaa6da35; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeStats & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, ::tl2::tasks::QueueTypeStats & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::tasks::QueueTypeStats & result);
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::tasks::QueueTypeStats & result);

	friend std::ostream& operator<<(std::ostream& s, const GetStats& rhs) {
		rhs.write_json(s);
		return s;
	}
};

} // namespace tl2

