# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-09-30  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124182381	         9.662 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2541961	       476.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	10269609	       111.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   52578	     24149 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	633317733	         1.930 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	285206377	         4.147 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295749367	         4.053 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	348924680	         3.436 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21147457	        57.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1133 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1288268	       929.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   65457	     19134 ns/op	   15865 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	274361191	         4.364 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	471043140	         2.504 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11319506	       106.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	12353662	        89.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	480862692	         2.493 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11119189	       106.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	159560292	         7.507 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9638751	       125.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4493911	       262.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   66746	     18833 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	480452194	         2.492 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11310369	       105.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	418908711	         2.817 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10940053	       110.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	182282949	         6.561 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8110002	       144.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36071828	        33.30 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8624440	       139.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4108018	       295.3 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   59984	     18730 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202533580	         5.924 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8277070	       144.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	48374509	        23.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10081760	       119.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4190654	       286.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157010586	         7.643 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	11959680	        99.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9346436	       129.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   68929	     19273 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	308326153	         3.786 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8020429	       149.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	641916962	         1.867 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   72993	     18010 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	19427481	        57.76 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2340690	       518.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102561	     11882 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   63831	     18557 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	22464633	        48.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2686218	       449.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3357324	       355.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   67230	     17261 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.817 / 0 allocs | 110.3 / 0 allocs | **39.2x** | N/A | N/A | N/A | N/A |
| Enum | 4.364 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.41 / 0 allocs | 1133 / 88 B / 5 allocs | **19.7x** | 929.5 / 0 allocs | **16.2x** | 19134 / 15865 B / 76 allocs | **333.3x** |
| GTE | 2.493 / 0 allocs | 106.3 / 0 allocs | **42.6x** | N/A | N/A | N/A | N/A |
| MinLength | 23.08 / 0 allocs | 119.1 / 0 allocs | **5.2x** | 286.9 / 32 B / 2 allocs | **12.4x** | N/A | N/A |
| UUID | 48.86 / 0 allocs | 449.4 / 0 allocs | **9.2x** | 355.4 / 0 allocs | **7.3x** | 17261 / 15542 B / 76 allocs | **353.3x** |
| MaxItems | 6.561 / 0 allocs | 144.1 / 0 allocs | **22.0x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.30 / 0 allocs | 139.1 / 0 allocs | **4.2x** | 295.3 / 32 B / 2 allocs | **8.9x** | 18730 / 15648 B / 81 allocs | **562.5x** |
| LT | 2.492 / 0 allocs | 105.8 / 0 allocs | **42.5x** | N/A | N/A | N/A | N/A |
| MinItems | 5.924 / 0 allocs | 144.8 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| Alpha | 9.662 / 0 allocs | 476.5 / 0 allocs | **49.3x** | 111.8 / 0 allocs | **11.6x** | 24149 / 16937 B / 101 allocs | **2499.4x** |
| Required | 3.786 / 0 allocs | 149.7 / 0 allocs | **39.5x** | 1.867 / 0 allocs | **0.5x** | 18010 / 15488 B / 73 allocs | **4757.0x** |
| Length | 7.507 / 0 allocs | 125.6 / 0 allocs | **16.7x** | 262.2 / 32 B / 2 allocs | **34.9x** | 18833 / 15616 B / 79 allocs | **2508.7x** |
| URL | 57.76 / 0 allocs | 518.2 / 144 B / 1 allocs | **9.0x** | 11882 / 146 B / 1 allocs | **205.7x** | 18557 / 15681 B / 77 allocs | **321.3x** |
| Numeric | 7.643 / 0 allocs | 99.48 / 0 allocs | **13.0x** | 129.6 / 0 allocs | **17.0x** | 19273 / 15574 B / 78 allocs | **2521.7x** |
| GT | 2.504 / 0 allocs | 106.6 / 0 allocs | **42.6x** | 89.80 / 0 allocs | **35.9x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.930 | 0 allocs |
| CELMultipleExpressions | 4.147 | 0 allocs |
| CELBasic | 4.053 | 0 allocs |
| CELCrossField | 3.436 | 0 allocs |
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
