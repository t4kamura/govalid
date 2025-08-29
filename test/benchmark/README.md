# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-08-29  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124493932	         9.638 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2604486	       467.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11243094	       108.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56419	     21129 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	644449602	         1.863 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	295622466	         4.064 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295765201	         4.052 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	349386493	         3.430 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21681102	        57.29 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1090 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1357944	       883.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69332	     17176 ns/op	   15846 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	274995867	         4.365 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481775254	         2.494 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11321208	       106.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13578297	        88.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	482516252	         2.487 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11379686	       105.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160695715	         7.468 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9624538	       124.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4848124	       247.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73526	     16222 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	480476317	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11351236	       105.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	427403314	         2.803 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11009536	       109.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183726565	         6.530 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8281614	       143.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35674026	        33.40 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8632646	       139.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4213498	       284.9 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   74631	     16160 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202898401	         5.912 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8371674	       144.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	52081485	        23.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10122384	       118.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4238194	       283.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	155921322	         7.674 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12156308	        98.73 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9388420	       128.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   72325	     16624 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	320884311	         3.738 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8026161	       149.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	643046090	         1.866 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75903	     15680 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	19548649	        61.86 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2478184	       482.9 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103156	     11685 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70719	     16813 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24753098	        48.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2679643	       455.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3583272	       335.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70315	     16814 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.803 / 0 allocs | 109.5 / 0 allocs | **39.1x** | N/A | N/A | N/A | N/A |
| Enum | 4.365 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 57.29 / 0 allocs | 1090 / 89 B / 5 allocs | **19.0x** | 883.8 / 0 allocs | **15.4x** | 17176 / 15846 B / 76 allocs | **299.8x** |
| GTE | 2.487 / 0 allocs | 105.7 / 0 allocs | **42.5x** | N/A | N/A | N/A | N/A |
| MinLength | 23.07 / 0 allocs | 118.8 / 0 allocs | **5.1x** | 283.4 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.53 / 0 allocs | 455.4 / 0 allocs | **9.4x** | 335.7 / 0 allocs | **6.9x** | 16814 / 15542 B / 76 allocs | **346.5x** |
| MaxItems | 6.530 / 0 allocs | 143.9 / 0 allocs | **22.0x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.40 / 0 allocs | 139.6 / 0 allocs | **4.2x** | 284.9 / 32 B / 2 allocs | **8.5x** | 16160 / 15648 B / 81 allocs | **483.8x** |
| LT | 2.491 / 0 allocs | 105.7 / 0 allocs | **42.4x** | N/A | N/A | N/A | N/A |
| MinItems | 5.912 / 0 allocs | 144.1 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| Alpha | 9.638 / 0 allocs | 467.5 / 0 allocs | **48.5x** | 108.0 / 0 allocs | **11.2x** | 21129 / 16937 B / 101 allocs | **2192.3x** |
| Required | 3.738 / 0 allocs | 149.3 / 0 allocs | **39.9x** | 1.866 / 0 allocs | **0.5x** | 15680 / 15488 B / 73 allocs | **4194.8x** |
| Length | 7.468 / 0 allocs | 124.8 / 0 allocs | **16.7x** | 247.9 / 32 B / 2 allocs | **33.2x** | 16222 / 15616 B / 79 allocs | **2172.2x** |
| URL | 61.86 / 0 allocs | 482.9 / 144 B / 1 allocs | **7.8x** | 11685 / 145 B / 1 allocs | **188.9x** | 16813 / 15681 B / 77 allocs | **271.8x** |
| Numeric | 7.674 / 0 allocs | 98.73 / 0 allocs | **12.9x** | 128.5 / 0 allocs | **16.7x** | 16624 / 15574 B / 78 allocs | **2166.3x** |
| GT | 2.494 / 0 allocs | 106.1 / 0 allocs | **42.5x** | 88.22 / 0 allocs | **35.4x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.863 | 0 allocs |
| CELMultipleExpressions | 4.064 | 0 allocs |
| CELBasic | 4.052 | 0 allocs |
| CELCrossField | 3.430 | 0 allocs |
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
