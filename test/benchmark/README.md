# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-09-08  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	120180116	         9.929 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2580452	       470.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11287618	       108.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55956	     21074 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	644297460	         1.862 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296106288	         4.052 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296034838	         4.053 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	350225348	         3.428 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21720427	        59.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1083 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1283360	       929.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   70029	     16957 ns/op	   15850 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275277420	         4.357 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481400116	         2.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11291918	       106.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13694738	        88.21 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	480994658	         2.502 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11336304	       105.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160487818	         7.473 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9648940	       124.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4850866	       246.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73317	     16173 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	482088342	         2.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11353918	       105.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	427497583	         2.804 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10936252	       109.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	176637453	         6.831 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8379212	       143.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35931000	        33.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8621817	       139.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4221574	       283.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73671	     16156 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202995938	         5.910 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8309212	       144.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	51655111	        23.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10140795	       118.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4248636	       283.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	158039511	         7.602 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12166989	        98.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9430893	       128.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   72685	     16592 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321278953	         3.736 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8049316	       148.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	642666561	         1.870 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76299	     15520 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20799498	        57.89 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2482684	       483.4 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	   95500	     12526 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71446	     16706 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24694974	        48.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2691108	       454.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3583906	       335.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71426	     16664 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.804 / 0 allocs | 109.9 / 0 allocs | **39.2x** | N/A | N/A | N/A | N/A |
| Enum | 4.357 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 59.44 / 0 allocs | 1083 / 89 B / 5 allocs | **18.2x** | 929.2 / 0 allocs | **15.6x** | 16957 / 15850 B / 76 allocs | **285.3x** |
| GTE | 2.502 / 0 allocs | 105.6 / 0 allocs | **42.2x** | N/A | N/A | N/A | N/A |
| MinLength | 23.02 / 0 allocs | 118.7 / 0 allocs | **5.2x** | 283.2 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.67 / 0 allocs | 454.6 / 0 allocs | **9.3x** | 335.8 / 0 allocs | **6.9x** | 16664 / 15542 B / 76 allocs | **342.4x** |
| MaxItems | 6.831 / 0 allocs | 143.3 / 0 allocs | **21.0x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.33 / 0 allocs | 139.3 / 0 allocs | **4.2x** | 283.4 / 32 B / 2 allocs | **8.5x** | 16156 / 15648 B / 81 allocs | **484.7x** |
| LT | 2.489 / 0 allocs | 105.6 / 0 allocs | **42.4x** | N/A | N/A | N/A | N/A |
| MinItems | 5.910 / 0 allocs | 144.3 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| Alpha | 9.929 / 0 allocs | 470.2 / 0 allocs | **47.4x** | 108.6 / 0 allocs | **10.9x** | 21074 / 16937 B / 101 allocs | **2122.5x** |
| Required | 3.736 / 0 allocs | 148.8 / 0 allocs | **39.8x** | 1.870 / 0 allocs | **0.5x** | 15520 / 15488 B / 73 allocs | **4154.2x** |
| Length | 7.473 / 0 allocs | 124.3 / 0 allocs | **16.6x** | 246.4 / 32 B / 2 allocs | **33.0x** | 16173 / 15616 B / 79 allocs | **2164.2x** |
| URL | 57.89 / 0 allocs | 483.4 / 144 B / 1 allocs | **8.4x** | 12526 / 146 B / 1 allocs | **216.4x** | 16706 / 15681 B / 77 allocs | **288.6x** |
| Numeric | 7.602 / 0 allocs | 98.54 / 0 allocs | **13.0x** | 128.2 / 0 allocs | **16.9x** | 16592 / 15574 B / 78 allocs | **2182.6x** |
| GT | 2.489 / 0 allocs | 106.0 / 0 allocs | **42.6x** | 88.21 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.862 | 0 allocs |
| CELMultipleExpressions | 4.052 | 0 allocs |
| CELBasic | 4.053 | 0 allocs |
| CELCrossField | 3.428 | 0 allocs |
| CELStringLength | 1.000 | 0 allocs |
| CELNumericComparison | 1.000 | 0 allocs |

CEL (Common Expression Language) support allows complex runtime expressions with near-zero overhead.

## Running Benchmarks

```bash
# Update all benchmark documentation
make sync-benchmarks

# Run benchmarks manually
cd test
go test ./benchmark/ -bench=. -benchmem -benchtime=10s

# Run specific validator benchmarks
go test ./benchmark/ -bench=BenchmarkGoValid{ValidatorName} -benchmem
```
