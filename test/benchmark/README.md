# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-09-03  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	123721154	         9.701 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2572880	       469.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11002708	       110.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   48118	     22172 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	588715506	         2.113 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	294723667	         4.056 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	294695918	         4.063 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	345763069	         3.444 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21490412	        57.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	  969456	      1123 ns/op	      88 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1359522	       883.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   68055	     17466 ns/op	   15856 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275468059	         4.367 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481775340	         2.493 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11174151	       107.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13610720	        88.68 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	481051732	         2.507 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11314927	       106.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160613922	         7.469 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9499260	       125.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4680392	       252.2 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   71640	     16559 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	482103265	         2.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11346373	       106.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428448943	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10903593	       110.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183603441	         6.540 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8153713	       144.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	34508887	        33.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8597667	       139.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4140181	       287.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   72699	     16589 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	201691770	         5.933 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8328439	       144.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52018488	        23.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9956606	       118.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4137242	       289.7 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	156056712	         7.709 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12107311	       100.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9393145	       129.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   65227	     18241 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321140385	         3.737 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8063205	       148.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	643154853	         1.865 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   73345	     16266 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20660281	        57.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2391883	       502.5 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  100922	     11892 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   68923	     17346 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24756632	        48.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2669275	       449.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3420913	       351.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   68983	     17242 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.801 / 0 allocs | 110.2 / 0 allocs | **39.3x** | N/A | N/A | N/A | N/A |
| Enum | 4.367 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.80 / 0 allocs | 1123 / 88 B / 5 allocs | **19.4x** | 883.0 / 0 allocs | **15.3x** | 17466 / 15856 B / 76 allocs | **302.2x** |
| GTE | 2.507 / 0 allocs | 106.4 / 0 allocs | **42.4x** | N/A | N/A | N/A | N/A |
| MinLength | 23.02 / 0 allocs | 118.8 / 0 allocs | **5.2x** | 289.7 / 32 B / 2 allocs | **12.6x** | N/A | N/A |
| UUID | 48.61 / 0 allocs | 449.9 / 0 allocs | **9.3x** | 351.9 / 0 allocs | **7.2x** | 17242 / 15542 B / 76 allocs | **354.7x** |
| MaxItems | 6.540 / 0 allocs | 144.0 / 0 allocs | **22.0x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.46 / 0 allocs | 139.8 / 0 allocs | **4.2x** | 287.6 / 32 B / 2 allocs | **8.6x** | 16589 / 15648 B / 81 allocs | **495.8x** |
| LT | 2.489 / 0 allocs | 106.0 / 0 allocs | **42.6x** | N/A | N/A | N/A | N/A |
| MinItems | 5.933 / 0 allocs | 144.5 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| Alpha | 9.701 / 0 allocs | 469.3 / 0 allocs | **48.4x** | 110.3 / 0 allocs | **11.4x** | 22172 / 16937 B / 101 allocs | **2285.5x** |
| Required | 3.737 / 0 allocs | 148.9 / 0 allocs | **39.8x** | 1.865 / 0 allocs | **0.5x** | 16266 / 15488 B / 73 allocs | **4352.7x** |
| Length | 7.469 / 0 allocs | 125.0 / 0 allocs | **16.7x** | 252.2 / 32 B / 2 allocs | **33.8x** | 16559 / 15616 B / 79 allocs | **2217.0x** |
| URL | 57.60 / 0 allocs | 502.5 / 144 B / 1 allocs | **8.7x** | 11892 / 146 B / 1 allocs | **206.5x** | 17346 / 15681 B / 77 allocs | **301.1x** |
| Numeric | 7.709 / 0 allocs | 100.4 / 0 allocs | **13.0x** | 129.3 / 0 allocs | **16.8x** | 18241 / 15574 B / 78 allocs | **2366.2x** |
| GT | 2.493 / 0 allocs | 107.1 / 0 allocs | **43.0x** | 88.68 / 0 allocs | **35.6x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 2.113 | 0 allocs |
| CELMultipleExpressions | 4.056 | 0 allocs |
| CELBasic | 4.063 | 0 allocs |
| CELCrossField | 3.444 | 0 allocs |
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
