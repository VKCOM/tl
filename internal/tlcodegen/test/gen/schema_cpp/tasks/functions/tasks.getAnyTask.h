#pragma once

#include "../../basictl/io_streams.h"
#include "../types/tasks.taskInfo.h"


namespace tl2 { namespace tasks { 
struct GetAnyTask {

	std::string_view tl_name() const { return "tasks.getAnyTask"; }
	uint32_t tl_tag() const { return 0x4a9c7dbb; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s) noexcept;
	bool write(::basictl::tl_ostream & s)const noexcept;

	void read_or_throw(::basictl::tl_throwable_istream & s);
	void write_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s) noexcept;
	bool write_boxed(::basictl::tl_ostream & s)const noexcept;
	
	void read_boxed_or_throw(::basictl::tl_throwable_istream & s);
	void write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const;

	bool read_result(::basictl::tl_istream & s, std::optional<::tl2::tasks::TaskInfo> & result) noexcept;
	bool write_result(::basictl::tl_ostream & s, std::optional<::tl2::tasks::TaskInfo> & result) noexcept;

	void read_result_or_throw(::basictl::tl_throwable_istream & s, std::optional<::tl2::tasks::TaskInfo> & result) noexcept;
	void write_result_or_throw(::basictl::tl_throwable_ostream & s, std::optional<::tl2::tasks::TaskInfo> & result) noexcept;

	friend std::ostream& operator<<(std::ostream& s, const GetAnyTask& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

