# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-09-20  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124449044	         9.642 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2603110	       471.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11159906	       108.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   56764	     21034 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	644513460	         1.864 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	295833932	         4.066 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	295279279	         4.052 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	349866393	         3.428 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21750018	        55.22 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1093 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1335039	       910.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   70284	     17149 ns/op	   15864 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	274897767	         4.409 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481254470	         2.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11326743	       106.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13375290	        89.90 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	482567649	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11380311	       105.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160500085	         7.478 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9659570	       124.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4877341	       246.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73953	     16150 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481942939	         2.496 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11197694	       107.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	428176833	         2.801 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10899594	       110.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	183710517	         6.536 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8248106	       145.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	35898661	        33.33 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8639619	       138.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4226721	       283.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   73005	     16290 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202770000	         5.915 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8233383	       145.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	51761948	        23.02 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10116156	       118.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4260620	       282.4 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157676466	         7.601 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12166359	        98.95 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9417615	       128.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71077	     16611 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	321045036	         3.743 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8032476	       150.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	643370888	         1.866 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   76177	     15606 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20829110	        57.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2467112	       486.6 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  103206	     11692 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71444	     16827 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24684006	        48.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2693738	       447.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3573091	       343.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   71116	     16858 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.801 / 0 allocs | 110.5 / 0 allocs | **39.5x** | N/A | N/A | N/A | N/A |
| Enum | 4.409 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.22 / 0 allocs | 1093 / 89 B / 5 allocs | **19.8x** | 910.6 / 0 allocs | **16.5x** | 17149 / 15864 B / 76 allocs | **310.6x** |
| GTE | 2.491 / 0 allocs | 105.7 / 0 allocs | **42.4x** | N/A | N/A | N/A | N/A |
| MinLength | 23.02 / 0 allocs | 118.5 / 0 allocs | **5.1x** | 282.4 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.62 / 0 allocs | 447.1 / 0 allocs | **9.2x** | 343.4 / 0 allocs | **7.1x** | 16858 / 15542 B / 76 allocs | **346.7x** |
| MaxItems | 6.536 / 0 allocs | 145.8 / 0 allocs | **22.3x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.33 / 0 allocs | 138.9 / 0 allocs | **4.2x** | 283.1 / 32 B / 2 allocs | **8.5x** | 16290 / 15648 B / 81 allocs | **488.7x** |
| LT | 2.496 / 0 allocs | 107.3 / 0 allocs | **43.0x** | N/A | N/A | N/A | N/A |
| MinItems | 5.915 / 0 allocs | 145.6 / 0 allocs | **24.6x** | N/A | N/A | N/A | N/A |
| Alpha | 9.642 / 0 allocs | 471.0 / 0 allocs | **48.8x** | 108.1 / 0 allocs | **11.2x** | 21034 / 16937 B / 101 allocs | **2181.5x** |
| Required | 3.743 / 0 allocs | 150.2 / 0 allocs | **40.1x** | 1.866 / 0 allocs | **0.5x** | 15606 / 15488 B / 73 allocs | **4169.4x** |
| Length | 7.478 / 0 allocs | 124.6 / 0 allocs | **16.7x** | 246.0 / 32 B / 2 allocs | **32.9x** | 16150 / 15616 B / 79 allocs | **2159.7x** |
| URL | 57.65 / 0 allocs | 486.6 / 144 B / 1 allocs | **8.4x** | 11692 / 146 B / 1 allocs | **202.8x** | 16827 / 15681 B / 77 allocs | **291.9x** |
| Numeric | 7.601 / 0 allocs | 98.95 / 0 allocs | **13.0x** | 128.2 / 0 allocs | **16.9x** | 16611 / 15574 B / 78 allocs | **2185.4x** |
| GT | 2.489 / 0 allocs | 106.0 / 0 allocs | **42.6x** | 89.90 / 0 allocs | **36.1x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.864 | 0 allocs |
| CELMultipleExpressions | 4.066 | 0 allocs |
| CELBasic | 4.052 | 0 allocs |
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
