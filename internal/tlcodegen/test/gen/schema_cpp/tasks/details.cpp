#include "headers/tasks.TaskStatus.h"
#include "headers/tasks.TaskStatusItems.h"
#include "headers/tasks.taskInfo.h"
#include "headers/tasks.queueTypeInfo.h"
#include "headers/tasks.queueTypeSettings.h"
#include "headers/tasks.queueStats.h"
#include "headers/tasks.getTaskFromQueue.h"
#include "headers/tasks.getQueueTypes.h"
#include "headers/tasks.getQueueSize.h"
#include "headers/tasks.getAnyTask.h"
#include "headers/tasks.fullFilledCron.h"
#include "headers/tasks.cronTaskWithId.h"
#include "headers/tasks.cronTask.h"
#include "headers/tasks.cronTime.h"
#include "headers/tasks.addTask.h"
#include "headers/tasks.task.h"
#include "headers/tasks.queueTypeStats.h"
#include "../__common_namespace/headers/int.h"
#include "../__common_namespace/headers/long.h"
#include "../__common_namespace/headers/Bool.h"


void tl2::details::BuiltinVectorTasksQueueTypeInfoReset(std::vector<::tl2::tasks::QueueTypeInfo>& item) {
	item.resize(0); // TODO - unwrap
}

bool tl2::details::BuiltinVectorTasksQueueTypeInfoWriteJSON(std::ostream & s, const std::vector<::tl2::tasks::QueueTypeInfo>& item) {
	s << "[";
	size_t index = 0;
	for(const auto & el : item) {
		if (!::tl2::details::TasksQueueTypeInfoWriteJSON(s, el)) { return false; }
		if (index != item.size() - 1) {
			s << ",";
		}
		index++;
	}
	s << "]";
	return true;
}

bool tl2::details::BuiltinVectorTasksQueueTypeInfoRead(::basictl::tl_istream & s, std::vector<::tl2::tasks::QueueTypeInfo>& item) {
	uint32_t len = 0;
	if (!s.nat_read(len)) { return false; }
	// TODO - check length sanity
	item.resize(len);
	for(auto && el : item) {
		if (!::tl2::details::TasksQueueTypeInfoRead(s, el)) { return false; }
	}
	return true;
}

bool tl2::details::BuiltinVectorTasksQueueTypeInfoWrite(::basictl::tl_ostream & s, const std::vector<::tl2::tasks::QueueTypeInfo>& item) {
	if (!s.nat_write(item.size())) { return false; }
	for(const auto & el : item) {
		if (!::tl2::details::TasksQueueTypeInfoWrite(s, el)) { return false; }
	}
	return true;
}

bool tl2::tasks::AddTask::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksAddTaskWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::AddTask::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksAddTaskRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::AddTask::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksAddTaskWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::AddTask::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::AddTask::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::AddTask::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksAddTaskReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::AddTask::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksAddTaskWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::AddTask::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::AddTask::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksAddTaskReset(::tl2::tasks::AddTask& item) {
	item.type_name.clear();
	item.queue_id.clear();
	::tl2::details::TasksTaskReset(item.task);
}

bool tl2::details::TasksAddTaskWriteJSON(std::ostream& s, const ::tl2::tasks::AddTask& item) {
	auto add_comma = false;
	s << "{";
	if (item.type_name.size() != 0) {
		add_comma = true;
		s << "\"type_name\":";
		s << "\"" << item.type_name << "\"";
	}
	if (item.queue_id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"queue_id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.queue_id)) { return false; }
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"task\":";
	if (!::tl2::details::TasksTaskWriteJSON(s, item.task)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::TasksAddTaskRead(::basictl::tl_istream & s, ::tl2::tasks::AddTask& item) {
	if (!s.string_read(item.type_name)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.queue_id)) { return false; }
	if (!::tl2::details::TasksTaskRead(s, item.task)) { return false; }
	return true;
}

bool tl2::details::TasksAddTaskWrite(::basictl::tl_ostream & s, const ::tl2::tasks::AddTask& item) {
	if (!s.string_write(item.type_name)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.queue_id)) { return false; }
	if (!::tl2::details::TasksTaskWrite(s, item.task)) { return false; }
	return true;
}

bool tl2::details::TasksAddTaskReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::AddTask& item) {
	if (!s.nat_read_exact_tag(0x2ca073d5)) { return false; }
	return tl2::details::TasksAddTaskRead(s, item);
}

bool tl2::details::TasksAddTaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::AddTask& item) {
	if (!s.nat_write(0x2ca073d5)) { return false; }
	return tl2::details::TasksAddTaskWrite(s, item);
}

bool tl2::details::TasksAddTaskReadResult(::basictl::tl_istream & s, tl2::tasks::AddTask& item, int64_t& result) {
	if (!s.nat_read_exact_tag(0x22076cba)) { return false;}
	if (!s.long_read(result)) { return false; }
	return true;
}
bool tl2::details::TasksAddTaskWriteResult(::basictl::tl_ostream & s, tl2::tasks::AddTask& item, int64_t& result) {
	if (!s.nat_write(0x22076cba)) { return false; }
	if (!s.long_write(result)) { return false;}
	return true;
}

bool tl2::tasks::AddTask::read_result(::basictl::tl_istream & s, int64_t & result) noexcept {
	bool success = tl2::details::TasksAddTaskReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::tasks::AddTask::write_result(::basictl::tl_ostream & s, int64_t & result) noexcept {
	bool success = tl2::details::TasksAddTaskWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::tasks::AddTask::read_result_or_throw(::basictl::tl_throwable_istream & s, int64_t & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::tasks::AddTask::write_result_or_throw(::basictl::tl_throwable_ostream & s, int64_t & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::tasks::CronTask::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksCronTaskWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::CronTask::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksCronTaskRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::CronTask::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksCronTaskWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::CronTask::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::CronTask::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::CronTask::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksCronTaskReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::CronTask::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksCronTaskWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::CronTask::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::CronTask::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksCronTaskReset(::tl2::tasks::CronTask& item) {
	item.type_name.clear();
	item.queue_id.clear();
	::tl2::details::TasksTaskReset(item.task);
	::tl2::details::TasksCronTimeReset(item.time);
}

bool tl2::details::TasksCronTaskWriteJSON(std::ostream& s, const ::tl2::tasks::CronTask& item) {
	auto add_comma = false;
	s << "{";
	if (item.type_name.size() != 0) {
		add_comma = true;
		s << "\"type_name\":";
		s << "\"" << item.type_name << "\"";
	}
	if (item.queue_id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"queue_id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.queue_id)) { return false; }
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"task\":";
	if (!::tl2::details::TasksTaskWriteJSON(s, item.task)) { return false; }
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"time\":";
	if (!::tl2::details::TasksCronTimeWriteJSON(s, item.time)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::TasksCronTaskRead(::basictl::tl_istream & s, ::tl2::tasks::CronTask& item) {
	if (!s.string_read(item.type_name)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.queue_id)) { return false; }
	if (!::tl2::details::TasksTaskRead(s, item.task)) { return false; }
	if (!::tl2::details::TasksCronTimeRead(s, item.time)) { return false; }
	return true;
}

bool tl2::details::TasksCronTaskWrite(::basictl::tl_ostream & s, const ::tl2::tasks::CronTask& item) {
	if (!s.string_write(item.type_name)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.queue_id)) { return false; }
	if (!::tl2::details::TasksTaskWrite(s, item.task)) { return false; }
	if (!::tl2::details::TasksCronTimeWrite(s, item.time)) { return false; }
	return true;
}

bool tl2::details::TasksCronTaskReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::CronTask& item) {
	if (!s.nat_read_exact_tag(0xc90cf28a)) { return false; }
	return tl2::details::TasksCronTaskRead(s, item);
}

bool tl2::details::TasksCronTaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::CronTask& item) {
	if (!s.nat_write(0xc90cf28a)) { return false; }
	return tl2::details::TasksCronTaskWrite(s, item);
}

bool tl2::tasks::CronTaskWithId::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksCronTaskWithIdWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::CronTaskWithId::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksCronTaskWithIdRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::CronTaskWithId::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksCronTaskWithIdWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::CronTaskWithId::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::CronTaskWithId::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::CronTaskWithId::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksCronTaskWithIdReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::CronTaskWithId::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksCronTaskWithIdWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::CronTaskWithId::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::CronTaskWithId::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksCronTaskWithIdReset(::tl2::tasks::CronTaskWithId& item) {
	item.id = 0;
	item.next_time = 0;
	::tl2::details::TasksCronTaskReset(item.task);
}

bool tl2::details::TasksCronTaskWithIdWriteJSON(std::ostream& s, const ::tl2::tasks::CronTaskWithId& item) {
	auto add_comma = false;
	s << "{";
	if (item.id != 0) {
		add_comma = true;
		s << "\"id\":";
		s << item.id;
	}
	if (item.next_time != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"next_time\":";
		s << item.next_time;
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"task\":";
	if (!::tl2::details::TasksCronTaskWriteJSON(s, item.task)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::TasksCronTaskWithIdRead(::basictl::tl_istream & s, ::tl2::tasks::CronTaskWithId& item) {
	if (!s.int_read(item.id)) { return false; }
	if (!s.int_read(item.next_time)) { return false; }
	if (!::tl2::details::TasksCronTaskRead(s, item.task)) { return false; }
	return true;
}

bool tl2::details::TasksCronTaskWithIdWrite(::basictl::tl_ostream & s, const ::tl2::tasks::CronTaskWithId& item) {
	if (!s.int_write(item.id)) { return false;}
	if (!s.int_write(item.next_time)) { return false;}
	if (!::tl2::details::TasksCronTaskWrite(s, item.task)) { return false; }
	return true;
}

bool tl2::details::TasksCronTaskWithIdReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::CronTaskWithId& item) {
	if (!s.nat_read_exact_tag(0x3a958001)) { return false; }
	return tl2::details::TasksCronTaskWithIdRead(s, item);
}

bool tl2::details::TasksCronTaskWithIdWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::CronTaskWithId& item) {
	if (!s.nat_write(0x3a958001)) { return false; }
	return tl2::details::TasksCronTaskWithIdWrite(s, item);
}

bool tl2::tasks::CronTime::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksCronTimeWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::CronTime::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksCronTimeRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::CronTime::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksCronTimeWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::CronTime::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::CronTime::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::CronTime::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksCronTimeReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::CronTime::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksCronTimeWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::CronTime::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::CronTime::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksCronTimeReset(::tl2::tasks::CronTime& item) {
	item.fields_mask = 0;
	item.seconds.clear();
	item.minutes.clear();
	item.hours.clear();
	item.days_of_week.clear();
	item.days.clear();
	item.months.clear();
}

bool tl2::details::TasksCronTimeWriteJSON(std::ostream& s, const ::tl2::tasks::CronTime& item) {
	auto add_comma = false;
	s << "{";
	if (item.fields_mask != 0) {
		add_comma = true;
		s << "\"fields_mask\":";
		s << item.fields_mask;
	}
	if ((item.fields_mask & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"seconds\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.seconds)) { return false; }
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"minutes\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.minutes)) { return false; }
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"hours\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.hours)) { return false; }
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"days_of_week\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.days_of_week)) { return false; }
	}
	if ((item.fields_mask & (1<<4)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"days\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.days)) { return false; }
	}
	if ((item.fields_mask & (1<<5)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"months\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.months)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::TasksCronTimeRead(::basictl::tl_istream & s, ::tl2::tasks::CronTime& item) {
	if (!s.nat_read(item.fields_mask)) { return false; }
	if ((item.fields_mask & (1<<0)) != 0) {
		if (!::tl2::details::BuiltinVectorIntRead(s, item.seconds)) { return false; }
	} else {
			item.seconds.clear();
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (!::tl2::details::BuiltinVectorIntRead(s, item.minutes)) { return false; }
	} else {
			item.minutes.clear();
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (!::tl2::details::BuiltinVectorIntRead(s, item.hours)) { return false; }
	} else {
			item.hours.clear();
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (!::tl2::details::BuiltinVectorIntRead(s, item.days_of_week)) { return false; }
	} else {
			item.days_of_week.clear();
	}
	if ((item.fields_mask & (1<<4)) != 0) {
		if (!::tl2::details::BuiltinVectorIntRead(s, item.days)) { return false; }
	} else {
			item.days.clear();
	}
	if ((item.fields_mask & (1<<5)) != 0) {
		if (!::tl2::details::BuiltinVectorIntRead(s, item.months)) { return false; }
	} else {
			item.months.clear();
	}
	return true;
}

bool tl2::details::TasksCronTimeWrite(::basictl::tl_ostream & s, const ::tl2::tasks::CronTime& item) {
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!::tl2::details::BuiltinVectorIntWrite(s, item.seconds)) { return false; }
	}
	if ((item.fields_mask & (1<<1)) != 0) {
			if (!::tl2::details::BuiltinVectorIntWrite(s, item.minutes)) { return false; }
	}
	if ((item.fields_mask & (1<<2)) != 0) {
			if (!::tl2::details::BuiltinVectorIntWrite(s, item.hours)) { return false; }
	}
	if ((item.fields_mask & (1<<3)) != 0) {
			if (!::tl2::details::BuiltinVectorIntWrite(s, item.days_of_week)) { return false; }
	}
	if ((item.fields_mask & (1<<4)) != 0) {
			if (!::tl2::details::BuiltinVectorIntWrite(s, item.days)) { return false; }
	}
	if ((item.fields_mask & (1<<5)) != 0) {
			if (!::tl2::details::BuiltinVectorIntWrite(s, item.months)) { return false; }
	}
	return true;
}

bool tl2::details::TasksCronTimeReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::CronTime& item) {
	if (!s.nat_read_exact_tag(0xd4177d7f)) { return false; }
	return tl2::details::TasksCronTimeRead(s, item);
}

bool tl2::details::TasksCronTimeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::CronTime& item) {
	if (!s.nat_write(0xd4177d7f)) { return false; }
	return tl2::details::TasksCronTimeWrite(s, item);
}

bool tl2::tasks::FullFilledCron::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksFullFilledCronWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::FullFilledCron::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksFullFilledCronRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::FullFilledCron::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksFullFilledCronWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::FullFilledCron::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::FullFilledCron::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::FullFilledCron::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksFullFilledCronReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::FullFilledCron::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksFullFilledCronWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::FullFilledCron::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::FullFilledCron::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksFullFilledCronReset(::tl2::tasks::FullFilledCron& item) {
	item.fields_mask = 0;
	item.a0 = 0;
	item.a1 = 0;
	item.a2 = 0;
	item.a3 = 0;
	item.a4 = 0;
	item.a5 = 0;
	item.a6 = 0;
	item.a7 = 0;
	item.a8 = 0;
	item.a9 = 0;
	item.a10 = 0;
	item.a11 = 0;
	item.a12 = 0;
	item.a13 = 0;
	item.a14 = 0;
	item.a15 = 0;
	item.a16 = 0;
	item.a17 = 0;
	item.a18 = 0;
	item.a19 = 0;
	item.a20 = 0;
	item.a21 = 0;
	item.a22 = 0;
	item.a23 = 0;
	item.a24 = 0;
	item.a25 = 0;
	item.a26 = 0;
	item.a27 = 0;
	item.a28 = 0;
	item.a29 = 0;
	item.a30 = 0;
	item.a31 = 0;
}

bool tl2::details::TasksFullFilledCronWriteJSON(std::ostream& s, const ::tl2::tasks::FullFilledCron& item) {
	auto add_comma = false;
	s << "{";
	if (item.fields_mask != 0) {
		add_comma = true;
		s << "\"fields_mask\":";
		s << item.fields_mask;
	}
	if ((item.fields_mask & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a0\":";
		s << item.a0;
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a1\":";
		s << item.a1;
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a2\":";
		s << item.a2;
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a3\":";
		s << item.a3;
	}
	if ((item.fields_mask & (1<<4)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a4\":";
		s << item.a4;
	}
	if ((item.fields_mask & (1<<5)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a5\":";
		s << item.a5;
	}
	if ((item.fields_mask & (1<<6)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a6\":";
		s << item.a6;
	}
	if ((item.fields_mask & (1<<7)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a7\":";
		s << item.a7;
	}
	if ((item.fields_mask & (1<<8)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a8\":";
		s << item.a8;
	}
	if ((item.fields_mask & (1<<9)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a9\":";
		s << item.a9;
	}
	if ((item.fields_mask & (1<<10)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a10\":";
		s << item.a10;
	}
	if ((item.fields_mask & (1<<11)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a11\":";
		s << item.a11;
	}
	if ((item.fields_mask & (1<<12)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a12\":";
		s << item.a12;
	}
	if ((item.fields_mask & (1<<13)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a13\":";
		s << item.a13;
	}
	if ((item.fields_mask & (1<<14)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a14\":";
		s << item.a14;
	}
	if ((item.fields_mask & (1<<15)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a15\":";
		s << item.a15;
	}
	if ((item.fields_mask & (1<<16)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a16\":";
		s << item.a16;
	}
	if ((item.fields_mask & (1<<17)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a17\":";
		s << item.a17;
	}
	if ((item.fields_mask & (1<<18)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a18\":";
		s << item.a18;
	}
	if ((item.fields_mask & (1<<19)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a19\":";
		s << item.a19;
	}
	if ((item.fields_mask & (1<<20)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a20\":";
		s << item.a20;
	}
	if ((item.fields_mask & (1<<21)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a21\":";
		s << item.a21;
	}
	if ((item.fields_mask & (1<<22)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a22\":";
		s << item.a22;
	}
	if ((item.fields_mask & (1<<23)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a23\":";
		s << item.a23;
	}
	if ((item.fields_mask & (1<<24)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a24\":";
		s << item.a24;
	}
	if ((item.fields_mask & (1<<25)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a25\":";
		s << item.a25;
	}
	if ((item.fields_mask & (1<<26)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a26\":";
		s << item.a26;
	}
	if ((item.fields_mask & (1<<27)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a27\":";
		s << item.a27;
	}
	if ((item.fields_mask & (1<<28)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a28\":";
		s << item.a28;
	}
	if ((item.fields_mask & (1<<29)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a29\":";
		s << item.a29;
	}
	if ((item.fields_mask & (1<<30)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a30\":";
		s << item.a30;
	}
	if ((item.fields_mask & (1<<31)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"a31\":";
		s << item.a31;
	}
	s << "}";
	return true;
}

bool tl2::details::TasksFullFilledCronRead(::basictl::tl_istream & s, ::tl2::tasks::FullFilledCron& item) {
	if (!s.nat_read(item.fields_mask)) { return false; }
	if ((item.fields_mask & (1<<0)) != 0) {
		if (!s.int_read(item.a0)) { return false; }
	} else {
			item.a0 = 0;
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (!s.int_read(item.a1)) { return false; }
	} else {
			item.a1 = 0;
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (!s.int_read(item.a2)) { return false; }
	} else {
			item.a2 = 0;
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (!s.int_read(item.a3)) { return false; }
	} else {
			item.a3 = 0;
	}
	if ((item.fields_mask & (1<<4)) != 0) {
		if (!s.int_read(item.a4)) { return false; }
	} else {
			item.a4 = 0;
	}
	if ((item.fields_mask & (1<<5)) != 0) {
		if (!s.int_read(item.a5)) { return false; }
	} else {
			item.a5 = 0;
	}
	if ((item.fields_mask & (1<<6)) != 0) {
		if (!s.int_read(item.a6)) { return false; }
	} else {
			item.a6 = 0;
	}
	if ((item.fields_mask & (1<<7)) != 0) {
		if (!s.int_read(item.a7)) { return false; }
	} else {
			item.a7 = 0;
	}
	if ((item.fields_mask & (1<<8)) != 0) {
		if (!s.int_read(item.a8)) { return false; }
	} else {
			item.a8 = 0;
	}
	if ((item.fields_mask & (1<<9)) != 0) {
		if (!s.int_read(item.a9)) { return false; }
	} else {
			item.a9 = 0;
	}
	if ((item.fields_mask & (1<<10)) != 0) {
		if (!s.int_read(item.a10)) { return false; }
	} else {
			item.a10 = 0;
	}
	if ((item.fields_mask & (1<<11)) != 0) {
		if (!s.int_read(item.a11)) { return false; }
	} else {
			item.a11 = 0;
	}
	if ((item.fields_mask & (1<<12)) != 0) {
		if (!s.int_read(item.a12)) { return false; }
	} else {
			item.a12 = 0;
	}
	if ((item.fields_mask & (1<<13)) != 0) {
		if (!s.int_read(item.a13)) { return false; }
	} else {
			item.a13 = 0;
	}
	if ((item.fields_mask & (1<<14)) != 0) {
		if (!s.int_read(item.a14)) { return false; }
	} else {
			item.a14 = 0;
	}
	if ((item.fields_mask & (1<<15)) != 0) {
		if (!s.int_read(item.a15)) { return false; }
	} else {
			item.a15 = 0;
	}
	if ((item.fields_mask & (1<<16)) != 0) {
		if (!s.int_read(item.a16)) { return false; }
	} else {
			item.a16 = 0;
	}
	if ((item.fields_mask & (1<<17)) != 0) {
		if (!s.int_read(item.a17)) { return false; }
	} else {
			item.a17 = 0;
	}
	if ((item.fields_mask & (1<<18)) != 0) {
		if (!s.int_read(item.a18)) { return false; }
	} else {
			item.a18 = 0;
	}
	if ((item.fields_mask & (1<<19)) != 0) {
		if (!s.int_read(item.a19)) { return false; }
	} else {
			item.a19 = 0;
	}
	if ((item.fields_mask & (1<<20)) != 0) {
		if (!s.int_read(item.a20)) { return false; }
	} else {
			item.a20 = 0;
	}
	if ((item.fields_mask & (1<<21)) != 0) {
		if (!s.int_read(item.a21)) { return false; }
	} else {
			item.a21 = 0;
	}
	if ((item.fields_mask & (1<<22)) != 0) {
		if (!s.int_read(item.a22)) { return false; }
	} else {
			item.a22 = 0;
	}
	if ((item.fields_mask & (1<<23)) != 0) {
		if (!s.int_read(item.a23)) { return false; }
	} else {
			item.a23 = 0;
	}
	if ((item.fields_mask & (1<<24)) != 0) {
		if (!s.int_read(item.a24)) { return false; }
	} else {
			item.a24 = 0;
	}
	if ((item.fields_mask & (1<<25)) != 0) {
		if (!s.int_read(item.a25)) { return false; }
	} else {
			item.a25 = 0;
	}
	if ((item.fields_mask & (1<<26)) != 0) {
		if (!s.int_read(item.a26)) { return false; }
	} else {
			item.a26 = 0;
	}
	if ((item.fields_mask & (1<<27)) != 0) {
		if (!s.int_read(item.a27)) { return false; }
	} else {
			item.a27 = 0;
	}
	if ((item.fields_mask & (1<<28)) != 0) {
		if (!s.int_read(item.a28)) { return false; }
	} else {
			item.a28 = 0;
	}
	if ((item.fields_mask & (1<<29)) != 0) {
		if (!s.int_read(item.a29)) { return false; }
	} else {
			item.a29 = 0;
	}
	if ((item.fields_mask & (1<<30)) != 0) {
		if (!s.int_read(item.a30)) { return false; }
	} else {
			item.a30 = 0;
	}
	if ((item.fields_mask & (1<<31)) != 0) {
		if (!s.int_read(item.a31)) { return false; }
	} else {
			item.a31 = 0;
	}
	return true;
}

bool tl2::details::TasksFullFilledCronWrite(::basictl::tl_ostream & s, const ::tl2::tasks::FullFilledCron& item) {
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!s.int_write(item.a0)) { return false;}
	}
	if ((item.fields_mask & (1<<1)) != 0) {
			if (!s.int_write(item.a1)) { return false;}
	}
	if ((item.fields_mask & (1<<2)) != 0) {
			if (!s.int_write(item.a2)) { return false;}
	}
	if ((item.fields_mask & (1<<3)) != 0) {
			if (!s.int_write(item.a3)) { return false;}
	}
	if ((item.fields_mask & (1<<4)) != 0) {
			if (!s.int_write(item.a4)) { return false;}
	}
	if ((item.fields_mask & (1<<5)) != 0) {
			if (!s.int_write(item.a5)) { return false;}
	}
	if ((item.fields_mask & (1<<6)) != 0) {
			if (!s.int_write(item.a6)) { return false;}
	}
	if ((item.fields_mask & (1<<7)) != 0) {
			if (!s.int_write(item.a7)) { return false;}
	}
	if ((item.fields_mask & (1<<8)) != 0) {
			if (!s.int_write(item.a8)) { return false;}
	}
	if ((item.fields_mask & (1<<9)) != 0) {
			if (!s.int_write(item.a9)) { return false;}
	}
	if ((item.fields_mask & (1<<10)) != 0) {
			if (!s.int_write(item.a10)) { return false;}
	}
	if ((item.fields_mask & (1<<11)) != 0) {
			if (!s.int_write(item.a11)) { return false;}
	}
	if ((item.fields_mask & (1<<12)) != 0) {
			if (!s.int_write(item.a12)) { return false;}
	}
	if ((item.fields_mask & (1<<13)) != 0) {
			if (!s.int_write(item.a13)) { return false;}
	}
	if ((item.fields_mask & (1<<14)) != 0) {
			if (!s.int_write(item.a14)) { return false;}
	}
	if ((item.fields_mask & (1<<15)) != 0) {
			if (!s.int_write(item.a15)) { return false;}
	}
	if ((item.fields_mask & (1<<16)) != 0) {
			if (!s.int_write(item.a16)) { return false;}
	}
	if ((item.fields_mask & (1<<17)) != 0) {
			if (!s.int_write(item.a17)) { return false;}
	}
	if ((item.fields_mask & (1<<18)) != 0) {
			if (!s.int_write(item.a18)) { return false;}
	}
	if ((item.fields_mask & (1<<19)) != 0) {
			if (!s.int_write(item.a19)) { return false;}
	}
	if ((item.fields_mask & (1<<20)) != 0) {
			if (!s.int_write(item.a20)) { return false;}
	}
	if ((item.fields_mask & (1<<21)) != 0) {
			if (!s.int_write(item.a21)) { return false;}
	}
	if ((item.fields_mask & (1<<22)) != 0) {
			if (!s.int_write(item.a22)) { return false;}
	}
	if ((item.fields_mask & (1<<23)) != 0) {
			if (!s.int_write(item.a23)) { return false;}
	}
	if ((item.fields_mask & (1<<24)) != 0) {
			if (!s.int_write(item.a24)) { return false;}
	}
	if ((item.fields_mask & (1<<25)) != 0) {
			if (!s.int_write(item.a25)) { return false;}
	}
	if ((item.fields_mask & (1<<26)) != 0) {
			if (!s.int_write(item.a26)) { return false;}
	}
	if ((item.fields_mask & (1<<27)) != 0) {
			if (!s.int_write(item.a27)) { return false;}
	}
	if ((item.fields_mask & (1<<28)) != 0) {
			if (!s.int_write(item.a28)) { return false;}
	}
	if ((item.fields_mask & (1<<29)) != 0) {
			if (!s.int_write(item.a29)) { return false;}
	}
	if ((item.fields_mask & (1<<30)) != 0) {
			if (!s.int_write(item.a30)) { return false;}
	}
	if ((item.fields_mask & (1<<31)) != 0) {
			if (!s.int_write(item.a31)) { return false;}
	}
	return true;
}

bool tl2::details::TasksFullFilledCronReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::FullFilledCron& item) {
	if (!s.nat_read_exact_tag(0xd4177d7e)) { return false; }
	return tl2::details::TasksFullFilledCronRead(s, item);
}

bool tl2::details::TasksFullFilledCronWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::FullFilledCron& item) {
	if (!s.nat_write(0xd4177d7e)) { return false; }
	return tl2::details::TasksFullFilledCronWrite(s, item);
}

bool tl2::tasks::GetAnyTask::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksGetAnyTaskWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::GetAnyTask::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksGetAnyTaskRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::GetAnyTask::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksGetAnyTaskWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::GetAnyTask::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::GetAnyTask::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::GetAnyTask::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksGetAnyTaskReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::GetAnyTask::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksGetAnyTaskWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::GetAnyTask::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::GetAnyTask::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksGetAnyTaskReset(::tl2::tasks::GetAnyTask& item) {
}

bool tl2::details::TasksGetAnyTaskWriteJSON(std::ostream& s, const ::tl2::tasks::GetAnyTask& item) {
	s << "true";
	return true;
}

bool tl2::details::TasksGetAnyTaskRead(::basictl::tl_istream & s, ::tl2::tasks::GetAnyTask& item) {
	return true;
}

bool tl2::details::TasksGetAnyTaskWrite(::basictl::tl_ostream & s, const ::tl2::tasks::GetAnyTask& item) {
	return true;
}

bool tl2::details::TasksGetAnyTaskReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::GetAnyTask& item) {
	if (!s.nat_read_exact_tag(0x4a9c7dbb)) { return false; }
	return tl2::details::TasksGetAnyTaskRead(s, item);
}

bool tl2::details::TasksGetAnyTaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::GetAnyTask& item) {
	if (!s.nat_write(0x4a9c7dbb)) { return false; }
	return tl2::details::TasksGetAnyTaskWrite(s, item);
}

bool tl2::details::TasksGetAnyTaskReadResult(::basictl::tl_istream & s, tl2::tasks::GetAnyTask& item, std::optional<::tl2::tasks::TaskInfo>& result) {
	if (!::tl2::details::TasksTaskInfoMaybeReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::TasksGetAnyTaskWriteResult(::basictl::tl_ostream & s, tl2::tasks::GetAnyTask& item, std::optional<::tl2::tasks::TaskInfo>& result) {
	if (!::tl2::details::TasksTaskInfoMaybeWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::tasks::GetAnyTask::read_result(::basictl::tl_istream & s, std::optional<::tl2::tasks::TaskInfo> & result) noexcept {
	bool success = tl2::details::TasksGetAnyTaskReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::tasks::GetAnyTask::write_result(::basictl::tl_ostream & s, std::optional<::tl2::tasks::TaskInfo> & result) noexcept {
	bool success = tl2::details::TasksGetAnyTaskWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::tasks::GetAnyTask::read_result_or_throw(::basictl::tl_throwable_istream & s, std::optional<::tl2::tasks::TaskInfo> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::tasks::GetAnyTask::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::optional<::tl2::tasks::TaskInfo> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::tasks::GetQueueSize::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksGetQueueSizeWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::GetQueueSize::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksGetQueueSizeRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::GetQueueSize::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksGetQueueSizeWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::GetQueueSize::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::GetQueueSize::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::GetQueueSize::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksGetQueueSizeReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::GetQueueSize::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksGetQueueSizeWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::GetQueueSize::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::GetQueueSize::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksGetQueueSizeReset(::tl2::tasks::GetQueueSize& item) {
	item.type_name.clear();
	item.queue_id.clear();
	item.fields_mask = 0;
	item.local_dep = 0;
}

bool tl2::details::TasksGetQueueSizeWriteJSON(std::ostream& s, const ::tl2::tasks::GetQueueSize& item) {
	auto add_comma = false;
	s << "{";
	if (item.type_name.size() != 0) {
		add_comma = true;
		s << "\"type_name\":";
		s << "\"" << item.type_name << "\"";
	}
	if (item.queue_id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"queue_id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.queue_id)) { return false; }
	}
	if (item.fields_mask != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"fields_mask\":";
		s << item.fields_mask;
	}
	if ((item.fields_mask & (1<<4)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"local_dep\":";
		s << item.local_dep;
	}
	s << "}";
	return true;
}

bool tl2::details::TasksGetQueueSizeRead(::basictl::tl_istream & s, ::tl2::tasks::GetQueueSize& item) {
	if (!s.string_read(item.type_name)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.queue_id)) { return false; }
	if (!s.nat_read(item.fields_mask)) { return false; }
	if ((item.fields_mask & (1<<4)) != 0) {
		if (!s.int_read(item.local_dep)) { return false; }
	} else {
			item.local_dep = 0;
	}
	return true;
}

bool tl2::details::TasksGetQueueSizeWrite(::basictl::tl_ostream & s, const ::tl2::tasks::GetQueueSize& item) {
	if (!s.string_write(item.type_name)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.queue_id)) { return false; }
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<4)) != 0) {
			if (!s.int_write(item.local_dep)) { return false;}
	}
	return true;
}

bool tl2::details::TasksGetQueueSizeReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::GetQueueSize& item) {
	if (!s.nat_read_exact_tag(0x6abbb057)) { return false; }
	return tl2::details::TasksGetQueueSizeRead(s, item);
}

bool tl2::details::TasksGetQueueSizeWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::GetQueueSize& item) {
	if (!s.nat_write(0x6abbb057)) { return false; }
	return tl2::details::TasksGetQueueSizeWrite(s, item);
}

bool tl2::details::TasksGetQueueSizeReadResult(::basictl::tl_istream & s, tl2::tasks::GetQueueSize& item, ::tl2::tasks::QueueStats& result) {
	if (!::tl2::details::TasksQueueStatsReadBoxed(s, result, item.fields_mask)) { return false; }
	return true;
}
bool tl2::details::TasksGetQueueSizeWriteResult(::basictl::tl_ostream & s, tl2::tasks::GetQueueSize& item, ::tl2::tasks::QueueStats& result) {
	if (!::tl2::details::TasksQueueStatsWriteBoxed(s, result, item.fields_mask)) { return false; }
	return true;
}

bool tl2::tasks::GetQueueSize::read_result(::basictl::tl_istream & s, ::tl2::tasks::QueueStats & result) noexcept {
	bool success = tl2::details::TasksGetQueueSizeReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::tasks::GetQueueSize::write_result(::basictl::tl_ostream & s, ::tl2::tasks::QueueStats & result) noexcept {
	bool success = tl2::details::TasksGetQueueSizeWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::tasks::GetQueueSize::read_result_or_throw(::basictl::tl_throwable_istream & s, ::tl2::tasks::QueueStats & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::tasks::GetQueueSize::write_result_or_throw(::basictl::tl_throwable_ostream & s, ::tl2::tasks::QueueStats & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::tasks::GetQueueTypes::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksGetQueueTypesWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::GetQueueTypes::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksGetQueueTypesRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::GetQueueTypes::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksGetQueueTypesWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::GetQueueTypes::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::GetQueueTypes::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::GetQueueTypes::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksGetQueueTypesReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::GetQueueTypes::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksGetQueueTypesWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::GetQueueTypes::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::GetQueueTypes::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksGetQueueTypesReset(::tl2::tasks::GetQueueTypes& item) {
	item.settings_mask = 0;
	item.stats_mask = 0;
}

bool tl2::details::TasksGetQueueTypesWriteJSON(std::ostream& s, const ::tl2::tasks::GetQueueTypes& item) {
	auto add_comma = false;
	s << "{";
	if (item.settings_mask != 0) {
		add_comma = true;
		s << "\"settings_mask\":";
		s << item.settings_mask;
	}
	if (item.stats_mask != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"stats_mask\":";
		s << item.stats_mask;
	}
	s << "}";
	return true;
}

bool tl2::details::TasksGetQueueTypesRead(::basictl::tl_istream & s, ::tl2::tasks::GetQueueTypes& item) {
	if (!s.nat_read(item.settings_mask)) { return false; }
	if (!s.nat_read(item.stats_mask)) { return false; }
	return true;
}

bool tl2::details::TasksGetQueueTypesWrite(::basictl::tl_ostream & s, const ::tl2::tasks::GetQueueTypes& item) {
	if (!s.nat_write(item.settings_mask)) { return false;}
	if (!s.nat_write(item.stats_mask)) { return false;}
	return true;
}

bool tl2::details::TasksGetQueueTypesReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::GetQueueTypes& item) {
	if (!s.nat_read_exact_tag(0x5434457a)) { return false; }
	return tl2::details::TasksGetQueueTypesRead(s, item);
}

bool tl2::details::TasksGetQueueTypesWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::GetQueueTypes& item) {
	if (!s.nat_write(0x5434457a)) { return false; }
	return tl2::details::TasksGetQueueTypesWrite(s, item);
}

bool tl2::details::TasksGetQueueTypesReadResult(::basictl::tl_istream & s, tl2::tasks::GetQueueTypes& item, std::vector<::tl2::tasks::QueueTypeInfo>& result) {
	if (!s.nat_read_exact_tag(0x1cb5c415)) { return false;}
	if (!::tl2::details::BuiltinVectorTasksQueueTypeInfoRead(s, result)) { return false; }
	return true;
}
bool tl2::details::TasksGetQueueTypesWriteResult(::basictl::tl_ostream & s, tl2::tasks::GetQueueTypes& item, std::vector<::tl2::tasks::QueueTypeInfo>& result) {
	if (!s.nat_write(0x1cb5c415)) { return false; }
	if (!::tl2::details::BuiltinVectorTasksQueueTypeInfoWrite(s, result)) { return false; }
	return true;
}

bool tl2::tasks::GetQueueTypes::read_result(::basictl::tl_istream & s, std::vector<::tl2::tasks::QueueTypeInfo> & result) noexcept {
	bool success = tl2::details::TasksGetQueueTypesReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::tasks::GetQueueTypes::write_result(::basictl::tl_ostream & s, std::vector<::tl2::tasks::QueueTypeInfo> & result) noexcept {
	bool success = tl2::details::TasksGetQueueTypesWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::tasks::GetQueueTypes::read_result_or_throw(::basictl::tl_throwable_istream & s, std::vector<::tl2::tasks::QueueTypeInfo> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::tasks::GetQueueTypes::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::vector<::tl2::tasks::QueueTypeInfo> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::tasks::GetTaskFromQueue::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksGetTaskFromQueueWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::GetTaskFromQueue::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksGetTaskFromQueueRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::GetTaskFromQueue::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksGetTaskFromQueueWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::GetTaskFromQueue::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::GetTaskFromQueue::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::GetTaskFromQueue::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksGetTaskFromQueueReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::GetTaskFromQueue::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksGetTaskFromQueueWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::GetTaskFromQueue::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::GetTaskFromQueue::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksGetTaskFromQueueReset(::tl2::tasks::GetTaskFromQueue& item) {
	item.type_name.clear();
	item.queue_id.clear();
}

bool tl2::details::TasksGetTaskFromQueueWriteJSON(std::ostream& s, const ::tl2::tasks::GetTaskFromQueue& item) {
	auto add_comma = false;
	s << "{";
	if (item.type_name.size() != 0) {
		add_comma = true;
		s << "\"type_name\":";
		s << "\"" << item.type_name << "\"";
	}
	if (item.queue_id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"queue_id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.queue_id)) { return false; }
	}
	s << "}";
	return true;
}

bool tl2::details::TasksGetTaskFromQueueRead(::basictl::tl_istream & s, ::tl2::tasks::GetTaskFromQueue& item) {
	if (!s.string_read(item.type_name)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.queue_id)) { return false; }
	return true;
}

bool tl2::details::TasksGetTaskFromQueueWrite(::basictl::tl_ostream & s, const ::tl2::tasks::GetTaskFromQueue& item) {
	if (!s.string_write(item.type_name)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.queue_id)) { return false; }
	return true;
}

bool tl2::details::TasksGetTaskFromQueueReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::GetTaskFromQueue& item) {
	if (!s.nat_read_exact_tag(0x6a52b698)) { return false; }
	return tl2::details::TasksGetTaskFromQueueRead(s, item);
}

bool tl2::details::TasksGetTaskFromQueueWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::GetTaskFromQueue& item) {
	if (!s.nat_write(0x6a52b698)) { return false; }
	return tl2::details::TasksGetTaskFromQueueWrite(s, item);
}

bool tl2::details::TasksGetTaskFromQueueReadResult(::basictl::tl_istream & s, tl2::tasks::GetTaskFromQueue& item, std::optional<::tl2::tasks::TaskInfo>& result) {
	if (!::tl2::details::TasksTaskInfoMaybeReadBoxed(s, result)) { return false; }
	return true;
}
bool tl2::details::TasksGetTaskFromQueueWriteResult(::basictl::tl_ostream & s, tl2::tasks::GetTaskFromQueue& item, std::optional<::tl2::tasks::TaskInfo>& result) {
	if (!::tl2::details::TasksTaskInfoMaybeWriteBoxed(s, result)) { return false; }
	return true;
}

bool tl2::tasks::GetTaskFromQueue::read_result(::basictl::tl_istream & s, std::optional<::tl2::tasks::TaskInfo> & result) noexcept {
	bool success = tl2::details::TasksGetTaskFromQueueReadResult(s, *this, result);
	s.last_release();
	return success;
}
bool tl2::tasks::GetTaskFromQueue::write_result(::basictl::tl_ostream & s, std::optional<::tl2::tasks::TaskInfo> & result) noexcept {
	bool success = tl2::details::TasksGetTaskFromQueueWriteResult(s, *this, result);
	s.last_release();
	return success;
}

void tl2::tasks::GetTaskFromQueue::read_result_or_throw(::basictl::tl_throwable_istream & s, std::optional<::tl2::tasks::TaskInfo> & result) {
	::basictl::tl_istream s2(s);
	this->read_result(s2, result);
	s2.pass_data(s);
}
void tl2::tasks::GetTaskFromQueue::write_result_or_throw(::basictl::tl_throwable_ostream & s, std::optional<::tl2::tasks::TaskInfo> & result) {
	::basictl::tl_ostream s2(s);
	this->write_result(s2, result);
	s2.pass_data(s);
}

bool tl2::tasks::QueueStats::write_json(std::ostream& s, uint32_t nat_fields_mask)const {
	if (!::tl2::details::TasksQueueStatsWriteJSON(s, *this, nat_fields_mask)) { return false; }
	return true;
}

bool tl2::tasks::QueueStats::read(::basictl::tl_istream & s, uint32_t nat_fields_mask) noexcept {
	if (!::tl2::details::TasksQueueStatsRead(s, *this, nat_fields_mask)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::QueueStats::write(::basictl::tl_ostream & s, uint32_t nat_fields_mask)const noexcept {
	if (!::tl2::details::TasksQueueStatsWrite(s, *this, nat_fields_mask)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::QueueStats::read_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_fields_mask) {
	::basictl::tl_istream s2(s);
	this->read(s2, nat_fields_mask);
	s2.pass_data(s);
}

void tl2::tasks::QueueStats::write_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_fields_mask)const {
	::basictl::tl_ostream s2(s);
	this->write(s2, nat_fields_mask);
	s2.pass_data(s);
}

bool tl2::tasks::QueueStats::read_boxed(::basictl::tl_istream & s, uint32_t nat_fields_mask) noexcept {
	if (!::tl2::details::TasksQueueStatsReadBoxed(s, *this, nat_fields_mask)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::QueueStats::write_boxed(::basictl::tl_ostream & s, uint32_t nat_fields_mask)const noexcept {
	if (!::tl2::details::TasksQueueStatsWriteBoxed(s, *this, nat_fields_mask)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::QueueStats::read_boxed_or_throw(::basictl::tl_throwable_istream & s, uint32_t nat_fields_mask) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2, nat_fields_mask);
	s2.pass_data(s);
}

void tl2::tasks::QueueStats::write_boxed_or_throw(::basictl::tl_throwable_ostream & s, uint32_t nat_fields_mask)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2, nat_fields_mask);
	s2.pass_data(s);
}

void tl2::details::TasksQueueStatsReset(::tl2::tasks::QueueStats& item) {
	item.waiting_size = 0;
	item.scheduled_size = 0;
	item.in_progress_size = 0;
}

bool tl2::details::TasksQueueStatsWriteJSON(std::ostream& s, const ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask) {
	auto add_comma = false;
	s << "{";
	if ((nat_fields_mask & (1<<0)) != 0) {
		add_comma = true;
		s << "\"waiting_size\":";
		s << item.waiting_size;
	}
	if ((nat_fields_mask & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"scheduled_size\":";
		s << item.scheduled_size;
	}
	if ((nat_fields_mask & (1<<2)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"in_progress_size\":";
		s << item.in_progress_size;
	}
	s << "}";
	return true;
}

bool tl2::details::TasksQueueStatsRead(::basictl::tl_istream & s, ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask) {
	if ((nat_fields_mask & (1<<0)) != 0) {
		if (!s.int_read(item.waiting_size)) { return false; }
	} else {
			item.waiting_size = 0;
	}
	if ((nat_fields_mask & (1<<1)) != 0) {
		if (!s.int_read(item.scheduled_size)) { return false; }
	} else {
			item.scheduled_size = 0;
	}
	if ((nat_fields_mask & (1<<2)) != 0) {
		if (!s.int_read(item.in_progress_size)) { return false; }
	} else {
			item.in_progress_size = 0;
	}
	return true;
}

bool tl2::details::TasksQueueStatsWrite(::basictl::tl_ostream & s, const ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask) {
	if ((nat_fields_mask & (1<<0)) != 0) {
			if (!s.int_write(item.waiting_size)) { return false;}
	}
	if ((nat_fields_mask & (1<<1)) != 0) {
			if (!s.int_write(item.scheduled_size)) { return false;}
	}
	if ((nat_fields_mask & (1<<2)) != 0) {
			if (!s.int_write(item.in_progress_size)) { return false;}
	}
	return true;
}

bool tl2::details::TasksQueueStatsReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask) {
	if (!s.nat_read_exact_tag(0x1d942543)) { return false; }
	return tl2::details::TasksQueueStatsRead(s, item, nat_fields_mask);
}

bool tl2::details::TasksQueueStatsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::QueueStats& item, uint32_t nat_fields_mask) {
	if (!s.nat_write(0x1d942543)) { return false; }
	return tl2::details::TasksQueueStatsWrite(s, item, nat_fields_mask);
}

bool tl2::tasks::QueueTypeInfo::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksQueueTypeInfoWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::QueueTypeInfo::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksQueueTypeInfoRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::QueueTypeInfo::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksQueueTypeInfoWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::QueueTypeInfo::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::QueueTypeInfo::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::QueueTypeInfo::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksQueueTypeInfoReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::QueueTypeInfo::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksQueueTypeInfoWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::QueueTypeInfo::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::QueueTypeInfo::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksQueueTypeInfoReset(::tl2::tasks::QueueTypeInfo& item) {
	item.type_name.clear();
	::tl2::details::TasksQueueTypeSettingsReset(item.settings);
	::tl2::details::TasksQueueTypeStatsReset(item.stats);
}

bool tl2::details::TasksQueueTypeInfoWriteJSON(std::ostream& s, const ::tl2::tasks::QueueTypeInfo& item) {
	auto add_comma = false;
	s << "{";
	if (item.type_name.size() != 0) {
		add_comma = true;
		s << "\"type_name\":";
		s << "\"" << item.type_name << "\"";
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"settings\":";
	if (!::tl2::details::TasksQueueTypeSettingsWriteJSON(s, item.settings)) { return false; }
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"stats\":";
	if (!::tl2::details::TasksQueueTypeStatsWriteJSON(s, item.stats)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::TasksQueueTypeInfoRead(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeInfo& item) {
	if (!s.string_read(item.type_name)) { return false; }
	if (!::tl2::details::TasksQueueTypeSettingsRead(s, item.settings)) { return false; }
	if (!::tl2::details::TasksQueueTypeStatsRead(s, item.stats)) { return false; }
	return true;
}

bool tl2::details::TasksQueueTypeInfoWrite(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeInfo& item) {
	if (!s.string_write(item.type_name)) { return false;}
	if (!::tl2::details::TasksQueueTypeSettingsWrite(s, item.settings)) { return false; }
	if (!::tl2::details::TasksQueueTypeStatsWrite(s, item.stats)) { return false; }
	return true;
}

bool tl2::details::TasksQueueTypeInfoReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeInfo& item) {
	if (!s.nat_read_exact_tag(0x38d38d3e)) { return false; }
	return tl2::details::TasksQueueTypeInfoRead(s, item);
}

bool tl2::details::TasksQueueTypeInfoWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeInfo& item) {
	if (!s.nat_write(0x38d38d3e)) { return false; }
	return tl2::details::TasksQueueTypeInfoWrite(s, item);
}

bool tl2::tasks::QueueTypeSettings::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksQueueTypeSettingsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::QueueTypeSettings::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksQueueTypeSettingsRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::QueueTypeSettings::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksQueueTypeSettingsWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::QueueTypeSettings::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::QueueTypeSettings::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::QueueTypeSettings::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksQueueTypeSettingsReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::QueueTypeSettings::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksQueueTypeSettingsWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::QueueTypeSettings::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::QueueTypeSettings::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksQueueTypeSettingsReset(::tl2::tasks::QueueTypeSettings& item) {
	item.fields_mask = 0;
	item.is_enabled = false;
	item.is_persistent = false;
	item.priority = 0;
	item.default_retry_time = 0;
	item.default_retry_num = 0;
	item.move_to_queue_type_on_error.clear();
	item.is_blocking = false;
	item.timelimit = 0;
	item.max_queue_size = 0;
}

bool tl2::details::TasksQueueTypeSettingsWriteJSON(std::ostream& s, const ::tl2::tasks::QueueTypeSettings& item) {
	auto add_comma = false;
	s << "{";
	if (item.fields_mask != 0) {
		add_comma = true;
		s << "\"fields_mask\":";
		s << item.fields_mask;
	}
	if ((item.fields_mask & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"is_enabled\":";
		if (!::tl2::details::BoolWriteJSON(s, item.is_enabled)) { return false; }
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"is_persistent\":";
		if (!::tl2::details::BoolWriteJSON(s, item.is_persistent)) { return false; }
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"priority\":";
		s << item.priority;
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"default_retry_time\":";
		s << item.default_retry_time;
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"default_retry_num\":";
		s << item.default_retry_num;
	}
	if ((item.fields_mask & (1<<4)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"move_to_queue_type_on_error\":";
		s << "\"" << item.move_to_queue_type_on_error << "\"";
	}
	if ((item.fields_mask & (1<<5)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"is_blocking\":";
		if (!::tl2::details::BoolWriteJSON(s, item.is_blocking)) { return false; }
	}
	if ((item.fields_mask & (1<<6)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"timelimit\":";
		s << item.timelimit;
	}
	if ((item.fields_mask & (1<<7)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"max_queue_size\":";
		s << item.max_queue_size;
	}
	s << "}";
	return true;
}

bool tl2::details::TasksQueueTypeSettingsRead(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeSettings& item) {
	if (!s.nat_read(item.fields_mask)) { return false; }
	if ((item.fields_mask & (1<<0)) != 0) {
		if (!::tl2::details::BoolReadBoxed(s, item.is_enabled)) { return false; }
	} else {
			item.is_enabled = false;
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (!::tl2::details::BoolReadBoxed(s, item.is_persistent)) { return false; }
	} else {
			item.is_persistent = false;
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (!s.int_read(item.priority)) { return false; }
	} else {
			item.priority = 0;
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (!s.int_read(item.default_retry_time)) { return false; }
	} else {
			item.default_retry_time = 0;
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (!s.int_read(item.default_retry_num)) { return false; }
	} else {
			item.default_retry_num = 0;
	}
	if ((item.fields_mask & (1<<4)) != 0) {
		if (!s.string_read(item.move_to_queue_type_on_error)) { return false; }
	} else {
			item.move_to_queue_type_on_error.clear();
	}
	if ((item.fields_mask & (1<<5)) != 0) {
		if (!::tl2::details::BoolReadBoxed(s, item.is_blocking)) { return false; }
	} else {
			item.is_blocking = false;
	}
	if ((item.fields_mask & (1<<6)) != 0) {
		if (!s.int_read(item.timelimit)) { return false; }
	} else {
			item.timelimit = 0;
	}
	if ((item.fields_mask & (1<<7)) != 0) {
		if (!s.int_read(item.max_queue_size)) { return false; }
	} else {
			item.max_queue_size = 0;
	}
	return true;
}

bool tl2::details::TasksQueueTypeSettingsWrite(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeSettings& item) {
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!::tl2::details::BoolWriteBoxed(s, item.is_enabled)) { return false; }
	}
	if ((item.fields_mask & (1<<1)) != 0) {
			if (!::tl2::details::BoolWriteBoxed(s, item.is_persistent)) { return false; }
	}
	if ((item.fields_mask & (1<<2)) != 0) {
			if (!s.int_write(item.priority)) { return false;}
	}
	if ((item.fields_mask & (1<<3)) != 0) {
			if (!s.int_write(item.default_retry_time)) { return false;}
	}
	if ((item.fields_mask & (1<<3)) != 0) {
			if (!s.int_write(item.default_retry_num)) { return false;}
	}
	if ((item.fields_mask & (1<<4)) != 0) {
			if (!s.string_write(item.move_to_queue_type_on_error)) { return false;}
	}
	if ((item.fields_mask & (1<<5)) != 0) {
			if (!::tl2::details::BoolWriteBoxed(s, item.is_blocking)) { return false; }
	}
	if ((item.fields_mask & (1<<6)) != 0) {
			if (!s.int_write(item.timelimit)) { return false;}
	}
	if ((item.fields_mask & (1<<7)) != 0) {
			if (!s.int_write(item.max_queue_size)) { return false;}
	}
	return true;
}

bool tl2::details::TasksQueueTypeSettingsReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeSettings& item) {
	if (!s.nat_read_exact_tag(0x561fbc09)) { return false; }
	return tl2::details::TasksQueueTypeSettingsRead(s, item);
}

bool tl2::details::TasksQueueTypeSettingsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeSettings& item) {
	if (!s.nat_write(0x561fbc09)) { return false; }
	return tl2::details::TasksQueueTypeSettingsWrite(s, item);
}

bool tl2::tasks::QueueTypeStats::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksQueueTypeStatsWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::QueueTypeStats::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksQueueTypeStatsRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::QueueTypeStats::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksQueueTypeStatsWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::QueueTypeStats::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::QueueTypeStats::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::QueueTypeStats::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksQueueTypeStatsReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::QueueTypeStats::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksQueueTypeStatsWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::QueueTypeStats::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::QueueTypeStats::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksQueueTypeStatsReset(::tl2::tasks::QueueTypeStats& item) {
	item.fields_mask = 0;
	item.waiting_size = 0;
	item.scheduled_size = 0;
	item.in_progress_size = 0;
	item.num_queues = 0;
}

bool tl2::details::TasksQueueTypeStatsWriteJSON(std::ostream& s, const ::tl2::tasks::QueueTypeStats& item) {
	auto add_comma = false;
	s << "{";
	if (item.fields_mask != 0) {
		add_comma = true;
		s << "\"fields_mask\":";
		s << item.fields_mask;
	}
	if ((item.fields_mask & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"waiting_size\":";
		s << item.waiting_size;
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"scheduled_size\":";
		s << item.scheduled_size;
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"in_progress_size\":";
		s << item.in_progress_size;
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"num_queues\":";
		s << item.num_queues;
	}
	s << "}";
	return true;
}

bool tl2::details::TasksQueueTypeStatsRead(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeStats& item) {
	if (!s.nat_read(item.fields_mask)) { return false; }
	if ((item.fields_mask & (1<<0)) != 0) {
		if (!s.long_read(item.waiting_size)) { return false; }
	} else {
			item.waiting_size = 0;
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (!s.long_read(item.scheduled_size)) { return false; }
	} else {
			item.scheduled_size = 0;
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (!s.long_read(item.in_progress_size)) { return false; }
	} else {
			item.in_progress_size = 0;
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (!s.int_read(item.num_queues)) { return false; }
	} else {
			item.num_queues = 0;
	}
	return true;
}

bool tl2::details::TasksQueueTypeStatsWrite(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeStats& item) {
	if (!s.nat_write(item.fields_mask)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!s.long_write(item.waiting_size)) { return false;}
	}
	if ((item.fields_mask & (1<<1)) != 0) {
			if (!s.long_write(item.scheduled_size)) { return false;}
	}
	if ((item.fields_mask & (1<<2)) != 0) {
			if (!s.long_write(item.in_progress_size)) { return false;}
	}
	if ((item.fields_mask & (1<<3)) != 0) {
			if (!s.int_write(item.num_queues)) { return false;}
	}
	return true;
}

bool tl2::details::TasksQueueTypeStatsReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::QueueTypeStats& item) {
	if (!s.nat_read_exact_tag(0xe1b785f2)) { return false; }
	return tl2::details::TasksQueueTypeStatsRead(s, item);
}

bool tl2::details::TasksQueueTypeStatsWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::QueueTypeStats& item) {
	if (!s.nat_write(0xe1b785f2)) { return false; }
	return tl2::details::TasksQueueTypeStatsWrite(s, item);
}

bool tl2::tasks::Task::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksTaskWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::Task::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksTaskRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::Task::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksTaskWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::Task::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::Task::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::Task::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksTaskReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::Task::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksTaskWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::Task::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::Task::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksTaskReset(::tl2::tasks::Task& item) {
	item.fields_mask = 0;
	item.flags = 0;
	item.tag.clear();
	item.data.clear();
	item.id = 0;
	item.retries = 0;
	item.scheduled_time = 0;
	item.deadline = 0;
}

bool tl2::details::TasksTaskWriteJSON(std::ostream& s, const ::tl2::tasks::Task& item) {
	auto add_comma = false;
	s << "{";
	if (item.fields_mask != 0) {
		add_comma = true;
		s << "\"fields_mask\":";
		s << item.fields_mask;
	}
	if (item.flags != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"flags\":";
		s << item.flags;
	}
	if (item.tag.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"tag\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.tag)) { return false; }
	}
	if (item.data.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"data\":";
		s << "\"" << item.data << "\"";
	}
	if ((item.fields_mask & (1<<0)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"id\":";
		s << item.id;
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"retries\":";
		s << item.retries;
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"scheduled_time\":";
		s << item.scheduled_time;
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"deadline\":";
		s << item.deadline;
	}
	s << "}";
	return true;
}

bool tl2::details::TasksTaskRead(::basictl::tl_istream & s, ::tl2::tasks::Task& item) {
	if (!s.nat_read(item.fields_mask)) { return false; }
	if (!s.int_read(item.flags)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.tag)) { return false; }
	if (!s.string_read(item.data)) { return false; }
	if ((item.fields_mask & (1<<0)) != 0) {
		if (!s.long_read(item.id)) { return false; }
	} else {
			item.id = 0;
	}
	if ((item.fields_mask & (1<<1)) != 0) {
		if (!s.int_read(item.retries)) { return false; }
	} else {
			item.retries = 0;
	}
	if ((item.fields_mask & (1<<2)) != 0) {
		if (!s.int_read(item.scheduled_time)) { return false; }
	} else {
			item.scheduled_time = 0;
	}
	if ((item.fields_mask & (1<<3)) != 0) {
		if (!s.int_read(item.deadline)) { return false; }
	} else {
			item.deadline = 0;
	}
	return true;
}

bool tl2::details::TasksTaskWrite(::basictl::tl_ostream & s, const ::tl2::tasks::Task& item) {
	if (!s.nat_write(item.fields_mask)) { return false;}
	if (!s.int_write(item.flags)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.tag)) { return false; }
	if (!s.string_write(item.data)) { return false;}
	if ((item.fields_mask & (1<<0)) != 0) {
			if (!s.long_write(item.id)) { return false;}
	}
	if ((item.fields_mask & (1<<1)) != 0) {
			if (!s.int_write(item.retries)) { return false;}
	}
	if ((item.fields_mask & (1<<2)) != 0) {
			if (!s.int_write(item.scheduled_time)) { return false;}
	}
	if ((item.fields_mask & (1<<3)) != 0) {
			if (!s.int_write(item.deadline)) { return false;}
	}
	return true;
}

bool tl2::details::TasksTaskReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::Task& item) {
	if (!s.nat_read_exact_tag(0x7c23bc2c)) { return false; }
	return tl2::details::TasksTaskRead(s, item);
}

bool tl2::details::TasksTaskWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::Task& item) {
	if (!s.nat_write(0x7c23bc2c)) { return false; }
	return tl2::details::TasksTaskWrite(s, item);
}

bool tl2::tasks::TaskInfo::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksTaskInfoWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::TaskInfo::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksTaskInfoRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::TaskInfo::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksTaskInfoWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::TaskInfo::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::TaskInfo::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::TaskInfo::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksTaskInfoReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::TaskInfo::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksTaskInfoWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::TaskInfo::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::TaskInfo::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksTaskInfoReset(::tl2::tasks::TaskInfo& item) {
	item.type_name.clear();
	item.queue_id.clear();
	::tl2::details::TasksTaskReset(item.task);
}

bool tl2::details::TasksTaskInfoWriteJSON(std::ostream& s, const ::tl2::tasks::TaskInfo& item) {
	auto add_comma = false;
	s << "{";
	if (item.type_name.size() != 0) {
		add_comma = true;
		s << "\"type_name\":";
		s << "\"" << item.type_name << "\"";
	}
	if (item.queue_id.size() != 0) {
		if (add_comma) {
			s << ",";
		}
		add_comma = true;
		s << "\"queue_id\":";
		if (!::tl2::details::BuiltinVectorIntWriteJSON(s, item.queue_id)) { return false; }
	}
	if (add_comma) {
		s << ",";
	}
	add_comma = true;
	s << "\"task\":";
	if (!::tl2::details::TasksTaskWriteJSON(s, item.task)) { return false; }
	s << "}";
	return true;
}

bool tl2::details::TasksTaskInfoRead(::basictl::tl_istream & s, ::tl2::tasks::TaskInfo& item) {
	if (!s.string_read(item.type_name)) { return false; }
	if (!::tl2::details::BuiltinVectorIntRead(s, item.queue_id)) { return false; }
	if (!::tl2::details::TasksTaskRead(s, item.task)) { return false; }
	return true;
}

bool tl2::details::TasksTaskInfoWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskInfo& item) {
	if (!s.string_write(item.type_name)) { return false;}
	if (!::tl2::details::BuiltinVectorIntWrite(s, item.queue_id)) { return false; }
	if (!::tl2::details::TasksTaskWrite(s, item.task)) { return false; }
	return true;
}

bool tl2::details::TasksTaskInfoReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskInfo& item) {
	if (!s.nat_read_exact_tag(0x06f0c6a6)) { return false; }
	return tl2::details::TasksTaskInfoRead(s, item);
}

bool tl2::details::TasksTaskInfoWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskInfo& item) {
	if (!s.nat_write(0x06f0c6a6)) { return false; }
	return tl2::details::TasksTaskInfoWrite(s, item);
}

bool tl2::details::TasksTaskInfoMaybeWriteJSON(std::ostream & s, const std::optional<::tl2::tasks::TaskInfo>& item) {
	s << "{";
	if (item) {
		s << "\"ok\":true";
		s << ",\"value\":";
		if (!::tl2::details::TasksTaskInfoWriteJSON(s, *item)) { return false; }
	}
	s << "}";
	return true;
}
bool tl2::details::TasksTaskInfoMaybeReadBoxed(::basictl::tl_istream & s, std::optional<::tl2::tasks::TaskInfo>& item) {
	bool has_item = false;
	if (!s.bool_read(has_item, 0x27930a7b, 0x3f9c8ef8)) { return false; }
	if (has_item) {
		if (!item) {
			item.emplace();
		}
		if (!::tl2::details::TasksTaskInfoRead(s, *item)) { return false; }
		return true;
	}
	item.reset();
	return true;
}

bool tl2::details::TasksTaskInfoMaybeWriteBoxed(::basictl::tl_ostream & s, const std::optional<::tl2::tasks::TaskInfo>& item) {
	if (!s.nat_write(item ? 0x3f9c8ef8 : 0x27930a7b)) { return false; }
	if (item) {
		if (!::tl2::details::TasksTaskInfoWrite(s, *item)) { return false; }
	}
	return true;
}

static const std::string_view TasksTaskStatus_tbl_tl_name[]{"tasks.taskStatusNotCurrentlyInEngine", "tasks.taskStatusScheduled", "tasks.taskStatusWaiting", "tasks.taskStatusInProgress"};
static const uint32_t TasksTaskStatus_tbl_tl_tag[]{0xb207caaa, 0x0aca80a9, 0x16739c2c, 0x06ef70e7};

bool tl2::tasks::TaskStatus::write_json(std::ostream & s)const {
	if (!::tl2::details::TasksTaskStatusWriteJSON(s, *this)) { return false; }
	return true;
}
bool tl2::tasks::TaskStatus::read_boxed(::basictl::tl_istream & s) {
	if (!::tl2::details::TasksTaskStatusReadBoxed(s, *this)) { return false; }
	return true;
}
bool tl2::tasks::TaskStatus::write_boxed(::basictl::tl_ostream & s)const {
	if (!::tl2::details::TasksTaskStatusWriteBoxed(s, *this)) { return false; }
	return true;
}
std::string_view tl2::tasks::TaskStatus::tl_name() const {
	return TasksTaskStatus_tbl_tl_name[value.index()];
}
uint32_t tl2::tasks::TaskStatus::tl_tag() const {
	return TasksTaskStatus_tbl_tl_tag[value.index()];
}


void tl2::details::TasksTaskStatusReset(::tl2::tasks::TaskStatus& item) {
	item.value.emplace<0>(); // TODO - optimize, if already 0, call Reset function
}

bool tl2::details::TasksTaskStatusWriteJSON(std::ostream & s, const ::tl2::tasks::TaskStatus& item) {
	s << "\"" << TasksTaskStatus_tbl_tl_name[item.value.index()] << "\"";
	return true;
}
bool tl2::details::TasksTaskStatusReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatus& item) {
	uint32_t nat;
	s.nat_read(nat);
	switch (nat) {
	case 0xb207caaa:
		if (item.value.index() != 0) { item.value.emplace<0>(); }
		break;
	case 0x0aca80a9:
		if (item.value.index() != 1) { item.value.emplace<1>(); }
		break;
	case 0x16739c2c:
		if (item.value.index() != 2) { item.value.emplace<2>(); }
		break;
	case 0x06ef70e7:
		if (item.value.index() != 3) { item.value.emplace<3>(); }
		break;
	default:
		return s.set_error_union_tag();
    }
	return true;
}

bool tl2::details::TasksTaskStatusWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatus& item) {
	s.nat_write(TasksTaskStatus_tbl_tl_tag[item.value.index()]);
	switch (item.value.index()) {
	}
	return true;
}

bool tl2::tasks::TaskStatusInProgress::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksTaskStatusInProgressWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::TaskStatusInProgress::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksTaskStatusInProgressRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::TaskStatusInProgress::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksTaskStatusInProgressWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::TaskStatusInProgress::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::TaskStatusInProgress::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::TaskStatusInProgress::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksTaskStatusInProgressReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::TaskStatusInProgress::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksTaskStatusInProgressWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::TaskStatusInProgress::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::TaskStatusInProgress::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksTaskStatusInProgressReset(::tl2::tasks::TaskStatusInProgress& item) {
}

bool tl2::details::TasksTaskStatusInProgressWriteJSON(std::ostream& s, const ::tl2::tasks::TaskStatusInProgress& item) {
	s << "true";
	return true;
}

bool tl2::details::TasksTaskStatusInProgressRead(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusInProgress& item) {
	return true;
}

bool tl2::details::TasksTaskStatusInProgressWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusInProgress& item) {
	return true;
}

bool tl2::details::TasksTaskStatusInProgressReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusInProgress& item) {
	if (!s.nat_read_exact_tag(0x06ef70e7)) { return false; }
	return tl2::details::TasksTaskStatusInProgressRead(s, item);
}

bool tl2::details::TasksTaskStatusInProgressWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusInProgress& item) {
	if (!s.nat_write(0x06ef70e7)) { return false; }
	return tl2::details::TasksTaskStatusInProgressWrite(s, item);
}

bool tl2::tasks::TaskStatusNotCurrentlyInEngine::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksTaskStatusNotCurrentlyInEngineWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::TaskStatusNotCurrentlyInEngine::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksTaskStatusNotCurrentlyInEngineRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::TaskStatusNotCurrentlyInEngine::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksTaskStatusNotCurrentlyInEngineWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::TaskStatusNotCurrentlyInEngine::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::TaskStatusNotCurrentlyInEngine::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::TaskStatusNotCurrentlyInEngine::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksTaskStatusNotCurrentlyInEngineReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::TaskStatusNotCurrentlyInEngine::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksTaskStatusNotCurrentlyInEngineWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::TaskStatusNotCurrentlyInEngine::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::TaskStatusNotCurrentlyInEngine::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksTaskStatusNotCurrentlyInEngineReset(::tl2::tasks::TaskStatusNotCurrentlyInEngine& item) {
}

bool tl2::details::TasksTaskStatusNotCurrentlyInEngineWriteJSON(std::ostream& s, const ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item) {
	s << "true";
	return true;
}

bool tl2::details::TasksTaskStatusNotCurrentlyInEngineRead(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item) {
	return true;
}

bool tl2::details::TasksTaskStatusNotCurrentlyInEngineWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item) {
	return true;
}

bool tl2::details::TasksTaskStatusNotCurrentlyInEngineReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item) {
	if (!s.nat_read_exact_tag(0xb207caaa)) { return false; }
	return tl2::details::TasksTaskStatusNotCurrentlyInEngineRead(s, item);
}

bool tl2::details::TasksTaskStatusNotCurrentlyInEngineWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusNotCurrentlyInEngine& item) {
	if (!s.nat_write(0xb207caaa)) { return false; }
	return tl2::details::TasksTaskStatusNotCurrentlyInEngineWrite(s, item);
}

bool tl2::tasks::TaskStatusScheduled::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksTaskStatusScheduledWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::TaskStatusScheduled::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksTaskStatusScheduledRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::TaskStatusScheduled::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksTaskStatusScheduledWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::TaskStatusScheduled::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::TaskStatusScheduled::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::TaskStatusScheduled::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksTaskStatusScheduledReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::TaskStatusScheduled::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksTaskStatusScheduledWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::TaskStatusScheduled::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::TaskStatusScheduled::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksTaskStatusScheduledReset(::tl2::tasks::TaskStatusScheduled& item) {
}

bool tl2::details::TasksTaskStatusScheduledWriteJSON(std::ostream& s, const ::tl2::tasks::TaskStatusScheduled& item) {
	s << "true";
	return true;
}

bool tl2::details::TasksTaskStatusScheduledRead(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusScheduled& item) {
	return true;
}

bool tl2::details::TasksTaskStatusScheduledWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusScheduled& item) {
	return true;
}

bool tl2::details::TasksTaskStatusScheduledReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusScheduled& item) {
	if (!s.nat_read_exact_tag(0x0aca80a9)) { return false; }
	return tl2::details::TasksTaskStatusScheduledRead(s, item);
}

bool tl2::details::TasksTaskStatusScheduledWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusScheduled& item) {
	if (!s.nat_write(0x0aca80a9)) { return false; }
	return tl2::details::TasksTaskStatusScheduledWrite(s, item);
}

bool tl2::tasks::TaskStatusWaiting::write_json(std::ostream& s)const {
	if (!::tl2::details::TasksTaskStatusWaitingWriteJSON(s, *this)) { return false; }
	return true;
}

bool tl2::tasks::TaskStatusWaiting::read(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksTaskStatusWaitingRead(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::TaskStatusWaiting::write(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksTaskStatusWaitingWrite(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::TaskStatusWaiting::read_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read(s2);
	s2.pass_data(s);
}

void tl2::tasks::TaskStatusWaiting::write_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write(s2);
	s2.pass_data(s);
}

bool tl2::tasks::TaskStatusWaiting::read_boxed(::basictl::tl_istream & s) noexcept {
	if (!::tl2::details::TasksTaskStatusWaitingReadBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

bool tl2::tasks::TaskStatusWaiting::write_boxed(::basictl::tl_ostream & s)const noexcept {
	if (!::tl2::details::TasksTaskStatusWaitingWriteBoxed(s, *this)) { return false; }
	s.last_release();
	return true;
}

void tl2::tasks::TaskStatusWaiting::read_boxed_or_throw(::basictl::tl_throwable_istream & s) {
	::basictl::tl_istream s2(s);
	this->read_boxed(s2);
	s2.pass_data(s);
}

void tl2::tasks::TaskStatusWaiting::write_boxed_or_throw(::basictl::tl_throwable_ostream & s)const {
	::basictl::tl_ostream s2(s);
	this->write_boxed(s2);
	s2.pass_data(s);
}

void tl2::details::TasksTaskStatusWaitingReset(::tl2::tasks::TaskStatusWaiting& item) {
}

bool tl2::details::TasksTaskStatusWaitingWriteJSON(std::ostream& s, const ::tl2::tasks::TaskStatusWaiting& item) {
	s << "true";
	return true;
}

bool tl2::details::TasksTaskStatusWaitingRead(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusWaiting& item) {
	return true;
}

bool tl2::details::TasksTaskStatusWaitingWrite(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusWaiting& item) {
	return true;
}

bool tl2::details::TasksTaskStatusWaitingReadBoxed(::basictl::tl_istream & s, ::tl2::tasks::TaskStatusWaiting& item) {
	if (!s.nat_read_exact_tag(0x16739c2c)) { return false; }
	return tl2::details::TasksTaskStatusWaitingRead(s, item);
}

bool tl2::details::TasksTaskStatusWaitingWriteBoxed(::basictl::tl_ostream & s, const ::tl2::tasks::TaskStatusWaiting& item) {
	if (!s.nat_write(0x16739c2c)) { return false; }
	return tl2::details::TasksTaskStatusWaitingWrite(s, item);
}
