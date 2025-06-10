# compiler options
CC = g++
CFLAGS = -std=c++20 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror -Wunused-parameter

# compile all object files together
all: __build/main.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o __build/a.o __build/ab.o __build/au.o __build/cd.o __build/cyc1.o __build/cyc2.o __build/cyc3.o __build/service5.o __build/service5Long.o __build/usefulService.o 
	@mkdir -p __build
	$(CC) $(CFLAGS) -o all __build/main.o __build/io_streams.o __build/io_throwable_streams.o __build/__common_namespace.o __build/a.o __build/ab.o __build/au.o __build/cd.o __build/cyc1.o __build/cyc2.o __build/cyc3.o __build/service5.o __build/service5Long.o __build/usefulService.o 

__build/main.o: main.cpp
	@mkdir -p __build
	$(CC) $(CFLAGS) -c main.cpp -o __build/main.o

# compile meta data collection
__build/__meta.o: __meta/headers.h __meta/details.cpp __build
	$(CC) $(CFLAGS) -I. -o __build/__meta.o -c __meta/details.cpp

# compile objects factories
__build/__factory.o: __factory/headers.h __factory/details.cpp usefulService/functions/usefulService.getUserEntity.h __common_namespace/types/useTrue.h __common_namespace/types/useResponse.h __common_namespace/types/useDictUgly.h __common_namespace/types/unionArgsUse.h __common_namespace/types/typeD.h __common_namespace/types/typeC.h __common_namespace/types/typeB.h __common_namespace/types/true.h __common_namespace/types/testMaybe2.h __common_namespace/types/testMaybe.h service5/types/service5.stringOutput.h service5Long/types/service5Long.stringOutput.h service5Long/types/service5Long.emptyOutput.h service5/types/service5.emptyOutput.h __common_namespace/types/replace18.h __common_namespace/types/replace17.h __common_namespace/types/replace.h __common_namespace/types/replace9.h __common_namespace/types/replace8.h __common_namespace/types/replace7.h __common_namespace/types/replace6.h __common_namespace/types/replace5.h __common_namespace/types/replace3.h __common_namespace/types/replace2.h __common_namespace/types/replace15.h __common_namespace/types/replace12.h __common_namespace/types/replace10.h __common_namespace/types/nativeWrappers.h __common_namespace/types/myPlus3.h __common_namespace/types/myZero3.h __common_namespace/types/myPlus.h __common_namespace/types/myZero.h __common_namespace/types/myNat.h __common_namespace/types/multiPoint.h __common_namespace/types/maybeTest1.h __common_namespace/types/hren.h __common_namespace/types/cycleTuple.h cd/types/cd.useCycle.h cd/types/cd.typeD.h cd/types/cd.typeC.h cd/types/cd.typeB.h cd/types/cd.topLevel3.h __common_namespace/types/halfStr.h ab/types/ab.useDictString.h ab/types/ab.useCycle.h cyc1/types/cyc1.myCycle.h cyc2/types/cyc2.myCycle.h cyc3/types/cyc3.myCycle.h ab/types/ab.typeD.h ab/types/ab.typeC.h ab/types/ab.typeB.h ab/types/ab.topLevel2.h __common_namespace/types/useStr.h ab/types/ab.topLevel1.h __common_namespace/types/noStr.h ab/types/ab.testMaybe.h cd/types/cd.myType.h cd/types/cd.response.h ab/types/ab.myType.h ab/types/ab.empty.h ab/types/ab.counterChangeRequestPeriodsOne.h ab/types/ab.counterChangeRequestPeriodsMany.h ab/types/ab.code.h ab/functions/ab.call9.h ab/functions/ab.call8.h ab/functions/ab.call7.h __common_namespace/types/typeA.h ab/functions/ab.call6.h ab/functions/ab.call5.h ab/functions/ab.call4.h cd/types/cd.typeA.h ab/functions/ab.call3.h ab/functions/ab.call2.h ab/functions/ab.call11.h ab/functions/ab.call10.h ab/functions/ab.call1.h ab/types/ab.typeA.h a/types/a.top2.h au/types/au.nionA.h a/types/a.uNionA.h a/types/a.Color.h b/types/a.ColorItems.h a/types/a.ColorItems.h __build
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

# build object files for individual namespaces
__build/__common_namespace.o: __common_namespace/details.cpp  __common_namespace/headers/Bool.h __common_namespace/headers/MyNat2.h __common_namespace/headers/MyNat3.h __common_namespace/headers/UnionArgsXXX.h __common_namespace/headers/cycleTuple.h __common_namespace/headers/dictionaryElemStrange.h __common_namespace/headers/dictionaryField.h __common_namespace/headers/hren.h __common_namespace/headers/inner.h __common_namespace/headers/innerMaybe.h __common_namespace/headers/int.h __common_namespace/headers/int32.h __common_namespace/headers/int64.h __common_namespace/headers/long.h __common_namespace/headers/maybeTest1.h __common_namespace/headers/maybeWrapper.h __common_namespace/headers/myInt32.h __common_namespace/headers/myInt64.h __common_namespace/headers/myNat.h __common_namespace/headers/myPlus.h __common_namespace/headers/myPlus3.h __common_namespace/headers/myZero.h __common_namespace/headers/myZero3.h __common_namespace/headers/nativeWrappers.h __common_namespace/headers/noStr.h __common_namespace/headers/pair.h __common_namespace/headers/replace.h __common_namespace/headers/replace1.h __common_namespace/headers/replace10.h __common_namespace/headers/replace10Elem.h __common_namespace/headers/replace11.h __common_namespace/headers/replace11Elem.h __common_namespace/headers/replace12.h __common_namespace/headers/replace12Elem.h __common_namespace/headers/replace13.h __common_namespace/headers/replace13Elem.h __common_namespace/headers/replace14.h __common_namespace/headers/replace14Elem.h __common_namespace/headers/replace14ElemElem.h __common_namespace/headers/replace15.h __common_namespace/headers/replace15Elem.h __common_namespace/headers/replace15Elem1.h __common_namespace/headers/replace15Elem2.h __common_namespace/headers/replace17.h __common_namespace/headers/replace18.h __common_namespace/headers/replace2.h __common_namespace/headers/replace3.h __common_namespace/headers/replace4.h __common_namespace/headers/replace5.h __common_namespace/headers/replace6.h __common_namespace/headers/replace7.h __common_namespace/headers/replace8.h __common_namespace/headers/replace9.h __common_namespace/headers/testMaybe.h __common_namespace/headers/true.h __common_namespace/headers/tuple.h __common_namespace/headers/typeA.h __common_namespace/headers/unionArgsUse.h __common_namespace/headers/unionArgsXXX1.h __common_namespace/headers/unionArgsXXX2.h __common_namespace/headers/useStr.h __common_namespace/headers/useTrue.h __common_namespace/headers/vector.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/__common_namespace.o -c __common_namespace/details.cpp
__build/a.o: a/details.cpp  a/headers/a.Color.h a/headers/a.ColorItems.h a/headers/a.MyBool.h a/headers/a.MyUnion.h a/headers/a.inner.h a/headers/a.middle.h a/headers/a.top2.h a/headers/a.uNionA.h a/headers/a_a.ColorItems.h a/headers/a_dictionaryElem.h a/headers/a_dictionaryElemStrange.h a/headers/a_dictionaryElemUgly.h a/headers/a_int32.h a/headers/a_multiPoint.h a/headers/a_myInt32.h a/headers/a_pair.h a/headers/a_string.h a/headers/a_tuple.h a/headers/a_useDictUgly.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/a.o -c a/details.cpp
__build/ab.o: ab/details.cpp  ab/headers/ab.CounterChangeRequestPeriods.h ab/headers/ab.Response.h ab/headers/ab.alias.h ab/headers/ab.call1.h ab/headers/ab.call10.h ab/headers/ab.call11.h ab/headers/ab.call2.h ab/headers/ab.call3.h ab/headers/ab.call4.h ab/headers/ab.call5.h ab/headers/ab.call6.h ab/headers/ab.call7.h ab/headers/ab.call8.h ab/headers/ab.call9.h ab/headers/ab.code.h ab/headers/ab.counterChangeRequestPeriodsMany.h ab/headers/ab.counterChangeRequestPeriodsOne.h ab/headers/ab.empty.h ab/headers/ab.myType.h ab/headers/ab.testMaybe.h ab/headers/ab.topLevel1.h ab/headers/ab.topLevel2.h ab/headers/ab.typeA.h ab/headers/ab.typeB.h ab/headers/ab.typeC.h ab/headers/ab.typeD.h ab/headers/ab.useCycle.h ab/headers/ab.useDictString.h ab/headers/ab_a.Color.h ab/headers/ab_cd.myType.h ab/headers/ab_dictionary.h ab/headers/ab_dictionaryField.h ab/headers/ab_testMaybe2.h ab/headers/ab_typeB.h ab/headers/ab_typeC.h ab/headers/ab_useResponse.h ab/headers/ab_vector.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/ab.o -c ab/details.cpp
__build/au.o: au/details.cpp  au/headers/au.nionA.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/au.o -c au/details.cpp
__build/cd.o: cd/details.cpp  cd/headers/cd.myType.h cd/headers/cd.response.h cd/headers/cd.topLevel3.h cd/headers/cd.typeA.h cd/headers/cd.typeB.h cd/headers/cd.typeC.h cd/headers/cd.typeD.h cd/headers/cd.useCycle.h cd/headers/cd_a.Color.h cd/headers/cd_halfStr.h cd/headers/cd_typeD.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/cd.o -c cd/details.cpp
__build/cyc1.o: cyc1/details.cpp  cyc1/headers/cyc1.myCycle.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/cyc1.o -c cyc1/details.cpp
__build/cyc2.o: cyc2/details.cpp  cyc2/headers/cyc2.myCycle.h cyc2/headers/cyc2_cyc3.myCycle.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/cyc2.o -c cyc2/details.cpp
__build/cyc3.o: cyc3/details.cpp  cyc3/headers/cyc3.myCycle.h cyc3/headers/cyc3_cyc1.myCycle.h cyc3/headers/cyc3_vector.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/cyc3.o -c cyc3/details.cpp
__build/service5.o: service5/details.cpp  service5/headers/service5.Output.h service5/headers/service5.emptyOutput.h service5/headers/service5.stringOutput.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/service5.o -c service5/details.cpp
__build/service5Long.o: service5Long/details.cpp  service5Long/headers/service5Long.Output.h service5Long/headers/service5Long.emptyOutput.h service5Long/headers/service5Long.stringOutput.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/service5Long.o -c service5Long/details.cpp
__build/usefulService.o: usefulService/details.cpp  usefulService/headers/usefulService.getUserEntity.h usefulService/headers/usefulService.getUserEntityResult.h usefulService/headers/usefulService.userEntityPaymentItem.h usefulService/headers/usefulService.userEntityPaymentItemPromo.h usefulService/headers/usefulService_dictionary.h usefulService/headers/usefulService_dictionaryField.h usefulService/headers/usefulService_vector.h
	@mkdir -p __build
	$(CC) $(CFLAGS) -I. -o __build/usefulService.o -c usefulService/details.cpp
