
CC = g++
CFLAGS = -std=c++20 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror=return-type -Wno-unused-parameter
all: build/tl main.o build/basictl.obuild/tl/__common_namespace.o build/tl/benchmarks.o build/tl/cases.o build/tl/cases_bytes.o 
	$(CC) $(CFLAGS) -o all main.o build/basictl.o build/tl/__common_namespace.o build/tl/benchmarks.o build/tl/cases.o build/tl/cases_bytes.o 

build/tl:
	mkdir -p build/tl

main.o: main.cpp
	$(CC) $(CFLAGS) -c main.cpp

build/meta.o: meta/headers.h meta/details.cpp
	$(CC) $(CFLAGS) -o build/meta.o -c meta/details.cpp

build/factory.o: factory/headers.h factory/details.cpp tl/cases/types/cases.testVector.h tl/cases/types/cases.testUnionContainer.h tl/cases/types/cases.testUnion2.h tl/cases/types/cases.testUnion1.h tl/cases/types/cases.testTuple.h tl/cases/types/cases.testRecursiveFieldMask.h tl/cases/types/cases.testOutFieldMaskContainer.h tl/cases/types/cases.testMaybe.h tl/cases/types/cases.testLocalFieldmask.h tl/cases/types/cases.testEnumContainer.h tl/cases/types/cases.testDictString.h tl/cases/types/cases.testDictInt.h tl/cases/types/cases.testDictAny.h tl/cases/types/cases.testBeforeReadBitValidation.h tl/cases/types/cases.testArray.h tl/cases/types/cases.testAllPossibleFieldConfigsContainer.h tl/cases/types/cases.replace7plusplus.h tl/cases/types/cases.replace7plus.h tl/cases/types/cases.replace7.h tl/cases/types/cases.myCycle1.h tl/cases/types/cases.myCycle2.h tl/cases/types/cases.myCycle3.h tl/cases_bytes/types/cases_bytes.testVector.h tl/cases_bytes/types/cases_bytes.testTuple.h tl/cases_bytes/types/cases_bytes.testEnumContainer.h tl/cases/types/cases.TestEnumItems.h tl/cases_bytes/types/cases_bytes.TestEnumItems.h tl/cases_bytes/types/cases_bytes.testDictStringString.h tl/cases_bytes/types/cases_bytes.testDictString.h tl/cases_bytes/types/cases_bytes.testDictInt.h tl/cases_bytes/types/cases_bytes.testDictAny.h tl/cases_bytes/types/cases_bytes.testArray.h tl/benchmarks/types/benchmarks.vrutoyTopLevelContainerWithDependency.h tl/benchmarks/types/benchmarks.vrutoyTopLevelContainer.h tl/benchmarks/types/benchmarks.vrutoytopLevelUnionEmpty.h tl/benchmarks/types/benchmarks.vrutoytopLevelUnionBig.h tl/benchmarks/types/benchmarks.vruposition.h tl/__common_namespace/types/true.h tl/benchmarks/types/benchmarks.vruhash.h
	$(CC) $(CFLAGS) -o build/factory.o -c factory/details.cpp

build/basictl.o: basics/basictl.cpp basics/basictl.h
	$(CC) $(CFLAGS) -o build/basictl.o -c basics/basictl.cpp

build/string_io.o: basics/string_io.cpp basics/string_io.h
	$(CC) $(CFLAGS) -o build/string_io.o -c basics/string_io.cpp

build/tl/__common_namespace.o: tl/__common_namespace/details.cpp tl/__common_namespace/details.cpp tl/__common_namespace/headers/dictionary.h tl/__common_namespace/headers/dictionaryAny.h tl/__common_namespace/headers/dictionaryField.h tl/__common_namespace/headers/dictionaryFieldAny.h tl/__common_namespace/headers/int.h tl/__common_namespace/headers/int32.h tl/__common_namespace/headers/int64.h tl/__common_namespace/headers/long.h tl/__common_namespace/headers/string.h tl/__common_namespace/headers/true.h tl/__common_namespace/headers/vector.h
	$(CC) $(CFLAGS) -o build/tl/__common_namespace.o -c tl/__common_namespace/details.cpp
build/tl/benchmarks.o: tl/benchmarks/details.cpp tl/benchmarks/details.cpp tl/benchmarks/headers/benchmarks.VrutoyTopLevelUnion.h tl/benchmarks/headers/benchmarks.vruhash.h tl/benchmarks/headers/benchmarks.vruposition.h tl/benchmarks/headers/benchmarks.vrutoyPositions.h tl/benchmarks/headers/benchmarks.vrutoyTopLevelContainer.h tl/benchmarks/headers/benchmarks.vrutoyTopLevelContainerWithDependency.h tl/benchmarks/headers/benchmarks.vrutoytopLevelUnionBig.h tl/benchmarks/headers/benchmarks.vrutoytopLevelUnionEmpty.h tl/benchmarks/headers/benchmarks_vector.h
	$(CC) $(CFLAGS) -o build/tl/benchmarks.o -c tl/benchmarks/details.cpp
build/tl/cases.o: tl/cases/details.cpp tl/cases/details.cpp tl/cases/headers/cases.TestEnum.h tl/cases/headers/cases.TestEnumItems.h tl/cases/headers/cases.TestUnion.h tl/cases/headers/cases.myCycle1.h tl/cases/headers/cases.myCycle2.h tl/cases/headers/cases.myCycle3.h tl/cases/headers/cases.replace7.h tl/cases/headers/cases.replace7plus.h tl/cases/headers/cases.replace7plusplus.h tl/cases/headers/cases.testAllPossibleFieldConfigs.h tl/cases/headers/cases.testAllPossibleFieldConfigsContainer.h tl/cases/headers/cases.testArray.h tl/cases/headers/cases.testBeforeReadBitValidation.h tl/cases/headers/cases.testDictAny.h tl/cases/headers/cases.testDictInt.h tl/cases/headers/cases.testDictString.h tl/cases/headers/cases.testEnumContainer.h tl/cases/headers/cases.testLocalFieldmask.h tl/cases/headers/cases.testMaybe.h tl/cases/headers/cases.testOutFieldMask.h tl/cases/headers/cases.testOutFieldMaskContainer.h tl/cases/headers/cases.testRecursiveFieldMask.h tl/cases/headers/cases.testTuple.h tl/cases/headers/cases.testUnion1.h tl/cases/headers/cases.testUnion2.h tl/cases/headers/cases.testUnionContainer.h tl/cases/headers/cases.testVector.h tl/cases/headers/cases_int.h tl/cases/headers/cases_tuple.h tl/cases/headers/cases_vector.h
	$(CC) $(CFLAGS) -o build/tl/cases.o -c tl/cases/details.cpp
build/tl/cases_bytes.o: tl/cases_bytes/details.cpp tl/cases_bytes/details.cpp tl/cases_bytes/headers/cases_bytes.TestEnum.h tl/cases_bytes/headers/cases_bytes.TestEnumItems.h tl/cases_bytes/headers/cases_bytes.testArray.h tl/cases_bytes/headers/cases_bytes.testDictAny.h tl/cases_bytes/headers/cases_bytes.testDictInt.h tl/cases_bytes/headers/cases_bytes.testDictString.h tl/cases_bytes/headers/cases_bytes.testDictStringString.h tl/cases_bytes/headers/cases_bytes.testEnumContainer.h tl/cases_bytes/headers/cases_bytes.testTuple.h tl/cases_bytes/headers/cases_bytes.testVector.h tl/cases_bytes/headers/cases_bytes_dictionary.h tl/cases_bytes/headers/cases_bytes_dictionaryField.h tl/cases_bytes/headers/cases_bytes_string.h tl/cases_bytes/headers/cases_bytes_tuple.h tl/cases_bytes/headers/cases_bytes_vector.h
	$(CC) $(CFLAGS) -o build/tl/cases_bytes.o -c tl/cases_bytes/details.cpp
