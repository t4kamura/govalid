# Benchmark Results

Performance comparison between govalid and popular Go validation libraries.

## Latest Results

**Benchmarked on:** 2025-09-02  
**Platform:** Linux 6.11.0-1018-azure x86_64  
**Go version:** go1.24.3

## Raw Benchmark Data

```
BenchmarkGoValidAlpha-4                    	124295905	         9.655 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundAlpha-4               	 2601813	       467.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkAsaskevichGovalidatorAlpha-4      	11161824	       108.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateAlpha-4             	   55626	     21548 ns/op	   16937 B/op	     101 allocs/op
BenchmarkGoValidCELConcurrent-4            	645205753	         1.860 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELMultipleExpressions-4   	296221032	         4.058 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELBasic-4                 	296236803	         4.053 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELCrossField-4            	349650669	         3.431 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELStringLength-4          	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidCELNumericComparison-4     	1000000000	         1.000 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidEmail-4                    	21586681	        55.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundEmail-4               	 1000000	      1115 ns/op	      89 B/op	       5 allocs/op
BenchmarkGoValidatorEmail-4                	 1271743	       931.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateEmail-4             	   68436	     17352 ns/op	   15869 B/op	      76 allocs/op
BenchmarkGoValidEnum-4                     	275078018	         4.361 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGT-4                       	481476141	         2.492 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGT-4                  	11263264	       106.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorGT-4                   	13606564	        89.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidGTE-4                      	479721094	         2.501 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundGTE-4                 	11335855	       106.0 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLength-4                   	160409286	         7.478 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLength-4              	 9640614	       124.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorLength-4               	 4803214	       249.6 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateLength-4            	   72363	     16363 ns/op	   15616 B/op	      79 allocs/op
BenchmarkGoValidLT-4                       	481229320	         2.494 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLT-4                  	11356444	       105.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidLTE-4                      	427975206	         2.805 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundLTE-4                 	11007324	       109.5 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxItems-4                 	182694579	         6.563 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxItems-4            	 8286483	       144.7 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMaxLength-4                	36006061	        33.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMaxLength-4           	 8598453	       139.3 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMaxLength-4            	 4175902	       286.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGookitValidateMaxLength-4         	   72754	     16496 ns/op	   15648 B/op	      81 allocs/op
BenchmarkGoValidMinItems-4                 	202793950	         5.920 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinItems-4            	 8347078	       144.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidMinLength-4                	51986899	        23.05 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundMinLength-4           	10005403	       119.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorMinLength-4            	 4193406	       286.0 ns/op	      32 B/op	       2 allocs/op
BenchmarkGoValidNumeric-4                  	157363160	         7.613 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundNumeric-4             	12090448	        98.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorNumeric-4              	 9458725	       127.8 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateNumeric-4           	   71224	     16763 ns/op	   15574 B/op	      78 allocs/op
BenchmarkGoValidRequired-4                 	320896393	         3.738 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundRequired-4            	 8045037	       149.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorRequired-4             	637254807	         1.869 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateRequired-4          	   75968	     15753 ns/op	   15488 B/op	      73 allocs/op
BenchmarkGoValidURL-4                      	20803777	        57.67 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundURL-4                 	 2423863	       493.2 ns/op	     144 B/op	       1 allocs/op
BenchmarkGoValidatorURL-4                  	  102288	     11741 ns/op	     145 B/op	       1 allocs/op
BenchmarkGookitValidateURL-4               	   70700	     16956 ns/op	   15681 B/op	      77 allocs/op
BenchmarkGoValidUUID-4                     	24674838	        48.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoPlaygroundUUID-4                	 2678340	       448.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkGoValidatorUUID-4                 	 3585114	       342.4 ns/op	       0 B/op	       0 allocs/op
BenchmarkGookitValidateUUID-4              	   70704	     16803 ns/op	   15542 B/op	      76 allocs/op
```

## Performance Comparison

| Validator | govalid | go-playground | vs go-playground | asaskevich/govalidator | vs asaskevich | gookit/validate | vs gookit |
|-----------|---------|---------------|------------------|----------------------|---------------|----------------|----------|
| LTE | 2.805 / 0 allocs | 109.5 / 0 allocs | **39.0x** | N/A | N/A | N/A | N/A |
| Enum | 4.361 / 0 allocs | N/A | N/A | N/A | N/A | N/A | N/A |
| Email | 55.28 / 0 allocs | 1115 / 89 B / 5 allocs | **20.2x** | 931.7 / 0 allocs | **16.9x** | 17352 / 15869 B / 76 allocs | **313.9x** |
| GTE | 2.501 / 0 allocs | 106.0 / 0 allocs | **42.4x** | N/A | N/A | N/A | N/A |
| MinLength | 23.05 / 0 allocs | 119.1 / 0 allocs | **5.2x** | 286.0 / 32 B / 2 allocs | **12.4x** | N/A | N/A |
| UUID | 48.54 / 0 allocs | 448.6 / 0 allocs | **9.2x** | 342.4 / 0 allocs | **7.1x** | 16803 / 15542 B / 76 allocs | **346.2x** |
| MaxItems | 6.563 / 0 allocs | 144.7 / 0 allocs | **22.0x** | N/A | N/A | N/A | N/A |
| MaxLength | 33.42 / 0 allocs | 139.3 / 0 allocs | **4.2x** | 286.0 / 32 B / 2 allocs | **8.6x** | 16496 / 15648 B / 81 allocs | **493.6x** |
| LT | 2.494 / 0 allocs | 105.6 / 0 allocs | **42.3x** | N/A | N/A | N/A | N/A |
| MinItems | 5.920 / 0 allocs | 144.1 / 0 allocs | **24.3x** | N/A | N/A | N/A | N/A |
| Alpha | 9.655 / 0 allocs | 467.4 / 0 allocs | **48.4x** | 108.3 / 0 allocs | **11.2x** | 21548 / 16937 B / 101 allocs | **2231.8x** |
| Required | 3.738 / 0 allocs | 149.6 / 0 allocs | **40.0x** | 1.869 / 0 allocs | **0.5x** | 15753 / 15488 B / 73 allocs | **4214.3x** |
| Length | 7.478 / 0 allocs | 124.4 / 0 allocs | **16.6x** | 249.6 / 32 B / 2 allocs | **33.4x** | 16363 / 15616 B / 79 allocs | **2188.2x** |
| URL | 57.67 / 0 allocs | 493.2 / 144 B / 1 allocs | **8.6x** | 11741 / 145 B / 1 allocs | **203.6x** | 16956 / 15681 B / 77 allocs | **294.0x** |
| Numeric | 7.613 / 0 allocs | 98.82 / 0 allocs | **13.0x** | 127.8 / 0 allocs | **16.8x** | 16763 / 15574 B / 78 allocs | **2201.9x** |
| GT | 2.492 / 0 allocs | 106.6 / 0 allocs | **42.8x** | 89.60 / 0 allocs | **36.0x** | N/A | N/A |

## CEL Expression Validation (govalid Exclusive)

| CEL Validator | govalid (ns/op) | Allocations |
|---------------|-----------------|-------------|
| CELConcurrent | 1.860 | 0 allocs |
| CELMultipleExpressions | 4.058 | 0 allocs |
| CELBasic | 4.053 | 0 allocs |
| CELCrossField | 3.431 | 0 allocs |
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
