# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-08-19  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	120519493	         9.957 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2256994	       504.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11194527	       107.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   57205	     22196 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	547406389	         2.191 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	275166184	         4.369 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	274782466	         4.367 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	349923906	         3.429 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21701155	        55.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1084 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1311884	       914.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   69510	     17177 ns/op	   15847 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	296530171	         4.046 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	427484780	         2.803 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11266093	       106.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13723641	        87.80 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	385268398	         3.115 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11355837	       105.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160140888	         7.482 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9961230	       120.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4694109	       247.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   73000	     16285 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	384067227	         3.116 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11394109	       105.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	479389819	         2.491 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	10850708	       110.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	201611233	         5.949 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8326194	       145.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	37589929	        31.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8544754	       140.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4292182	       280.1 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   74527	     16288 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	192682791	         6.224 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8348097	       144.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	51993190	        23.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	 9927046	       121.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4250890	       284.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	124279102	         9.654 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	13664283	        88.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9432690	       128.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   72015	     16637 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	296373217	         4.048 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 7981424	       149.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	551214344	         2.179 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   77073	     15617 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20842056	        57.84 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2514458	       476.8 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102918	     11710 ns/op	     146 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   71031	     16792 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24505596	        48.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2686041	       447.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3582979	       348.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70938	     16791 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.491 / 0 allocs | 110.6 / 0 allocs | **44.4x** | N/A | N/A | N/A | N/A |
| Enum | 4.046 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.28 / 0 allocs | 1084 / 89 B / 5 allocs | **19.6x** | 914.6 / 0 allocs | **16.5x** | 17177 / 15847 B / 76 allocs | **310.7x** |
| GTE | 3.115 / 0 allocs | 105.6 / 0 allocs | **33.9x** | N/A | N/A | N/A | N/A |
| MinLength | 23.05 / 0 allocs | 121.2 / 0 allocs | **5.3x** | 284.0 / 32 B / 2 allocs | **12.3x** | N/A | N/A |
| UUID | 48.62 / 0 allocs | 447.9 / 0 allocs | **9.2x** | 348.9 / 0 allocs | **7.2x** | 16791 / 15542 B / 76 allocs | **345.4x** |
| MaxItems | 5.949 / 0 allocs | 145.4 / 0 allocs | **24.4x** | N/A | N/A | N/A | N/A |
| MaxLength | 31.72 / 0 allocs | 140.7 / 0 allocs | **4.4x** | 280.1 / 32 B / 2 allocs | **8.8x** | 16288 / 15648 B / 81 allocs | **513.5x** |
| LT | 3.116 / 0 allocs | 105.5 / 0 allocs | **33.9x** | N/A | N/A | N/A | N/A |
| MinItems | 6.224 / 0 allocs | 144.1 / 0 allocs | **23.2x** | N/A | N/A | N/A | N/A |
| Alpha | 9.957 / 0 allocs | 504.4 / 0 allocs | **50.7x** | 107.7 / 0 allocs | **10.8x** | 22196 / 16937 B / 101 allocs | **2229.2x** |
| Required | 4.048 / 0 allocs | 149.7 / 0 allocs | **37.0x** | 2.179 / 0 allocs | **0.5x** | 15617 / 15488 B / 73 allocs | **3858.0x** |
| Length | 7.482 / 0 allocs | 120.3 / 0 allocs | **16.1x** | 247.0 / 32 B / 2 allocs | **33.0x** | 16285 / 15616 B / 79 allocs | **2176.6x** |
| URL | 57.84 / 0 allocs | 476.8 / 144 B / 1 allocs | **8.2x** | 11710 / 146 B / 1 allocs | **202.5x** | 16792 / 15681 B / 77 allocs | **290.3x** |
| Numeric | 9.654 / 0 allocs | 88.50 / 0 allocs | **9.2x** | 128.0 / 0 allocs | **13.3x** | 16637 / 15574 B / 78 allocs | **1723.3x** |
| GT | 2.803 / 0 allocs | 106.8 / 0 allocs | **38.1x** | 87.80 / 0 allocs | **31.3x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 2.191 | 0 allocs |
| CELMultipleExpressions | 4.369 | 0 allocs |
| CELBasic | 4.367 | 0 allocs |
| CELCrossField | 3.429 | 0 allocs |
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
