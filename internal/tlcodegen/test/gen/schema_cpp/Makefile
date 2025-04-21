
CC = g++
CFLAGS = -std=c++20 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror=return-type -Wno-unused-parameter
# compile all object files together
all: main.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o __build/antispam.o __build/pkg2.o __build/service1.o __build/service2.o __build/service3.o __build/service4.o __build/service5.o __build/service6.o __build/tasks.o __build/tree_stats.o __build/unique.o 
	@mkdir -p __build
	$(CC) $(CFLAGS) -o all main.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o __build/antispam.o __build/pkg2.o __build/service1.o __build/service2.o __build/service3.o __build/service4.o __build/service5.o __build/service6.o __build/tasks.o __build/tree_stats.o __build/unique.o 

main.o: main.cpp
	@mkdir -p __build
	$(CC) $(CFLAGS) -c main.cpp

# compile meta data collection
__build/__meta.o: __meta/headers.h __meta/details.cpp __build
	$(CC) $(CFLAGS) -I. -o __build/__meta.o -c __meta/details.cpp

# compile objects factories
__build/__factory.o: __factory/headers.h __factory/details.cpp __common_namespace/types/withFloat.h __common_namespace/types/vector.h unique/functions/unique.stringToInt.h unique/functions/unique.get.h __common_namespace/types/tuple.h tree_stats/types/tree_stats.objectLimitValueLong.h tasks/types/tasks.TaskStatusItems.h tasks/types/tasks.taskInfo.h tasks/types/tasks.queueTypeInfo.h tasks/types/tasks.queueTypeSettings.h tasks/types/tasks.queueStats.h tasks/functions/tasks.getTaskFromQueue.h tasks/functions/tasks.getQueueTypes.h tasks/functions/tasks.getQueueSize.h tasks/functions/tasks.getAnyTask.h tasks/types/tasks.fullFilledCron.h tasks/types/tasks.cronTaskWithId.h tasks/types/tasks.cronTask.h tasks/types/tasks.cronTime.h tasks/functions/tasks.addTask.h tasks/types/tasks.task.h __common_namespace/types/string.h __common_namespace/types/statOne.h service6/functions/service6.multiFindWithBounds.h service6/functions/service6.multiFind.h service6/types/service6.findWithBoundsResult.h service5/functions/service5.query.h service5/functions/service5.performQuery.h service5/types/service5.params.h service5/types/service5.Output.h service5/types/service5.stringOutput.h service5/functions/service5.insert.h service5/types/service5.emptyOutput.h service4/types/service4.modifiedNewsEntry.h service4/types/service4.object.h service3/functions/service3.setLimits.h service3/functions/service3.setLastVisitTimestamp.h service3/functions/service3.restoreProduct.h service3/functions/service3.restoreGroupedProducts.h service3/functions/service3.restoreAllProducts.h service3/types/service3.productStatsOld.h service3/types/service3.limits.h service3/types/service3.groupSizeLimit.h service3/types/service3.groupCountLimit.h service3/functions/service3.getScheduledProducts.h service3/functions/service3.getProducts.h service3/functions/service3.getProductStats.h service3/functions/service3.getLimits.h service3/functions/service3.getLastVisitTimestamp.h service3/functions/service3.deleteProduct.h service3/functions/service3.deleteGroupedProducts.h service3/functions/service3.deleteAllProducts.h service3/functions/service3.createProduct.h service2/functions/service2.setObjectTtl.h service2/functions/service2.set.h service2/functions/service2.addOrIncrMany.h service1/functions/service1.touch.h service1/functions/service1.setOrIncr.h service1/functions/service1.set.h service1/functions/service1.replaceOrIncr.h service1/functions/service1.replace.h service1/types/service1.keysStat.h service1/functions/service1.incr.h service1/functions/service1.getWildcardWithFlags.h service1/functions/service1.getWildcardList.h service1/functions/service1.getWildcardDict.h service1/functions/service1.getWildcard.h service1/functions/service1.getKeysStatPeriods.h service1/functions/service1.getKeysStat.h service1/functions/service1.getExpireTime.h service1/functions/service1.get.h service1/functions/service1.exists.h service1/functions/service1.enableKeysStat.h service1/functions/service1.enableExpiration.h service1/functions/service1.disableKeysStat.h service1/functions/service1.disableExpiration.h service1/functions/service1.delete.h service1/functions/service1.decr.h service1/functions/service1.cas.h service1/functions/service1.append.h service1/functions/service1.addOrIncr.h service1/functions/service1.addOrGet.h service1/functions/service1.add.h __common_namespace/types/rpcInvokeReqExtra.h __common_namespace/types/true.h pkg2/types/pkg2.t1.h pkg2/types/pkg2.foo.h __common_namespace/types/nonOptNat.h __common_namespace/types/myTwoDicts.h __common_namespace/types/myMcValueVector.h __common_namespace/types/myMcValueTuple.h __common_namespace/types/myMcValue.h __common_namespace/types/myBoxedVectorSlice.h __common_namespace/types/long.h __common_namespace/types/issue3498.h service6/types/service6.findResultRow.h service6/types/service6.error.h __common_namespace/types/int.h __common_namespace/functions/getStats.h tasks/types/tasks.queueTypeStats.h __common_namespace/functions/getNonOptNat.h __common_namespace/functions/getMyValue.h __common_namespace/types/MyValue.h __common_namespace/types/myString.h __common_namespace/types/myInt.h __common_namespace/functions/getMyDouble.h __common_namespace/types/myDouble.h __common_namespace/functions/getMyDictOfInt.h __common_namespace/types/myDictOfInt.h __common_namespace/functions/getMaybeIface.h service1/types/service1.Value.h service1/types/service1.strvalueWithTime.h service1/types/service1.strvalue.h service1/types/service1.not_found.h service1/types/service1.longvalueWithTime.h service1/types/service1.longvalue.h __common_namespace/functions/getFloat.h __common_namespace/functions/getDouble.h __common_namespace/functions/get_arrays.h __common_namespace/types/float.h __common_namespace/types/fieldConflict4.h __common_namespace/types/fieldConflict3.h __common_namespace/types/fieldConflict2.h __common_namespace/types/fieldConflict1.h __common_namespace/types/double.h __common_namespace/types/dictionary.h __common_namespace/functions/boxedVector64.h __common_namespace/functions/boxedVector32BoxedElem.h __common_namespace/functions/boxedVector32.h __common_namespace/functions/boxedTupleSlice3.h __common_namespace/functions/boxedTupleSlice2.h __common_namespace/types/myBoxedTupleSlice.h __common_namespace/functions/boxedTupleSlice1.h __common_namespace/functions/boxedTuple.h __common_namespace/functions/boxedString.h __common_namespace/functions/boxedInt.h __common_namespace/functions/boxedArray.h __common_namespace/types/myBoxedArray.h __common_namespace/types/boolStat.h __common_namespace/types/Bool.h __common_namespace/types/benchObject.h __common_namespace/types/integer.h antispam/types/antispam.PatternFull.h antispam/types/antispam.patternNotFound.h antispam/types/antispam.patternFound.h antispam/functions/antispam.getPattern.h __build
	$(CC) $(CFLAGS) -I. -o __build/__factory.o -c __factory/details.cpp

# compile streams which are used to work with io
__build/io_streams.o: basictl/constants.h basictl/errors.h basictl/io_connectors.h basictl/io_streams.cpp basictl/io_streams.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/io_streams.o -c basictl/io_streams.cpp

__build/io_throwable_streams.o: basictl/constants.h basictl/errors.h basictl/io_connectors.h basictl/io_throwable_streams.cpp basictl/io_throwable_streams.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/io_throwable_streams.o -c basictl/io_throwable_streams.cpp

__build/string_io.o: basictl/io_connectors.h basictl/impl/string_io.cpp basictl/impl/string_io.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/string_io.o -c basictl/impl/string_io.cpp

# create object files for individual namespaces
__build/__common_namespace.o: __common_namespace/details.cpp __common_namespace/details.cpp __common_namespace/headers/Bool.h __common_namespace/headers/MyValue.h __common_namespace/headers/benchObject.h __common_namespace/headers/boxedArray.h __common_namespace/headers/boxedInt.h __common_namespace/headers/boxedString.h __common_namespace/headers/boxedTuple.h __common_namespace/headers/boxedTupleSlice1.h __common_namespace/headers/boxedTupleSlice2.h __common_namespace/headers/boxedTupleSlice3.h __common_namespace/headers/boxedVector32.h __common_namespace/headers/boxedVector32BoxedElem.h __common_namespace/headers/boxedVector64.h __common_namespace/headers/dictionary.h __common_namespace/headers/dictionaryField.h __common_namespace/headers/double.h __common_namespace/headers/fieldConflict1.h __common_namespace/headers/fieldConflict2.h __common_namespace/headers/fieldConflict3.h __common_namespace/headers/fieldConflict4.h __common_namespace/headers/float.h __common_namespace/headers/getDouble.h __common_namespace/headers/getFloat.h __common_namespace/headers/getMyDictOfInt.h __common_namespace/headers/getMyDouble.h __common_namespace/headers/getMyValue.h __common_namespace/headers/getNonOptNat.h __common_namespace/headers/get_arrays.h __common_namespace/headers/int.h __common_namespace/headers/integer.h __common_namespace/headers/long.h __common_namespace/headers/myBoxedArray.h __common_namespace/headers/myBoxedTupleSlice.h __common_namespace/headers/myBoxedVectorSlice.h __common_namespace/headers/myDictOfInt.h __common_namespace/headers/myDouble.h __common_namespace/headers/myInt.h __common_namespace/headers/myMaybe1.h __common_namespace/headers/myMaybe2.h __common_namespace/headers/myString.h __common_namespace/headers/myTuple.h __common_namespace/headers/myTwoDicts.h __common_namespace/headers/nonOptNat.h __common_namespace/headers/rpcInvokeReqExtra.h __common_namespace/headers/statOne.h __common_namespace/headers/string.h __common_namespace/headers/true.h __common_namespace/headers/tuple.h __common_namespace/headers/vector.h __common_namespace/headers/withFloat.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/__common_namespace.o -c __common_namespace/details.cpp
__build/antispam.o: antispam/details.cpp antispam/details.cpp antispam/headers/antispam.PatternFull.h antispam/headers/antispam.getPattern.h antispam/headers/antispam.patternFound.h antispam/headers/antispam.patternNotFound.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/antispam.o -c antispam/details.cpp
__build/pkg2.o: pkg2/details.cpp pkg2/details.cpp pkg2/headers/pkg2.foo.h pkg2/headers/pkg2.t1.h pkg2/headers/pkg2.t2.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/pkg2.o -c pkg2/details.cpp
__build/service1.o: service1/details.cpp service1/details.cpp service1/headers/service1.Value.h service1/headers/service1.add.h service1/headers/service1.addOrGet.h service1/headers/service1.addOrIncr.h service1/headers/service1.append.h service1/headers/service1.cas.h service1/headers/service1.decr.h service1/headers/service1.delete.h service1/headers/service1.disableExpiration.h service1/headers/service1.disableKeysStat.h service1/headers/service1.enableExpiration.h service1/headers/service1.enableKeysStat.h service1/headers/service1.exists.h service1/headers/service1.get.h service1/headers/service1.getExpireTime.h service1/headers/service1.getKeysStat.h service1/headers/service1.getKeysStatPeriods.h service1/headers/service1.getWildcard.h service1/headers/service1.getWildcardDict.h service1/headers/service1.getWildcardList.h service1/headers/service1.getWildcardWithFlags.h service1/headers/service1.incr.h service1/headers/service1.keysStat.h service1/headers/service1.longvalue.h service1/headers/service1.longvalueWithTime.h service1/headers/service1.not_found.h service1/headers/service1.replace.h service1/headers/service1.replaceOrIncr.h service1/headers/service1.set.h service1/headers/service1.setOrIncr.h service1/headers/service1.strvalue.h service1/headers/service1.strvalueWithTime.h service1/headers/service1.touch.h service1/headers/service1_dictionary.h service1/headers/service1_dictionaryField.h service1/headers/service1_getMaybeIface.h service1/headers/service1_map.h service1/headers/service1_myAnonMcValue.h service1/headers/service1_myMcValue.h service1/headers/service1_myMcValueTuple.h service1/headers/service1_myMcValueVector.h service1/headers/service1_tuple.h service1/headers/service1_vector.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/service1.o -c service1/details.cpp
__build/service2.o: service2/details.cpp service2/details.cpp service2/headers/service2.addOrIncrMany.h service2/headers/service2.counterSet.h service2/headers/service2.deltaSet.h service2/headers/service2.objectId.h service2/headers/service2.set.h service2/headers/service2.setObjectTtl.h service2/headers/service2_double.h service2/headers/service2_tuple.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/service2.o -c service2/details.cpp
__build/service3.o: service3/details.cpp service3/details.cpp service3/headers/service3.createProduct.h service3/headers/service3.deleteAllProducts.h service3/headers/service3.deleteGroupedProducts.h service3/headers/service3.deleteProduct.h service3/headers/service3.getLastVisitTimestamp.h service3/headers/service3.getLimits.h service3/headers/service3.getProductStats.h service3/headers/service3.getProducts.h service3/headers/service3.getScheduledProducts.h service3/headers/service3.groupCountLimit.h service3/headers/service3.groupSizeLimit.h service3/headers/service3.limits.h service3/headers/service3.product.h service3/headers/service3.productStatsOld.h service3/headers/service3.restoreAllProducts.h service3/headers/service3.restoreGroupedProducts.h service3/headers/service3.restoreProduct.h service3/headers/service3.setLastVisitTimestamp.h service3/headers/service3.setLimits.h service3/headers/service3_boolStat.h service3/headers/service3_vector.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/service3.o -c service3/details.cpp
__build/service4.o: service4/details.cpp service4/details.cpp service4/headers/service4.modifiedNewsEntry.h service4/headers/service4.object.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/service4.o -c service4/details.cpp
__build/service5.o: service5/details.cpp service5/details.cpp service5/headers/service5.Output.h service5/headers/service5.emptyOutput.h service5/headers/service5.insert.h service5/headers/service5.params.h service5/headers/service5.performQuery.h service5/headers/service5.query.h service5/headers/service5.stringOutput.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/service5.o -c service5/details.cpp
__build/service6.o: service6/details.cpp service6/details.cpp service6/headers/service6.error.h service6/headers/service6.findResultRow.h service6/headers/service6.findWithBoundsResult.h service6/headers/service6.multiFind.h service6/headers/service6.multiFindWithBounds.h service6/headers/service6_Either.h service6/headers/service6_issue3498.h service6/headers/service6_left.h service6/headers/service6_right.h service6/headers/service6_vector.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/service6.o -c service6/details.cpp
__build/tasks.o: tasks/details.cpp tasks/details.cpp tasks/headers/tasks.TaskStatus.h tasks/headers/tasks.TaskStatusItems.h tasks/headers/tasks.addTask.h tasks/headers/tasks.cronTask.h tasks/headers/tasks.cronTaskWithId.h tasks/headers/tasks.cronTime.h tasks/headers/tasks.fullFilledCron.h tasks/headers/tasks.getAnyTask.h tasks/headers/tasks.getQueueSize.h tasks/headers/tasks.getQueueTypes.h tasks/headers/tasks.getTaskFromQueue.h tasks/headers/tasks.queueStats.h tasks/headers/tasks.queueTypeInfo.h tasks/headers/tasks.queueTypeSettings.h tasks/headers/tasks.queueTypeStats.h tasks/headers/tasks.task.h tasks/headers/tasks.taskInfo.h tasks/headers/tasks_getStats.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/tasks.o -c tasks/details.cpp
__build/tree_stats.o: tree_stats/details.cpp tree_stats/details.cpp tree_stats/headers/tree_stats.ObjectLimitValue.h tree_stats/headers/tree_stats.objectLimitValueDouble.h tree_stats/headers/tree_stats.objectLimitValueLong.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/tree_stats.o -c tree_stats/details.cpp
__build/unique.o: unique/details.cpp unique/details.cpp unique/headers/unique.get.h unique/headers/unique.stringToInt.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/unique.o -c unique/details.cpp
# compile individual namespaces
__compile/__common_namespace: main.o __build/__common_namespace.o __build/io_streams.o __build/io_throwable_streams.o
	@mkdir -p __compile
	$(CC) $(CFLAGS) -o __compile/__common_namespace main.o __build/__common_namespace.o __build/io_streams.o __build/io_throwable_streams.o

__compile/antispam: main.o __build/antispam.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o
	@mkdir -p __compile
	$(CC) $(CFLAGS) -o __compile/antispam main.o __build/antispam.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o

__compile/pkg2: main.o __build/pkg2.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o
	@mkdir -p __compile
	$(CC) $(CFLAGS) -o __compile/pkg2 main.o __build/pkg2.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o

__compile/service1: main.o __build/service1.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o
	@mkdir -p __compile
	$(CC) $(CFLAGS) -o __compile/service1 main.o __build/service1.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o

__compile/service2: main.o __build/service2.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o
	@mkdir -p __compile
	$(CC) $(CFLAGS) -o __compile/service2 main.o __build/service2.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o

__compile/service3: main.o __build/service3.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o
	@mkdir -p __compile
	$(CC) $(CFLAGS) -o __compile/service3 main.o __build/service3.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o

__compile/service4: main.o __build/service4.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o
	@mkdir -p __compile
	$(CC) $(CFLAGS) -o __compile/service4 main.o __build/service4.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o

__compile/service5: main.o __build/service5.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o
	@mkdir -p __compile
	$(CC) $(CFLAGS) -o __compile/service5 main.o __build/service5.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o

__compile/service6: main.o __build/service6.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o
	@mkdir -p __compile
	$(CC) $(CFLAGS) -o __compile/service6 main.o __build/service6.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o

__compile/tasks: main.o __build/tasks.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o
	@mkdir -p __compile
	$(CC) $(CFLAGS) -o __compile/tasks main.o __build/tasks.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o

__compile/tree_stats: main.o __build/tree_stats.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o
	@mkdir -p __compile
	$(CC) $(CFLAGS) -o __compile/tree_stats main.o __build/tree_stats.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o

__compile/unique: main.o __build/unique.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o
	@mkdir -p __compile
	$(CC) $(CFLAGS) -o __compile/unique main.o __build/unique.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o

# compile all namespaces (for test purposes)
compile-all: __compile/__common_namespace __compile/antispam __compile/pkg2 __compile/service1 __compile/service2 __compile/service3 __compile/service4 __compile/service5 __compile/service6 __compile/tasks __compile/tree_stats __compile/unique
