# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-08-15  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	100000000	        10.25 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2513702	       483.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10666924	       112.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   53282	     22411 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	519052612	         2.277 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	268293296	         4.566 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	262907355	         4.557 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	333640660	         3.569 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	19327504	        57.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	  862719	      1176 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1265624	       950.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   64688	     18485 ns/op	   15857 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	283458189	         4.242 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	409061706	         2.919 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	10872572	       111.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	12774891	        92.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	368692372	         3.250 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	10792124	       111.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	152152006	         7.855 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9721320	       123.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4480106	       266.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   69578	     17251 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	374337532	         3.179 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	10753572	       111.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	450776774	         2.611 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10510958	       114.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	194776780	         6.162 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8046544	       148.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36840680	        32.19 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8438376	       141.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 3983755	       304.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   69100	     17506 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	182254294	         6.607 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8201967	       150.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	50087284	        23.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9445496	       125.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 3855003	       308.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	100000000	        10.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12836562	        91.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9032996	       134.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   65862	     17987 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	280244545	         4.238 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7705500	       157.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	529326152	         2.241 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   70465	     16820 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	18756708	        62.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2349837	       509.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	   95130	     12666 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   66793	     18030 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	23199217	        50.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2528600	       473.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3331137	       359.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   66866	     18126 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.611 / 0 allocs | 114.5 / 0 allocs | **43.9x** | N/A | N/A | N/A | N/A |
| Enum | 4.242 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.45 / 0 allocs | 1176 / 89 B / 5 allocs | **20.5x** | 950.2 / 0 allocs | **16.5x** | 18485 / 15857 B / 76 allocs | **321.8x** |
| GTE | 3.250 / 0 allocs | 111.7 / 0 allocs | **34.4x** | N/A | N/A | N/A | N/A |
| MinLength | 23.64 / 0 allocs | 125.5 / 0 allocs | **5.3x** | 308.1 / 32 B / 2 allocs | **13.0x** | N/A | N/A |
| UUID | 50.50 / 0 allocs | 473.7 / 0 allocs | **9.4x** | 359.3 / 0 allocs | **7.1x** | 18126 / 15542 B / 76 allocs | **358.9x** |
| MaxItems | 6.162 / 0 allocs | 148.9 / 0 allocs | **24.2x** | N/A | N/A | N/A | N/A |
| MaxLength | 32.19 / 0 allocs | 141.8 / 0 allocs | **4.4x** | 304.1 / 32 B / 2 allocs | **9.4x** | 17506 / 15648 B / 81 allocs | **543.8x** |
| LT | 3.179 / 0 allocs | 111.2 / 0 allocs | **35.0x** | N/A | N/A | N/A | N/A |
| MinItems | 6.607 / 0 allocs | 150.1 / 0 allocs | **22.7x** | N/A | N/A | N/A | N/A |
| Alpha | 10.25 / 0 allocs | 483.7 / 0 allocs | **47.2x** | 112.6 / 0 allocs | **11.0x** | 22411 / 16937 B / 101 allocs | **2186.4x** |
| Required | 4.238 / 0 allocs | 157.8 / 0 allocs | **37.2x** | 2.241 / 0 allocs | **0.5x** | 16820 / 15488 B / 73 allocs | **3968.9x** |
| Length | 7.855 / 0 allocs | 123.1 / 0 allocs | **15.7x** | 266.4 / 32 B / 2 allocs | **33.9x** | 17251 / 15616 B / 79 allocs | **2196.2x** |
| URL | 62.20 / 0 allocs | 509.8 / 144 B / 1 allocs | **8.2x** | 12666 / 145 B / 1 allocs | **203.6x** | 18030 / 15681 B / 77 allocs | **289.9x** |
| Numeric | 10.10 / 0 allocs | 91.46 / 0 allocs | **9.1x** | 134.8 / 0 allocs | **13.3x** | 17987 / 15574 B / 78 allocs | **1780.9x** |
| GT | 2.919 / 0 allocs | 111.8 / 0 allocs | **38.3x** | 92.67 / 0 allocs | **31.7x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 2.277 | 0 allocs |
| CELMultipleExpressions | 4.566 | 0 allocs |
| CELBasic | 4.557 | 0 allocs |
| CELCrossField | 3.569 | 0 allocs |
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
