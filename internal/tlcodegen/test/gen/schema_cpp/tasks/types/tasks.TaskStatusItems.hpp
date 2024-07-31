#pragma once

#include "../../a_tlgen_helpers_code.hpp"


namespace tl2 { namespace tasks { 
struct TaskStatusInProgress {

	std::string_view tl_name() const { return "tasks.taskStatusInProgress"; }
	uint32_t tl_tag() const { return 0x06ef70e7; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TaskStatusInProgress& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

namespace tl2 { namespace tasks { 
struct TaskStatusNotCurrentlyInEngine {

	std::string_view tl_name() const { return "tasks.taskStatusNotCurrentlyInEngine"; }
	uint32_t tl_tag() const { return 0xb207caaa; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TaskStatusNotCurrentlyInEngine& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

namespace tl2 { namespace tasks { 
struct TaskStatusScheduled {

	std::string_view tl_name() const { return "tasks.taskStatusScheduled"; }
	uint32_t tl_tag() const { return 0x0aca80a9; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TaskStatusScheduled& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

namespace tl2 { namespace tasks { 
struct TaskStatusWaiting {

	std::string_view tl_name() const { return "tasks.taskStatusWaiting"; }
	uint32_t tl_tag() const { return 0x16739c2c; }

	bool write_json(std::ostream& s)const;

	bool read(::basictl::tl_istream & s);
	bool write(::basictl::tl_ostream & s)const;

	bool read_boxed(::basictl::tl_istream & s);
	bool write_boxed(::basictl::tl_ostream & s)const;

	friend std::ostream& operator<<(std::ostream& s, const TaskStatusWaiting& rhs) {
		rhs.write_json(s);
		return s;
	}
};

}} // namespace tl2::tasks

