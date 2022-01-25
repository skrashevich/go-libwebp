// Code generated by 'ccgo -pkgname main -trace-translation-units -o ../libwebp.go cdb.json src/.libs/libwebp.a', DO NOT EDIT.

package main

var CAPI = map[string]struct{}{
	"InitGammaTables":                      {},
	"InitGammaTablesS":                     {},
	"InitGetCoeffs":                        {},
	"QuantizeLevels":                       {},
	"VP8AccumulateSSE":                     {},
	"VP8AdjustFilterStrength":              {},
	"VP8ApplyNearLossless":                 {},
	"VP8BitReaderSetBuffer":                {},
	"VP8BitWriterAppend":                   {},
	"VP8BitWriterFinish":                   {},
	"VP8BitWriterInit":                     {},
	"VP8BitWriterWipeOut":                  {},
	"VP8CalculateLevelCosts":               {},
	"VP8Cat3":                              {},
	"VP8Cat4":                              {},
	"VP8Cat5":                              {},
	"VP8Cat6":                              {},
	"VP8CheckSignature":                    {},
	"VP8Clear":                             {},
	"VP8CodeIntraModes":                    {},
	"VP8CoeffsProba0":                      {},
	"VP8CoeffsUpdateProba":                 {},
	"VP8CollectHistogram":                  {},
	"VP8Copy16x8":                          {},
	"VP8Copy4x4":                           {},
	"VP8Decimate":                          {},
	"VP8Decode":                            {},
	"VP8DecodeMB":                          {},
	"VP8DecompressAlphaRows":               {},
	"VP8DefaultProbas":                     {},
	"VP8Delete":                            {},
	"VP8DitherCombine8x8":                  {},
	"VP8DspInit":                           {},
	"VP8DspInitMIPS32":                     {},
	"VP8DspInitMIPSdspR2":                  {},
	"VP8DspInitMSA":                        {},
	"VP8DspInitNEON":                       {},
	"VP8DspInitSSE2":                       {},
	"VP8DspInitSSE41":                      {},
	"VP8DspScan":                           {},
	"VP8EmitTokens":                        {},
	"VP8EncAnalyze":                        {},
	"VP8EncBands":                          {},
	"VP8EncDeleteAlpha":                    {},
	"VP8EncDspCostInit":                    {},
	"VP8EncDspCostInitMIPS32":              {},
	"VP8EncDspCostInitMIPSdspR2":           {},
	"VP8EncDspCostInitNEON":                {},
	"VP8EncDspCostInitSSE2":                {},
	"VP8EncDspInit":                        {},
	"VP8EncDspInitMIPS32":                  {},
	"VP8EncDspInitMIPSdspR2":               {},
	"VP8EncDspInitMSA":                     {},
	"VP8EncDspInitNEON":                    {},
	"VP8EncDspInitSSE2":                    {},
	"VP8EncDspInitSSE41":                   {},
	"VP8EncFinishAlpha":                    {},
	"VP8EncFreeBitWriters":                 {},
	"VP8EncInitAlpha":                      {},
	"VP8EncLoop":                           {},
	"VP8EncPredChroma8":                    {},
	"VP8EncPredLuma16":                     {},
	"VP8EncPredLuma4":                      {},
	"VP8EncQuantize2Blocks":                {},
	"VP8EncQuantizeBlock":                  {},
	"VP8EncQuantizeBlockWHT":               {},
	"VP8EncStartAlpha":                     {},
	"VP8EncTokenLoop":                      {},
	"VP8EncWrite":                          {},
	"VP8EnterCritical":                     {},
	"VP8EntropyCost":                       {},
	"VP8EstimateTokenSize":                 {},
	"VP8ExitCritical":                      {},
	"VP8FTransform":                        {},
	"VP8FTransform2":                       {},
	"VP8FTransformWHT":                     {},
	"VP8FilterStrengthFromDelta":           {},
	"VP8FiltersInit":                       {},
	"VP8FiltersInitMIPSdspR2":              {},
	"VP8FiltersInitMSA":                    {},
	"VP8FiltersInitNEON":                   {},
	"VP8FiltersInitSSE2":                   {},
	"VP8FixedCostsI16":                     {},
	"VP8FixedCostsI4":                      {},
	"VP8FixedCostsUV":                      {},
	"VP8GetCPUInfo":                        {},
	"VP8GetCostLuma16":                     {},
	"VP8GetCostLuma4":                      {},
	"VP8GetCostUV":                         {},
	"VP8GetHeaders":                        {},
	"VP8GetInfo":                           {},
	"VP8GetResidualCost":                   {},
	"VP8GetSignedValue":                    {},
	"VP8GetThreadMethod":                   {},
	"VP8GetValue":                          {},
	"VP8HFilter16":                         {},
	"VP8HFilter16i":                        {},
	"VP8HFilter8":                          {},
	"VP8HFilter8i":                         {},
	"VP8I16ModeOffsets":                    {},
	"VP8I4ModeOffsets":                     {},
	"VP8ITransform":                        {},
	"VP8InitBitReader":                     {},
	"VP8InitClipTables":                    {},
	"VP8InitDithering":                     {},
	"VP8InitFilter":                        {},
	"VP8InitFrame":                         {},
	"VP8InitIoInternal":                    {},
	"VP8InitRandom":                        {},
	"VP8InitResidual":                      {},
	"VP8InitScanline":                      {},
	"VP8IteratorBytesToNz":                 {},
	"VP8IteratorExport":                    {},
	"VP8IteratorImport":                    {},
	"VP8IteratorInit":                      {},
	"VP8IteratorIsDone":                    {},
	"VP8IteratorNext":                      {},
	"VP8IteratorNzToBytes":                 {},
	"VP8IteratorProgress":                  {},
	"VP8IteratorReset":                     {},
	"VP8IteratorRotateI4":                  {},
	"VP8IteratorSaveBoundary":              {},
	"VP8IteratorSetCountDown":              {},
	"VP8IteratorSetRow":                    {},
	"VP8IteratorStartI4":                   {},
	"VP8LAddGreenToBlueAndRed":             {},
	"VP8LAddGreenToBlueAndRed_C":           {},
	"VP8LAddVector":                        {},
	"VP8LAddVectorEq":                      {},
	"VP8LAllocateHistogram":                {},
	"VP8LAllocateHistogramSet":             {},
	"VP8LBackwardReferencesTraceBackwards": {},
	"VP8LBackwardRefsClear":                {},
	"VP8LBackwardRefsCursorAdd":            {},
	"VP8LBackwardRefsInit":                 {},
	"VP8LBitEntropyInit":                   {},
	"VP8LBitReaderSetBuffer":               {},
	"VP8LBitWriterClone":                   {},
	"VP8LBitWriterFinish":                  {},
	"VP8LBitWriterInit":                    {},
	"VP8LBitWriterReset":                   {},
	"VP8LBitWriterSwap":                    {},
	"VP8LBitWriterWipeOut":                 {},
	"VP8LBitsEntropy":                      {},
	"VP8LBitsEntropyUnrefined":             {},
	"VP8LBuildHuffmanTable":                {},
	"VP8LBundleColorMap":                   {},
	"VP8LBundleColorMap_C":                 {},
	"VP8LCheckSignature":                   {},
	"VP8LClear":                            {},
	"VP8LClearBackwardRefs":                {},
	"VP8LCollectColorBlueTransforms":       {},
	"VP8LCollectColorBlueTransforms_C":     {},
	"VP8LCollectColorRedTransforms":        {},
	"VP8LCollectColorRedTransforms_C":      {},
	"VP8LColorCacheClear":                  {},
	"VP8LColorCacheCopy":                   {},
	"VP8LColorCacheInit":                   {},
	"VP8LColorIndexInverseTransformAlpha":  {},
	"VP8LColorSpaceTransform":              {},
	"VP8LCombinedShannonEntropy":           {},
	"VP8LConvertBGRAToBGR":                 {},
	"VP8LConvertBGRAToBGR_C":               {},
	"VP8LConvertBGRAToRGB":                 {},
	"VP8LConvertBGRAToRGB565":              {},
	"VP8LConvertBGRAToRGB565_C":            {},
	"VP8LConvertBGRAToRGBA":                {},
	"VP8LConvertBGRAToRGBA4444":            {},
	"VP8LConvertBGRAToRGBA4444_C":          {},
	"VP8LConvertBGRAToRGBA_C":              {},
	"VP8LConvertBGRAToRGB_C":               {},
	"VP8LConvertFromBGRA":                  {},
	"VP8LCreateCompressedHuffmanTree":      {},
	"VP8LCreateHuffmanTree":                {},
	"VP8LDecodeAlphaHeader":                {},
	"VP8LDecodeAlphaImageStream":           {},
	"VP8LDecodeHeader":                     {},
	"VP8LDecodeImage":                      {},
	"VP8LDelete":                           {},
	"VP8LDistanceToPlaneCode":              {},
	"VP8LDoFillBitWindow":                  {},
	"VP8LDspInit":                          {},
	"VP8LDspInitMIPSdspR2":                 {},
	"VP8LDspInitMSA":                       {},
	"VP8LDspInitNEON":                      {},
	"VP8LDspInitSSE2":                      {},
	"VP8LDspInitSSE41":                     {},
	"VP8LEncDspInit":                       {},
	"VP8LEncDspInitMIPS32":                 {},
	"VP8LEncDspInitMIPSdspR2":              {},
	"VP8LEncDspInitMSA":                    {},
	"VP8LEncDspInitNEON":                   {},
	"VP8LEncDspInitSSE2":                   {},
	"VP8LEncDspInitSSE41":                  {},
	"VP8LEncodeImage":                      {},
	"VP8LEncodeStream":                     {},
	"VP8LExtraCost":                        {},
	"VP8LExtraCostCombined":                {},
	"VP8LFastLog2Slow":                     {},
	"VP8LFastSLog2Slow":                    {},
	"VP8LFreeHistogram":                    {},
	"VP8LFreeHistogramSet":                 {},
	"VP8LGetBackwardReferences":            {},
	"VP8LGetCombinedEntropyUnrefined":      {},
	"VP8LGetEntropyUnrefined":              {},
	"VP8LGetHistoImageSymbols":             {},
	"VP8LGetHistogramSize":                 {},
	"VP8LGetInfo":                          {},
	"VP8LHashChainClear":                   {},
	"VP8LHashChainFill":                    {},
	"VP8LHashChainInit":                    {},
	"VP8LHistogramAdd":                     {},
	"VP8LHistogramAddSinglePixOrCopy":      {},
	"VP8LHistogramCreate":                  {},
	"VP8LHistogramEstimateBits":            {},
	"VP8LHistogramInit":                    {},
	"VP8LHistogramSetClear":                {},
	"VP8LHistogramStoreRefs":               {},
	"VP8LHtreeGroupsFree":                  {},
	"VP8LHtreeGroupsNew":                   {},
	"VP8LInitBitReader":                    {},
	"VP8LInverseTransform":                 {},
	"VP8LMapColor32b":                      {},
	"VP8LMapColor8b":                       {},
	"VP8LNew":                              {},
	"VP8LPredictor0_C":                     {},
	"VP8LPredictor10_C":                    {},
	"VP8LPredictor11_C":                    {},
	"VP8LPredictor12_C":                    {},
	"VP8LPredictor13_C":                    {},
	"VP8LPredictor1_C":                     {},
	"VP8LPredictor2_C":                     {},
	"VP8LPredictor3_C":                     {},
	"VP8LPredictor4_C":                     {},
	"VP8LPredictor5_C":                     {},
	"VP8LPredictor6_C":                     {},
	"VP8LPredictor7_C":                     {},
	"VP8LPredictor8_C":                     {},
	"VP8LPredictor9_C":                     {},
	"VP8LPredictors":                       {},
	"VP8LPredictorsAdd":                    {},
	"VP8LPredictorsAdd_C":                  {},
	"VP8LPredictorsSub":                    {},
	"VP8LPredictorsSub_C":                  {},
	"VP8LPutBitsFlushBits":                 {},
	"VP8LPutBitsInternal":                  {},
	"VP8LReadBits":                         {},
	"VP8LRefsCursorInit":                   {},
	"VP8LRefsCursorNextBlock":              {},
	"VP8LResidualImage":                    {},
	"VP8LSubtractGreenFromBlueAndRed":      {},
	"VP8LSubtractGreenFromBlueAndRed_C":    {},
	"VP8LTransformColor":                   {},
	"VP8LTransformColorInverse":            {},
	"VP8LTransformColorInverse_C":          {},
	"VP8LTransformColor_C":                 {},
	"VP8LVectorMismatch":                   {},
	"VP8LevelCodes":                        {},
	"VP8LevelFixedCosts":                   {},
	"VP8LoadFinalBytes":                    {},
	"VP8MakeChroma8Preds":                  {},
	"VP8MakeIntra4Preds":                   {},
	"VP8MakeLuma16Preds":                   {},
	"VP8Mean16x4":                          {},
	"VP8New":                               {},
	"VP8ParseIntraModeRow":                 {},
	"VP8ParseProba":                        {},
	"VP8ParseQuant":                        {},
	"VP8PredChroma8":                       {},
	"VP8PredLuma16":                        {},
	"VP8PredLuma4":                         {},
	"VP8ProcessRow":                        {},
	"VP8PutBit":                            {},
	"VP8PutBitUniform":                     {},
	"VP8PutBits":                           {},
	"VP8PutSignedBits":                     {},
	"VP8RecordCoeffTokens":                 {},
	"VP8RecordCoeffs":                      {},
	"VP8RemapBitReader":                    {},
	"VP8ResetProba":                        {},
	"VP8SSE16x16":                          {},
	"VP8SSE16x8":                           {},
	"VP8SSE4x4":                            {},
	"VP8SSE8x8":                            {},
	"VP8SSIMDspInit":                       {},
	"VP8SSIMDspInitSSE2":                   {},
	"VP8SSIMFromStats":                     {},
	"VP8SSIMFromStatsClipped":              {},
	"VP8SSIMGet":                           {},
	"VP8SSIMGetClipped":                    {},
	"VP8Scan":                              {},
	"VP8SetError":                          {},
	"VP8SetHistogramData":                  {},
	"VP8SetIntra16Mode":                    {},
	"VP8SetIntra4Mode":                     {},
	"VP8SetIntraUVMode":                    {},
	"VP8SetResidualCoeffs":                 {},
	"VP8SetSegment":                        {},
	"VP8SetSegmentParams":                  {},
	"VP8SetSkip":                           {},
	"VP8SimpleHFilter16":                   {},
	"VP8SimpleHFilter16i":                  {},
	"VP8SimpleVFilter16":                   {},
	"VP8SimpleVFilter16i":                  {},
	"VP8Status":                            {},
	"VP8StatusMessage":                     {},
	"VP8StoreFilterStats":                  {},
	"VP8TBufferClear":                      {},
	"VP8TBufferInit":                       {},
	"VP8TDisto16x16":                       {},
	"VP8TDisto4x4":                         {},
	"VP8Transform":                         {},
	"VP8TransformAC3":                      {},
	"VP8TransformDC":                       {},
	"VP8TransformDCUV":                     {},
	"VP8TransformUV":                       {},
	"VP8TransformWHT":                      {},
	"VP8UVModeOffsets":                     {},
	"VP8VFilter16":                         {},
	"VP8VFilter16i":                        {},
	"VP8VFilter8":                          {},
	"VP8VFilter8i":                         {},
	"VP8WriteProbas":                       {},
	"VP8kabs0":                             {},
	"VP8kclip1":                            {},
	"VP8ksclip1":                           {},
	"VP8ksclip2":                           {},
	"WebPAllocateDecBuffer":                {},
	"WebPAlphaReplace":                     {},
	"WebPApplyAlphaMultiply":               {},
	"WebPApplyAlphaMultiply4444":           {},
	"WebPAvoidSlowMemory":                  {},
	"WebPBlendAlpha":                       {},
	"WebPCheckCropDimensions":              {},
	"WebPCleanupTransparentArea":           {},
	"WebPConfigInitInternal":               {},
	"WebPConfigLosslessPreset":             {},
	"WebPConvertARGBToUV":                  {},
	"WebPConvertARGBToUV_C":                {},
	"WebPConvertARGBToY":                   {},
	"WebPConvertBGR24ToY":                  {},
	"WebPConvertRGB24ToY":                  {},
	"WebPConvertRGBA32ToUV":                {},
	"WebPConvertRGBA32ToUV_C":              {},
	"WebPCopyDecBuffer":                    {},
	"WebPCopyDecBufferPixels":              {},
	"WebPCopyPixels":                       {},
	"WebPCopyPlane":                        {},
	"WebPDeallocateAlphaMemory":            {},
	"WebPDecode":                           {},
	"WebPDecodeARGB":                       {},
	"WebPDecodeARGBInto":                   {},
	"WebPDecodeBGR":                        {},
	"WebPDecodeBGRA":                       {},
	"WebPDecodeBGRAInto":                   {},
	"WebPDecodeBGRInto":                    {},
	"WebPDecodeRGB":                        {},
	"WebPDecodeRGBA":                       {},
	"WebPDecodeRGBAInto":                   {},
	"WebPDecodeRGBInto":                    {},
	"WebPDecodeYUV":                        {},
	"WebPDecodeYUVInto":                    {},
	"WebPDequantizeLevels":                 {},
	"WebPDispatchAlpha":                    {},
	"WebPDispatchAlphaToGreen":             {},
	"WebPEncode":                           {},
	"WebPEncodeBGR":                        {},
	"WebPEncodeBGRA":                       {},
	"WebPEncodeLosslessBGR":                {},
	"WebPEncodeLosslessBGRA":               {},
	"WebPEncodeLosslessRGB":                {},
	"WebPEncodeLosslessRGBA":               {},
	"WebPEncodeRGB":                        {},
	"WebPEncodeRGBA":                       {},
	"WebPEncodingSetError":                 {},
	"WebPEstimateBestFilter":               {},
	"WebPExtractAlpha":                     {},
	"WebPExtractGreen":                     {},
	"WebPFilters":                          {},
	"WebPFlipBuffer":                       {},
	"WebPFree":                             {},
	"WebPFreeDecBuffer":                    {},
	"WebPGetColorPalette":                  {},
	"WebPGetDecoderVersion":                {},
	"WebPGetEncoderVersion":                {},
	"WebPGetFeaturesInternal":              {},
	"WebPGetInfo":                          {},
	"WebPGetLinePairConverter":             {},
	"WebPGetWorkerInterface":               {},
	"WebPGrabDecBuffer":                    {},
	"WebPHasAlpha32b":                      {},
	"WebPHasAlpha8b":                       {},
	"WebPIAppend":                          {},
	"WebPIDecGetRGB":                       {},
	"WebPIDecGetYUVA":                      {},
	"WebPIDecode":                          {},
	"WebPIDecodedArea":                     {},
	"WebPIDelete":                          {},
	"WebPINewDecoder":                      {},
	"WebPINewRGB":                          {},
	"WebPINewYUV":                          {},
	"WebPINewYUVA":                         {},
	"WebPISetIOHooks":                      {},
	"WebPIUpdate":                          {},
	"WebPInitAlphaProcessing":              {},
	"WebPInitAlphaProcessingMIPSdspR2":     {},
	"WebPInitAlphaProcessingNEON":          {},
	"WebPInitAlphaProcessingSSE2":          {},
	"WebPInitAlphaProcessingSSE41":         {},
	"WebPInitConvertARGBToYUV":             {},
	"WebPInitConvertARGBToYUVNEON":         {},
	"WebPInitConvertARGBToYUVSSE2":         {},
	"WebPInitConvertARGBToYUVSSE41":        {},
	"WebPInitCustomIo":                     {},
	"WebPInitDecBufferInternal":            {},
	"WebPInitDecoderConfigInternal":        {},
	"WebPInitSamplers":                     {},
	"WebPInitSamplersMIPS32":               {},
	"WebPInitSamplersMIPSdspR2":            {},
	"WebPInitSamplersSSE2":                 {},
	"WebPInitSamplersSSE41":                {},
	"WebPInitSharpYUVNEON":                 {},
	"WebPInitSharpYUVSSE2":                 {},
	"WebPInitUpsamplers":                   {},
	"WebPInitUpsamplersMIPSdspR2":          {},
	"WebPInitUpsamplersMSA":                {},
	"WebPInitUpsamplersNEON":               {},
	"WebPInitUpsamplersSSE2":               {},
	"WebPInitUpsamplersSSE41":              {},
	"WebPInitYUV444Converters":             {},
	"WebPInitYUV444ConvertersMIPSdspR2":    {},
	"WebPInitYUV444ConvertersSSE2":         {},
	"WebPInitYUV444ConvertersSSE41":        {},
	"WebPIoInitFromOptions":                {},
	"WebPMalloc":                           {},
	"WebPMemoryWrite":                      {},
	"WebPMemoryWriterClear":                {},
	"WebPMemoryWriterInit":                 {},
	"WebPMultARGBRow":                      {},
	"WebPMultARGBRow_C":                    {},
	"WebPMultARGBRows":                     {},
	"WebPMultRow":                          {},
	"WebPMultRow_C":                        {},
	"WebPMultRows":                         {},
	"WebPPackRGB":                          {},
	"WebPParseHeaders":                     {},
	"WebPPictureARGBToYUVA":                {},
	"WebPPictureARGBToYUVADithered":        {},
	"WebPPictureAlloc":                     {},
	"WebPPictureAllocARGB":                 {},
	"WebPPictureAllocYUVA":                 {},
	"WebPPictureCopy":                      {},
	"WebPPictureCrop":                      {},
	"WebPPictureDistortion":                {},
	"WebPPictureFree":                      {},
	"WebPPictureHasTransparency":           {},
	"WebPPictureImportBGR":                 {},
	"WebPPictureImportBGRA":                {},
	"WebPPictureImportBGRX":                {},
	"WebPPictureImportRGB":                 {},
	"WebPPictureImportRGBA":                {},
	"WebPPictureImportRGBX":                {},
	"WebPPictureInitInternal":              {},
	"WebPPictureIsView":                    {},
	"WebPPictureRescale":                   {},
	"WebPPictureResetBuffers":              {},
	"WebPPictureSharpARGBToYUVA":           {},
	"WebPPictureSmartARGBToYUVA":           {},
	"WebPPictureView":                      {},
	"WebPPictureYUVAToARGB":                {},
	"WebPPlaneDistortion":                  {},
	"WebPReplaceTransparentPixels":         {},
	"WebPReportProgress":                   {},
	"WebPRescaleNeededLines":               {},
	"WebPRescalerDspInit":                  {},
	"WebPRescalerDspInitMIPS32":            {},
	"WebPRescalerDspInitMIPSdspR2":         {},
	"WebPRescalerDspInitMSA":               {},
	"WebPRescalerDspInitNEON":              {},
	"WebPRescalerDspInitSSE2":              {},
	"WebPRescalerExport":                   {},
	"WebPRescalerExportRow":                {},
	"WebPRescalerExportRowExpand":          {},
	"WebPRescalerExportRowExpand_C":        {},
	"WebPRescalerExportRowShrink":          {},
	"WebPRescalerExportRowShrink_C":        {},
	"WebPRescalerGetScaledDimensions":      {},
	"WebPRescalerImport":                   {},
	"WebPRescalerImportRow":                {},
	"WebPRescalerImportRowExpand":          {},
	"WebPRescalerImportRowExpand_C":        {},
	"WebPRescalerImportRowShrink":          {},
	"WebPRescalerImportRowShrink_C":        {},
	"WebPRescalerInit":                     {},
	"WebPResetDecParams":                   {},
	"WebPSafeCalloc":                       {},
	"WebPSafeFree":                         {},
	"WebPSafeMalloc":                       {},
	"WebPSamplerProcessPlane":              {},
	"WebPSamplers":                         {},
	"WebPSetWorkerInterface":               {},
	"WebPSharpYUVFilterRow":                {},
	"WebPSharpYUVUpdateRGB":                {},
	"WebPSharpYUVUpdateY":                  {},
	"WebPUnfilters":                        {},
	"WebPUpsamplers":                       {},
	"WebPValidateConfig":                   {},
	"WebPYUV444Converters":                 {},
	"WebPYuv444ToArgb_C":                   {},
	"WebPYuv444ToBgr_C":                    {},
	"WebPYuv444ToBgra_C":                   {},
	"WebPYuv444ToRgb565_C":                 {},
	"WebPYuv444ToRgb_C":                    {},
	"WebPYuv444ToRgba4444_C":               {},
	"WebPYuv444ToRgba_C":                   {},
	"kLog2Table":                           {},
	"kPrefixEncodeCode":                    {},
	"kPrefixEncodeExtraBitsValue":          {},
	"kSLog2Table":                          {},
	"kVP8Log2Range":                        {},
	"kVP8NewRange":                         {},
}
